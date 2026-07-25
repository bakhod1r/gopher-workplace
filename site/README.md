# Gopher Workplace — Site

In-browser Go challenge runner. Pure static: Go code is interpreted client-side
via [yaegi](https://github.com/traefik/yaegi) compiled to WebAssembly — no
server, no backend.

## Structure

```
site/
├── runner/                     # Go — the wasm runner (its own module)
│   ├── go.mod
│   ├── runner.go               # yaegi interpret + suite/custom runners
│   └── cmd/
│       ├── wasm/main.go        # js/wasm bridge: gopherRunDedupe, ...Custom
│       └── prove/main.go       # native sanity harness
├── web/                        # static frontend (deploy this dir)
│   ├── index.html
│   └── assets/
│       ├── css/styles.css      # themes + layout
│       ├── js/
│       │   ├── problems.js     # problem catalog (data-driven)
│       │   ├── app.js          # editor, run/submit, custom test, drawer
│       │   └── wasm_exec.js    # Go wasm glue (copied from GOROOT)
│       └── wasm/
│           ├── gopher.wasm     # built runner (~38M, gz ~8M)
│           └── gopher.wasm.gz
└── scripts/
    ├── build.sh                # build wasm + stage assets
    └── serve.sh                # local static server
```

## Develop

```bash
./scripts/build.sh            # rebuild wasm into web/assets/
./scripts/serve.sh 8145       # quick python static server (no gzip)
./scripts/serve-prod.sh 8080  # Go server: gzip-serves wasm (~8M on the wire)
```

## Deploy (real site)

The site is fully static — deploy the `web/` directory anywhere. The wasm is
large, so **serve it gzipped**.

- **Own box / container:** run `server/` (Go, stdlib only). It serves `web/`,
  sets `application/wasm`, and streams `gopher.wasm.gz` with
  `Content-Encoding: gzip` when the client accepts it (40M → ~8M).
  ```bash
  cd site && ./scripts/build.sh && ./scripts/serve-prod.sh 8080
  ```
- **Netlify / Cloudflare Pages:** `netlify.toml` builds via `scripts/build.sh`
  and publishes `web/`; `web/_headers` sets wasm MIME + cache. These hosts
  gzip/brotli automatically.
- **GitHub Pages / plain CDN:** upload `web/`. Ensure the host sends
  `Content-Type: application/wasm` and gzip; otherwise the 40M raw file is
  downloaded uncompressed.

## Local runner (full fidelity)

The wasm interpreter can't do `go test -race`, GC/finalizer timing, or
benchmarks, so senior/staff puzzles (e.g. `leak`, `collect`) are locked in the
pure-static site. Start the **optional local runner** — a small stdlib HTTP
server that drives the *real* Go toolchain — and the site unlocks every level.

```bash
go run ./site/cmd/localrunner          # serves :7070, auto-finds challenges/
go run ./site/cmd/localrunner -port 9090
GW_RUNNER_PORT=9090 GW_ROOT=/path/to/repo go run ./site/cmd/localrunner
```

Run it from the repo root (it walks up to find `challenges/`). The browser
probes `GET /health` on load; when it answers, a `● local runner` badge appears
in the nav and **every** Run/Submit routes to it (real `go test`, `-race` added
automatically for concurrency puzzles). When it's absent, the site falls back to
wasm (junior/middle) and shows a “start it with: `go run ./site/cmd/localrunner`”
hint on locked puzzles. The Report JSON shape is unchanged, so results render
identically either way.

**What it unlocks:** senior/staff levels, `-race`, GC/finalizer/`SetFinalizer`
tests, benchmarks — the real `make verify` per challenge.

**Endpoints:** `GET /health`, `POST /run` (`{challengeId, src}` or
`{challengeId, files}`), `POST /vet`, `POST /fmt`, `GET /history?challengeId=`.

**Persistence:** submissions are stored in SQLite (`~/.gopher-workplace/runner.db`,
override with `-db`/`GW_DB`) and kept **30 days** (swept on startup + hourly).
The db driver is `modernc.org/sqlite` (pure Go, no cgo) — the one dependency
outside stdlib, contained to this module.

**Security — it runs arbitrary user Go code on your machine:**
- **localhost only** — never expose the port to untrusted networks.
- Each request runs in a throwaway `os.MkdirTemp` module, deleted after.
- 20s context timeout; the whole process group is `SIGKILL`ed on expiry (an
  infinite-loop submission is killed, not hung).
- Network disabled (`GOPROXY=off`, `GOSUMDB=off`); isolated `GOCACHE`/`GOPATH`.
- Paths are validated to stay inside `challenges/` (no traversal).

## Add a puzzle (pure-Go only)

1. Add suite/custom runner funcs in `runner/runner.go` and expose them in
   `runner/cmd/wasm/main.go`.
2. Add an entry to `window.PROBLEMS` and `window.CATALOG` in
   `web/assets/js/problems.js` (title, starter, description, runner fn names).
3. `./scripts/build.sh`.

## Limits

- yaegi interprets pure Go, no cgo. Junior/middle slice puzzles work.
- Senior (`runtime.SetFinalizer`/GC timing) and staff (`-race`) are **not**
  supported by the interpreter — start the [Local runner](#local-runner-full-fidelity)
  to unlock them with the real toolchain.
- `gopher.wasm` is large; host with `Content-Encoding: gzip` (serve the `.gz`).

## Themes

hacker (default) · matrix · amber · synthwave · dracula · one dark · paper.
Persisted in `localStorage` under `gw-theme`.
