"""10-advanced-topics — rotation 3: 5 puzzles each for middle, senior, staff."""

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
    name="mapappend",
    title="The Append That Never Reached The Map",
    mode="bug",
    sig="func Add(m map[string][]int, key string, v int)",
    doc="""Add appends v to the slice stored under key, creating the entry when it
is missing.

A map value is not addressable: appending to m[key] produces a new slice
header that has to be stored back.

Examples:

	m := map[string][]int{}; Add(m, "a", 1) => m["a"] is [1]""",
    buggy="""if m == nil {
	return
}
_ = append(m[key], v)""",
    solution="""if m == nil {
	return
}
m[key] = append(m[key], v)""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	m := map[string][]int{}
	Add(m, "a", 1)
	Add(m, "a", 2)
	Add(m, "b", 3)
	want := map[string][]int{"a": {1, 2}, "b": {3}}
	if !reflect.DeepEqual(m, want) {
		t.Errorf("m = %v, want %v", m, want)
	}
}

func TestAddCreatesTheEntry(t *testing.T) {
	m := map[string][]int{}
	Add(m, "new", 7)
	got, ok := m["new"]
	if !ok {
		t.Fatal("the key was not created")
	}
	if !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("m[new] = %v, want [7]", got)
	}
}

func TestAddNilMap(t *testing.T) {
	Add(nil, "a", 1)
}

func TestAddManyValues(t *testing.T) {
	m := map[string][]int{}
	for i := 0; i < 100; i++ {
		Add(m, "k", i)
	}
	if len(m["k"]) != 100 {
		t.Errorf("len = %d, want 100", len(m["k"]))
	}
	for i, v := range m["k"] {
		if v != i {
			t.Fatalf("m[k][%d] = %d, want %d", i, v, i)
		}
	}
}
""",
    context="A grouping helper appends into a map of slices. It compiles, it runs, and every bucket comes out empty.",
    task=[
        "Append `v` to the slice stored under `key`, creating the entry if absent.",
        "A nil map must not panic.",
        "Fix the single bug so the appended value reaches the map.",
    ],
    examples=[
        ('Add(m, "a", 1); Add(m, "a", 2)', 'm["a"] is [1 2]', None),
        ('Add(m, "new", 7)', "the key is created", "A missing key reads as a nil slice, which append handles."),
        ("Add(nil, \"a\", 1)", "no panic", None),
    ],
    topics=[
        ("Map values are not addressable", "`m[k]` is a copy; there is no slot you can append through."),
        ("append returns a new header", "Its result is the only valid slice afterwards."),
        ("The nil slice appends fine", "A missing key needs no special case."),
    ],
    hint="`append` gives you something back. Where does it go?",
    intuition="Reading `m[key]` copies the slice header out of the map. `append` may reallocate, and even when it does not, the new length lives only in the copy — so unless you store it back, the map keeps the old header.",
    approach=[
        "Guard the nil map.",
        "`m[key] = append(m[key], v)`.",
    ],
    walkthrough="For a missing key, `m[key]` is a nil slice; `append` allocates and returns a one-element slice, which the assignment stores. Without the assignment the new slice is discarded and the key never appears.",
    pitfalls=[
        "Checking `if _, ok := m[key]; !ok` and pre-creating the slice — harmless, and it does not fix the missing assignment.",
        "Assuming a map of pointers would behave the same; there the value is a pointer you can write through.",
    ],
)

P(
    "middle",
    name="arrayptr",
    title="A Pointer To An Array Is Not A Slice",
    sig="func Sum(a *[8]int) int",
    doc="""Sum totals the array a points at.

A pointer to an array carries the length in its type, so it can be
indexed and ranged directly, with no header and no allocation.

Examples:

	Sum(&[8]int{1, 2}) => 3""",
    solution="""if a == nil {
	return 0
}
total := 0
for _, v := range a {
	total += v
}
return total""",
    tests="""
import "testing"

var sink int

func TestSum(t *testing.T) {
	a := [8]int{1, 2, 3}
	if got := Sum(&a); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
	if got := Sum(&[8]int{}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, want 0", got)
	}
}

func TestSumSeesLaterWrites(t *testing.T) {
	a := [8]int{}
	p := &a
	a[0] = 5
	if got := Sum(p); got != 5 {
		t.Errorf("Sum = %d, want 5: the pointer must reach the caller's array", got)
	}
}

func TestSumAllocatesNothing(t *testing.T) {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	if n := testing.AllocsPerRun(200, func() { sink = Sum(&a) }); n != 0 {
		t.Errorf("Sum made %v allocations, want 0", n)
	}
}

func TestSumNegative(t *testing.T) {
	a := [8]int{-4, 4}
	if got := Sum(&a); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}
""",
    context="A fixed-size buffer is passed as `[8]int` and copied on every call. Switching to a slice would work and would also lose the compiler's guarantee that the length is exactly eight.",
    task=[
        "Total the array `a` points at.",
        "A nil pointer totals 0.",
        "Zero allocations; the caller's array must be the one read.",
    ],
    examples=[
        ("Sum(&[8]int{1,2})", "3", None),
        ("a[0] = 5 after taking &a", "Sum sees 5", "The pointer is a live view."),
        ("Sum(nil)", "0", None),
    ],
    topics=[
        ("*[N]T can be ranged", "The length is in the type, so `range a` works on the pointer directly."),
        ("No header, no copy", "One word is passed, and the array is not duplicated."),
        ("Length as a type guarantee", "A `*[8]int` cannot be given a seven-element array."),
    ],
    hint="`for _, v := range a` works on the pointer. No dereference needed.",
    intuition="Go lets a pointer to an array be indexed and ranged as if it were the array. That gives you the cheapness of a pointer with the compile-time length check a slice cannot provide.",
    approach=[
        "Return 0 for a nil pointer.",
        "Range the pointer and accumulate.",
    ],
    walkthrough="Passing `&a` moves one word. `range a` iterates the eight elements in the caller's storage — no copy, no header, no allocation.",
    pitfalls=[
        "`(*a)[i]` works but is noise; the pointer indexes directly.",
        "Forgetting the nil check — ranging a nil array pointer panics.",
    ],
)

