package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// fixture builds a repo root holding one real, solvable puzzle and returns a
// server wired to it. The puzzle is tiny so `go test` against it stays fast.
func fixture(t *testing.T) (*server, string) {
	t.Helper()
	root := t.TempDir()
	id := "junior/01-basics/01-vars/double"
	dir := filepath.Join(root, "challenges", filepath.FromSlash(id))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("go.mod", "module double\n\ngo 1.26\n")
	write("double.go", "package double\n\nfunc Double(n int) int {\n\treturn 0\n}\n")
	write("double_test.go", `package double

import "testing"

func TestDouble(t *testing.T) {
	if got := Double(2); got != 4 {
		t.Errorf("Double(2) = %d, want 4", got)
	}
}
`)
	s := &server{
		cfg:     config{root: root},
		store:   testStore(t),
		limiter: newLimiter(2),
	}
	return s, id
}

func postJSON(t *testing.T, h http.HandlerFunc, path string, body any) *httptest.ResponseRecorder {
	t.Helper()
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(string(b)))
	rec := httptest.NewRecorder()
	h(rec, req)
	return rec
}

func decodeReport(t *testing.T, rec *httptest.ResponseRecorder) report {
	t.Helper()
	var rep report
	if err := json.Unmarshal(rec.Body.Bytes(), &rep); err != nil {
		t.Fatalf("decode report: %v\nbody: %s", err, rec.Body.String())
	}
	return rep
}

func TestHandleRunGreenSubmitIsRecordedAsSolved(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	src := "package double\n\nfunc Double(n int) int {\n\treturn n * 2\n}\n"

	rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{ChallengeID: id, Src: src, Submit: true}))
	if !rep.OK || !rep.CompileOK {
		t.Fatalf("ok=%v compileOk=%v error=%q cases=%+v", rep.OK, rep.CompileOK, rep.Error, rep.Cases)
	}
	if len(rep.Warnings) != 0 {
		t.Errorf("warnings = %q, want none for gofmt-clean passing code", rep.Warnings)
	}
	solved, err := s.store.solvedIDs()
	if err != nil {
		t.Fatal(err)
	}
	if len(solved) != 1 || solved[0] != id {
		t.Errorf("solvedIDs = %q, want [%s]", solved, id)
	}
}

func TestHandleRunRedIsNotSolved(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	src := "package double\n\nfunc Double(n int) int {\n\treturn n\n}\n" // wrong

	rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{ChallengeID: id, Src: src, Submit: true}))
	if rep.OK {
		t.Error("ok = true for a wrong answer")
	}
	if !rep.CompileOK {
		t.Errorf("compileOk = false; it compiles fine: %q", rep.Error)
	}
	if len(rep.Cases) == 0 || rep.Cases[0].Got != "2" || rep.Cases[0].Want != "4" {
		t.Errorf("cases = %+v, want got/want 2/4", rep.Cases)
	}
	if solved, _ := s.store.solvedIDs(); len(solved) != 0 {
		t.Errorf("solvedIDs = %q, want none", solved)
	}
}

// A passing but unformatted Submit must be reported green yet NOT counted as
// solved — warnings block acceptance.
func TestHandleRunUnformattedSubmitWarnsAndIsNotSolved(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	src := "package double\nfunc Double(n int) int {   return n*2 }\n"

	rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{ChallengeID: id, Src: src, Submit: true}))
	if !rep.OK {
		t.Fatalf("ok = false: %q %+v", rep.Error, rep.Cases)
	}
	if len(rep.Warnings) == 0 {
		t.Fatal("no warnings for non-gofmt-clean source")
	}
	if !strings.Contains(rep.Warnings[0], "gofmt") {
		t.Errorf("warning = %q", rep.Warnings[0])
	}
	if solved, _ := s.store.solvedIDs(); len(solved) != 0 {
		t.Errorf("solvedIDs = %q; a warned submit must not count", solved)
	}
}

func TestHandleRunCompileError(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	rep := decodeReport(t, postJSON(t, s.handleRun, "/run",
		runReq{ChallengeID: id, Src: "package double\n\nfunc Double(n int) int { return nope }\n"}))
	if rep.CompileOK {
		t.Error("compileOk = true for undefined identifier")
	}
	if !strings.Contains(rep.Error, "nope") {
		t.Errorf("error = %q, want the compiler message", rep.Error)
	}
}

// The files form overwrites named candidate files instead of the starter.
func TestHandleRunFilesForm(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)
	rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{
		ChallengeID: id,
		Files:       map[string]string{"double.go": "package double\n\nfunc Double(n int) int {\n\treturn n + n\n}\n"},
		Submit:      true,
	}))
	if !rep.OK {
		t.Fatalf("ok = false: %q %+v", rep.Error, rep.Cases)
	}
}

