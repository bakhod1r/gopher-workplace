// Command server is a small production static file server for the Gopher
// Workplace site. It serves ./web, sets correct MIME types, and transparently
// serves pre-compressed *.gz assets (e.g. gopher.wasm.gz) with
// Content-Encoding: gzip so the ~38M wasm ships as ~8M.
//
// Usage:
//
//	go run ./server                 # serves ../web on :8080
//	PORT=9000 ROOT=web go run .     # override
package main

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := env("ROOT", "web")
	port := env("PORT", "8080")

	if abs, err := filepath.Abs(root); err == nil {
		root = abs
	}
	if _, err := os.Stat(root); err != nil {
		log.Fatalf("root %q not found: %v", root, err)
	}

	_ = mime.AddExtensionType(".wasm", "application/wasm")

	log.Printf("gopher-workplace serving %s on http://localhost:%s", root, port)
	log.Fatal(http.ListenAndServe(":"+port, handler(root)))
}

func handler(root string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(root))
	return func(w http.ResponseWriter, r *http.Request) {
		clean := filepath.Clean(r.URL.Path)
		full := filepath.Join(root, clean)

		// Serve a pre-gzipped twin when present and accepted.
		if strings.HasSuffix(clean, ".wasm") && acceptsGzip(r) {
			if _, err := os.Stat(full + ".gz"); err == nil {
				w.Header().Set("Content-Encoding", "gzip")
				w.Header().Set("Content-Type", "application/wasm")
				w.Header().Set("Vary", "Accept-Encoding")
				w.Header().Set("Cache-Control", "public, max-age=3600")
				http.ServeFile(w, r, full+".gz")
				return
			}
		}
		// Long cache for hashed-ish static assets; short for html.
		if strings.HasSuffix(clean, ".wasm") || strings.HasPrefix(clean, "/assets/") {
			w.Header().Set("Cache-Control", "public, max-age=3600")
		}
		fs.ServeHTTP(w, r)
	}
}

func acceptsGzip(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
