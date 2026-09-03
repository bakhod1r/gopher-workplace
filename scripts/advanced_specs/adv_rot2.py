"""10-advanced-topics — rotation 2: 5 puzzles each for middle, senior, staff.

Themes stay the four this topic opened with: memory management, escape
analysis, reflection and unsafe.
"""

SPECS = []


def P(level, **kw):
    kw["level"] = level
    kw.setdefault("sub", "10-advanced-topics")
    kw.setdefault("mode", "stub")
    kw.setdefault("stub", 'panic("not implemented")')
    kw.setdefault("imports", [])
    kw.setdefault("extra", "")
    SPECS.append(kw)


# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="shiftleft",
    title="Move Elements Over Themselves",
    sig="func Shift(s []int, n int) []int",
    doc="""Shift drops the first n elements of s by moving the rest to the front,
in place, and returns the shortened slice.

The source and destination ranges overlap, which copy handles correctly.
n is clamped into [0, len(s)].

Examples:

	Shift([]int{1, 2, 3, 4}, 2) => []int{3, 4}""",
    solution="""if n < 0 {
	n = 0
}
if n > len(s) {
	n = len(s)
}
k := copy(s, s[n:])
return s[:k]""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	if got := Shift([]int{1, 2, 3, 4}, 2); !reflect.DeepEqual(got, []int{3, 4}) {
		t.Errorf("Shift = %v, want [3 4]", got)
	}
	if got := Shift([]int{1, 2, 3}, 0); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Shift = %v, want [1 2 3]", got)
	}
	if got := Shift([]int{1, 2}, 9); len(got) != 0 {
		t.Errorf("Shift = %v, want empty", got)
	}
	if got := Shift([]int{1, 2}, -1); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("Shift = %v, want [1 2]", got)
	}
	if got := Shift(nil, 1); len(got) != 0 {
		t.Errorf("Shift = %v, want empty", got)
	}
}

func TestShiftOverlapIsHandled(t *testing.T) {
	s := make([]int, 1000)
	for i := range s {
		s[i] = i
	}
	got := Shift(s, 1)
	for i, v := range got {
		if v != i+1 {
			t.Fatalf("got[%d] = %d, want %d: the overlapping ranges were not copied correctly", i, v, i+1)
		}
	}
}

func TestShiftAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	if n := testing.AllocsPerRun(100, func() { _ = Shift(s, 1) }); n != 0 {
		t.Errorf("Shift made %v allocations, want 0", n)
	}
}
""",
    context="A queue drains its head by allocating a new slice for the tail. The tail is almost the whole queue, so the drain costs a full copy into fresh memory every time.",
    task=[
        "Drop the first `n` elements by moving the rest to the front of the same array.",
        "Clamp `n` into `[0, len(s)]`; return the shortened slice.",
        "Zero allocations, and correct when the ranges overlap.",
    ],
    examples=[
        ("Shift([]int{1,2,3,4}, 2)", "[3 4]", None),
        ("Shift([]int{1,2}, 9)", "[]", "n is clamped to the length."),
        ("Shift(s, 1) over 1000 elements", "every element moved down by one", None),
    ],
    topics=[
        ("copy handles overlap", "It behaves like memmove, so a left shift over itself is well defined."),
        ("copy returns the count", "The number of elements moved is the new length."),
        ("In-place over reallocation", "The array is already there; only the view changes."),
    ],
    hint="`copy(s, s[n:])` — and its return value is the answer's length.",
    intuition="A left shift moves every element to a lower index, so a forward copy never overwrites something it has not yet read. `copy` is specified to handle that, which is why no temporary is needed.",
    approach=[
        "Clamp `n`.",
        "`k := copy(s, s[n:])`.",
        "Return `s[:k]`.",
    ],
    walkthrough="For [1 2 3 4] and n = 2: `copy` moves 3 into index 0 and 4 into index 1, returning 2. The result is `s[:2]` = [3 4].",
    pitfalls=[
        "Copying into a temporary first — correct and pointless for a left shift.",
        "Returning `s[:len(s)-n]` without clamping, which is negative for a large `n`.",
    ],
)

P(
    "middle",
    name="boxvalues",
    title="The Detour Through Interface Values",
    mode="bug",
    sig="func Total(vals []int) int64",
    doc="""Total sums vals.

Passing the values through []any boxes every element: an interface value
needs a word to point at, so each int gets a heap home it never needed.

Examples:

	Total([]int{1, 2, 3}) => 6""",
    buggy="""boxed := make([]any, 0, len(vals))
for _, v := range vals {
	boxed = append(boxed, v)
}
var total int64
for _, b := range boxed {
	total += int64(b.(int))
}
return total""",
    solution="""var total int64
for _, v := range vals {
	total += int64(v)
}
return total""",
    tests="""
import "testing"

var sink int64

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2, 3}); got != 6 {
		t.Errorf("Total = %d, want 6", got)
	}
	if got := Total(nil); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
	if got := Total([]int{-5, 5}); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
}

func TestTotalWideAccumulator(t *testing.T) {
	vals := make([]int, 8)
	for i := range vals {
		vals[i] = 1 << 40
	}
	if got := Total(vals); got != 8<<40 {
		t.Errorf("Total = %d, want %d", got, int64(8)<<40)
	}
}

func TestTotalAllocatesNothing(t *testing.T) {
	vals := make([]int, 64)
	for i := range vals {
		vals[i] = 1000 + i
	}
	if n := testing.AllocsPerRun(100, func() { sink = Total(vals) }); n != 0 {
		t.Errorf("Total made %v allocations, want 0: the values are being boxed", n)
	}
}

func BenchmarkTotal(b *testing.B) {
	vals := make([]int, 4096)
	for i := range vals {
		vals[i] = i
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Total(vals)
	}
}
""",
    context="A summing helper was made \"generic\" by routing the values through []any so one function could serve several element types. The generality was never used, and the allocation profile pays for it on every call.",
    task=[
        "Sum `vals` and return the total as an int64.",
        "Fix the single bug so the function allocates nothing.",
        "The accumulator must stay wide enough that large values do not overflow.",
    ],
    examples=[
        ("Total([]int{1,2,3})", "6", None),
        ("Total(nil)", "0", None),
        ("64 values", "0 allocations, not 65", "One for the slice, one per boxed element."),
    ],
    topics=[
        ("Interface boxing", "Storing an int in an `any` needs a heap word for the data half of the interface."),
        ("The small-integer cache", "Values 0-255 reuse a runtime table, which is why the cost hides in toy tests."),
        ("Escape through a container", "The `[]any` outlives each element's scope, so every box escapes."),
        ("Generality has a price", "An unused abstraction still costs what it would have cost."),
    ],
    hint="The second loop already has everything it needs. What is the first loop for?",
    intuition="An interface value is a type word plus a data word, and the data word must be a pointer. Putting an int behind `any` therefore means allocating somewhere for the int to live -- once per element, for a generality nobody asked for.",
    approach=[
        "Delete the boxing pass.",
        "Range `vals` directly and accumulate into an int64.",
    ],
    walkthrough="For 64 values above 255, the boxed version allocates the `[]any` plus one box per element -- 65 allocations. The direct loop allocates nothing and keeps the accumulator in a register.",
    pitfalls=[
        "Testing with small values, where the runtime's 0-255 cache hides most of the cost.",
        "Accumulating in `int`, which is fine on 64-bit and not on every target.",
    ],
)