func TestHandleRunBadRequests(t *testing.T) {
	s, id := fixture(t)

	// Not POST.
	rec := httptest.NewRecorder()
	s.handleRun(rec, httptest.NewRequest(http.MethodGet, "/run", nil))
	if rep := decodeReport(t, rec); !strings.Contains(rep.Error, "POST") {
		t.Errorf("GET /run error = %q", rep.Error)
	}

	// Malformed JSON.
	rec = httptest.NewRecorder()
	s.handleRun(rec, httptest.NewRequest(http.MethodPost, "/run", strings.NewReader("{oops")))
	if rep := decodeReport(t, rec); !strings.Contains(rep.Error, "bad body") {
		t.Errorf("bad json error = %q", rep.Error)
	}

	// Missing fields.
	if rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{Src: "x"})); !strings.Contains(rep.Error, "challengeId") {
		t.Errorf("missing id error = %q", rep.Error)
	}
	if rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{ChallengeID: id})); !strings.Contains(rep.Error, "files or src") {
		t.Errorf("missing src error = %q", rep.Error)
	}

	// Unknown challenge, and a traversal attempt.
	for _, bad := range []string{"junior/nope", "../../etc"} {
		rep := decodeReport(t, postJSON(t, s.handleRun, "/run", runReq{ChallengeID: bad, Src: "x"}))
		if rep.Error == "" {
			t.Errorf("challengeId %q accepted", bad)
		}
	}
}

func TestHandleVet(t *testing.T) {
	if testing.Short() {
		t.Skip("shells out to the Go toolchain")
	}
	s, id := fixture(t)

	clean := decodeReport(t, postJSON(t, s.handleVet, "/vet",
		runReq{ChallengeID: id, Src: "package double\n\nfunc Double(n int) int {\n\treturn n * 2\n}\n"}))
	if !clean.OK || !clean.CompileOK {
		t.Errorf("clean vet: ok=%v compileOk=%v error=%q", clean.OK, clean.CompileOK, clean.Error)
	}

	// Printf arg mismatch is a classic vet catch.
	bad := decodeReport(t, postJSON(t, s.handleVet, "/vet", runReq{
		ChallengeID: id,
		Src:         "package double\n\nimport \"fmt\"\n\nfunc Double(n int) int {\n\tfmt.Printf(\"%d %d\\n\", n)\n\treturn n * 2\n}\n",
	}))
	if bad.OK {
		t.Errorf("vet passed a Printf arg mismatch: %+v", bad)
	}
}

func TestHandleVetBadRequests(t *testing.T) {
	s, _ := fixture(t)
	if rep := decodeReport(t, postJSON(t, s.handleVet, "/vet", runReq{Src: "x"})); rep.Error == "" {
		t.Error("missing challengeId accepted")
	}
	if rep := decodeReport(t, postJSON(t, s.handleVet, "/vet", runReq{ChallengeID: "nope/nope", Src: "x"})); rep.Error == "" {
		t.Error("unknown challenge accepted")
	}
	rec := httptest.NewRecorder()
	s.handleVet(rec, httptest.NewRequest(http.MethodGet, "/vet", nil))
	if rep := decodeReport(t, rec); rep.Error == "" {
		t.Error("GET /vet accepted")
	}
}

func TestHandleFmt(t *testing.T) {
	s := &server{}
	rec := postJSON(t, s.handleFmt, "/fmt", map[string]string{"src": "package p\nfunc  F( ) int {return 1}\n"})
	var got map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["ok"] != true {
		t.Fatalf("ok = %v (%v)", got["ok"], got["error"])
	}
	if !strings.Contains(got["source"].(string), "func F() int { return 1 }") {
		t.Errorf("source = %q", got["source"])
	}

	// Unparseable source reports the error instead of mangling the input.
	rec = postJSON(t, s.handleFmt, "/fmt", map[string]string{"src": "package p\nfunc {"})
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["ok"] != false || got["error"] == "" {
		t.Errorf("bad source: %#v", got)
	}

	// Wrong method.
	rec = httptest.NewRecorder()
	s.handleFmt(rec, httptest.NewRequest(http.MethodGet, "/fmt", nil))
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["ok"] != false {
		t.Errorf("GET /fmt: %#v", got)
	}
}

