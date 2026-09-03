"""10-advanced-topics — rotation 6: 5 puzzles per level."""

SPECS = []


def P(level, **kw):
    kw["level"] = level
    kw.setdefault("sub", "10-advanced-topics")
    kw.setdefault("mode", "stub")
    kw.setdefault("stub", 'panic("not implemented")')
    kw.setdefault("imports", [])
    kw.setdefault("extra", "")
    SPECS.append(kw)


# ---------------------------------------------------------------- junior -----

P(
    "junior",
    name="rowat",
    title="Reach Into A Slice Of Slices",
    sig="func Row(g [][]int, i int) ([]int, bool)",
    doc="""Row returns the i-th row of g and whether it exists.

An out-of-range index is a missing row, not a panic. The row is returned
as a view, so writes through it reach g.

Examples:

	Row([][]int{{1}, {2}}, 1) => []int{2}, true""",
    solution="""if i < 0 || i >= len(g) {
	return nil, false
}
return g[i], true""",
    tests="""
import "testing"

func TestRow(t *testing.T) {
	g := [][]int{{1, 2}, {3}}
	got, ok := Row(g, 0)
	if !ok || len(got) != 2 || got[0] != 1 {
		t.Errorf("Row = %v, %v, want [1 2], true", got, ok)
	}
	got, ok = Row(g, 1)
	if !ok || len(got) != 1 || got[0] != 3 {
		t.Errorf("Row = %v, %v, want [3], true", got, ok)
	}
}

func TestRowOutOfRange(t *testing.T) {
	g := [][]int{{1}}
	for _, i := range []int{-1, 1, 99} {
		if _, ok := Row(g, i); ok {
			t.Errorf("Row(_, %d) reported ok, want false", i)
		}
	}
	if _, ok := Row(nil, 0); ok {
		t.Error("Row(nil, 0) reported ok, want false")
	}
}

func TestRowIsAView(t *testing.T) {
	g := [][]int{{1, 2}}
	row, _ := Row(g, 0)
	row[0] = 99
	if g[0][0] != 99 {
		t.Error("the row is a copy; it must be a view into g")
	}
}

func TestRowNilRow(t *testing.T) {
	g := [][]int{nil}
	got, ok := Row(g, 0)
	if !ok {
		t.Error("a nil row still exists")
	}
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}
""",
    context="A grid helper indexes rows directly. A request with a bad row number takes the whole service down with an index-out-of-range panic.",
    task=[
        "Return the `i`-th row of `g` and whether it exists.",
        "An out-of-range index reports false instead of panicking.",
        "The row is a view — writes through it must reach `g`.",
    ],
    examples=[
        ("Row([][]int{{1},{2}}, 1)", "[2], true", None),
        ("Row(g, -1)", "nil, false", None),
        ("row[0] = 99 after Row(g, 0)", "g[0][0] is 99", "The row is not a copy."),
    ],
    topics=[
        ("Bounds checks you write yourself", "The runtime's check panics; the caller wanted an answer."),
        ("Slice of slices", "The outer slice holds headers; indexing it copies a header, not the elements."),
        ("Views share storage", "Writing through the returned row reaches the grid."),
    ],
    hint="Check both ends of the range before indexing.",
    intuition="Indexing is only safe inside the range, and the range is `0` up to `len(g)`. Checking it yourself turns a crash into a value the caller can branch on.",
    approach=[
        "Return `nil, false` for `i < 0` or `i >= len(g)`.",
        "Return `g[i], true`.",
    ],
    walkthrough="`Row(nil, 0)` sees `len(g) == 0`, so 0 is already out of range and the nil check is not needed separately.",
    pitfalls=[
        "Checking only `i >= len(g)` and letting a negative index panic.",
        "Copying the row to \"be safe\", which breaks the caller's writes and allocates.",
    ],
)

P(
    "junior",
    name="mapget",
    title="Missing Or Zero",
    sig="func Get(m map[string]int, key string) (int, bool)",
    doc="""Get returns the value stored under key and whether the key was present.

A missing key reads as 0, which is also a value a key can hold — only the
second result tells them apart.

Examples:

	Get(map[string]int{"a": 0}, "a") => 0, true""",
    solution="""v, ok := m[key]
return v, ok""",
    tests="""
import "testing"

func TestGet(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	if v, ok := Get(m, "a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := Get(m, "zero"); !ok || v != 0 {
		t.Errorf("Get(zero) = %d, %v, want 0, true: a stored zero is present", v, ok)
	}
	if v, ok := Get(m, "missing"); ok || v != 0 {
		t.Errorf("Get(missing) = %d, %v, want 0, false", v, ok)
	}
}

func TestGetNilMap(t *testing.T) {
	if v, ok := Get(nil, "a"); ok || v != 0 {
		t.Errorf("Get(nil) = %d, %v, want 0, false", v, ok)
	}
}

func TestGetEmptyKey(t *testing.T) {
	m := map[string]int{"": 5}
	if v, ok := Get(m, ""); !ok || v != 5 {
		t.Errorf("Get(\\"\\") = %d, %v, want 5, true", v, ok)
	}
}
""",
    context="A feature-flag lookup treats a zero count as \"not configured\" and silently re-enables a flag someone deliberately set to zero.",
    task=[
        "Return the value under `key` and whether the key was present.",
        "A stored zero must report present.",
        "A nil map reports absent without panicking.",
    ],
    examples=[
        ('Get(map[string]int{"a":0}, "a")', "0, true", "The zero was stored."),
        ('Get(m, "missing")', "0, false", None),
        ('Get(nil, "a")', "0, false", None),
    ],
    topics=[
        ("The comma-ok form", "A map index in a two-value assignment reports presence."),
        ("The zero value is ambiguous", "Without the boolean, absent and zero are the same reading."),
        ("Reading a nil map is legal", "It behaves as an empty map; only writing panics."),
    ],
    hint="One statement, two results.",
    intuition="A map index always produces a value — the zero value when the key is missing. The second result is the only thing that distinguishes \"stored zero\" from \"not there\", which is exactly the distinction most map bugs turn on.",
    approach=[
        "`v, ok := m[key]` and return both.",
    ],
    walkthrough="Reading a nil map yields the zero value and false, so no separate nil check is needed.",
    pitfalls=[
        "`if m[key] != 0` as a presence test, which is the bug this closes.",
        "Writing to a nil map, which does panic — only reads are safe.",
    ],
)

P(
    "junior",
    name="fieldcount",
    title="How Many Fields Does This Have",
    sig="func FieldCount(v any) (total, exported int)",
    doc="""FieldCount returns how many fields v's struct type has in total, and
how many of them are exported.

A non-struct, or a nil interface, reports 0, 0.

Examples:

	FieldCount(rec{}) => 3, 2""",
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil || t.Kind() != reflect.Struct {
	return 0, 0
}
total = t.NumField()
for i := 0; i < total; i++ {
	if t.Field(i).IsExported() {
		exported++
	}
}
return total, exported""",
    tests="""
import "testing"

type rec struct {
	A      int
	B      string
	hidden bool
}

func TestFieldCount(t *testing.T) {
	total, exported := FieldCount(rec{})
	if total != 3 || exported != 2 {
		t.Errorf("FieldCount = %d, %d, want 3, 2", total, exported)
	}
}

func TestFieldCountEmptyStruct(t *testing.T) {
	total, exported := FieldCount(struct{}{})
	if total != 0 || exported != 0 {
		t.Errorf("FieldCount = %d, %d, want 0, 0", total, exported)
	}
}

func TestFieldCountAllUnexported(t *testing.T) {
	type private struct{ a, b int }
	total, exported := FieldCount(private{})
	if total != 2 || exported != 0 {
		t.Errorf("FieldCount = %d, %d, want 2, 0", total, exported)
	}
}

func TestFieldCountNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, "s", []int{1}, &rec{}} {
		total, exported := FieldCount(in)
		if total != 0 || exported != 0 {
			t.Errorf("FieldCount(%#v) = %d, %d, want 0, 0", in, total, exported)
		}
	}
}
""",
    context="A serialiser reports how much of each struct it can actually see. Unexported fields are invisible to it, and the counts explain a lot of confused bug reports.",
    task=[
        "Return the total field count and the exported field count of `v`'s struct type.",
        "Report 0, 0 for a non-struct, a pointer to a struct, or a nil interface.",
    ],
    examples=[
        ("FieldCount(rec{})", "3, 2", "One unexported field."),
        ("FieldCount(struct{}{})", "0, 0", None),
        ("FieldCount(&rec{})", "0, 0", "A pointer is not a struct."),
    ],
    topics=[
        ("Type.NumField", "The field count is part of the type, not the value."),
        ("StructField.IsExported", "Export status is metadata, not something to infer from the name."),
        ("Guard the kind first", "`NumField` panics on anything but a struct."),
    ],
    hint="Named results let you count in place without extra variables.",
    intuition="Everything about a struct's shape is available at run time — how many fields, what they are called, and which of them the rest of the program is allowed to touch.",
    approach=[
        "Take `reflect.TypeOf(v)`; bail out for nil or non-struct.",
        "Set `total` from `NumField`, then count the exported ones.",
    ],
    walkthrough="`rec` has three fields; `hidden` starts with a lower-case letter, so `IsExported` is false and the exported count is 2.",
    pitfalls=[
        "Inferring export status from the first character instead of asking; the rule involves Unicode, not just ASCII.",
        "Forgetting that a pointer must be dereferenced first — here the spec says not to.",
    ],
)

P(
    "junior",
    name="sliceandarray",
    title="A Slice Header Versus An Array",
    sig="func Sizes() (arr, sl uintptr)",
    doc="""Sizes returns the size of a [8]int array and of an []int slice header.

The array's size is its contents; the slice's is three words, whatever it
points at.

