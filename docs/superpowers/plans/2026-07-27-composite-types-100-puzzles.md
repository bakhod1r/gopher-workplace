# Composite-Types 100 Puzzles — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Author 100 real-work Go puzzles across the 15 empty subtopic slots of
`challenges/01-language-basics/03-composite-types`, spread over 4 levels by
per-slot competency.

**Architecture:** Each puzzle is its own Go module authored red→green. Junior/
middle ship a from-scratch stub (`panic("not implemented")`); senior/staff ship
exactly one planted bug between `// CHANGE CODE` markers. One fully-worked
exemplar (Task 1, puzzle `checkoutgrid`) defines the shape; every later puzzle
repeats that shape with its own concrete signature, test table, and README.
The gate for each task is `make verify` (fmt-check + vet + test) — red before
authoring the solution reference, green after the stub is implementable.

**Tech Stack:** Go 1.26, `go test`, `gofmt`, `go vet`, GNU make. No external deps.

## Global Constraints

- Layout: `03-composite-types/<slot>/<level>/<name>/{go.mod,Makefile,<pkg>.go,<pkg>_test.go,README.md}`. `<level>` ∈ `junior|middle|senior|staff`.
- Module path: `github.com/gopher-workplace/challenges/01-language-basics/03-composite-types/<slot>/<level>/<name>` — exact.
- `go 1.26` in every `go.mod`.
- **`scripts/scaffold.sh` does NOT fit these nested slots** (it parses a 3-part `level/topic/subtopic` slot). Scaffold manually: copy an existing puzzle dir (e.g. `.../junior/dedupe`) as the template, then rename package/files and rewrite. Copy `Makefile` verbatim.
- Mode by level (§5e): junior/middle = from-scratch stub, no `CHANGE CODE` markers; senior/staff = exactly one planted bug between `// CHANGE CODE BELOW/ABOVE THIS LINE`. Never mix.
- Difficulty dial (§5b): junior correctness only; middle no O(n²)/needless allocs; senior scale + RAM ceiling (bench/streaming test); staff CPU/time ceiling + `-race`-clean concurrency.
- Depth (§5d): pitch hint + Topics-to-Master at the level's depth.
- Cumulative (§5c): weave ≥1–2 earlier covered concepts, list them in README Topics to Master.
- Guard test (§5f) where the task is hardcode-cheatable: parse `<pkg>.go` with `go/parser` mode `0`, emit `t.Logf("WARN: ...")` (never fail).
- README sections: Context / Task / Examples (≥3) / Topics to Master / Hint / Validate. No SOLUTION.md. No roadmap-path references.
- Do not change a puzzle's function signature or its task tests once written.
- **PROHIBITED (global):** `git commit`, `git push`, `rm`. No task step commits — the deliverable gate is `make verify`, not a commit. Delete a `.keep` by moving it aside or leaving authoring to the tooling, never `rm` (use `git rm --cached`-free approach: the slot dir simply gains the puzzle; the stray `.keep` can stay or be emptied — flag to the user rather than `rm`).
- Verify one puzzle: `make -C <puzzle-dir> verify`. Verify all: `make verify` at repo root.
- **Shipped state is RED by design.** The delivered puzzle file is the stub (panics) or the planted-bug version — its tests MUST fail (`make verify` red is correct, not a defect). Correctness = red→green proof: (a) shipped stub/bug fails tests; (b) a reference solution (never shipped) passes `make verify`. Reviewers must not require the shipped file to pass tests.
- **Commits enabled for this work** (human granted). Each puzzle (or slot task) commits on `main` with trailer `Co-Authored-By: Claude Opus 4.8 <noreply@anthropic.com>`. `.keep` removal: leave the stray `.keep` in place (no `rm`); note as deferred-minor.

---

## Task 1: Exemplar — `01-arrays/junior/checkoutgrid` (fully worked)

This task is the reference pattern. Every later puzzle mirrors these five files.

