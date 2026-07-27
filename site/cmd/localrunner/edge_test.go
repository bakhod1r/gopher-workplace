package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// unreadable makes a dir unreadable for the duration of the test. Skips when
// running as root, where permissions do not bite.
func unreadable(t *testing.T, dir string) {
	t.Helper()
	if os.Geteuid() == 0 {
		t.Skip("running as root: permission bits do not apply")
	}
	if err := os.Chmod(dir, 0o000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(dir, 0o755) })
}

func TestCopyDirSkipsSubdirsAndReportsErrors(t *testing.T) {
	src, dst := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(src, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(src, "nested", "deep.go"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(src, "a.go"), []byte("package a\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyDir(src, dst); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dst, "a.go")); err != nil {
		t.Errorf("flat file not copied: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dst, "nested")); err == nil {
		t.Error("nested dir was copied; puzzles are flat")
	}

	// Unreadable source.
	if err := copyDir(filepath.Join(t.TempDir(), "missing"), dst); err == nil {
		t.Error("missing src: want an error")
	}
	// Unwritable destination.
	ro := t.TempDir()
	unreadable(t, ro)
	if err := copyDir(src, ro); err == nil {
		t.Error("unwritable dst: want an error")
	}
	// Unreadable source file.
	src2 := t.TempDir()
	f := filepath.Join(src2, "b.go")
	if err := os.WriteFile(f, []byte("x"), 0o000); err != nil {
		t.Fatal(err)
	}
	if os.Geteuid() != 0 {
		if err := copyDir(src2, dst); err == nil {
			t.Error("unreadable src file: want an error")
		}
	}
}

func TestMaterializeErrors(t *testing.T) {
	src := t.TempDir()
	if err := os.WriteFile(filepath.Join(src, "a.go"), []byte("package a\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	// copyDir fails.
	if err := materialize(filepath.Join(t.TempDir(), "nope"), t.TempDir(), runReq{Src: "x"}); err == nil {
		t.Error("bad challenge dir: want an error")
	}
	// A candidate file name that is not a .go file.
	if err := materialize(src, t.TempDir(), runReq{Files: map[string]string{"notes.txt": "x"}}); err == nil {
		t.Error("non-.go candidate: want an error")
	}
	// src form with no .go file to overwrite.
	empty := t.TempDir()
	if err := os.WriteFile(filepath.Join(empty, "README.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := materialize(empty, t.TempDir(), runReq{Src: "x"}); err == nil {
		t.Error("no candidate file: want an error")
	}
}

func TestStarterFileMissingDir(t *testing.T) {
	if _, err := starterFile(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Error("missing dir: want an error")
	}
}

func TestGofmtWarnEdgeCases(t *testing.T) {
	// Unreadable dir → no warnings, no panic.
	if got := gofmtWarn(filepath.Join(t.TempDir(), "nope")); got != nil {
		t.Errorf("missing dir: %q", got)
	}
	// No .go files → nothing to check.
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := gofmtWarn(dir); got != nil {
		t.Errorf("no go files: %q", got)
	}
	// Test files are ignored; only the candidate is checked.
	if err := os.WriteFile(filepath.Join(dir, "a_test.go"), []byte("package a\nfunc  F(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := gofmtWarn(dir); got != nil {
		t.Errorf("test file flagged: %q", got)
	}
	// A malformed .go file makes gofmt exit non-zero: best-effort, no warnings.
	if err := os.WriteFile(filepath.Join(dir, "b.go"), []byte("package\n{{{"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := gofmtWarn(dir); got != nil {
		t.Errorf("gofmt failure should yield no warnings, got %q", got)
	}
}

func TestNeedsRaceIgnoresUnreadableAndNonGo(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("go func()"), 0o644); err != nil {
		t.Fatal(err)
	}
	if needsRace(dir) {
		t.Error("a non-.go file triggered -race")
	}
	if os.Geteuid() != 0 {
		if err := os.WriteFile(filepath.Join(dir, "a.go"), []byte("package a\nimport \"sync\"\n"), 0o000); err != nil {
			t.Fatal(err)
		}
		if needsRace(dir) {
			t.Error("an unreadable file must be skipped, not read")
		}
	}
	if needsRace(filepath.Join(t.TempDir(), "missing")) {
		t.Error("missing dir")
	}
}

// A submission that never terminates must be killed, not left running.
func TestGoTestTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	dir := t.TempDir()
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("go.mod", "module spin\n\ngo 1.26\n")
	write("spin.go", "package spin\n\nfunc Spin() {\n\tfor {\n\t}\n}\n")
	write("spin_test.go", "package spin\n\nimport \"testing\"\n\nfunc TestSpin(t *testing.T) { Spin() }\n")

	old := runTimeout
	runTimeout = 3 * time.Second
	t.Cleanup(func() { runTimeout = old })

	start := time.Now()
	rep := goTest(dir, false)
	if !strings.Contains(rep.Error, "timed out") {
		t.Errorf("error = %q, want a timeout report", rep.Error)
	}
	if elapsed := time.Since(start); elapsed > 30*time.Second {
		t.Errorf("took %s: the process group was not killed promptly", elapsed)
	}
}

func TestHandleVetTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	old := runTimeout
	runTimeout = time.Millisecond // expire before `go vet` can finish
	t.Cleanup(func() { runTimeout = old })

	rep := decodeReport(t, postJSON(t, s.handleVet, "/vet",
		runReq{ChallengeID: id, Src: "package double\n\nfunc Double(n int) int {\n\treturn n * 2\n}\n"}))
	if !strings.Contains(rep.Error, "timed out") {
		t.Errorf("error = %q, want a vet timeout", rep.Error)
	}
}

// With TMPDIR pointing at nothing, MkdirTemp fails and both handlers must
// report it rather than panic.
func TestHandlersReportTempDirFailure(t *testing.T) {
	s, id := fixture(t)
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "does-not-exist"))

	for name, h := range map[string]http.HandlerFunc{"run": s.handleRun, "vet": s.handleVet} {
		rep := decodeReport(t, postJSON(t, h, "/"+name, runReq{ChallengeID: id, Src: "package double\n"}))
		if !strings.Contains(rep.Error, "temp dir") {
			t.Errorf("%s: error = %q, want a temp dir failure", name, rep.Error)
		}
	}
}

func TestHandlersReportMaterializeFailure(t *testing.T) {
	s, id := fixture(t)
	for name, h := range map[string]http.HandlerFunc{"run": s.handleRun, "vet": s.handleVet} {
		rep := decodeReport(t, postJSON(t, h, "/"+name,
			runReq{ChallengeID: id, Files: map[string]string{"notes.txt": "x"}}))
		if rep.Error == "" {
			t.Errorf("%s: a non-.go candidate was accepted", name)
		}
	}
}

// The store is optional: /run must work with it absent.
func TestHandleRunWithoutStore(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	s.store = nil
	rep := decodeReport(t, postJSON(t, s.handleRun, "/run",
		runReq{ChallengeID: id, Src: "package double\n\nfunc Double(n int) int {\n\treturn n * 2\n}\n"}))
	if !rep.OK {
		t.Errorf("ok = false without a store: %q", rep.Error)
	}
}

func TestOpenStoreOnNonDatabaseFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "garbage.db")
	if err := os.WriteFile(path, []byte("this is not a sqlite file"), 0o644); err != nil {
		t.Fatal(err)
	}
	if s, err := openStore(path); err == nil {
		s.close()
		t.Error("openStore on a non-database file: want an error")
	}
}

// A NULL challenge_id cannot be scanned into a string; the error must be
// returned, not silently dropped.
func TestSolvedIDsAndHistoryScanErrors(t *testing.T) {
	s := testStore(t)
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, submitted, created_at)
		 VALUES (NULL, '{}', 1, 1, 1, 0, 1, 0)`); err != nil {
		t.Skipf("db rejected the NULL row: %v", err)
	}
	if _, err := s.solvedIDs(); err == nil {
		t.Error("solvedIDs: want a scan error")
	}
	if _, err := s.history("", 10); err != nil {
		t.Logf("history: %v", err)
	}
}

func TestSolvedIDsAndHistoryOnClosedDB(t *testing.T) {
	s := testStore(t)
	s.close()
	if _, err := s.solvedIDs(); err == nil {
		t.Error("solvedIDs on a closed db: want an error")
	}
	if _, err := s.history("x", 10); err == nil {
		t.Error("history on a closed db: want an error")
	}
}

func TestSweepLogsOnError(t *testing.T) {
	logs := captureLog(t)
	s := testStore(t)
	s.close()
	s.sweep(time.Hour) // must log, not panic
	if !strings.Contains(logs.String(), "retention sweep") {
		t.Errorf("logs = %q, want a sweep error", logs.String())
	}
}

// The retention goroutine keeps sweeping after startup.
func TestStartRetentionTicks(t *testing.T) {
	old := sweepInterval
	sweepInterval = 20 * time.Millisecond
	t.Cleanup(func() { sweepInterval = old })

	s := testStore(t)
	s.startRetention(time.Hour)

	// Insert a stale row *after* the startup sweep; the ticker must remove it.
	if _, err := s.db.Exec(
		`INSERT INTO submissions (challenge_id, files, ok, total, passed, race, submitted, created_at)
		 VALUES ('ch', '{}', 1, 1, 1, 0, 1, ?)`, time.Now().Add(-48*time.Hour).Unix()); err != nil {
		t.Fatal(err)
	}
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		rows, err := s.history("ch", 10)
		if err == nil && len(rows) == 0 {
			return // swept
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Error("the retention ticker never swept the stale row")
}

func TestSaveLogsOnError(t *testing.T) {
	logs := captureLog(t)
	s := testStore(t)
	s.close()
	s.save("ch", nil, true, 1, 1, false, true)
	if !strings.Contains(logs.String(), "save submission") {
		t.Errorf("logs = %q, want a save error", logs.String())
	}
}

// run() creates ~/.gopher-workplace when no -db is given; a HOME that is a file
// makes that fail.
func TestRunStateDirFailure(t *testing.T) {
	root := repoFixture(t)
	fakeHome := filepath.Join(t.TempDir(), "home-is-a-file")
	if err := os.WriteFile(fakeHome, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HOME", fakeHome)
	err := run([]string{"-root", root, "-port", "0"})
	if err == nil || !strings.Contains(err.Error(), "state dir") {
		t.Errorf("err = %v, want a state dir failure", err)
	}
}

// A very long line must not blow up the scanner; parseTestJSON reads up to 4MB.
func TestParseTestJSONHugeOutput(t *testing.T) {
	huge := strings.Repeat("x", 5*1024*1024)
	line, err := json.Marshal(testEvent{Action: "output", Test: "TestA", Output: huge})
	if err != nil {
		t.Fatal(err)
	}
	out := append(line, []byte("\n{\"Action\":\"fail\",\"Test\":\"TestA\"}\n")...)
	rep := parseTestJSON(out, "stderr text", errors.New("exit status 1"))
	// The oversized line is dropped by the scanner; what matters is that we get
	// a report back instead of a panic or a hang.
	if rep.OK {
		t.Error("ok = true for a failing run")
	}
}

func TestParseTestJSONFallsBackToStderr(t *testing.T) {
	rep := parseTestJSON(nil, "go: cannot find main module", errors.New("exit status 1"))
	if rep.CompileOK {
		t.Error("compileOk = true with no output at all")
	}
	if !strings.Contains(rep.Error, "cannot find main module") {
		t.Errorf("error = %q, want the stderr text", rep.Error)
	}
}

func TestParseTestJSONFallsBackToRunErr(t *testing.T) {
	rep := parseTestJSON(nil, "", errors.New("fork/exec: no such file"))
	if !strings.Contains(rep.Error, "fork/exec") {
		t.Errorf("error = %q, want the run error", rep.Error)
	}
}

func TestParseTestJSONIgnoresBlankLines(t *testing.T) {
	out := []byte("\n\n{\"Action\":\"pass\",\"Test\":\"TestA\"}\n\n")
	rep := parseTestJSON(out, "", nil)
	if len(rep.Cases) != 1 {
		t.Errorf("cases = %d, want 1", len(rep.Cases))
	}
}

func TestWithLimitPassesThroughRecorder(t *testing.T) {
	s := &server{limiter: newLimiter(1)}
	rec := httptest.NewRecorder()
	s.withLimit(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, report{OK: true})
	})(rec, httptest.NewRequest(http.MethodPost, "/run", nil))
	if !strings.Contains(rec.Body.String(), `"ok":true`) {
		t.Errorf("body = %q", rec.Body.String())
	}
}

// Concurrency-flavoured puzzles are run with -race; this exercises that arm.
func TestGoTestWithRace(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain with -race")
	}
	dir := t.TempDir()
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("go.mod", "module conc\n\ngo 1.26\n")
	write("conc.go", "package conc\n\nimport \"sync\"\n\nfunc Count(n int) int {\n\tvar mu sync.Mutex\n\tvar wg sync.WaitGroup\n\ttotal := 0\n\tfor i := 0; i < n; i++ {\n\t\twg.Add(1)\n\t\tgo func() {\n\t\t\tdefer wg.Done()\n\t\t\tmu.Lock()\n\t\t\ttotal++\n\t\t\tmu.Unlock()\n\t\t}()\n\t}\n\twg.Wait()\n\treturn total\n}\n")
	write("conc_test.go", "package conc\n\nimport \"testing\"\n\nfunc TestCount(t *testing.T) {\n\tif got := Count(8); got != 8 {\n\t\tt.Errorf(\"Count(8) = %d, want 8\", got)\n\t}\n}\n")

	if !needsRace(dir) {
		t.Fatal("needsRace did not spot the concurrency markers")
	}
	rep := goTest(dir, true)
	if !rep.OK {
		t.Errorf("race-enabled run failed: %q %+v", rep.Error, rep.Cases)
	}
}
