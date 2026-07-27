package main

import (
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// repoFixture lays out a minimal repo root: one puzzle plus a site/web dir.
func repoFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "double")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module double\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	web := filepath.Join(root, "site", "web")
	if err := os.MkdirAll(web, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(web, "index.html"), []byte("<h1>hi</h1>"), 0o644); err != nil {
		t.Fatal(err)
	}
	return root
}

// A full server boot: real listener, real routes, static file serving, then a
// clean shutdown.
func TestRunServesAPIAndStaticSite(t *testing.T) {
	root := repoFixture(t)
	db := filepath.Join(t.TempDir(), "run.db")

	type result struct{ base string }
	got := make(chan result, 1)
	done := make(chan error, 1)

	go func() {
		done <- run([]string{"-host", "127.0.0.1", "-port", "0", "-root", root, "-db", db},
			func(ln net.Listener, srv *http.Server) {
				base := "http://" + ln.Addr().String()
				got <- result{base}
				// Give the caller time to make its requests, then shut down.
				<-time.After(3 * time.Second)
				_ = srv.Close()
			})
	}()

	var base string
	select {
	case r := <-got:
		base = r.base
	case err := <-done:
		t.Fatalf("server exited early: %v", err)
	case <-time.After(10 * time.Second):
		t.Fatal("server never started")
	}

	body, code := httpGet(t, base+"/health")
	if code != http.StatusOK || !strings.Contains(body, `"ok":true`) {
		t.Errorf("/health = %d %q", code, body)
	}
	body, code = httpGet(t, base+"/")
	if code != http.StatusOK || !strings.Contains(body, "hi") {
		t.Errorf("static index = %d %q", code, body)
	}
	// The limiter is wired into /run: a GET is rejected by decodeRun, which
	// proves the request reached the handler through the middleware chain.
	body, code = httpGet(t, base+"/run")
	if code != http.StatusOK || !strings.Contains(body, "POST") {
		t.Errorf("/run = %d %q", code, body)
	}

	if err := <-done; err != nil {
		t.Errorf("run returned %v", err)
	}
}

func httpGet(t *testing.T, url string) (string, int) {
	t.Helper()
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(b), res.StatusCode
}

// Without site/web the API still comes up; only static serving is skipped.
func TestRoutesWithoutWebDir(t *testing.T) {
	s := &server{cfg: config{root: t.TempDir()}, limiter: newLimiter(1)}
	h := s.routes()
	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	rec := newRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("/health = %d", rec.Code)
	}
	// No file server registered, so an unknown path 404s.
	req, _ = http.NewRequest(http.MethodGet, "/nope", nil)
	rec = newRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("/nope = %d, want 404", rec.Code)
	}
}

func TestRunErrors(t *testing.T) {
	root := repoFixture(t)

	// Unknown flag.
	if err := run([]string{"-nope"}); err == nil {
		t.Error("unknown flag: want an error")
	}
	// No challenges/ anywhere under -root.
	if err := run([]string{"-root", filepath.Join(t.TempDir(), "empty")}); err == nil {
		t.Error("bad root: want an error")
	}
	// A db path that cannot be opened.
	err := run([]string{"-root", root, "-db", filepath.Join(t.TempDir(), "missing-dir", "x.db")})
	if err == nil || !strings.Contains(err.Error(), "open db") {
		t.Errorf("bad db path: err = %v", err)
	}
	// A port already in use.
	ln, lerr := net.Listen("tcp", "127.0.0.1:0")
	if lerr != nil {
		t.Fatal(lerr)
	}
	defer ln.Close()
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	err = run([]string{"-root", root, "-db", filepath.Join(t.TempDir(), "a.db"), "-port", port})
	if err == nil || !strings.Contains(err.Error(), "listen") {
		t.Errorf("port in use: err = %v", err)
	}
}

// Binding off-loopback is allowed but must be announced; this exercises the
// warning path.
func TestRunWarnsWhenNotLoopback(t *testing.T) {
	root := repoFixture(t)
	logs := captureLog(t)

	done := make(chan error, 1)
	started := make(chan struct{})
	go func() {
		done <- run([]string{"-host", "0.0.0.0", "-port", "0", "-root", root,
			"-db", filepath.Join(t.TempDir(), "w.db")},
			func(ln net.Listener, srv *http.Server) {
				close(started)
				<-time.After(200 * time.Millisecond)
				_ = srv.Close()
			})
	}()
	select {
	case <-started:
	case err := <-done:
		t.Fatalf("server exited early: %v", err)
	case <-time.After(10 * time.Second):
		t.Fatal("server never started")
	}
	<-done

	if !strings.Contains(logs.String(), "WARNING") {
		t.Errorf("no warning for a non-loopback bind:\n%s", logs.String())
	}
}

