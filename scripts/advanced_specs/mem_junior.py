"""10-advanced-topics / 01-memory-management-in-depth / junior — 20 puzzles."""

SUB = "01-memory-management-in-depth"
LEVEL = "junior"

SPECS = []


def P(**kw):
    kw.setdefault("sub", SUB)
    kw.setdefault("level", LEVEL)
    kw.setdefault("mode", "stub")
    kw.setdefault("stub", 'panic("not implemented")')
    kw.setdefault("imports", [])
    kw.setdefault("extra", "")
    SPECS.append(kw)


P(
    name="prealloc",
    title="Allocate The Slice Once",
    sig="func Squares(n int) []int",
    doc="""Squares returns the squares 0..n-1 in order.

The result must be built with a single allocation: give the slice its
final length up front instead of growing it element by element.

Examples:

	Squares(3) => []int{0, 1, 4}""",
    solution="""out := make([]int, n)
for i := range out {
	out[i] = i * i
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestSquares(t *testing.T) {
	if got := Squares(4); !reflect.DeepEqual(got, []int{0, 1, 4, 9}) {
		t.Errorf("Squares(4) = %v, want [0 1 4 9]", got)
	}
	if got := Squares(0); len(got) != 0 {
		t.Errorf("Squares(0) = %v, want empty", got)
	}
}

func TestSquaresAllocatesOnce(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { _ = Squares(64) }); n > 1 {
		t.Errorf("Squares made %v allocations, want 1: size the slice up front", n)
	}
}
""",
    context="A report builder grows its result slice one `append` at a time. Under profiling most of its time is spent copying the backing array as it doubles.",
    task=[
        "Return the squares of `0..n-1` in order.",
        "The whole result must come from one allocation — `AllocsPerRun` must see at most 1.",
    ],
    examples=[
        ("Squares(4)", "[0 1 4 9]", None),
        ("Squares(1)", "[0]", None),
        ("Squares(0)", "[]", "n == 0 gives an empty, non-nil-or-nil slice; length is what is checked."),
    ],
    topics=[
        ("make with a length", "`make([]T, n)` reserves the final size in one allocation."),
        ("append growth", "Appending past cap allocates a bigger array and copies — repeatedly."),
        ("Allocation counting", "`testing.AllocsPerRun` grades allocation behaviour, not just output."),
    ],
    hint="You already know the final length before the loop starts.",
    intuition="`append` cannot see the future: each time it runs out of capacity it allocates a larger array and copies everything over. When the final size is known, ask for it once.",
    approach=[
        "Allocate `make([]int, n)`.",
        "Fill each index with `i * i`.",
        "Return the slice.",
    ],
    walkthrough="For n = 64, appending from a nil slice reallocates about seven times (1, 2, 4, 8, 16, 32, 64) and copies 63 elements in total. `make([]int, 64)` allocates once and copies nothing.",
    pitfalls=[
        "`make([]int, n)` versus `make([]int, 0, n)` — the first is already filled with zeros and is indexed, the second must be appended to.",
        "Using `append` on a slice made with a length appends *after* the zeros.",
    ],
)

P(
    name="cloneints",
    title="A Copy That Owns Its Memory",
    sig="func Clone(s []int) []int",
    doc="""Clone returns a copy of s that shares no storage with s.

Writes to the result must not be visible through s, and writes to s must
not be visible through the result.

Examples:

	Clone([]int{1, 2}) => []int{1, 2}""",
    solution="""out := make([]int, len(s))
copy(out, s)
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestCloneContents(t *testing.T) {
	if got := Clone([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Clone = %v, want [1 2 3]", got)
	}
	if got := Clone(nil); len(got) != 0 {
		t.Errorf("Clone(nil) = %v, want empty", got)
	}
}

func TestCloneIsIndependent(t *testing.T) {
	s := []int{1, 2, 3}
	c := Clone(s)
	s[0] = 99
	if c[0] != 1 {
		t.Errorf("c[0] = %d, want 1: the clone still views s", c[0])
	}
	c[1] = 42
	if s[1] != 2 {
		t.Errorf("s[1] = %d, want 2: s still views the clone", s[1])
	}
}
""",
    context="A cache hands callers the slice it stores internally. Callers mutate what they get back, and the cached entry changes underneath the cache.",
    task=[
        "Return a copy of `s` that shares no backing array with it.",
        "Handle a nil or empty input without panicking.",
    ],
    examples=[
        ("Clone([]int{1,2,3})", "[1 2 3]", None),
        ("c := Clone(s); s[0] = 99", "c[0] is unchanged", "The copy owns its own array."),
        ("Clone(nil)", "[]", None),
    ],
    topics=[
        ("Slice header vs storage", "A slice is a pointer, length and capacity — assigning one copies the header, not the elements."),
        ("copy", "`copy(dst, src)` moves elements up to the shorter length."),
        ("Aliasing", "Two slices over one array see each other's writes."),
    ],
    hint="`d := s` copies three words. What copies the elements?",
    intuition="A slice value is a small header pointing at an array. Handing it out hands out the array. To break the link you must allocate a new array and copy the elements into it.",
    approach=[
        "Allocate a destination of `len(s)`.",
        "`copy` the elements across.",
        "Return the destination.",
    ],
    walkthrough="`s := []int{1,2,3}` allocates one array. `c := s` yields a second header over the *same* array, so `s[0] = 99` shows up in `c`. After `make` + `copy` there are two arrays and the write is invisible.",
    pitfalls=[
        "`s[:]` is not a copy.",
        "`copy(out, s)` into an out of length 0 copies nothing — the destination's length is what limits it.",
    ],
)

P(
    name="clearmap",
    title="Empty The Map, Keep The Map",
    sig="func Reset(m map[string]int)",
    doc="""Reset removes every entry from m without replacing the map.

Callers hold the same map value, so the entries must be deleted in place
rather than by assigning a fresh map.

Examples:

	m := map[string]int{"a": 1}; Reset(m) => len(m) == 0""",
    solution="""clear(m)""",
    tests="""
import "testing"

func TestResetEmpties(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	Reset(m)
	if len(m) != 0 {
		t.Errorf("len(m) = %d, want 0", len(m))
	}
}

func TestResetIsVisibleToTheCaller(t *testing.T) {
	m := map[string]int{"a": 1}
	alias := m
	Reset(m)
	if len(alias) != 0 {
		t.Errorf("len(alias) = %d, want 0: the caller's map was not emptied", len(alias))
	}
	m["c"] = 3
	if alias["c"] != 3 {
		t.Error("the map was replaced, not emptied")
	}
}

func TestResetNilMap(t *testing.T) {
	Reset(nil)
}
""",
    context="A worker reuses one map between batches. Assigning a fresh map inside the helper leaves every other holder of the old map looking at stale data.",
    task=[
        "Delete every entry of `m` in place.",
        "The caller's map value must stay the same map — do not assign a new one.",
        "A nil map must not panic.",
    ],
    examples=[
        ('Reset(map[string]int{"a":1})', "len == 0", None),
        ("alias := m; Reset(m)", "len(alias) == 0", "Both names still refer to one map."),
        ("Reset(nil)", "no panic", None),
    ],
    topics=[
        ("clear", "`clear(m)` removes all entries from the map you already have."),
        ("Maps are reference-like", "A map value is a pointer to a runtime structure; the parameter is a copy of that pointer."),
        ("Reuse over reallocation", "Emptying keeps the buckets already paid for."),
    ],
    hint="Assigning `m = map[string]int{}` only rebinds the local parameter.",
    intuition="The map parameter is a copy of the caller's map *handle*. Writing through the handle (deleting keys) is visible to the caller; overwriting the handle is not.",
    approach=[
        "Call the builtin `clear` on the map.",
        "`clear` on a nil map is a no-op, so nothing extra is needed.",
    ],
    walkthrough="`m := map[string]int{\"a\":1}; alias := m` — two handles, one table. `clear(m)` empties the table, so `len(alias) == 0`. `m = map[string]int{}` would give `m` a new table and leave `alias` holding the old one.",
    pitfalls=[
        "Assigning a new map inside the function — the caller never sees it.",
        "Ranging and deleting works, but `clear` says it in one line and handles nil.",
    ],
)