P(
    "middle",
    name="cutbytes",
    title="Split Once, Copy Nothing",
    sig="func Cut(s string, sep byte) (before, after string, found bool)",
    doc="""Cut splits s around the first occurrence of sep.

When sep is absent, before is s and after is empty. Both results are
substrings, so nothing is copied.

Examples:

	Cut("a=b", '=') => "a", "b", true""",
    solution="""for i := 0; i < len(s); i++ {
	if s[i] == sep {
		return s[:i], s[i+1:], true
	}
}
return s, "", false""",
    tests="""
import (
	"testing"
	"unsafe"
)

var (
	sinkA, sinkB string
	sinkOK       bool
)

func TestCut(t *testing.T) {
	cases := []struct {
		in             string
		before, after  string
		found          bool
	}{
		{"a=b", "a", "b", true},
		{"a=b=c", "a", "b=c", true},
		{"=x", "", "x", true},
		{"x=", "x", "", true},
		{"abc", "abc", "", false},
		{"", "", "", false},
		{"=", "", "", true},
	}
	for _, c := range cases {
		b, a, ok := Cut(c.in, '=')
		if b != c.before || a != c.after || ok != c.found {
			t.Errorf("Cut(%q) = %q, %q, %v, want %q, %q, %v",
				c.in, b, a, ok, c.before, c.after, c.found)
		}
	}
}

func TestCutResultsAreSubstrings(t *testing.T) {
	s := "key=value"
	before, after, ok := Cut(s, '=')
	if !ok {
		t.Fatal("Cut reported not found")
	}
	if unsafe.StringData(before) != unsafe.StringData(s) {
		t.Error("before is a copy; it must be a substring of s")
	}
	want := (*byte)(unsafe.Add(unsafe.Pointer(unsafe.StringData(s)), 4))
	if unsafe.StringData(after) != want {
		t.Error("after is a copy; it must be a substring of s")
	}
}

func TestCutAllocatesNothing(t *testing.T) {
	s := "a-very-long-configuration-key=and-an-equally-long-value"
	n := testing.AllocsPerRun(200, func() { sinkA, sinkB, sinkOK = Cut(s, '=') })
	if n != 0 {
		t.Errorf("Cut made %v allocations, want 0", n)
	}
}
""",
    context="A config parser splits `key=value` lines with a helper that returns freshly allocated strings. There are thousands of lines and every one of them allocates twice for text it already has.",
    task=[
        "Split `s` around the first `sep`.",
        "When `sep` is absent, return `s`, `\"\"`, false.",
        "Both results must be substrings of `s` — zero allocations.",
    ],
    examples=[
        ('Cut("a=b", \'=\')', '"a", "b", true', None),
        ('Cut("a=b=c", \'=\')', '"a", "b=c", true', "Only the first separator splits."),
        ('Cut("abc", \'=\')', '"abc", "", false', None),
    ],
    topics=[
        ("Substrings share bytes", "Slicing a string is a new header over the same immutable memory."),
        ("First occurrence only", "The loop returns as soon as it finds one."),
        ("The not-found contract", "Returning `s` unchanged lets callers ignore `found` when convenient."),
    ],
    hint="Two slice expressions and an early return.",
    intuition="Strings are immutable, so a piece of one can point straight into it. Splitting is entirely a question of where the boundary is — nothing has to move.",
    approach=[
        "Scan for the first `sep`.",
        "Return `s[:i]`, `s[i+1:]`, true.",
        "After the loop, return `s`, `\"\"`, false.",
    ],
    walkthrough='For "key=value", the separator is at index 3, so the results are `s[:3]` and `s[4:]` — two headers over the original bytes.',
    pitfalls=[
        "Returning `s[i:]` for `after`, which keeps the separator.",
        "Building the halves with concatenation or `[]byte` round-trips.",
    ],
)

P(
    "middle",
    name="fieldkinds",
    title="Describe A Struct's Shape",
    sig="func FieldKinds(v any) []string",
    doc="""FieldKinds returns "Name:kind" for each exported field of v, in
declaration order.

A non-struct, or a nil interface, yields nil.

Examples:

	FieldKinds(row{}) => []string{"ID:int", "Name:string"}""",
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
	out = append(out, f.Name+":"+f.Type.Kind().String())
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

type row struct {
	ID     int
	Name   string
	Tags   []string
	Ratio  float64
	hidden bool
}

func TestFieldKinds(t *testing.T) {
	want := []string{"ID:int", "Name:string", "Tags:slice", "Ratio:float64"}
	if got := FieldKinds(row{}); !reflect.DeepEqual(got, want) {
		t.Errorf("FieldKinds = %v, want %v", got, want)
	}
}

func TestFieldKindsNested(t *testing.T) {
	type outer struct {
		In  row
		Ptr *row
	}
	want := []string{"In:struct", "Ptr:ptr"}
	if got := FieldKinds(outer{}); !reflect.DeepEqual(got, want) {
		t.Errorf("FieldKinds = %v, want %v", got, want)
	}
}

func TestFieldKindsRejectsNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, []int{1}, map[string]int{}, &row{}} {
		if got := FieldKinds(in); got != nil {
			t.Errorf("FieldKinds(%#v) = %v, want nil", in, got)
		}
	}
}

func TestFieldKindsEmptyStruct(t *testing.T) {
	if got := FieldKinds(struct{}{}); got != nil {
		t.Errorf("FieldKinds = %v, want nil", got)
	}
}
""",
    context="A schema documentation generator prints each config struct's fields and their shapes. Keeping the document in step with the code by hand lasted exactly one release.",
    task=[
        "Return `\"Name:kind\"` for each exported field, in declaration order.",
        "Report the field's kind, not its declared type name.",
        "Return nil for a non-struct or a nil interface.",
    ],
    examples=[
        ("FieldKinds(row{})", "[ID:int Name:string Tags:slice Ratio:float64]", "`hidden` is unexported."),
        ("a field of type *row", "Ptr:ptr", "The kind, not the type name."),
        ("FieldKinds(3)", "<nil>", None),
    ],
    topics=[
        ("StructField.Type", "Each field carries its own `reflect.Type`."),
        ("Kind vs Type name", "`[]string` has kind slice; its type string is \"[]string\"."),
        ("Declaration order", "`Field(i)` walks the struct in source order."),
    ],
    hint="`f.Type.Kind().String()` is the right-hand half of each entry.",
    intuition="A struct type is a table the runtime carries: names, types, tags, export status. Rendering it is a loop, and the result stays correct as the struct changes.",
    approach=[
        "Reject a nil or non-struct type.",
        "Loop the fields, skipping unexported ones.",
        "Append `Name + \":\" + Type.Kind().String()`.",
    ],
    walkthrough="`row` has five fields; `hidden` is skipped, and `Tags` reports kind slice rather than the type string \"[]string\".",
    pitfalls=[
        "Using `f.Type.String()`, which gives the type name instead of the kind.",
        "Calling `NumField` before checking the kind.",
    ],
)

P(
    "middle",
    name="padwaste",
    title="How Many Bytes Is The Padding",
    sig="func Waste(v any) uintptr",
    doc="""Waste returns how many bytes of v's struct type are padding: its total
size minus the sum of its fields' sizes.

A non-struct wastes nothing.