P(
    "middle",
    name="setallints",
    title="Write Every Int Field At Once",
    sig="func SetAllInts(ptr any, v int) (int, error)",
    doc="""SetAllInts sets every settable int field of the struct ptr points at to
v, and reports how many fields it wrote.

Unexported fields and fields of other kinds are skipped.

Examples:

	SetAllInts(&rec{}, 7) => 2, nil for a struct with two int fields""",
    imports=['"errors"', '"reflect"'],
    extra="""// ErrTarget is returned when ptr is not a non-nil pointer to a struct.
var ErrTarget = errors.New("target must be a non-nil pointer to a struct")""",
    solution="""rv := reflect.ValueOf(ptr)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return 0, ErrTarget
}
rv = rv.Elem()
if rv.Kind() != reflect.Struct {
	return 0, ErrTarget
}
n := 0
for i := 0; i < rv.NumField(); i++ {
	f := rv.Field(i)
	if f.Kind() == reflect.Int && f.CanSet() {
		f.SetInt(int64(v))
		n++
	}
}
return n, nil""",
    tests="""
import (
	"errors"
	"testing"
)

type rec struct {
	A      int
	B      int
	Name   string
	hidden int
	Ratio  float64
}

func TestSetAllInts(t *testing.T) {
	r := &rec{Name: "keep", hidden: 1}
	n, err := SetAllInts(r, 7)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Errorf("count = %d, want 2", n)
	}
	if r.A != 7 || r.B != 7 {
		t.Errorf("r = %+v, want A and B set to 7", *r)
	}
	if r.Name != "keep" || r.hidden != 1 || r.Ratio != 0 {
		t.Errorf("r = %+v: other fields must be untouched", *r)
	}
}

func TestSetAllIntsNoIntFields(t *testing.T) {
	var s struct {
		A string
		B bool
	}
	n, err := SetAllInts(&s, 1)
	if err != nil || n != 0 {
		t.Errorf("count = %d, err = %v, want 0, nil", n, err)
	}
}

func TestSetAllIntsBadTarget(t *testing.T) {
	for _, c := range []any{rec{}, nil, (*rec)(nil), new(int)} {
		if _, err := SetAllInts(c, 1); !errors.Is(err, ErrTarget) {
			t.Errorf("SetAllInts(%#v) = %v, want ErrTarget", c, err)
		}
	}
}

func TestSetAllIntsOverwrites(t *testing.T) {
	r := &rec{A: 1, B: 2}
	if _, err := SetAllInts(r, 0); err != nil {
		t.Fatal(err)
	}
	if r.A != 0 || r.B != 0 {
		t.Errorf("r = %+v, want both zeroed", *r)
	}
}
""",
    context="A test helper resets every counter in a stats struct before each case. Adding a counter means remembering to reset it, and nobody does.",
    task=[
        "Set every settable int field of the struct to `v`.",
        "Return how many fields were written.",
        "Skip unexported fields and every other kind.",
        "Return `ErrTarget` unless `ptr` is a non-nil pointer to a struct.",
    ],
    examples=[
        ("SetAllInts(&rec{}, 7)", "2, nil", "Only the two int fields."),
        ("a struct with no int fields", "0, nil", None),
        ("SetAllInts(rec{}, 1)", "ErrTarget", None),
    ],
    topics=[
        ("CanSet", "One check covers addressability and export status."),
        ("Kind filtering", "`SetInt` panics on the wrong kind, so filter before writing."),
        ("Counting the work", "Returning the count makes the helper testable without inspecting the struct."),
    ],
    hint="`CanSet` is the only export check you need on the Value side.",
    intuition="Once you have an addressable struct, each field is either writable or not, and reflection will tell you which. That turns \"reset all counters\" into a loop the struct definition drives.",
    approach=[
        "Validate the pointer and step to the struct.",
        "For each field, write it when the kind is int and `CanSet` is true.",
        "Return the count.",
    ],
    walkthrough="`rec` has five fields; `A` and `B` pass both checks, `Name` and `Ratio` fail the kind check, and `hidden` fails `CanSet`.",
    pitfalls=[
        "Checking `IsExported` but not the kind, which panics on the string field.",
        "Using `reflect.ValueOf(ptr).Field(i)` without `Elem`, which is not a struct at all.",
    ],
)

P(
    "middle",
    name="bytesequal",
    title="Compare Bytes To A String Without Converting",
    sig="func EqualString(b []byte, s string) bool",
    doc="""EqualString reports whether b's bytes are exactly s.

Neither side may be converted: a conversion in either direction copies
the payload just to throw the copy away.

Examples:

	EqualString([]byte("hi"), "hi") => true""",
    imports=['"unsafe"'],
    solution="""if len(b) != len(s) {
	return false
}
if len(b) == 0 {
	return true
}
return unsafe.String(unsafe.SliceData(b), len(b)) == s""",
    tests="""
import (
	"bytes"
	"testing"
)

var sink bool

func TestEqualString(t *testing.T) {
	cases := []struct {
		b, s string
		want bool
	}{
		{"hi", "hi", true},
		{"hi", "ho", false},
		{"hi", "his", false},
		{"his", "hi", false},
		{"", "", true},
		{"", "x", false},
		{"x", "", false},
	}
	for _, c := range cases {
		if got := EqualString([]byte(c.b), c.s); got != c.want {
			t.Errorf("EqualString(%q, %q) = %v, want %v", c.b, c.s, got, c.want)
		}
	}
}

func TestEqualStringNil(t *testing.T) {
	if !EqualString(nil, "") {
		t.Error("EqualString(nil, \\"\\") = false, want true")
	}
	if EqualString(nil, "x") {
		t.Error("EqualString(nil, \\"x\\") = true, want false")
	}
}

func TestEqualStringAllocatesNothing(t *testing.T) {
	b := bytes.Repeat([]byte("payload"), 512)
	s := string(b)
	if n := testing.AllocsPerRun(200, func() { sink = EqualString(b, s) }); n != 0 {
		t.Errorf("EqualString made %v allocations, want 0", n)
	}
}

func TestEqualStringDoesNotRetainTheView(t *testing.T) {
	b := []byte("mutable")
	if !EqualString(b, "mutable") {
		t.Fatal("EqualString = false, want true")
	}
	b[0] = 'M'
	if EqualString(b, "mutable") {
		t.Error("EqualString = true after the bytes changed, want false")
	}
}
""",
    context="A hot dispatcher compares an incoming frame's type field against a handful of known names. Every comparison copies the frame's bytes into a string first.",
    task=[
        "Report whether `b`'s bytes equal `s`.",
        "Convert nothing — zero allocations for any length.",
        "Compare the lengths before the contents.",
    ],
    examples=[
        ('EqualString([]byte("hi"), "hi")', "true", None),
        ('EqualString([]byte("his"), "hi")', "false", "Different lengths cannot match."),
        ('EqualString(nil, "")', "true", None),
    ],
    topics=[
        ("A borrowed view is safe here", "The string dies inside the call, so nothing can observe the aliasing."),
        ("Length first", "It is one comparison and it settles most cases."),
        ("bytes.Equal is the real answer", "This puzzle is about why it can be allocation-free."),
    ],
    hint="Wrap the bytes in a string that lives only long enough to be compared.",
    intuition="Zero-copy is a question of lifetime. A string view built purely to feed a comparison never escapes the function, so nothing can ever see that it aliased a mutable slice.",
    approach=[
        "Return false on a length mismatch, true for two empties.",
        "Build a string view over `b` and compare it with `s`.",
    ],
    walkthrough="For a 3584-byte payload, `string(b) == s` allocates and copies 3584 bytes. The view compares the same bytes in place and allocates nothing.",
    pitfalls=[
        "Returning or storing the view — the safety argument depends on it dying here.",
        "Skipping the empty guard; a nil data pointer with a non-zero length is invalid.",
    ],
)

