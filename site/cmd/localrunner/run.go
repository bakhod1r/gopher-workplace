package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"
)

// runTimeout bounds a single toolchain invocation. A var, not a const, so
// tests can shrink it instead of waiting out the real budget.
var runTimeout = 20 * time.Second

// runReq is the POST /run body. Either files (name->content) or src (single
// candidate file) must be provided, plus the challengeId (path under
// challenges/).
type runReq struct {
	ChallengeID string            `json:"challengeId"`
	Files       map[string]string `json:"files"`
	Src         string            `json:"src"`
	Submit      bool              `json:"submit"` // true = Submit (counts toward solved), false = Run
}

// caseResult / report mirror EXACTLY the JSON the frontend render() consumes;
// do not change these field names.
type caseResult struct {
	Name string `json:"name"`
	Pass bool   `json:"pass"`
	Got  string `json:"got"`
	Want string `json:"want"`
}

type report struct {
	OK        bool         `json:"ok"`
	CompileOK bool         `json:"compileOk"`
	Error     string       `json:"error"`
	Cases     []caseResult `json:"cases"`
	Warnings  []string     `json:"warnings,omitempty"` // non-blocking: hardcode / clean-code hints
	// Output is everything the run printed: fmt.Print* to stdout, log.Print* to
	// stderr, plus the test framing lines. Shown verbatim in the console's
	// output tab so print-debugging works.
	Output string `json:"output,omitempty"`
}

func (s *server) handleRun(w http.ResponseWriter, r *http.Request) {
	req, err := decodeRun(r)
	if err != nil {
		writeJSON(w, report{Error: err.Error()})
		return
	}
	chDir, err := s.challengeDir(req.ChallengeID)
	if err != nil {
		writeJSON(w, report{Error: err.Error()})
		return
	}

	tmp, err := os.MkdirTemp("", "gw-run-*")
	if err != nil {
		writeJSON(w, report{Error: "temp dir: " + err.Error()})
		return
	}
	defer os.RemoveAll(tmp)

	if err := materialize(chDir, tmp, req); err != nil {
		writeJSON(w, report{Error: err.Error()})
		return
	}

	race := needsRace(tmp)
	rep := goTest(tmp, race)

	// Clean-code hint: flag a non-gofmt-formatted submission. Combined with the
	// WARN: lines from guard tests, these warnings do NOT block a Run, but they
	// DO block a Submit from being accepted (counting as solved).
	if rep.CompileOK {
		rep.Warnings = append(rep.Warnings, gofmtWarn(tmp)...)
	}

	// A Submit is accepted (solved) only when it passes AND is warning-free.
	accepted := req.Submit && rep.OK && len(rep.Warnings) == 0

	// Persist the run (best-effort; never fail the request on db error). Only an
	// accepted submit is stored as submitted=1, so /solved excludes hardcoded or
	// unformatted answers.
	if s.store != nil {
		files := req.Files
		if files == nil && req.Src != "" {
			files = map[string]string{"src": req.Src}
		}
		s.store.save(req.ChallengeID, files, rep.OK, len(rep.Cases), passCount(rep), race, accepted)
	}

	writeJSON(w, rep)
}

// gofmtWarn returns a warning for each top-level non-test .go file in dir that
// is not gofmt-formatted. Best-effort: returns nil if gofmt is unavailable.
func gofmtWarn(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var files []string
	for _, e := range entries {
		n := e.Name()
		if strings.HasSuffix(n, ".go") && !strings.HasSuffix(n, "_test.go") {
			files = append(files, filepath.Join(dir, n))
		}
	}
	if len(files) == 0 {
		return nil
	}
	out, err := exec.Command("gofmt", append([]string{"-l"}, files...)...).Output()
	if err != nil {
		return nil
	}
	var w []string
	for _, f := range strings.Fields(string(out)) {
		w = append(w, "not gofmt-clean: "+filepath.Base(f)+" — use Format")
	}
	return w
}