P(
    name="nilout",
    title="Let The Dropped Elements Be Collected",
    sig="func DropAll(s []*Node)",
    doc="""DropAll clears every element of s to nil, in place.

The length of s must not change; only the pointers it holds are released
so the nodes they referenced become unreachable.

Examples:

	s := []*Node{{1}}; DropAll(s) => s[0] == nil""",
    extra="""// Node is one payload the slice points at.
type Node struct {
	ID int
}""",
    solution="""clear(s)""",
    tests="""
import "testing"

func TestDropAllNilsEveryElement(t *testing.T) {
	s := []*Node{{ID: 1}, {ID: 2}, {ID: 3}}
	DropAll(s)
	if len(s) != 3 {
		t.Fatalf("len = %d, want 3: the length must not change", len(s))
	}
	for i, p := range s {
		if p != nil {
			t.Errorf("s[%d] = %v, want nil", i, p)
		}
	}
}

func TestDropAllIsVisibleToTheCaller(t *testing.T) {
	s := []*Node{{ID: 1}}
	DropAll(s[:1])
	if s[0] != nil {
		t.Error("the caller's element was not cleared")
	}
}

func TestDropAllEmpty(t *testing.T) {
	DropAll(nil)
}
""",
    context="A pool keeps a fixed-size slice of task pointers and reuses it between rounds. Memory grows round over round even though the tasks are finished.",
    task=[
        "Set every element of `s` to nil.",
        "Keep `len(s)` unchanged — the slice itself is reused.",
        "A nil slice must not panic.",
    ],
    examples=[
        ("s := []*Node{{1},{2}}; DropAll(s)", "[<nil> <nil>]", None),
        ("len(s) after DropAll", "unchanged", "Only the elements are released, not the slice."),
        ("DropAll(nil)", "no panic", None),
    ],
    topics=[
        ("Reachability", "The collector frees what nothing points at; a stale pointer in a live slice is a pointer."),
        ("clear on slices", "`clear(s)` writes the zero value into every element."),
        ("Slices share storage", "Writing through the parameter is visible to the caller."),
    ],
    hint="Reslicing to zero length hides the pointers from you, not from the collector.",
    intuition="A slice you keep alive keeps every pointer inside it alive too. Shortening the slice does not erase the elements past the new length — they are still in the array, still reachable, still pinning their nodes.",
    approach=[
        "Call `clear(s)` to write nil into every element.",
        "Do not reslice or reassign — the caller shares this array.",
    ],
    walkthrough="Three nodes, three pointers. After `clear(s)` the array holds three nils, nothing else references the nodes, and the next collection reclaims them.",
    pitfalls=[
        "`s = s[:0]` — the pointers are still in the backing array.",
        "`s = nil` — rebinds the local parameter only.",
    ],
)

P(
    name="truncate",
    title="Cut The Tail Without Pinning It",
    sig="func Truncate(s []*Node, n int) []*Node",
    doc="""Truncate returns the first n elements of s, clearing the elements it
drops so they no longer keep their payloads reachable.

n is clamped into [0, len(s)]. The result reuses s's storage.

Examples:

	Truncate([]*Node{{1}, {2}}, 1) => the first element only""",
    extra="""// Node is one payload the slice points at.
type Node struct {
	ID int
}""",
    solution="""if n < 0 {
	n = 0
}
if n > len(s) {
	n = len(s)
}
clear(s[n:])
return s[:n]""",
    tests="""
import "testing"

func TestTruncateLength(t *testing.T) {
	s := []*Node{{ID: 1}, {ID: 2}, {ID: 3}}
	if got := Truncate(s, 2); len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
	if got := Truncate(s, 9); len(got) != 3 {
		t.Errorf("len = %d, want 3 when n exceeds the input", len(got))
	}
	if got := Truncate(s, -1); len(got) != 0 {
		t.Errorf("len = %d, want 0 when n is negative", len(got))
	}
}

func TestTruncateClearsTheTail(t *testing.T) {
	s := []*Node{{ID: 1}, {ID: 2}, {ID: 3}}
	Truncate(s, 1)
	for i := 1; i < len(s); i++ {
		if s[i] != nil {
			t.Errorf("s[%d] still holds %v: the dropped payload stays reachable", i, s[i])
		}
	}
	if s[0] == nil {
		t.Error("s[0] was cleared, but it is kept")
	}
}
""",
    context="A queue trims its backlog with `q = q[:n]`. The dropped jobs are gone from the caller's view but the process still holds their memory for as long as the queue lives.",
    task=[
        "Clamp `n` into `[0, len(s)]`.",
        "Clear the elements from `n` onward before returning.",
        "Return `s[:n]` — the storage is reused, not reallocated.",
    ],
    examples=[
        ("Truncate([]*Node{{1},{2},{3}}, 1)", "length 1", "Indices 1 and 2 of the array become nil."),
        ("Truncate(s, 9)", "length 3", "n is clamped up to len(s)."),
        ("Truncate(s, -1)", "length 0", "n is clamped to 0."),
    ],
    topics=[
        ("Reslicing does not erase", "Elements past the new length stay in the backing array."),
        ("clear on a sub-slice", "`clear(s[n:])` releases exactly the dropped range."),
        ("Clamping", "Reject out-of-range n before it indexes anything."),
    ],
    hint="`s[:n]` is the answer; something has to happen to `s[n:]` first.",
    intuition="Length is a view, not a fence. The array still holds the tail, and the collector follows the array, not your view of it. Clear the tail, then narrow the view.",
    approach=[
        "Clamp `n` low and high.",
        "`clear(s[n:])` so the dropped pointers go to nil.",
        "Return `s[:n]`.",
    ],
    walkthrough="With three nodes and n = 1: `clear(s[1:])` nils indices 1 and 2, then `s[:1]` is returned. Node 1 stays reachable, nodes 2 and 3 do not.",
    pitfalls=[
        "Clearing after reslicing — `s[:n][n:]` is empty, so it clears nothing.",
        "Forgetting the clamp: `s[n:]` with n > len(s) panics.",
    ],
)

