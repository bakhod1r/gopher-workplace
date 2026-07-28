# Composite-Types — 100 Real-Work Puzzles (design)

**Date:** 2026-07-27
**Scope:** `challenges/01-language-basics/03-composite-types` — its 15 empty
subtopic slots.
**Goal:** author **100** puzzles that read like real production code, spread
across the 4 levels (junior/middle/senior/staff) by **per-slot competency**
(§2 scope-law, §5d depth ladder).

---

## 1. Level model — per-slot competency

A slot is authored at the level(s) its concept genuinely belongs to. Basics
(arrays, strings, maps, structs) are junior→middle; the `-internals` /
`memory-layout` slots carry the senior/staff depth. We do **not** force all 4
levels into every slot — that would violate §2 (a junior composite slot has no
senior concept available at its position).

Directory convention (matches existing `dedupe`):

```
03-composite-types/<slot>/<level>/<name>/{go.mod,Makefile,<pkg>.go,<pkg>_test.go,README.md}
module: github.com/gopher-workplace/challenges/01-language-basics/03-composite-types/<slot>/<level>/<name>
```

## 2. Level distribution (sums to 100)

| # | slot | junior | middle | senior | staff | total |
|---|------|:--:|:--:|:--:|:--:|:--:|
| 1 | 01-arrays | 4 | 3 | – | – | 7 |
| 2 | 02-slices/01-capacity-and-growth | 2 | 3 | 2 | – | 7 |
| 3 | 02-slices/02-make | 3 | 2 | – | – | 5 |
| 4 | 02-slices/03-slice-to-array-conversion | 2 | 3 | – | – | 5 |
| 5 | 02-slices/04-array-to-slice-conversion | 2 | 3 | – | – | 5 |
| 6 | 02-slices/05-slice-header-internals | – | 2 | 3 | 2 | 7 |
| 7 | 02-slices/06-slice-tricks | 3 | 4 | 2 | – | 9 |
| 8 | 03-strings | 3 | 5 | 1 | – | 9 |
| 9 | 04-maps/01-comma-ok-idiom | 3 | 3 | – | – | 6 |
| 10 | 04-maps/02-map-internals | – | 2 | 3 | 2 | 7 |
| 11 | 05-structs/01-struct-tags-and-json | 3 | 4 | – | – | 7 |
| 12 | 05-structs/02-embedding-structs | 2 | 4 | 1 | – | 7 |
| 13 | 05-structs/03-memory-layout | – | 1 | 3 | 3 | 7 |
| 14 | 06-empty-struct | 2 | 3 | 1 | – | 6 |
| 15 | 07-anonymous-structs | 3 | 3 | – | – | 6 |
| | **total** | **32** | **45** | **16** | **7** | **100** |

Average 25/level; real spread thins toward staff because composite-types is a
junior topic — only its internals reach the memory model.

## 3. Mode + difficulty per level (§5b, §5d, §5e)

| level | mode (§5e) | dial (§5b) | depth (§5d) |
|-------|-----------|-----------|-------------|
| junior | from-scratch stub | correctness, small input | surface / API |
| middle | mostly stub; some planted-bug | basic efficiency, no O(n²)/extra allocs | mechanics (backing array, cap growth) |
| senior | planted-bug | scale + RAM ceiling (stream, bounded allocs) | runtime / cost model (escape, GC) |
| staff | planted-bug | CPU/time ceiling + race-free concurrency | memory model / internals |

Junior & middle from-scratch puzzles **may** ship an optional `<name>.debug.txt`
sibling (§5f, learn⇄debug toggle) — not counted as separate puzzles.

Every puzzle: weave ≥1–2 earlier covered concepts (§5c), list them in README
**Topics to Master**. Guard tests (§5f) where a task is hardcode-cheatable.

## 4. Puzzle roster (real-work scenarios)

Each line: `name — scenario — target(+woven priors)`.