func decodeRun(r *http.Request) (runReq, error) {
	if r.Method != http.MethodPost {
		return runReq{}, errors.New("POST required")
	}
	var req runReq
	if err := json.NewDecoder(io.LimitReader(r.Body, 1<<20)).Decode(&req); err != nil {
		return runReq{}, fmt.Errorf("bad body: %w", err)
	}
	if req.ChallengeID == "" {
		return runReq{}, errors.New("challengeId required")
	}
	if len(req.Files) == 0 && req.Src == "" {
		return runReq{}, errors.New("files or src required")
	}
	return req, nil
}

// challengeDir validates the id and returns the real challenge directory. It
// rejects any path escaping challenges/.
func (s *server) challengeDir(id string) (string, error) {
	clean := filepath.Clean("/" + id)     // force-rooted, collapses ..
	rel := strings.TrimPrefix(clean, "/") // drop leading slash
	if rel == "" || strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("invalid challengeId %q", id)
	}
	base := filepath.Join(s.cfg.root, "challenges")
	dir := filepath.Join(base, rel)

	// The lexical check above stops `..` in the id, but a symlink *inside*
	// challenges/ could still point anywhere on disk. Compare fully resolved
	// paths so the real target has to live under challenges/ too.
	realBase, err := filepath.EvalSymlinks(base)
	if err != nil {
		return "", fmt.Errorf("challenges/ not readable under %q", s.cfg.root)
	}
	realDir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return "", fmt.Errorf("no such challenge %q", id)
	}
	if !strings.HasPrefix(realDir, realBase+string(os.PathSeparator)) {
		return "", fmt.Errorf("path escapes challenges/: %q", id)
	}
	if _, err := os.Stat(filepath.Join(dir, "go.mod")); err != nil {
		return "", fmt.Errorf("no such challenge %q", id)
	}
	return dir, nil
}

// materialize copies the real challenge dir into tmp, then overwrites the
// candidate file(s) with submitted source.
func materialize(chDir, tmp string, req runReq) error {
	if err := copyDir(chDir, tmp); err != nil {
		return fmt.Errorf("copy challenge: %w", err)
	}
	if len(req.Files) > 0 {
		for name, content := range req.Files {
			if err := writeCandidate(tmp, name, content); err != nil {
				return err
			}
		}
		return nil
	}
	// src form: overwrite the starter file (shortest non-test .go name).
	starter, err := starterFile(tmp)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(tmp, starter), []byte(req.Src), 0o644)
}

func writeCandidate(tmp, name, content string) error {
	clean := filepath.Base(name) // no subdirs / traversal for candidate files
	if clean == "" || clean == "." || !strings.HasSuffix(clean, ".go") {
		return fmt.Errorf("invalid file name %q", name)
	}
	return os.WriteFile(filepath.Join(tmp, clean), []byte(content), 0o644)
}

// starterFile picks the candidate file the same way gen-problems.py does: the
// shortest-named non-test .go file.
func starterFile(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	best := ""
	for _, e := range entries {
		n := e.Name()
		if !strings.HasSuffix(n, ".go") || strings.HasSuffix(n, "_test.go") {
			continue
		}
		if best == "" || len(n) < len(best) {
			best = n
		}
	}
	if best == "" {
		return "", errors.New("no candidate .go file in challenge")
	}
	return best, nil
}

func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue // challenge puzzles are flat; skip nested dirs
		}
		b, err := os.ReadFile(filepath.Join(src, e.Name()))
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dst, e.Name()), b, 0o644); err != nil {
			return err
		}
	}
	return nil
}

var concurrencyRe = regexp.MustCompile(`\bgo func\b|\bgo [A-Za-z_]|"sync"|"sync/atomic"|goroutine`)