P(
    name="resetbuf",
    title="Keep The Capacity, Drop The Contents",
    sig="func Reset(buf []byte) []byte",
    doc="""Reset returns buf emptied for reuse, keeping the capacity it already
has so the next writer does not have to allocate again.

Examples:

	Reset(make([]byte, 8, 64)) => length 0, capacity 64""",
    solution="""return buf[:0]""",
    tests="""
import "testing"

func TestResetEmptiesAndKeepsCapacity(t *testing.T) {
	buf := make([]byte, 8, 64)
	out := Reset(buf)
	if len(out) != 0 {
		t.Errorf("len = %d, want 0", len(out))
	}
	if cap(out) != 64 {
		t.Errorf("cap = %d, want 64: the capacity must survive the reset", cap(out))
	}
}

func TestResetReusesTheSameArray(t *testing.T) {
	buf := make([]byte, 0, 8)
	buf = append(buf, 'a', 'b')
	out := append(Reset(buf), 'z')
	if len(out) != 1 || out[0] != 'z' {
		t.Fatalf("out = %q, want \\"z\\"", out)
	}
	if &buf[:1][0] != &out[0] {
		t.Error("the reset buffer allocated a new array instead of reusing buf")
	}
}

func TestResetNil(t *testing.T) {
	if got := Reset(nil); len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}
""",
    context="A line-oriented writer allocates a fresh `[]byte` for every record. The records are small, the record count is not, and the allocator shows up at the top of the profile.",
    task=[
        "Return a slice of length 0 that still owns `buf`'s array and capacity.",
        "A nil input must return an empty result without panicking.",
    ],
    examples=[
        ("Reset(make([]byte, 8, 64))", "len 0, cap 64", None),
        ("append(Reset(buf), 'z')", "writes into buf's existing array", "No allocation for the next record."),
        ("Reset(nil)", "len 0", None),
    ],
    topics=[
        ("Length vs capacity", "Length is what you can index; capacity is what you already own."),
        ("Buffer reuse", "Resetting to `[:0]` is the standard reuse idiom."),
        ("Zero-value slices", "`nil[:0]` is legal and yields an empty slice."),
    ],
    hint="One expression. It is a reslice.",
    intuition="Allocating is the expensive part, not zeroing. `buf[:0]` throws away the *view* while keeping the array, so the next round of appends writes into memory you have already paid for.",
    approach=[
        "Return `buf[:0]`.",
    ],
    walkthrough="`make([]byte, 8, 64)` owns a 64-byte array. `buf[:0]` still points at it with cap 64, so the next 64 bytes of appends need no allocation.",
    pitfalls=[
        "`return nil` or `return []byte{}` — correct length, capacity thrown away.",
        "Reusing a buffer whose old contents someone still holds a view of.",
    ],
)

P(
    name="chunk",
    title="Batch A Slice Into Fixed-Size Windows",
    sig="func Chunk(s []int, n int) [][]int",
    doc="""Chunk splits s into consecutive groups of at most n elements.

The last group holds the remainder. For n <= 0 the result is nil. The
groups are views into s — no element is copied.

Examples:

	Chunk([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}""",
    solution="""if n <= 0 {
	return nil
}
out := make([][]int, 0, (len(s)+n-1)/n)
for i := 0; i < len(s); i += n {
	end := i + n
	if end > len(s) {
		end = len(s)
	}
	out = append(out, s[i:end])
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chunk = %v, want %v", got, want)
	}
	if got := Chunk([]int{1, 2}, 5); !reflect.DeepEqual(got, [][]int{{1, 2}}) {
		t.Errorf("Chunk = %v, want [[1 2]]", got)
	}
	if got := Chunk(nil, 3); len(got) != 0 {
		t.Errorf("Chunk(nil) = %v, want empty", got)
	}
	if got := Chunk([]int{1}, 0); got != nil {
		t.Errorf("Chunk(s, 0) = %v, want nil", got)
	}
}

func TestChunkGroupsAreViews(t *testing.T) {
	s := []int{1, 2, 3, 4}
	g := Chunk(s, 2)
	g[0][0] = 99
	if s[0] != 99 {
		t.Error("the groups copied the elements; they must be views into s")
	}
}
""",
    context="A batch uploader sends records in groups of 500. The current splitter copies every record into a fresh group and doubles the job's peak memory.",
    task=[
        "Split `s` into consecutive groups of at most `n`.",
        "The groups must be views into `s`, not copies.",
        "Return nil when `n <= 0`; preallocate the outer slice to its exact group count.",
    ],
    examples=[
        ("Chunk([]int{1,2,3,4,5}, 2)", "[[1 2] [3 4] [5]]", "The last group holds the remainder."),
        ("Chunk([]int{1,2}, 5)", "[[1 2]]", "n larger than the input gives one group."),
        ("Chunk([]int{1}, 0)", "<nil>", None),
    ],
    topics=[
        ("Sub-slicing", "`s[i:end]` is a view — no elements move."),
        ("Ceiling division", "`(len+n-1)/n` is the group count, so the outer slice is sized once."),
        ("Boundary clamping", "The final window must stop at `len(s)`."),
    ],
    hint="Two allocations at most: the outer slice, and nothing else.",
    intuition="Splitting is a question of where the boundaries are, not of moving data. Every group can point into the original array; only the little slice of headers is new.",
    approach=[
        "Reject `n <= 0`.",
        "Preallocate the outer slice with the ceiling-divided group count.",
        "Step `i` by `n`, clamp `end` to `len(s)`, append `s[i:end]`.",
    ],
    walkthrough="For five elements and n = 2, the group count is (5+1)/2 = 3. The windows are [0:2], [2:4] and [4:5] — the last one clamped.",
    pitfalls=[
        "`s[i : i+n]` without the clamp panics on the last group.",
        "Copying each group into a fresh slice — correct output, wrong memory behaviour.",
    ],
)

P(
    name="rows",
    title="One Allocation For The Whole Grid",
    sig="func Rows(r, c int) [][]int",
    doc="""Rows returns an r-by-c grid of zeros whose rows are consecutive
windows into a single backing array.

Allocating each row separately costs r allocations and scatters the grid
across the heap; this must cost two.

Examples:

	Rows(2, 3) => a 2x3 grid, rows 0 and 1 adjacent in memory""",
    solution="""if r <= 0 || c <= 0 {
	return nil
}
flat := make([]int, r*c)
out := make([][]int, r)
for i := range out {
	out[i] = flat[i*c : (i+1)*c : (i+1)*c]
}
return out""",
    tests="""
import "testing"

func TestRowsShape(t *testing.T) {
	g := Rows(2, 3)
	if len(g) != 2 {
		t.Fatalf("rows = %d, want 2", len(g))
	}
	for i, row := range g {
		if len(row) != 3 {
			t.Fatalf("row %d has length %d, want 3", i, len(row))
		}
		for _, v := range row {
			if v != 0 {
				t.Fatalf("row %d is not zeroed", i)
			}
		}
	}
	if g := Rows(0, 3); g != nil {
		t.Errorf("Rows(0,3) = %v, want nil", g)
	}
}

func TestRowsShareOneArray(t *testing.T) {
	g := Rows(2, 3)
	if &g[0][0] == &g[1][0] {
		t.Fatal("rows overlap")
	}
	g[0] = append(g[0], 7)
	if g[1][0] == 7 {
		t.Error("row 0 spilled into row 1: cap each row with a three-index slice")
	}
}

func TestRowsAllocatesTwice(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { _ = Rows(32, 32) }); n > 2 {
		t.Errorf("Rows made %v allocations, want at most 2", n)
	}
}
""",
    context="An image filter builds its working buffer as a slice of independently allocated rows. The rows land all over the heap and the filter's inner loop misses cache on every row change.",
    task=[
        "Return an `r` by `c` grid of zeros.",
        "All rows must be windows into one flat array — at most two allocations total.",
        "Return nil when `r <= 0` or `c <= 0`.",
    ],
    examples=[
        ("Rows(2, 3)", "[[0 0 0] [0 0 0]]", None),
        ("&g[0][2] and &g[1][0]", "adjacent addresses", "Row 1 begins where row 0 ends."),
        ("Rows(0, 3)", "<nil>", None),
    ],
    topics=[
        ("Backing arrays", "Many slices can be windows into one allocation."),
        ("Three-index slicing", "`flat[a:b:b]` caps each row so an append cannot spill into the next."),
        ("Locality", "Contiguous memory is what makes the traversal cache-friendly."),
    ],
    hint="Allocate `r*c` ints first, then hand out windows of `c`.",
    intuition="A `[][]int` is a slice of headers. Nothing says those headers must point at different allocations — carving one flat array into `r` windows gives the same API with one allocation and perfect locality.",
    approach=[
        "Reject non-positive dimensions.",
        "Allocate `flat := make([]int, r*c)` and `out := make([][]int, r)`.",
        "Point row `i` at `flat[i*c : (i+1)*c : (i+1)*c]`.",
    ],
    walkthrough="Rows(2,3) allocates six ints. Row 0 is `flat[0:3:3]`, row 1 is `flat[3:6:6]`, so `&g[0][2]` and `&g[1][0]` are neighbours.",
    pitfalls=[
        "Omitting the capacity bound — an `append` to row 0 would overwrite row 1.",
        "`make([][]int, r)` alone leaves every row nil.",
    ],
)