Examples:

	Waste(gappy{}) => 14 for a byte, an int64 and a byte""",
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil || t.Kind() != reflect.Struct {
	return 0
}
var used uintptr
for i := 0; i < t.NumField(); i++ {
	used += t.Field(i).Type.Size()
}
return t.Size() - used""",
    tests="""
import (
	"testing"
	"unsafe"
)

type gappy struct {
	A byte
	B int64
	C byte
}

type packed struct {
	B int64
	A byte
	C byte
}

type none struct {
	A int64
	B int64
}

func TestWasteGappy(t *testing.T) {
	var g gappy
	want := unsafe.Sizeof(g) - (unsafe.Sizeof(g.A) + unsafe.Sizeof(g.B) + unsafe.Sizeof(g.C))
	if got := Waste(gappy{}); got != want {
		t.Errorf("Waste = %d, want %d", got, want)
	}
	if got := Waste(gappy{}); got == 0 {
		t.Error("Waste = 0, want a positive number for a badly ordered struct")
	}
}

func TestWastePackedIsSmaller(t *testing.T) {
	if Waste(packed{}) >= Waste(gappy{}) {
		t.Errorf("packed wastes %d and gappy wastes %d, want the packed layout to waste less",
			Waste(packed{}), Waste(gappy{}))
	}
}

func TestWasteNone(t *testing.T) {
	if got := Waste(none{}); got != 0 {
		t.Errorf("Waste = %d, want 0: same-width fields need no padding", got)
	}
}

func TestWasteNonStructs(t *testing.T) {
	for _, in := range []any{nil, 3, "s", []int{1}} {
		if got := Waste(in); got != 0 {
			t.Errorf("Waste(%#v) = %d, want 0", in, got)
		}
	}
}

func TestWasteEmptyStruct(t *testing.T) {
	if got := Waste(struct{}{}); got != 0 {
		t.Errorf("Waste = %d, want 0", got)
	}
}
""",
    context="A team is told to reorder struct fields for memory, and nobody can say which structs are actually wasteful. A number would settle it in a minute.",
    task=[
        "Return the struct's size minus the sum of its field sizes.",
        "Include unexported fields — they occupy space too.",
        "Return 0 for a non-struct or a nil interface.",
    ],
    examples=[
        ("Waste(gappy{})", "14", "byte, int64, byte on a 64-bit build."),
        ("Waste(packed{})", "less than gappy", "Widest field first."),
        ("Waste(none{})", "0", "Two int64s need no padding."),
    ],
    topics=[
        ("Type.Size", "The reflective twin of `unsafe.Sizeof`, available for a type known at run time."),
        ("Padding is size minus content", "Internal gaps and tail padding both show up in this difference."),
        ("Unexported fields count", "They cannot be read, but they still occupy bytes."),
    ],
    hint="Sum the field sizes, subtract from the struct's size.",
    intuition="Padding is not stored anywhere you can query directly — it is the difference between what the fields need and what the struct occupies. Reflection gives you both numbers.",
    approach=[
        "Reject non-structs.",
        "Sum `t.Field(i).Type.Size()` over every field.",
        "Return `t.Size()` minus that sum.",
    ],
    walkthrough="`gappy` needs 1 + 8 + 1 = 10 bytes of fields and occupies 24: seven bytes after `A` and seven after `C`. `packed` occupies 16, wasting 6.",
    pitfalls=[
        "Skipping unexported fields, which undercounts the content and overstates the waste.",
        "Using `Sizeof` on the interface value, which measures the interface header, not the struct.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="popretain",
    title="The Popped Element That Never Left",
    mode="bug",
    sig="func Pop(s []*Job) (*Job, []*Job)",
    doc="""Pop removes and returns the last element of s.

The shortened slice must not keep the popped job reachable: the element
is still in the backing array until it is cleared.

Examples:

	Pop([]*Job{a, b}) => b, a slice holding only a""",
    extra="""// Job is one queued unit of work.
type Job struct {
	ID  int
	Pad [1024]byte
}""",
    buggy="""if len(s) == 0 {
	return nil, s
}
last := s[len(s)-1]
return last, s[:len(s)-1]""",
    solution="""if len(s) == 0 {
	return nil, s
}
i := len(s) - 1
last := s[i]
s[i] = nil
return last, s[:i]""",
    tests="""
import "testing"

func TestPop(t *testing.T) {
	a, b := &Job{ID: 1}, &Job{ID: 2}
	got, rest := Pop([]*Job{a, b})
	if got != b {
		t.Errorf("Pop returned %v, want the last job", got)
	}
	if len(rest) != 1 || rest[0] != a {
		t.Errorf("rest = %v, want [a]", rest)
	}
}

func TestPopEmpty(t *testing.T) {
	got, rest := Pop(nil)
	if got != nil || len(rest) != 0 {
		t.Errorf("Pop(nil) = %v, %v, want nil, empty", got, rest)
	}
}

func TestPopClearsTheSlot(t *testing.T) {
	s := []*Job{{ID: 1}, {ID: 2}}
	_, rest := Pop(s)
	if s[1] != nil {
		t.Error("the popped slot still holds the job: it stays reachable through the backing array")
	}
	if len(rest) != 1 {
		t.Errorf("len = %d, want 1", len(rest))
	}
}

func TestPopRepeatedly(t *testing.T) {
	s := make([]*Job, 8)
	for i := range s {
		s[i] = &Job{ID: i}
	}
	backing := s
	for i := 7; i >= 0; i-- {
		var got *Job
		got, s = Pop(s)
		if got == nil || got.ID != i {
			t.Fatalf("pop %d returned %v", i, got)
		}
	}
	for i, p := range backing[:8] {
		if p != nil {
			t.Fatalf("slot %d still holds job %d after every pop", i, p.ID)
		}
	}
}
""",
    context="A worker pops jobs off a stack that is reused for the process's lifetime. Each job carries a kilobyte of payload, and the heap never comes back down after a burst.",
    task=[
        "Return the last element and the shortened slice.",
        "An empty input returns nil and the slice unchanged.",
        "Fix the single bug so the popped element stops being reachable through the array.",
    ],
    examples=[
        ("Pop([]*Job{a, b})", "b, [a]", None),
        ("s[1] after Pop(s)", "nil", "The vacated slot is cleared."),
        ("Pop(nil)", "nil, []", None),
    ],
    topics=[
        ("Reslicing does not erase", "The element past the new length is still in the array."),
        ("Long-lived containers leak", "A reused stack keeps every slot it ever filled."),
        ("Clear before shortening", "Writing nil is the only way to release the reference."),
    ],
    hint="The returned slice is right. What is still sitting at index `len(s)-1`?",
    intuition="Shortening a slice changes your view, not the memory. As long as the backing array is alive, every pointer still in it is alive too — so a stack that is reused forever retains everything it ever held.",
    approach=[
        "Handle the empty case.",
        "Read the last element, write nil into its slot.",
        "Return the element and `s[:i]`.",
    ],
    walkthrough="Popping eight jobs from an eight-slot stack leaves eight kilobytes reachable before the fix and none after it — the array is still there, but every slot is nil.",
    pitfalls=[
        "Clearing after reslicing, which no longer has access to the slot.",
        "Assuming the collector can free part of an array; it frees allocations, not ranges.",
    ],
)

P(
    "senior",
    name="ptrperitem",
    title="One Allocation, Not One Per Element",
    mode="bug",
    sig="func Build(n int) []*Node",
    doc="""Build returns n nodes, each pointed at by one element of the result.

Allocating each node separately costs n allocations and scatters them
across the heap; one backing array plus n pointers into it costs two.

Examples:

	Build(3) => three nodes with IDs 0, 1, 2""",
    extra="""// Node is one element of the built collection.
type Node struct {
	ID   int
	Next *Node
}""",
    buggy="""if n <= 0 {
	return nil
}
out := make([]*Node, 0, n)
for i := 0; i < n; i++ {
	out = append(out, &Node{ID: i})
}
return out""",
    solution="""if n <= 0 {
	return nil
}
block := make([]Node, n)
out := make([]*Node, n)
for i := 0; i < n; i++ {
	block[i].ID = i
	out[i] = &block[i]
}
return out""",
    tests="""
import "testing"

var sink []*Node

func TestBuild(t *testing.T) {
	got := Build(3)
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for i, p := range got {
		if p == nil || p.ID != i {
			t.Fatalf("got[%d] = %v, want a node with ID %d", i, p, i)
		}
	}
	if got := Build(0); got != nil {
		t.Errorf("Build(0) = %v, want nil", got)
	}
	if got := Build(-1); got != nil {
		t.Errorf("Build(-1) = %v, want nil", got)
	}
}

func TestBuildNodesAreDistinct(t *testing.T) {
	got := Build(4)
	for i := 0; i < len(got); i++ {
		for j := i + 1; j < len(got); j++ {
			if got[i] == got[j] {
				t.Fatalf("nodes %d and %d are the same pointer", i, j)
			}
		}
	}
	got[0].ID = 99
	if got[1].ID != 1 {
		t.Error("writing one node changed another")
	}
}

func TestBuildIsWritable(t *testing.T) {
	got := Build(2)
	got[0].Next = got[1]
	if got[0].Next.ID != 1 {
		t.Error("the nodes are not linkable")
	}
}

func TestBuildAllocationsDoNotScaleWithN(t *testing.T) {
	n := testing.AllocsPerRun(50, func() { sink = Build(256) })
	if n > 4 {
		t.Errorf("Build made %v allocations for 256 nodes, want a handful: allocate one block", n)
	}
}
""",
    context="A graph loader allocates a node per vertex. For a million vertices that is a million small allocations, a million objects for the collector to scan, and a traversal that misses cache on every hop.",
    task=[
        "Return `n` nodes with IDs `0..n-1`, each addressed by one element of the result.",
        "The nodes must be distinct and writable, including linking one to another.",
        "Fix the single bug so the allocation count does not scale with `n`.",
        "`n <= 0` returns nil.",
    ],
    examples=[
        ("Build(3)", "three nodes, IDs 0, 1, 2", None),
        ("256 nodes", "a handful of allocations, not 256", None),
        ("got[0].Next = got[1]", "links fine", "The nodes are ordinary addressable values."),
    ],
    topics=[
        ("Block allocation", "One `[]Node` plus pointers into it replaces n separate objects."),
        ("Elements of a slice are addressable", "`&block[i]` is a normal pointer, valid as long as the array lives."),
        ("Locality", "Contiguous nodes make a traversal cache-friendly."),
        ("Lifetime coupling", "Every node stays alive as long as any pointer into the block does."),
    ],
    hint="The nodes do not have to be allocated one at a time to be pointed at one at a time.",
    intuition="A pointer needs an address, not its own allocation. Carving one array into n addressable elements gives the same API with two allocations, and the collector sees one object instead of n.",
    approach=[
        "Allocate `block := make([]Node, n)` and `out := make([]*Node, n)`.",
        "Set each node's ID and point `out[i]` at `&block[i]`.",
    ],
    walkthrough="Building 256 nodes costs 256 allocations before the fix and 2 after. The trade-off is that the block is freed only when the last pointer into it is gone.",
    pitfalls=[
        "Taking `&node` of a loop variable and expecting distinct nodes — every iteration must address a distinct element.",
        "Applying this when nodes have wildly different lifetimes; the block is all-or-nothing.",
    ],
)

