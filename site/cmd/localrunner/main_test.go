package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestAllowedOrigin(t *testing.T) {
	allow := []string{
		"http://localhost:7070",
		"http://localhost:8080",
		"http://127.0.0.1:9090",
		"https://localhost:7070",
		"http://[::1]:7070",
		"http://localhost",
	}
	for _, o := range allow {
		if !allowedOrigin(o) {
			t.Errorf("allowedOrigin(%q) = false, want true", o)
		}
	}
	// /run executes arbitrary code, so anything not served from this machine
	// must be refused — including hosts that merely *contain* a loopback name.
	deny := []string{
		"https://evil.example",
		"http://localhost.evil.example",
		"http://localhost.attacker.com:7070",
		"http://127.0.0.1.evil.example",
		"https://gopher-workplace.netlify.app",
		"null",
		"",
	}
	for _, o := range deny {
		if allowedOrigin(o) {
			t.Errorf("allowedOrigin(%q) = true, want false", o)
		}
	}
}

func TestWithCORSRefusesForeignOrigin(t *testing.T) {
	s := &server{}
	called := false
	h := s.withCORS(func(w http.ResponseWriter, r *http.Request) { called = true })

	req := httptest.NewRequest(http.MethodPost, "/run", nil)
	req.Header.Set("Origin", "https://evil.example")
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusForbidden)
	}
	if called {
		t.Error("handler ran for a foreign origin; it must not be reached")
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Errorf("Access-Control-Allow-Origin = %q, want empty", got)
	}
}

func TestWithCORSAllowsLoopbackAndSameOrigin(t *testing.T) {
	s := &server{}
	for _, origin := range []string{"", "http://localhost:8080"} {
		called := false
		h := s.withCORS(func(w http.ResponseWriter, r *http.Request) { called = true })
		req := httptest.NewRequest(http.MethodPost, "/run", nil)
		if origin != "" {
			req.Header.Set("Origin", origin)
		}
		rec := httptest.NewRecorder()
		h(rec, req)
		if !called {
			t.Errorf("origin %q: handler not reached", origin)
		}
		if got := rec.Header().Get("Access-Control-Allow-Origin"); got != origin {
			t.Errorf("origin %q: Access-Control-Allow-Origin = %q", origin, got)
		}
	}
}

func TestConfigLoopback(t *testing.T) {
	cases := map[string]bool{
		"127.0.0.1": true,
		"localhost": true,
		"::1":       true,
		"0.0.0.0":   false,
		"10.0.0.5":  false,
	}
	for host, want := range cases {
		if got := (config{host: host}).loopback(); got != want {
			t.Errorf("loopback(%q) = %v, want %v", host, got, want)
		}
	}
}

func TestConfigAddr(t *testing.T) {
	if got := (config{host: "127.0.0.1", port: "7070"}).addr(); got != "127.0.0.1:7070" {
		t.Errorf("addr = %q", got)
	}
	if got := (config{host: "::1", port: "7070"}).addr(); got != "[::1]:7070" {
		t.Errorf("addr = %q, want brackets around the IPv6 host", got)
	}
}