P(
    name="mapprealloc",
    title="Size The Map Before Filling It",
    sig="func Count(words []string) map[string]int",
    doc="""Count returns how many times each word appears in words.

The map must be created with a size hint so it does not rehash its way up
from nothing while the loop runs.

Examples:

	Count([]string{"a", "b", "a"}) => map[a:2 b:1]""",
    solution="""m := make(map[string]int, len(words))
for _, w := range words {
	m[w]++
}
return m""",
    tests="""
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	got := Count([]string{"a", "b", "a", "c", "a"})
	want := map[string]int{"a": 3, "b": 1, "c": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Count = %v, want %v", got, want)
	}
	if got := Count(nil); len(got) != 0 {
		t.Errorf("Count(nil) = %v, want empty", got)
	}
}

func TestCountLarge(t *testing.T) {
	in := make([]string, 0, 1000)
	for i := 0; i < 1000; i++ {
		in = append(in, fmt.Sprintf("w%d", i%50))
	}
	got := Count(in)
	if len(got) != 50 {
		t.Fatalf("distinct = %d, want 50", len(got))
	}
	if got["w0"] != 20 {
		t.Errorf("count = %d, want 20", got["w0"])
	}
}

func TestSizeHint(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "mapprealloc.go", nil, 0)
	if err != nil {
		return
	}
	sized := false
	ast.Inspect(f, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "make" && len(c.Args) >= 2 {
				sized = true
			}
		}
		return true
	})
	if !sized {
		t.Logf("WARN: give make a size hint — make(map[string]int, len(words))")
	}
}
""",
    context="A log analyser counts a few million tokens. The counting map starts empty and rehashes a dozen times on the way up, copying every bucket each time.",
    task=[
        "Return the count of each distinct word.",
        "Create the map with a size hint so it does not grow from empty.",
        "A nil input returns an empty map.",
    ],
    examples=[
        ('Count([]string{"a","b","a"})', "map[a:2 b:1]", None),
        ("Count(nil)", "map[]", None),
        ('Count([]string{"x"})', "map[x:1]", None),
    ],
    topics=[
        ("Map size hints", "`make(map[K]V, n)` allocates enough buckets for n entries up front."),
        ("Rehashing", "Growing a map reallocates the bucket array and moves entries."),
        ("Zero value of a map entry", "`m[w]++` works on a missing key because the value starts at 0."),
    ],
    hint="`make` takes a second argument for maps too.",
    intuition="A map grows by allocating a bigger bucket array and migrating entries. If you know roughly how many keys are coming, saying so up front skips every intermediate size.",
    approach=[
        "`make(map[string]int, len(words))`.",
        "Loop and `m[w]++`.",
        "Return the map.",
    ],
    walkthrough="For 1000 tokens over 50 distinct words, an unsized map grows through several bucket-array sizes; a sized one starts big enough and never migrates.",
    pitfalls=[
        "Over-hinting with a huge number wastes memory — `len(words)` is an upper bound, and that is fine here.",
        "The hint is advice, not a cap; the map still grows if you exceed it.",
    ],
)

P(
    name="builder",
    title="Join Without The Intermediate Strings",
    sig="func Join(parts []string, sep string) string",
    doc="""Join concatenates parts separated by sep.

Strings are immutable, so `s += p` allocates a new string every round.
Build the result in one growing buffer instead.

Examples:

	Join([]string{"a", "b"}, "-") => "a-b" """,
    imports=['"strings"'],
    solution="""if len(parts) == 0 {
	return ""
}
n := len(sep) * (len(parts) - 1)
for _, p := range parts {
	n += len(p)
}
var b strings.Builder
b.Grow(n)
for i, p := range parts {
	if i > 0 {
		b.WriteString(sep)
	}
	b.WriteString(p)
}
return b.String()""",
    tests="""
import "testing"

func TestJoin(t *testing.T) {
	if got := Join([]string{"a", "b", "c"}, "-"); got != "a-b-c" {
		t.Errorf("Join = %q, want \\"a-b-c\\"", got)
	}
	if got := Join([]string{"solo"}, ","); got != "solo" {
		t.Errorf("Join = %q, want \\"solo\\"", got)
	}
	if got := Join(nil, ","); got != "" {
		t.Errorf("Join = %q, want empty", got)
	}
	if got := Join([]string{"a", "b"}, ""); got != "ab" {
		t.Errorf("Join = %q, want \\"ab\\"", got)
	}
}

func TestJoinAllocationsAreBounded(t *testing.T) {
	parts := make([]string, 64)
	for i := range parts {
		parts[i] = "chunk"
	}
	if n := testing.AllocsPerRun(50, func() { _ = Join(parts, ", ") }); n > 3 {
		t.Errorf("Join made %v allocations, want at most 3: grow the buffer once", n)
	}
}
""",
    context="A CSV header is assembled with `line += col + \",\"`. With 200 columns that is 400 short-lived strings per row, and the rows are the whole file.",
    task=[
        "Concatenate `parts` with `sep` between them.",
        "Size the buffer once before writing — at most 3 allocations for a 64-part join.",
        "An empty input returns the empty string.",
    ],
    examples=[
        ('Join([]string{"a","b","c"}, "-")', '"a-b-c"', None),
        ('Join([]string{"solo"}, ",")', '"solo"', "No separator with a single part."),
        ("Join(nil, \",\")", '""', None),
    ],
    topics=[
        ("String immutability", "`+=` cannot extend a string; it allocates a new one and copies both sides."),
        ("strings.Builder", "A growable byte buffer that hands out its bytes as a string without a final copy."),
        ("Grow", "One `Grow(n)` replaces every doubling step."),
    ],
    hint="Compute the exact final length before you write the first byte.",
    intuition="Every `+=` on a string is an allocate-and-copy of everything so far, so joining n parts costs O(n²) bytes copied. A builder writes into one buffer that only grows — and if you pre-size it, it never grows at all.",
    approach=[
        "Return early for an empty input.",
        "Sum the part lengths plus `len(sep) * (len(parts)-1)`.",
        "`b.Grow(n)`, then write parts with separators.",
        "Return `b.String()`.",
    ],
    walkthrough="Joining 64 five-byte parts with a two-byte separator needs 64*5 + 63*2 = 446 bytes. One `Grow(446)` allocates once; `+=` would have allocated 64 times and copied about 14 KB.",
    pitfalls=[
        "Adding the separator after every part and trimming the tail — an extra copy for nothing.",
        "Forgetting that `len` on a string is bytes, which is exactly what the buffer needs.",
    ],
)

