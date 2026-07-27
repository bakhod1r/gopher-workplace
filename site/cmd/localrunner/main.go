// Command localrunner is the backend for Gopher Workplace. It runs the real Go
// toolchain (go test, -race, benchmarks, GC/finalizer timing) against submitted
// solutions, and serves the static site from the same origin so the page and
// the API never disagree about ports.
//
// The frontend probes GET /health on load and routes every Run/Submit to
// POST /run.
//
// Usage (its own module — run it from this directory):
//
//	go run .                          # serves 127.0.0.1:7070, auto-finds challenges/
//	go run . -port 9090
//	GW_RUNNER_PORT=9090 GW_ROOT=/path/to/repo go run .
//
// SAFETY: this executes arbitrary user Go code on your machine, as you, with
// your file access and your network. It is NOT a sandbox — it is a convenience
// for running code you wrote yourself. Mitigations: it binds 127.0.0.1 only,
// rejects cross-origin browser requests, runs each request in a throwaway temp
// module with a 20s timeout (process group killed on expiry) and an isolated
// GOCACHE/GOPATH, and caps concurrent runs. Do not point it at code you would
// not run by hand.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

const version = "0.1.0"

// config holds resolved server settings.
type config struct {
	host  string // listen address; 127.0.0.1 unless deliberately overridden
	port  string
	root  string // repo root: the dir containing challenges/
	dbDSN string
}

// addr is the host:port the server binds.
func (c config) addr() string { return net.JoinHostPort(c.host, c.port) }

// loopback reports whether the bind address is reachable only from this
// machine. Anything else means arbitrary code execution is exposed to the
// network, so we warn loudly.
func (c config) loopback() bool {
	ip := net.ParseIP(c.host)
	return c.host == "localhost" || (ip != nil && ip.IsLoopback())
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

// run is main's body, split out so tests can drive a real server on an
// ephemeral port and shut it down. ready, when non-nil, is closed once the
// listener is up; stop ends the server.
func run(args []string, hooks ...serveHook) error {
	fs := flag.NewFlagSet("localrunner", flag.ContinueOnError)
	var (
		portFlag = fs.String("port", envOr("GW_RUNNER_PORT", "7070"), "listen port")
		hostFlag = fs.String("host", envOr("GW_HOST", "127.0.0.1"), "listen address; loopback only unless you know what you are doing")
		rootFlag = fs.String("root", os.Getenv("GW_ROOT"), "repo root (dir containing challenges/); auto-detected if empty")
		dbFlag   = fs.String("db", os.Getenv("GW_DB"), "sqlite db path (default ~/.gopher-workplace/runner.db)")
	)
	if err := fs.Parse(args); err != nil {
		return err
	}

	root, err := findRoot(*rootFlag)
	if err != nil {
		return fmt.Errorf("cannot locate repo root: %w", err)
	}
	dbPath := *dbFlag
	if dbPath == "" {
		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, ".gopher-workplace")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create state dir: %w", err)
		}
		dbPath = filepath.Join(dir, "runner.db")
	}

	cfg := config{host: *hostFlag, port: *portFlag, root: root, dbDSN: dbPath}

	store, err := openStore(dbPath)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer store.close()
	store.startRetention(30 * 24 * time.Hour) // keep submissions 30 days

	srv := &server{cfg: cfg, store: store, limiter: newLimiter(maxConcurrentRuns)}

	ln, err := net.Listen("tcp", cfg.addr())
	if err != nil {
		return fmt.Errorf("listen on %s: %w", cfg.addr(), err)
	}

	log.Printf("gopher-workplace local runner %s", version)
	log.Printf("  root:  %s", root)
	log.Printf("  db:    %s (30-day retention)", dbPath)
	log.Printf("  limit: %d concurrent runs", maxConcurrentRuns)
	log.Printf("  listening on http://%s", ln.Addr())
	if !cfg.loopback() {
		log.Printf("  WARNING: bound to %s, not loopback — anyone who can reach this", cfg.host)
		log.Printf("           port can execute arbitrary Go code as %s. Use -host 127.0.0.1.", os.Getenv("USER"))
	}

	httpSrv := &http.Server{Handler: srv.routes()}
	for _, h := range hooks {
		go h(ln, httpSrv)
	}
	if err := httpSrv.Serve(ln); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// serveHook receives the live listener and server; tests use it to talk to the