P(
    "middle",
    name="groupby",
    title="Bucket Pairs Without Growing Every Bucket",
    sig="func Group(pairs [][2]int) map[int][]int",
    doc="""Group collects the second element of each pair into a bucket keyed by
the first, preserving input order within a bucket.

Both the map and each bucket should be sized from what the input already
tells you, instead of growing from nothing.

Examples:

	Group([][2]int{{1, 10}, {1, 11}, {2, 20}}) => map[1:[10 11] 2:[20]]""",
    solution="""if len(pairs) == 0 {
	return map[int][]int{}
}
counts := make(map[int]int, len(pairs))
for _, p := range pairs {
	counts[p[0]]++
}
out := make(map[int][]int, len(counts))
for k, n := range counts {
	out[k] = make([]int, 0, n)
}
for _, p := range pairs {
	out[p[0]] = append(out[p[0]], p[1])
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestGroup(t *testing.T) {
	got := Group([][2]int{{1, 10}, {2, 20}, {1, 11}})
	want := map[int][]int{1: {10, 11}, 2: {20}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Group = %v, want %v", got, want)
	}
}

func TestGroupEmpty(t *testing.T) {
	got := Group(nil)
	if got == nil {
		t.Fatal("Group(nil) = nil, want an empty map")
	}
	if len(got) != 0 {
		t.Errorf("Group = %v, want empty", got)
	}
}

func TestGroupPreservesOrder(t *testing.T) {
	pairs := make([][2]int, 0, 100)
	for i := 0; i < 100; i++ {
		pairs = append(pairs, [2]int{i % 3, i})
	}
	got := Group(pairs)
	for k, vs := range got {
		for i := 1; i < len(vs); i++ {
			if vs[i] <= vs[i-1] {
				t.Fatalf("bucket %d is out of order: %v", k, vs)
			}
		}
	}
}

func TestGroupBucketsAreRightSized(t *testing.T) {
	pairs := make([][2]int, 0, 300)
	for i := 0; i < 300; i++ {
		pairs = append(pairs, [2]int{i % 5, i})
	}
	got := Group(pairs)
	for k, vs := range got {
		if cap(vs) != len(vs) {
			t.Errorf("bucket %d has len %d and cap %d: size the buckets from the counts", k, len(vs), cap(vs))
		}
	}
}
""",
    context="A grouping step over a few million rows spends most of its time in `growslice`: every bucket doubles its way up from nothing, and there are thousands of buckets.",
    task=[
        "Collect each pair's second element into a bucket keyed by its first.",
        "Preserve input order within a bucket.",
        "Size the map and every bucket up front — each bucket's capacity must equal its final length.",
        "An empty input returns an empty, non-nil map.",
    ],
    examples=[
        ("Group([][2]int{{1,10},{2,20},{1,11}})", "map[1:[10 11] 2:[20]]", None),
        ("Group(nil)", "map[]", "Empty, not nil."),
        ("cap of each bucket", "equal to its length", None),
    ],
    topics=[
        ("Two-pass sizing", "Counting first turns every bucket's growth into one allocation."),
        ("Map size hints", "The distinct-key count is known after the counting pass."),
        ("Append order", "A single forward pass preserves the input order for free."),
    ],
    hint="Count first, allocate second, fill third.",
    intuition="Growth is only expensive when the size is a surprise. One cheap counting pass turns thousands of doubling chains into one exactly-sized allocation each.",
    approach=[
        "Count occurrences per key.",
        "Allocate the result map sized to the distinct-key count, and each bucket sized to its count.",
        "Walk the pairs again, appending into the pre-sized buckets.",
    ],
    walkthrough="With 300 pairs over 5 keys, the counting pass finds 60 per key; each bucket is allocated once at capacity 60 and never grows.",
    pitfalls=[
        "Returning nil for an empty input, which is a different value from an empty map.",
        "Sizing the buckets to `len(pairs)`, which is correct and wastes most of it.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="poolbound",
    title="The Pool That Kept Every Oversized Buffer",
    mode="bug",
    sig="func Render(size int) int",
    doc="""Render borrows a scratch buffer, fills size bytes of it, returns the
buffer to the pool and reports how many bytes it wrote.

Occasional huge requests must not leave the pool holding huge buffers
forever: a buffer larger than maxScratch is dropped instead of returned.

Examples:

	Render(16) => 16""",
    imports=['"sync"'],
    extra="""// maxScratch is the largest buffer worth keeping in the pool.
const maxScratch = 4096

var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}

// PooledCap reports the capacity of a buffer currently in the pool, or 0.
func PooledCap() int {
	v := pool.Get()
	if v == nil {
		return 0
	}
	b := v.([]byte)
	c := cap(b)
	pool.Put(b) //nolint:staticcheck // the puzzle keeps the pool API simple
	return c
}""",
    buggy="""if size < 0 {
	size = 0
}
buf := pool.Get().([]byte)[:0]
for i := 0; i < size; i++ {
	buf = append(buf, byte(i))
}
n := len(buf)
pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
return n""",
    solution="""if size < 0 {
	size = 0
}
buf := pool.Get().([]byte)[:0]
for i := 0; i < size; i++ {
	buf = append(buf, byte(i))
}
n := len(buf)
if cap(buf) <= maxScratch {
	pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
}
return n""",
    tests="""
import "testing"

func TestRender(t *testing.T) {
	if got := Render(16); got != 16 {
		t.Errorf("Render = %d, want 16", got)
	}
	if got := Render(0); got != 0 {
		t.Errorf("Render = %d, want 0", got)
	}
	if got := Render(-5); got != 0 {
		t.Errorf("Render = %d, want 0", got)
	}
}

func TestRenderLarge(t *testing.T) {
	if got := Render(1 << 20); got != 1<<20 {
		t.Errorf("Render = %d, want %d", got, 1<<20)
	}
}

func TestOversizedBuffersAreNotPooled(t *testing.T) {
	// drain whatever is in the pool
	for i := 0; i < 8; i++ {
		PooledCap()
	}
	Render(1 << 20)
	for i := 0; i < 8; i++ {
		if c := PooledCap(); c > maxScratch {
			t.Fatalf("the pool holds a %d-byte buffer, want at most %d: drop oversized buffers", c, maxScratch)
		}
	}
}

func TestSmallBuffersAreStillPooled(t *testing.T) {
	Render(32)
	if c := PooledCap(); c == 0 {
		t.Error("the pool is empty after a small render: normal buffers must go back")
	}
}

func TestRenderStaysCorrectAfterALargeRequest(t *testing.T) {
	Render(1 << 20)
	for i := 0; i < 50; i++ {
		if got := Render(8); got != 8 {
			t.Fatalf("call %d: Render = %d, want 8", i, got)
		}
	}
}
""",
    context="A service pools its render buffers and runs flat for weeks. One customer uploads a 40 MB document, and from then on the process holds 40 MB per pooled buffer forever.",
    task=[
        "Fill `size` bytes into a pooled buffer and return the count.",
        "Return the buffer to the pool only when its capacity is at most `maxScratch`.",
        "Fix the single bug; small buffers must still be recycled.",
    ],
    examples=[
        ("Render(16)", "16", None),
        ("Render(1<<20) then PooledCap()", "at most 4096", "The huge buffer was dropped."),
        ("Render(32) then PooledCap()", "non-zero", "Normal buffers still go back."),
    ],
    topics=[
        ("Pools have no size policy", "`sync.Pool` recycles whatever you hand it, however large."),
        ("Capacity is what persists", "The length resets; the allocation does not."),
        ("Dropping is free", "An unreturned buffer is simply collected."),
    ],
    hint="Everything about the borrow is right. What condition should guard the return?",
    intuition="A pool is a cache of allocations, and a cache with no eviction rule keeps its worst entry forever. One outlier request permanently raises the memory floor unless the return path checks the size.",
    approach=[
        "Fill the buffer as before.",
        "Return it to the pool only when `cap(buf) <= maxScratch`.",
    ],
    walkthrough="A 1 MiB render grows the borrowed buffer to at least 1 MiB. Putting it back leaves that megabyte pinned by the pool; dropping it lets the collector reclaim it and the next `Get` starts from 64 bytes again.",
    pitfalls=[
        "Checking `len(buf)` instead of `cap(buf)` — the length is zeroed on the next borrow.",
        "Dropping every buffer, which disables the pool entirely.",
    ],
)

P(
    "senior",
    name="closureretain",
    title="The Callback That Held The Whole Batch",
    mode="bug",
    sig="func Summarize(batch []Record) func() int",
    doc="""Summarize returns a function reporting the batch's total size.

The returned function outlives the batch, so it must capture the answer
rather than the data: a closure over the slice keeps every record alive
for as long as the callback exists.

Examples:

	f := Summarize(batch); f() => the total size""",
    extra="""// Record is one ingested item.
type Record struct {
	Size int
	Pad  [256]byte
}""",
    buggy="""return func() int {
	total := 0
	for _, r := range batch {
		total += r.Size
	}
	return total
}""",
    solution="""total := 0
for _, r := range batch {
	total += r.Size
}
return func() int { return total }""",
    tests="""
import (
	"runtime"
	"testing"
)

func TestSummarize(t *testing.T) {
	batch := []Record{{Size: 1}, {Size: 2}, {Size: 3}}
	f := Summarize(batch)
	if got := f(); got != 6 {
		t.Errorf("f() = %d, want 6", got)
	}
	if got := f(); got != 6 {
		t.Errorf("second call = %d, want 6", got)
	}
}

func TestSummarizeEmpty(t *testing.T) {
	if got := Summarize(nil)(); got != 0 {
		t.Errorf("f() = %d, want 0", got)
	}
}

func TestSummarizeSnapshotsTheBatch(t *testing.T) {
	batch := []Record{{Size: 1}, {Size: 2}}
	f := Summarize(batch)
	batch[0].Size = 100
	if got := f(); got != 3 {
		t.Errorf("f() = %d, want 3: the total must be computed before the callback is returned", got)
	}
}

func TestSummarizeReleasesTheBatch(t *testing.T) {
	makeCallback := func() func() int {
		batch := make([]Record, 8192)
		for i := range batch {
			batch[i].Size = 1
		}
		return Summarize(batch)
	}

	runtime.GC()
	var before runtime.MemStats
	runtime.ReadMemStats(&before)

	f := makeCallback()
	if got := f(); got != 8192 {
		t.Fatalf("f() = %d, want 8192", got)
	}

	runtime.GC()
	var after runtime.MemStats
	runtime.ReadMemStats(&after)
	runtime.KeepAlive(f)

	if after.HeapAlloc > before.HeapAlloc+(1<<20) {
		t.Errorf("heap grew by %d bytes with the callback alive, want under 1 MiB: the closure retains the batch",
			after.HeapAlloc-before.HeapAlloc)
	}
}
""",
    context="A pipeline stage returns a small \"report this later\" callback per batch. The callbacks are tiny, there are thousands of them, and the heap holds every batch that ever passed through.",
    task=[
        "Return a function reporting the batch's total `Size`.",
        "The total must be fixed when `Summarize` returns — later edits to the batch are invisible.",
        "Fix the single bug so the callback does not keep the batch alive.",
    ],
    examples=[
        ("f := Summarize(batch); f()", "the total", None),
        ("batch[0].Size = 100 after Summarize", "f() unchanged", "The total is a snapshot."),
        ("heap with the callback alive", "the batch is collectable", None),
    ],
    topics=[
        ("Closures capture variables, not values", "Capturing the slice keeps its backing array reachable."),
        ("Retention through a callback", "A one-word result held by a closure over megabytes still costs megabytes."),
        ("Compute eagerly to release early", "Capturing the answer breaks the reference to the data."),
    ],
    hint="The loop is in the wrong place. Where does the answer have to be computed for the batch to be droppable?",
    intuition="Whatever a closure mentions, it keeps. Deferring the computation looks lazy and cheap, but it converts a temporary batch into state with the callback's lifetime.",
    approach=[
        "Compute the total before constructing the closure.",
        "Return a closure that captures only the total.",
    ],
    walkthrough="8192 records of about 264 bytes is roughly 2 MiB. The lazy closure pins all of it; the eager version captures one int and the batch is collected as soon as `Summarize` returns.",
    pitfalls=[
        "Copying the batch into the closure — same retention, plus a copy.",
        "Assuming laziness is free; it moves work later and lifetime longer.",
    ],
)

P(
    "senior",
    name="tagvalidate",
    title="Reject A Bad Schema Before It Ships",
    sig="func Validate(v any) []string",
    doc="""Validate returns the problems with v's `col` tags, in field order.

Every exported field must carry a non-empty col tag, no two fields may
share a tag, and the tag must contain no comma. Each problem is reported
as "FieldName: reason".

Examples:

	Validate(bad{}) => []string{"B: duplicate tag \\"a\\""}""",
    imports=['"reflect"'],
    solution="""rt := reflect.TypeOf(v)
if rt == nil || rt.Kind() != reflect.Struct {
	return []string{"not a struct"}
}
var out []string
seen := make(map[string]string, rt.NumField())
for i := 0; i < rt.NumField(); i++ {
	f := rt.Field(i)
	if !f.IsExported() {
		continue
	}
	tag, ok := f.Tag.Lookup("col")
	if !ok || tag == "" {
		out = append(out, f.Name+": missing col tag")
		continue
	}
	bad := false
	for j := 0; j < len(tag); j++ {
		if tag[j] == ',' {
			bad = true
			break
		}
	}
	if bad {
		out = append(out, f.Name+": tag contains a comma")
		continue
	}
	if prev, dup := seen[tag]; dup {
		out = append(out, f.Name+": duplicate tag of "+prev)
		continue
	}
	seen[tag] = f.Name
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type good struct {
	A      int    `col:"a"`
	B      string `col:"b"`
	hidden int
}

type missing struct {
	A int `col:"a"`
	B int
}

type empty struct {
	A int `col:""`
}

type dup struct {
	A int `col:"x"`
	B int `col:"x"`
}

type comma struct {
	A int `col:"a,omitempty"`
}

func TestValidateGood(t *testing.T) {
	if got := Validate(good{}); len(got) != 0 {
		t.Errorf("Validate = %v, want no problems", got)
	}
}

func TestValidateMissing(t *testing.T) {
	if got := Validate(missing{}); !reflect.DeepEqual(got, []string{"B: missing col tag"}) {
		t.Errorf("Validate = %v, want [B: missing col tag]", got)
	}
	if got := Validate(empty{}); !reflect.DeepEqual(got, []string{"A: missing col tag"}) {
		t.Errorf("Validate = %v, want [A: missing col tag]", got)
	}
}

func TestValidateDuplicate(t *testing.T) {
	if got := Validate(dup{}); !reflect.DeepEqual(got, []string{"B: duplicate tag of A"}) {
		t.Errorf("Validate = %v, want [B: duplicate tag of A]", got)
	}
}

func TestValidateComma(t *testing.T) {
	if got := Validate(comma{}); !reflect.DeepEqual(got, []string{"A: tag contains a comma"}) {
		t.Errorf("Validate = %v, want [A: tag contains a comma]", got)
	}
}

func TestValidateNonStruct(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}} {
		if got := Validate(in); len(got) != 1 || got[0] != "not a struct" {
			t.Errorf("Validate(%#v) = %v, want [not a struct]", in, got)
		}
	}
}

func TestValidateReportsInFieldOrder(t *testing.T) {
	type multi struct {
		A int `col:"x"`
		B int
		C int `col:"x"`
	}
	want := []string{"B: missing col tag", "C: duplicate tag of A"}
	if got := Validate(multi{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Validate = %v, want %v", got, want)
	}
}
""",
    context="A tag typo maps two struct fields to one database column. The mistake is invisible in review and shows up as data loss in production.",
    task=[
        "Report every problem with `v`'s `col` tags, in field order.",
        "A missing or empty tag, a comma in the tag, and a duplicate tag are each a problem.",
        "Skip unexported fields; report `\"not a struct\"` for anything that is not a struct.",
        "One problem per field: report the first one found and move on.",
    ],
    examples=[
        ("Validate(good{})", "[]", "No problems."),
        ("Validate(missing{})", "[B: missing col tag]", None),
        ("Validate(dup{})", "[B: duplicate tag of A]", "The later field is the one reported."),
    ],
    topics=[
        ("Schema validation at run time", "The struct is the schema, so it can check itself."),
        ("Tag.Lookup", "Distinguishes an absent tag from an empty one — both are problems here."),
        ("Deterministic reporting", "Field order makes the output stable and diffable."),
        ("First problem per field", "Continuing after a report keeps each field to one line."),
    ],
    hint="One pass, a `seen` map from tag to the field that claimed it.",
    intuition="A tag-driven mapping is a schema the compiler does not check. Walking the type once at start-up turns a silent production bug into a startup failure.",
    approach=[
        "Reject non-structs.",
        "For each exported field: look the tag up, then check empty, comma and duplicate in that order.",
        "Record the tag's owner in `seen` only when the field is otherwise valid.",
    ],
    walkthrough="For `multi`: A claims \"x\"; B has no tag and is reported; C's \"x\" is already claimed by A, so it is reported as a duplicate of A.",
    pitfalls=[
        "Recording a tag in `seen` before validating it, so a later field duplicates an invalid tag.",
        "Using `Tag.Get`, which cannot tell an empty tag from a missing one.",
    ],
)

P(
    "senior",
    name="alignassert",
    title="Check The Alignment You Depend On",
    mode="bug",
    sig="func Check() bool",
    doc="""Check reports whether Counter's Value field is aligned well enough for
64-bit atomic operations.

The requirement is the type's own alignment, which unsafe.Alignof
reports. Hard-coding a number is how this check passes on the machine it
was written on and nowhere else.

Examples:

	Check() => true for a correctly laid out Counter""",
    imports=['"unsafe"'],
    extra="""// Counter is a struct whose Value field is updated atomically.
type Counter struct {
	Value int64
	Name  string
}""",
    buggy="""var c Counter
return uintptr(unsafe.Pointer(&c.Value))%8 == 0 && unsafe.Offsetof(c.Value) == 0""",
    solution="""var c Counter
want := unsafe.Alignof(c.Value)
return unsafe.Offsetof(c.Value)%want == 0 &&
	uintptr(unsafe.Pointer(&c.Value))%want == 0""",
    tests="""
import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
	"unsafe"
)

func TestCheck(t *testing.T) {
	if !Check() {
		t.Error("Check = false, want true: Value is at offset 0 of a 64-bit-aligned struct")
	}
}

func TestCheckIsRepeatable(t *testing.T) {
	for i := 0; i < 100; i++ {
		if !Check() {
			t.Fatalf("run %d: Check = false, want true", i)
		}
	}
}

func TestFixtureIsUnchanged(t *testing.T) {
	var c Counter
	if unsafe.Offsetof(c.Value) != 0 {
		t.Error("Value must stay the first field")
	}
	if unsafe.Sizeof(c.Value) != 8 {
		t.Error("Value must stay an int64")
	}
}

func TestCheckDerivesTheRequirement(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "alignassert.go", nil, 0)
	if err != nil {
		t.Skipf("cannot parse the source: %v", err)
	}
	usesAlignof := false
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Name == "unsafe" && sel.Sel.Name == "Alignof" {
				usesAlignof = true
			}
		}
		return true
	})
	if !usesAlignof {
		t.Error("the requirement must come from unsafe.Alignof, not from a hard-coded number")
	}
}
""",
    context="A struct is laid out so its counter can be updated atomically. The invariant is asserted with a literal 8, and on a 32-bit build the assertion is both wrong and green.",
    task=[
        "Report whether `Counter.Value` is aligned for atomic access.",
        "Derive the requirement from `unsafe.Alignof`, not from a literal.",
        "Check the field's offset within the struct and the address of an actual instance.",
        "Fix the single bug.",
    ],
    examples=[
        ("Check()", "true", "Value is the first field of a well-aligned struct."),
        ("100 runs", "true every time", None),
        ("the requirement's source", "unsafe.Alignof", None),
    ],
    topics=[
        ("unsafe.Alignof", "The type's required alignment, as a compile-time constant."),
        ("Alignment is per platform", "A 64-bit value is not 8-byte aligned on every architecture by default."),
        ("Offset and address", "Both matter: a well-placed field in a misaligned struct is still misaligned."),
    ],
    hint="The number 8 is an answer, not a question. Which call asks the question?",
    intuition="Hard-coded layout constants encode one machine's answer into code that will run on others. `Alignof` asks the compiler for the real requirement, so the check travels with the build.",
    approach=[
        "Take `want := unsafe.Alignof(c.Value)`.",
        "Check that the field's offset is a multiple of `want`.",
        "Check that an instance's field address is a multiple of `want`.",
    ],
    walkthrough="On a 64-bit build `Alignof(int64)` is 8 and both checks pass. On a platform where it is 4, the literal 8 would have demanded more than the platform provides — the assertion would fail for the wrong reason.",
    pitfalls=[
        "Checking only the offset; the struct itself may sit at a misaligned address inside another struct.",
        "Using `Sizeof` in place of `Alignof` — they agree for int64 and not in general.",
    ],
)

P(
    "senior",
    name="builderreuse",
    title="One Builder, Many Lines",
    sig="func RenderLines(rows [][]int) []string",
    doc="""RenderLines renders each row as its values joined by '-'.

The builder is per-call state: reset it between rows instead of
constructing one per row, and reserve its capacity once.

Examples:

	RenderLines([][]int{{1, 2}}) => []string{"1-2"}""",
    imports=['"strconv"', '"strings"'],
    solution="""out := make([]string, 0, len(rows))
widest := 0
for _, row := range rows {
	if n := len(row) * 12; n > widest {
		widest = n
	}
}
var b strings.Builder
b.Grow(widest)
for _, row := range rows {
	b.Reset()
	for i, v := range row {
		if i > 0 {
			b.WriteByte('-')
		}
		b.WriteString(strconv.Itoa(v))
	}
	out = append(out, b.String())
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestRenderLines(t *testing.T) {
	got := RenderLines([][]int{{1, 2}, {3}, {}})
	want := []string{"1-2", "3", ""}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RenderLines = %q, want %q", got, want)
	}
	if got := RenderLines(nil); len(got) != 0 {
		t.Errorf("RenderLines = %q, want empty", got)
	}
}

func TestRenderLinesNegatives(t *testing.T) {
	got := RenderLines([][]int{{-1, 2, -3}})
	if len(got) != 1 || got[0] != "-1-2--3" {
		t.Errorf("RenderLines = %q, want [-1-2--3]", got)
	}
}

func TestRenderLinesDoesNotLeakBetweenRows(t *testing.T) {
	rows := make([][]int, 50)
	for i := range rows {
		rows[i] = []int{i}
	}
	got := RenderLines(rows)
	for i, s := range got {
		want := strconvItoa(i)
		if s != want {
			t.Fatalf("line %d = %q, want %q: the builder was not reset", i, s, want)
		}
	}
}

func strconvItoa(n int) string {
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

func TestRenderLinesAllocationsScaleWithRows(t *testing.T) {
	rows := make([][]int, 64)
	for i := range rows {
		rows[i] = []int{i, i + 1, i + 2}
	}
	n := testing.AllocsPerRun(50, func() { _ = RenderLines(rows) })
	if n > float64(len(rows))*2+8 {
		t.Errorf("RenderLines made %v allocations for %d rows, want about %d: reuse one builder",
			n, len(rows), len(rows)*2)
	}
}
""",
    context="A CSV writer constructs a `strings.Builder` per row. Each one starts empty, grows through several sizes, and is discarded a line later.",
    task=[
        "Render each row as its values joined by `-`.",
        "Use one builder for the whole call, reset between rows.",
        "Reserve its capacity once, before the loop.",
        "An empty row renders as the empty string.",
    ],
    examples=[
        ("RenderLines([][]int{{1,2},{3}})", '["1-2" "3"]', None),
        ("RenderLines([][]int{{}})", '[""]', None),
        ("RenderLines([][]int{{-1,2,-3}})", '["-1-2--3"]', "The separator and the minus sign both appear."),
    ],
    topics=[
        ("Builder.Reset", "Empties the builder while keeping the buffer it has grown."),
        ("Grow once", "Reserving the widest row's worth removes every intermediate growth."),
        ("String extraction copies", "`b.String()` is the per-row allocation you cannot avoid."),
    ],
    hint="Construct and `Grow` above the loop; `Reset` inside it.",
    intuition="A builder's value is the buffer it accumulates. Constructing one per row throws that buffer away every time — the reuse is the entire point of the type.",
    approach=[
        "Preallocate the result slice.",
        "Estimate the widest row and `Grow` the builder once.",
        "Per row: `Reset`, write the values with separators, append `b.String()`.",
    ],
    walkthrough="For 64 rows of three values, one builder is grown once and reset 64 times. Only the 64 result strings and the result slice are allocated.",
    pitfalls=[
        "Forgetting `Reset`, which concatenates every row onto the previous ones.",
        "Reusing a builder across calls — `strings.Builder` must not be copied, and sharing one is not concurrency-safe.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="boundedcache",
    title="A Cache That Cannot Outgrow Its Limit",
    sig="func (c *Cache) Put(key string, val []byte)",
    doc="""Put stores a copy of val under key, evicting the oldest entry when the
cache is at capacity.

The stored value must own its bytes — callers reuse their buffers — and
the cache must never hold more than limit entries.

Examples:

	c := NewCache(2); c.Put("a", v) => Get("a") returns a copy of v""",
    imports=['"sync"'],
    extra="""// Cache is a bounded, concurrency-safe byte cache with FIFO eviction.
type Cache struct {
	mu    sync.Mutex
	limit int
	items map[string][]byte
	order []string
}

// NewCache returns a cache holding at most limit entries.
func NewCache(limit int) *Cache {
	if limit < 1 {
		limit = 1
	}
	return &Cache{limit: limit, items: make(map[string][]byte, limit), order: make([]string, 0, limit)}
}

// Get returns the stored bytes for key, if present.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.items[key]
	return v, ok
}

// Len reports how many entries the cache holds.
func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}""",
    solution="""owned := make([]byte, len(val))
copy(owned, val)

c.mu.Lock()
defer c.mu.Unlock()

if _, exists := c.items[key]; !exists {
	if len(c.order) >= c.limit {
		oldest := c.order[0]
		c.order = append(c.order[:0], c.order[1:]...)
		delete(c.items, oldest)
	}
	c.order = append(c.order, key)
}
c.items[key] = owned""",
    tests="""
import (
	"bytes"
	"sync"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	c := NewCache(4)
	c.Put("a", []byte("one"))
	got, ok := c.Get("a")
	if !ok || !bytes.Equal(got, []byte("one")) {
		t.Errorf("Get = %q, %v, want \\"one\\", true", got, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Error("Get(missing) reported ok, want false")
	}
}

func TestPutCopiesTheValue(t *testing.T) {
	c := NewCache(4)
	buf := []byte("first")
	c.Put("k", buf)
	copy(buf, "SECON")
	got, _ := c.Get("k")
	if !bytes.Equal(got, []byte("first")) {
		t.Errorf("Get = %q, want \\"first\\": the cache stored the caller's buffer", got)
	}
}

func TestEvictsOldest(t *testing.T) {
	c := NewCache(2)
	c.Put("a", []byte("1"))
	c.Put("b", []byte("2"))
	c.Put("c", []byte("3"))
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
	if _, ok := c.Get("a"); ok {
		t.Error("the oldest entry was not evicted")
	}
	for _, k := range []string{"b", "c"} {
		if _, ok := c.Get(k); !ok {
			t.Errorf("%q was evicted, want it kept", k)
		}
	}
}

func TestOverwriteDoesNotEvict(t *testing.T) {
	c := NewCache(2)
	c.Put("a", []byte("1"))
	c.Put("b", []byte("2"))
	c.Put("a", []byte("updated"))
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
	got, ok := c.Get("a")
	if !ok || !bytes.Equal(got, []byte("updated")) {
		t.Errorf("Get(a) = %q, %v, want \\"updated\\", true", got, ok)
	}
	if _, ok := c.Get("b"); !ok {
		t.Error("b was evicted by an overwrite")
	}
}

func TestStaysBoundedUnderLoad(t *testing.T) {
	c := NewCache(8)
	for i := 0; i < 5000; i++ {
		c.Put(string(rune('a'+i%26))+string(rune('a'+i/26%26)), []byte("payload"))
		if c.Len() > 8 {
			t.Fatalf("Len = %d after %d puts, want at most 8", c.Len(), i+1)
		}
	}
}

func TestConcurrentPuts(t *testing.T) {
	c := NewCache(16)
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			buf := make([]byte, 8)
			for i := 0; i < 500; i++ {
				for j := range buf {
					buf[j] = byte('a' + w)
				}
				c.Put(string(rune('a'+w))+string(rune('0'+i%10)), buf)
				c.Get(string(rune('a' + w)))
			}
		}(w)
	}
	wg.Wait()
	if c.Len() > 16 {
		t.Errorf("Len = %d, want at most 16", c.Len())
	}
}
""",
    context="A response cache is added with a size limit and a mutex. Memory still climbs: the limit counts entries in the map, and the eviction path was never reached because every key was new.",
    task=[
        "Store a copy of `val` under `key` — the caller reuses its buffer.",
        "Evict the oldest entry when inserting a new key would exceed `limit`.",
        "Overwriting an existing key must not evict anything.",
        "Safe for concurrent use; copy outside the lock.",
    ],
    examples=[
        ('c.Put("k", buf); copy(buf, "SECON")', "Get(\"k\") still returns the original", None),
        ("NewCache(2), put a, b, c", "a is evicted, Len is 2", None),
        ("overwriting an existing key", "no eviction", None),
    ],
    topics=[
        ("Ownership at the boundary", "A cache that stores a caller's slice stores a promise the caller will break."),
        ("Bounded state", "The eviction path must run on the insert of a new key, not on every put."),
        ("Lock scope", "Allocating and copying outside the lock keeps the critical section short."),
        ("FIFO order tracking", "The order slice is what makes \"oldest\" meaningful."),
    ],
    hint="Two decisions per put: does this key already exist, and is the cache full?",
    intuition="Bounded memory needs two things a map does not give you: a rule for what to drop and ownership of what you keep. Both are decided at insert time, and only for keys that are actually new.",
    approach=[
        "Copy `val` into a private slice before taking the lock.",
        "Under the lock, check whether the key exists.",
        "For a new key at capacity, drop the front of the order slice and delete it from the map.",
        "Record the new key in the order and store the copy.",
    ],
    walkthrough="With a limit of 2 and puts a, b, c: c is new and the cache is full, so a is evicted. Putting a again after that would evict b — but putting b again would only overwrite.",
    pitfalls=[
        "Appending to `order` on an overwrite, which lets the same key sit in the order twice.",
        "Storing `val` directly, which is invisible until the caller reuses the buffer.",
        "Holding the lock across the copy, which serialises every writer on a memcpy.",
    ],
)

P(
    "staff",
    name="hotaggregate",
    title="Aggregate A Stream With No Garbage At All",
    sig="func Aggregate(lines [][]byte, sep byte) (int64, int, error)",
    doc="""Aggregate sums the decimal integers across every line and reports the
total and the field count.

The whole aggregation must run without allocating: no conversions, no
split slices, no formatted errors.

Examples:

	Aggregate([][]byte{[]byte("1,2")}, ',') => 3, 2, nil""",
    imports=['"errors"'],
    extra="""// ErrSyntax marks a field that is not a decimal integer.
var ErrSyntax = errors.New("invalid integer")""",
    solution="""var total int64
count := 0
for _, line := range lines {
	if len(line) == 0 {
		continue
	}
	start := 0
	for i := 0; i <= len(line); i++ {
		if i < len(line) && line[i] != sep {
			continue
		}
		field := line[start:i]
		start = i + 1
		if len(field) == 0 {
			return 0, 0, ErrSyntax
		}
		neg := false
		j := 0
		if field[0] == '-' || field[0] == '+' {
			neg = field[0] == '-'
			j = 1
		}
		if j == len(field) {
			return 0, 0, ErrSyntax
		}
		var v int64
		for ; j < len(field); j++ {
			c := field[j]
			if c < '0' || c > '9' {
				return 0, 0, ErrSyntax
			}
			v = v*10 + int64(c-'0')
		}
		if neg {
			v = -v
		}
		total += v
		count++
	}
}
return total, count, nil""",
    tests="""
import (
	"bytes"
	"errors"
	"testing"
)

var (
	sinkT int64
	sinkC int
)

func TestAggregate(t *testing.T) {
	in := [][]byte{[]byte("1,2"), []byte("3"), []byte("-4,+5")}
	total, count, err := Aggregate(in, ',')
	if err != nil || total != 7 || count != 5 {
		t.Errorf("Aggregate = %d, %d, %v, want 7, 5, nil", total, count, err)
	}
}

func TestAggregateEmptyInputs(t *testing.T) {
	if total, count, err := Aggregate(nil, ','); err != nil || total != 0 || count != 0 {
		t.Errorf("Aggregate = %d, %d, %v, want 0, 0, nil", total, count, err)
	}
	if total, count, err := Aggregate([][]byte{nil, {}}, ','); err != nil || total != 0 || count != 0 {
		t.Errorf("Aggregate = %d, %d, %v, want 0, 0, nil", total, count, err)
	}
}

func TestAggregateSyntaxErrors(t *testing.T) {
	for _, in := range []string{"1,,2", "1,x", "-", "+", "a"} {
		if _, _, err := Aggregate([][]byte{[]byte(in)}, ','); !errors.Is(err, ErrSyntax) {
			t.Errorf("Aggregate(%q) = %v, want ErrSyntax", in, err)
		}
	}
}

func TestAggregateAllocatesNothing(t *testing.T) {
	lines := make([][]byte, 64)
	for i := range lines {
		lines[i] = bytes.Repeat([]byte("12345,"), 16)
		lines[i] = lines[i][:len(lines[i])-1]
	}
	n := testing.AllocsPerRun(50, func() { sinkT, sinkC, _ = Aggregate(lines, ',') })
	if n != 0 {
		t.Errorf("Aggregate made %v allocations, want 0", n)
	}
}

func TestAggregateErrorPathAllocatesNothing(t *testing.T) {
	lines := [][]byte{[]byte("1,x")}
	var err error
	n := testing.AllocsPerRun(50, func() { _, _, err = Aggregate(lines, ',') })
	_ = err
	if n != 0 {
		t.Errorf("the error path made %v allocations, want 0: return the sentinel", n)
	}
}

func BenchmarkAggregate(b *testing.B) {
	lines := make([][]byte, 256)
	for i := range lines {
		lines[i] = bytes.Repeat([]byte("12345,"), 16)
		lines[i] = lines[i][:len(lines[i])-1]
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkT, sinkC, _ = Aggregate(lines, ',')
	}
}
""",
    context="An ingest hot path parses a million lines a second. Every allocation it makes is one the collector has to chase, and the profile is all allocator.",
    task=[
        "Sum the decimal integers across every line, separated by `sep`.",
        "Return the total, the field count, and `ErrSyntax` for a malformed field.",
        "Empty lines are skipped; an empty field is a syntax error.",
        "Zero allocations — on the success path and the error path alike.",
    ],
    examples=[
        ('Aggregate([][]byte{[]byte("1,2"), []byte("3")}, \',\')', "6, 3, nil", None),
        ('Aggregate([][]byte{[]byte("1,,2")}, \',\')', "ErrSyntax", None),
        ("64 lines of 16 fields", "0 allocations", None),
    ],
    topics=[
        ("Parse in place", "Digits fold into an accumulator straight from the bytes."),
        ("Sentinel errors on the failure path", "A formatted error would allocate exactly when the input is hostile."),
        ("Virtual trailing separator", "Running the index to `len(line)` inclusive closes the final field."),
        ("Allocation as a graded property", "`AllocsPerRun` on both paths is the specification."),
    ],
    hint="Two nested loops and one accumulator. Nothing is ever built.",
    intuition="Every convenience in a parser — split, convert, format — is an allocation. At a million lines a second the only viable design is one that reads the caller's bytes and produces nothing but numbers.",
    approach=[
        "For each non-empty line, walk `i` to `len(line)` inclusive, treating the end as a separator.",
        "Per field: optional sign, then fold digits; reject empties, lone signs and non-digits.",
        "Accumulate the total and the count.",
    ],
    walkthrough='For "-4,+5": the first field folds to 4 and is negated; the second folds to 5. The total gains 1 and the count gains 2, with no memory touched beyond the caller\'s slices.',
    pitfalls=[
        "Returning `fmt.Errorf` for a bad field, which allocates on precisely the input an attacker controls.",
        "Stopping the inner loop at `len(line)-1`, which drops the last field.",
        "Treating an empty line as a syntax error; the spec skips it.",
    ],
)

P(
    "staff",
    name="shardedcounter",
    title="Shard The Counter, Fold It Once",
    sig="func (c *Counter) Add(key string, n int64)",
    doc="""Add increments the counter for key.

Counters are sharded by the key's hash so concurrent writers rarely touch
the same lock; Total folds the shards after the writers are done.

Examples:

	c := NewCounter(4); c.Add("a", 1) => Total()["a"] == 1""",
    imports=['"hash/maphash"', '"sync"'],
    extra="""// shard is one lock-protected slice of the key space, padded so two
// shards never share a cache line.
type shard struct {
	mu sync.Mutex
	m  map[string]int64
	_  [48]byte
}

// Counter is a sharded string counter.
type Counter struct {
	seed   maphash.Seed
	shards []shard
}

// NewCounter returns a counter with n shards (rounded up to at least 1).
func NewCounter(n int) *Counter {
	if n < 1 {
		n = 1
	}
	c := &Counter{seed: maphash.MakeSeed(), shards: make([]shard, n)}
	for i := range c.shards {
		c.shards[i].m = make(map[string]int64)
	}
	return c
}

// shardFor returns the shard that owns key.
func (c *Counter) shardFor(key string) *shard {
	h := maphash.String(c.seed, key)
	return &c.shards[h%uint64(len(c.shards))]
}

// Total folds every shard into one map. Call it after the writers are done.
func (c *Counter) Total() map[string]int64 {
	out := make(map[string]int64)
	for i := range c.shards {
		s := &c.shards[i]
		s.mu.Lock()
		for k, v := range s.m {
			out[k] += v
		}
		s.mu.Unlock()
	}
	return out
}""",
    solution="""s := c.shardFor(key)
s.mu.Lock()
s.m[key] += n
s.mu.Unlock()""",
    tests="""
import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestAddAndTotal(t *testing.T) {
	c := NewCounter(4)
	c.Add("a", 1)
	c.Add("a", 2)
	c.Add("b", 5)
	got := c.Total()
	if got["a"] != 3 || got["b"] != 5 {
		t.Errorf("Total = %v, want map[a:3 b:5]", got)
	}
}

func TestAddNegativeAndZero(t *testing.T) {
	c := NewCounter(2)
	c.Add("k", 5)
	c.Add("k", -5)
	if got := c.Total(); got["k"] != 0 {
		t.Errorf("Total = %v, want map[k:0]", got)
	}
}

func TestSingleShard(t *testing.T) {
	c := NewCounter(0)
	c.Add("a", 1)
	c.Add("b", 2)
	got := c.Total()
	if got["a"] != 1 || got["b"] != 2 {
		t.Errorf("Total = %v, want map[a:1 b:2]", got)
	}
}

func TestConcurrentAdds(t *testing.T) {
	const (
		workers = 16
		perTask = 1000
		keys    = 8
	)
	c := NewCounter(16)
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < perTask; i++ {
				c.Add(fmt.Sprintf("k%d", i%keys), 1)
			}
		}()
	}
	wg.Wait()
	got := c.Total()
	if len(got) != keys {
		t.Fatalf("Total has %d keys, want %d", len(got), keys)
	}
	want := int64(workers * perTask / keys)
	for k, v := range got {
		if v != want {
			t.Fatalf("Total[%q] = %d, want %d: increments were lost", k, v, want)
		}
	}
}

func TestSameKeyAlwaysHitsOneShard(t *testing.T) {
	c := NewCounter(8)
	first := c.shardFor("stable")
	for i := 0; i < 100; i++ {
		if c.shardFor("stable") != first {
			t.Fatal("shardFor is not deterministic for one counter")
		}
	}
}

func TestShardsDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(shard{}); got < 64 {
		t.Errorf("sizeof(shard) = %d, want at least 64: neighbouring shards share a cache line", got)
	}
}
""",
    context="A metrics counter behind one mutex is the bottleneck at sixteen cores. Replacing it with atomics per key is not possible — the keys are dynamic and the map itself needs protection.",
    task=[
        "Increment `key`'s counter in the shard that owns it.",
        "Hold only that shard's lock, and only for the update.",
        "Correct under concurrent use: no lost increments, no race.",
    ],
    examples=[
        ('c.Add("a", 1); c.Add("a", 2)', 'Total()["a"] == 3', None),
        ("16 workers x 1000 adds over 8 keys", "2000 per key", None),
        ("shardFor of one key", "always the same shard", None),
    ],
    topics=[
        ("Lock striping", "Splitting one lock into n reduces contention by roughly n."),
        ("Deterministic sharding", "The same key must always map to the same shard, or increments split and are folded wrongly."),
        ("Padding the shards", "Adjacent mutexes on one cache line reintroduce contention in hardware."),
        ("Fold after the writers", "`Total` is a read-side operation, not part of the hot path."),
    ],
    hint="One shard, one lock, one map update. The routing is already written.",
    intuition="Contention is about how many writers want the same lock, not how many writers there are. Hashing the key spreads them over independent locks, and correctness survives because each key always lands on the same one.",
    approach=[
        "`c.shardFor(key)` to route.",
        "Lock that shard, add to its map entry, unlock.",
    ],
    walkthrough="With 16 shards and 16 workers, writers collide about a sixteenth as often as with a single lock. `Total` then walks the shards once and sums the per-shard counts for each key.",
    pitfalls=[
        "Locking every shard, which is correct and no faster than one mutex.",
        "Holding the lock across the hash computation, which lengthens the critical section for nothing.",
        "Using a per-call random seed, which would scatter one key across shards.",
    ],
)

P(
    "staff",
    name="viewsafety",
    title="A View The Caller Cannot Corrupt",
    sig="func Window(b []byte, off, n int) ([]byte, bool)",
    doc="""Window returns the n bytes of b starting at off, as a view whose
capacity is exactly n.

The caller may append to the result, so the capacity must not let that
append reach the bytes after the window.

Examples:

	Window(buf, 2, 3) => buf[2:5] with capacity 3, true""",
    imports=['"unsafe"'],
    solution="""if off < 0 || n < 0 || off+n > len(b) {
	return nil, false
}
if n == 0 {
	return nil, true
}
p := unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), off)
return unsafe.Slice((*byte)(p), n)[:n:n], true""",
    tests="""
import (
	"bytes"
	"testing"
	"unsafe"
)

func TestWindow(t *testing.T) {
	b := []byte("abcdef")
	got, ok := Window(b, 2, 3)
	if !ok || !bytes.Equal(got, []byte("cde")) {
		t.Errorf("Window = %q, %v, want \\"cde\\", true", got, ok)
	}
}

func TestWindowCapacityIsExact(t *testing.T) {
	b := []byte("abcdef")
	got, ok := Window(b, 2, 3)
	if !ok {
		t.Fatal("Window reported false")
	}
	if cap(got) != 3 {
		t.Fatalf("cap = %d, want 3", cap(got))
	}
	got = append(got, 'Z')
	if b[5] == 'Z' {
		t.Error("the append wrote past the window into the caller's buffer")
	}
	if string(b) != "abcdef" {
		t.Errorf("b = %q, want \\"abcdef\\"", b)
	}
}

func TestWindowSharesUntilItGrows(t *testing.T) {
	b := []byte("abcdef")
	got, _ := Window(b, 0, 3)
	got[0] = 'X'
	if b[0] != 'X' {
		t.Error("the window does not share the buffer")
	}
}

func TestWindowBounds(t *testing.T) {
	b := []byte("abcdef")
	for _, c := range [][2]int{{-1, 2}, {0, -1}, {4, 3}, {7, 0}, {0, 7}} {
		if _, ok := Window(b, c[0], c[1]); ok {
			t.Errorf("Window(off=%d, n=%d) reported ok, want false", c[0], c[1])
		}
	}
}

func TestWindowZeroLength(t *testing.T) {
	b := []byte("abc")
	got, ok := Window(b, 1, 0)
	if !ok || len(got) != 0 {
		t.Errorf("Window = %q, %v, want empty, true", got, ok)
	}
}

func TestWindowAllocatesNothing(t *testing.T) {
	b := make([]byte, 4096)
	var sink []byte
	if n := testing.AllocsPerRun(200, func() { sink, _ = Window(b, 8, 1024) }); n != 0 {
		t.Errorf("Window made %v allocations, want 0", n)
	}
	_ = sink
}

func TestWindowStartsWhereAsked(t *testing.T) {
	b := []byte("abcdef")
	got, _ := Window(b, 2, 2)
	want := (*byte)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), 2))
	if unsafe.SliceData(got) != want {
		t.Error("the window does not start at the requested offset")
	}
}
""",
    context="A framing layer hands each parser a view of the shared read buffer. One parser appends to its view, the append fits in the spare capacity, and the next frame in the buffer is silently rewritten.",
    task=[
        "Return the `n` bytes of `b` starting at `off`, sharing the storage.",
        "The result's capacity must equal its length, so an append cannot reach the following bytes.",
        "Report false for a negative offset or length, or a window running past the end.",
        "Zero allocations; a zero-length window is legal.",
    ],
    examples=[
        ('Window([]byte("abcdef"), 2, 3)', '"cde", true', None),
        ("append to the result", "does not touch b[5]", "The capacity is exactly 3."),
        ("Window(b, 4, 3) on a 6-byte buffer", "nil, false", None),
    ],
    topics=[
        ("Three-index slicing", "`s[:n:n]` is what makes the capacity a boundary rather than a suggestion."),
        ("Spare capacity is someone else's data", "A view into a shared buffer must not expose the rest of it."),
        ("Explicit bounds checks", "`off+n > len(b)` is the check the runtime no longer performs for you."),
        ("unsafe.Add for the start", "The offset is applied in pointer space."),
    ],
    hint="Getting the right bytes is half the job. The other half is the third index.",
    intuition="A slice's capacity is a licence to write. Handing out a view with capacity beyond its length hands out a licence over memory the caller was never given — and `append` will use it without a word.",
    approach=[
        "Validate `off`, `n` and `off+n` against `len(b)`.",
        "Return nil for a zero-length window.",
        "Build the view from the offset pointer and cap it with `[:n:n]`.",
    ],
    walkthrough="Without the capacity cap, `Window(b, 2, 3)` on a six-byte buffer yields capacity 4, so one append overwrites `b[5]`. With `[:n:n]` the append is forced to allocate and the buffer is untouched.",
    pitfalls=[
        "`b[off : off+n]` alone — correct bytes, wrong capacity.",
        "Checking `off < len(b)` instead of `off+n <= len(b)`.",
        "Building the window with a nil data pointer when `n` is 0.",
    ],
)

P(
    "staff",
    name="lazyinit",
    title="Build It Once, However Many Ask",
    sig="func (t *Table) Lookup(k string) (int, bool)",
    doc="""Lookup returns the value for k, building the table's index on first use.

The index is expensive, the callers are concurrent, and it must be built
exactly once — every later lookup should be a plain map read.

Examples:

	t := NewTable(pairs); t.Lookup("a") => the value for "a" """,
    imports=['"sync"', '"sync/atomic"'],
    extra="""// Builds counts how many times the index has been constructed.
var Builds atomic.Int64

// Table indexes a slice of pairs lazily.
type Table struct {
	once  sync.Once
	pairs [][2]string
	index map[string]int
}

// NewTable returns a table over pairs, without building the index.
func NewTable(pairs [][2]string) *Table {
	return &Table{pairs: pairs}
}

// build constructs the index. It must run at most once per table.
func (t *Table) build() {
	Builds.Add(1)
	t.index = make(map[string]int, len(t.pairs))
	for i, p := range t.pairs {
		t.index[p[0]] = i
	}
}""",
    solution="""t.once.Do(t.build)
i, ok := t.index[k]
return i, ok""",
    tests="""
import (
	"strconv"
	"sync"
	"testing"
)

func pairs(n int) [][2]string {
	out := make([][2]string, n)
	for i := range out {
		out[i] = [2]string{"k" + strconv.Itoa(i), "v"}
	}
	return out
}

func TestLookup(t *testing.T) {
	tbl := NewTable([][2]string{{"a", "1"}, {"b", "2"}})
	if i, ok := tbl.Lookup("a"); !ok || i != 0 {
		t.Errorf("Lookup(a) = %d, %v, want 0, true", i, ok)
	}
	if i, ok := tbl.Lookup("b"); !ok || i != 1 {
		t.Errorf("Lookup(b) = %d, %v, want 1, true", i, ok)
	}
	if _, ok := tbl.Lookup("missing"); ok {
		t.Error("Lookup(missing) reported ok, want false")
	}
}

func TestLookupEmptyTable(t *testing.T) {
	tbl := NewTable(nil)
	if _, ok := tbl.Lookup("a"); ok {
		t.Error("Lookup on an empty table reported ok, want false")
	}
}

func TestBuildsOncePerTable(t *testing.T) {
	tbl := NewTable(pairs(100))
	before := Builds.Load()
	for i := 0; i < 500; i++ {
		tbl.Lookup("k1")
	}
	if got := Builds.Load() - before; got != 1 {
		t.Errorf("the index was built %d times, want 1", got)
	}
}

func TestSeparateTablesBuildSeparately(t *testing.T) {
	before := Builds.Load()
	a := NewTable(pairs(4))
	b := NewTable(pairs(4))
	a.Lookup("k0")
	b.Lookup("k0")
	if got := Builds.Load() - before; got != 2 {
		t.Errorf("built %d times, want 2: the once must be per table", got)
	}
}

func TestConcurrentFirstUse(t *testing.T) {
	tbl := NewTable(pairs(1000))
	before := Builds.Load()
	var wg sync.WaitGroup
	const workers = 32
	wg.Add(workers)
	results := make([]bool, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			_, ok := tbl.Lookup("k" + strconv.Itoa(w))
			results[w] = ok
		}(w)
	}
	wg.Wait()
	if got := Builds.Load() - before; got != 1 {
		t.Errorf("the index was built %d times under 32 concurrent first uses, want 1", got)
	}
	for w, ok := range results {
		if !ok {
			t.Fatalf("worker %d did not find its key: it read the index before it was built", w)
		}
	}
}
""",
    context="A lookup table is built on first use behind a \"if index == nil\" check. Under a cold-start burst several goroutines see nil at once, three indexes are built, and two of them are thrown away while readers are using them.",
    task=[
        "Return the index of `k` in the table's pairs, building the index on first use.",
        "The index must be built exactly once per table, even under concurrent first use.",
        "Every caller must observe a fully built index.",
    ],
    examples=[
        ('NewTable([][2]string{{"a","1"}}).Lookup("a")', "0, true", None),
        ("500 lookups", "one build", None),
        ("32 concurrent first lookups", "one build, all succeed", None),
    ],
    topics=[
        ("sync.Once", "Runs the function once and blocks the others until it has finished."),
        ("The happens-before it provides", "Every `Do` returns after the initialisation's writes are visible."),
        ("Check-then-act is not atomic", "A nil check plus an assignment is a race with a benign-looking symptom."),
        ("Per-instance state", "The `Once` lives in the table, so two tables build independently."),
    ],
    hint="One line before the map read. The build function is already written.",
    intuition="Lazy initialisation is a race unless the check and the build are one indivisible step. `sync.Once` provides exactly that, and it also provides the memory ordering that makes the built map safe to read afterwards.",
    approach=[
        "`t.once.Do(t.build)`.",
        "Read `t.index[k]` and return the comma-ok result.",
    ],
    walkthrough="Under 32 concurrent first calls, one goroutine runs `build` and the other 31 block inside `Do` until it returns — so every one of them reads a complete index.",
    pitfalls=[
        "`if t.index == nil { t.build() }` — the race the `Once` exists to remove.",
        "Copying the `Table` after first use; a `sync.Once` must not be copied.",
        "Putting the `Once` in a package-level variable, which would build only the first table's index.",
    ],
)