### slot 1 — 01-arrays
- j `checkoutgrid` — fixed 7×N seating map, mark taken seats — array value vs slice, indexing
- j `rgbahist` — `[256]int` color histogram from pixel bytes — array as fixed table, range
- j `weekload` — `[7]float64` per-weekday request load, peak day — array comparison, zero value
- j `boardhash` — compare two `[9]byte` tic-tac-toe boards for equality — arrays are comparable (== )
- m `rollingavg` — fixed-window `[N]float64` ring buffer of latencies — array + modulo index (+slice view)
- m `dnscache4` — `[4]byte` IPv4 key in a map — array as map key (+maps)
- m `matmul` — `[3][3]int` transform multiply — multidim arrays (+nested loops)

### slot 2 — 02-slices/01-capacity-and-growth
- j `growlog` — append events, report len vs cap after each — len/cap observation
- j `preallocbatch` — build N-row batch with correct `make(,0,N)` — prealloc avoids regrow
- m `appendaliasbug` — sub-batch shares backing, overwrites parent (planted-bug) — cap reuse aliasing (+slice header)
- m `dedupeinplace` — filter-in-place keeping cap — `s[:0]` reuse (+append)
- m `csvcols` — split rows, grow columns, no O(n²) copies — amortized growth
- s `streamdedup` — dedupe 100M-row stream under RAM ceiling (planted-bug: buffers all) — bounded alloc streaming (+map set)
- s `nocopygrow` — hot-path append that must not reallocate per call (bench ns/op) — cap budgeting (+escape)

### slot 3 — 02-slices/02-make
- j `zerobuf` — allocate zeroed `[]byte` frame of size n — make len semantics
- j `gridrows` — `make([][]int, rows)` then per-row make — 2D slice construction
- j `capvslen` — build slice with make(,len,cap), fill spare via append — make three-arg
- m `readnbytes` — pre-size read buffer, avoid over-alloc (planted-bug: make(,n) vs (,0,n)) — len vs cap trap
- m `poolframe` — reusable frame buffer `make` once, reset per iter — reuse over realloc (+capacity)

### slot 4 — 02-slices/03-slice-to-array-conversion
- j `ipv4octets` — `[]byte` → `[4]byte` for a fixed IPv4 — slice→array (Go 1.20+)
- j `sha256fix` — hash slice → `[32]byte` fixed digest — array pointer conversion
- m `headerframe` — parse fixed 8-byte header from stream slice (planted-bug: wrong len panic) — length-check before convert (+capacity)
- m `chunk16` — AES-block `[16]byte` blocks from payload — slice→array in loop (+slicing)
- m `magicbytes` — file signature `[4]byte` compare — array equality after convert (+arrays)

### slot 5 — 02-slices/04-array-to-slice-conversion
- j `arrslice` — expose `[N]int` config as read-only `[]int` — array→slice `[:]`
- j `bufwriter` — write into `arr[:]` view — slicing an array
- m `aliasarray` — mutate through slice reflects in array (planted-bug: expected copy) — array→slice aliases (+slice header)
- m `windowview` — sliding windows over a fixed `[N]byte` buffer — sub-slicing an array
- m `passbyslice` — pass big array to func cheaply via slice — avoid array copy cost (+make)

### slot 6 — 02-slices/05-slice-header-internals
- m `headerprint` — implement `len/cap/ptr`-style introspection via reslicing — header = ptr+len+cap
- m `truncleak` — truncated slice still pins backing array elems (planted-bug) — reslice keeps cap
- s `subslicealias` — API returns sub-slice, caller corrupts parent (planted-bug) — shared backing (+capacity)
- s `clipthree` — full-slice-expr `s[a:b:c]` to hand out safe sub-view (planted-bug: 2-index) — three-index slice
- s `gcpin` — small head-slice pins huge backing, memory leak (planted-bug) — copy to release (+GC reachability)
- st `racyappend` — concurrent appends to shared-backing sub-slices data-race (planted-bug, `-race`) — header not concurrency-safe (+goroutines)
- st `zerocopyparse` — parse under time ceiling with zero-copy sub-slices, no aliasing bug (planted-bug) — header cost model (+bench)