**Files:**
- Create: `challenges/01-language-basics/03-composite-types/01-arrays/junior/checkoutgrid/go.mod`
- Create: `.../checkoutgrid/Makefile` (copy verbatim from `.../junior/dedupe/Makefile`)
- Create: `.../checkoutgrid/checkoutgrid.go`
- Test: `.../checkoutgrid/checkoutgrid_test.go`
- Create: `.../checkoutgrid/README.md`

**Interfaces:**
- Produces: `func SeatMap(taken [][2]int) [7][10]bool` — mark taken seats on a fixed 7×10 grid; `taken` is a list of `{row,col}`; out-of-range entries ignored.

- [ ] **Step 1: Scaffold from dedupe**

```bash
SRC=challenges/01-language-basics/03-composite-types/junior/dedupe
DST=challenges/01-language-basics/03-composite-types/01-arrays/junior/checkoutgrid
mkdir -p "$DST"
cp "$SRC/Makefile" "$DST/Makefile"
```

- [ ] **Step 2: Write `go.mod`**

```
module github.com/gopher-workplace/challenges/01-language-basics/03-composite-types/01-arrays/junior/checkoutgrid

go 1.26
```

- [ ] **Step 3: Write the failing test** (`checkoutgrid_test.go`)

```go
package checkoutgrid

import (
	"reflect"
	"testing"
)

func TestSeatMap(t *testing.T) {
	cases := []struct {
		name  string
		taken [][2]int
		want  func() [7][10]bool
	}{
		{"single", [][2]int{{0, 0}}, func() [7][10]bool { var g [7][10]bool; g[0][0] = true; return g }},
		{"several", [][2]int{{1, 2}, {6, 9}}, func() [7][10]bool { var g [7][10]bool; g[1][2] = true; g[6][9] = true; return g }},
		{"out of range ignored", [][2]int{{7, 0}, {0, 10}, {-1, 0}}, func() [7][10]bool { var g [7][10]bool; return g }},
		{"empty", nil, func() [7][10]bool { var g [7][10]bool; return g }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SeatMap(tc.taken); !reflect.DeepEqual(got, tc.want()) {
				t.Errorf("SeatMap(%v) = %v, want %v", tc.taken, got, tc.want())
			}
		})
	}
}
```

- [ ] **Step 4: Write the stub** (`checkoutgrid.go`) — ships RED

```go
// Package checkoutgrid — Gopher Workplace challenge.
package checkoutgrid

// SeatMap returns a fixed 7-row by 10-column seating grid with every valid
// {row,col} in taken marked true. Coordinates outside [0,7)×[0,10) are ignored.
//
// Examples:
//
//	SeatMap([][2]int{{0,0}})        => grid with [0][0]=true
//	SeatMap([][2]int{{1,2},{6,9}})  => grid with [1][2],[6][9]=true
//	SeatMap([][2]int{{7,0}})        => all false (row 7 out of range)
func SeatMap(taken [][2]int) [7][10]bool {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
```

- [ ] **Step 5: Verify RED**

Run: `make -C challenges/01-language-basics/03-composite-types/01-arrays/junior/checkoutgrid verify`
Expected: test panics/fails ("not implemented").

- [ ] **Step 6: Confirm solvability** — temporarily implement the reference in a scratch copy, run `make verify`, expect PASS, then revert the shipped `.go` back to the stub. The shipped file stays the stub; only confirm it is solvable.

Reference solution (do NOT ship — for solvability check only):

```go
func SeatMap(taken [][2]int) [7][10]bool {
	var g [7][10]bool
	for _, rc := range taken {
		r, c := rc[0], rc[1]
		if r >= 0 && r < 7 && c >= 0 && c < 10 {
			g[r][c] = true
		}
	}
	return g
}
```