P(
    "senior",
    name="deepequalcost",
    title="A Comparison That Should Not Reflect",
    mode="bug",
    sig="func Changed(a, b Config) bool",
    doc="""Changed reports whether the two configs differ.

Config is a comparable struct, so == does the whole job. Reflecting over
it boxes both operands and walks the fields at run time.

Examples:

	Changed(Config{Retries: 1}, Config{Retries: 2}) => true""",
    imports=['"reflect"'],
    sol_imports=[],
    extra="""// Config is a comparable settings block.
type Config struct {
	Retries int
	Timeout int
	Name    string
	Debug   bool
}""",
    buggy="""return !reflect.DeepEqual(a, b)""",
    solution="""return a != b""",
    tests="""
import "testing"

var sink bool

func TestChanged(t *testing.T) {
	base := Config{Retries: 1, Timeout: 2, Name: "n", Debug: true}
	if Changed(base, base) {
		t.Error("Changed = true for identical configs, want false")
	}
	other := base
	other.Retries = 9
	if !Changed(base, other) {
		t.Error("Changed = false for differing configs, want true")
	}
}

func TestChangedEveryField(t *testing.T) {
	base := Config{}
	cases := []Config{
		{Retries: 1},
		{Timeout: 1},
		{Name: "x"},
		{Debug: true},
	}
	for i, c := range cases {
		if !Changed(base, c) {
			t.Errorf("case %d: Changed = false, want true", i)
		}
	}
}

func TestChangedZeroValues(t *testing.T) {
	if Changed(Config{}, Config{}) {
		t.Error("Changed = true for two zero configs, want false")
	}
}

func TestChangedAllocatesNothing(t *testing.T) {
	a := Config{Retries: 1, Name: "left"}
	b := Config{Retries: 2, Name: "right"}
	if n := testing.AllocsPerRun(200, func() { sink = Changed(a, b) }); n != 0 {
		t.Errorf("Changed made %v allocations, want 0: the struct is comparable", n)
	}
}
""",
    context="A config watcher compares the old and new settings on every poll. `reflect.DeepEqual` was the obvious spelling, and the watcher now allocates twice a second forever.",
    task=[
        "Report whether the two configs differ.",
        "Fix the single bug so the comparison allocates nothing.",
        "Every field must participate in the comparison.",
    ],
    examples=[
        ("Changed(Config{Retries:1}, Config{Retries:2})", "true", None),
        ("Changed(Config{}, Config{})", "false", None),
        ("allocations per call", "0", "`==` compiles to a direct comparison."),
    ],
    topics=[
        ("Comparable structs", "A struct of comparable fields supports `==` field by field."),
        ("DeepEqual boxes its arguments", "Both operands become `any`, which allocates."),
        ("Reflection is for unknown types", "Here the type is known at compile time."),
        ("When DeepEqual is still right", "Slices, maps and functions are not comparable with `==`."),
    ],
    hint="The struct has no slices, maps or functions in it. What does that make it?",
    intuition="`DeepEqual` exists for types `==` cannot handle. Reaching for it on a comparable struct pays for a run-time type walk and two boxes to answer a question the compiler could have answered directly.",
    approach=[
        "Return `a != b`.",
    ],
    walkthrough="`Config` holds two ints, a string and a bool — all comparable, so `!=` compares them inline. `DeepEqual` boxes both structs and walks four fields reflectively on every poll.",
    pitfalls=[
        "Adding a slice field later, which makes the struct non-comparable and `==` a compile error — that is the signal to reconsider, not to reach back for DeepEqual.",
        "Comparing field by field manually, which is correct and drifts as fields are added.",
    ],
)

P(
    "senior",
    name="concatexact",
    title="Join With Exactly One Allocation",
    sig="func Concat(parts []string) string",
    doc="""Concat returns the parts joined end to end.

The final length is the sum of the parts' lengths, so the result can be
built in one allocation and handed out without a second copy.

Examples:

	Concat([]string{"a", "bc"}) => "abc" """,
    imports=['"unsafe"'],
    solution="""n := 0
for _, p := range parts {
	n += len(p)
}
if n == 0 {
	return ""
}
buf := make([]byte, 0, n)
for _, p := range parts {
	buf = append(buf, p...)
}
return unsafe.String(unsafe.SliceData(buf), len(buf))""",
    tests="""
import "testing"

var sink string

func TestConcat(t *testing.T) {
	if got := Concat([]string{"a", "bc", "d"}); got != "abcd" {
		t.Errorf("Concat = %q, want \\"abcd\\"", got)
	}
	if got := Concat(nil); got != "" {
		t.Errorf("Concat = %q, want empty", got)
	}
	if got := Concat([]string{"", ""}); got != "" {
		t.Errorf("Concat = %q, want empty", got)
	}
	if got := Concat([]string{"solo"}); got != "solo" {
		t.Errorf("Concat = %q, want \\"solo\\"", got)
	}
}

func TestConcatLong(t *testing.T) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "chunk"
	}
	got := Concat(parts)
	if len(got) != 500 {
		t.Fatalf("len = %d, want 500", len(got))
	}
	for i := 0; i < len(got); i += 5 {
		if got[i:i+5] != "chunk" {
			t.Fatalf("at %d: %q, want \\"chunk\\"", i, got[i:i+5])
		}
	}
}

func TestConcatAllocatesOnce(t *testing.T) {
	parts := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	if n := testing.AllocsPerRun(200, func() { sink = Concat(parts) }); n > 1 {
		t.Errorf("Concat made %v allocations, want 1: size the buffer, then wrap it", n)
	}
}

func TestConcatResultIsIndependent(t *testing.T) {
	parts := []string{"ab", "cd"}
	first := Concat(parts)
	parts[0] = "XY"
	if first != "abcd" {
		t.Errorf("first = %q, want \\"abcd\\"", first)
	}
}
""",
    context="A hot path joins five short strings per request. `strings.Join` copies once into its buffer and `string(buf)` copies again — the second copy is pure waste on a buffer nothing else can see.",
    task=[
        "Return the parts joined end to end.",
        "Exactly one allocation: size the buffer from the total length, then hand it out without copying again.",
        "An empty input, or all-empty parts, returns the empty string.",
    ],
    examples=[
        ('Concat([]string{"a","bc","d"})', '"abcd"', None),
        ("Concat(nil)", '""', None),
        ("allocations per call", "1", "The buffer, and nothing else."),
    ],
    topics=[
        ("Exact sizing", "Summing the lengths removes every growth step."),
        ("unsafe.String over your own buffer", "Legal precisely because nothing else can write to it."),
        ("string(buf) is the second copy", "It is required when the buffer is shared, and wasted when it is not."),
        ("Local ownership", "The buffer never leaves the function except as the string."),
    ],
    hint="Build it in a `[]byte` you allocated, then stop copying.",
    intuition="The rule for `unsafe.String` is that the bytes must never change. A buffer allocated inside the function, written once, and never referenced again satisfies that by construction — so the final copy buys nothing.",
    approach=[
        "Sum the parts' lengths; return `\"\"` for zero.",
        "`make([]byte, 0, n)` and append every part.",
        "Wrap it with `unsafe.String`.",
    ],
    walkthrough="Joining five parts totalling 27 bytes allocates one 27-byte buffer. `string(buf)` would allocate a second 27 bytes and copy — for bytes only this function has ever seen.",
    pitfalls=[
        "Keeping a reference to `buf` after wrapping it; any later write breaks the string's immutability.",
        "Skipping the zero-length guard, which would wrap a nil data pointer.",
    ],
)