### slot 7 — 02-slices/06-slice-tricks
- j `insertat` — insert element at index i — insert idiom
- j `deleteat` — delete index i preserving order — delete idiom
- j `reverseinplace` — reverse in place — two-pointer swap
- m `filterinplace` — drop nils reusing backing (planted-bug: leaks pointers) — `s[:0]` + zero tail (+GC)
- m `dedupsorted` — dedupe a sorted slice in place — adjacent-compare
- m `movetofront` — LRU-style move element to front, no realloc — rotate trick (+capacity)
- m `batchsplit` — split into N-sized chunks lazily — chunk idiom (+make)
- s `bigdelete` — delete matching rows from 100M slice, bounded allocs (planted-bug) — in-place compact at scale
- s `stablepartition` — partition huge slice under RAM ceiling (planted-bug) — in-place, no aux buffer (+bench)

### slot 8 — 03-strings
- j `slugify` — title → url slug — string building, byte vs rune
- j `maskcard` — mask all but last-4 of a card number — indexing, concat
- j `countwords` — whitespace word count — range over string (bytes)
- m `runereverse` — reverse a UTF-8 string correctly (planted-bug: byte reverse mojibake) — runes vs bytes (+slices)
- m `builderjoin` — join fields via `strings.Builder`, no `+=` O(n²) — Builder over concat
- m `trimparse` — parse `key=value` lines, trim safely — slicing strings (+maps)
- m `bytestostr` — reuse `[]byte` ↔ string without extra copy where safe (planted-bug) — conversion cost
- m `csvquote` — escape quotes/commas per RFC4180 — byte scanning (+builder)
- s `bigconcat` — assemble 100MB report streaming, bounded allocs (planted-bug: naive `+`) — Builder growth at scale (+capacity)

### slot 9 — 04-maps/01-comma-ok-idiom
- j `featureflag` — flag lookup with default when absent — comma-ok read
- j `wordcount` — frequency map, increment-or-init — zero-value + comma-ok
- j `firstseen` — record first index of each key, ignore later — comma-ok guard
- m `sessionget` — distinguish "missing" vs "present but zero" TTL (planted-bug: uses `==0`) — comma-ok vs zero value
- m `graphadj` — adjacency map, lazy-init neighbor slices — comma-ok + append (+slices)
- m `dedupemap` — dedupe stream via `map[T]struct{}` set — set idiom (+empty-struct)

### slot 10 — 04-maps/02-map-internals
- m `iterorder` — code that wrongly relies on map range order (planted-bug) — unordered iteration
- m `deletewhileiter` — safe delete during range (planted-bug) — iteration + delete rules
- s `preallocmap` — `make(map,hint)` to cut rehash on 10M inserts (planted-bug: no hint, bench) — sizing, load factor
- s `nilmapwrite` — nil-map write panic in a lazy cache (planted-bug) — nil vs empty map (+comma-ok)
- s `keyalloc` — struct-key map avoiding per-lookup alloc (planted-bug, `-benchmem`) — hashing + escape (+structs)
- st `concurrentmap` — sharded map under concurrent load, race-free (planted-bug, `-race`) — map not concurrency-safe (+goroutines)
- st `syncmapswap` — swap plain map+mutex vs `sync.Map` under CPU ceiling (planted-bug) — internals + contention (+bench)

### slot 11 — 05-structs/01-struct-tags-and-json
- j `apiuser` — tag struct for exact JSON field names — json tags
- j `omitempty` — omit zero fields from payload — omitempty semantics
- j `renamekeys` — snake_case wire ↔ Go fields — tag mapping
- m `unmarshalpartial` — decode only known fields, ignore extras — decoder + tags (+maps)
- m `nestedconfig` — nested/embedded JSON config (planted-bug: unexported field silently dropped) — exported + tags
- m `customtime` — RFC3339 time field via tag + wrapper — tag + type (+strings)
- m `rawdefer` — `json.RawMessage` defer-decode by type field — tags + two-pass (+comma-ok)

### slot 12 — 05-structs/02-embedding-structs
- j `basemodel` — embed `BaseModel{ID,CreatedAt}` in entities — struct embedding, promotion
- j `logfields` — embed a `Fields` struct to inherit helpers — promoted methods intro
- m `overridemethod` — outer shadows embedded method (planted-bug: calls wrong one) — method promotion + shadow
- m `mixinvalidate` — compose behavior from two embedded structs — multiple embedding
- m `embedconflict` — ambiguous promoted field selector (planted-bug: compile/selector) — conflict resolution
- m `jsonembed` — embedded struct flattening in JSON tags — embedding + tags (+json)
- s `embedptr` — pointer-embed sharing state, unexpected aliasing (planted-bug) — embed pointer vs value (+slice header)