P(
    name="inplacefilter",
    title="Filter Without A Second Slice",
    sig="func KeepEven(s []int) []int",
    doc="""KeepEven returns the even elements of s, in order, reusing s's own
storage rather than allocating a result.

The elements of s beyond the returned length are unspecified.

Examples:

	KeepEven([]int{1, 2, 3, 4}) => []int{2, 4}""",
    solution="""k := 0
for _, v := range s {
	if v%2 == 0 {
		s[k] = v
		k++
	}
}
return s[:k]""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestKeepEven(t *testing.T) {
	if got := KeepEven([]int{1, 2, 3, 4, 6}); !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("KeepEven = %v, want [2 4 6]", got)
	}
	if got := KeepEven([]int{1, 3}); len(got) != 0 {
		t.Errorf("KeepEven = %v, want empty", got)
	}
	if got := KeepEven(nil); len(got) != 0 {
		t.Errorf("KeepEven(nil) = %v, want empty", got)
	}
	if got := KeepEven([]int{-2, -1, 0}); !reflect.DeepEqual(got, []int{-2, 0}) {
		t.Errorf("KeepEven = %v, want [-2 0]", got)
	}
}

func TestKeepEvenAllocatesNothing(t *testing.T) {
	s := make([]int, 256)
	for i := range s {
		s[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { _ = KeepEven(s) }); n != 0 {
		t.Errorf("KeepEven made %v allocations, want 0: filter in place", n)
	}
}
""",
    context="A stream stage filters batches of a few hundred thousand records and allocates a fresh result slice for every batch. The batches are short-lived and the collector is doing all the work.",
    task=[
        "Return the even elements of `s` in order.",
        "Reuse `s`'s storage — the function must allocate nothing.",
    ],
    examples=[
        ("KeepEven([]int{1,2,3,4,6})", "[2 4 6]", None),
        ("KeepEven([]int{1,3})", "[]", "Nothing survives; the length is 0."),
        ("KeepEven([]int{-2,-1,0})", "[-2 0]", "0 and negative evens count."),
    ],
    topics=[
        ("In-place compaction", "A write index trails the read index and never overtakes it."),
        ("Reslicing as the result", "`s[:k]` reports the kept count without allocating."),
        ("Destructive helpers", "Reusing the caller's array is a documented contract, not a secret."),
    ],
    hint="Two indices, one array: where you are reading and where you are writing.",
    intuition="The kept elements can only move left, never right, so the read cursor is always at or ahead of the write cursor. That means you can overwrite as you go without ever losing an element you still need.",
    approach=[
        "Keep a write index `k` starting at 0.",
        "Range over `s`; when the element qualifies, store it at `s[k]` and bump `k`.",
        "Return `s[:k]`.",
    ],
    walkthrough="For [1 2 3 4 6]: 1 is skipped; 2 goes to s[0], k=1; 3 skipped; 4 goes to s[1], k=2; 6 goes to s[2], k=3. Return s[:3] = [2 4 6].",
    pitfalls=[
        "Appending to a fresh slice — correct output, one allocation per call.",
        "Assuming the tail of `s` is untouched afterwards; it is not.",
    ],
)

P(
    name="copyarray",
    title="Arrays Are Values, Slices Are Not",
    sig="func Bump(a [4]int) [4]int",
    doc="""Bump returns a copy of a with every element increased by one.

The caller's array must not change: an array parameter is a value, and
the returned array is a separate value again.

Examples:

	Bump([4]int{1, 2, 3, 4}) => [4]int{2, 3, 4, 5}""",
    solution="""for i := range a {
	a[i]++
}
return a""",
    tests="""
import "testing"

func TestBump(t *testing.T) {
	if got := Bump([4]int{1, 2, 3, 4}); got != [4]int{2, 3, 4, 5} {
		t.Errorf("Bump = %v, want [2 3 4 5]", got)
	}
	if got := Bump([4]int{}); got != [4]int{1, 1, 1, 1} {
		t.Errorf("Bump = %v, want [1 1 1 1]", got)
	}
	if got := Bump([4]int{-1, 0, 0, 0}); got != [4]int{0, 1, 1, 1} {
		t.Errorf("Bump = %v, want [0 1 1 1]", got)
	}
}

func TestBumpDoesNotTouchTheCaller(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	Bump(a)
	if a != [4]int{1, 2, 3, 4} {
		t.Errorf("a = %v, want [1 2 3 4]: the caller's array changed", a)
	}
}

func TestBumpAllocatesNothing(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	if n := testing.AllocsPerRun(100, func() { _ = Bump(a) }); n != 0 {
		t.Errorf("Bump made %v allocations, want 0: arrays are values", n)
	}
}
""",
    context="A geometry helper takes a fixed-size array and mutates it, expecting the caller to see the change. The caller never does, and the confusion costs an afternoon.",
    task=[
        "Return an array whose elements are one greater than `a`'s.",
        "The caller's array must be unchanged.",
        "No allocations — the array lives on the stack.",
    ],
    examples=[
        ("Bump([4]int{1,2,3,4})", "[2 3 4 5]", None),
        ("a := [4]int{1,2,3,4}; Bump(a)", "a is still [1 2 3 4]", "The parameter was a copy."),
        ("Bump([4]int{})", "[1 1 1 1]", None),
    ],
    topics=[
        ("Array value semantics", "Passing or assigning an array copies every element."),
        ("Array vs slice", "A slice copies a header; an array copies the data."),
        ("Stack allocation", "A small fixed-size value needs no heap."),
    ],
    hint="You may mutate the parameter freely — it is already your copy.",
    intuition="`[4]int` is a value like an `int` is. The function receives its own copy, so mutating it is private, and returning it copies again. Nothing is shared, and nothing is allocated.",
    approach=[
        "Range over the parameter and increment each element.",
        "Return the parameter.",
    ],
    walkthrough="`Bump(a)` copies four ints onto the stack, increments them, and copies four ints back out. `a` in the caller is untouched throughout.",
    pitfalls=[
        "Expecting a mutation to reach the caller — take `*[4]int` or a slice for that.",
        "Allocating a fresh array to fill; the parameter already is one.",
    ],
)

P(
    name="sumlocal",
    title="A Function That Touches No Heap",
    sig="func Sum(s []int) int",
    doc="""Sum returns the total of s.

The function must not allocate: every value it needs fits in a local
variable, and the input is only read.

Examples:

	Sum([]int{1, 2, 3}) => 6""",
    solution="""total := 0
for _, v := range s {
	total += v
}
return total""",
    tests="""
import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, want 0", got)
	}
	if got := Sum([]int{-3, 3}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}

func TestSumAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	for i := range s {
		s[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { _ = Sum(s) }); n != 0 {
		t.Errorf("Sum made %v allocations, want 0", n)
	}
}
""",
    context="A hot metrics loop calls a sum helper millions of times. The helper looks trivial, yet the allocation profile blames it.",
    task=[
        "Return the sum of the elements of `s`.",
        "The function must make zero allocations.",
        "A nil input sums to 0.",
    ],
    examples=[
        ("Sum([]int{1,2,3})", "6", None),
        ("Sum(nil)", "0", "The empty sum."),
        ("Sum([]int{-3,3})", "0", None),
    ],
    topics=[
        ("Stack vs heap", "A local that never outlives the call stays on the stack."),
        ("Reading a slice", "Ranging a slice copies elements into a loop variable — still no heap."),
        ("AllocsPerRun", "Zero allocations is a testable property, not a hope."),
    ],
    hint="An `int` accumulator. Nothing else.",
    intuition="Nothing about summing needs memory that outlives the call. Keep the running total in a plain local and the compiler keeps it in a register.",
    approach=[
        "Declare a local `total`.",
        "Range and add.",
        "Return `total`.",
    ],
    walkthrough="Summing 512 ints touches 4 KB of the caller's array and one stack slot. No allocation is ever requested.",
    pitfalls=[
        "`total := new(int)` — now the pointer may escape and the count is 1.",
        "Building an intermediate slice or using `fmt` anywhere in a hot helper.",
    ],
)

