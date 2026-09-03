"""10-advanced-topics — rotation 5: 5 puzzles per level."""

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
    name="capgrew",
    title="Did That Append Reallocate",
    sig="func Appended(s []int, v int) ([]int, bool)",
    doc="""Appended appends v to s and reports whether the append had to grow the
capacity.

Growing means a new array was allocated and the old contents copied.

Examples:

	Appended(make([]int, 0, 4), 1) => a one-element slice, false""",
    solution="""before := cap(s)
out := append(s, v)
return out, cap(out) != before""",
    tests="""
import "testing"

func TestAppendedValue(t *testing.T) {
	got, _ := Appended([]int{1, 2}, 3)
	if len(got) != 3 || got[2] != 3 {
		t.Errorf("Appended = %v, want [1 2 3]", got)
	}
}

func TestAppendedWithRoom(t *testing.T) {
	s := make([]int, 0, 4)
	got, grew := Appended(s, 1)
	if grew {
		t.Error("grew = true, want false: the capacity was sufficient")
	}
	if len(got) != 1 || cap(got) != 4 {
		t.Errorf("len, cap = %d, %d, want 1, 4", len(got), cap(got))
	}
}

func TestAppendedWithoutRoom(t *testing.T) {
	s := make([]int, 1, 1)
	got, grew := Appended(s, 2)
	if !grew {
		t.Error("grew = false, want true: the capacity was exhausted")
	}
	if cap(got) <= 1 {
		t.Errorf("cap = %d, want more than 1", cap(got))
	}
}

func TestAppendedFromNil(t *testing.T) {
	got, grew := Appended(nil, 5)
	if !grew {
		t.Error("grew = false, want true: a nil slice has no capacity")
	}
	if len(got) != 1 || got[0] != 5 {
		t.Errorf("Appended = %v, want [5]", got)
	}
}
""",
    context="A reviewer claims a hot loop never reallocates because it appends to a pre-sized slice. Nobody has actually measured it.",
    task=[
        "Append `v` to `s` and return the result.",
        "Report whether the capacity changed — that is, whether a new array was allocated.",
    ],
    examples=[
        ("Appended(make([]int,0,4), 1)", "[1], false", "The room was already there."),
        ("Appended(make([]int,1,1), 2)", "[.. 2], true", "The capacity was exhausted."),
        ("Appended(nil, 5)", "[5], true", None),
    ],
    topics=[
        ("Capacity is the growth signal", "`append` reallocates exactly when the length would exceed the capacity."),
        ("Length always changes", "Only the capacity distinguishes a copy from an in-place write."),
        ("Reading cap before and after", "The comparison is the whole measurement."),
    ],
    hint="Record something before the append and compare it afterwards.",
    intuition="Whether `append` copied is invisible in the result's contents — the length changes either way. The capacity is the observable that separates \"wrote into spare room\" from \"allocated a bigger array\".",
    approach=[
        "Save `cap(s)`.",
        "Append.",
        "Return the result and whether the capacity differs.",
    ],
    walkthrough="A slice with len 0 and cap 4 has room, so the capacity stays 4 and `grew` is false. A slice with len 1 and cap 1 has none, so `append` allocates a larger array and the capacity changes.",
    pitfalls=[
        "Comparing lengths, which always differ by one.",
        "Comparing `cap(s)` after the append with itself — `s` still holds the old header.",
    ],
)

P(
    "junior",
    name="bytesandrunes",
    title="Bytes Are Not Characters",
    sig="func Counts(s string) (bytes, runes int)",
    doc="""Counts returns how many bytes and how many characters s holds.

len gives bytes; ranging a string yields characters.

Examples:

	Counts("héllo") => 6, 5""",
    solution="""bytes = len(s)
for range s {
	runes++
}
return bytes, runes""",
    tests="""
import "testing"

var sinkA, sinkB int

func TestCounts(t *testing.T) {
	cases := []struct {
		in           string
		bytes, runes int
	}{
		{"hello", 5, 5},
		{"héllo", 6, 5},
		{"日本語", 9, 3},
		{"", 0, 0},
		{"a\\u00e9\\u65e5", 6, 3},
	}
	for _, c := range cases {
		b, r := Counts(c.in)
		if b != c.bytes || r != c.runes {
			t.Errorf("Counts(%q) = %d, %d, want %d, %d", c.in, b, r, c.bytes, c.runes)
		}
	}
}

func TestCountsASCIIAgree(t *testing.T) {
	b, r := Counts("plain ascii")
	if b != r {
		t.Errorf("bytes = %d, runes = %d, want them equal for ASCII", b, r)
	}
}

func TestCountsAllocatesNothing(t *testing.T) {
	s := "a string with sömé nön-ascii characters in it"
	if n := testing.AllocsPerRun(200, func() { sinkA, sinkB = Counts(s) }); n != 0 {
		t.Errorf("Counts made %v allocations, want 0: do not convert to []rune", n)
	}
}
""",
    context="A field is limited to twenty characters. The check uses `len(s)`, and users with accented names are rejected for a name that fits.",
    task=[
        "Return the byte length and the character count of `s`.",
        "Zero allocations — do not convert to `[]rune`.",
    ],
    examples=[
        ('Counts("hello")', "5, 5", "ASCII: one byte per character."),
        ('Counts("héllo")', "6, 5", "é is two bytes."),
        ('Counts("日本語")', "9, 3", "Three bytes each."),
    ],
    topics=[
        ("len is bytes", "A string's length is its storage size, not its character count."),
        ("range decodes UTF-8", "Ranging a string yields one iteration per character."),
        ("[]rune(s) allocates", "It builds a whole new slice of 32-bit values."),
    ],
    hint="`range` over the string, and count the iterations.",
    intuition="Go strings are UTF-8 bytes. `len` answers a storage question and `range` answers a text question — they agree only for ASCII, which is why the bug survives every English test case.",
    approach=[
        "`bytes = len(s)`.",
        "Range over `s`, incrementing a counter.",
    ],
    walkthrough='"héllo" is six bytes: h, two for é, then l, l, o. Ranging it yields five iterations, one per character.',
    pitfalls=[
        "`utf8.RuneCountInString` is the real answer and is also allocation-free — the point here is why `len` differs.",
        "`for i := range s` gives byte offsets, not consecutive indices.",
    ],
)

