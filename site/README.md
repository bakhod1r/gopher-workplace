# Gopher Workplace — Site

The web UI for the puzzles: pick one from the sidebar, edit in-page, run the
real test suite. There is no in-browser interpreter — every Run and Submit is
executed by [`cmd/localrunner`](cmd/localrunner), a small localhost server that
drives the genuine Go toolchain. That is what makes `-race`, GC/finalizer
timing, and benchmarks work at all.

The runner also serves this directory, so the page and the API share one origin.
From the repo root:

```bash
make dev          # http://localhost:7070
```

## Structure

```
site/
├── cmd/
│   ├── localrunner/            # Go — the backend (its own module)
│   │   ├── main.go             # config, routes, CORS, static file serving
│   │   ├── run.go              # materialize a puzzle, go test, parse -json
│   │   ├── tools.go            # /vet, /fmt, /history
│   │   ├── limit.go            # concurrency cap for toolchain work
│   │   └── store.go            # SQLite: submissions, solved set, retention
│   └── gencatalog/             # Go — generates web/assets/js/problems.js
├── web/                        # the frontend (served by the runner)
│   ├── index.html  playground.html  problemset.html  roadmap.html
│   └── assets/
│       ├── css/styles.css      # themes + layout
│       └── js/
│           ├── problems.js     # GENERATED catalog — do not edit by hand
│           ├── app.js          # editor, run/submit, custom test, drawer
│           └── localrunner.js  # backend bridge: /health probe, /run, /solved
├── server/                     # optional standalone static file server
└── scripts/
    ├── build.sh                # regenerate problems.js (wraps gencatalog)
    └── serve.sh                # static-only server, no backend (Python)
```

## Develop

```bash
make dev                  # from the repo root: runner + UI on :7070
make catalog              # regenerate problems.js after adding/editing a puzzle
make site-test            # vet + test the runner and the generator

./scripts/serve.sh 8145   # static files only, no backend (frontend work)
```

`problems.js` is generated from the `challenges/` tree — edit the puzzle's
`README.md` and starter file, then re-run `make catalog`. CI fails if the
committed catalog is stale.

## Backend API

All endpoints are on the runner's port; anything not matching is served from
`web/`.

| Endpoint | Purpose |
|----------|---------|
| `GET /health` | liveness probe; the nav badge is driven by this |
| `POST /run` | `{challengeId, src \| files, submit}` → Report JSON |
| `POST /vet` | same body, runs `go vet` |
| `POST /fmt` | `{src}` → gofmt'd source |
| `GET /solved` | challenge ids with at least one accepted Submit |
| `GET /history?challengeId=` | recent submissions |

A Submit only counts as solved when it passes **and** is warning-free — guard
tests emit `WARN:` lines for hardcoded answers, and non-gofmt-clean code is
flagged the same way.

**Persistence:** SQLite at `~/.gopher-workplace/runner.db` (`-db`/`GW_DB`),
kept forever by default (`-retention 720h` opts into pruning, swept on startup
and hourly). Old level-first challenge ids are rewritten to the current
topic-first layout at startup, so a tree reshuffle never loses solves. The driver is
`modernc.org/sqlite` (pure Go, no cgo) — the one dependency outside stdlib.

**Security:** the runner executes arbitrary Go code as you. It binds `127.0.0.1`,
refuses cross-origin browser requests, caps concurrent runs, and kills a run's
whole process group after 20s. It is **not** a sandbox — see the [security
note](../README.md#security-note) in the root README.

## Add a puzzle

Author it under `challenges/` (see
[challenges/GENERATION.md](../challenges/GENERATION.md)), then `make catalog`.
Nothing in `site/` needs editing: the sidebar, description, starter code, and
level all come from the puzzle's files.

## Themes

hacker (default) · matrix · amber · synthwave · dracula · one dark · paper.
Persisted in `localStorage` under `gw-theme`.