### slot 13 — 05-structs/03-memory-layout
- m `fieldorder` — reorder fields to shrink struct size (bench sizeof) — alignment/padding
- s `padcut` — cut a hot 10M-elem `[]struct` footprint via ordering (planted-bug, `-benchmem`) — padding at scale
- s `alignatomic` — 64-bit atomic field alignment on struct (planted-bug) — alignment guarantees (+atomics intro)
- s `cachefriendly` — SoA vs AoS layout for scan speed under time ceiling (planted-bug) — cache lines (+slices)
- st `falseshare` — two counters on one cache line, false sharing (planted-bug, bench) — false sharing + padding (+goroutines)
- st `emptytail` — zero-size trailing field address gotcha (planted-bug) — struct addressing rules (+empty-struct)
- st `unsafeoffset` — `unsafe.Offsetof` layout assumption breaks (planted-bug) — memory layout + unsafe

### slot 14 — 06-empty-struct
- j `stringset` — `map[string]struct{}` set add/has/delete — empty struct as set value
- j `donesignal` — `chan struct{}` completion signal — zero-size signaling (intro)
- m `visitedset` — graph visited-set, no bool waste (planted-bug: `map[T]bool` leak intent) — struct{} 0 bytes (+maps)
- m `enumset` — bitset-like membership via struct{} map — set ops (+comma-ok)
- m `keyonlyindex` — presence index over 10M keys, memory win vs bool — footprint (+map internals)
- s `signalfanout` — close(chan struct{}) fan-out to N waiters (planted-bug) — broadcast semantics (+goroutines)

### slot 15 — 07-anonymous-structs
- j `tablecase` — table-driven test rows as `[]struct{...}` — anonymous struct literal
- j `configlit` — one-off nested config literal — anonymous nested struct
- j `pairreturn` — group locals in an anonymous struct — ad-hoc grouping
- m `jsonshape` — decode into anonymous struct for a one-shot endpoint — anon + json tags (+json)
- m `groupby` — `map[key]struct{sum,count int}` aggregation — anon struct map value (+maps)
- m `sortkey` — decorate-sort with anonymous `{item,key}` slice — anon + sort (+slices)

## 5. Per-puzzle authoring loop (each of the 100)

1. `scripts/coverage.sh 01-language-basics/03-composite-types/<slot>` → covered set.
2. Scaffold from `_template/` into `<slot>/<level>/<name>/`; inject `COVERED:`.
3. Stage red state by mode (§5e): junior/middle stub, senior/staff planted-bug.
4. Write table tests; add §5b constraint tests (bench/RAM/`-race`) for senior/staff.
5. Add §5f guard test where the task is hardcode-cheatable.
6. README: Context / Task / ≥3 Examples / Topics to Master (target + woven) / hint at correct depth (§5d).
7. `make verify` red→green; delete `.keep`.

## 6. Batch plan (authoring order)

Author by slot, junior→staff within slot, in the table order (1→15). Suggested
batches of ~10 for review checkpoints:

B1 arrays(7)+make(5)…, B2 capacity+conversions, B3 header-internals+slice-tricks,
B4 strings, B5 maps(both), B6 struct-tags+embedding, B7 memory-layout,
B8 empty-struct+anonymous-structs. `make verify` green after each batch.

## 7. Open items / risks

- **Scope creep:** several middle/senior scenarios lean on `strings`/`maps`
  which sit *inside* this same topic — order slots so a woven concept is
  authored (or at least covered) first. Confirm `coverage.sh` returns them.
- **Layout mismatch:** on-disk grid is topic-first, but §1 says level-first.
  We follow the existing on-disk `<slot>/<level>/<name>` convention (as
  `dedupe` did). Flag if a canonical relayout is wanted instead.
- Names/scenarios are provisional; each may shift during authoring to keep the
  bug natural (§5c "never contrive it").
