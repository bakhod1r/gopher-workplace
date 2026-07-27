package main

import (
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestMain lets the test binary re-exec itself as the real command: when
// GW_TEST_MAIN is set we call main() and exit, which is the only way to cover
// main() itself.
func TestMain(m *testing.M) {
	if os.Getenv("GW_TEST_MAIN") != "" {
		os.Args = append([]string{"localrunner"}, strings.Fields(os.Getenv("GW_TEST_MAIN_ARGS"))...)
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

// main() must surface a startup failure as a non-zero exit, not a silent hang.
func TestMainExitsOnStartupFailure(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-challenges-here")
	if err := os.MkdirAll(missing, 0o755); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainExitsOnStartupFailure")
	cmd.Env = append(os.Environ(),
		"GW_TEST_MAIN=1",
		"GW_TEST_MAIN_ARGS=-root "+missing,
	)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("exit status 0; want failure. output:\n%s", out)
	}
	if !strings.Contains(string(out), "cannot locate repo root") {
		t.Errorf("output = %q, want the root error", out)
	}
}

// main() on a healthy config serves until killed.
func TestMainServes(t *testing.T) {
	if testing.Short() {
		t.Skip("spawns a subprocess")
	}
	root := repoFixture(t)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	ln.Close() // free the port for the child

	cmd := exec.Command(os.Args[0], "-test.run=TestMainServes")
	cmd.Env = append(os.Environ(),
		"GW_TEST_MAIN=1",
		"GW_TEST_MAIN_ARGS=-root "+root+" -port "+port+" -db "+filepath.Join(t.TempDir(), "m.db"),
	)
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_, _ = cmd.Process.Wait()
	})

	base := "http://127.0.0.1:" + port
	deadline := time.Now().Add(20 * time.Second)
	for time.Now().Before(deadline) {
		res, err := http.Get(base + "/health")
		if err == nil {
			res.Body.Close()
			if res.StatusCode == http.StatusOK {
				return // main() came up and served
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("main() never served /health")
}

// Serve returning something other than ErrServerClosed must propagate: closing
// the listener out from under it is the cheapest way to produce that.
func TestRunPropagatesServeError(t *testing.T) {
	root := repoFixture(t)
	done := make(chan error, 1)
	go func() {
		done <- run([]string{"-root", root, "-port", "0", "-db", filepath.Join(t.TempDir(), "s.db")},
			func(ln net.Listener, srv *http.Server) {
				time.Sleep(100 * time.Millisecond)
				_ = ln.Close()
			})
	}()
	select {
	case err := <-done:
		if err == nil {
			t.Error("run returned nil after the listener died; want the Serve error")
		}
	case <-time.After(15 * time.Second):
		t.Fatal("run never returned")
	}
}

// findRoot's last resort is the executable's own directory: a runner binary
// sitting inside the repo works even when started from an unrelated cwd. Cover
// it by copying this test binary into a fixture repo and running it from a
// directory with no challenges/ above it.
func TestFindRootFallsBackToExecutableDir(t *testing.T) {
	if testing.Short() {
		t.Skip("spawns a subprocess")
	}
	root := repoFixture(t)
	binDir := filepath.Join(root, "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	self, err := os.ReadFile(os.Args[0])
	if err != nil {
		t.Skipf("cannot read the test binary: %v", err)
	}
	bin := filepath.Join(binDir, "localrunner.test")
	if err := os.WriteFile(bin, self, 0o755); err != nil {
		t.Fatal(err)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	ln.Close()

	cmd := exec.Command(bin, "-test.run=TestFindRootFallsBackToExecutableDir")
	// No -root, and a cwd with no challenges/ above it: only the executable's
	// own location can resolve the repo.
	cmd.Dir = t.TempDir()
	cmd.Env = append(os.Environ(),
		"GW_TEST_MAIN=1",
		"GW_ROOT=",
		"HOME="+t.TempDir(),
		"GW_TEST_MAIN_ARGS=-port "+port,
	)
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = cmd.Process.Kill()
		_, _ = cmd.Process.Wait()
	})

	deadline := time.Now().Add(20 * time.Second)
	for time.Now().Before(deadline) {
		res, err := http.Get("http://127.0.0.1:" + port + "/health")
		if err == nil {
			res.Body.Close()
			if res.StatusCode == http.StatusOK {
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	t.Fatal("the runner never came up from its executable-relative root")
}