// challengeDir is the only thing standing between a request and the rest of the
// filesystem, so traversal must be rejected however it is spelled.
func TestChallengeDirRejectsTraversal(t *testing.T) {
	root := t.TempDir()
	ch := filepath.Join(root, "challenges", "junior", "01-topic", "01-sub", "puzzle")
	if err := os.MkdirAll(ch, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(ch, "go.mod"), []byte("module x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// A sibling of challenges/ that must stay unreachable.
	secret := filepath.Join(root, "secret")
	if err := os.MkdirAll(secret, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(secret, "go.mod"), []byte("module s\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	s := &server{cfg: config{root: root}}

	good := "junior/01-topic/01-sub/puzzle"
	if got, err := s.challengeDir(good); err != nil || got != ch {
		t.Errorf("challengeDir(%q) = %q, %v; want %q, nil", good, got, err, ch)
	}

	bad := []string{
		"../secret",
		"../../secret",
		"junior/../../secret",
		"/../secret",
		"",
		"/",
		"junior/01-topic/01-sub/nope",
	}
	for _, id := range bad {
		if got, err := s.challengeDir(id); err == nil {
			t.Errorf("challengeDir(%q) = %q, nil; want an error", id, got)
		}
	}
}

func TestStarterFilePicksShortestNonTest(t *testing.T) {
	dir := t.TempDir()
	for _, n := range []string{"dedupe.go", "dedupe_test.go", "helpers_extra.go"} {
		if err := os.WriteFile(filepath.Join(dir, n), []byte("package p\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	got, err := starterFile(dir)
	if err != nil {
		t.Fatal(err)
	}
	if got != "dedupe.go" {
		t.Errorf("starterFile = %q, want dedupe.go", got)
	}

	if _, err := starterFile(t.TempDir()); err == nil {
		t.Error("starterFile on an empty dir: want an error")
	}
}

func TestWriteCandidateRejectsPathsAndNonGo(t *testing.T) {
	dir := t.TempDir()
	for _, name := range []string{"../escape.go", "sub/dir/x.go", "notgo.txt", ""} {
		err := writeCandidate(dir, name, "package p\n")
		if name == "../escape.go" || name == "sub/dir/x.go" {
			// filepath.Base strips the path, so these land in dir as a plain
			// file — the point is that nothing is written outside dir.
			if err != nil {
				continue
			}
			base := filepath.Base(name)
			if _, statErr := os.Stat(filepath.Join(dir, base)); statErr != nil {
				t.Errorf("%q: expected the file to be written inside dir", name)
			}
			if _, statErr := os.Stat(filepath.Join(filepath.Dir(dir), "escape.go")); statErr == nil {
				t.Errorf("%q: escaped the temp dir", name)
			}
			continue
		}
		if err == nil {
			t.Errorf("writeCandidate(%q): want an error", name)
		}
	}
}

func TestNeedsRace(t *testing.T) {
	cases := map[string]bool{
		"package p\nfunc f() { go func() {}() }\n": true,
		"package p\nimport \"sync\"\n":             true,
		"package p\nfunc f() int { return 1 }\n":   false,
	}
	for src, want := range cases {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "a.go"), []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
		if got := needsRace(dir); got != want {
			t.Errorf("needsRace(%q) = %v, want %v", src, got, want)
		}
	}
}

// A symlink inside challenges/ pointing outside it must be refused: the lexical
// `..` check cannot see through symlinks, so the resolved path is compared too.
func TestChallengeDirRejectsSymlinkEscape(t *testing.T) {
	root := t.TempDir()
	base := filepath.Join(root, "challenges")
	if err := os.MkdirAll(base, 0o755); err != nil {
		t.Fatal(err)
	}
	// A real puzzle living outside challenges/.
	outside := filepath.Join(t.TempDir(), "elsewhere")
	if err := os.MkdirAll(outside, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(outside, "go.mod"), []byte("module x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(base, "sneaky")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	// A symlink that goes nowhere.
	if err := os.Symlink(filepath.Join(root, "nope"), filepath.Join(base, "dangling")); err != nil {
		t.Fatal(err)
	}

	s := &server{cfg: config{root: root}}
	if got, err := s.challengeDir("sneaky"); err == nil {
		t.Errorf("challengeDir(sneaky) = %q, nil; a symlink out of challenges/ must be refused", got)
	}
	if _, err := s.challengeDir("dangling"); err == nil {
		t.Error("challengeDir(dangling): want an error")
	}

	// A symlink that stays inside challenges/ is fine.
	real := filepath.Join(base, "real")
	if err := os.MkdirAll(real, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(real, "go.mod"), []byte("module y\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(real, filepath.Join(base, "alias")); err != nil {
		t.Fatal(err)
	}
	if _, err := s.challengeDir("alias"); err != nil {
		t.Errorf("challengeDir(alias) = %v; an in-tree symlink is legitimate", err)
	}
}

// No challenges/ at all: the resolve fails before anything touches the disk
// below it.
func TestChallengeDirWithoutChallengesDir(t *testing.T) {
	s := &server{cfg: config{root: t.TempDir()}}
	if _, err := s.challengeDir("junior/a/b/c"); err == nil {
		t.Error("want an error when challenges/ does not exist")
	}
}