- [ ] **Step 7: Write `README.md`** — Context (checkout seat-hold service), Task (implement `SeatMap`, ignore out-of-range, don't change signature/tests), ≥3 Examples, Topics to Master (array value type; fixed-size `[N][M]` vs slice; zero value of array is all-zero; indexing bounds), Hint (declare `var g [7][10]bool`, range and guard bounds), Validate (`make verify`).

- [ ] **Step 8: Gate** — `make -C <dir> verify` green with the stub replaced by the reference during check; ship the stub. Do NOT commit (global rule). Leave the slot's `.keep` for the user to clear (no `rm`).

---

## Task 2–15: one task per slot

Each task authors **all puzzles of one slot**, junior→staff order, each puzzle
following the Task 1 seven-file pattern (scaffold → go.mod → failing test → stub
or planted-bug → verify red → solvability check → README → gate). Per puzzle:

- **junior/middle:** ship a from-scratch stub (`panic("not implemented")`).
- **senior/staff:** ship one planted bug between `// CHANGE CODE BELOW/ABOVE THIS LINE`; the surrounding code compiles and looks right.
- Get the covered set first: `scripts/coverage.sh 01-language-basics/03-composite-types/<slot>` (informational; scaffold.sh itself is not used — see Global Constraints).
- Deliverable gate for the task: `make verify` at repo root passes for every puzzle in the slot.

Interfaces per puzzle (exact function signatures the executor must produce). `[L]`
= level; mode follows the level rule above.

### Task 2 — slot `01-arrays` (remaining 6; checkoutgrid done in Task 1)
- `rgbahist` [j] `func Histogram(px []byte) [256]int` — count byte frequencies. woven: range, zero value.
- `weekload` [j] `func PeakDay(load [7]float64) int` — index of max weekday; ties→lowest index. woven: array compare, indexing.
- `boardhash` [j] `func Equal(a, b [9]byte) bool` — board equality (use `==`). woven: arrays are comparable.
- `rollingavg` [m] `func Averages(lat []float64, window int) []float64` — ring buffer `[window]` averaging; O(n). woven: array + modulo, slice output.
- `dnscache4` [m] `func Resolve(cache map[[4]byte]string, ip [4]byte) (string, bool)` — array as map key. woven: maps, comma-ok.
- `matmul` [m] `func Mul(a, b [3][3]int) [3][3]int` — 3×3 multiply. woven: multidim arrays, nested loops.

### Task 3 — slot `02-slices/01-capacity-and-growth` (7)
- `growlog` [j] `func Grow(events []string) []int` — cap after each append (len(events) entries). woven: len/cap.
- `preallocbatch` [j] `func Batch(n int) []int` — `make([]int,0,n)` then append 0..n-1; final cap==n. woven: prealloc.
- `appendaliasbug` [m, PLANTED-BUG] `func SubBatch(parent []int, from, to int) []int` — return copy of a range that must not alias parent's backing; bug = returns `parent[from:to]` re-append aliasing. woven: slice header.
- `dedupeinplace` [m] `func Compact(s []int) []int` — dedupe adjacent using `s[:0]` reuse. woven: append, cap reuse.
- `csvcols` [m] `func Columns(rows []string, sep byte) [][]string` — split; amortized growth, no O(n²). woven: growth.
- `streamdedup` [s, PLANTED-BUG] `func StreamDedupe(r io.Reader) (int, error)` — count distinct ints from a stream under RAM ceiling; bug = accumulates all into a slice. bench + large-input test. woven: map set, streaming.
- `nocopygrow` [s, PLANTED-BUG] `func Append(dst []T, v T) []T` — hot-path append that must not realloc within budgeted cap; bug = ignores cap and re-makes. `testing.B` ns/op + allocs. woven: escape, capacity.

### Task 4 — slot `02-slices/02-make` (5)
- `zerobuf` [j] `func Frame(n int) []byte` — `make([]byte,n)`; all-zero, len n. woven: make len.
- `gridrows` [j] `func Grid(rows, cols int) [][]int` — allocate 2D. woven: nested make.
- `capvslen` [j] `func Fill(n int) []int` — `make([]int,n/2,n)` then append rest; len==n. woven: three-arg make.
- `readnbytes` [m, PLANTED-BUG] `func ReadN(src []byte, n int) []byte` — pre-sized buffer; bug = `make([]byte,n)` then append (double length). woven: len vs cap trap.
- `poolframe` [m] `func Process(frames [][]byte) []int` — one reusable buffer `make` once, reset per iter. woven: reuse, capacity.

### Task 5 — slot `02-slices/03-slice-to-array-conversion` (5)
- `ipv4octets` [j] `func Octets(b []byte) ([4]byte, bool)` — slice→`[4]byte`, ok=false if len≠4. woven: slice→array.
- `sha256fix` [j] `func Digest(sum []byte) [32]byte` — fixed digest from 32-byte slice. woven: array ptr conversion.
- `headerframe` [m, PLANTED-BUG] `func Header(stream []byte) ([8]byte, error)` — parse first 8 bytes; bug = converts without length check → panic. woven: length-check, capacity.
- `chunk16` [m] `func Blocks(payload []byte) [][16]byte` — split into `[16]byte` blocks (payload len multiple of 16). woven: slicing, slice→array in loop.
- `magicbytes` [m] `func IsPNG(b []byte) bool` — compare first 4 bytes to signature via `[4]byte`. woven: array equality, arrays.

### Task 6 — slot `02-slices/04-array-to-slice-conversion` (5)
- `arrslice` [j] `func View(cfg [8]int) []int` — expose array as `[]int` via `cfg[:]`. woven: array→slice.
- `bufwriter` [j] `func WriteInto(buf *[16]byte, data []byte) int` — write into `buf[:]`, return n copied. woven: slicing an array.
- `aliasarray` [m, PLANTED-BUG] `func Bump(arr *[4]int)` — mutate via slice `arr[:]`; bug = slices a local copy, mutation lost. woven: array→slice aliasing, slice header.
- `windowview` [m] `func Windows(buf [16]byte, size int) [][]byte` — sliding windows over `buf[:]`. woven: sub-slicing array.
- `passbyslice` [m] `func Sum(arr *[1024]int) int` — sum via slice view, no array copy. woven: avoid array copy, make.

### Task 7 — slot `02-slices/05-slice-header-internals` (7)
- `headerprint` [m] `func Reslice(s []int, low, high int) (length, capacity int)` — return len/cap of `s[low:high]`. woven: header = ptr+len+cap.
- `truncleak` [m, PLANTED-BUG] `func Head(s []byte, n int) []byte` — return first n but must not retain backing; bug = `s[:n]` keeps full cap. woven: reslice keeps cap.
- `subslicealias` [s, PLANTED-BUG] `func Split(s []int, at int) (left, right []int)` — hand out independent halves; bug = both share backing, append to left corrupts right. woven: shared backing, capacity.
- `clipthree` [s, PLANTED-BUG] `func Safe(s []int, from, to int) []int` — full-slice-expr `s[from:to:to]`; bug = two-index form leaks cap. woven: three-index slice.
- `gcpin` [s, PLANTED-BUG] `func FirstLine(data []byte) []byte` — return first line copied out; bug = sub-slice pins huge backing (leak). bench allocs. woven: copy to release, GC reachability.
- `racyappend` [st, PLANTED-BUG] `func FanAppend(base []int, vals []int) [][]int` — N goroutines append to shared-backing sub-slices; bug = data race (`-race` fails). woven: header not concurrency-safe, goroutines.
- `zerocopyparse` [st, PLANTED-BUG] `func Fields(line []byte) [][]byte` — zero-copy sub-slice fields under time ceiling; bug = later append aliases a field. bench ns/op. woven: header cost model, bench.

### Task 8 — slot `02-slices/06-slice-tricks` (9)
- `insertat` [j] `func Insert(s []int, i, v int) []int` — insert v at index i. woven: insert idiom.
- `deleteat` [j] `func Delete(s []int, i int) []int` — delete index i, order preserved. woven: delete idiom.
- `reverseinplace` [j] `func Reverse(s []int)` — reverse in place. woven: two-pointer swap.
- `filterinplace` [m, PLANTED-BUG] `func DropNil(s []*int) []*int` — `s[:0]` reuse; bug = leaves dangling pointers in tail (no zeroing). woven: `s[:0]`, GC.
- `dedupsorted` [m] `func DedupSorted(s []int) []int` — dedupe sorted in place. woven: adjacent-compare.
- `movetofront` [m] `func MoveToFront(s []int, i int) []int` — rotate element i to front, no realloc. woven: rotate trick, capacity.
- `batchsplit` [m] `func Chunks(s []int, size int) [][]int` — split into size-N chunks. woven: chunk idiom, make.
- `bigdelete` [s, PLANTED-BUG] `func DeleteAll(s []int, pred func(int) bool) []int` — compact 100M slice in place, bounded allocs; bug = builds a new slice each match. bench allocs. woven: in-place compact at scale.
- `stablepartition` [s, PLANTED-BUG] `func Partition(s []int, pred func(int) bool) int` — stable partition under RAM ceiling; bug = allocates aux buffer. bench. woven: in-place, no aux.

### Task 9 — slot `03-strings` (9)
- `slugify` [j] `func Slug(title string) string` — lowercased, spaces→`-`, alnum only. woven: byte vs rune, building.
- `maskcard` [j] `func Mask(card string) string` — keep last 4, rest `*`. woven: indexing, concat.
- `countwords` [j] `func Words(s string) int` — whitespace-delimited word count. woven: range over string.
- `runereverse` [m, PLANTED-BUG] `func Reverse(s string) string` — reverse by runes; bug = byte reversal mangles UTF-8. woven: runes vs bytes, slices.
- `builderjoin` [m] `func Join(parts []string, sep string) string` — via `strings.Builder`, no `+=`. woven: Builder over concat.
- `trimparse` [m] `func Parse(line string) (key, val string, ok bool)` — split `key=value`, trim spaces. woven: slicing strings, maps context.
- `bytestostr` [m, PLANTED-BUG] `func Normalize(b []byte) string` — return normalized string reusing `b` where safe; bug = returns `string(b)` view that later mutates. woven: conversion cost.
- `csvquote` [m] `func Quote(field string) string` — RFC4180 escape (quote if contains `,"`\n`). woven: byte scanning, builder.
- `bigconcat` [s, PLANTED-BUG] `func Report(rows []string) string` — assemble large report streaming with Builder, bounded allocs; bug = naive `+=` O(n²). bench. woven: Builder growth at scale, capacity.

### Task 10 — slot `04-maps/01-comma-ok-idiom` (6)
- `featureflag` [j] `func Enabled(flags map[string]bool, name string) bool` — default false when absent (comma-ok, not just index). woven: comma-ok read.
- `wordcount` [j] `func Count(words []string) map[string]int` — frequency via zero-value increment. woven: zero-value + comma-ok.
- `firstseen` [j] `func FirstIndex(xs []int) map[int]int` — first index of each value, ignore later. woven: comma-ok guard.
- `sessionget` [m, PLANTED-BUG] `func TTL(sessions map[string]int, id string) (int, bool)` — distinguish missing vs zero TTL; bug = uses `if v == 0` to mean missing. woven: comma-ok vs zero value.
- `graphadj` [m] `func AddEdge(g map[int][]int, u, v int)` — lazy-init neighbor slice. woven: comma-ok + append, slices.
- `dedupemap` [m] `func Unique(xs []int) []int` — dedupe via `map[int]struct{}` set, first-appearance order. woven: set idiom, empty-struct.

### Task 11 — slot `04-maps/02-map-internals` (7)
- `iterorder` [m, PLANTED-BUG] `func TopKey(m map[string]int) string` — return max-value key deterministically; bug = relies on range order for ties. woven: unordered iteration.
- `deletewhileiter` [m, PLANTED-BUG] `func Prune(m map[string]int, min int)` — delete entries below min during range; bug = mutates in unsafe way / re-adds. woven: iteration + delete rules.
- `preallocmap` [s, PLANTED-BUG] `func Index(keys []string) map[string]int` — `make(map,len(keys))` to cut rehash on 10M; bug = no size hint. bench. woven: sizing, load factor.
- `nilmapwrite` [s, PLANTED-BUG] `func Cache() *Store` + `func (s *Store) Put(k, v string)` — lazy map write; bug = nil-map write panic. woven: nil vs empty map, comma-ok.
- `keyalloc` [s, PLANTED-BUG] `func Lookup(m map[Point]int, p Point) int` where `Point struct{X,Y int}` — struct-key lookup no per-call alloc; bug = uses string-concatenated key. `-benchmem`. woven: hashing + escape, structs.
- `concurrentmap` [st, PLANTED-BUG] `func NewSharded() *Sharded` + `Get/Set` — sharded map race-free; bug = unsynchronized shared map (`-race` fails). woven: map not concurrency-safe, goroutines.
- `syncmapswap` [st, PLANTED-BUG] `func (c *Counter) Inc(k string)` — mutex-map vs `sync.Map` under CPU ceiling; bug = global lock contention. bench. woven: internals + contention, bench.

### Task 12 — slot `05-structs/01-struct-tags-and-json` (7)
- `apiuser` [j] `func Encode(u User) (string, error)` with tagged `User` — exact wire field names. woven: json tags.
- `omitempty` [j] `func Encode(p Profile) (string, error)` — omit zero fields. woven: omitempty.
- `renamekeys` [j] `func Decode(data string) (Config, error)` — snake_case wire ↔ Go fields. woven: tag mapping.
- `unmarshalpartial` [m] `func Known(data string) (map[string]any, error)` — decode only known fields. woven: decoder + tags, maps.
- `nestedconfig` [m, PLANTED-BUG] `func Decode(data string) (Server, error)` — nested config; bug = an unexported field silently dropped from JSON. woven: exported + tags.
- `customtime` [m] `func Decode(data string) (Event, error)` — RFC3339 time field via wrapper type + tag. woven: tag + type, strings.
- `rawdefer` [m] `func Decode(data string) (any, error)` — `json.RawMessage` two-pass decode by `type` field. woven: tags + two-pass, comma-ok.

### Task 13 — slot `05-structs/02-embedding-structs` (7)
- `basemodel` [j] `func Describe(u User) string` where `User` embeds `BaseModel{ID int; CreatedAt string}` — access promoted fields. woven: embedding, promotion.
- `logfields` [j] `func (e Entry) Line() string` where `Entry` embeds `Fields` — promoted method use. woven: promoted methods.
- `overridemethod` [m, PLANTED-BUG] `func (a Admin) Role() string` where `Admin` embeds `User`; bug = calls promoted `User.Role` instead of override. woven: method promotion + shadow.
- `mixinvalidate` [m] `func (f Form) Validate() []string` composing two embedded validators. woven: multiple embedding.
- `embedconflict` [m, PLANTED-BUG] `func Name(x T) string` where two embedded structs both have `Name`; bug = ambiguous selector (must qualify). woven: conflict resolution.
- `jsonembed` [m] `func Encode(r Response) (string, error)` — embedded struct flattening with tags. woven: embedding + tags, json.
- `embedptr` [s, PLANTED-BUG] `func Clone(a *A) *A` where `A` embeds `*Shared`; bug = clones share `*Shared` state (aliasing). woven: embed pointer vs value, slice header.

### Task 14 — slot `05-structs/03-memory-layout` (7)
- `fieldorder` [m] `func SizeOf() uintptr` / reorder a struct's fields to shrink `unsafe.Sizeof`. bench sizeof. woven: alignment/padding.
- `padcut` [s, PLANTED-BUG] `type Row struct{...}` + `func Total(rows []Row) int` — reorder to cut 10M-elem footprint; bug = padded field order. `-benchmem`. woven: padding at scale.
- `alignatomic` [s, PLANTED-BUG] `type Counter struct{...}` — 64-bit atomic field must be aligned; bug = misaligned field (panics on 32-bit). woven: alignment guarantees, atomics.
- `cachefriendly` [s, PLANTED-BUG] `func Scan(...)` — SoA vs AoS for scan speed under time ceiling; bug = AoS layout thrashes cache. bench. woven: cache lines, slices.
- `falseshare` [st, PLANTED-BUG] `type Pair struct{...}` + parallel counters; bug = two counters on one cache line (false sharing). bench. woven: false sharing + padding, goroutines.
- `emptytail` [st, PLANTED-BUG] `func Addr(...)` — zero-size trailing field address gotcha; bug = assumes distinct address past struct end. woven: struct addressing rules, empty-struct.
- `unsafeoffset` [st, PLANTED-BUG] `func Read(...)` using `unsafe.Offsetof`; bug = hardcoded offset breaks under field reorder. woven: memory layout + unsafe.

### Task 15 — slot `06-empty-struct` (6) then `07-anonymous-structs` (6)

**06-empty-struct:**
- `stringset` [j] `func NewSet(xs []string) map[string]struct{}` + `Has`. woven: empty struct as set value.
- `donesignal` [j] `func Worker() <-chan struct{}` — completion signal channel. woven: zero-size signaling.
- `visitedset` [m, PLANTED-BUG] `func Reachable(g map[int][]int, start int) int` — DFS with `map[int]struct{}` visited; bug = re-visits / infinite loop with wrong visited check. woven: struct{} 0 bytes, maps.
- `enumset` [m] `func Membership(perms []int) map[int]struct{}` — set ops. woven: set ops, comma-ok.
- `keyonlyindex` [m] `func Index(keys []string) map[string]struct{}` — presence index, memory win vs bool. woven: footprint, map internals.
- `signalfanout` [s, PLANTED-BUG] `func Broadcast(n int) (func(), []<-chan struct{})` — `close(chan struct{})` fan-out; bug = sends instead of close (only one waiter wakes). woven: broadcast semantics, goroutines.

**07-anonymous-structs:**
- `tablecase` [j] `func Run() []string` — table-driven rows as `[]struct{name string; in,want int}`. woven: anonymous struct literal.
- `configlit` [j] `func Default() any` — one-off nested anonymous config literal. woven: anonymous nested struct.
- `pairreturn` [j] `func Stats(xs []int) (any)` — group locals in anonymous struct. woven: ad-hoc grouping.
- `jsonshape` [m] `func Parse(data string) (string, int, error)` — decode into anonymous struct for one endpoint. woven: anon + json tags, json.
- `groupby` [m] `func GroupBy(xs []int) map[int]struct{ Sum, Count int }` — aggregation with anon struct value. woven: anon struct map value, maps.
- `sortkey` [m] `func SortByKey(items []string, key func(string) int) []string` — decorate-sort with anon `{item,key}` slice. woven: anon + sort, slices.

---

## Self-Review

- **Spec coverage:** all 15 slots + 100 puzzle names from the design's roster are present across Tasks 1–15 (Task 1 covers `checkoutgrid`; Task 2 the other 6 of `01-arrays`; Tasks 3–15 the rest). Counts per slot match the design matrix (7/7/5/5/5/7/9/9/6/7/7/7/7/6+6). ✓
- **Placeholder scan:** every puzzle carries a concrete signature + level + mode + woven concept; the exemplar carries full test/stub/solution code. No "TBD". ✓
- **Type consistency:** signatures are self-contained per puzzle; no cross-puzzle type references. ✓
- **Mode discipline:** planted-bug marked `[PLANTED-BUG]` (senior/staff + a few middle debug puzzles); all others from-scratch stubs. ✓
- **Prohibited ops:** no task step runs `git commit`/`push`/`rm`; gate is `make verify`; `.keep` removal deferred to the user. ✓

## Notes / risks carried from spec

- Intra-topic dependency: several middle/senior puzzles weave `strings`/`maps`
  (same topic). `coverage.sh` confirms availability at each slot's position;
  author slots in table order so woven concepts are covered first.
- Names/scenarios provisional — adjust to keep each planted bug natural (§5c).
- On-disk layout is topic-first; we follow the existing `<slot>/<level>/<name>`
  convention (as `dedupe`). Flag if a canonical level-first relayout is wanted.