P(
    name="repeatbytes",
    title="Repeat With The Length You Already Know",
    sig="func Repeat(b []byte, n int) []byte",
    doc="""Repeat returns n concatenated copies of b in a freshly allocated
slice that shares nothing with b.

For n <= 0 the result is empty. The allocation must happen once, at the
final size.

Examples:

	Repeat([]byte("ab"), 2) => []byte("abab")""",
    solution="""if n <= 0 || len(b) == 0 {
	return []byte{}
}
out := make([]byte, len(b)*n)
for i := 0; i < n; i++ {
	copy(out[i*len(b):], b)
}
return out""",
    tests="""
import (
	"bytes"
	"testing"
)

func TestRepeat(t *testing.T) {
	if got := Repeat([]byte("ab"), 3); !bytes.Equal(got, []byte("ababab")) {
		t.Errorf("Repeat = %q, want \\"ababab\\"", got)
	}
	if got := Repeat([]byte("x"), 0); len(got) != 0 {
		t.Errorf("Repeat = %q, want empty", got)
	}
	if got := Repeat(nil, 4); len(got) != 0 {
		t.Errorf("Repeat = %q, want empty", got)
	}
	if got := Repeat([]byte("ab"), -1); len(got) != 0 {
		t.Errorf("Repeat = %q, want empty", got)
	}
}

func TestRepeatIsIndependent(t *testing.T) {
	b := []byte("ab")
	out := Repeat(b, 2)
	b[0] = 'z'
	if out[0] != 'a' {
		t.Error("the result shares storage with the input")
	}
}

func TestRepeatAllocatesOnce(t *testing.T) {
	b := []byte("abcdefgh")
	if n := testing.AllocsPerRun(100, func() { _ = Repeat(b, 64) }); n > 1 {
		t.Errorf("Repeat made %v allocations, want 1", n)
	}
}
""",
    context="A padding helper builds its filler by appending the pattern in a loop. For a 64 KB pad it reallocates seventeen times.",
    task=[
        "Return `n` copies of `b` concatenated.",
        "The result must share no storage with `b`.",
        "One allocation only; `n <= 0` or an empty `b` gives an empty result.",
    ],
    examples=[
        ('Repeat([]byte("ab"), 3)', '"ababab"', None),
        ('Repeat([]byte("x"), 0)', '""', None),
        ("Repeat(nil, 4)", '""', None),
    ],
    topics=[
        ("Known output size", "`len(b)*n` is the exact final length."),
        ("copy into a window", "`copy(out[i*len(b):], b)` writes each repetition in place."),
        ("Independence", "A fresh array means later writes to `b` are invisible."),
    ],
    hint="Allocate `len(b)*n` bytes first; then it is `n` copies into windows.",
    intuition="The output length is completely determined by the inputs, so there is no reason to discover it by growing. Allocate it, then fill it.",
    approach=[
        "Handle the empty cases.",
        "`make([]byte, len(b)*n)`.",
        "`copy` `b` into the window starting at `i*len(b)` for each i.",
    ],
    walkthrough="`Repeat([]byte(\"ab\"), 3)` allocates six bytes, then copies \"ab\" at offsets 0, 2 and 4.",
    pitfalls=[
        "`append(out, b...)` in a loop — correct, but that is the reallocating version.",
        "Forgetting that `len(b) == 0` makes `len(b)*n` zero, which is fine, but `n <= 0` must be checked too.",
    ],
)

P(
    name="growslice",
    title="Make Room Before The Appends",
    sig="func Grow(s []int, n int) []int",
    doc="""Grow returns s with capacity for at least n more elements, without
changing its length or contents.

If s already has the room, it is returned untouched and nothing is
allocated. n < 0 is treated as 0.

Examples:

	Grow(make([]int, 2, 2), 8) => length 2, capacity at least 10""",
    solution="""if n < 0 {
	n = 0
}
if cap(s)-len(s) >= n {
	return s
}
out := make([]int, len(s), len(s)+n)
copy(out, s)
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestGrowCapacity(t *testing.T) {
	s := make([]int, 2, 2)
	s[0], s[1] = 1, 2
	out := Grow(s, 8)
	if len(out) != 2 {
		t.Errorf("len = %d, want 2: Grow must not change the length", len(out))
	}
	if cap(out) < 10 {
		t.Errorf("cap = %d, want at least 10", cap(out))
	}
	if !reflect.DeepEqual(out, []int{1, 2}) {
		t.Errorf("contents = %v, want [1 2]", out)
	}
}

func TestGrowIsANoOpWhenTheRoomExists(t *testing.T) {
	s := make([]int, 1, 32)
	if n := testing.AllocsPerRun(100, func() { _ = Grow(s, 4) }); n != 0 {
		t.Errorf("Grow made %v allocations, want 0 when the capacity already fits", n)
	}
}

func TestGrowNegative(t *testing.T) {
	s := make([]int, 1, 1)
	if out := Grow(s, -5); len(out) != 1 {
		t.Errorf("len = %d, want 1", len(out))
	}
}
""",
    context="A decoder knows the record count from the header but still lets `append` discover the size, paying a reallocation and a full copy at every doubling.",
    task=[
        "Return a slice with the same length and contents as `s` and room for `n` more elements.",
        "Allocate nothing when the spare capacity already covers `n`.",
        "Treat `n < 0` as 0.",
    ],
    examples=[
        ("Grow(make([]int,2,2), 8)", "len 2, cap >= 10", None),
        ("Grow(make([]int,1,32), 4)", "the same slice, no allocation", "The room is already there."),
        ("Grow(s, -5)", "s unchanged", None),
    ],
    topics=[
        ("Spare capacity", "`cap(s) - len(s)` is what an append can use for free."),
        ("make with length and capacity", "`make([]int, len, cap)` keeps the contents indexable and reserves the rest."),
        ("Amortised growth", "Reserving once beats doubling repeatedly."),
    ],
    hint="Compare `cap(s)-len(s)` with `n` before doing anything.",
    intuition="Growth is only expensive when it is a surprise. Once you can say how much more is coming, one allocation covers all of it and the appends after it are pure writes.",
    approach=[
        "Clamp `n` at 0.",
        "If `cap(s)-len(s) >= n`, return `s`.",
        "Otherwise allocate `make([]int, len(s), len(s)+n)`, copy, and return it.",
    ],
    walkthrough="A slice of len 2, cap 2 needs room for 8: spare is 0, so a new array of cap 10 is allocated and the two elements copied. A slice of len 1, cap 32 needs room for 4: spare is 31, so it is returned as is.",
    pitfalls=[
        "`make([]int, len(s)+n)` — that changes the length, not just the capacity.",
        "Always reallocating; the no-op case is what the test measures.",
    ],
)

