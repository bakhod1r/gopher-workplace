# 11 Performance Engineering — build plan (mirror of 03-generics)

Goal: turn `11-performance-engineering/` from 7 empty `.keep` subtopic slots
into the same shape as `03-generics/`: 4 level dirs, 50 puzzles each, 200 total.

## 1. Target layout

```
11-performance-engineering/
    TOPICS.md
    junior/<puzzle>/    50   implement-from-scratch stub
    middle/<puzzle>/    50   implement-from-scratch stub
    senior/<puzzle>/    50   one planted bug
    staff/<puzzle>/     50   one planted bug + difficulty dial (scale/alloc/time ceiling)
```

Subtopic dirs (`01-cpu-profiling` … `07-trace-tool`) are deleted; their content
is absorbed into blocks inside every level, exactly like 03-generics folded its
17 roadmap subtopics into 4 blocks.

Each puzzle dir holds 6 files, copied from `03-generics` shape:

| File | Notes |
|------|-------|
| `go.mod` | `module github.com/gopher-workplace/challenges/11-performance-engineering/<level>/<name>`, `go 1.26` |
| `Makefile` | identical to existing (`verify` = fmt-check + vet + test) |
| `<name>.go` | package doc, exported func, stub or `CHANGE CODE BELOW/ABOVE THIS LINE` block |
| `<name>_test.go` | deterministic tests; senior/staff add the bug-exposing test |
| `README.md` | Context / Task / Examples(≥3) / Topics to Master table / Hint / Validate |
| `EDUCATION.md` | Intuition / Approach / Solution / Walkthrough / Pitfalls |

## 2. Block map (7 roadmap subtopics → 4 blocks)

| Block | Covers | Roadmap subtopics |
|-------|--------|-------------------|
| 1 | Measuring | benchmarking-strategy, cpu-profiling |
| 2 | Memory | memory-profiling, pprof-deep |
| 3 | Contention | mutex-block-profiling, trace-tool |
| 4 | Optimizing | optimization-workflow |

Each level walks all 4 blocks in learning-path order (~12/13 puzzles per block).

## 3. Determinism rule (the one thing generics did not need)

Profilers and wall-clock timing are not reproducible in CI. So puzzles teach
performance *mechanics*, asserted deterministically:

- allocation counts via `testing.AllocsPerRun` (exact integer ceiling)
- `b.ReportAllocs` / benchmark code as the *subject under test*, correctness of
  the harness asserted by unit tests (`b.N` misuse, `b.ResetTimer` placement,
  loop-invariant hoisting, sink variables that stop dead-code elimination)
- pprof/trace puzzles operate on **parsed profile data structures** (flat/cum
  aggregation, self vs total, sample weighting, stack folding) — pure functions
  over fixture data, no live profiler
- contention puzzles assert observable behaviour (ordering, lost updates,
  `-race` clean) not timing
- staff dial = `AllocsPerRun` ceilings, O(n) vs O(n²) at scale with a wall-clock
  budget wide enough to be stable, and `-race`

Scope law (GENERATION.md §2): level 11 may use everything from 01–10, so
concurrency, reflection, unsafe-free runtime knowledge are all fair game.

## 4. Puzzle themes per level

- **junior (stub)** — write a correct benchmark loop, `ReportAllocs`, sink var,
  `ResetTimer`, `b.Run` subbenchmarks, byte/op math, flat-vs-cum sum over a
  fixture profile, top-N hot functions, prealloc `make([]T,0,n)`,
  `strings.Builder` vs `+=`, map prealloc, sample-count aggregation.
- **middle (stub)** — stack folding, self vs cumulative time, profile diffing,
  percentile latency, alloc-per-op accounting, escape-avoiding APIs (accept
  buffer arg), `sync.Pool` wrapper, batching, worker pool sizing, trace event
  span merging, block-profile aggregation by call site.
- **senior (planted bug)** — timer started before setup, `b.N` ignored or used
  as data size, result not sunk (dead-code elimination hides the work),
  double-count in cum aggregation, pool object not reset before reuse, slice
  reused across goroutines, `defer` in hot loop, wrong percentile index,
  block-profile rate treated as count, trace span end before start.
- **staff (planted bug + dial)** — quadratic append in a "prealloc" path, map
  key allocating per lookup, `sync.Pool` leaking capacity growth, false sharing
  in a counter array, lock held across a slow call, pooled buffer aliased after
  `Put`, profile merge losing samples at scale, O(n²) stack folding, race under
  `-race`, alloc ceiling failure that only shows at n=100k.

## 5. Execution order

1. Write `11-performance-engineering/TOPICS.md` (mirror of the generics one:
   block table, mode note, 4 name lists, progress table).
2. Fix the 200 puzzle names first, all 4 lists, before any code — names are the
   contract for the rest of the work.
3. Generate scaffolding with a script (`gen.sh` in this dir) that stamps the 6
   files per puzzle from the `03-generics` shape.
4. Author in batches of ~10 puzzles: code + tests + README + EDUCATION,
   `make -C <dir> verify` green after the fix, red before it.
5. Delete the 7 `.keep` subtopic dirs once junior block 1 lands; delete this
   file when all 200 are authored.
6. Final gate: `make verify` from repo root.

## 6. Batch checklist (per puzzle)

- [ ] red state correct (stub panics / bug fails exactly one named test)
- [ ] fix is a one-liner-ish, inside markers for senior/staff
- [ ] tests deterministic — no wall-clock assertion under 100ms of slack
- [ ] README has ≥3 examples and the Topics to Master table
- [ ] `make verify` passes on the fixed copy