Examples:

	Sizes() => 64, 24 on a 64-bit build""",
    imports=['"unsafe"'],
    solution="""var (
	a [8]int
	s []int
)
return unsafe.Sizeof(a), unsafe.Sizeof(s)""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestSizes(t *testing.T) {
	var (
		a [8]int
		s []int
	)
	arr, sl := Sizes()
	if arr != unsafe.Sizeof(a) {
		t.Errorf("array = %d, want %d", arr, unsafe.Sizeof(a))
	}
	if sl != unsafe.Sizeof(s) {
		t.Errorf("slice = %d, want %d", sl, unsafe.Sizeof(s))
	}
}

func TestArrayIsBigger(t *testing.T) {
	arr, sl := Sizes()
	if arr <= sl {
		t.Errorf("array = %d, slice = %d, want the array to be larger", arr, sl)
	}
}

func TestSliceSizeIsIndependentOfLength(t *testing.T) {
	_, sl := Sizes()
	long := make([]int, 100000)
	if got := unsafe.Sizeof(long); got != sl {
		t.Errorf("a 100000-element slice header is %d, want %d: the header does not grow", got, sl)
	}
}

func TestArrayIsEightWords(t *testing.T) {
	arr, _ := Sizes()
	var one int
	if arr != 8*unsafe.Sizeof(one) {
		t.Errorf("array = %d, want %d", arr, 8*unsafe.Sizeof(one))
	}
}
""",
    context="A memory estimate multiplies the record count by the size of a struct containing a slice. The estimate is off by the entire payload, because the slice's elements were never counted.",
    task=[
        "Return `unsafe.Sizeof` for a `[8]int` and for an `[]int`.",
        "Derive both from declared variables rather than writing numbers.",
    ],
    examples=[
        ("Sizes()", "64, 24", "Eight ints; a pointer, a length and a capacity."),
        ("Sizeof of a 100000-element slice", "still 24", "The header does not grow."),
        ("array vs slice", "the array is larger", None),
    ],
    topics=[
        ("Arrays carry their contents", "The length is in the type, so the size includes every element."),
        ("Slices are three words", "Pointer, length and capacity — the elements live elsewhere."),
        ("Sizeof measures the type", "Never what a pointer or header refers to."),
    ],
    hint="Declare one of each and measure them.",
    intuition="This is the difference between the two types in one number. An array *is* its data; a slice is a small description of data that lives somewhere else — which is why passing an array copies everything and passing a slice copies three words.",
    approach=[
        "Declare `var a [8]int` and `var s []int`.",
        "Return `unsafe.Sizeof` of each.",
    ],
    walkthrough="On a 64-bit build an int is 8 bytes, so the array is 64 and the slice header is 24 — regardless of how many elements the slice has.",
    pitfalls=[
        "Expecting `Sizeof` on a slice to grow with its length.",
        "Writing 64 and 24 as literals; both depend on the word size.",
    ],
)

P(
    "junior",
    name="swapptr",
    title="Exchange Two Values Through Pointers",
    sig="func Swap(a, b *int)",
    doc="""Swap exchanges the values a and b point at.

If either pointer is nil, nothing happens. Nothing is allocated.

Examples:

	x, y := 1, 2; Swap(&x, &y) => x is 2, y is 1""",
    solution="""if a == nil || b == nil {
	return
}
*a, *b = *b, *a""",
    tests="""
import "testing"

func TestSwap(t *testing.T) {
	x, y := 1, 2
	Swap(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("x, y = %d, %d, want 2, 1", x, y)
	}
}

func TestSwapSamePointer(t *testing.T) {
	x := 5
	Swap(&x, &x)
	if x != 5 {
		t.Errorf("x = %d, want 5", x)
	}
}

func TestSwapNil(t *testing.T) {
	x := 1
	Swap(&x, nil)
	Swap(nil, &x)
	Swap(nil, nil)
	if x != 1 {
		t.Errorf("x = %d, want 1", x)
	}
}

func TestSwapAllocatesNothing(t *testing.T) {
	x, y := 1, 2
	if n := testing.AllocsPerRun(200, func() { Swap(&x, &y) }); n != 0 {
		t.Errorf("Swap made %v allocations, want 0", n)
	}
}
""",
    context="A sort helper takes the values instead of pointers. It swaps its own copies, the slice never changes, and the sort loops forever.",
    task=[
        "Exchange the values `a` and `b` point at.",
        "Do nothing when either pointer is nil.",
        "Zero allocations.",
    ],
    examples=[
        ("x, y := 1, 2; Swap(&x, &y)", "x is 2, y is 1", None),
        ("Swap(&x, &x)", "x unchanged", "Swapping a value with itself is a no-op."),
        ("Swap(&x, nil)", "nothing happens", None),
    ],
    topics=[
        ("Writing through a pointer", "`*p = v` reaches the caller's variable."),
        ("Tuple assignment", "`*a, *b = *b, *a` evaluates the right side before assigning."),
        ("Aliasing is fine here", "The tuple form is correct even when both pointers are the same."),
    ],
    hint="One statement does it, and it happens to handle the aliased case too.",
    intuition="Go evaluates the whole right-hand side before assigning any of it, so a tuple swap needs no temporary and stays correct even when both pointers address the same variable.",
    approach=[
        "Return early when either pointer is nil.",
        "`*a, *b = *b, *a`.",
    ],
    walkthrough="With `Swap(&x, &x)` the right side reads 5 and 5, then writes 5 and 5 — unchanged, as it should be. A sequential swap with a temporary is also correct; the tuple form is just shorter.",
    pitfalls=[
        "Taking `int` parameters, which swaps copies and does nothing.",
        "Writing `*a = *b` then `*b = *a`, which loses the original value of `*a`.",
    ],
)

# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="insertat",
    title="Make Room In The Middle",
    sig="func InsertAt(s []int, i, v int) []int",
    doc="""InsertAt inserts v at index i, shifting the rest up, and returns the
extended slice.

i is clamped into [0, len(s)]. The shift must not lose the element it
overwrites.

Examples:

	InsertAt([]int{1, 3}, 1, 2) => []int{1, 2, 3}""",
    solution="""if i < 0 {
	i = 0
}
if i > len(s) {
	i = len(s)
}
s = append(s, 0)
copy(s[i+1:], s[i:])
s[i] = v
return s""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	if got := InsertAt([]int{1, 3}, 1, 2); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{2, 3}, 0, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{1, 2}, 2, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
}

func TestInsertAtClamps(t *testing.T) {
	if got := InsertAt([]int{1, 2}, 99, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
	if got := InsertAt([]int{2, 3}, -5, 1); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("InsertAt = %v, want [1 2 3]", got)
	}
}

func TestInsertAtEmpty(t *testing.T) {
	if got := InsertAt(nil, 0, 7); !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("InsertAt = %v, want [7]", got)
	}
}

func TestInsertAtLongShift(t *testing.T) {
	s := make([]int, 0, 100)
	for i := 0; i < 50; i++ {
		s = append(s, i)
	}
	got := InsertAt(s, 0, -1)
	if got[0] != -1 {
		t.Fatalf("got[0] = %d, want -1", got[0])
	}
	for i := 0; i < 50; i++ {
		if got[i+1] != i {
			t.Fatalf("got[%d] = %d, want %d: the shift lost an element", i+1, got[i+1], i)
		}
	}
}

func TestInsertAtReusesCapacity(t *testing.T) {
	s := make([]int, 2, 8)
	if n := testing.AllocsPerRun(100, func() { _ = InsertAt(s[:2], 1, 9) }); n != 0 {
		t.Errorf("InsertAt made %v allocations, want 0 when the capacity allows", n)
	}
}
""",
    context="An ordered list inserts by building a new slice from two sub-slices. Every insertion allocates and copies the whole list twice.",
    task=[
        "Insert `v` at index `i`, shifting the rest up.",
        "Clamp `i` into `[0, len(s)]`.",
        "With spare capacity, allocate nothing.",
    ],
    examples=[
        ("InsertAt([]int{1,3}, 1, 2)", "[1 2 3]", None),
        ("InsertAt([]int{1,2}, 99, 3)", "[1 2 3]", "The index is clamped to the end."),
        ("InsertAt(nil, 0, 7)", "[7]", None),
    ],
    topics=[
        ("Grow first, then shift", "Appending a placeholder makes room for the copy."),
        ("copy handles overlap", "A right shift over itself works because `copy` behaves like memmove."),
        ("Capacity reuse", "The append is free when the room is already there."),
    ],
    hint="Three steps: make room, shift, then write.",
    intuition="Insertion is a shift, and a shift needs somewhere to shift into. Appending a placeholder first turns the problem into one `copy` and one assignment.",
    approach=[
        "Clamp `i`.",
        "`append` a zero to extend the length.",
        "`copy(s[i+1:], s[i:])` to shift the tail up.",
        "Write `v` at `i` and return.",
    ],
    walkthrough="For [1 3] and i = 1: appending gives [1 3 0]; the copy shifts 3 up to index 2 giving [1 3 3]; writing 2 at index 1 gives [1 2 3].",
    pitfalls=[
        "Shifting before extending, which writes past the length.",
        "`append(s[:i], append([]int{v}, s[i:]...)...)` — correct, and it allocates twice.",
    ],
)

P(
    "middle",
    name="concatloop",
    title="The Concatenation That Copied Everything Each Time",
    mode="bug",
    sig="func Join(parts []string) string",
    doc="""Join concatenates parts end to end.

Strings are immutable, so += allocates a new string and copies both sides
every round — quadratic in the total length.

Examples:

	Join([]string{"a", "bc"}) => "abc" """,
    imports=['"strings"'],
    buggy="""out := ""
for _, p := range parts {
	out += p
}
return out""",
    solution="""n := 0
for _, p := range parts {
	n += len(p)
}
var b strings.Builder
b.Grow(n)
for _, p := range parts {
	b.WriteString(p)
}
return b.String()""",
    tests="""
import (
	"strings"
	"testing"
)

var sink string