P(
    "senior",
    name="flushwriter",
    title="The Buffered Writer Nobody Flushed",
    mode="bug",
    sig="func WriteAll(w io.Writer, lines []string) error",
    doc="""WriteAll writes each line followed by '\\n' through a buffered writer.

A buffered writer holds the tail of the output until it is flushed; the
last partial buffer is lost otherwise.

Examples:

	WriteAll(&buf, []string{"a"}) => buf holds "a\\n" """,
    imports=['"bufio"', '"io"'],
    buggy="""bw := bufio.NewWriter(w)
for _, l := range lines {
	if _, err := bw.WriteString(l); err != nil {
		return err
	}
	if err := bw.WriteByte('\\n'); err != nil {
		return err
	}
}
return nil""",
    solution="""bw := bufio.NewWriter(w)
for _, l := range lines {
	if _, err := bw.WriteString(l); err != nil {
		return err
	}
	if err := bw.WriteByte('\\n'); err != nil {
		return err
	}
}
return bw.Flush()""",
    tests="""
import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestWriteAll(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteAll(&buf, []string{"a", "b"}); err != nil {
		t.Fatal(err)
	}
	if got := buf.String(); got != "a\\nb\\n" {
		t.Errorf("buf = %q, want \\"a\\\\nb\\\\n\\"", got)
	}
}

func TestWriteAllEmpty(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteAll(&buf, nil); err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 0 {
		t.Errorf("buf = %q, want empty", buf.String())
	}
}

func TestWriteAllLargeOutput(t *testing.T) {
	var buf bytes.Buffer
	lines := make([]string, 500)
	for i := range lines {
		lines[i] = strings.Repeat("x", 40)
	}
	if err := WriteAll(&buf, lines); err != nil {
		t.Fatal(err)
	}
	if got := buf.Len(); got != 500*41 {
		t.Errorf("wrote %d bytes, want %d: the last buffer was never flushed", got, 500*41)
	}
}

type failWriter struct{}

func (failWriter) Write(p []byte) (int, error) { return 0, errors.New("boom") }

func TestWriteAllPropagatesErrors(t *testing.T) {
	lines := make([]string, 500)
	for i := range lines {
		lines[i] = strings.Repeat("y", 100)
	}
	if err := WriteAll(failWriter{}, lines); err == nil {
		t.Error("want the writer's error, got nil")
	}
}
""",
    context="An export job writes a million rows and the file is short by a few kilobytes. The truncation is at the end, it varies run to run, and nothing reports an error.",
    task=[
        "Write each line followed by a newline through a buffered writer.",
        "Fix the single bug so no output is lost.",
        "Return any error, including one that only surfaces at the end.",
    ],
    examples=[
        ('WriteAll(&buf, []string{"a","b"})', '"a\\nb\\n"', None),
        ("500 lines of 40 bytes", "20500 bytes written", "The final partial buffer must reach the writer."),
        ("a failing writer", "an error", None),
    ],
    topics=[
        ("Buffered writers hold output", "Bytes stay in the buffer until it fills or is flushed."),
        ("Flush is where errors surface", "A write into the buffer succeeds even when the underlying writer will not."),
        ("Silent truncation", "The missing bytes are always the last ones, which is why it looks random."),
    ],
    hint="Every byte written so far is accounted for. What about the ones still in the buffer?",
    intuition="A buffered writer trades immediate writes for batched ones, which means the last batch only leaves when you say so. It is also the point at which the underlying writer's failures finally become visible.",
    approach=[
        "Write the lines as before.",
        "Return `bw.Flush()` instead of nil.",
    ],
    walkthrough="500 lines of 41 bytes is 20500, and the default buffer is 4096 — so 20480 bytes leave on their own and the last 20 sit in the buffer forever without the flush.",
    pitfalls=[
        "`defer bw.Flush()` — it flushes, and it discards the error.",
        "Assuming a small output is safe; a small output is entirely inside the buffer.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="copyonwrite",
    title="Publish A New Map, Never Edit The Old One",
    sig="func (s *Store) Set(key string, val int)",
    doc="""Set publishes a new snapshot of the map with key set to val.

Readers hold whatever snapshot was current when they loaded it, so a
published map must never be modified again: build a copy, then swap.

Examples:

	s.Set("a", 1); s.Get("a") => 1, true""",
    imports=['"sync"', '"sync/atomic"'],
    extra="""// Store is a read-mostly map published by pointer swap.
type Store struct {
	mu sync.Mutex // serialises writers only
	m  atomic.Pointer[map[string]int]
}

// Get reads from the current snapshot without locking.
func (s *Store) Get(key string) (int, bool) {
	p := s.m.Load()
	if p == nil {
		return 0, false
	}
	v, ok := (*p)[key]
	return v, ok
}

// Len reports the current snapshot's size.
func (s *Store) Len() int {
	p := s.m.Load()
	if p == nil {
		return 0
	}
	return len(*p)
}""",
    solution="""s.mu.Lock()
defer s.mu.Unlock()

var old map[string]int
if p := s.m.Load(); p != nil {
	old = *p
}
next := make(map[string]int, len(old)+1)
for k, v := range old {
	next[k] = v
}
next[key] = val
s.m.Store(&next)""",
    tests="""
import (
	"sync"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	var s Store
	if _, ok := s.Get("a"); ok {
		t.Error("Get on an empty store reported ok, want false")
	}
	s.Set("a", 1)
	s.Set("b", 2)
	if v, ok := s.Get("a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := s.Get("b"); !ok || v != 2 {
		t.Errorf("Get(b) = %d, %v, want 2, true", v, ok)
	}
	if s.Len() != 2 {
		t.Errorf("Len = %d, want 2", s.Len())
	}
}

func TestSetOverwrites(t *testing.T) {
	var s Store
	s.Set("a", 1)
	s.Set("a", 9)
	if v, _ := s.Get("a"); v != 9 {
		t.Errorf("Get(a) = %d, want 9", v)
	}
	if s.Len() != 1 {
		t.Errorf("Len = %d, want 1", s.Len())
	}
}

func TestOldSnapshotsAreImmutable(t *testing.T) {
	var s Store
	s.Set("a", 1)
	before := s.m.Load()
	s.Set("b", 2)
	if len(*before) != 1 {
		t.Errorf("the previous snapshot grew to %d entries: it was modified in place", len(*before))
	}
	if _, ok := (*before)["b"]; ok {
		t.Error("the previous snapshot gained a key")
	}
}

func TestConcurrentReadersAndWriters(t *testing.T) {
	var s Store
	s.Set("seed", 0)
	var wg sync.WaitGroup
	stop := make(chan struct{})

	for r := 0; r < 8; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				if v, ok := s.Get("seed"); !ok || v != 0 {
					panic("seed changed")
				}
			}
		}()
	}

	var wwg sync.WaitGroup
	for w := 0; w < 4; w++ {
		wwg.Add(1)
		go func(w int) {
			defer wwg.Done()
			for i := 0; i < 200; i++ {
				s.Set(string(rune('a'+w)), i)
			}
		}(w)
	}
	wwg.Wait()
	close(stop)
	wg.Wait()

	if s.Len() != 5 {
		t.Errorf("Len = %d, want 5", s.Len())
	}
}
""",
    context="A read-mostly routing table is guarded by an RWMutex. Reads dominate by a thousand to one, and the read lock's atomic traffic is the top entry in the profile.",
    task=[
        "Publish a new snapshot with `key` set to `val`.",
        "A published map must never be modified — readers may still be holding it.",
        "Serialise writers with the mutex; readers take no lock.",
        "Correct with concurrent readers and writers.",
    ],
    examples=[
        ('s.Set("a", 1); s.Get("a")', "1, true", None),
        ("a snapshot loaded before a Set", "unchanged afterwards", None),
        ("4 writers x 200 sets, 8 readers", "no torn or lost state", None),
    ],
    topics=[
        ("Copy-on-write", "Writers pay a full copy so readers pay nothing."),
        ("Atomic pointer publication", "One word swap replaces the whole map indivisibly."),
        ("Immutable after publication", "The published map is shared with every reader, forever."),
        ("The writer lock", "It serialises the read-copy-swap sequence, which is not atomic on its own."),
    ],
    hint="Load, copy, modify the copy, store. The lock is around all four.",
    intuition="When reads vastly outnumber writes, moving the cost to the writer is a good trade. The invariant that makes it safe is absolute: once a map is published, it is frozen, because you can never know who is still reading it.",
    approach=[
        "Take the writer mutex.",
        "Load the current snapshot, copy it into a new map sized `len(old)+1`.",
        "Set the key in the copy and `Store` its address.",
    ],
    walkthrough="A reader that loaded the old pointer keeps reading the old map, which nobody touches. The next reader loads the new pointer. There is no moment at which any map is both published and being written.",
    pitfalls=[
        "Writing into the loaded map before swapping, which is a data race with every current reader.",
        "Dropping the writer lock, which lets two writers each copy the same old map and lose one update.",
        "Using this shape for a write-heavy map, where copying per write is far worse than a mutex.",
    ],
)

P(
    "staff",
    name="splitviews",
    title="Fields As Views, Into The Caller's Slice",
    sig="func Fields(dst [][]byte, line []byte, sep byte) [][]byte",
    doc="""Fields appends each sep-separated field of line to dst as a view and
returns the extended slice.

The fields share line's storage, and dst lets the caller reuse the header
slice, so a steady-state call allocates nothing at all.

Examples:

	Fields(nil, []byte("a,b"), ',') => [][]byte{"a", "b"}""",
    solution="""if len(line) == 0 {
	return dst
}
start := 0
for i := 0; i <= len(line); i++ {
	if i < len(line) && line[i] != sep {
		continue
	}
	dst = append(dst, line[start:i:i])
	start = i + 1
}
return dst""",
    tests="""
import (
	"bytes"
	"testing"
)

var sink [][]byte

func TestFields(t *testing.T) {
	got := Fields(nil, []byte("a,bb,c"), ',')
	want := [][]byte{[]byte("a"), []byte("bb"), []byte("c")}
	if len(got) != len(want) {
		t.Fatalf("got %d fields, want %d", len(got), len(want))
	}
	for i := range want {
		if !bytes.Equal(got[i], want[i]) {
			t.Errorf("field %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFieldsEdges(t *testing.T) {
	if got := Fields(nil, nil, ','); len(got) != 0 {
		t.Errorf("Fields(nil) = %q, want empty", got)
	}
	got := Fields(nil, []byte("a,,b"), ',')
	if len(got) != 3 || len(got[1]) != 0 {
		t.Errorf("Fields = %q, want three fields with an empty middle", got)
	}
	got = Fields(nil, []byte(",x"), ',')
	if len(got) != 2 || len(got[0]) != 0 || !bytes.Equal(got[1], []byte("x")) {
		t.Errorf("Fields = %q, want an empty field then x", got)
	}
	got = Fields(nil, []byte("x,"), ',')
	if len(got) != 2 || len(got[1]) != 0 {
		t.Errorf("Fields = %q, want x then an empty field", got)
	}
}

func TestFieldsAppendsToDst(t *testing.T) {
	dst := [][]byte{[]byte("keep")}
	got := Fields(dst, []byte("a"), ',')
	if len(got) != 2 || !bytes.Equal(got[0], []byte("keep")) {
		t.Errorf("Fields = %q, want [keep a]", got)
	}
}

func TestFieldsAreViews(t *testing.T) {
	line := []byte("ab,cd")
	got := Fields(nil, line, ',')
	got[0][0] = 'X'
	if line[0] != 'X' {
		t.Error("the fields copied the bytes; they must be views into line")
	}
}

func TestFieldsCapAtTheBoundary(t *testing.T) {
	line := []byte("ab,cd")
	got := Fields(nil, line, ',')
	got[0] = append(got[0], 'Z')
	if line[2] == 'Z' {
		t.Error("appending to a field wrote into the next one: cap each view at its own end")
	}
	if string(line) != "ab,cd" {
		t.Errorf("line = %q, want \\"ab,cd\\"", line)
	}
}

func TestFieldsSteadyStateAllocatesNothing(t *testing.T) {
	line := bytes.Repeat([]byte("field,"), 32)
	line = line[:len(line)-1]
	dst := make([][]byte, 0, 64)
	Fields(dst[:0], line, ',')
	if n := testing.AllocsPerRun(100, func() { sink = Fields(dst[:0], line, ',') }); n != 0 {
		t.Errorf("Fields made %v allocations, want 0 when dst has room", n)
	}
}
""",
    context="A record splitter allocates a fresh `[][]byte` and copies every field on every line. At a million lines a second it is the whole cost of the parser.",
    task=[
        "Append each `sep`-separated field of `line` to `dst` as a view, and return the extended slice.",
        "Empty fields count, including leading and trailing ones; an empty line adds nothing.",
        "Each view's capacity must stop at its own end, so appending to one cannot reach the next.",
        "With room in `dst`, allocate nothing.",
    ],
    examples=[
        ('Fields(nil, []byte("a,bb,c"), \',\')', '["a" "bb" "c"]', None),
        ('Fields(nil, []byte("a,,b"), \',\')', '["a" "" "b"]', "The empty middle field is a field."),
        ("appending to field 0", "does not touch field 1", "Each view is capacity-capped."),
    ],
    topics=[
        ("Views instead of copies", "The fields point into the caller's line."),
        ("Three-index slicing", "`line[start:i:i]` stops an append from spilling into the next field."),
        ("Append-style APIs", "A `dst` parameter lets the caller own the header slice across lines."),
        ("Virtual trailing separator", "Running to `len(line)` inclusive closes the last field."),
    ],
    hint="Two things per field: where it starts and ends, and what its capacity should be.",
    intuition="Splitting moves no data — it only decides boundaries. The only allocation a splitter needs is the slice of headers, and even that can be the caller's if the API asks for it.",
    approach=[
        "Return `dst` unchanged for an empty line.",
        "Walk `i` to `len(line)` inclusive, treating the end as a separator.",
        "Append `line[start:i:i]` and move `start` past the separator.",
    ],
    walkthrough='For "ab,cd", the boundaries are at 2 and 5. The views are `line[0:2:2]` and `line[3:5:5]`, so appending to the first must reallocate rather than overwrite the comma.',
    pitfalls=[
        "`line[start:i]` without the third index — the field's capacity runs to the end of the line.",
        "Stopping at `len(line)-1`, which drops the last field.",
        "Treating an empty line as one empty field; the spec adds nothing.",
    ],
)

P(
    "staff",
    name="encodecache",
    title="Encode Any Struct, Resolve It Once",
    sig="func Encode(dst []byte, v any) ([]byte, error)",
    doc="""Encode appends "name=value" for each exported string field of v to dst,
separated by '&', and returns the extended slice.

The per-type field list is resolved once and cached, so repeated
encodings of a known type cost a map lookup and some appends.

Examples:

	Encode(nil, user{Name: "a"}) => []byte("Name=a")""",
    imports=['"errors"', '"reflect"', '"sync"'],
    extra="""// ErrKind is returned when v is not a struct.
var ErrKind = errors.New("v must be a struct")

// layouts caches each struct type's exported string field indices.
var layouts sync.Map // reflect.Type -> []fieldRef

type fieldRef struct {
	name  string
	index int
}

// layoutOf resolves and caches the encodable fields of t.
func layoutOf(t reflect.Type) []fieldRef {
	if v, ok := layouts.Load(t); ok {
		return v.([]fieldRef)
	}
	refs := make([]fieldRef, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.IsExported() && f.Type.Kind() == reflect.String {
			refs = append(refs, fieldRef{name: f.Name, index: i})
		}
	}
	actual, _ := layouts.LoadOrStore(t, refs)
	return actual.([]fieldRef)
}""",
    solution="""rv := reflect.ValueOf(v)
if !rv.IsValid() || rv.Kind() != reflect.Struct {
	return dst, ErrKind
}
refs := layoutOf(rv.Type())
for i, r := range refs {
	if i > 0 {
		dst = append(dst, '&')
	}
	dst = append(dst, r.name...)
	dst = append(dst, '=')
	dst = append(dst, rv.Field(r.index).String()...)
}
return dst, nil""",
    tests="""
import (
	"bytes"
	"errors"
	"sync"
	"testing"
)

type user struct {
	Name   string
	Email  string
	Age    int
	hidden string
}

func TestEncode(t *testing.T) {
	got, err := Encode(nil, user{Name: "ann", Email: "a@b", Age: 3, hidden: "x"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, []byte("Name=ann&Email=a@b")) {
		t.Errorf("Encode = %q, want \\"Name=ann&Email=a@b\\"", got)
	}
}

func TestEncodeAppendsToDst(t *testing.T) {
	got, err := Encode([]byte("pre:"), user{Name: "x"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, []byte("pre:Name=x&Email=")) {
		t.Errorf("Encode = %q, want \\"pre:Name=x&Email=\\"", got)
	}
}

func TestEncodeNoStringFields(t *testing.T) {
	type nums struct{ A, B int }
	got, err := Encode(nil, nums{1, 2})
	if err != nil || len(got) != 0 {
		t.Errorf("Encode = %q, %v, want empty, nil", got, err)
	}
}

func TestEncodeBadKind(t *testing.T) {
	for _, v := range []any{nil, 3, []int{1}, &user{}} {
		if _, err := Encode(nil, v); !errors.Is(err, ErrKind) {
			t.Errorf("Encode(%#v) = %v, want ErrKind", v, err)
		}
	}
}

func TestEncodeConcurrent(t *testing.T) {
	const workers = 16
	var wg sync.WaitGroup
	wg.Add(workers)
	errs := make([]error, workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			u := user{Name: string(rune('a' + w)), Email: "e"}
			want := []byte("Name=" + string(rune('a'+w)) + "&Email=e")
			buf := make([]byte, 0, 64)
			for i := 0; i < 200; i++ {
				got, err := Encode(buf[:0], u)
				if err != nil {
					errs[w] = err
					return
				}
				if !bytes.Equal(got, want) {
					errs[w] = errors.New("wrong output: " + string(got))
					return
				}
			}
		}(w)
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestEncodeSteadyStateAllocatesNothing(t *testing.T) {
	u := user{Name: "ann", Email: "a@b"}
	buf := make([]byte, 0, 128)
	Encode(buf[:0], u)
	var sink []byte
	n := testing.AllocsPerRun(200, func() { sink, _ = Encode(buf[:0], u) })
	_ = sink
	if n != 0 {
		t.Errorf("Encode made %v allocations, want 0: use the cached layout and the caller's buffer", n)
	}
}
""",
    context="A form encoder walks the struct's field table on every call and builds a string per field. It is the slowest and the most allocating step in a request that does nothing else interesting.",
    task=[
        "Append `name=value` for each exported string field, separated by `&`.",
        "Use the cached layout so the field table is walked once per type.",
        "Append into `dst`; with room, the call must allocate nothing.",
        "Return `ErrKind` for anything that is not a struct.",
    ],
    examples=[
        ('Encode(nil, user{Name:"ann", Email:"a@b"})', '"Name=ann&Email=a@b"', None),
        ('Encode([]byte("pre:"), user{Name:"x"})', '"pre:Name=x&Email="', "dst is extended."),
        ("16 goroutines x 200 encodes", "every result correct", None),
    ],
    topics=[
        ("Cache the type, not the value", "Field indices are per type and never change."),
        ("Append everything", "Writing into the caller's buffer removes the per-field string."),
        ("sync.Map for a read-mostly cache", "Loads after the first write take no lock."),
        ("Value.String without boxing", "It reads a string field directly, unlike `Interface()`."),
    ],
    hint="The cache is written for you. Validate, fetch the layout, append.",
    intuition="Reflection is expensive where it inspects and cheap where it accesses. Resolving the layout once moves all the inspection to the first call, leaving a loop of appends that costs about what hand-written code would.",
    approach=[
        "Reject a non-struct.",
        "Fetch the cached field references for the type.",
        "For each, append the separator, the name, `=`, and the field's string value.",
    ],
    walkthrough="The first encode of `user` walks four fields and stores two references. Every later one — from any goroutine — does a `sync.Map` load and four appends into the caller's buffer.",
    pitfalls=[
        "Ranging the struct's fields instead of the cached layout, which restores the per-call walk.",
        "`rv.Field(i).Interface().(string)`, which boxes and allocates on every field.",
        "Caching a `reflect.Value`; Values are bound to one variable, indices are not.",
    ],
)

P(
    "staff",
    name="singleflight",
    title="One Miss, One Fetch, Many Waiters",
    sig="func (g *Group) Do(key string, fn func() int) int",
    doc="""Do runs fn for key and returns its result, sharing one in-flight call
among every concurrent caller for that key.

A thundering herd on a cold cache must not become N identical fetches,
each allocating its own result.

Examples:

	g.Do("a", expensive) from 32 goroutines => expensive runs once""",
    imports=['"sync"'],
    extra="""// call is one in-flight or completed fetch.
type call struct {
	wg  sync.WaitGroup
	val int
}

// Group deduplicates concurrent calls by key.
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}""",
    solution="""g.mu.Lock()
if g.m == nil {
	g.m = make(map[string]*call)
}
if c, ok := g.m[key]; ok {
	g.mu.Unlock()
	c.wg.Wait()
	return c.val
}
c := new(call)
c.wg.Add(1)
g.m[key] = c
g.mu.Unlock()

c.val = fn()
c.wg.Done()

g.mu.Lock()
delete(g.m, key)
g.mu.Unlock()

return c.val""",
    tests="""
import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDoReturnsTheResult(t *testing.T) {
	var g Group
	if got := g.Do("a", func() int { return 42 }); got != 42 {
		t.Errorf("Do = %d, want 42", got)
	}
}

func TestDoRunsAgainAfterCompletion(t *testing.T) {
	var g Group
	var calls atomic.Int64
	fn := func() int { calls.Add(1); return 1 }
	g.Do("a", fn)
	g.Do("a", fn)
	if got := calls.Load(); got != 2 {
		t.Errorf("fn ran %d times, want 2: Do must not cache across completed calls", got)
	}
}

func TestDoDeduplicatesConcurrentCallers(t *testing.T) {
	var g Group
	var calls atomic.Int64
	release := make(chan struct{})
	fn := func() int {
		calls.Add(1)
		<-release
		return 7
	}

	const workers = 32
	var wg sync.WaitGroup
	results := make([]int, workers)
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wg.Done()
			results[i] = g.Do("hot", fn)
		}(i)
	}

	// let every goroutine reach Do before the fetch completes
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()

	if got := calls.Load(); got != 1 {
		t.Errorf("fn ran %d times, want 1: the callers were not deduplicated", got)
	}
	for i, v := range results {
		if v != 7 {
			t.Fatalf("worker %d got %d, want 7", i, v)
		}
	}
}

func TestDoDifferentKeysRunInParallel(t *testing.T) {
	var g Group
	var calls atomic.Int64
	var wg sync.WaitGroup
	const keys = 8
	wg.Add(keys)
	for i := 0; i < keys; i++ {
		go func(i int) {
			defer wg.Done()
			g.Do(string(rune('a'+i)), func() int { calls.Add(1); return i })
		}(i)
	}
	wg.Wait()
	if got := calls.Load(); got != keys {
		t.Errorf("fn ran %d times, want %d: distinct keys must not block each other", got, keys)
	}
}

func TestDoMixedKeysUnderLoad(t *testing.T) {
	var g Group
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				key := string(rune('a' + i%4))
				want := i % 4
				if got := g.Do(key, func() int { return want }); got != want {
					panic("wrong value for key")
				}
			}
		}(w)
	}
	wg.Wait()
}
""",
    context="A cache entry expires and two hundred requests miss it at once. Two hundred identical database queries run, each building the same multi-megabyte result, and the database falls over.",
    task=[
        "Run `fn` for `key` and return its result.",
        "Concurrent callers for the same key must share one execution and all receive its result.",
        "Different keys must not block each other.",
        "After a call completes, the next `Do` for that key runs `fn` again — this is deduplication, not caching.",
    ],
    examples=[
        ('32 goroutines calling g.Do("hot", fn)', "fn runs once, all get its result", None),
        ("two sequential Do calls", "fn runs twice", "No caching across completed calls."),
        ("8 goroutines on 8 keys", "fn runs 8 times", None),
    ],
    topics=[
        ("Deduplication vs caching", "One shares an in-flight call; the other stores a finished result."),
        ("WaitGroup as a one-shot broadcast", "Waiters block until the owner calls Done."),
        ("Lock scope", "The map is guarded; `fn` runs outside the lock or every key serialises."),
        ("Cleanup after completion", "Removing the entry is what makes the next call run again."),
    ],
    hint="The first caller for a key owns the call. Everyone else waits on it and reads its result.",
    intuition="The herd is a coordination problem, not a caching one: the work is already happening, and the other callers just need to be told when it is done. A per-key `WaitGroup` is exactly that signal.",
    approach=[
        "Under the lock, look the key up. If a call exists, unlock, `Wait`, return its value.",
        "Otherwise create the call, `Add(1)`, register it, and unlock.",
        "Run `fn`, store the result, `Done`.",
        "Take the lock again to delete the entry, then return.",
    ],
    walkthrough="Of 32 callers, one finds no entry and becomes the owner; the other 31 find it, release the lock and block in `Wait`. When the owner finishes, all 32 return the same value and the entry is removed.",
    pitfalls=[
        "Holding `g.mu` across `fn()`, which serialises every key and defeats the purpose.",
        "Writing `c.val` after `Done`, so waiters can read it before it is set.",
        "Never deleting the entry, which silently turns this into a cache that never expires.",
    ],
)

P(
    "staff",
    name="fanin",
    title="Merge Channels Without Leaving Goroutines Behind",
    sig="func Merge(done <-chan struct{}, ins ...<-chan int) <-chan int",
    doc="""Merge returns a channel carrying every value from ins, closed once all
inputs are drained or done is closed.

Every goroutine Merge starts must exit: an abandoned consumer must not
leave forwarders blocked on a send forever.

Examples:

	Merge(done, a, b) => a channel with everything from a and b""",
    imports=['"sync"'],
    solution="""out := make(chan int)
var wg sync.WaitGroup
wg.Add(len(ins))
for _, in := range ins {
	go func(in <-chan int) {
		defer wg.Done()
		for v := range in {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
	}(in)
}
go func() {
	wg.Wait()
	close(out)
}()
return out""",
    tests="""
import (
	"runtime"
	"testing"
	"time"
)

func send(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestMergeDeliversEverything(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Merge(done, send(1, 2), send(3), send())
	sum := 0
	count := 0
	for v := range out {
		sum += v
		count++
	}
	if count != 3 || sum != 6 {
		t.Errorf("got %d values summing to %d, want 3 and 6", count, sum)
	}
}

func TestMergeClosesWhenInputsDrain(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Merge(done, send(1))
	<-out
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced an extra value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed after the inputs drained")
	}
}

func TestMergeNoInputs(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Merge(done)
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced a value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed with no inputs")
	}
}

func TestMergeAbandonedConsumerDoesNotLeak(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()

	for round := 0; round < 20; round++ {
		done := make(chan struct{})
		ins := make([]<-chan int, 4)
		for i := range ins {
			ch := make(chan int)
			go func(ch chan int, i int) {
				for j := 0; ; j++ {
					select {
					case ch <- i*100 + j:
					case <-done:
						close(ch)
						return
					}
				}
			}(ch, i)
			ins[i] = ch
		}
		out := Merge(done, ins...)
		<-out // take one value, then walk away
		close(done)
	}

	deadline := time.Now().Add(3 * time.Second)
	for runtime.NumGoroutine() > base+4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, want about %d: forwarders are still blocked on the send", got, base)
	}
}
""",
    context="A fan-in helper merges worker channels. The consumer often stops early on the first error, and the process accumulates blocked forwarder goroutines until it is restarted.",
    task=[
        "Return a channel carrying every value from `ins`.",
        "Close it once every input is drained.",
        "Every goroutine must exit when `done` is closed, even if the consumer has stopped reading.",
        "Zero inputs closes the output immediately.",
    ],
    examples=[
        ("Merge(done, send(1,2), send(3))", "a channel yielding 1, 2 and 3", None),
        ("inputs drained", "out is closed", None),
        ("consumer abandons out, done closed", "no goroutine left behind", None),
    ],
    topics=[
        ("A blocked send is a live goroutine", "It holds its stack and everything its frame references."),
        ("select with a cancellation channel", "Makes every send abandonable."),
        ("WaitGroup then close", "The closer must run after all forwarders, in its own goroutine."),
        ("Only the sender closes", "The forwarders send; the extra goroutine closes."),
    ],
    hint="Each forwarder's send is the thing that can block forever. What else should it be able to do?",
    intuition="A goroutine blocked on a send to a channel nobody reads never returns and never gets collected. Every send in a fan-in has to be paired with an escape route, and the output can only be closed once all of them are gone.",
    approach=[
        "Start one forwarder per input, tracked by a `WaitGroup`.",
        "Forward with `select` over the send and `<-done`.",
        "In a separate goroutine, `Wait` then `close(out)`.",
    ],
    walkthrough="With four inputs and a consumer that reads once and leaves, closing `done` lets all four forwarders return from their `select`; the closer's `Wait` then returns and `out` is closed.",
    pitfalls=[
        "Calling `wg.Wait()` in `Merge` itself, which blocks until the output is fully consumed.",
        "Closing `out` in the forwarders, which panics on the second close.",
        "A bare `out <- v` with no `select`, which is exactly the leak.",
    ],
)
