# Gopher Workplace — project guide

Practice repo mirroring the Go roadmap as **broken-code-fix puzzles**. Each
puzzle is its own Go module the learner drives red→green — either by
implementing a stub from scratch (junior/concept) or fixing one planted bug
(senior/staff); see GENERATION.md §5e.

Full spec lives in [challenges/PLAN.md](challenges/PLAN.md) (structure) and
[challenges/GENERATION.md](challenges/GENERATION.md) (authoring law). Read both
before authoring. This file is the short orientation.

## Layout

```
challenges/<level>/<NN-topic>/<MM-subtopic>/<puzzle>/
    go.mod  Makefile  <pkg>.go  <pkg>_test.go  README.md
```

- Levels (outermost, ordered): `junior < middle < senior < staff`.
- Grid is **ragged** — a level only holds topics its subtopics need.
- Topics + subtopics **renumber from 01 within each level** in learning-path order.
- Empty slots hold a `.keep`; authoring a puzzle deletes it.
- `_template/` = the shape to copy.
- 4 levels · 209 subtopic slots (35/62/75/37). Authored so far, all under
  `junior/01-language-basics`: `01-variables-and-constants` (swap, byteunits,
  temperature, retries, plan-limits), `02-data-types` (average, checksum,
  truncate, runecount, narrowing, almostequal), `03-composite-types` (dedupe,
  wordfreq, reverse). Most ship a from-scratch stub plus a `<name>.debug.txt`
  planted-bug variant (see GENERATION.md §5e–5f).

## The one law (scope, GENERATION.md §2)

Learning path is a total order: `level → topic → subtopic`. A puzzle at position
`P` may only rely on Go concepts introduced **at or before `P`**. The target
subtopic is the concept taught (must be exercised); everything before is fair
game; everything after is forbidden. Level fully precedes the next, so a middle
puzzle may use every junior topic.

Three independent axes rise with level:
- **scope** §2 — which concepts allowed.
- **difficulty dial** §5b — resource/scale pressure (junior=correctness only →
  staff=CPU/time ceiling + race-free concurrency). Encode as graded tests, not prose.
- **depth** §5d — how far under the language the bug lives (surface → memory model).

Also §5c **cumulative rule**: weave in ≥1–2 earlier covered concepts, listed in
README **Topics to Master**.

## Authoring a puzzle

1. Pick a `.keep` slot. Get its allowed history:
   `scripts/coverage.sh <level>/<NN-topic>/<MM-subtopic>`
2. Scaffold: `scripts/scaffold.sh <slot> <name> <Func>` — copies `_template/`,
   injects the covered set as a `COVERED:` block.
3. Module path: `github.com/gopher-workplace/challenges/<level>/<topic>/<subtopic>/<name>`
4. Stage the red state by the puzzle's mode (GENERATION.md §5e): **implement-
   from-scratch** stub (`panic("not implemented")`, no `CHANGE CODE` markers) for
   junior/concept puzzles, or **one planted bug** between `// CHANGE CODE
   BELOW/ABOVE THIS LINE` for senior/staff/debugging puzzles. Don't change the
   function signature or the task's tests.
5. Verify red→green. Delete the `.keep`.

Conventions: one mode per puzzle; each puzzle own `go.mod` (`go 1.26`); README has
Context / Task / Examples(≥3) / Topics to Master / hint; no `SOLUTION.md`; no
roadmap-path references in READMEs.

## Run

```bash
make -C challenges/<level>/<topic>/<subtopic>/<name> verify   # one puzzle
make verify        # every authored puzzle (verify = fmt-check + vet + test)
make list          # list all challenge modules
bash scripts/reconstruct.sh   # regenerate grid (idempotent, keeps puzzles)
```

## Rules of engagement

- Never run `git commit`, `git push`, or `rm` (global limit).
- `bypassPermissions` default is on — act without asking.
