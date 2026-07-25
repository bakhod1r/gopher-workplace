# Gopher Workplace — Challenge Plan

Practice repo mirroring the Go roadmap as broken-code-fix **puzzles**.

## Structure

```
challenges/<level>/<NN-topic>/<NN-subtopic>/<puzzle>/
                                              ├── go.mod
                                              ├── Makefile
                                              ├── <pkg>.go        # from-scratch stub, or one planted bug (§5e)
                                              ├── <pkg>_test.go   # red until fixed
                                              └── README.md       # story + task + topics to master
```

- **level** ∈ `junior middle senior staff` — the TOP dimension.
- Each subtopic is assigned to a level by **competency**, not by quartile — the
  level whose profile (see [GENERATION.md](GENERATION.md) §5-profiles) the
  concept belongs to. So a topic spans **only** the levels its subtopics
  actually need: foundational topics (language-basics) collapse into junior,
  deep topics (runtime-internals, advanced) sit at senior/staff. The grid is
  **ragged** — levels do NOT all have 17 topics.
- A subtopic lives under exactly one level; the puzzle(s) live inside it.
- **Topics are renumbered from 01 within each level**, contiguous (no gaps),
  in global learning-path order. So the same topic can carry a different NN per
  level (error-handling = 03 junior, 02 middle, 01 staff). Subtopics likewise
  **renumber from 01 within each (level, topic)**, roadmap order preserved.
- Learning-path total order is `level → topic (global order) → subtopic MM`.
  Level dominates, and per-level topic numbering follows the global order, so
  the order stays sound; the table below is the canonical global sequence.
- Empty subtopics hold a `.keep`; authoring a puzzle removes it.
- Regenerate anytime: `bash scripts/reconstruct.sh` (explicit J/M/S/T level map
  in the script; idempotent, preserves existing puzzles).

## Level split (competency-based, ragged)

Assignment by level profile, not quartile. `07-concurrency/00-introduction`
dropped (not a coding puzzle). Totals below.

| Level | topics present | subtopics |
|-------|----------------|-----------|
| 🟢 junior | 8  | 35 |
| 🔵 middle | 12 | 70 |
| 🟠 senior | 17 | 80 |
| 🔴 staff  | 9  | 37 |
| **total** |    | **223** |

Topics per level:

- 🟢 **junior:** language-basics, methods-and-interfaces, error-handling,
  code-organization, testing-and-benchmarking, standard-library, go-toolchain,
  modern-language-features.
- 🔵 **middle:** + generics, concurrency, web-development, design-patterns,
  observability (and more subtopics of the junior topics).
- 🟠 **senior:** + advanced-topics, performance-engineering, runtime-and-internals,
  go-source-reading, webassembly (nearly every topic reaches senior).
- 🔴 **staff:** error-handling, generics, testing, concurrency, advanced-topics,
  performance-engineering, runtime-and-internals, go-source-reading, webassembly
  — only the topics with genuine internals/memory-model depth.

## Topic order (learning path)

Reordered so dependencies flow: fundamentals → interfaces → errors → stdlib →
organization → toolchain → testing → modern syntax → generics → concurrency,
then the deep/internals tail. Topic NN is the learning-path key.

| # | Topic | Roadmap source | subs |
|---|-------|----------------|------|
| 01 | language-basics | 02-language-basics | 7 |
| 02 | methods-and-interfaces | 03-methods-and-interfaces | 19 |
| 03 | error-handling | 05-error-handling | 13 |
| 04 | standard-library | 08-standard-library | 21 |
| 05 | code-organization | 06-code-organization | 9 |
| 06 | go-toolchain | 10-go-toolchain | 11 |
| 07 | testing-and-benchmarking | 09-testing-and-benchmarking | 17 |
| 08 | modern-language-features | 18-modern-language-features | 5 |
| 09 | generics | 04-generics | 17 |
| 10 | concurrency | 07-concurrency | 28* |
| 11 | web-development | 19-web-development (added) | 9 |
| 12 | design-patterns-in-go | 13-design-patterns-in-go | 20 |
| 13 | observability | 17-observability-and-runtime-introspection | 5 |
| 14 | performance-engineering | 12-performance-engineering | 7 |
| 15 | advanced-topics | 11-advanced-topics | 16 |
| 16 | runtime-and-internals | 14-runtime-and-internals | 6 |
| 17 | go-source-reading | 15-go-source-reading | 6 |
| 18 | webassembly | 16-webassembly-and-alternative-targets | 5 |

*concurrency 28 = roadmap 26 − `00-introduction` (dropped) + `select-and-buffering`,
`worker-pools`, `race-detection` (added from current roadmap.sh).

## Current state

- 4 levels · ragged grid · **223** subtopics total (36/70/80/37)
- 4 authored puzzles (the `slices` exemplar), inside
  `junior/01-language-basics/03-composite-types/{dedupe, chunk, leak, collect}`
- rest are `.keep` slots — to be authored

## Authoring a puzzle

1. Pick a `.keep` subtopic. Copy the shape of a `slices` puzzle.
2. `go mod init github.com/gopher-workplace/challenges/<level>/<topic>/<subtopic>/<name>`
3. Intended behaviour in the doc comment; stage the red state per mode
   (GENERATION.md §5e): from-scratch stub (`panic("not implemented")`) for
   junior/concept puzzles, or one planted bug between
   `// CHANGE CODE BELOW/ABOVE THIS LINE` for senior/staff/debugging puzzles.
4. Drive it from `<pkg>_test.go` (red until fixed).
5. Reuse the standard `Makefile` (`verify` = fmt-check + vet + test).
6. Delete the subtopic's `.keep`.

## Run

```bash
make -C challenges/junior/01-language-basics/03-composite-types/chunk verify
make verify        # every authored puzzle (root Makefile)
```