P(
    name="fillshared",
    title="Write Through The Caller's Array",
    sig="func Fill(s []int, v int)",
    doc="""Fill sets every element of s to v.

The parameter is a view onto the caller's array, so the writes must be
visible to the caller. Nothing is allocated and nothing is returned.

Examples:

	s := []int{1, 2}; Fill(s, 7) => s is [7 7]""",
    solution="""for i := range s {
	s[i] = v
}""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	s := []int{1, 2, 3}
	Fill(s, 7)
	if !reflect.DeepEqual(s, []int{7, 7, 7}) {
		t.Errorf("s = %v, want [7 7 7]", s)
	}
}

func TestFillWritesThroughAView(t *testing.T) {
	s := []int{1, 2, 3, 4}
	Fill(s[1:3], 0)
	if !reflect.DeepEqual(s, []int{1, 0, 0, 4}) {
		t.Errorf("s = %v, want [1 0 0 4]: only the view's range may be written", s)
	}
}

func TestFillEmpty(t *testing.T) {
	Fill(nil, 1)
	Fill([]int{}, 1)
}

func TestFillAllocatesNothing(t *testing.T) {
	s := make([]int, 256)
	if n := testing.AllocsPerRun(100, func() { Fill(s, 3) }); n != 0 {
		t.Errorf("Fill made %v allocations, want 0", n)
	}
}
""",
    context="A helper meant to zero a scratch region builds a fresh slice of zeros and assigns it to its parameter. The scratch region is never actually zeroed.",
    task=[
        "Set every element of `s` to `v`.",
        "The writes must be visible through the caller's slice, including when `s` is a sub-slice view.",
        "No allocations, no return value.",
    ],
    examples=[
        ("s := []int{1,2,3}; Fill(s, 7)", "s is [7 7 7]", None),
        ("s := []int{1,2,3,4}; Fill(s[1:3], 0)", "s is [1 0 0 4]", "Only the view's range is written."),
        ("Fill(nil, 1)", "no panic", None),
    ],
    topics=[
        ("Slices share storage", "The parameter is a copy of the header, pointing at the same array."),
        ("Views", "`s[1:3]` writes into the middle of the original array."),
        ("Assignment vs mutation", "`s = ...` rebinds the local; `s[i] = ...` mutates the array."),
    ],
    hint="Index into the parameter. Do not assign to it.",
    intuition="Passing a slice copies three words but shares one array. That is why mutating elements reaches the caller and reassigning the parameter does not.",
    approach=[
        "Range over the indices of `s`.",
        "Assign `v` to each index.",
    ],
    walkthrough="`Fill(s[1:3], 0)` receives a header with pointer `&s[1]` and length 2, so it writes indices 1 and 2 of the caller's array and leaves 0 and 3 alone.",
    pitfalls=[
        "`s = make([]int, len(s))` — allocates and is invisible outside.",
        "Using `range s` with the value variable and assigning to it; that writes to the loop copy.",
    ],
)

P(
    name="upperascii",
    title="Upper-Case The Bytes You Were Given",
    sig="func Upper(b []byte) []byte",
    doc="""Upper upper-cases the ASCII letters of b in place and returns b.

Non-letters and non-ASCII bytes are left alone. Nothing is allocated —
the caller's buffer is the working buffer.

Examples:

	Upper([]byte("go1")) => []byte("GO1")""",
    solution="""for i, c := range b {
	if c >= 'a' && c <= 'z' {
		b[i] = c - 'a' + 'A'
	}
}
return b""",
    tests="""
import (
	"bytes"
	"testing"
)

func TestUpper(t *testing.T) {
	if got := Upper([]byte("go1 x")); !bytes.Equal(got, []byte("GO1 X")) {
		t.Errorf("Upper = %q, want \\"GO1 X\\"", got)
	}
	if got := Upper([]byte("ALREADY")); !bytes.Equal(got, []byte("ALREADY")) {
		t.Errorf("Upper = %q, want \\"ALREADY\\"", got)
	}
	if got := Upper(nil); len(got) != 0 {
		t.Errorf("Upper(nil) = %q, want empty", got)
	}
}

func TestUpperIsInPlace(t *testing.T) {
	b := []byte("abc")
	out := Upper(b)
	if !bytes.Equal(b, []byte("ABC")) {
		t.Errorf("b = %q, want \\"ABC\\": the caller's buffer was not modified", b)
	}
	if &out[0] != &b[0] {
		t.Error("Upper returned a different array")
	}
}

func TestUpperAllocatesNothing(t *testing.T) {
	b := bytes.Repeat([]byte("abcdef"), 32)
	if n := testing.AllocsPerRun(100, func() { _ = Upper(b) }); n != 0 {
		t.Errorf("Upper made %v allocations, want 0", n)
	}
}
""",
    context="A header normaliser converts each `[]byte` to a string, upper-cases it, and converts back. Two allocations per header, on every request.",
    task=[
        "Upper-case the ASCII letters `a`-`z` of `b`, in place.",
        "Return `b` itself — the same array the caller passed.",
        "No allocations; leave every other byte untouched.",
    ],
    examples=[
        ('Upper([]byte("go1 x"))', '"GO1 X"', "Digits and spaces are untouched."),
        ('Upper([]byte("ALREADY"))', '"ALREADY"', None),
        ("Upper(nil)", '""', None),
    ],
    topics=[
        ("[]byte is mutable, string is not", "Working on the byte slice avoids both conversions."),
        ("ASCII arithmetic", "`c - 'a' + 'A'` is the case flip for ASCII letters."),
        ("In-place transforms", "Returning the input makes the call chainable without copying."),
    ],
    hint="`strings.ToUpper` would cost you two conversions. You already hold the bytes.",
    intuition="`[]byte(s)` and `string(b)` each copy. When the data is already a mutable byte slice, transform it where it is.",
    approach=[
        "Range over `b` by index and value.",
        "When the byte is in `a`..`z`, write back the upper-case byte.",
        "Return `b`.",
    ],
    walkthrough="\"go1 x\": 'g' and 'o' are shifted by -32, '1' and ' ' fail the range test, 'x' is shifted. The array now reads \"GO1 X\" and no new memory was touched.",
    pitfalls=[
        "Ranging a `[]byte` with `for i, c := range` gives bytes, not runes — which is what is wanted here, but the same loop over a `string` gives runes.",
        "Touching bytes above 0x7F; multi-byte UTF-8 must be left alone.",
    ],
)

P(
    name="dedupesorted",
    title="Collapse Runs Without New Memory",
    sig="func Dedupe(s []int) []int",
    doc="""Dedupe removes consecutive duplicates from the sorted slice s,
reusing s's storage, and returns the deduplicated prefix.

Elements past the returned length are unspecified.

Examples:

	Dedupe([]int{1, 1, 2, 3, 3}) => []int{1, 2, 3}""",
    solution="""if len(s) == 0 {
	return s
}
k := 1
for i := 1; i < len(s); i++ {
	if s[i] != s[k-1] {
		s[k] = s[i]
		k++
	}
}
return s[:k]""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	if got := Dedupe([]int{1, 1, 2, 3, 3, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Dedupe = %v, want [1 2 3]", got)
	}
	if got := Dedupe([]int{5}); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("Dedupe = %v, want [5]", got)
	}
	if got := Dedupe(nil); len(got) != 0 {
		t.Errorf("Dedupe(nil) = %v, want empty", got)
	}
	if got := Dedupe([]int{2, 2, 2}); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Dedupe = %v, want [2]", got)
	}
}

func TestDedupeAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	for i := range s {
		s[i] = i / 2
	}
	if n := testing.AllocsPerRun(100, func() { _ = Dedupe(s) }); n != 0 {
		t.Errorf("Dedupe made %v allocations, want 0", n)
	}
}
""",
    context="A sorted index is deduplicated into a new slice on every rebuild. The index is large, the duplicates are few, and the copy dominates the rebuild time.",
    task=[
        "Collapse runs of equal elements in the sorted input `s`.",
        "Reuse `s`'s storage and return the deduplicated prefix.",
        "Zero allocations.",
    ],
    examples=[
        ("Dedupe([]int{1,1,2,3,3,3})", "[1 2 3]", None),
        ("Dedupe([]int{2,2,2})", "[2]", "A single run collapses to one element."),
        ("Dedupe(nil)", "[]", None),
    ],
    topics=[
        ("Two-cursor compaction", "The write cursor never passes the read cursor, so overwriting is safe."),
        ("Sorted input", "Duplicates are adjacent, so one comparison against the last kept element suffices."),
        ("Prefix results", "`s[:k]` is the answer without a second array."),
    ],
    hint="Compare each element with the last one you decided to keep, not with its neighbour.",
    intuition="Sorting has already grouped the duplicates. One pass with a write cursor rewrites the survivors to the front, and the prefix is the result.",
    approach=[
        "Return early for an empty slice.",
        "Start `k` at 1 — the first element always survives.",
        "For each later element, keep it when it differs from `s[k-1]`.",
        "Return `s[:k]`.",
    ],
    walkthrough="[1 1 2 3 3 3]: i=1 equals s[0], skipped. i=2 (2) differs, written to s[1], k=2. i=3 (3) differs from s[1], written to s[2], k=3. The rest match s[2]. Result s[:3] = [1 2 3].",
    pitfalls=[
        "Comparing `s[i]` with `s[i-1]` — correct only until the array has been overwritten behind you.",
        "Starting `k` at 0 and dropping the first element.",
    ],
)

