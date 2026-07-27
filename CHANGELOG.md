# Changelog

All notable changes to this project are documented here. Dates are ISO 8601.

## [0.3.0] — 2026-07-27

Teaching material next to every puzzle, three new constants puzzles, and one
fewer mode to explain.

### Added

- **An `education` tab beside the description.** Each puzzle may ship an
  `EDUCATION.md` — the concept explained properly: the idea, why it matters, the
  traps, and a short "try it yourself". The catalog generator renders it and the
  playground shows it as a second tab; puzzles without one simply hide the tab.
  All ten `01-variables-and-constants` puzzles now have it.
- The education text is **not copyable** — `user-select: none` plus blocked
  `copy`/`cut`/`contextmenu`/`dragstart`. It is there to be read and retyped, not
  lifted. The description stays copyable.
- **Three puzzles**, taking `junior/01-language-basics/01-variables-and-constants`
  to ten, each on a concept the other seven do not cover:
  - `typedconst` — typed vs untyped constants, and why conversion *direction*
    decides correctness (`byte(256)` wraps to 0).
  - `discard` — the blank identifier: why Go forces you to receive every value,
    and how `_` discards one.
  - `endpoint` — `const` vs package-level `var`, deriving a value instead of
    pasting it.

### Changed

- **The learn ⇄ debug toggle is gone.** Every puzzle now opens on its stub and is
  built from scratch. The editor is simpler and there is one less mode to
  explain. Draft storage keys lose their mode suffix.
- Markdown rendering understands `*italic*`.

### Fixed

- Built binaries (`gencatalog`, `localrunner`, `server`) are gitignored instead
  of landing in a commit.

[0.3.0]: https://github.com/bakhod1r/gopher-workplace/releases/tag/v0.3.0

## [0.2.0] — 2026-07-27

The runner grew up: it is now safe to leave running, it serves the UI itself,
and the tooling around it is tested and gated in CI.

### Security

- **The runner binds `127.0.0.1` instead of every interface.** It previously
  listened on `0.0.0.0`, so anyone on the same network — a café, an office, a
  hotel — could `POST /run` and execute arbitrary Go code as the user running
  it. `-host`/`GW_HOST` can still change the bind address, and prints a loud
  warning when the result is not loopback.
- **Cross-origin browser requests are refused.** Any origin that was not on a
  short allow-list still received `Access-Control-Allow-Origin: *`, which let
  any page a user happened to visit drive `/run` in the background. Non-loopback
  origins now get `403`, and the allow-list is anchored so hosts like
  `localhost.attacker.com` no longer match.
- **Symlinks in `challenges/` cannot escape it.** The path check was lexical, so
  it could not see a symlink inside `challenges/` pointing at, say, `/etc`.
  Resolved paths are now compared; an in-tree symlink still works.
- **Concurrent toolchain runs are capped** (2–4, by CPU count) and excess
  requests are shed with `503` instead of piling up `go test -race` processes.
- Documentation no longer claims this is a sandbox. It is not: submissions run
  with the user's file access and network. `GOPROXY=off` blocks module
  downloads, not sockets.

### Changed

- **The runner serves the web UI.** One process, one port, one origin — no
  separate static server, no port to keep in sync. `make dev` is the whole
  thing; the UI lives at <http://localhost:7070>.
- **The catalog generator is Go, not Python** (`site/cmd/gencatalog`). Output is
  byte-identical to the script it replaces. Go is now the only hard requirement
  to install and run the project.
- **The published site is explicitly browse-only.** With no runner detected, the
  UI shows a banner explaining that code runs on your own machine, disables Run
  and Submit, and links to the setup instructions instead of failing on click.
- Failure output in the UI reads correctly for the common Go idiom
  `Call(args) = X, want Y`; the got/want columns were previously empty for
  nearly every puzzle, which only matched `got X want Y`.
- `netlify.toml` and `web/_headers` no longer reference the wasm build that the
  project dropped.

### Added

- `make setup`, `make dev`, `make update`, `make catalog`, `make site-test` —
  and `make help` lists them. Install, configuration, and troubleshooting are
  documented in the README.
- A tool check that reads the required Go version out of the module, so it stays
  honest as the project moves.
- CI (GitHub Actions): gofmt, vet, tests, a **100% statement coverage gate** on
  both tooling modules, a stale-catalog check, and a per-challenge build.

### Notes

- Puzzles ship red on purpose. A failing `make verify` on an unsolved puzzle is
  the expected state, not a broken install.
- Solve history lives at `~/.gopher-workplace/runner.db`, outside the repo, and
  survives updates. It is swept after 30 days.

[0.2.0]: https://github.com/bakhod1r/gopher-workplace/releases/tag/v0.2.0
