<p align="center">
  <img src="assets/logo.png" alt="Gopher Workplace" width="420">
</p>

<h1 align="center">Gopher Workplace</h1>

<p align="center"><em>Practice · Build · Grow</em></p>

<p align="center">
  A practice repo mirroring the Go roadmap as <strong>broken-code-fix puzzles</strong>.<br>
  Drive each puzzle red&nbsp;→&nbsp;green — implement a stub from scratch, or fix one planted bug.
</p>

---

## What is this

Every puzzle is its own Go module. You start from a failing state and make the
tests pass — either by:

- **implement-from-scratch** — flesh out a `panic("not implemented")` stub
  (junior / concept puzzles), or
- **fix one planted bug** — the single defect sits between
  `// CHANGE CODE BELOW/ABOVE THIS LINE` markers (senior / staff / debugging).

Puzzles follow the Go learning path in a strict total order: `level → topic →
subtopic`. A puzzle may only rely on concepts introduced **at or before** its
position — so the difficulty rises honestly as you climb.

## Levels

```
junior  <  middle  <  senior  <  staff
```

4 levels · 209 subtopic slots (35 / 62 / 75 / 37). Three axes rise with level:

| Axis | junior | staff |
|------|--------|-------|
| **scope** — concepts allowed | language basics | full language |
| **difficulty** — resource pressure | correctness only | CPU/time ceilings, race-free concurrency |
| **depth** — where the bug lives | surface | memory model |

## Layout

```
challenges/<level>/<NN-topic>/<MM-subtopic>/<puzzle>/
    go.mod  Makefile  <pkg>.go  <pkg>_test.go  README.md
```

Grid is **ragged** — a level only holds topics its subtopics need. Empty slots
hold a `.keep`; authoring a puzzle deletes it. `_template/` is the shape to copy.

Full spec: [challenges/PLAN.md](challenges/PLAN.md) (structure) and
[challenges/GENERATION.md](challenges/GENERATION.md) (authoring law).

## How to work a puzzle

Each puzzle is one Go module — solve it red → green:

1. **Pick one.** `make list` shows every module. `cd` into one.

   ```bash
   cd challenges/junior/01-language-basics/01-variables-and-constants/plan-limits
   ```

2. **Read the brief.** `README.md` in the folder = Context / Task / Examples /
   Topics to Master / hint.

3. **Confirm it's red.** Tests fail out of the box.

   ```bash
   make test
   ```

4. **Fix it.** Open the `<pkg>.go` file and either:
   - fill in the `panic("not implemented")` stub, or
   - fix the single planted bug between the
     `// CHANGE CODE BELOW/ABOVE THIS LINE` markers.

   Don't touch the function signature or the test file.

5. **Go green.** Full gate = fmt-check + vet + test.

   ```bash
   make verify      # PASS: challenge validated
   ```

Handy per-puzzle targets: `make test-v` (verbose), `make fmt` (format),
`make vet`, `make clean` (drop test cache).

## Run everything

```bash
make list          # list all challenge modules
make verify        # verify every authored puzzle (fmt-check + vet + test)

# single puzzle from repo root
make -C challenges/<level>/<topic>/<subtopic>/<name> verify

bash scripts/reconstruct.sh   # regenerate grid (idempotent, keeps puzzles)
```

## Play in the browser

A LeetCode-style web UI ships in [site/](site/): pick a puzzle from the sidebar,
edit in-page, and run the real test suite. Every puzzle runs against the genuine
Go toolchain via a small localhost backend (`site/cmd/localrunner`) — no wasm
sandbox, so all levels (including `-race` and GC-sensitive ones) work. Solve
history is kept in SQLite so your "submitted" state survives reloads.

```bash
# 1. regenerate the problem catalog from the challenges/ tree
bash site/scripts/build.sh

# 2. start the local runner (serves :7070, auto-finds challenges/)
go run ./site/cmd/localrunner

# 3. open the UI
open site/web/index.html          # or serve site/web/ however you like
```

Editor shortcuts: **Ctrl/Cmd+Enter** run · **Ctrl/Cmd+S** or
**Ctrl/Cmd+Shift+Enter** submit · **Ctrl/Cmd+/** toggle comment. Re-run
`build.sh` whenever you add or edit a puzzle so the catalog picks it up.

## Authoring a puzzle

1. Pick a `.keep` slot, get its allowed history:
   `scripts/coverage.sh <level>/<NN-topic>/<MM-subtopic>`
2. Scaffold: `scripts/scaffold.sh <slot> <name> <Func>`
3. Stage the red state (stub **or** one planted bug), don't touch the signature
   or the task's tests.
4. Verify red → green. Delete the `.keep`.

See [CLAUDE.md](CLAUDE.md) for the short orientation and rules of engagement.