P(
    name="poolscratch",
    title="Borrow A Buffer Instead Of Allocating One",
    sig="func Encode(vals []int) string",
    doc="""Encode returns vals rendered as decimal numbers joined by ','.

The scratch buffer used to build the text must come from the package's
sync.Pool and go back into it, so repeated calls do not each allocate a
buffer.

Examples:

	Encode([]int{1, 2}) => "1,2" """,
    imports=['"strconv"', '"sync"'],
    extra="""// pool hands out reusable scratch buffers.
var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}""",
    solution="""buf := pool.Get().([]byte)[:0]
for i, v := range vals {
	if i > 0 {
		buf = append(buf, ',')
	}
	buf = strconv.AppendInt(buf, int64(v), 10)
}
out := string(buf)
pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
return out""",
    tests="""
import "testing"

func TestEncode(t *testing.T) {
	if got := Encode([]int{1, 2, 3}); got != "1,2,3" {
		t.Errorf("Encode = %q, want \\"1,2,3\\"", got)
	}
	if got := Encode(nil); got != "" {
		t.Errorf("Encode = %q, want empty", got)
	}
	if got := Encode([]int{-7}); got != "-7" {
		t.Errorf("Encode = %q, want \\"-7\\"", got)
	}
}

func TestEncodeRepeatedCallsStayCorrect(t *testing.T) {
	for i := 0; i < 100; i++ {
		if got := Encode([]int{i, i + 1}); got != itoa(i)+","+itoa(i+1) {
			t.Fatalf("call %d: Encode = %q", i, got)
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func TestEncodeUsesThePool(t *testing.T) {
	Encode([]int{1})
	got := pool.Get()
	if got == nil {
		t.Fatal("the pool is empty: the scratch buffer was never returned")
	}
	if b, ok := got.([]byte); !ok || cap(b) == 0 {
		t.Errorf("pooled value = %T, want a []byte with capacity", got)
	}
}
""",
    context="A serialiser allocates a fresh scratch buffer for every record it renders. The buffers are identical in size and die immediately — textbook pool material.",
    task=[
        "Render `vals` as decimal numbers joined by `,`.",
        "Take the scratch buffer from `pool` and put it back before returning.",
        "Reset the borrowed buffer's length before writing into it.",
    ],
    examples=[
        ("Encode([]int{1,2,3})", '"1,2,3"', None),
        ("Encode(nil)", '""', None),
        ("Encode([]int{-7})", '"-7"', "The sign is part of the rendering."),
    ],
    topics=[
        ("sync.Pool", "A free list of reusable values; `Get` may return a recycled one or call `New`."),
        ("Resetting a borrowed buffer", "A pooled buffer arrives with whatever length it was returned at."),
        ("strconv.Append*", "Appends the rendering into an existing buffer instead of allocating a string."),
    ],
    hint="What is the length of a buffer that came back from the pool?",
    intuition="A pool turns \"allocate, use, drop\" into \"borrow, use, return\". The catch is that a borrowed buffer is not empty — it is whatever the last borrower left, so you must reset its length before you write.",
    approach=[
        "`pool.Get().([]byte)` and reslice to `[:0]`.",
        "Append the numbers and separators.",
        "Convert to a string, put the buffer back, return the string.",
    ],
    walkthrough="The first call finds the pool empty, so `New` makes a 64-byte buffer. Every later call reuses it: 100 calls, one buffer.",
    pitfalls=[
        "Returning `buf` itself instead of a string — the caller would hold memory the pool is handing to someone else.",
        "Forgetting `[:0]`, which appends onto the previous call's output.",
    ],
)

P(
    name="shrink",
    title="Give Back The Capacity You Stopped Using",
    sig="func Shrink(s []int) []int",
    doc="""Shrink returns a copy of s sized exactly to its length when s is
holding on to far more capacity than it uses, and returns s unchanged
otherwise.

"Far more" means the capacity is more than twice the length.

Examples:

	Shrink(make([]int, 2, 64)) => a slice of length 2 and capacity 2""",
    solution="""if cap(s) <= 2*len(s) {
	return s
}
out := make([]int, len(s))
copy(out, s)
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestShrinkReleasesTheSpare(t *testing.T) {
	s := make([]int, 2, 64)
	s[0], s[1] = 1, 2
	out := Shrink(s)
	if !reflect.DeepEqual(out, []int{1, 2}) {
		t.Errorf("contents = %v, want [1 2]", out)
	}
	if cap(out) != 2 {
		t.Errorf("cap = %d, want 2", cap(out))
	}
	if &out[0] == &s[0] {
		t.Error("the result still points at the oversized array")
	}
}

func TestShrinkKeepsATightSlice(t *testing.T) {
	s := make([]int, 8, 10)
	out := Shrink(s)
	if &out[0] != &s[0] {
		t.Error("a tight slice was copied for nothing")
	}
	if n := testing.AllocsPerRun(100, func() { _ = Shrink(s) }); n != 0 {
		t.Errorf("Shrink made %v allocations on a tight slice, want 0", n)
	}
}

func TestShrinkEmpty(t *testing.T) {
	if got := Shrink(nil); len(got) != 0 {
		t.Errorf("Shrink(nil) = %v, want empty", got)
	}
}
""",
    context="A long-lived index is built by appending millions of entries, then filtered down to a few thousand. The filtered index keeps the whole original array alive for the life of the process.",
    task=[
        "Return a right-sized copy when `cap(s) > 2*len(s)`.",
        "Return `s` untouched — and allocate nothing — otherwise.",
        "A nil or empty input must not panic.",
    ],
    examples=[
        ("Shrink(make([]int,2,64))", "len 2, cap 2, new array", None),
        ("Shrink(make([]int,8,10))", "the same slice", "10 is not more than twice 8, so nothing is copied."),
        ("Shrink(nil)", "[]", None),
    ],
    topics=[
        ("Capacity outlives length", "A short slice over a huge array pins the whole array."),
        ("Right-sizing", "`make([]int, len(s))` plus `copy` releases the spare on the next collection."),
        ("Copy only when it pays", "The threshold keeps the common case allocation-free."),
    ],
    hint="`cap(s) <= 2*len(s)` is the cheap case. Return early.",
    intuition="The collector frees allocations, not the unused part of one. Once a slice is a small window on a big array, the only way to release the rest is to copy the window somewhere its own size.",
    approach=[
        "Compare `cap(s)` with `2*len(s)`; if it is not bigger, return `s`.",
        "Otherwise allocate `make([]int, len(s))`, copy, return the copy.",
    ],
    walkthrough="A slice of len 2, cap 64 wastes 62 slots. Copying two ints into a two-int array lets the 64-int array be collected. A slice of len 8, cap 10 fails the test and is returned as is.",
    pitfalls=[
        "Shrinking unconditionally — every call then allocates.",
        "`s[:len(s):len(s)]` caps the capacity but keeps pointing at the same big array.",
    ],
)
