# Changelog

All notable changes to this project are documented here. Dates are ISO 8601.

## [0.6.0] — 2026-07-29

Every puzzle in `01-language-basics` now teaches like LeetCode — a full worked
solution beside three-plus examples — and the problemset filters by multiple
values at once.

### Changed

- **All 450 `01-language-basics` puzzles rewritten to a LeetCode-style
  `EDUCATION.md`**: five fixed sections — **Intuition / Approach / Solution /
  Walkthrough / Pitfalls** — where **Solution** is the complete, verified-green
  code (the stub implemented, or the one planted bug fixed), with the
  `// CHANGE CODE` markers and bug-narration comments stripped. The retired shape
  (The idea / Why it matters / Watch out / Try it yourself) is gone.
- **Every README now carries ≥3 LeetCode-style examples** (`Input:` / `Output:` /
  optional `_Explanation:_`), covering the interesting edges. Covers
  `01-variables-and-constants` (40), `02-data-types` (70), `03-composite-types`
  (100), `04-functions` (120), `05-pointers` (120).
- Each puzzle's Solution was re-derived and **verified red→green + gofmt-clean**;
  puzzles still ship RED (the stub/planted bug is untouched).
- **Multi-select problemset filters.** Level, Tag, and Status are now checkbox
  dropdowns — filter by several values at once — replacing the single-choice
  selects.
- `GENERATION.md` §5g documents the required README-examples and EDUCATION shape.

### Fixed