func TestHandleHistory(t *testing.T) {
	s, id := fixture(t)
	s.store.save(id, map[string]string{"double.go": "x"}, true, 1, 1, false, true)

	rec := httptest.NewRecorder()
	s.handleHistory(rec, httptest.NewRequest(http.MethodGet, "/history?challengeId="+id, nil))
	var rows []historyRow
	if err := json.Unmarshal(rec.Body.Bytes(), &rows); err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || !rows[0].OK {
		t.Errorf("rows = %+v", rows)
	}

	// No id, and no store: both answer an empty list rather than an error.
	rec = httptest.NewRecorder()
	s.handleHistory(rec, httptest.NewRequest(http.MethodGet, "/history", nil))
	if strings.TrimSpace(rec.Body.String()) != "[]" {
		t.Errorf("no id: body = %q", rec.Body.String())
	}
	rec = httptest.NewRecorder()
	(&server{}).handleHistory(rec, httptest.NewRequest(http.MethodGet, "/history?challengeId=x", nil))
	if strings.TrimSpace(rec.Body.String()) != "[]" {
		t.Errorf("no store: body = %q", rec.Body.String())
	}
}

func TestHandleHistoryStoreError(t *testing.T) {
	s, id := fixture(t)
	s.store.close() // force a query error
	rec := httptest.NewRecorder()
	s.handleHistory(rec, httptest.NewRequest(http.MethodGet, "/history?challengeId="+id, nil))
	if !strings.Contains(rec.Body.String(), "error") {
		t.Errorf("body = %q, want an error payload", rec.Body.String())
	}
}

func TestHandleHealth(t *testing.T) {
	rec := httptest.NewRecorder()
	(&server{}).handleHealth(rec, httptest.NewRequest(http.MethodGet, "/health", nil))
	var got map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got["ok"] != true || got["version"] != version {
		t.Errorf("health = %#v", got)
	}
}

func TestHandleSolved(t *testing.T) {
	s, id := fixture(t)
	s.store.save(id, nil, true, 1, 1, false, true)

	rec := httptest.NewRecorder()
	s.handleSolved(rec, httptest.NewRequest(http.MethodGet, "/solved", nil))
	var got struct {
		Solved []string `json:"solved"`
		Error  string   `json:"error"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Solved) != 1 || got.Solved[0] != id {
		t.Errorf("solved = %q", got.Solved)
	}

	// No store at all: an empty set, not an error.
	rec = httptest.NewRecorder()
	(&server{}).handleSolved(rec, httptest.NewRequest(http.MethodGet, "/solved", nil))
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Solved) != 0 {
		t.Errorf("solved = %q, want empty", got.Solved)
	}

	// Store error surfaces.
	s.store.close()
	rec = httptest.NewRecorder()
	s.handleSolved(rec, httptest.NewRequest(http.MethodGet, "/solved", nil))
	if !strings.Contains(rec.Body.String(), "error") {
		t.Errorf("body = %q, want an error payload", rec.Body.String())
	}
}

func TestWithCORSPreflight(t *testing.T) {
	s := &server{}
	called := false
	h := s.withCORS(func(w http.ResponseWriter, r *http.Request) { called = true })
	req := httptest.NewRequest(http.MethodOptions, "/run", nil)
	req.Header.Set("Origin", "http://localhost:7070")
	rec := httptest.NewRecorder()
	h(rec, req)
	if rec.Code != http.StatusNoContent {
		t.Errorf("preflight status = %d, want 204", rec.Code)
	}
	if called {
		t.Error("preflight must not reach the handler")
	}
	if rec.Header().Get("Access-Control-Allow-Methods") == "" {
		t.Error("missing Access-Control-Allow-Methods")
	}
}

// The cap must shed load rather than let unbounded `go test` processes pile up.
func TestLimiterSheds(t *testing.T) {
	s := &server{limiter: newLimiter(1)}
	release := make(chan struct{})
	blocked := make(chan struct{})
	slow := s.withLimit(func(w http.ResponseWriter, r *http.Request) {
		close(blocked)
		<-release
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		slow(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/run", nil))
	}()
	<-blocked

	// The only slot is taken; a second request sheds after admitWait.
	start := time.Now()
	rec := httptest.NewRecorder()
	s.withLimit(func(w http.ResponseWriter, r *http.Request) {
		t.Error("handler ran while over capacity")
	})(rec, httptest.NewRequest(http.MethodPost, "/run", nil))

	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want 503", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "busy") {
		t.Errorf("body = %q", rec.Body.String())
	}
	if waited := time.Since(start); waited < admitWait {
		t.Errorf("shed after %s, want at least %s of grace", waited, admitWait)
	}

	close(release)
	wg.Wait()

	// The slot is free again.
	ran := false
	rec = httptest.NewRecorder()
	s.withLimit(func(w http.ResponseWriter, r *http.Request) { ran = true })(rec, httptest.NewRequest(http.MethodPost, "/run", nil))
	if !ran {
		t.Error("handler did not run after the slot was released")
	}
}

func TestLimiterAcquireRelease(t *testing.T) {
	l := newLimiter(2)
	if !l.acquire() || !l.acquire() {
		t.Fatal("first two acquires must succeed")
	}
	l.release()
	if !l.acquire() {
		t.Error("acquire after release must succeed")
	}
}