func TestJoin(t *testing.T) {
	if got := Join([]string{"a", "bc", "d"}); got != "abcd" {
		t.Errorf("Join = %q, want \\"abcd\\"", got)
	}
	if got := Join(nil); got != "" {
		t.Errorf("Join = %q, want empty", got)
	}
	if got := Join([]string{"", ""}); got != "" {
		t.Errorf("Join = %q, want empty", got)
	}
	if got := Join([]string{"solo"}); got != "solo" {
		t.Errorf("Join = %q, want \\"solo\\"", got)
	}
}

func TestJoinLong(t *testing.T) {
	parts := make([]string, 200)
	for i := range parts {
		parts[i] = "chunk"
	}
	got := Join(parts)
	if got != strings.Repeat("chunk", 200) {
		t.Errorf("Join produced %d bytes, want %d", len(got), 200*5)
	}
}

func TestJoinAllocationsAreBounded(t *testing.T) {
	parts := make([]string, 128)
	for i := range parts {
		parts[i] = "chunk"
	}
	n := testing.AllocsPerRun(20, func() { sink = Join(parts) })
	if n > 3 {
		t.Errorf("Join made %v allocations for 128 parts, want at most 3: build in one buffer", n)
	}
}
""",
    context="A report builder concatenates a few hundred fragments with `+=`. The fragments are short, the report is not, and the function dominates the profile.",
    task=[
        "Concatenate the parts end to end.",
        "Fix the single bug so the work is linear: at most a handful of allocations for 128 parts.",
        "An empty input returns the empty string.",
    ],
    examples=[
        ('Join([]string{"a","bc","d"})', '"abcd"', None),
        ("128 parts", "at most 3 allocations", "Not 128."),
        ("Join(nil)", '""', None),
    ],
    topics=[
        ("String immutability", "`+=` cannot extend a string; it builds a new one and copies both sides."),
        ("Quadratic accumulation", "Joining n parts copies O(n²) bytes."),
        ("strings.Builder", "One growing buffer, handed out as a string without a final copy."),
        ("Grow", "Sizing it once removes even the doubling."),
    ],
    hint="The loop is fine. What it accumulates into is not.",
    intuition="Each `+=` allocates a string as long as everything so far and copies it. By the last part you have copied the whole report once per fragment — the cost is invisible at three parts and fatal at three hundred.",
    approach=[
        "Sum the parts' lengths.",
        "`Grow` a `strings.Builder` to that size.",
        "Write every part and return `b.String()`.",
    ],
    walkthrough="128 five-byte parts total 640 bytes. `+=` allocates 128 strings and copies about 41,000 bytes; the builder allocates once and copies 640.",
    pitfalls=[
        "Using a builder without `Grow`, which is linear but still reallocates a few times.",
        "`strings.Join(parts, \"\")` is the real-world answer — the point is why `+=` is not.",
    ],
)

P(
    "middle",
    name="reflectsum",
    title="Total Any Slice Of Integers",
    sig="func Sum(v any) (int64, bool)",
    doc="""Sum totals v when it is a slice or array of a signed integer kind, and
reports whether it could.

Examples:

	Sum([]int32{1, 2}) => 3, true""",
    imports=['"reflect"'],
    solution="""rv := reflect.ValueOf(v)
if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
	return 0, false
}
switch rv.Type().Elem().Kind() {
case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
default:
	return 0, false
}
var total int64
for i := 0; i < rv.Len(); i++ {
	total += rv.Index(i).Int()
}
return total, true""",
    tests="""
import "testing"

type myInt int

func TestSum(t *testing.T) {
	cases := []struct {
		in    any
		total int64
		ok    bool
	}{
		{[]int{1, 2, 3}, 6, true},
		{[]int32{-1, 1}, 0, true},
		{[]int64{1 << 40, 1 << 40}, 1 << 41, true},
		{[3]int{1, 2, 3}, 6, true},
		{[]myInt{2, 3}, 5, true},
		{[]int{}, 0, true},
		{[]string{"a"}, 0, false},
		{[]uint8{1}, 0, false},
		{3, 0, false},
		{nil, 0, false},
		{map[string]int{"a": 1}, 0, false},
	}
	for _, c := range cases {
		total, ok := Sum(c.in)
		if total != c.total || ok != c.ok {
			t.Errorf("Sum(%#v) = %d, %v, want %d, %v", c.in, total, ok, c.total, c.ok)
		}
	}
}

func TestSumWideAccumulator(t *testing.T) {
	in := make([]int32, 8)
	for i := range in {
		in[i] = 1 << 30
	}
	if total, ok := Sum(in); !ok || total != 8<<30 {
		t.Errorf("Sum = %d, %v, want %d, true: the total must not overflow", total, ok, int64(8)<<30)
	}
}
""",
    context="A metrics helper sums whichever integer slice a plugin returns. The plugin API is `any`, and a type switch over every width is four near-identical branches.",
    task=[
        "Total `v` when it is a slice or array of a signed integer kind.",
        "Report false for any other kind, including unsigned elements and non-sequences.",
        "Accumulate in int64 so wide values cannot overflow.",
    ],
    examples=[
        ("Sum([]int32{-1,1})", "0, true", None),
        ("Sum([]myInt{2,3})", "5, true", "A named type's kind is still int."),
        ("Sum([]uint8{1})", "0, false", "Unsigned is out of scope."),
    ],
    topics=[
        ("Value.Int", "Reads any signed integer kind as an int64."),
        ("Type.Elem", "A slice or array type knows its element type."),
        ("Kind covers named types", "`myInt`'s kind is int, so it is handled without a special case."),
        ("Index into a Value", "`rv.Index(i)` works for both slices and arrays."),
    ],
    hint="Two checks: the container's kind, then the element's.",
    intuition="Reflection collapses \"one branch per integer width\" into a single loop, because `Int()` normalises every signed kind to int64. The type switch a caller would have written becomes two kind checks.",
    approach=[
        "Reject anything that is not a slice or array.",
        "Reject an element kind that is not a signed integer.",
        "Loop with `rv.Index(i).Int()` into an int64.",
    ],
    walkthrough="`[]myInt{2,3}` has element kind int, so `Int()` reads each element and the total is 5 — with no mention of `myInt` anywhere.",
    pitfalls=[
        "Calling `Int()` on an unsigned element, which panics — `Uint()` is the accessor for those.",
        "Accumulating in the element's own width, which overflows for large int32 values.",
    ],
)

P(
    "middle",
    name="samestring",
    title="Do These Two Strings Share Their Bytes",
    sig="func SameBytes(a, b string) bool",
    doc="""SameBytes reports whether a and b are the same length and start at the
same address — that is, whether they share their storage.

Two equal strings may or may not share; this asks about identity, not
equality.

Examples:

	s := "abc"; SameBytes(s, s) => true""",
    imports=['"unsafe"'],
    solution="""if len(a) != len(b) {
	return false
}
if len(a) == 0 {
	return false
}
return unsafe.StringData(a) == unsafe.StringData(b)""",
    tests="""
import (
	"strings"
	"testing"
)

func TestSameBytesIdentical(t *testing.T) {
	s := strings.Repeat("x", 32)
	if !SameBytes(s, s) {
		t.Error("SameBytes(s, s) = false, want true")
	}
	if !SameBytes(s, s[:32]) {
		t.Error("SameBytes(s, s[:32]) = false, want true")
	}
}

func TestSameBytesSubstrings(t *testing.T) {
	s := strings.Repeat("y", 8)
	if SameBytes(s, s[:4]) {
		t.Error("different lengths must report false")
	}
	if SameBytes(s[1:], s[:7]) {
		t.Error("different starts must report false")
	}
}

func TestSameBytesSeparateCopies(t *testing.T) {
	a := strings.Repeat("z", 16)
	b := string([]byte(a))
	if a != b {
		t.Fatal("the fixture is wrong: the strings must be equal")
	}
	if SameBytes(a, b) {
		t.Error("two separate copies must report false")
	}
}

func TestSameBytesEmpty(t *testing.T) {
	if SameBytes("", "") {
		t.Error("empty strings have no storage to share")
	}
	if SameBytes("a", "") {
		t.Error("different lengths must report false")
	}
}
""",
    context="A cache is supposed to hand out the interned copy of a key. Proving it does — rather than returning an equal but separate string — needs identity, which `==` cannot express.",
    task=[
        "Report whether `a` and `b` have the same length and the same start address.",
        "Two equal but separately allocated strings report false.",
        "Empty strings report false — there is no storage to share.",
    ],
    examples=[
        ("s := \"abc\"; SameBytes(s, s)", "true", None),
        ("SameBytes(s[1:], s[:7])", "false", "Different starts."),
        ('a and string([]byte(a))', "false", "Equal, not identical."),
    ],
    topics=[
        ("unsafe.StringData", "The address of a string's first byte."),
        ("Equality versus identity", "`==` compares bytes; this compares storage."),
        ("Interning is about identity", "The whole point is that repeats share one allocation."),
    ],
    hint="Length first, then the data pointers.",
    intuition="Go gives no way to ask whether two strings are the same object, because normally it should not matter. It matters exactly when you are verifying that a cache or interner is doing its job.",
    approach=[
        "Return false on a length mismatch.",
        "Return false for empty strings.",
        "Compare `unsafe.StringData(a)` with `unsafe.StringData(b)`.",
    ],
    walkthrough="`s` and `s[:32]` on a 32-byte string share both the length and the start, so the answer is true. A round trip through `[]byte` allocates a new array and the answer is false.",
    pitfalls=[
        "Skipping the length check, which would call two different-length strings identical.",
        "Reading anything into a false result — sharing is an implementation detail, not a guarantee.",
    ],
)

P(
    "middle",
    name="deleteinrange",
    title="Delete While You Range",
    sig="func RemoveEven(m map[int]int) int",
    doc="""RemoveEven deletes every entry whose key is even and returns how many
were removed.

Deleting during a range is defined: an entry not yet reached that is
deleted will not be produced.