- `02-data-types/junior/vowelcount` test asserted `Vowels("café") == 2`, but `é`
  is not an ASCII vowel (the test's own comment said so) — corrected to `1`.
- `03-composite-types/junior/checkoutgrid` module path carried a stale
  `01-arrays/` segment; flattened to match its location.

### Notes

- Version bumped to 0.6.0; asset cache-buster to `?v=27`; `problems.js`
  regenerated (450 puzzles).

## [0.5.0] — 2026-07-28

The `05-pointers` topic lands complete — **120 puzzles**, 30 per level — in the
same flat, level-only layout as `04-functions`.

### Added

- **`01-language-basics/05-pointers` — 120 new puzzles** (junior 30, middle 30,
  senior 30, staff 30) spanning pointer basics, pointers with structs, pointers
  with maps and slices, memory management, `unsafe.Pointer`, and nil-pointer
  dereference:
  - **junior (30)** — deref/address-of, increment/negate/toggle/clamp through a
    pointer, swap via pointers, allocation, nil-safe deref, slice/map of pointers,
    struct field mutation, double pointers, linked-list length, pointer identity.
  - **middle (30)** — pointer vs value receivers, linked-list ops (prepend,
    reverse, append via `**Node`, delete, middle via slow/fast, dedup, concat,
    nth-from-end, cycle detection, merge, rotate), BST insert/search/height/leaves/
    sum/mirror, method values, deep copy, slice-pointer growth, `new` vs `&T{}`.
  - **senior (30, planted-bug)** — value-receiver mutation loss, delete-head
    ignoring the new head, swapping pointers not values, discarded recursive
    returns, missing slow/fast nil guard, deref-before-nil-check, reversal order,
    double-pointer reassignment, shallow struct/list/tree copies, gap off-by-one,
    merge remainder, hoisted-variable address, stale element pointer, identity vs
    value confusion, nested-map init, nil-out-on-shrink, and more.
  - **staff (30, planted-bug, internals)** — struct field-order padding
    (`unsafe.Sizeof`), `unsafe.Add` stride, sub-slice/`s[:k:k]` retention, pointer
    vs pointee size, slice/string/array data pointers, `Offsetof` vs `Sizeof`,
    `Alignof`, `unsafe.Slice` length, reinterpret width/layout, method
    expressions, per-iteration allocation, deep pointer-field copy, `delete` vs
    nil, `clear`, pool reset, and nil `unsafe.Pointer`/`SliceData` guards.
  - Each ships an `EDUCATION.md`. Junior/middle are implement-from-scratch stubs
    (red → green); senior/staff are single-planted-bug puzzles between
    `// CHANGE CODE …` markers.
- Every puzzle verified: `gofmt` clean, `go vet` passes (0 issues); stubs and
  planted bugs each go red → green with the documented fix.

### Changed

- Regenerated `challenges/HIERARCHY.md` and `site/web/assets/js/problems.js`
  (now 453 puzzles). Version bumped to 0.5.0; asset cache-buster to `?v=24`.

## [0.4.0] — 2026-07-28

The `04-functions` topic lands complete — **120 puzzles**, 30 per level — and the
topic adopts a flat level-only layout.

### Added

- **`01-language-basics/04-functions` — 120 new puzzles** (junior 30, middle 30,
  senior 30, staff 30), covering functions, closures/defer, and control flow
  together:
  - **junior (30)** — multiple return, variadic sum/join/average, div-mod,
    clamp, min/max, for-range sums, countdown, FizzBuzz, 2D grid sum, reverse,
    GCD, sign/day/parity switches, leap year, abs, call-by-value and pointer
    basics.
  - **middle (30)** — closures (counter, adder, memoize, once, compose, curry,
    pipeline, group-by, take/drop-while, zip-with, shared-state tracker,
    per-iteration capture), higher-order map/filter/reduce, `defer` ordering,
    argument-snapshot timing, defer-mutates-named-return, defer-in-loop,
    `recover`, named returns, variadic forwarding to `append`, labeled break,
    `goto`, and `fallthrough`.
  - **senior (30, planted-bug)** — shadowed named-return error, recover outside
    defer, defer argument timing, sliding-window/look-ahead/binary-search bounds,
    accidental `fallthrough`, nil-map write, recursion base case, stale element
    pointer after realloc, append clobbering shared capacity, missing `default`,
    inverted `continue` guard, shared-loop-variable capture, defer overwriting a
    result, boundary comparison, all/any early return, range-copy no-op mutation,
    ordered delete off-by-one, rune vs byte length, `copy` bounded by length,
    discarded `append` result, missing spread, slice-alias mutation, signed-modulo
    parity, nested-map init, two-pointer index bound, defer-fires-at-exit, prefix
    sum, insert shift direction.
  - **staff (30, planted-bug, internals)** — selective recover/re-panic,
    defer-wraps-the-snapshot error idiom, panic-during-unwind masking, shared
    closure cell, accumulator reset in recursion, method-value receiver binding,
    recovered error lost to a local, two-append shared capacity, stale element
    pointer, `s[:k:k]` capacity retention, deferred LIFO push, uint8 accumulator
    overflow, self-referential loop bound, bare-return with deferred adjust,
    defer-before-acquisition, short-circuit nil guard, labeled break vs continue,
    switch init statement, nil callback guard, comma-ok, deferred slice-header
    snapshot, append-to-nil, overlapping copy direction, typed recover assertion,
    reused-buffer aliasing, simultaneous assignment, make length vs capacity,
    rune indexing, per-iteration defer scope, immediately-invoked init.
  - Each ships an `EDUCATION.md`. Junior/middle are implement-from-scratch stubs
    (red → green); senior/staff are single-planted-bug puzzles between
    `// CHANGE CODE …` markers.
- Every puzzle verified: `gofmt` clean, `go vet` passes; stubs fail red and pass
  green on the reference implementation; each planted bug fails red and passes
  green once the documented one-line fix is applied.

### Changed

- **`04-functions` uses a flat, level-only layout** — `04-functions/<level>/<name>`
  — instead of per-subtopic nesting. The catalog generator already derives the
  level from the slug and the subtopic from the README, so grouping is unchanged.
- The five pre-existing `conditionals` puzzles moved under `04-functions/junior/`.
- Regenerated `challenges/HIERARCHY.md` and `site/web/assets/js/problems.js`
  (now 333 puzzles). Asset cache-buster bumped to `?v=23`.

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