func TestFindRoot(t *testing.T) {
	root := repoFixture(t)

	if got, err := findRoot(root); err != nil || got != root {
		t.Errorf("findRoot(root) = %q, %v", got, err)
	}
	// Walks up from a nested dir.
	nested := filepath.Join(root, "challenges", "junior", "01-a")
	if got, err := findRoot(nested); err != nil || got != root {
		t.Errorf("findRoot(nested) = %q, %v", got, err)
	}
	// An explicit hint with no challenges/ above it fails.
	if _, err := findRoot(filepath.Join(t.TempDir(), "nothing")); err == nil {
		t.Error("findRoot on an empty tree: want an error")
	}

	// No hint: falls back to the working directory.
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	got, err := findRoot("")
	if err != nil {
		t.Fatal(err)
	}
	// macOS resolves TempDir through /private, so compare resolved paths.
	if want, _ := filepath.EvalSymlinks(root); got != want && got != root {
		t.Errorf("findRoot(\"\") = %q, want %q", got, root)
	}

	// Last resort: the executable's own directory. A runner binary sitting in a
	// repo resolves it even when started from an unrelated cwd. Simulated by
	// planting a challenges/ dir next to this test binary.
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	exe, err := os.Executable()
	if err != nil {
		t.Skipf("os.Executable: %v", err)
	}
	exeDir := filepath.Dir(exe)
	planted := filepath.Join(exeDir, "challenges")
	if _, err := os.Stat(planted); err == nil {
		t.Skip("the test binary already sits next to a challenges/ dir")
	}
	if err := os.Mkdir(planted, 0o755); err != nil {
		t.Skipf("cannot plant a challenges/ dir next to the test binary: %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(planted) })

	got, err = findRoot("")
	if err != nil {
		t.Fatalf("findRoot(\"\") = %v; want the executable's dir", err)
	}
	if resolved, _ := filepath.EvalSymlinks(exeDir); got != exeDir && got != resolved {
		t.Errorf("findRoot(\"\") = %q, want %q", got, exeDir)
	}
}

func TestEnvOr(t *testing.T) {
	t.Setenv("GW_TEST_KEY", "set")
	if got := envOr("GW_TEST_KEY", "fallback"); got != "set" {
		t.Errorf("envOr = %q, want the env value", got)
	}
	if got := envOr("GW_TEST_MISSING", "fallback"); got != "fallback" {
		t.Errorf("envOr = %q, want the fallback", got)
	}
}

// With no -db flag the runner keeps state under $HOME/.gopher-workplace.
func TestRunUsesDefaultDBPath(t *testing.T) {
	root := repoFixture(t)
	home := t.TempDir()
	t.Setenv("HOME", home)

	done := make(chan error, 1)
	started := make(chan struct{})
	go func() {
		done <- run([]string{"-root", root, "-port", "0"},
			func(ln net.Listener, srv *http.Server) {
				close(started)
				_ = srv.Close()
			})
	}()
	select {
	case <-started:
	case err := <-done:
		t.Fatalf("server exited early: %v", err)
	case <-time.After(10 * time.Second):
		t.Fatal("server never started")
	}
	<-done

	if _, err := os.Stat(filepath.Join(home, ".gopher-workplace", "runner.db")); err != nil {
		t.Errorf("default db not created under HOME: %v", err)
	}
}

// A challenge dir that exists but holds no go.mod is not a challenge.
func TestChallengeDirRequiresGoMod(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "challenges", "junior", "01-a", "01-b", "empty")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	s := &server{cfg: config{root: root}}
	if _, err := s.challengeDir("junior/01-a/01-b/empty"); err == nil {
		t.Error("a dir without go.mod was accepted as a challenge")
	}
}

// Nothing to find: no hint, no challenges/ above cwd, none above the binary.
func TestFindRootFailsWhenNoRepoAnywhere(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(wd) })
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	if exe, err := os.Executable(); err == nil {
		if _, err := os.Stat(filepath.Join(filepath.Dir(exe), "challenges")); err == nil {
			t.Skip("the test binary sits next to a challenges/ dir")
		}
	}
	if root, err := findRoot(""); err == nil {
		t.Errorf("findRoot(\"\") = %q, nil; want an error", root)
	}
}
