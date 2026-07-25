// Command localrunner is the optional full-fidelity backend for Gopher
// Workplace. It runs the *real* Go toolchain (go test, -race, benchmarks,
// GC/finalizer timing) against submitted solutions — things the in-browser
// yaegi-wasm interpreter cannot do.
//
// The static site detects it via GET /health and, when present, routes every
// Run/Submit to POST /run instead of the wasm path. When absent, the site
// falls back to wasm (junior/middle only).
//
// Usage:
//
//	go run ./site/cmd/localrunner              # serves :7070, auto-finds challenges/
//	go run ./site/cmd/localrunner -port 9090
//	GW_RUNNER_PORT=9090 GW_ROOT=/path/to/repo go run ./site/cmd/localrunner
//
// SAFETY: this executes arbitrary user Go code on your machine. It is meant for
// localhost only. Each request runs in a throwaway temp module with a 20s
// timeout (process group killed on expiry), network disabled (GOPROXY=off), and
// an isolated GOCACHE/GOPATH. Do not expose this port to untrusted networks.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const version = "0.1.0"

// config holds resolved server settings.
type config struct {
	port  string
	root  string // repo root: the dir containing challenges/
	dbDSN string
}

func main() {
	var (
		portFlag = flag.String("port", envOr("GW_RUNNER_PORT", "7070"), "listen port")
		rootFlag = flag.String("root", os.Getenv("GW_ROOT"), "repo root (dir containing challenges/); auto-detected if empty")
		dbFlag   = flag.String("db", os.Getenv("GW_DB"), "sqlite db path (default ~/.gopher-workplace/runner.db)")
	)
	flag.Parse()

	root, err := findRoot(*rootFlag)
	if err != nil {
		log.Fatalf("cannot locate repo root: %v", err)
	}
	dbPath := *dbFlag
	if dbPath == "" {
		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, ".gopher-workplace")
		_ = os.MkdirAll(dir, 0o755)
		dbPath = filepath.Join(dir, "runner.db")
	}

	cfg := config{port: *portFlag, root: root, dbDSN: dbPath}

	store, err := openStore(dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer store.close()
	store.startRetention(30 * 24 * time.Hour) // keep submissions 30 days

	srv := &server{cfg: cfg, store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", srv.withCORS(srv.handleHealth))
	mux.HandleFunc("/run", srv.withCORS(srv.handleRun))
	mux.HandleFunc("/vet", srv.withCORS(srv.handleVet))
	mux.HandleFunc("/fmt", srv.withCORS(srv.handleFmt))
	mux.HandleFunc("/history", srv.withCORS(srv.handleHistory))
	mux.HandleFunc("/solved", srv.withCORS(srv.handleSolved))

	// Serve the static site from this same server so the whole app is one
	// origin: open http://localhost:PORT/playground.html and Run executes here.
	webDir := filepath.Join(root, "site", "web")
	if fi, err := os.Stat(webDir); err == nil && fi.IsDir() {
		mux.Handle("/", http.FileServer(http.Dir(webDir)))
		log.Printf("  web:   %s", webDir)
	} else {
		log.Printf("  web:   (not found at %s — static site not served)", webDir)
	}

	log.Printf("gopher-workplace local runner %s", version)
	log.Printf("  root:  %s", root)
	log.Printf("  db:    %s (30-day retention)", dbPath)
	log.Printf("  listening on http://localhost:%s", cfg.port)
	log.Fatal(http.ListenAndServe(":"+cfg.port, mux))
}

type server struct {
	cfg   config
	store *store
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

// withCORS wraps a handler with permissive CORS for the local static site and
// the deployed origin, and answers preflight requests.
func (s *server) withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
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

// allowedOrigin permits localhost (any port) and the deployed origin. Any other
// origin still gets a wildcard (see withCORS) because no credentials are used.
func allowedOrigin(o string) bool {
	if o == "" {
		return false
	}
	if strings.HasPrefix(o, "http://localhost") ||
		strings.HasPrefix(o, "http://127.0.0.1") ||
		strings.HasPrefix(o, "https://localhost") {
		return true
	}
	return o == "https://gopher-workplace.netlify.app"
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