// running instance and then shut it down.
type serveHook func(net.Listener, *http.Server)

// routes wires the API and the static site. Anything not matching an endpoint
// is served from site/web, so the page and the API share one origin.
func (s *server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.withCORS(s.handleHealth))
	mux.HandleFunc("/run", s.withCORS(s.withLimit(s.handleRun)))
	mux.HandleFunc("/vet", s.withCORS(s.withLimit(s.handleVet)))
	mux.HandleFunc("/fmt", s.withCORS(s.handleFmt))
	mux.HandleFunc("/history", s.withCORS(s.handleHistory))
	mux.HandleFunc("/solved", s.withCORS(s.handleSolved))

	webDir := filepath.Join(s.cfg.root, "site", "web")
	if fi, err := os.Stat(webDir); err == nil && fi.IsDir() {
		mux.Handle("/", http.FileServer(http.Dir(webDir)))
		log.Printf("  web:   %s", webDir)
	} else {
		log.Printf("  web:   (not found at %s — static site not served)", webDir)
	}
	return mux
}

type server struct {
	cfg     config
	store   *store
	limiter *limiter
}

// findRoot resolves the repo root: the nearest ancestor holding a challenges/
// dir. Tries the explicit hint, then cwd, then the executable's dir.
func findRoot(hint string) (string, error) {
	tryFrom := func(start string) (string, bool) {
		dir := start
		for {
			if fi, err := os.Stat(filepath.Join(dir, "challenges")); err == nil && fi.IsDir() {
				return dir, true
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				return "", false
			}
			dir = parent
		}
	}
	if hint != "" {
		if abs, err := filepath.Abs(hint); err == nil {
			if r, ok := tryFrom(abs); ok {
				return r, nil
			}
		}
		return "", fmt.Errorf("no challenges/ under -root %q", hint)
	}
	if wd, err := os.Getwd(); err == nil {
		if r, ok := tryFrom(wd); ok {
			return r, nil
		}
	}
	if exe, err := os.Executable(); err == nil {
		if r, ok := tryFrom(filepath.Dir(exe)); ok {
			return r, nil
		}
	}
	return "", fmt.Errorf("challenges/ not found (pass -root)")
}

// withCORS allows the local static site to call the API and rejects every other
// browser origin, then answers preflight requests.
//
// This is a security boundary, not a convenience: /run executes arbitrary Go
// code, so a wildcard would let any page you happen to visit drive it. A
// same-origin request (no Origin header, or one we serve ourselves) is fine; an
// unknown Origin is refused outright rather than answered without CORS headers,
// so the failure is visible instead of a confusing opaque error.
func (s *server) withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if !allowedOrigin(origin) {
				http.Error(w, "cross-origin request refused: this runner only "+
					"serves its own page (http://localhost:PORT)", http.StatusForbidden)
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h(w, r)
	}
}

var loopbackOrigin = regexp.MustCompile(`^https?://(localhost|127\.0\.0\.1|\[::1\])(:\d+)?$`)

// allowedOrigin permits only pages served from this machine: the runner's own
// origin, or a static server (make serve) on another loopback port. Everything
// else — including any public site — is refused.
func allowedOrigin(o string) bool {
	return loopbackOrigin.MatchString(o)
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{"ok": true, "version": version})
}

// handleSolved returns the authoritative solved set from SQLite: challenge ids
// with at least one passing Submit.
func (s *server) handleSolved(w http.ResponseWriter, r *http.Request) {
	ids := []string{}
	if s.store != nil {
		got, err := s.store.solvedIDs()
		if err != nil {
			writeJSON(w, map[string]any{"error": err.Error()})
			return
		}
		ids = got
	}
	writeJSON(w, map[string]any{"solved": ids})
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