Examples:

	RemoveEven(map[int]int{1: 1, 2: 2}) => 1""",
    solution="""removed := 0
for k := range m {
	if k%2 == 0 {
		delete(m, k)
		removed++
	}
}
return removed""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestRemoveEven(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	if got := RemoveEven(m); got != 2 {
		t.Errorf("RemoveEven = %d, want 2", got)
	}
	want := map[int]int{1: 1, 3: 3}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("m = %v, want %v", m, want)
	}
}

func TestRemoveEvenNegativeKeys(t *testing.T) {
	m := map[int]int{-2: 1, -1: 1, 0: 1}
	if got := RemoveEven(m); got != 2 {
		t.Errorf("RemoveEven = %d, want 2: -2 and 0 are even", got)
	}
	if _, ok := m[-1]; !ok || len(m) != 1 {
		t.Errorf("m = %v, want map[-1:1]", m)
	}
}

func TestRemoveEvenEmpty(t *testing.T) {
	if got := RemoveEven(nil); got != 0 {
		t.Errorf("RemoveEven(nil) = %d, want 0", got)
	}
	m := map[int]int{}
	if got := RemoveEven(m); got != 0 {
		t.Errorf("RemoveEven = %d, want 0", got)
	}
}

func TestRemoveEvenIsVisibleToTheCaller(t *testing.T) {
	m := map[int]int{2: 2}
	alias := m
	RemoveEven(m)
	if len(alias) != 0 {
		t.Error("the deletion was not applied to the caller's map")
	}
}

func TestRemoveEvenLarge(t *testing.T) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}
	if got := RemoveEven(m); got != 500 {
		t.Errorf("RemoveEven = %d, want 500", got)
	}
	if len(m) != 500 {
		t.Errorf("len = %d, want 500", len(m))
	}
	for k := range m {
		if k%2 == 0 {
			t.Fatalf("key %d survived", k)
		}
	}
}
""",
    context="A cleanup pass collects the keys to remove into a slice first, \"because deleting during iteration is unsafe\". The slice is as large as the map, and the caution was unnecessary.",
    task=[
        "Delete every entry with an even key and return the count.",
        "Negative even keys count; 0 is even.",
        "Modify the caller's map in place; a nil map removes nothing.",
    ],
    examples=[
        ("RemoveEven(map[int]int{1:1, 2:2})", "1", None),
        ("RemoveEven(map[int]int{-2:1, -1:1, 0:1})", "2", "-2 and 0 are even."),
        ("RemoveEven(nil)", "0", None),
    ],
    topics=[
        ("Deleting during range is defined", "The spec allows it; an unreached deleted entry is simply not produced."),
        ("Maps are reference-like", "The deletions reach the caller."),
        ("No second slice needed", "Collecting keys first doubles the memory for nothing."),
    ],
    hint="Range and delete. That is the whole function.",
    intuition="Go's map iteration is explicitly specified to tolerate deletion: entries removed before they are reached are not produced. The defensive two-pass version is a habit carried over from languages where it would crash.",
    approach=[
        "Range the keys.",
        "Delete the even ones and count them.",
    ],
    walkthrough="Over a 1000-entry map, 500 keys are deleted as the loop reaches them, and the loop still visits each surviving key exactly once.",
    pitfalls=[
        "Adding entries during iteration, which *is* unspecified — new keys may or may not be produced.",
        "`k%2 == 0` is correct for negatives too; `k&1 == 0` is as well, but `k%2 == 1` is not a test for odd.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="bufferreset",
    title="The Shared Buffer Nobody Emptied",
    mode="bug",
    sig="func Render(vals []int) string",
    doc="""Render returns vals as decimal numbers joined by '-'.

The package keeps one scratch buffer to avoid allocating per call. A
shared buffer has to be emptied before it is written to.

Examples:

	Render([]int{1, 2}) => "1-2" """,
    imports=['"bytes"', '"strconv"', '"sync"'],
    extra="""var (
	mu      sync.Mutex
	scratch bytes.Buffer
)""",
    buggy="""mu.Lock()
defer mu.Unlock()
for i, v := range vals {
	if i > 0 {
		scratch.WriteByte('-')
	}
	scratch.WriteString(strconv.Itoa(v))
}
return scratch.String()""",
    solution="""mu.Lock()
defer mu.Unlock()
scratch.Reset()
for i, v := range vals {
	if i > 0 {
		scratch.WriteByte('-')
	}
	scratch.WriteString(strconv.Itoa(v))
}
return scratch.String()""",
    tests="""
import (
	"sync"
	"testing"
)

func TestRenderOnce(t *testing.T) {
	if got := Render([]int{1, 2, 3}); got != "1-2-3" {
		t.Errorf("Render = %q, want \\"1-2-3\\"", got)
	}
}

func TestRenderRepeatedly(t *testing.T) {
	for i := 0; i < 200; i++ {
		if got := Render([]int{7}); got != "7" {
			t.Fatalf("call %d: Render = %q, want \\"7\\": the shared buffer was not reset", i, got)
		}
	}
}

func TestRenderEmpty(t *testing.T) {
	if got := Render(nil); got != "" {
		t.Errorf("Render = %q, want empty", got)
	}
}

func TestRenderStaysBounded(t *testing.T) {
	for i := 0; i < 2000; i++ {
		Render([]int{1, 2, 3, 4})
	}
	if got := len(Render([]int{9})); got != 1 {
		t.Errorf("Render returned %d bytes, want 1: the buffer keeps growing", got)
	}
}

func TestRenderConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	bad := make([]string, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			want := strconvItoa(w)
			for i := 0; i < 200; i++ {
				if got := Render([]int{w}); got != want {
					bad[w] = got
					return
				}
			}
		}(w)
	}
	wg.Wait()
	for w, got := range bad {
		if got != "" {
			t.Fatalf("worker %d saw %q", w, got)
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
""",
    context="A renderer keeps one package-level buffer behind a mutex. The first response is correct and every response after it contains the whole history of the process.",
    task=[
        "Render `vals` as decimal numbers joined by `-`.",
        "Fix the single bug so each call starts from an empty buffer.",
        "The mutex must stay — the buffer is shared.",
    ],
    examples=[
        ("Render([]int{1,2,3})", '"1-2-3"', None),
        ("200 calls of Render([]int{7})", 'every call returns "7"', None),
        ("Render(nil)", '""', None),
    ],
    topics=[
        ("Shared state needs resetting", "A package-level buffer keeps whatever the last caller left."),
        ("Buffer.Reset", "Empties it while keeping the memory it has grown."),
        ("Unbounded growth", "Without the reset the buffer grows forever."),
        ("The mutex is not the bug", "Serialisation was correct; the state was not."),
    ],
    hint="The lock is right and the loop is right. What is the buffer's length when the loop starts?",
    intuition="Sharing a buffer to save allocations means inheriting its contents. The lock stops two callers writing at once; it does nothing about the bytes the previous caller left behind.",
    approach=[
        "Take the lock as before.",
        "`scratch.Reset()` before writing.",
        "Render and return the string.",
    ],
    walkthrough='The first call writes "7" and leaves it there. The second appends, returning "77". After 2000 calls the buffer holds thousands of bytes and every response carries all of them.',
    pitfalls=[
        "Resetting after `String()` instead of before writing — a concurrent caller may already be inside the lock's queue with stale expectations.",
        "Removing the shared buffer entirely, which fixes the symptom and gives up the optimisation.",
    ],
)

P(
    "senior",
    name="bigscratch",
    title="The Scratch Array That Went To The Heap",
    mode="bug",
    sig="func Format(v int64) []byte",
    doc="""Format returns v's decimal digits.

The result escapes, so whatever it points into escapes with it — sizing
the scratch buffer for the worst imaginable case then costs that much
heap on every call.

Examples:

	Format(42) => []byte("42")""",
    imports=['"strconv"'],
    buggy="""var buf [4096]byte
return strconv.AppendInt(buf[:0], v, 10)""",
    solution="""return strconv.AppendInt(make([]byte, 0, 20), v, 10)""",
    tests="""
import (
	"bytes"
	"runtime"
	"testing"
)

var sink []byte

func TestFormat(t *testing.T) {
	cases := map[int64]string{
		0: "0", 42: "42", -7: "-7",
		9223372036854775807:  "9223372036854775807",
		-9223372036854775808: "-9223372036854775808",
	}
	for in, want := range cases {
		if got := Format(in); !bytes.Equal(got, []byte(want)) {
			t.Errorf("Format(%d) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatResultsAreIndependent(t *testing.T) {
	a := Format(11)
	b := Format(22)
	if !bytes.Equal(a, []byte("11")) {
		t.Errorf("a = %q, want \\"11\\"", a)
	}
	if !bytes.Equal(b, []byte("22")) {
		t.Errorf("b = %q, want \\"22\\"", b)
	}
}

func TestFormatDoesNotAllocateAKilobyte(t *testing.T) {
	const runs = 2000
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	for i := 0; i < runs; i++ {
		sink = Format(int64(i))
	}
	runtime.ReadMemStats(&after)
	used := after.TotalAlloc - before.TotalAlloc
	if used > runs*128 {
		t.Errorf("allocated %d bytes over %d calls (%d per call), want well under 128 per call: the scratch buffer escapes",
			used, runs, used/runs)
	}
}
""",
    context="A formatter declares a generously sized local array so it never has to grow. The allocation count looks fine and the process allocates four kilobytes for every number it prints.",
    task=[
        "Return `v`'s decimal digits, including the sign.",
        "Fix the single bug so a call allocates on the order of the digits, not kilobytes.",
        "Results from separate calls must be independent.",
    ],
    examples=[
        ("Format(42)", '"42"', None),
        ("Format(-9223372036854775808)", "the full minimum int64", "Twenty bytes is enough for any int64."),
        ("2000 calls", "well under 128 bytes each", None),
    ],
    topics=[
        ("Escape analysis follows the result", "The returned slice points into the scratch, so the scratch escapes."),
        ("A local array is only free while it stays local", "Once it escapes, its full size is allocated."),
        ("Size the buffer to the problem", "Twenty bytes covers every int64."),
        ("Allocation count hides allocation size", "One 4 KiB allocation counts the same as one 20-byte one."),
    ],
    hint="The allocation count is already 1. Look at how many bytes that one allocation is.",
    intuition="A stack array costs nothing until something outlives the frame. Returning a slice of it forces the whole array onto the heap — so the \"generous\" size is not insurance, it is the per-call cost.",
    approach=[
        "Allocate a right-sized buffer instead: twenty bytes holds any int64 with its sign.",
        "Append the digits into it and return the result.",
    ],
    walkthrough="Two thousand calls allocate about 8 MB with the 4 KiB array and about 40 KB with a twenty-byte buffer, for identical output.",
    pitfalls=[
        "Trusting `AllocsPerRun` alone; it counts allocations, not bytes.",
        "Keeping the array and copying out of it, which allocates twice.",
    ],
)

P(
    "senior",
    name="embeddedfields",
    title="Fields That Came From Somewhere Else",
    sig="func Paths(v any) []string",
    doc="""Paths returns the dotted path of every exported leaf field of v,
descending through embedded structs.

An embedded struct contributes its fields' paths under its own name;
named struct fields are not descended into.

Examples:

	Paths(User{}) => []string{"Base.ID", "Name"}""",
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil || t.Kind() != reflect.Struct {
	return nil
}
var out []string
for i := 0; i < t.NumField(); i++ {
	f := t.Field(i)
	if !f.IsExported() {
		continue
	}
	if f.Anonymous && f.Type.Kind() == reflect.Struct {
		for _, sub := range Paths(reflect.New(f.Type).Elem().Interface()) {
			out = append(out, f.Name+"."+sub)
		}
		continue
	}
	out = append(out, f.Name)
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type Base struct {
	ID     int
	hidden int
}

type Meta struct {
	Tag string
}

type User struct {
	Base
	Name  string
	Extra Meta
}

type Deep struct {
	User
	Note string
}

func TestPaths(t *testing.T) {
	want := []string{"Base.ID", "Name", "Extra"}
	if got := Paths(User{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Paths = %v, want %v", got, want)
	}
}

func TestPathsNested(t *testing.T) {
	want := []string{"User.Base.ID", "User.Name", "User.Extra", "Note"}
	if got := Paths(Deep{}); !reflect.DeepEqual(got, want) {
		t.Errorf("Paths = %v, want %v", got, want)
	}
}

func TestPathsSkipsUnexported(t *testing.T) {
	got := Paths(Base{})
	if !reflect.DeepEqual(got, []string{"ID"}) {
		t.Errorf("Paths = %v, want [ID]", got)
	}
}

func TestPathsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &User{}} {
		if got := Paths(in); got != nil {
			t.Errorf("Paths(%#v) = %v, want nil", in, got)
		}
	}
}

func TestPathsNamedStructFieldIsALeaf(t *testing.T) {
	got := Paths(User{})
	for _, p := range got {
		if p == "Extra.Tag" {
			t.Error("a named struct field must not be descended into")
		}
	}
}
""",
    context="A mapping layer walks struct fields to build column names. Embedding a shared `Base` struct silently produced a column called `Base` instead of `Base.ID`, and the migration ran anyway.",
    task=[
        "Return the dotted path of every exported leaf field, in declaration order.",
        "Descend through embedded structs, prefixing with the embedded type's name.",
        "Treat a named struct field as a leaf — do not descend into it.",
        "Skip unexported fields; return nil for a non-struct.",
    ],
    examples=[
        ("Paths(User{})", "[Base.ID Name Extra]", "`Base` is embedded; `Extra` is named."),
        ("Paths(Deep{})", "[User.Base.ID User.Name User.Extra Note]", "Embedding nests."),
        ("Paths(&User{})", "<nil>", None),
    ],
    topics=[
        ("StructField.Anonymous", "Marks an embedded field — the flag Go uses for promotion."),
        ("Embedding is not inheritance", "The field is still a field; only its name is implicit."),
        ("Recursion on the type", "`reflect.New(f.Type).Elem().Interface()` gives a zero value to recurse on."),
        ("Leaf choice is a design decision", "Descending into every struct would flatten too much."),
    ],
    hint="`f.Anonymous` is the only thing that distinguishes `Base` from `Extra`.",
    intuition="Embedding is a naming rule, not a new kind of field. Reflection shows it as an ordinary field with the `Anonymous` flag set — so \"descend or not\" is one boolean, and everything else is an ordinary walk.",
    approach=[
        "Reject non-structs.",
        "For each exported field: if it is anonymous and a struct, recurse and prefix with its name.",
        "Otherwise append the field's name as a leaf.",
    ],
    walkthrough="`user` embeds `Base`, so the walk recurses and prefixes, giving `Base.ID`. `Extra` is a named struct field, so it stops there.",
    pitfalls=[
        "Treating every struct field as embedded, which flattens `Extra` too.",
        "Using `FieldByName` for promoted fields, which finds them without telling you where they came from.",
    ],
)

P(
    "senior",
    name="alignedcopy",
    title="Copy Words Only When The Buffer Allows",
    sig="func CopyWords(dst []uint32, src []byte) (int, bool)",
    doc="""CopyWords copies as many whole uint32 values as fit from src into dst
and returns how many were copied.

It reports false when src's length is not a multiple of four or its start
is not 4-byte aligned; nothing is copied in that case.

Examples:

	CopyWords(make([]uint32, 2), eightBytes) => 2, true""",
    imports=['"unsafe"'],
    solution="""if len(src) == 0 || len(src)%4 != 0 {
	return 0, false
}
p := unsafe.Pointer(unsafe.SliceData(src))
if uintptr(p)&3 != 0 {
	return 0, false
}
view := unsafe.Slice((*uint32)(p), len(src)/4)
return copy(dst, view), true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func wordBytes(vals []uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(vals))), len(vals)*4)
}

func TestCopyWords(t *testing.T) {
	src := wordBytes([]uint32{1, 2, 3})
	dst := make([]uint32, 3)
	n, ok := CopyWords(dst, src)
	if !ok || n != 3 {
		t.Fatalf("CopyWords = %d, %v, want 3, true", n, ok)
	}
	if dst[0] != 1 || dst[1] != 2 || dst[2] != 3 {
		t.Errorf("dst = %v, want [1 2 3]", dst)
	}
}

func TestCopyWordsShortDst(t *testing.T) {
	src := wordBytes([]uint32{1, 2, 3})
	dst := make([]uint32, 2)
	n, ok := CopyWords(dst, src)
	if !ok || n != 2 {
		t.Errorf("CopyWords = %d, %v, want 2, true", n, ok)
	}
}

func TestCopyWordsIsACopy(t *testing.T) {
	vals := []uint32{1, 2}
	src := wordBytes(vals)
	dst := make([]uint32, 2)
	CopyWords(dst, src)
	vals[0] = 99
	if dst[0] != 1 {
		t.Error("dst aliases src; it must be a copy")
	}
}

func TestCopyWordsBadShapes(t *testing.T) {
	src := wordBytes([]uint32{1, 2, 3, 4})
	dst := make([]uint32, 4)
	for _, c := range []struct {
		name string
		in   []byte
	}{
		{"nil", nil},
		{"empty", src[:0]},
		{"length not a multiple of 4", src[:6]},
		{"misaligned", src[1:13]},
	} {
		n, ok := CopyWords(dst, c.in)
		if ok || n != 0 {
			t.Errorf("%s: CopyWords = %d, %v, want 0, false", c.name, n, ok)
		}
	}
}

func TestCopyWordsAllocatesNothing(t *testing.T) {
	src := wordBytes(make([]uint32, 1024))
	dst := make([]uint32, 1024)
	var n int
	if a := testing.AllocsPerRun(100, func() { n, _ = CopyWords(dst, src) }); a != 0 {
		t.Errorf("CopyWords made %v allocations, want 0", a)
	}
	_ = n
}
""",
    context="A decoder copies a received byte buffer into a word array with a loop and four shifts per word. On a same-endian link the shifts are pure overhead.",
    task=[
        "Copy as many whole `uint32` values as fit from `src` into `dst`.",
        "Return the count, and false when `src`'s length is not a multiple of four or its start is misaligned.",
        "`dst` must not alias `src`; allocate nothing.",
    ],
    examples=[
        ("CopyWords(make([]uint32,3), bytesOf(1,2,3))", "3, true", None),
        ("a 2-element dst", "2, true", "`copy` stops at the shorter side."),
        ("src[1:13]", "0, false", "Misaligned start."),
    ],
    topics=[
        ("Reinterpret then copy", "The view costs nothing; the copy is what makes `dst` independent."),
        ("copy bounds both sides", "It moves `min(len(dst), len(view))` elements."),
        ("Two preconditions", "Length divisibility and address alignment."),
        ("Not a wire format", "The interpretation is the host's byte order."),
    ],
    hint="Build the view, then let `copy` do the bounding.",
    intuition="Reinterpretation and copying compose: the view gives you typed elements without moving anything, and `copy` then moves exactly as many as both sides can hold.",
    approach=[
        "Reject an empty or non-multiple-of-four `src`, and a misaligned start.",
        "Build `unsafe.Slice((*uint32)(p), len(src)/4)`.",
        "Return `copy(dst, view), true`.",
    ],
    walkthrough="Twelve source bytes become a three-element view; a two-element `dst` receives two words and the count is 2.",
    pitfalls=[
        "Passing `len(src)` as the element count, which builds a view four times too long.",
        "Returning the view itself, which would alias the caller's bytes.",
    ],
)

P(
    "senior",
    name="encodeexact",
    title="Serialise With The Size You Can Compute",
    sig="func Encode(recs []Rec) []byte",
    doc="""Encode renders each record as "id:name" separated by '\\n'.

The output's length is determined by the input, so it should be
allocated once at exactly that size.

Examples:

	Encode([]Rec{{1, "a"}}) => []byte("1:a")""",
    imports=['"strconv"'],
    extra="""// Rec is one record to encode.
type Rec struct {
	ID   int
	Name string
}""",
    solution="""if len(recs) == 0 {
	return nil
}
n := len(recs) - 1
for _, r := range recs {
	n += len(strconv.Itoa(r.ID)) + 1 + len(r.Name)
}
out := make([]byte, 0, n)
for i, r := range recs {
	if i > 0 {
		out = append(out, '\\n')
	}
	out = strconv.AppendInt(out, int64(r.ID), 10)
	out = append(out, ':')
	out = append(out, r.Name...)
}
return out""",
    tests="""
import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

var sink []byte

func TestEncode(t *testing.T) {
	got := Encode([]Rec{{ID: 1, Name: "a"}, {ID: 22, Name: "bb"}})
	if !bytes.Equal(got, []byte("1:a\\n22:bb")) {
		t.Errorf("Encode = %q, want \\"1:a\\\\n22:bb\\"", got)
	}
}

func TestEncodeEmpty(t *testing.T) {
	if got := Encode(nil); len(got) != 0 {
		t.Errorf("Encode = %q, want empty", got)
	}
	if got := Encode([]Rec{}); len(got) != 0 {
		t.Errorf("Encode = %q, want empty", got)
	}
}

func TestEncodeEdges(t *testing.T) {
	got := Encode([]Rec{{ID: -5, Name: ""}})
	if !bytes.Equal(got, []byte("-5:")) {
		t.Errorf("Encode = %q, want \\"-5:\\"", got)
	}
}

func TestEncodeLarge(t *testing.T) {
	recs := make([]Rec, 500)
	var want strings.Builder
	for i := range recs {
		recs[i] = Rec{ID: i * 37, Name: strings.Repeat("n", i%5)}
		if i > 0 {
			want.WriteByte('\\n')
		}
		want.WriteString(strconv.Itoa(recs[i].ID))
		want.WriteByte(':')
		want.WriteString(recs[i].Name)
	}
	if got := Encode(recs); string(got) != want.String() {
		t.Errorf("Encode produced %d bytes, want %d", len(got), want.Len())
	}
}

func TestEncodeAllocatesOnce(t *testing.T) {
	recs := make([]Rec, 64)
	for i := range recs {
		recs[i] = Rec{ID: i, Name: "name"}
	}
	n := testing.AllocsPerRun(50, func() { sink = Encode(recs) })
	if n > 1 {
		t.Errorf("Encode made %v allocations, want 1: compute the length first", n)
	}
}
""",
    context="An export path builds its payload by appending to a nil slice. For a large batch that is a dozen reallocations and a full copy at each one.",
    task=[
        "Render each record as `id:name`, separated by `\\n`.",
        "Compute the exact output length first and allocate once.",
        "An empty input returns an empty result.",
    ],
    examples=[
        ('Encode([]Rec{{1,"a"},{22,"bb"}})', '"1:a\\n22:bb"', None),
        ('Encode([]Rec{{-5,""}})', '"-5:"', "The sign counts; an empty name is legal."),
        ("64 records", "1 allocation", None),
    ],
    topics=[
        ("Two passes beat doubling", "One cheap sizing pass removes every reallocation."),
        ("Digit counting", "`len(strconv.Itoa(id))` is the simplest correct width, sign included."),
        ("Separator arithmetic", "n records need n-1 separators."),
        ("strconv.Append*", "Writes digits straight into the buffer."),
    ],
    hint="The separators are part of the length too.",
    intuition="Every byte of the output is determined by the input, so the size is knowable before the first write. Once it is, the encoder is a straight line of appends into memory that never moves.",
    approach=[
        "Return nil for an empty input.",
        "Sum each record's digits, one colon, and its name, plus `len(recs)-1` separators.",
        "Allocate at that capacity and append the records.",
    ],
    walkthrough="Two records \"1:a\" and \"22:bb\" total 3 + 5 plus one separator — nine bytes, allocated once.",
    pitfalls=[
        "Forgetting the separators, so the last append reallocates.",
        "Counting digits with a hand-rolled loop that mishandles negative IDs or zero.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="shardedset",
    title="A Set That Many Goroutines Can Share",
    sig="func (s *Set) Add(key string) bool",
    doc="""Add inserts key and reports whether it was newly added.

The set is striped across shards so concurrent writers of different keys
rarely contend, and each shard is padded onto its own cache line.

Examples:

	s := NewSet(4); s.Add("a") => true, then false""",
    imports=['"hash/maphash"', '"sync"'],
    extra="""// bucket is one shard of the set, padded to a cache line.
type bucket struct {
	mu sync.Mutex
	m  map[string]struct{}
	_  [48]byte
}

// Set is a striped concurrent string set.
type Set struct {
	seed    maphash.Seed
	buckets []bucket
}

// NewSet returns a set with n shards.
func NewSet(n int) *Set {
	if n < 1 {
		n = 1
	}
	s := &Set{seed: maphash.MakeSeed(), buckets: make([]bucket, n)}
	for i := range s.buckets {
		s.buckets[i].m = make(map[string]struct{})
	}
	return s
}

// bucketFor returns the shard owning key.
func (s *Set) bucketFor(key string) *bucket {
	h := maphash.String(s.seed, key)
	return &s.buckets[h%uint64(len(s.buckets))]
}

// Has reports whether key is present.
func (s *Set) Has(key string) bool {
	b := s.bucketFor(key)
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.m[key]
	return ok
}

// Len reports the total number of keys.
func (s *Set) Len() int {
	n := 0
	for i := range s.buckets {
		b := &s.buckets[i]
		b.mu.Lock()
		n += len(b.m)
		b.mu.Unlock()
	}
	return n
}""",
    solution="""b := s.bucketFor(key)
b.mu.Lock()
defer b.mu.Unlock()
if _, ok := b.m[key]; ok {
	return false
}
b.m[key] = struct{}{}
return true""",
    tests="""
import (
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestAdd(t *testing.T) {
	s := NewSet(4)
	if !s.Add("a") {
		t.Error("the first Add reported false, want true")
	}
	if s.Add("a") {
		t.Error("the second Add reported true, want false")
	}
	if !s.Has("a") {
		t.Error("Has(a) = false, want true")
	}
	if s.Has("b") {
		t.Error("Has(b) = true, want false")
	}
	if s.Len() != 1 {
		t.Errorf("Len = %d, want 1", s.Len())
	}
}

func TestAddManyKeys(t *testing.T) {
	s := NewSet(8)
	for i := 0; i < 1000; i++ {
		if !s.Add(strconv.Itoa(i)) {
			t.Fatalf("Add(%d) reported false on a new key", i)
		}
	}
	if s.Len() != 1000 {
		t.Errorf("Len = %d, want 1000", s.Len())
	}
}

func TestSingleShard(t *testing.T) {
	s := NewSet(0)
	if !s.Add("a") || s.Add("a") {
		t.Error("a single-shard set must still deduplicate")
	}
}

func TestAddIsExactlyOnceUnderConcurrency(t *testing.T) {
	const (
		workers = 16
		keys    = 200
	)
	s := NewSet(16)
	var added atomic.Int64
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < keys; i++ {
				if s.Add(strconv.Itoa(i)) {
					added.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	if got := added.Load(); got != keys {
		t.Errorf("Add reported true %d times, want %d: the check and the insert must be atomic", got, keys)
	}
	if s.Len() != keys {
		t.Errorf("Len = %d, want %d", s.Len(), keys)
	}
}

func TestBucketsDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(bucket{}); got < 64 {
		t.Errorf("sizeof(bucket) = %d, want at least 64", got)
	}
}
""",
    context="A deduplicating crawler guards its seen-set with one mutex. At thirty-two workers the mutex is the bottleneck, and a first attempt at striping reported the same URL as new twice.",
    task=[
        "Insert `key` into its shard and report whether it was newly added.",
        "The presence check and the insert must be one atomic step.",
        "Hold only that shard's lock.",
        "Correct under concurrent use: a key added by many goroutines reports true exactly once.",
    ],
    examples=[
        ('s.Add("a") twice', "true, then false", None),
        ("16 workers adding 200 shared keys", "exactly 200 trues", None),
        ("NewSet(0)", "still deduplicates", None),
    ],
    topics=[
        ("Check-then-act must be atomic", "Releasing the lock between them lets two callers both win."),
        ("Lock striping", "Independent shards mean independent locks."),
        ("Deterministic routing", "The same key must always reach the same shard."),
        ("Padding the shard", "Neighbouring mutexes on one line contend in hardware."),
    ],
    hint="One lock, held across both the lookup and the insert.",
    intuition="Striping is about which lock, never about whether. The whole correctness of `Add` rests on the presence check and the insert happening without a gap — that is what makes \"newly added\" meaningful.",
    approach=[
        "Route to the shard.",
        "Lock it; return false if the key is present.",
        "Insert and return true.",
    ],
    walkthrough="Sixteen workers racing on one key all reach the same shard. One holds the lock, finds it absent and inserts; the rest find it present and report false.",
    pitfalls=[
        "Checking with `Has` and then calling a separate insert — two lock acquisitions with a race between them.",
        "Locking every shard, which is correct and no better than one mutex.",
        "Re-seeding the hash per call, which would scatter one key across shards.",
    ],
)

P(
    "staff",
    name="boundedqueue",
    title="Block The Producer, Do Not Buffer Forever",
    sig="func (q *Queue) Put(done <-chan struct{}, v int) bool",
    doc="""Put appends v to the queue, waiting while the queue is full, and
reports whether it was accepted.

Waiting is what applies backpressure; the queue must never grow past its
capacity, and a cancelled producer must not wait forever.

Examples:

	q := NewQueue(2); q.Put(done, 1) => true""",
    extra="""// Queue is a bounded FIFO with blocking producers.
type Queue struct {
	ch chan int
}

// NewQueue returns a queue holding at most n values.
func NewQueue(n int) *Queue {
	if n < 1 {
		n = 1
	}
	return &Queue{ch: make(chan int, n)}
}

// Take removes and returns the oldest value, waiting if the queue is empty.
func (q *Queue) Take(done <-chan struct{}) (int, bool) {
	select {
	case v := <-q.ch:
		return v, true
	case <-done:
		return 0, false
	}
}

// Len reports how many values are queued.
func (q *Queue) Len() int { return len(q.ch) }

// Cap reports the queue's capacity.
func (q *Queue) Cap() int { return cap(q.ch) }""",
    solution="""select {
case q.ch <- v:
	return true
case <-done:
	return false
}""",
    tests="""
import (
	"sync"
	"testing"
	"time"
)

func TestPutAndTake(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	q := NewQueue(2)
	if !q.Put(done, 1) || !q.Put(done, 2) {
		t.Fatal("the first two puts must succeed")
	}
	if v, ok := q.Take(done); !ok || v != 1 {
		t.Errorf("Take = %d, %v, want 1, true", v, ok)
	}
	if v, ok := q.Take(done); !ok || v != 2 {
		t.Errorf("Take = %d, %v, want 2, true: the queue must be FIFO", v, ok)
	}
}

func TestPutBlocksWhenFull(t *testing.T) {
	done := make(chan struct{})
	q := NewQueue(1)
	q.Put(done, 1)
	got := make(chan bool, 1)
	go func() { got <- q.Put(done, 2) }()
	select {
	case <-got:
		t.Fatal("Put returned while the queue was full")
	case <-time.After(50 * time.Millisecond):
	}
	close(done)
	if ok := <-got; ok {
		t.Error("the cancelled Put reported true, want false")
	}
}

func TestPutResumesAfterTake(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	q := NewQueue(1)
	q.Put(done, 1)
	got := make(chan bool, 1)
	go func() { got <- q.Put(done, 2) }()
	time.Sleep(20 * time.Millisecond)
	if v, _ := q.Take(done); v != 1 {
		t.Fatalf("Take = %d, want 1", v)
	}
	select {
	case ok := <-got:
		if !ok {
			t.Error("the waiting Put reported false, want true")
		}
	case <-time.After(2 * time.Second):
		t.Error("the waiting Put was never released")
	}
}

func TestQueueNeverExceedsCapacity(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	q := NewQueue(4)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if !q.Put(done, i) {
				return
			}
			if q.Len() > q.Cap() {
				panic("the queue grew past its capacity")
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if _, ok := q.Take(done); !ok {
				return
			}
		}
	}()
	wg.Wait()
}
""",
    context="A producer sends into an unbounded queue whenever the consumer lags. The lag is normal and the queue is the reason the process runs out of memory during the nightly batch.",
    task=[
        "Append `v`, waiting while the queue is full.",
        "Report false without appending when `done` closes first.",
        "The queue must never hold more than its capacity.",
    ],
    examples=[
        ("NewQueue(2), two Puts", "true, true", None),
        ("a third Put while full", "blocks, then false when done closes", None),
        ("a Take while a Put waits", "the Put proceeds", None),
    ],
    topics=[
        ("Backpressure by blocking", "A full buffer must slow the producer, not grow."),
        ("Buffered channel as the queue", "Its capacity is the bound and its blocking is the pressure."),
        ("Cancellable waiting", "`select` over the send and `done`."),
        ("Append nothing on cancellation", "Returning false must leave the queue untouched."),
    ],
    hint="A `select` with the send and the cancellation.",
    intuition="An unbounded queue converts a throughput mismatch into a memory leak. A bounded one converts it into a delay the producer can see and react to — which is what backpressure is.",
    approach=[
        "`select` on `q.ch <- v` and `<-done`.",
        "Return true for the send, false for the cancellation.",
    ],
    walkthrough="With a capacity of 4, the producer's fifth `Put` blocks until the consumer takes one. The length never exceeds 4, whatever the relative speeds.",
    pitfalls=[
        "Adding a `default` branch, which turns backpressure into silent data loss.",
        "Checking `len(q.ch) < cap(q.ch)` first — the state can change before the send.",
        "Growing the buffer when it fills, which is the original bug.",
    ],
)

P(
    "staff",
    name="typedclone",
    title="Clone A Map Without Reflection",
    sig="func CloneMap[K comparable, V any](m map[K]V) map[K]V",
    doc="""CloneMap returns a shallow copy of m with the same entries.

A type parameter keeps the keys and values concrete, so nothing is boxed
and the copy costs one allocation plus the entries.

Examples:

	CloneMap(map[string]int{"a": 1}) => a new map with the same entry""",
    solution="""if m == nil {
	return nil
}
out := make(map[K]V, len(m))
for k, v := range m {
	out[k] = v
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type point struct{ X, Y int }

func TestCloneMapContents(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := CloneMap(m)
	if !reflect.DeepEqual(got, m) {
		t.Errorf("CloneMap = %v, want %v", got, m)
	}
}

func TestCloneMapIsIndependent(t *testing.T) {
	m := map[string]int{"a": 1}
	got := CloneMap(m)
	got["b"] = 2
	if _, ok := m["b"]; ok {
		t.Error("the clone shares the original map")
	}
	m["c"] = 3
	if _, ok := got["c"]; ok {
		t.Error("the original shares the clone")
	}
}

func TestCloneMapNil(t *testing.T) {
	var m map[string]int
	if got := CloneMap(m); got != nil {
		t.Errorf("CloneMap(nil) = %v, want nil", got)
	}
}

func TestCloneMapEmpty(t *testing.T) {
	got := CloneMap(map[string]int{})
	if got == nil {
		t.Fatal("CloneMap of an empty map returned nil, want an empty map")
	}
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}

func TestCloneMapOtherTypes(t *testing.T) {
	m := map[int]point{1: {X: 1, Y: 2}}
	got := CloneMap(m)
	if got[1] != (point{X: 1, Y: 2}) {
		t.Errorf("got[1] = %v, want {1 2}", got[1])
	}
	got[1] = point{X: 9}
	if m[1].X != 1 {
		t.Error("struct values are shared")
	}
}

func TestCloneMapIsShallow(t *testing.T) {
	m := map[string][]int{"a": {1, 2}}
	got := CloneMap(m)
	got["a"][0] = 99
	if m["a"][0] != 99 {
		t.Error("the slice value was copied; a shallow clone shares it")
	}
}

func TestCloneMapAllocationsAreBounded(t *testing.T) {
	m := make(map[int]int, 64)
	for i := 0; i < 64; i++ {
		m[i] = i
	}
	var sink map[int]int
	if n := testing.AllocsPerRun(50, func() { sink = CloneMap(m) }); n > 6 {
		t.Errorf("CloneMap made %v allocations for 64 entries, want a handful: size the map up front", n)
	}
	_ = sink
}
""",
    context="A codebase has four hand-written map cloners, one per key-value pair it needed, and a fifth written with reflection that boxes every key and value.",
    task=[
        "Return a shallow copy of `m` with the same entries.",
        "A nil map clones to nil; an empty map clones to an empty, non-nil map.",
        "Size the new map up front; nothing may be boxed.",
    ],
    examples=[
        ('CloneMap(map[string]int{"a":1})', "a new map with the same entry", None),
        ("CloneMap(nil)", "<nil>", "nil is preserved, not turned into an empty map."),
        ('got["a"][0] = 99 on a map of slices', "the original sees it", "The clone is shallow."),
    ],
    topics=[
        ("Type parameters over reflection", "The compiler emits code for the concrete types; nothing is boxed."),
        ("comparable as the key constraint", "Exactly the types Go allows as map keys."),
        ("nil versus empty", "They behave differently and a clone must preserve which one it had."),
        ("Shallow by contract", "Reference values inside are shared, and that must be documented."),
    ],
    hint="Preserve nil, size the result, and copy.",
    intuition="Generics give one implementation the compiler specialises per instantiation — the same code a hand-written cloner would be, without the four copies or the boxing a reflective version needs.",
    approach=[
        "Return nil for a nil map.",
        "`make(map[K]V, len(m))`.",
        "Copy every entry and return.",
    ],
    walkthrough="Cloning a 64-entry map allocates one sized map; the entries are copied by value, so a struct value is independent and a slice value is shared.",
    pitfalls=[
        "Returning an empty map for a nil input, which changes `== nil` behaviour downstream.",
        "Omitting the size hint and rehashing on the way in.",
        "Documenting it as a deep copy; nested slices and maps are shared.",
    ],
)

P(
    "staff",
    name="arena",
    title="Carve Aligned Blocks From One Allocation",
    sig="func (a *Arena) Alloc(n int, align uintptr) ([]byte, bool)",
    doc="""Alloc returns the next n bytes of the arena, starting at an offset that
is a multiple of align.

The arena never grows: when the remaining space cannot satisfy the
request, Alloc reports false.

Examples:

	a := NewArena(64); a.Alloc(8, 8) => an 8-byte block, true""",
    extra="""// Arena hands out blocks of one fixed backing allocation.
type Arena struct {
	buf  []byte
	used uintptr
}

// NewArena returns an arena of size bytes.
func NewArena(size int) *Arena {
	if size < 0 {
		size = 0
	}
	return &Arena{buf: make([]byte, size)}
}

// Used reports how many bytes have been handed out, including padding.
func (a *Arena) Used() int { return int(a.used) }

// Cap reports the arena's total size.
func (a *Arena) Cap() int { return len(a.buf) }""",
    solution="""if n < 0 || align == 0 || align&(align-1) != 0 {
	return nil, false
}
start := (a.used + align - 1) &^ (align - 1)
if start > uintptr(len(a.buf)) || uintptr(n) > uintptr(len(a.buf))-start {
	return nil, false
}
a.used = start + uintptr(n)
return a.buf[start : start+uintptr(n) : start+uintptr(n)], true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestAllocSequential(t *testing.T) {
	a := NewArena(64)
	b1, ok := a.Alloc(8, 1)
	if !ok || len(b1) != 8 {
		t.Fatalf("Alloc = %d bytes, %v, want 8, true", len(b1), ok)
	}
	b2, ok := a.Alloc(4, 1)
	if !ok || len(b2) != 4 {
		t.Fatalf("Alloc = %d bytes, %v, want 4, true", len(b2), ok)
	}
	b1[0] = 1
	b2[0] = 2
	if b1[0] != 1 || b2[0] != 2 {
		t.Error("the blocks overlap")
	}
}

func TestAllocAligns(t *testing.T) {
	a := NewArena(64)
	a.Alloc(1, 1)
	b, ok := a.Alloc(8, 8)
	if !ok {
		t.Fatal("Alloc reported false")
	}
	if uintptr(unsafe.Pointer(&b[0]))%8 != 0 {
		t.Error("the block is not 8-byte aligned")
	}
	if a.Used() < 16 {
		t.Errorf("Used = %d, want at least 16: the padding counts", a.Used())
	}
}

func TestAllocCapacityIsExact(t *testing.T) {
	a := NewArena(64)
	b, _ := a.Alloc(8, 1)
	if cap(b) != 8 {
		t.Fatalf("cap = %d, want 8", cap(b))
	}
	b = append(b, 'x')
	c, _ := a.Alloc(8, 1)
	if c[0] == 'x' {
		t.Error("appending to one block wrote into the next")
	}
}

func TestAllocRefusesWhenFull(t *testing.T) {
	a := NewArena(16)
	if _, ok := a.Alloc(16, 1); !ok {
		t.Fatal("a request for the whole arena must succeed")
	}
	if _, ok := a.Alloc(1, 1); ok {
		t.Error("Alloc reported true for a full arena")
	}
}

func TestAllocRejectsBadArguments(t *testing.T) {
	a := NewArena(64)
	for _, c := range []struct {
		n     int
		align uintptr
	}{
		{-1, 1}, {8, 0}, {8, 3}, {8, 6}, {65, 1},
	} {
		if _, ok := a.Alloc(c.n, c.align); ok {
			t.Errorf("Alloc(%d, %d) reported true, want false", c.n, c.align)
		}
	}
	if a.Used() != 0 {
		t.Errorf("Used = %d, want 0: a rejected request must consume nothing", a.Used())
	}
}

func TestAllocZeroBytes(t *testing.T) {
	a := NewArena(8)
	b, ok := a.Alloc(0, 1)
	if !ok || len(b) != 0 {
		t.Errorf("Alloc(0) = %d bytes, %v, want 0, true", len(b), ok)
	}
}
""",
    context="A parser allocates thousands of small nodes per document. They all die together when the document is finished, which is exactly the shape a bump allocator serves best.",
    task=[
        "Return the next `n` bytes at an offset that is a multiple of `align`.",
        "Report false for a negative size, an `align` that is not a power of two, or a request that does not fit.",
        "A rejected request must consume nothing.",
        "Each block's capacity must equal its length, so an append cannot reach the next block.",
    ],
    examples=[
        ("NewArena(64).Alloc(8, 8)", "an 8-byte aligned block, true", None),
        ("Alloc(1,1) then Alloc(8,8)", "the second starts at offset 8", "Padding is consumed."),
        ("Alloc(8, 3)", "nil, false", "3 is not a power of two."),
    ],
    topics=[
        ("Bump allocation", "One pointer moves forward; there is no per-object bookkeeping."),
        ("Rounding up to an alignment", "`(x + a - 1) &^ (a - 1)` for a power-of-two `a`."),
        ("Three-index slicing", "Caps each block so appends cannot spill."),
        ("Bulk lifetime", "The whole arena is freed at once, or not at all."),
    ],
    hint="Round the cursor up first, then check whether the block fits without overflowing.",
    intuition="An arena trades per-object freeing for near-zero allocation cost: a rounded-up cursor and a bounds check. Everything it hands out lives exactly as long as the arena does.",
    approach=[
        "Validate `n` and the power-of-two `align`.",
        "Round `used` up to the alignment.",
        "Reject the request if the block does not fit, comparing without overflowing.",
        "Advance `used` and return the capped sub-slice.",
    ],
    walkthrough="After a one-byte block, the cursor is 1; a request for 8 bytes with alignment 8 rounds to 8, hands out bytes 8..15, and leaves `used` at 16 — the seven skipped bytes are gone.",
    pitfalls=[
        "`start + n > len(buf)` can overflow for a huge `n`; compare against the remaining space instead.",
        "Omitting the capacity cap, so an append to one block overwrites the next.",
        "Advancing `used` before deciding the request fits.",
    ],
)

P(
    "staff",
    name="batchedcounter",
    title="Accumulate Locally, Publish Rarely",
    sig="func (c *Counter) Add(local *Local, n int64)",
    doc="""Add adds n to the caller's local accumulator, flushing it into the
shared total when it reaches the batch threshold.

The shared atomic is the contended resource; touching it once per batch
instead of once per event is the whole point.

Examples:

	c.Add(local, 1) a thousand times => the total is 1000 after Flush""",
    imports=['"sync/atomic"'],
    extra="""// batchSize is how much a Local may accumulate before publishing.
const batchSize = 64

// Local is one goroutine's private accumulator. It must not be shared.
type Local struct {
	n int64
}

// Counter is a shared total fed by batched local accumulators.
type Counter struct {
	total atomic.Int64
}

// Flush publishes whatever the local still holds.
func (c *Counter) Flush(local *Local) {
	if local.n != 0 {
		c.total.Add(local.n)
		local.n = 0
	}
}

// Total returns the published total.
func (c *Counter) Total() int64 { return c.total.Load() }""",
    solution="""local.n += n
if local.n >= batchSize || local.n <= -batchSize {
	c.total.Add(local.n)
	local.n = 0
}""",
    tests="""
import (
	"sync"
	"testing"
)

func TestAddAccumulatesLocally(t *testing.T) {
	var c Counter
	var l Local
	for i := 0; i < batchSize-1; i++ {
		c.Add(&l, 1)
	}
	if got := c.Total(); got != 0 {
		t.Errorf("Total = %d, want 0: nothing should have been published yet", got)
	}
	c.Flush(&l)
	if got := c.Total(); got != batchSize-1 {
		t.Errorf("Total = %d, want %d", got, batchSize-1)
	}
}

func TestAddPublishesAtTheThreshold(t *testing.T) {
	var c Counter
	var l Local
	for i := 0; i < batchSize; i++ {
		c.Add(&l, 1)
	}
	if got := c.Total(); got != batchSize {
		t.Errorf("Total = %d, want %d: the batch should have been published", got, batchSize)
	}
	if l.n != 0 {
		t.Errorf("local = %d, want 0 after publishing", l.n)
	}
}

func TestAddHandlesLargeIncrements(t *testing.T) {
	var c Counter
	var l Local
	c.Add(&l, 1000)
	if got := c.Total(); got != 1000 {
		t.Errorf("Total = %d, want 1000", got)
	}
}

func TestAddHandlesNegatives(t *testing.T) {
	var c Counter
	var l Local
	c.Add(&l, -1000)
	c.Flush(&l)
	if got := c.Total(); got != -1000 {
		t.Errorf("Total = %d, want -1000", got)
	}
}

func TestNoIncrementIsLost(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	const (
		workers = 16
		each    = 1000
	)
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			var l Local
			for i := 0; i < each; i++ {
				c.Add(&l, 1)
			}
			c.Flush(&l)
		}()
	}
	wg.Wait()
	if got := c.Total(); got != workers*each {
		t.Errorf("Total = %d, want %d", got, workers*each)
	}
}

func BenchmarkAdd(b *testing.B) {
	var c Counter
	b.RunParallel(func(pb *testing.PB) {
		var l Local
		for pb.Next() {
			c.Add(&l, 1)
		}
		c.Flush(&l)
	})
}
""",
    context="A request counter uses one atomic increment per event. At several million events a second the atomic's cache line bounces between every core in the machine.",
    task=[
        "Add `n` to the caller's local accumulator.",
        "Publish into the shared total when the local reaches `batchSize` in either direction, then reset it.",
        "No increment may be lost: `Flush` plus the batched publishes must account for everything.",
        "Correct under concurrent use, with one `Local` per goroutine.",
    ],
    examples=[
        ("63 Adds of 1, then Total()", "0", "Still local."),
        ("64 Adds of 1", "the total is 64", "The batch published."),
        ("16 workers x 1000 Adds, each flushed", "16000", None),
    ],
    topics=[
        ("Contention scales with sharing", "Batching cuts atomic traffic by the batch size."),
        ("Private state needs no synchronisation", "The `Local` belongs to one goroutine."),
        ("Threshold in both directions", "A negative accumulator must publish too."),
        ("Flush is part of the contract", "A partial batch is only counted when the caller flushes."),
    ],
    hint="Add to the local, then decide whether it is time to publish.",
    intuition="An atomic increment is cheap in isolation and ruinous under contention, because every core wants the same cache line. Accumulating privately turns sixty-four coherence events into one.",
    approach=[
        "Add `n` to `local.n`.",
        "If the magnitude has reached `batchSize`, `total.Add(local.n)` and zero the local.",
    ],
    walkthrough="Sixteen workers each accumulate to 64 before touching the shared counter, so a million increments cost about sixteen thousand atomic operations instead of a million.",
    pitfalls=[
        "Checking only `>= batchSize`, so a decreasing counter never publishes.",
        "Sharing one `Local` between goroutines, which is a data race the type's documentation forbids.",
        "Forgetting `Flush`, which loses up to `batchSize-1` per goroutine.",
    ],
)