// needsRace scans the module's .go sources for concurrency markers; if found we
// run with -race.
func needsRace(dir string) bool {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		if concurrencyRe.Match(b) {
			return true
		}
	}
	return false
}

// goTest runs the challenge's real verification and maps -json output into the
// frontend Report shape.
func goTest(dir string, race bool) report {
	ctx, cancel := context.WithTimeout(context.Background(), runTimeout)
	defer cancel()

	args := []string{"test", "-run", ".", "-count=1", "-json"}
	if race {
		args = append(args, "-race")
	}
	args = append(args, "./...")

	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = dir
	cmd.Env = sandboxEnv(dir)
	// Own process group so we can kill the whole tree (test binary + children)
	// on timeout, not just `go`.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		if cmd.Process != nil {
			_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return cmd.Process.Kill()
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return report{Error: fmt.Sprintf("timed out after %s (possible infinite loop) — process killed", runTimeout)}
	}
	return parseTestJSON(stdout.Bytes(), stderr.String(), err)
}

// sandboxEnv returns a locked-down environment: no network, isolated caches.
func sandboxEnv(dir string) []string {
	env := []string{
		"GOPROXY=off",
		"GOFLAGS=-mod=mod",
		"GOSUMDB=off",
		"GOCACHE=" + filepath.Join(dir, ".gocache"),
		"GOPATH=" + filepath.Join(dir, ".gopath"),
		"GOTMPDIR=" + dir,
		"CGO_ENABLED=1", // -race needs cgo
	}
	// Preserve PATH, HOME, GOROOT so `go` and the toolchain resolve.
	for _, k := range []string{"PATH", "HOME", "GOROOT", "USER", "TMPDIR"} {
		if v := os.Getenv(k); v != "" {
			env = append(env, k+"="+v)
		}
	}
	return env
}

type testEvent struct {
	Action string `json:"Action"`
	Test   string `json:"Test"`
	Output string `json:"Output"`
}

var (
	compileErrRe = regexp.MustCompile(`\.go:\d+:\d+:`)

	// Failure messages come in two idioms. Explicit "got X want Y" first, then
	// the more common Go form "Call(args) = X, want Y" — most puzzle tests use
	// the latter, so without it the UI's got/want columns stay empty.
	gotWantREs = []*regexp.Regexp{
		regexp.MustCompile(`(?i)got[:= ]+(.+?)[,;]?\s+want[:= ]+(.+)`),
		regexp.MustCompile(`(?i)=\s*(.+?)[,;]\s*want[:= ]+(.+)`),
	}
)