P(
    "junior",
    name="typename",
    title="What Type Is In This Interface",
    sig="func TypeName(v any) string",
    doc="""TypeName returns the name of v's dynamic type.

A nil interface holds no type at all, so it reports "<nil>".

Examples:

	TypeName(3) => "int" """,
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil {
	return "<nil>"
}
return t.String()""",
    tests="""
import "testing"

type widget struct{ A int }

func TestTypeName(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{3, "int"},
		{"s", "string"},
		{3.5, "float64"},
		{[]int{1}, "[]int"},
		{map[string]int{}, "map[string]int"},
		{widget{}, "typename.widget"},
		{&widget{}, "*typename.widget"},
		{nil, "<nil>"},
		{[]byte("x"), "[]uint8"},
	}
	for _, c := range cases {
		if got := TypeName(c.in); got != c.want {
			t.Errorf("TypeName(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestTypeNameNamedTypes(t *testing.T) {
	type alias = int
	type named int
	if got := TypeName(alias(1)); got != "int" {
		t.Errorf("TypeName = %q, want \\"int\\": an alias is not a new type", got)
	}
	if got := TypeName(named(1)); got != "typename.named" {
		t.Errorf("TypeName = %q, want \\"typename.named\\"", got)
	}
}
""",
    context="A debug handler logs what it was given. `%T` works in a format string, and the handler needs the name as a value it can compare and store.",
    task=[
        "Return the name of `v`'s dynamic type.",
        "A nil interface reports `\"<nil>\"`.",
        "A named type reports its own name, not its underlying kind.",
    ],
    examples=[
        ("TypeName(3)", '"int"', None),
        ("TypeName(&widget{})", '"*typename.widget"', "The pointer is part of the type."),
        ("TypeName(nil)", '"<nil>"', None),
    ],
    topics=[
        ("reflect.TypeOf", "Returns nil for a nil interface, and a Type otherwise."),
        ("Type.String vs Kind", "`String` names the type; `Kind` names its shape."),
        ("Aliases are not types", "`type alias = int` is another spelling of int."),
    ],
    hint="`TypeOf(nil)` is nil, and calling a method on it panics.",
    intuition="An interface value carries a type word. Reflection reads it — and when the interface is nil there is nothing to read, which is why the nil check comes first.",
    approach=[
        "`reflect.TypeOf(v)`.",
        "Return `\"<nil>\"` when it is nil.",
        "Return `t.String()`.",
    ],
    walkthrough="`TypeName(named(1))` reports the declared name qualified by its package, while `TypeName(alias(1))` reports \"int\" because an alias declares no new type.",
    pitfalls=[
        "Calling `t.String()` without the nil check — that is a nil-pointer panic.",
        "Expecting `[]byte` to print as \"[]byte\"; `byte` is an alias for `uint8`.",
    ],
)

P(
    "junior",
    name="alignments",
    title="What Each Type Must Line Up On",
    sig="func Alignments() (b, i32, i64, s uintptr)",
    doc="""Alignments returns the alignment requirement of a byte, an int32, an
int64 and a string.

A type's alignment is the boundary its address must be a multiple of.

Examples:

	Alignments() => 1, 4, 8, 8 on a 64-bit build""",
    imports=['"unsafe"'],
    solution="""var (
	vb  byte
	v32 int32
	v64 int64
	vs  string
)
return unsafe.Alignof(vb), unsafe.Alignof(v32), unsafe.Alignof(v64), unsafe.Alignof(vs)""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestAlignments(t *testing.T) {
	var (
		vb  byte
		v32 int32
		v64 int64
		vs  string
	)
	b, i32, i64, s := Alignments()
	if b != unsafe.Alignof(vb) {
		t.Errorf("byte = %d, want %d", b, unsafe.Alignof(vb))
	}
	if i32 != unsafe.Alignof(v32) {
		t.Errorf("int32 = %d, want %d", i32, unsafe.Alignof(v32))
	}
	if i64 != unsafe.Alignof(v64) {
		t.Errorf("int64 = %d, want %d", i64, unsafe.Alignof(v64))
	}
	if s != unsafe.Alignof(vs) {
		t.Errorf("string = %d, want %d", s, unsafe.Alignof(vs))
	}
}

func TestAlignmentsAreAscending(t *testing.T) {
	b, i32, i64, _ := Alignments()
	if b != 1 {
		t.Errorf("byte alignment = %d, want 1", b)
	}
	if i32 < b || i64 < i32 {
		t.Errorf("alignments %d, %d, %d are not ascending", b, i32, i64)
	}
}

func TestAlignmentsArePowersOfTwo(t *testing.T) {
	for _, a := range []uintptr{mustAlign(0), mustAlign(1), mustAlign(2), mustAlign(3)} {
		if a == 0 || a&(a-1) != 0 {
			t.Errorf("alignment %d is not a power of two", a)
		}
	}
}

func mustAlign(i int) uintptr {
	b, i32, i64, s := Alignments()
	return []uintptr{b, i32, i64, s}[i]
}
""",
    context="A struct is laid out by hand for a binary format. Guessing where each field lands works until the first field wider than a byte.",
    task=[
        "Return the alignment of `byte`, `int32`, `int64` and `string`.",
        "Derive them with `unsafe.Alignof` rather than writing numbers.",
    ],
    examples=[
        ("Alignments()", "1, 4, 8, 8", "On a 64-bit build."),
        ("byte's alignment", "1", "A byte can sit anywhere."),
        ("every result", "a power of two", None),
    ],
    topics=[
        ("unsafe.Alignof", "The boundary a type's address must be a multiple of."),
        ("Alignment drives padding", "A field's offset is rounded up to its alignment."),
        ("Composite alignment", "A struct's alignment is its widest field's."),
    ],
    hint="`Alignof` takes a value, so declare one of each.",
    intuition="Alignment is a hardware requirement the compiler enforces: a wide load wants an address divisible by its width. Every layout question — offsets, padding, struct size — follows from these numbers.",
    approach=[
        "Declare a variable of each type.",
        "Return `unsafe.Alignof` of each.",
    ],
    walkthrough="A string is a pointer plus a length, so its alignment is the pointer's — 8 on a 64-bit build, even though it occupies 16 bytes.",
    pitfalls=[
        "Confusing `Alignof` with `Sizeof`; a string is 16 bytes and 8-aligned.",
        "Hard-coding the numbers, which are architecture-specific.",
    ],
)

P(
    "junior",
    name="derefsafe",
    title="Read Through A Pointer That May Be Nil",
    sig="func Value(p *int) int",
    doc="""Value returns what p points at, or 0 when p is nil.

A nil pointer is a valid value; dereferencing one is not.

Examples:

	Value(nil) => 0""",
    solution="""if p == nil {
	return 0
}
return *p""",
    tests="""
import "testing"

var sink int

func TestValue(t *testing.T) {
	n := 42
	if got := Value(&n); got != 42 {
		t.Errorf("Value = %d, want 42", got)
	}
	if got := Value(nil); got != 0 {
		t.Errorf("Value(nil) = %d, want 0", got)
	}
	zero := 0
	if got := Value(&zero); got != 0 {
		t.Errorf("Value = %d, want 0", got)
	}
}

func TestValueSeesLaterWrites(t *testing.T) {
	n := 1
	p := &n
	n = 7
	if got := Value(p); got != 7 {
		t.Errorf("Value = %d, want 7: the pointer is a live view", got)
	}
}

func TestValueAllocatesNothing(t *testing.T) {
	n := 3
	p := &n
	if got := testing.AllocsPerRun(200, func() { sink = Value(p) }); got != 0 {
		t.Errorf("Value made %v allocations, want 0", got)
	}
}
""",
    context="An optional field is modelled as `*int`. Every read site dereferences it, and the one that forgot the nil check is the one that runs in production.",
    task=[
        "Return what `p` points at.",
        "Return 0 when `p` is nil.",
        "Zero allocations.",
    ],
    examples=[
        ("n := 42; Value(&n)", "42", None),
        ("Value(nil)", "0", "No panic."),
        ("n changed after taking &n", "Value sees the new value", None),
    ],
    topics=[
        ("nil is a valid pointer value", "Holding one is fine; dereferencing one is a panic."),
        ("Pointers are live views", "They read the variable, not a snapshot of it."),
        ("The zero value as a default", "0 for a missing int is a decision, not an accident."),
    ],
    hint="One comparison, then the dereference.",
    intuition="A pointer type has one more state than its base type: absent. The dereference is only defined for the other states, so the check is not defensive noise — it is the missing case.",
    approach=[
        "Return 0 when `p` is nil.",
        "Otherwise return `*p`.",
    ],
    walkthrough="`Value(&n)` reads the current contents of `n`, so a write to `n` between taking the address and calling is visible. `Value(nil)` returns the documented default instead of faulting.",
    pitfalls=[
        "Returning `*p` after a `p != nil` check written the wrong way round.",
        "Treating 0 as \"absent\" at the call site — `&zero` is present and zero.",
    ],
)

# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="copyinto",
    title="Copy Stops At The Shorter Side",
    sig="func CopyInto(dst, src []int) int",
    doc="""CopyInto copies as many elements as fit from src into dst and returns
how many were copied.

Neither slice is resized: the copy is bounded by the shorter of the two.

Examples:

	CopyInto(make([]int, 2), []int{1, 2, 3}) => 2""",
    solution="""return copy(dst, src)""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestCopyIntoShortDst(t *testing.T) {
	dst := make([]int, 2)
	if got := CopyInto(dst, []int{1, 2, 3}); got != 2 {
		t.Errorf("CopyInto = %d, want 2", got)
	}
	if !reflect.DeepEqual(dst, []int{1, 2}) {
		t.Errorf("dst = %v, want [1 2]", dst)
	}
}

func TestCopyIntoShortSrc(t *testing.T) {
	dst := []int{9, 9, 9}
	if got := CopyInto(dst, []int{1}); got != 1 {
		t.Errorf("CopyInto = %d, want 1", got)
	}
	if !reflect.DeepEqual(dst, []int{1, 9, 9}) {
		t.Errorf("dst = %v, want [1 9 9]", dst)
	}
}

func TestCopyIntoEmpty(t *testing.T) {
	if got := CopyInto(nil, []int{1}); got != 0 {
		t.Errorf("CopyInto = %d, want 0", got)
	}
	if got := CopyInto(make([]int, 3), nil); got != 0 {
		t.Errorf("CopyInto = %d, want 0", got)
	}
	dst := make([]int, 0, 8)
	if got := CopyInto(dst, []int{1, 2}); got != 0 {
		t.Errorf("CopyInto = %d, want 0: capacity is not length", got)
	}
}

func TestCopyIntoOverlapping(t *testing.T) {
	s := []int{1, 2, 3, 4}
	if got := CopyInto(s, s[1:]); got != 3 {
		t.Errorf("CopyInto = %d, want 3", got)
	}
	if !reflect.DeepEqual(s, []int{2, 3, 4, 4}) {
		t.Errorf("s = %v, want [2 3 4 4]", s)
	}
}

func TestCopyIntoAllocatesNothing(t *testing.T) {
	dst := make([]int, 128)
	src := make([]int, 256)
	if n := testing.AllocsPerRun(100, func() { _ = CopyInto(dst, src) }); n != 0 {
		t.Errorf("CopyInto made %v allocations, want 0", n)
	}
}
""",
    context="A fixed-size window is filled from a larger stream. The first attempt indexed `src` up to `len(src)` and wrote past the end of `dst`.",
    task=[
        "Copy as many elements as fit from `src` into `dst`.",
        "Return the number copied.",
        "Nothing is resized, and nothing is allocated.",
    ],
    examples=[
        ("CopyInto(make([]int,2), []int{1,2,3})", "2", "dst is the limit."),
        ("CopyInto([]int{9,9,9}, []int{1})", "1", "src is the limit; the rest of dst is untouched."),
        ("CopyInto(make([]int,0,8), []int{1,2})", "0", "Capacity is not length."),
    ],
    topics=[
        ("copy is bounded by both", "It moves `min(len(dst), len(src))` elements."),
        ("Length, not capacity", "`copy` writes only what `dst`'s length allows."),
        ("Overlap is well defined", "`copy` behaves like memmove."),
    ],
    hint="The builtin already does exactly this, including the return value.",
    intuition="`copy` is the one builtin that refuses to write past either side. That makes \"fill what fits\" a single call, with the count falling out for free.",
    approach=[
        "Return `copy(dst, src)`.",
    ],
    walkthrough="A destination of length 2 and a source of length 3 copies 2. A destination of length 0 and capacity 8 copies nothing, because capacity is not writable through `copy`.",
    pitfalls=[
        "Looping with `len(src)` as the bound, which overruns a shorter `dst`.",
        "Expecting a `dst` with spare capacity to receive more than its length.",
    ],
)

P(
    "middle",
    name="rangecopy",
    title="The Loop Variable Is A Copy",
    mode="bug",
    sig="func Bump(items []Counter)",
    doc="""Bump increments every counter in items, in place.

Ranging by value copies each element; the increment has to reach the
slice's own storage.

Examples:

	items := []Counter{{N: 1}}; Bump(items) => items[0].N == 2""",
    extra="""// Counter is one element of the slice.
type Counter struct {
	N   int
	Pad [64]byte
}""",
    buggy="""for _, c := range items {
	c.N++
}""",
    solution="""for i := range items {
	items[i].N++
}""",
    tests="""
import "testing"

func TestBump(t *testing.T) {
	items := []Counter{{N: 1}, {N: 2}, {N: 3}}
	Bump(items)
	for i, c := range items {
		if c.N != i+2 {
			t.Errorf("items[%d].N = %d, want %d: the loop wrote to a copy", i, c.N, i+2)
		}
	}
}

func TestBumpEmpty(t *testing.T) {
	Bump(nil)
	Bump([]Counter{})
}

func TestBumpWritesThroughAView(t *testing.T) {
	items := []Counter{{N: 0}, {N: 0}, {N: 0}}
	Bump(items[1:2])
	if items[0].N != 0 || items[2].N != 0 {
		t.Errorf("items = %v, want only the middle element bumped", []int{items[0].N, items[1].N, items[2].N})
	}
	if items[1].N != 1 {
		t.Errorf("items[1].N = %d, want 1", items[1].N)
	}
}

func TestBumpAllocatesNothing(t *testing.T) {
	items := make([]Counter, 64)
	if n := testing.AllocsPerRun(100, func() { Bump(items) }); n != 0 {
		t.Errorf("Bump made %v allocations, want 0", n)
	}
}
""",
    context="A stats pass increments a counter in every element of a slice. It runs without error and the counters never move.",
    task=[
        "Increment `N` in every element of `items`, in place.",
        "The writes must be visible through the caller's slice.",
        "Fix the single bug; allocate nothing.",
    ],
    examples=[
        ("items := []Counter{{N:1}}; Bump(items)", "items[0].N is 2", None),
        ("Bump(items[1:2])", "only the middle element changes", None),
        ("Bump(nil)", "no panic", None),
    ],
    topics=[
        ("range copies the element", "The value variable is a fresh copy each iteration."),
        ("Index to write", "`items[i]` addresses the slice's own storage."),
        ("The copy can be expensive", "A 72-byte element is copied per iteration for nothing."),
    ],
    hint="Which variable is the slice's element, and which is a copy of it?",
    intuition="`for _, c := range s` binds `c` to a copy. Writing to it updates the copy and the copy is discarded — which is why the loop is silently a no-op rather than a compile error.",
    approach=[
        "Range over the indices only.",
        "Write through `items[i]`.",
    ],
    walkthrough="The buggy loop copies 72 bytes per element, increments the copy, and drops it. Indexing writes into the caller's array directly and copies nothing.",
    pitfalls=[
        "`for i, c := range items { items[i] = c }` after mutating `c` — correct, and it copies twice.",
        "Assuming a slice of pointers behaves the same; there the copy is a pointer you can write through.",
    ],
)

P(
    "middle",
    name="nonzerofields",
    title="Which Fields Were Actually Set",
    sig="func NonZero(v any) []string",
    doc="""NonZero returns the names of v's exported fields that hold something
other than their zero value, in declaration order.

A non-struct, or a nil interface, yields nil.

Examples:

	NonZero(patch{Name: "x"}) => []string{"Name"}""",
    imports=['"reflect"'],
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() || rv.Kind() != reflect.Struct {
	return nil
}
rt := rv.Type()
var out []string
for i := 0; i < rv.NumField(); i++ {
	if !rt.Field(i).IsExported() {
		continue
	}
	if !rv.Field(i).IsZero() {
		out = append(out, rt.Field(i).Name)
	}
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type patch struct {
	Name   string
	Count  int
	Active bool
	Tags   []string
	hidden int
}

func TestNonZero(t *testing.T) {
	got := NonZero(patch{Name: "x", Count: 3})
	if !reflect.DeepEqual(got, []string{"Name", "Count"}) {
		t.Errorf("NonZero = %v, want [Name Count]", got)
	}
}

func TestNonZeroAllZero(t *testing.T) {
	if got := NonZero(patch{}); got != nil {
		t.Errorf("NonZero = %v, want nil", got)
	}
}

func TestNonZeroSkipsUnexported(t *testing.T) {
	if got := NonZero(patch{hidden: 9}); got != nil {
		t.Errorf("NonZero = %v, want nil: unexported fields do not count", got)
	}
}

func TestNonZeroEmptySliceCounts(t *testing.T) {
	got := NonZero(patch{Tags: []string{}})
	if !reflect.DeepEqual(got, []string{"Tags"}) {
		t.Errorf("NonZero = %v, want [Tags]: an empty non-nil slice is not the zero value", got)
	}
}

func TestNonZeroFalseIsZero(t *testing.T) {
	if got := NonZero(patch{Active: false}); got != nil {
		t.Errorf("NonZero = %v, want nil", got)
	}
	if got := NonZero(patch{Active: true}); !reflect.DeepEqual(got, []string{"Active"}) {
		t.Errorf("NonZero = %v, want [Active]", got)
	}
}

func TestNonZeroNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, &patch{}} {
		if got := NonZero(in); got != nil {
			t.Errorf("NonZero(%#v) = %v, want nil", in, got)
		}
	}
}
""",
    context="A PATCH endpoint applies only the fields the caller sent. The struct cannot tell \"omitted\" from \"set to zero\" — but listing the non-zero ones is what the handler actually needs.",
    task=[
        "Return the names of exported fields whose value is not the zero value.",
        "Preserve declaration order; skip unexported fields.",
        "Return nil for a non-struct or a nil interface.",
    ],
    examples=[
        ('NonZero(patch{Name:"x", Count:3})', "[Name Count]", None),
        ("NonZero(patch{})", "<nil>", "Everything is zero."),
        ("NonZero(patch{Tags: []string{}})", "[Tags]", "An empty non-nil slice is not the zero slice."),
    ],
    topics=[
        ("Value.IsZero", "Compares against the type's zero value, whatever the field's type."),
        ("Type and Value in step", "Names come from the Type, values from the Value, at the same index."),
        ("nil vs empty", "A nil slice is zero; an allocated empty one is not."),
    ],
    hint="`rv.Field(i)` for the value, `rt.Field(i)` for the name.",
    intuition="Reflection lets one function answer \"what changed\" for every struct in the codebase, and `IsZero` handles each field's notion of empty without a type switch.",
    approach=[
        "Reject non-structs.",
        "Loop the fields, skipping unexported ones.",
        "Append the name when `rv.Field(i).IsZero()` is false.",
    ],
    walkthrough="`patch{Name:\"x\", Count:3}` has two non-zero fields; `Active` is false and `Tags` is nil, both of which are their types' zero values.",
    pitfalls=[
        "Comparing against a fresh zero struct field by field, which needs the type at compile time.",
        "Treating an empty slice as absent — `IsZero` deliberately does not.",
    ],
)

P(
    "middle",
    name="valuebytes",
    title="A Byte View Of One Value",
    sig="func Bytes(p *uint64) []byte",
    doc="""Bytes returns an 8-byte view of the uint64 p points at, sharing its
storage.

A nil pointer yields nil. The view is the machine's layout, so it is not
a portable encoding.

Examples:

	v := uint64(1); Bytes(&v) => 8 bytes sharing v""",
    imports=['"unsafe"'],
    solution="""if p == nil {
	return nil
}
return unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p))""",
    tests="""
import (
	"testing"
	"unsafe"
)

var sink []byte

func TestBytesShape(t *testing.T) {
	v := uint64(0)
	b := Bytes(&v)
	if len(b) != 8 {
		t.Errorf("len = %d, want 8", len(b))
	}
	if cap(b) != 8 {
		t.Errorf("cap = %d, want 8: an append must not run past the value", cap(b))
	}
}

func TestBytesSharesStorage(t *testing.T) {
	v := uint64(0)
	b := Bytes(&v)
	v = ^uint64(0)
	for i, x := range b {
		if x != 0xff {
			t.Fatalf("b[%d] = %#x, want 0xff: the view does not share v", i, x)
		}
	}
	b[0] = 0
	if v == ^uint64(0) {
		t.Error("writing through the view did not change v")
	}
}

func TestBytesNil(t *testing.T) {
	if got := Bytes(nil); got != nil {
		t.Errorf("Bytes(nil) = %v, want nil", got)
	}
}

func TestBytesAllocatesNothing(t *testing.T) {
	v := uint64(7)
	if n := testing.AllocsPerRun(200, func() { sink = Bytes(&v) }); n != 0 {
		t.Errorf("Bytes made %v allocations, want 0", n)
	}
}

func TestBytesMatchesTheSize(t *testing.T) {
	v := uint64(0)
	if uintptr(len(Bytes(&v))) != unsafe.Sizeof(v) {
		t.Error("the view's length must come from unsafe.Sizeof")
	}
}
""",
    context="A checksum routine takes `[]byte` and the caller has a counter. Encoding the counter into a temporary buffer allocates on every update.",
    task=[
        "Return an 8-byte view of the value `p` points at, sharing its storage.",
        "A nil pointer yields nil.",
        "The view's length and capacity must both come from `unsafe.Sizeof` — zero allocations.",
    ],
    examples=[
        ("v := uint64(0); b := Bytes(&v); v = ^uint64(0)", "b reads all 0xff", "The view is live."),
        ("Bytes(nil)", "<nil>", None),
        ("cap of the result", "8", "So an append cannot write past the value."),
    ],
    topics=[
        ("unsafe.Slice over a typed pointer", "Reinterprets one value as its bytes."),
        ("Sizeof for the length", "Hard-coding 8 breaks the moment the type changes."),
        ("Machine layout, not a wire format", "The byte order is the host's."),
    ],
    hint="Convert the pointer to `*byte`, and take the length from the type.",
    intuition="A value's bytes are already in memory in the machine's layout. Viewing them costs a slice header; copying them into a buffer costs an allocation and achieves the same reading.",
    approach=[
        "Return nil for a nil pointer.",
        "`unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p))`.",
    ],
    walkthrough="Setting `v` to all ones makes every byte of the view 0xff, because the view aliases `v` rather than copying it.",
    pitfalls=[
        "Writing 8 instead of `unsafe.Sizeof(*p)`, which is right until the type changes.",
        "Sending the view over a network — the byte order is the local machine's.",
    ],
)

P(
    "middle",
    name="mergemaps",
    title="Merge Into The Map You Were Given",
    sig="func Merge(dst, src map[string]int) int",
    doc="""Merge copies every entry of src into dst, overwriting existing keys,
and returns how many keys were newly added.

dst is modified in place; a nil dst adds nothing.

Examples:

	Merge(dst, map[string]int{"a": 1}) => 1 when dst lacked "a" """,
    solution="""if dst == nil {
	return 0
}
added := 0
for k, v := range src {
	if _, ok := dst[k]; !ok {
		added++
	}
	dst[k] = v
}
return added""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	dst := map[string]int{"a": 1}
	added := Merge(dst, map[string]int{"a": 9, "b": 2})
	if added != 1 {
		t.Errorf("added = %d, want 1", added)
	}
	want := map[string]int{"a": 9, "b": 2}
	if !reflect.DeepEqual(dst, want) {
		t.Errorf("dst = %v, want %v", dst, want)
	}
}

func TestMergeIsVisibleToTheCaller(t *testing.T) {
	dst := map[string]int{}
	alias := dst
	Merge(dst, map[string]int{"x": 1})
	if alias["x"] != 1 {
		t.Error("the merge was not applied to the caller's map")
	}
}

func TestMergeEmpty(t *testing.T) {
	dst := map[string]int{"a": 1}
	if got := Merge(dst, nil); got != 0 {
		t.Errorf("added = %d, want 0", got)
	}
	if len(dst) != 1 {
		t.Errorf("dst = %v, want it unchanged", dst)
	}
	if got := Merge(nil, map[string]int{"a": 1}); got != 0 {
		t.Errorf("added = %d, want 0 for a nil dst", got)
	}
}

func TestMergeDoesNotTouchSrc(t *testing.T) {
	src := map[string]int{"a": 1}
	Merge(map[string]int{"a": 5}, src)
	if src["a"] != 1 {
		t.Errorf("src[a] = %d, want 1", src["a"])
	}
}

func TestMergeCountsOnlyNewKeys(t *testing.T) {
	dst := map[string]int{"a": 1, "b": 2}
	if got := Merge(dst, map[string]int{"a": 9, "b": 9}); got != 0 {
		t.Errorf("added = %d, want 0", got)
	}
}
""",
    context="A settings loader merges defaults, a file and the environment. Each layer built a fresh map, so the process allocated four maps to end up with one.",
    task=[
        "Copy every entry of `src` into `dst`, overwriting existing keys.",
        "Return how many keys were newly added.",
        "Modify `dst` in place; a nil `dst` adds nothing and must not panic.",
        "`src` must not be modified.",
    ],
    examples=[
        ('Merge({"a":1}, {"a":9,"b":2})', '1, dst is {"a":9,"b":2}', "Only \"b\" is new."),
        ("Merge(dst, nil)", "0, dst unchanged", None),
        ("Merge(nil, src)", "0", None),
    ],
    topics=[
        ("Maps are reference-like", "Writing through the parameter reaches the caller's map."),
        ("Comma-ok for presence", "Distinguishes a missing key from a key holding zero."),
        ("In-place over rebuilding", "The destination already has its buckets."),
    ],
    hint="Check presence before writing, or the count is wrong.",
    intuition="The map parameter is a handle to the caller's table, so merging in place is both cheaper and visible. The only care needed is counting: the write itself cannot tell you whether the key was there.",
    approach=[
        "Return 0 for a nil `dst`.",
        "For each entry, test presence with comma-ok, then write.",
        "Return the count of keys that were absent.",
    ],
    walkthrough='Merging {"a":9,"b":2} into {"a":1}: "a" is present so the count stays 0 and the value is overwritten; "b" is absent, so the count becomes 1.',
    pitfalls=[
        "Counting `len(dst)` before and after — correct, and it hides the intent.",
        "Checking `dst[k] != 0`, which mistakes a stored zero for an absent key.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="scannerbuffer",
    title="The Scanner That Gave Up On A Long Line",
    mode="bug",
    sig="func LongestLine(r io.Reader) (int, error)",
    doc="""LongestLine returns the length of the longest line in r.

bufio.Scanner refuses tokens larger than its buffer limit, which defaults
to 64 KiB. A line longer than that is an error, not a truncation.

Examples:

	LongestLine(strings.NewReader("ab\\ncdef")) => 4, nil""",
    imports=['"bufio"', '"io"'],
    extra="""// maxLine is the longest line this reader must accept.
const maxLine = 4 << 20""",
    buggy="""sc := bufio.NewScanner(r)
best := 0
for sc.Scan() {
	if n := len(sc.Bytes()); n > best {
		best = n
	}
}
if err := sc.Err(); err != nil {
	return 0, err
}
return best, nil""",
    solution="""sc := bufio.NewScanner(r)
sc.Buffer(make([]byte, 0, 64*1024), maxLine)
best := 0
for sc.Scan() {
	if n := len(sc.Bytes()); n > best {
		best = n
	}
}
if err := sc.Err(); err != nil {
	return 0, err
}
return best, nil""",
    tests="""
import (
	"errors"
	"strings"
	"testing"
)

func TestLongestLineShort(t *testing.T) {
	got, err := LongestLine(strings.NewReader("ab\\ncdef\\ng"))
	if err != nil || got != 4 {
		t.Errorf("LongestLine = %d, %v, want 4, nil", got, err)
	}
}

func TestLongestLineEmpty(t *testing.T) {
	got, err := LongestLine(strings.NewReader(""))
	if err != nil || got != 0 {
		t.Errorf("LongestLine = %d, %v, want 0, nil", got, err)
	}
}

func TestLongestLineOverTheDefaultLimit(t *testing.T) {
	long := strings.Repeat("x", 200*1024)
	got, err := LongestLine(strings.NewReader("short\\n" + long + "\\n"))
	if err != nil {
		t.Fatalf("LongestLine returned %v, want nil: the scanner's buffer limit was never raised", err)
	}
	if got != len(long) {
		t.Errorf("LongestLine = %d, want %d", got, len(long))
	}
}

func TestLongestLineAtTheConfiguredMax(t *testing.T) {
	long := strings.Repeat("y", 1<<20)
	got, err := LongestLine(strings.NewReader(long))
	if err != nil {
		t.Fatalf("LongestLine returned %v, want nil", err)
	}
	if got != len(long) {
		t.Errorf("LongestLine = %d, want %d", got, len(long))
	}
}

type boom struct{}

func (boom) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestLongestLinePropagatesErrors(t *testing.T) {
	if _, err := LongestLine(boom{}); err == nil {
		t.Error("want the reader's error, got nil")
	}
}
""",
    context="A log analyser works on every sample file and fails on the production log with \"token too long\". The offending line is a stack trace.",
    task=[
        "Return the length of the longest line in `r`.",
        "Lines up to `maxLine` bytes must be accepted.",
        "Fix the single bug; propagate any real read error.",
    ],
    examples=[
        ('LongestLine(strings.NewReader("ab\\ncdef"))', "4, nil", None),
        ("a 200 KiB line", "its length, nil", "Above the default limit."),
        ("a failing reader", "the error", None),
    ],
    topics=[
        ("Scanner.Buffer", "Sets the initial buffer and the maximum token size."),
        ("The default 64 KiB cap", "Silent until a line exceeds it, then a hard error."),
        ("Scanner.Err", "Where `bufio.ErrTooLong` surfaces — the loop just ends."),
        ("Bounded by choice", "The limit protects against a hostile input; it should be explicit."),
    ],
    hint="The loop ends early and `Err` explains why. Which knob was never turned?",
    intuition="`bufio.Scanner` caps token size on purpose, so a malformed stream cannot exhaust memory. The default is a guess about your data, and when it is wrong the loop simply stops and reports the error afterwards.",
    approach=[
        "Call `sc.Buffer` with an initial buffer and `maxLine` as the cap.",
        "Scan as before and check `sc.Err()`.",
    ],
    walkthrough="A 200 KiB line exceeds the default 64 KiB, so `Scan` returns false immediately and `Err` is `bufio.ErrTooLong`. Raising the cap to 4 MiB lets the same input through.",
    pitfalls=[
        "Ignoring `sc.Err()`, which turns the failure into a silently short answer.",
        "Removing the cap entirely by passing a huge maximum — the limit is a safety feature.",
    ],
)

P(
    "senior",
    name="mapvalueretain",
    title="The Map Entry That Pinned The Whole Buffer",
    mode="bug",
    sig="func Index(m map[string][]byte, key string, batch []byte, off, n int)",
    doc="""Index stores the n bytes of batch at offset off under key.

The map outlives the batch, so the stored value must own its bytes: a
view keeps the entire batch reachable for as long as the entry lives.

Examples:

	Index(m, "a", batch, 0, 4) => m["a"] is a 4-byte copy""",
    buggy="""if m == nil || off < 0 || n < 0 || off+n > len(batch) {
	return
}
m[key] = batch[off : off+n : off+n]""",
    solution="""if m == nil || off < 0 || n < 0 || off+n > len(batch) {
	return
}
owned := make([]byte, n)
copy(owned, batch[off:off+n])
m[key] = owned""",
    tests="""
import (
	"bytes"
	"testing"
)

func TestIndexStoresTheBytes(t *testing.T) {
	m := map[string][]byte{}
	batch := []byte("hello world")
	Index(m, "a", batch, 0, 5)
	if !bytes.Equal(m["a"], []byte("hello")) {
		t.Errorf("m[a] = %q, want \\"hello\\"", m["a"])
	}
}

func TestIndexSurvivesBatchReuse(t *testing.T) {
	m := map[string][]byte{}
	batch := make([]byte, 16)
	copy(batch, "first-value")
	Index(m, "a", batch, 0, 5)
	copy(batch, "SECOND-VALUE")
	if !bytes.Equal(m["a"], []byte("first")) {
		t.Errorf("m[a] = %q, want \\"first\\": the entry views the reused batch", m["a"])
	}
}

func TestIndexReleasesTheBatch(t *testing.T) {
	m := map[string][]byte{}
	batch := make([]byte, 1<<20)
	Index(m, "a", batch, 0, 8)
	if cap(m["a"]) > 64 {
		t.Errorf("cap = %d, want a right-sized copy: the entry still owns the batch's array", cap(m["a"]))
	}
}

func TestIndexBadRanges(t *testing.T) {
	m := map[string][]byte{}
	batch := []byte("abcd")
	for _, c := range [][2]int{{-1, 2}, {0, -1}, {3, 3}, {5, 1}} {
		Index(m, "k", batch, c[0], c[1])
		if _, ok := m["k"]; ok {
			t.Fatalf("off=%d n=%d stored an entry, want none", c[0], c[1])
		}
	}
}

func TestIndexNilMap(t *testing.T) {
	Index(nil, "a", []byte("x"), 0, 1)
}
""",
    context="An indexer stores a few bytes per record out of a megabyte read buffer. The index is long-lived, the buffers are not, and resident memory grows by a megabyte per indexed record.",
    task=[
        "Store the `n` bytes of `batch` starting at `off` under `key`.",
        "The stored value must own its bytes — the batch is reused and then dropped.",
        "Ignore a nil map or an out-of-range range.",
        "Fix the single bug.",
    ],
    examples=[
        ('Index(m, "a", batch, 0, 5)', 'm["a"] is a 5-byte copy', None),
        ("the batch is overwritten afterwards", "the entry is unchanged", None),
        ("cap of the stored value", "the copy's size, not the batch's", None),
    ],
    topics=[
        ("Allocation-granular collection", "One live view pins the whole array."),
        ("Three-index slicing is not a copy", "It caps the capacity and still points at the batch."),
        ("Lifetime mismatch", "A short-lived buffer stored in a long-lived map is the whole bug."),
    ],
    hint="The capacity cap makes appends safe. What does it do about the pointer?",
    intuition="`batch[off:off+n:off+n]` protects the batch from being appended over, and does nothing about retention — the header still points into the megabyte. The collector frees allocations, not the parts of them nobody uses.",
    approach=[
        "Validate the range as before.",
        "Allocate `n` bytes, copy the slice into them, store the copy.",
    ],
    walkthrough="Storing 8 bytes out of a 1 MiB batch keeps 1 MiB reachable with a view, and 8 bytes with a copy.",
    pitfalls=[
        "Copying only when `n` is small; the entry's lifetime, not its size, is what matters.",
        "`append([]byte(nil), batch[off:off+n]...)` is also a copy — correct, just less explicit about the size.",
    ],
)

P(
    "senior",
    name="zerostrings",
    title="Blank Every String, However Deep",
    sig="func Redact(ptr any) error",
    doc="""Redact sets every exported string field of the struct ptr points at to
"", descending into nested structs and slices of structs.

Unexported fields are left alone.

Examples:

	Redact(&record{Name: "x"}) => nil, record.Name is "" """,
    imports=['"errors"', '"reflect"'],
    extra="""// ErrTarget is returned when ptr is not a non-nil pointer to a struct.
var ErrTarget = errors.New("target must be a non-nil pointer to a struct")

func redact(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.String:
		if rv.CanSet() {
			rv.SetString("")
		}
	case reflect.Pointer, reflect.Interface:
		if !rv.IsNil() {
			redact(rv.Elem())
		}
	case reflect.Struct:
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			if rt.Field(i).IsExported() {
				redact(rv.Field(i))
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			redact(rv.Index(i))
		}
	}
}""",
    solution="""rv := reflect.ValueOf(ptr)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return ErrTarget
}
rv = rv.Elem()
if rv.Kind() != reflect.Struct {
	return ErrTarget
}
redact(rv)
return nil""",
    tests="""
import (
	"errors"
	"testing"
)

type inner struct {
	Secret string
	Count  int
}

type record struct {
	Name   string
	In     inner
	Ptr    *inner
	List   []inner
	hidden string
}

func TestRedactFlat(t *testing.T) {
	r := &record{Name: "top", hidden: "keep"}
	if err := Redact(r); err != nil {
		t.Fatal(err)
	}
	if r.Name != "" {
		t.Errorf("Name = %q, want empty", r.Name)
	}
	if r.hidden != "keep" {
		t.Errorf("hidden = %q, want \\"keep\\"", r.hidden)
	}
}

func TestRedactNested(t *testing.T) {
	r := &record{
		Name: "top",
		In:   inner{Secret: "a", Count: 1},
		Ptr:  &inner{Secret: "b"},
		List: []inner{{Secret: "c"}, {Secret: "d"}},
	}
	if err := Redact(r); err != nil {
		t.Fatal(err)
	}
	if r.In.Secret != "" || r.Ptr.Secret != "" {
		t.Errorf("nested secrets survived: %+v", r)
	}
	for i, e := range r.List {
		if e.Secret != "" {
			t.Errorf("List[%d].Secret = %q, want empty", i, e.Secret)
		}
	}
	if r.In.Count != 1 {
		t.Errorf("Count = %d, want 1: non-string fields must be untouched", r.In.Count)
	}
}

func TestRedactNilPointerField(t *testing.T) {
	r := &record{Name: "x"}
	if err := Redact(r); err != nil {
		t.Fatal(err)
	}
	if r.Ptr != nil {
		t.Error("the nil pointer field was replaced")
	}
}

func TestRedactBadTarget(t *testing.T) {
	for _, c := range []any{record{}, nil, (*record)(nil), new(int)} {
		if err := Redact(c); !errors.Is(err, ErrTarget) {
			t.Errorf("Redact(%#v) = %v, want ErrTarget", c, err)
		}
	}
}
""",
    context="A crash reporter serialises the request struct into the report. Legal asks for every string field to be blanked first, and the struct has four levels of nesting that change every sprint.",
    task=[
        "Blank every exported string reachable from the struct `ptr` points at.",
        "Descend into nested structs, pointers, interfaces, slices and arrays.",
        "Leave unexported fields and non-string fields alone; a nil pointer field stays nil.",
        "Return `ErrTarget` unless `ptr` is a non-nil pointer to a struct.",
    ],
    examples=[
        ('Redact(&record{Name:"x"})', 'nil, Name is ""', None),
        ("a nested Ptr and a List of structs", "every Secret blanked", None),
        ("Redact(record{})", "ErrTarget", "A value cannot be written through."),
    ],
    topics=[
        ("Settability propagates", "Fields reached from an addressable struct are addressable too."),
        ("Recursive kind dispatch", "One case per container kind, recursing on the contents."),
        ("Slice elements are addressable", "`rv.Index(i)` of a slice can be set; of an array only when the array is."),
        ("CanSet as the final guard", "It covers both addressability and export status."),
    ],
    hint="The recursive helper is written for you. Validate, step through `Elem`, and let it walk.",
    intuition="Reflection's write rules follow the language's: you can only assign through something addressable. Starting from a pointer makes the whole tree beneath it addressable, so one recursive walk can rewrite every leaf.",
    approach=[
        "Verify `ptr` is a non-nil pointer to a struct.",
        "Step to the struct with `Elem`.",
        "Call the recursive helper.",
    ],
    walkthrough="From `&record{}`, `Elem` is addressable; each exported field inherits that, so the string leaves under `In`, `*Ptr` and every `List` element can be set.",
    pitfalls=[
        "Starting from `reflect.ValueOf(v)` of a value, where nothing is settable and every write is silently skipped.",
        "Recursing into unexported fields, whose values cannot be set and may panic on access.",
        "Following a nil pointer into `Elem`, which yields an invalid Value.",
    ],
)

P(
    "senior",
    name="structroundtrip",
    title="A Struct To Bytes And Back",
    sig="func Decode(b []byte) (Frame, bool)",
    doc="""Decode reinterprets b as a Frame, copying it out so the result does not
alias b.

The length must be exactly the frame's size and the start must be
correctly aligned; otherwise the second result is false.

Examples:

	Decode(encoded) => the frame, true""",
    imports=['"unsafe"'],
    extra="""// Frame is a fixed-layout record of scalars.
type Frame struct {
	Kind  uint32
	Seq   uint32
	Stamp int64
}

// Encode returns a byte view of f, for the tests to feed back in.
func Encode(f *Frame) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(f)), unsafe.Sizeof(*f))
}""",
    solution="""var zero Frame
size := unsafe.Sizeof(zero)
if uintptr(len(b)) != size {
	return zero, false
}
p := unsafe.Pointer(unsafe.SliceData(b))
if uintptr(p)%unsafe.Alignof(zero) != 0 {
	return zero, false
}
return *(*Frame)(p), true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestDecodeRoundTrip(t *testing.T) {
	in := Frame{Kind: 7, Seq: 9, Stamp: 1234567890}
	got, ok := Decode(Encode(&in))
	if !ok {
		t.Fatal("Decode reported false for a well-formed frame")
	}
	if got != in {
		t.Errorf("Decode = %+v, want %+v", got, in)
	}
}

func TestDecodeCopiesOut(t *testing.T) {
	in := Frame{Kind: 1}
	b := Encode(&in)
	got, ok := Decode(b)
	if !ok {
		t.Fatal("Decode reported false")
	}
	in.Kind = 99
	if got.Kind != 1 {
		t.Error("the result aliases the input bytes; it must be a copy")
	}
}

func TestDecodeWrongLength(t *testing.T) {
	in := Frame{}
	b := Encode(&in)
	for _, c := range [][]byte{nil, b[:4], b[:len(b)-1], append(append([]byte{}, b...), 0)} {
		if _, ok := Decode(c); ok {
			t.Errorf("Decode of %d bytes reported ok, want false", len(c))
		}
	}
}

func TestDecodeMisaligned(t *testing.T) {
	var zero Frame
	size := int(unsafe.Sizeof(zero))
	backing := make([]Frame, 2)
	all := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(backing))), size*2)
	if _, ok := Decode(all[1 : 1+size]); ok {
		t.Error("a misaligned frame reported ok, want false")
	}
}

func TestDecodeAllocatesNothing(t *testing.T) {
	in := Frame{Kind: 3}
	b := Encode(&in)
	var sink Frame
	if n := testing.AllocsPerRun(200, func() { sink, _ = Decode(b) }); n != 0 {
		t.Errorf("Decode made %v allocations, want 0", n)
	}
	_ = sink
}
""",
    context="A framing layer decodes fixed-size records field by field with `binary.LittleEndian`. Both ends are the same architecture, and the field-by-field decode is most of the receive path.",
    task=[
        "Reinterpret `b` as a `Frame` and return a copy of it.",
        "Report false unless `len(b)` is exactly the frame's size and the start is correctly aligned.",
        "The result must not alias `b`; allocate nothing.",
    ],
    examples=[
        ("Decode(Encode(&f))", "f, true", None),
        ("Decode(b[:4])", "zero Frame, false", "Wrong length."),
        ("a misaligned slice", "zero Frame, false", None),
    ],
    topics=[
        ("Dereference copies", "`*(*Frame)(p)` reads the struct out by value."),
        ("Size and alignment from the type", "`Sizeof` and `Alignof`, never literals."),
        ("Exact length", "Too short reads past the end; too long is a framing error."),
        ("Not portable", "The layout, padding and byte order are the local machine's."),
    ],
    hint="Two guards from the type, then one dereference.",
    intuition="Reinterpreting is only sound when the bytes really are a `Frame` — the right size, at the right alignment. The dereference then copies the struct out, which is what makes the result independent of the buffer.",
    approach=[
        "Compare `len(b)` with `unsafe.Sizeof` of a zero Frame.",
        "Check the data pointer against `unsafe.Alignof`.",
        "Return `*(*Frame)(p), true`.",
    ],
    walkthrough="A well-formed 16-byte aligned buffer decodes by one struct load. Changing the source afterwards cannot affect the result, because the dereference copied the fields out.",
    pitfalls=[
        "Returning `(*Frame)(p)` — a pointer into the caller's buffer, which is the aliasing the spec forbids.",
        "Accepting `len(b) >= size`, which silently ignores a framing error.",
        "Assuming the padding bytes carry information; they do not.",
    ],
)

P(
    "senior",
    name="dropwhenfull",
    title="Drop Instead Of Blocking",
    sig="func Offer(ch chan<- int, v int) bool",
    doc="""Offer sends v on ch if it can be accepted immediately, and reports
whether it was.

A metrics pipeline must never block its caller: when the buffer is full,
the sample is dropped.

Examples:

	Offer(ch, 1) => true when ch has room""",
    solution="""select {
case ch <- v:
	return true
default:
	return false
}""",
    tests="""
import (
	"sync"
	"testing"
	"time"
)

func TestOfferWithRoom(t *testing.T) {
	ch := make(chan int, 2)
	if !Offer(ch, 1) {
		t.Error("Offer = false, want true: the buffer had room")
	}
	if got := <-ch; got != 1 {
		t.Errorf("received %d, want 1", got)
	}
}

func TestOfferWhenFull(t *testing.T) {
	ch := make(chan int, 1)
	Offer(ch, 1)
	if Offer(ch, 2) {
		t.Error("Offer = true, want false: the buffer was full")
	}
	if got := <-ch; got != 1 {
		t.Errorf("received %d, want 1: the dropped value must not displace the first", got)
	}
}

func TestOfferDoesNotBlock(t *testing.T) {
	ch := make(chan int) // unbuffered, no receiver
	done := make(chan bool, 1)
	go func() { done <- Offer(ch, 1) }()
	select {
	case ok := <-done:
		if ok {
			t.Error("Offer = true on an unbuffered channel with no receiver")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Offer blocked, want an immediate false")
	}
}

func TestOfferUnbufferedWithReceiver(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	got := make(chan int, 1)
	go func() {
		defer wg.Done()
		got <- <-ch
	}()
	deadline := time.Now().Add(2 * time.Second)
	for !Offer(ch, 7) && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
	if v := <-got; v != 7 {
		t.Errorf("received %d, want 7", v)
	}
}

func TestOfferFillsExactly(t *testing.T) {
	ch := make(chan int, 4)
	accepted := 0
	for i := 0; i < 10; i++ {
		if Offer(ch, i) {
			accepted++
		}
	}
	if accepted != 4 {
		t.Errorf("accepted %d, want 4", accepted)
	}
}
""",
    context="A request handler publishes a metric on a buffered channel. The collector stalls, the buffer fills, and every request in the service blocks behind a statistics counter.",
    task=[
        "Send `v` on `ch` if it can be accepted immediately.",
        "Report whether it was sent.",
        "Never block, whatever the channel's state.",
    ],
    examples=[
        ("Offer(ch, 1) with room", "true", None),
        ("Offer on a full buffer", "false, the value is dropped", None),
        ("Offer on an unbuffered channel with no receiver", "false, immediately", None),
    ],
    topics=[
        ("select with default", "Makes a channel operation non-blocking."),
        ("Load shedding", "Dropping a sample beats stalling the request that produced it."),
        ("Unbuffered means synchronous", "It can only accept when a receiver is already waiting."),
    ],
    hint="One `select`, one case, one default.",
    intuition="A `select` with a `default` asks \"can this proceed right now\" instead of \"wait until it can\". That turns backpressure from a stall into a decision the caller gets to make.",
    approach=[
        "`select` with the send as the only case.",
        "`default` returns false.",
    ],
    walkthrough="With a buffer of 4 and ten offers, the first four are accepted and the rest take the default branch immediately.",
    pitfalls=[
        "Checking `len(ch) < cap(ch)` first — the state can change between the check and the send.",
        "Dropping silently in a context where the caller needed to know; the boolean is the point.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="lrucache",
    title="Evict The One Nobody Has Touched",
    sig="func (c *LRU) Get(key string) (int, bool)",
    doc="""Get returns the value for key and marks it as the most recently used.

A cache that evicts by insertion order throws away the hot entries; the
whole point of LRU is that a hit moves the entry to the front.

Examples:

	c.Put("a", 1); c.Get("a") => 1, true and "a" becomes newest""",
    imports=['"container/list"', '"sync"'],
    extra="""// entry is one cached pair, stored in the list.
type entry struct {
	key string
	val int
}

// LRU is a bounded, concurrency-safe least-recently-used cache.
type LRU struct {
	mu    sync.Mutex
	limit int
	ll    *list.List
	items map[string]*list.Element
}

// NewLRU returns a cache holding at most limit entries.
func NewLRU(limit int) *LRU {
	if limit < 1 {
		limit = 1
	}
	return &LRU{limit: limit, ll: list.New(), items: make(map[string]*list.Element, limit)}
}

// Put stores a value, evicting the least recently used entry if needed.
func (c *LRU) Put(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el)
		el.Value.(*entry).val = val
		return
	}
	if c.ll.Len() >= c.limit {
		oldest := c.ll.Back()
		if oldest != nil {
			c.ll.Remove(oldest)
			delete(c.items, oldest.Value.(*entry).key)
		}
	}
	c.items[key] = c.ll.PushFront(&entry{key: key, val: val})
}

// Len reports how many entries the cache holds.
func (c *LRU) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ll.Len()
}""",
    solution="""c.mu.Lock()
defer c.mu.Unlock()
el, ok := c.items[key]
if !ok {
	return 0, false
}
c.ll.MoveToFront(el)
return el.Value.(*entry).val, true""",
    tests="""
import (
	"sync"
	"testing"
)

func TestGetAndPut(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if _, ok := c.Get("missing"); ok {
		t.Error("Get(missing) reported ok, want false")
	}
}

func TestGetMarksAsRecentlyUsed(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	if _, ok := c.Get("a"); !ok {
		t.Fatal("a was not found")
	}
	c.Put("c", 3) // evicts the least recently used, which must be b
	if _, ok := c.Get("a"); !ok {
		t.Error("a was evicted, but Get(a) had just made it the newest")
	}
	if _, ok := c.Get("b"); ok {
		t.Error("b survived, want it evicted as the least recently used")
	}
}

func TestEvictsWithoutAnyGets(t *testing.T) {
	c := NewLRU(2)
	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)
	if _, ok := c.Get("a"); ok {
		t.Error("a survived, want it evicted")
	}
	if c.Len() != 2 {
		t.Errorf("Len = %d, want 2", c.Len())
	}
}

func TestStaysBounded(t *testing.T) {
	c := NewLRU(8)
	for i := 0; i < 2000; i++ {
		c.Put(string(rune('a'+i%26))+string(rune('a'+i/26%26)), i)
		c.Get(string(rune('a' + i%26)))
		if c.Len() > 8 {
			t.Fatalf("Len = %d, want at most 8", c.Len())
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	c := NewLRU(16)
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				k := string(rune('a' + i%10))
				c.Put(k, i)
				c.Get(k)
			}
		}(w)
	}
	wg.Wait()
	if c.Len() > 16 {
		t.Errorf("Len = %d, want at most 16", c.Len())
	}
}
""",
    context="A cache with a FIFO eviction rule has a hit rate of eleven percent. The working set fits comfortably in the limit, and the entries being evicted are the ones every request needs.",
    task=[
        "Return the value for `key` and whether it was present.",
        "A hit must make the entry the most recently used, so it is the last to be evicted.",
        "Hold the lock for the lookup and the reordering together.",
    ],
    examples=[
        ('c.Put("a",1); c.Get("a")', "1, true", None),
        ('Put a, Put b, Get a, Put c with limit 2', "b is evicted, a survives", "The Get rescued a."),
        ("no Gets, limit 2, Put a b c", "a is evicted", None),
    ],
    topics=[
        ("Recency needs a Get-side update", "Without it, LRU degenerates into FIFO."),
        ("list.MoveToFront", "O(1) reordering in a doubly linked list."),
        ("Map to list element", "The map gives O(1) lookup; the list gives O(1) ordering."),
        ("Reads mutate the structure", "So `Get` needs the write lock, not a read lock."),
    ],
    hint="A hit does two things, not one.",
    intuition="LRU's entire value is that a read is evidence of future use. If `Get` only reads, the ordering never reflects access and the cache evicts exactly the entries it should keep.",
    approach=[
        "Take the lock.",
        "Look the key up; return `0, false` on a miss.",
        "`MoveToFront` the element, then return its value.",
    ],
    walkthrough="With limit 2 holding a and b, `Get(\"a\")` moves a to the front. Putting c then evicts the back — b — instead of a.",
    pitfalls=[
        "Using an RWMutex and taking the read lock in `Get`; the reordering is a write.",
        "Returning the value before moving the element, then forgetting to move it at all.",
        "Storing the value in the map and the key in the list, which makes eviction unable to find the map entry.",
    ],
)

P(
    "staff",
    name="deadlinepool",
    title="Stop When The Deadline Passes",
    sig="func Process(ctx context.Context, items []int, workers int) ([]int, error)",
    doc="""Process doubles every item using workers goroutines and returns the
results in input order.

If ctx is cancelled or its deadline passes first, Process returns the
context's error, and every goroutine it started must have exited.

Examples:

	Process(ctx, []int{1, 2}, 2) => []int{2, 4}, nil""",
    imports=['"context"', '"sync"'],
    solution="""if workers < 1 {
	workers = 1
}
out := make([]int, len(items))
if len(items) == 0 {
	return out, ctx.Err()
}
if workers > len(items) {
	workers = len(items)
}

idx := make(chan int)
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	go func() {
		defer wg.Done()
		for {
			select {
			case i, ok := <-idx:
				if !ok {
					return
				}
				out[i] = items[i] * 2
			case <-ctx.Done():
				return
			}
		}
	}()
}

feed:
for i := range items {
	select {
	case idx <- i:
	case <-ctx.Done():
		break feed
	}
}
close(idx)
wg.Wait()

if err := ctx.Err(); err != nil {
	return nil, err
}
return out, nil""",
    tests="""
import (
	"context"
	"errors"
	"runtime"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	got, err := Process(context.Background(), []int{1, 2, 3}, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got[0] != 2 || got[1] != 4 || got[2] != 6 {
		t.Errorf("Process = %v, want [2 4 6]", got)
	}
}

func TestProcessEmpty(t *testing.T) {
	got, err := Process(context.Background(), nil, 4)
	if err != nil || len(got) != 0 {
		t.Errorf("Process = %v, %v, want empty, nil", got, err)
	}
}

func TestProcessWorkerCounts(t *testing.T) {
	items := make([]int, 1001)
	for i := range items {
		items[i] = i
	}
	for _, w := range []int{0, 1, 3, 64, 100000} {
		got, err := Process(context.Background(), items, w)
		if err != nil {
			t.Fatalf("workers=%d: %v", w, err)
		}
		for i := range items {
			if got[i] != items[i]*2 {
				t.Fatalf("workers=%d: got[%d] = %d, want %d", w, i, got[i], items[i]*2)
			}
		}
	}
}

func TestProcessCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Process(ctx, []int{1, 2, 3}, 2)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestProcessDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	time.Sleep(20 * time.Millisecond)
	items := make([]int, 10000)
	if _, err := Process(ctx, items, 4); !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want context.DeadlineExceeded", err)
	}
}

func TestProcessDoesNotLeak(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()
	for i := 0; i < 20; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		items := make([]int, 1000)
		Process(ctx, items, 8)
	}
	deadline := time.Now().Add(3 * time.Second)
	for runtime.NumGoroutine() > base+4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, want about %d", got, base)
	}
}
""",
    context="A batch endpoint fans work out to a worker pool. When the client disconnects the handler returns, and the workers keep going until the batch is finished.",
    task=[
        "Double every item using `workers` goroutines, results in input order.",
        "Return `ctx.Err()` when the context is done, with a nil result.",
        "Every goroutine must exit before `Process` returns, cancelled or not.",
        "`workers < 1` behaves as 1; more workers than items is legal.",
    ],
    examples=[
        ("Process(ctx, []int{1,2,3}, 2)", "[2 4 6], nil", None),
        ("an already-cancelled context", "nil, context.Canceled", None),
        ("20 cancelled runs of 1000 items", "no goroutines left", None),
    ],
    topics=[
        ("Cancellation on every blocking operation", "Both the feed's send and the workers' receive need a `ctx.Done()` case."),
        ("Disjoint slot writes", "`out[i]` from one worker needs no lock."),
        ("Close then Wait", "Closing the index channel ends the workers on the normal path."),
        ("Wait before returning", "It is what makes \"no goroutine outlives the call\" true."),
    ],
    hint="Three places can block: the send of an index, the receive of one, and the join.",
    intuition="A cancellable pool is not one that checks a flag between items — it is one where every point that can wait also watches the context. The join at the end is what turns \"they will stop eventually\" into a guarantee.",
    approach=[
        "Normalise `workers` and handle the empty input.",
        "Start workers that `select` on the index channel and `ctx.Done()`.",
        "Feed indices with a `select` that breaks out on cancellation.",
        "Close the channel, `Wait`, then return `ctx.Err()` or the results.",
    ],
    walkthrough="On cancellation the feed loop breaks, the channel is closed, and each worker returns from whichever `select` it is in. `Wait` then returns and the error is reported.",
    pitfalls=[
        "Returning on `ctx.Done()` without waiting, which leaves the workers running past the call.",
        "Feeding with a plain send, which blocks forever once the workers have exited.",
        "Returning partial results on cancellation when the caller expects all or nothing.",
    ],
)

P(
    "staff",
    name="mergesorted",
    title="Merge Runs Into The Caller's Buffer",
    sig="func Merge(dst []int, runs [][]int) []int",
    doc="""Merge appends every element of the sorted runs to dst in ascending
order and returns the extended slice.

With room in dst, the merge must allocate nothing: the cursors are the
only state it needs.

Examples:

	Merge(nil, [][]int{{1, 3}, {2}}) => []int{1, 2, 3}""",
    solution="""total := 0
for _, r := range runs {
	total += len(r)
}
if total == 0 {
	return dst
}
var cursors [16]int
pos := cursors[:0]
if len(runs) <= cap(pos) {
	pos = cursors[:len(runs)]
} else {
	pos = make([]int, len(runs))
}
for n := 0; n < total; n++ {
	best := -1
	for i, r := range runs {
		if pos[i] >= len(r) {
			continue
		}
		if best < 0 || r[pos[i]] < runs[best][pos[best]] {
			best = i
		}
	}
	dst = append(dst, runs[best][pos[best]])
	pos[best]++
}
return dst""",
    tests="""
import (
	"reflect"
	"sort"
	"testing"
)

var sink []int

func TestMerge(t *testing.T) {
	got := Merge(nil, [][]int{{1, 3, 5}, {2, 4}, {}})
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Merge = %v, want [1 2 3 4 5]", got)
	}
}

func TestMergeAppendsToDst(t *testing.T) {
	got := Merge([]int{0}, [][]int{{2}, {1}})
	if !reflect.DeepEqual(got, []int{0, 1, 2}) {
		t.Errorf("Merge = %v, want [0 1 2]", got)
	}
}

func TestMergeEmpty(t *testing.T) {
	if got := Merge(nil, nil); len(got) != 0 {
		t.Errorf("Merge = %v, want empty", got)
	}
	if got := Merge(nil, [][]int{{}, {}}); len(got) != 0 {
		t.Errorf("Merge = %v, want empty", got)
	}
}

func TestMergeDuplicates(t *testing.T) {
	got := Merge(nil, [][]int{{1, 1}, {1}})
	if !reflect.DeepEqual(got, []int{1, 1, 1}) {
		t.Errorf("Merge = %v, want [1 1 1]", got)
	}
}

func TestMergeMatchesSorting(t *testing.T) {
	runs := make([][]int, 7)
	var all []int
	for i := range runs {
		r := make([]int, 0, 20)
		for j := 0; j < 20; j++ {
			r = append(r, i+j*7)
		}
		sort.Ints(r)
		runs[i] = r
		all = append(all, r...)
	}
	sort.Ints(all)
	got := Merge(nil, runs)
	if !reflect.DeepEqual(got, all) {
		t.Errorf("Merge produced a different order than sorting the union")
	}
}

func TestMergeAllocatesNothingWithRoom(t *testing.T) {
	runs := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	dst := make([]int, 0, 32)
	if n := testing.AllocsPerRun(100, func() { sink = Merge(dst[:0], runs) }); n != 0 {
		t.Errorf("Merge made %v allocations, want 0", n)
	}
}
""",
    context="An external sort merges its runs by concatenating and re-sorting. The concatenation doubles peak memory and throws away the ordering the runs already have.",
    task=[
        "Append every element of the sorted `runs` to `dst` in ascending order.",
        "Preserve duplicates; an empty input returns `dst` unchanged.",
        "With room in `dst` and a modest number of runs, allocate nothing.",
    ],
    examples=[
        ("Merge(nil, [][]int{{1,3,5},{2,4}})", "[1 2 3 4 5]", None),
        ("Merge([]int{0}, [][]int{{2},{1}})", "[0 1 2]", "dst is extended."),
        ("Merge(nil, [][]int{{1,1},{1}})", "[1 1 1]", "Duplicates survive."),
    ],
    topics=[
        ("Merging preserves work", "The runs are already sorted; re-sorting throws that away."),
        ("Cursors as the only state", "One index per run, no data movement."),
        ("A stack array for small cases", "A fixed local array avoids allocating the cursor slice."),
        ("Append-style output", "The caller owns the destination."),
    ],
    hint="One cursor per run. Each step picks the smallest element still under a cursor.",
    intuition="A k-way merge is a repeated minimum over the run heads. Nothing needs copying and nothing needs sorting — the only state is where each run has got to, which is small enough to live on the stack.",
    approach=[
        "Sum the run lengths; return `dst` when the total is zero.",
        "Keep cursors in a fixed local array when the run count allows, otherwise allocate.",
        "Repeat `total` times: scan for the smallest head, append it, advance its cursor.",
    ],
    walkthrough="Merging {1,4,7}, {2,5,8} and {3,6,9} advances one cursor per step. The cursor array lives in the frame, so with room in `dst` the whole merge allocates nothing.",
    pitfalls=[
        "Allocating the cursor slice unconditionally, which is one allocation per call.",
        "Skipping exhausted runs incorrectly and indexing past a run's end.",
        "A heap is the right structure for many runs; the linear scan is fine for a few.",
    ],
)

P(
    "staff",
    name="semaphore",
    title="Bound The Work In Flight",
    sig="func (s *Sem) Acquire(done <-chan struct{}) bool",
    doc="""Acquire takes one slot, blocking until one is free, and reports whether
it got one.

It must also give up when done is closed, so a cancelled caller does not
wait forever on a saturated semaphore.

Examples:

	s := NewSem(2); s.Acquire(done) => true twice, then blocks""",
    extra="""// Sem is a counting semaphore of fixed capacity.
type Sem struct {
	slots chan struct{}
}

// NewSem returns a semaphore permitting n concurrent holders.
func NewSem(n int) *Sem {
	if n < 1 {
		n = 1
	}
	return &Sem{slots: make(chan struct{}, n)}
}

// Release returns one slot. It must only be called after a successful Acquire.
func (s *Sem) Release() {
	select {
	case <-s.slots:
	default:
	}
}

// Held reports how many slots are currently taken.
func (s *Sem) Held() int { return len(s.slots) }""",
    solution="""select {
case s.slots <- struct{}{}:
	return true
case <-done:
	return false
}""",
    tests="""
import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAcquireUpToCapacity(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	s := NewSem(2)
	if !s.Acquire(done) || !s.Acquire(done) {
		t.Fatal("the first two acquires must succeed")
	}
	if s.Held() != 2 {
		t.Errorf("Held = %d, want 2", s.Held())
	}
}

func TestAcquireBlocksWhenFull(t *testing.T) {
	done := make(chan struct{})
	s := NewSem(1)
	if !s.Acquire(done) {
		t.Fatal("the first acquire must succeed")
	}
	got := make(chan bool, 1)
	go func() { got <- s.Acquire(done) }()
	select {
	case <-got:
		t.Fatal("the second acquire returned while the semaphore was full")
	case <-time.After(50 * time.Millisecond):
	}
	close(done)
	if ok := <-got; ok {
		t.Error("the cancelled acquire reported true, want false")
	}
}

func TestReleaseUnblocksAWaiter(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	s := NewSem(1)
	s.Acquire(done)
	got := make(chan bool, 1)
	go func() { got <- s.Acquire(done) }()
	time.Sleep(20 * time.Millisecond)
	s.Release()
	select {
	case ok := <-got:
		if !ok {
			t.Error("the waiter reported false, want true")
		}
	case <-time.After(2 * time.Second):
		t.Error("the waiter was never released")
	}
}

func TestConcurrencyIsBounded(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	const limit = 4
	s := NewSem(limit)
	var inFlight, peak atomic.Int64
	var wg sync.WaitGroup
	const workers = 32
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				if !s.Acquire(done) {
					return
				}
				n := inFlight.Add(1)
				for {
					p := peak.Load()
					if n <= p || peak.CompareAndSwap(p, n) {
						break
					}
				}
				inFlight.Add(-1)
				s.Release()
			}
		}()
	}
	wg.Wait()
	if peak.Load() > limit {
		t.Errorf("peak concurrency was %d, want at most %d", peak.Load(), limit)
	}
}
""",
    context="A crawler starts a goroutine per URL. With a large frontier it opens ten thousand sockets, exhausts the file descriptor limit, and takes the rest of the process with it.",
    task=[
        "Take one slot, waiting until one is free.",
        "Report false and take nothing when `done` is closed first.",
        "Never allow more than the configured number of holders.",
    ],
    examples=[
        ("NewSem(2), two Acquires", "true, true", None),
        ("a third Acquire while full", "blocks, then false when done closes", None),
        ("Release while someone waits", "the waiter proceeds", None),
    ],
    topics=[
        ("A buffered channel is a semaphore", "Its capacity is the permit count."),
        ("Send to acquire, receive to release", "The buffer's occupancy is the held count."),
        ("Cancellable waiting", "`select` over the acquire and `done`."),
        ("Acquire nothing on cancellation", "Returning false must not leave a slot taken."),
    ],
    hint="Two cases in one `select`: taking a slot, and giving up.",
    intuition="A buffered channel already implements counting with blocking — sending is acquiring and the capacity is the limit. The only thing it lacks is a way out, and `select` supplies that.",
    approach=[
        "`select` on sending an empty struct into `s.slots` and on `<-done`.",
        "Return true for the first, false for the second.",
    ],
    walkthrough="With 32 workers and a limit of 4, at most four sends fit in the buffer; the rest block in the `select` until a `Release` receives one out, or `done` closes.",
    pitfalls=[
        "Checking `len(s.slots) < cap(s.slots)` first — the state can change before the send.",
        "Returning true on the `done` branch, so the caller releases a slot it never took.",
        "Releasing without a matching acquire, which the `default` in `Release` tolerates but which breaks the count.",
    ],
)

P(
    "staff",
    name="genericpool",
    title="A Typed Pool With No Assertions",
    sig="func (p *Pool[T]) Get() *T",
    doc="""Get returns a pointer to a zeroed T from the pool, or a new one when
the pool is empty.

The type parameter keeps the values typed on the way in and out, so no
caller ever writes a type assertion.

Examples:

	p := NewPool[Buffer](); p.Get() => a zeroed *Buffer""",
    imports=['"sync"'],
    extra="""// Pool is a typed wrapper around sync.Pool.
type Pool[T any] struct {
	inner sync.Pool
}

// NewPool returns a pool of T values.
func NewPool[T any]() *Pool[T] {
	return &Pool[T]{inner: sync.Pool{New: func() any { return new(T) }}}
}

// Put returns a value to the pool.
func (p *Pool[T]) Put(v *T) {
	if v == nil {
		return
	}
	p.inner.Put(v)
}""",
    solution="""v := p.inner.Get().(*T)
var zero T
*v = zero
return v""",
    tests="""
import (
	"sync"
	"testing"
)

type buffer struct {
	Name  string
	Count int
}

func TestGetReturnsAZeroValue(t *testing.T) {
	p := NewPool[buffer]()
	v := p.Get()
	if v == nil {
		t.Fatal("Get returned nil")
	}
	if *v != (buffer{}) {
		t.Errorf("Get = %+v, want the zero buffer", *v)
	}
}

func TestGetZeroesARecycledValue(t *testing.T) {
	p := NewPool[buffer]()
	v := p.Get()
	v.Name = "dirty"
	v.Count = 9
	p.Put(v)
	got := p.Get()
	if *got != (buffer{}) {
		t.Errorf("Get = %+v, want the zero buffer: a recycled value must be reset", *got)
	}
}

func TestGetIsTypedWithoutAssertions(t *testing.T) {
	p := NewPool[int]()
	v := p.Get()
	*v = 42
	if *v != 42 {
		t.Errorf("*v = %d, want 42", *v)
	}
	p.Put(v)
	if got := p.Get(); *got != 0 {
		t.Errorf("*got = %d, want 0", *got)
	}
}

func TestPutNil(t *testing.T) {
	p := NewPool[buffer]()
	p.Put(nil)
	if v := p.Get(); v == nil {
		t.Error("Get returned nil after Put(nil)")
	}
}

func TestConcurrentGetPut(t *testing.T) {
	p := NewPool[buffer]()
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				v := p.Get()
				if v.Count != 0 || v.Name != "" {
					panic("a dirty value came out of the pool")
				}
				v.Count = w
				v.Name = "in-use"
				if v.Count != w {
					panic("value shared between goroutines")
				}
				p.Put(v)
			}
		}(w)
	}
	wg.Wait()
}
""",
    context="Every pool in the codebase is a `sync.Pool` plus a type assertion at each call site. One of them asserts the wrong type, and the panic only fires under the load that fills the pool.",
    task=[
        "Return a pointer to a zeroed `T` from the pool.",
        "A recycled value must come back zeroed — callers must never see the previous holder's data.",
        "No type assertion may escape to the caller; safe for concurrent use.",
    ],
    examples=[
        ("NewPool[buffer]().Get()", "a zeroed *buffer", None),
        ("Put a dirty value, then Get", "zeroed again", None),
        ("NewPool[int]()", "works with any type", None),
    ],
    topics=[
        ("Generics over sync.Pool", "The assertion happens once, inside the wrapper."),
        ("Pooled values carry state", "Resetting is the wrapper's job, not the caller's."),
        ("The zero value as the reset", "`*v = zero` works for any T without knowing its fields."),
        ("sync.Pool may drop entries", "`New` covers the empty case, so `Get` never returns nil."),
    ],
    hint="The assertion is safe because `New` is the only thing that ever puts a value in — and then one more line makes it clean.",
    intuition="A pool's danger is stale state and its friction is type assertions. Generics remove the second, and a single assignment of the zero value removes the first for every type at once.",
    approach=[
        "`p.inner.Get().(*T)` — safe because only `*T` values ever enter the pool.",
        "Overwrite with the zero `T`.",
        "Return the pointer.",
    ],
    walkthrough="A `buffer` put back with `Name` set comes out of the pool with that name still in it; assigning the zero value clears every field without the wrapper knowing what they are.",
    pitfalls=[
        "Skipping the reset and documenting that callers must do it — one of them will not.",
        "Putting a value back while still holding a pointer to it.",
        "Assuming the pool retains what you put; entries can be dropped at any collection.",
    ],
)