// parseTestJSON converts `go test -json` output into a report. Build failures
// produce compileOk=false with the compiler output in Error.
func parseTestJSON(stdout []byte, stderr string, runErr error) report {
	type acc struct {
		out    strings.Builder
		result string
	}
	tests := map[string]*acc{}
	order := []string{}
	var pkgOut strings.Builder
	var stream strings.Builder // every printed line, in emission order
	buildFailed := false

	sc := bufio.NewScanner(bytes.NewReader(stdout))
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var ev testEvent
		if err := json.Unmarshal(line, &ev); err != nil {
			// Non-JSON: a raw build error emitted before test2json framing.
			pkgOut.Write(line)
			pkgOut.WriteByte('\n')
			stream.Write(line)
			stream.WriteByte('\n')
			continue
		}
		if ev.Action == "output" {
			stream.WriteString(ev.Output)
		}
		if ev.Test == "" {
			if ev.Output != "" {
				pkgOut.WriteString(ev.Output)
				if strings.Contains(ev.Output, "[build failed]") ||
					strings.Contains(ev.Output, "build failed") {
					buildFailed = true
				}
			}
			continue
		}
		a := tests[ev.Test]
		if a == nil {
			a = &acc{}
			tests[ev.Test] = a
			order = append(order, ev.Test)
		}
		switch ev.Action {
		case "output":
			a.out.WriteString(ev.Output)
		case "pass", "fail", "skip":
			a.result = ev.Action
		}
	}

	pkg := pkgOut.String()
	if compileErrRe.MatchString(pkg) && len(order) == 0 {
		buildFailed = true
	}

	if buildFailed || (len(order) == 0 && runErr != nil) {
		msg := strings.TrimSpace(pkg)
		if msg == "" {
			msg = strings.TrimSpace(stderr)
		}
		if msg == "" && runErr != nil {
			msg = runErr.Error()
		}
		return report{CompileOK: false, Error: msg, Output: printed(stream.String(), stderr)}
	}

	rep := report{CompileOK: true, OK: true, Output: printed(stream.String(), stderr)}
	for _, name := range order {
		a := tests[name]
		// Collect WARN: lines from any test's output (guard checks log these to
		// flag hardcoding / style without failing the submission).
		for _, ln := range strings.Split(a.out.String(), "\n") {
			if i := strings.Index(ln, "WARN:"); i >= 0 {
				rep.Warnings = append(rep.Warnings, strings.TrimSpace(ln[i+len("WARN:"):]))
			}
		}
		if a.result == "skip" {
			continue
		}
		pass := a.result == "pass"
		if !pass {
			rep.OK = false
		}
		c := caseResult{Name: name, Pass: pass}
		if !pass {
			c.Got, c.Want = extractGotWant(a.out.String())
		}
		rep.Cases = append(rep.Cases, c)
	}
	if len(rep.Cases) == 0 {
		// Nothing ran but no build error: surface package output as the message.
		rep.OK = false
		rep.Error = strings.TrimSpace(pkg)
	}
	return rep
}

// testFraming matches the lines `go test` prints about itself — the run/pass/
// fail scaffolding and the final package verdict. They belong in the test-result
// panel, not in the output tab, which is for what the candidate's own code
// printed with fmt and log.
var testFraming = regexp.MustCompile(`^\s*(=== (RUN|PAUSE|CONT|NAME)\b|--- (PASS|FAIL|SKIP|BENCH)\b|(PASS|FAIL|ok|\?)\s*$|(ok|FAIL|\?)\s+\S+\s)`)

// printed joins what the run wrote — the test2json output stream minus the test
// framing, plus anything the toolchain sent to stderr — and caps it so one
// runaway loop cannot ship a megabyte of text to the browser.
func printed(stream, stderr string) string {
	var b strings.Builder
	for _, ln := range strings.Split(stream, "\n") {
		if testFraming.MatchString(ln) {
			continue
		}
		b.WriteString(ln)
		b.WriteByte('\n')
	}
	out := b.String()
	if s := strings.TrimSpace(stderr); s != "" {
		out += s + "\n"
	}
	const max = 256 * 1024
	if len(out) > max {
		out = out[:max] + "\n… output truncated\n"
	}
	return strings.TrimRight(out, "\n")
}

// extractGotWant pulls a "got X want Y" pair from a test's failure output; if
// none is found it returns a condensed failure message as got, empty want.
func extractGotWant(out string) (got, want string) {
	for _, ln := range strings.Split(out, "\n") {
		for _, re := range gotWantREs {
			if m := re.FindStringSubmatch(ln); m != nil {
				return strings.TrimSpace(m[1]), strings.TrimSpace(m[2])
			}
		}
	}
	// Condense: first non-empty, non-framing line.
	for _, ln := range strings.Split(out, "\n") {
		t := strings.TrimSpace(ln)
		if t == "" || strings.HasPrefix(t, "=== RUN") || strings.HasPrefix(t, "--- FAIL") || strings.HasPrefix(t, "--- PASS") {
			continue
		}
		if len(t) > 200 {
			t = t[:200] + "…"
		}
		return t, ""
	}
	return "test failed", ""
}

func passCount(r report) int {
	n := 0
	for _, c := range r.Cases {
		if c.Pass {
			n++
		}
	}
	return n
}
