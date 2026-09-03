"""10-advanced-topics / 04-unsafe-package — rotation 1: 5 puzzles per level."""

SUB = "04-unsafe-package"

SPECS = []


def P(level, **kw):
    kw.setdefault("sub", SUB)
    kw["level"] = level
    kw.setdefault("mode", "stub")
    kw.setdefault("stub", 'panic("not implemented")')
    kw.setdefault("imports", [])
    kw.setdefault("extra", "")
    SPECS.append(kw)


# ---------------------------------------------------------------- junior -----

P(
    "junior",
    name="sizeof",
    title="How Big Is This, Really",
    sig="func Sizes() (header, id, name uintptr)",
    doc="""Sizes returns the size in bytes of the Header type and of its Id and
Name fields.

unsafe.Sizeof is a compile-time constant: it measures the type, not the
data a pointer or slice header refers to.

Examples:

	Sizes() => 40, 8, 16 on a 64-bit build""",
    imports=['"unsafe"'],
    extra="""// Header is the fixed part of a record.
type Header struct {
	Id   int64
	Name string
	Tags []string
}""",
    solution="""var h Header
return unsafe.Sizeof(h), unsafe.Sizeof(h.Id), unsafe.Sizeof(h.Name)""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestSizes(t *testing.T) {
	h, id, name := Sizes()
	var want Header
	if h != unsafe.Sizeof(want) {
		t.Errorf("header = %d, want %d", h, unsafe.Sizeof(want))
	}
	if id != unsafe.Sizeof(want.Id) {
		t.Errorf("id = %d, want %d", id, unsafe.Sizeof(want.Id))
	}
	if name != unsafe.Sizeof(want.Name) {
		t.Errorf("name = %d, want %d", name, unsafe.Sizeof(want.Name))
	}
}

func TestSizesAreTypeSizes(t *testing.T) {
	_, _, name := Sizes()
	long := Header{Name: "a very long name that changes nothing about the header"}
	if got := unsafe.Sizeof(long.Name); got != name {
		t.Errorf("string header = %d, want %d: Sizeof measures the type, not the bytes", got, name)
	}
}

func TestHeaderIsBiggerThanItsScalarField(t *testing.T) {
	h, id, _ := Sizes()
	if h <= id {
		t.Errorf("header = %d, id = %d: the struct must be larger than one field", h, id)
	}
}
""",
    context="A capacity estimate multiplies the record count by \"about the size of the struct\". The estimate is off by a factor of three because nobody measured the string and slice headers.",
    task=[
        "Return `unsafe.Sizeof` for the `Header` type, its `Id` field and its `Name` field.",
        "Do not hard-code numbers — the answers must follow the type.",
    ],
    examples=[
        ("Sizes()", "40, 8, 16", "On a 64-bit build; a string header is two words."),
        ("Sizeof of a long Name", "still 16", "The header size does not depend on the text."),
        ("header vs id", "header is larger", None),
    ],
    topics=[
        ("unsafe.Sizeof", "A compile-time constant giving the type's size in bytes."),
        ("Headers vs payload", "A string is a pointer plus a length; a slice adds a capacity."),
        ("Platform dependence", "Word size makes these numbers architecture-specific."),
    ],
    hint="You need a value to take `Sizeof` of. A zero `Header` will do.",
    intuition="`Sizeof` answers \"how many bytes does a variable of this type occupy\", which is not the same as \"how much memory does this value use\". A string field is always two words, whatever text it points at.",
    approach=[
        "Declare a zero `Header`.",
        "Return `unsafe.Sizeof` of the struct and of the two fields.",
    ],
    walkthrough="On a 64-bit build `Id` is 8 bytes, `Name` is 16 (pointer + length), `Tags` is 24 (pointer + length + capacity), so the struct is 48.",
    pitfalls=[
        "Writing the numbers as literals — they change with the architecture.",
        "Expecting `Sizeof` on a slice to include the elements; it does not.",
    ],
)

P(
    "junior",
    name="offsetof",
    title="Where Does This Field Start",
    sig="func Offsets() (a, b, c uintptr)",
    doc="""Offsets returns the byte offset of each field of Rec from the start of
the struct.

Offsets are decided by the compiler from the field order and the
alignment rules, not by the field sizes alone.

Examples:

	Offsets() => 0, 8, 16 for the declared layout""",
    imports=['"unsafe"'],
    extra="""// Rec is a record with mixed field widths.
type Rec struct {
	A byte
	B int64
	C byte
}""",
    solution="""var r Rec
return unsafe.Offsetof(r.A), unsafe.Offsetof(r.B), unsafe.Offsetof(r.C)""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestOffsets(t *testing.T) {
	var r Rec
	a, b, c := Offsets()
	if a != unsafe.Offsetof(r.A) || b != unsafe.Offsetof(r.B) || c != unsafe.Offsetof(r.C) {
		t.Errorf("Offsets = %d, %d, %d, want %d, %d, %d",
			a, b, c, unsafe.Offsetof(r.A), unsafe.Offsetof(r.B), unsafe.Offsetof(r.C))
	}
}

func TestFirstFieldStartsAtZero(t *testing.T) {
	if a, _, _ := Offsets(); a != 0 {
		t.Errorf("first offset = %d, want 0", a)
	}
}

func TestAlignmentCreatesAGap(t *testing.T) {
	var r Rec
	a, b, _ := Offsets()
	if b-a <= unsafe.Sizeof(r.A) {
		t.Errorf("B starts at %d, right after a %d-byte field: alignment should have pushed it further",
			b, unsafe.Sizeof(r.A))
	}
}

func TestOffsetsAreAscending(t *testing.T) {
	a, b, c := Offsets()
	if !(a < b && b < c) {
		t.Errorf("offsets %d, %d, %d are not ascending", a, b, c)
	}
}
""",
    context="A binary protocol is written by copying a struct's bytes. It works on one machine and produces garbage on another, because nobody checked where the fields actually sit.",
    task=[
        "Return the byte offset of each field of `Rec`.",
        "Derive them with `unsafe.Offsetof`, not by adding up sizes.",
    ],
    examples=[
        ("Offsets()", "0, 8, 16", "The one-byte `A` is followed by seven bytes of padding."),
        ("first offset", "0", "The first field always starts at the struct's address."),
        ("B - A", "8, not 1", "`int64` must be 8-byte aligned."),
    ],
    topics=[
        ("unsafe.Offsetof", "The field's byte offset within its struct, as a compile-time constant."),
        ("Alignment", "A field's offset is rounded up to a multiple of its alignment."),
        ("Padding", "The gap alignment leaves behind is real memory."),
    ],
    hint="`Offsetof` takes a field selector, so you need a variable to select from.",
    intuition="Fields are laid out in declaration order, but each one starts at an offset its type can be aligned to. That is why a byte followed by an int64 leaves seven bytes unused.",
    approach=[
        "Declare a zero `Rec`.",
        "Return `unsafe.Offsetof` for each of its three fields.",
    ],
    walkthrough="`A` is at 0. `B` is an int64 and needs 8-byte alignment, so it starts at 8, not 1. `C` follows at 16, and the struct is padded to 24.",
    pitfalls=[
        "Computing offsets as running sums of `Sizeof` — that ignores alignment.",
        "`Offsetof` needs the field selector `r.B`, not the type."
    ],
)

P(
    "junior",
    name="bytestostring",
    title="A String View Of Bytes You Own",
    sig="func Str(b []byte) string",
    doc="""Str returns a string that shares b's bytes instead of copying them.

The result is only valid while b is not written to again: a string is
supposed to be immutable, and this one is not.

Examples:

	Str([]byte("hi")) => "hi" """,
    imports=['"unsafe"'],
    solution="""if len(b) == 0 {
	return ""
}
return unsafe.String(unsafe.SliceData(b), len(b))""",
    tests="""
import (
	"testing"
	"unsafe"
)

var sink string

func TestStr(t *testing.T) {
	if got := Str([]byte("hello")); got != "hello" {
		t.Errorf("Str = %q, want \\"hello\\"", got)
	}
	if got := Str(nil); got != "" {
		t.Errorf("Str(nil) = %q, want empty", got)
	}
	if got := Str([]byte{}); got != "" {
		t.Errorf("Str([]) = %q, want empty", got)
	}
}

func TestStrSharesTheBytes(t *testing.T) {
	b := []byte("abc")
	s := Str(b)
	if unsafe.StringData(s) != unsafe.SliceData(b) {
		t.Error("Str copied the bytes; it must share them")
	}
}

func TestStrDoesNotAllocate(t *testing.T) {
	b := make([]byte, 4096)
	for i := range b {
		b[i] = 'x'
	}
	if n := testing.AllocsPerRun(200, func() { sink = Str(b) }); n != 0 {
		t.Errorf("Str made %v allocations, want 0", n)
	}
}
""",
    context="A parser converts every field to a string just to compare it. Each conversion copies bytes the parser already has and throws the copy away a line later.",
    task=[
        "Return a string sharing `b`'s bytes, with no copy.",
        "An empty or nil input returns the empty string.",
        "Zero allocations, whatever the length.",
    ],
    examples=[
        ('Str([]byte("hello"))', '"hello"', None),
        ("Str(nil)", '""', None),
        ("StringData(result) vs SliceData(input)", "the same pointer", "No copy happened."),
    ],
    topics=[
        ("unsafe.String", "Builds a string header over a pointer and a length."),
        ("unsafe.SliceData", "The address of a slice's first element, defined even for empty slices."),
        ("The obligation you take on", "The bytes must not change while the string is alive."),
    ],
    hint="`unsafe.String` needs a `*byte` and a length. `unsafe.SliceData` provides the first.",
    intuition="A string and a byte slice have the same bytes underneath; only the promise differs. `unsafe.String` makes the string header point at bytes you already have — and hands you the job of keeping that promise.",
    approach=[
        "Return `\"\"` for an empty input.",
        "`unsafe.String(unsafe.SliceData(b), len(b))`.",
    ],
    walkthrough="For a 4096-byte buffer, `string(b)` allocates and copies 4096 bytes. `unsafe.String` writes a two-word header and copies nothing.",
    pitfalls=[
        "Handing the result to a caller who outlives the buffer — that is a use-after-write, not a use-after-free, and it is silent.",
        "Skipping the empty-input guard; a nil slice's data pointer is nil, and a nil pointer with a non-zero length is invalid.",
    ],
)

P(
    "junior",
    name="stringtobytes",
    title="Read A String's Bytes Without Copying",
    sig="func Bytes(s string) []byte",
    doc="""Bytes returns a read-only byte view of s.

The bytes belong to the string and may live in read-only memory, so the
result must never be written to.

Examples:

	Bytes("hi") => []byte("hi"), sharing the string's bytes""",
    imports=['"unsafe"'],
    solution="""if len(s) == 0 {
	return nil
}
return unsafe.Slice(unsafe.StringData(s), len(s))""",
    tests="""
import (
	"bytes"
	"testing"
	"unsafe"
)

var sink []byte

func TestBytes(t *testing.T) {
	if got := Bytes("hello"); !bytes.Equal(got, []byte("hello")) {
		t.Errorf("Bytes = %q, want \\"hello\\"", got)
	}
	if got := Bytes(""); len(got) != 0 {
		t.Errorf("Bytes(\\"\\") = %q, want empty", got)
	}
}

func TestBytesSharesTheString(t *testing.T) {
	s := "shared"
	b := Bytes(s)
	if unsafe.SliceData(b) != unsafe.StringData(s) {
		t.Error("Bytes copied the string; it must share it")
	}
}

func TestBytesLengthAndCapacity(t *testing.T) {
	b := Bytes("abcd")
	if len(b) != 4 {
		t.Errorf("len = %d, want 4", len(b))
	}
	if cap(b) != 4 {
		t.Errorf("cap = %d, want 4: an append must not write past the string", cap(b))
	}
}

func TestBytesDoesNotAllocate(t *testing.T) {
	s := string(make([]byte, 4096))
	if n := testing.AllocsPerRun(200, func() { sink = Bytes(s) }); n != 0 {
		t.Errorf("Bytes made %v allocations, want 0", n)
	}
}
""",
    context="A checksum routine takes `[]byte` and every caller has a string. `[]byte(s)` copies the whole payload just so it can be read once.",
    task=[
        "Return a byte view of `s` without copying.",
        "The empty string yields an empty result.",
        "Zero allocations; the result's capacity must equal its length.",
    ],
    examples=[
        ('Bytes("hello")', '[]byte("hello")', "Sharing the string's bytes."),
        ('Bytes("")', "[]", None),
        ('cap(Bytes("abcd"))', "4", "So an append cannot write past the string."),
    ],
    topics=[
        ("unsafe.Slice", "Builds a slice header over a pointer and a length."),
        ("unsafe.StringData", "The address of a string's first byte."),
        ("Read-only in practice", "String literals may live in a read-only mapping; writing there faults."),
    ],
    hint="The mirror image of the byte-to-string direction.",
    intuition="Strings are immutable by contract, not by the type system. `unsafe.Slice` gives you a mutable-looking view of immutable bytes, which is exactly why the result must stay read-only.",
    approach=[
        "Return nil for the empty string.",
        "`unsafe.Slice(unsafe.StringData(s), len(s))`.",
    ],
    walkthrough="`unsafe.Slice(p, 4)` yields a slice of length and capacity 4 over the string's bytes — no allocation, and no room for `append` to overwrite anything after it.",
    pitfalls=[
        "Writing through the result — for a string literal that is a segmentation fault, and for any other string it corrupts a value the whole program believes is immutable.",
        "Passing the result to a function that appends to it.",
    ],
)

P(
    "junior",
    name="samearray",
    title="Do These Two Slices Share Storage",
    sig="func SameArray(a, b []int) bool",
    doc="""SameArray reports whether a and b start at the same element of the
same backing array.

Comparing slices with == is not allowed; comparing their data pointers
is.

Examples:

	s := []int{1, 2}; SameArray(s, s[:1]) => true""",
    imports=['"unsafe"'],
    solution="""if len(a) == 0 || len(b) == 0 {
	return false
}
return unsafe.SliceData(a) == unsafe.SliceData(b)""",
    tests="""
import "testing"

func TestSameArray(t *testing.T) {
	s := []int{1, 2, 3}
	if !SameArray(s, s) {
		t.Error("SameArray(s, s) = false, want true")
	}
	if !SameArray(s, s[:2]) {
		t.Error("SameArray(s, s[:2]) = false, want true")
	}
	if SameArray(s, s[1:]) {
		t.Error("SameArray(s, s[1:]) = true, want false: the start differs")
	}
}

func TestSameArrayDistinctSlices(t *testing.T) {
	a := []int{1, 2}
	b := []int{1, 2}
	if SameArray(a, b) {
		t.Error("SameArray = true for two separate arrays, want false")
	}
}

func TestSameArrayEmpty(t *testing.T) {
	s := []int{1}
	for _, c := range [][2][]int{
		{nil, nil}, {s, nil}, {nil, s}, {s[:0], s},
	} {
		if SameArray(c[0], c[1]) {
			t.Errorf("SameArray(%v, %v) = true, want false for an empty operand", c[0], c[1])
		}
	}
}

func TestSameArrayAfterCopy(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, 3)
	copy(b, a)
	if SameArray(a, b) {
		t.Error("a copy must not share storage")
	}
}
""",
    context="A helper documents that it may return its input or a fresh slice. A caller needs to know which happened before it mutates the result.",
    task=[
        "Report whether `a` and `b` begin at the same element of the same array.",
        "Any empty operand reports false.",
    ],
    examples=[
        ("s := []int{1,2,3}; SameArray(s, s[:2])", "true", "Same start, different length."),
        ("SameArray(s, s[1:])", "false", "The start differs."),
        ("SameArray(a, b) after copy", "false", None),
    ],
    topics=[
        ("unsafe.SliceData", "The slice's data pointer, comparable with =="),
        ("Slices are not comparable", "`a == b` does not compile for slices; only `== nil` is allowed."),
        ("Aliasing detection", "Shared storage is what makes one write visible through another slice."),
    ],
    hint="Two pointers, one comparison — plus a guard for empties.",
    intuition="A slice header's first word is the address of its first element. Two slices alias when that address is the same, which is a pointer comparison the language will not let you write without `unsafe`.",
    approach=[
        "Return false when either slice is empty.",
        "Compare `unsafe.SliceData(a)` with `unsafe.SliceData(b)`.",
    ],
    walkthrough="`s` and `s[:2]` both point at `&s[0]`, so the comparison is true. `s[1:]` points at `&s[1]`, so it is false — this is a same-start test, not an overlap test.",
    pitfalls=[
        "Treating this as an overlap check; overlapping slices with different starts report false.",
        "Skipping the empty guard, where the data pointer of a nil slice is nil and two nils would compare equal.",
    ],
)

# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="reorder",
    title="The Same Fields, Half The Bytes",
    mode="bug",
    sig="func Size() uintptr",
    doc="""Size returns the size of the Entry type.

Entry's fields are declared in an order that forces the compiler to
insert padding between them. Reordering them from widest to narrowest
removes it without changing what the struct holds.

Examples:

	Size() => 16 once the fields are ordered well""",
    imports=['"unsafe"'],
    extra="""// Entry is one cache record. Reorder its fields to remove the padding.
//
// CHANGE STRUCT BELOW THIS LINE
type Entry struct {
	Flag byte
	Ref  int64
	Kind byte
	Seq  int32
}

// CHANGE STRUCT ABOVE THIS LINE""",
    sol_extra="""// Entry is one cache record, ordered widest field first.
type Entry struct {
	Ref  int64
	Seq  int32
	Flag byte
	Kind byte
}""",
    buggy="""return unsafe.Sizeof(Entry{})""",
    solution="""return unsafe.Sizeof(Entry{})""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestSizeMatchesTheType(t *testing.T) {
	if got := Size(); got != unsafe.Sizeof(Entry{}) {
		t.Errorf("Size = %d, want %d", got, unsafe.Sizeof(Entry{}))
	}
}

func TestEntryIsPacked(t *testing.T) {
	if got := unsafe.Sizeof(Entry{}); got > 16 {
		t.Errorf("sizeof(Entry) = %d, want at most 16: reorder the fields widest first", got)
	}
}

func TestEntryStillHasEveryField(t *testing.T) {
	e := Entry{Flag: 1, Ref: 2, Kind: 3, Seq: 4}
	if e.Flag != 1 || e.Ref != 2 || e.Kind != 3 || e.Seq != 4 {
		t.Errorf("e = %+v: the fields and their types must not change", e)
	}
}

func TestFieldTypesAreUnchanged(t *testing.T) {
	e := Entry{}
	if unsafe.Sizeof(e.Ref) != 8 || unsafe.Sizeof(e.Seq) != 4 ||
		unsafe.Sizeof(e.Flag) != 1 || unsafe.Sizeof(e.Kind) != 1 {
		t.Error("a field's type changed; only the declaration order may move")
	}
}
""",
    context="A cache holds fifty million of one small struct. Its four fields need fourteen bytes and the struct occupies twenty-four, because they were declared in the order somebody thought of them.",
    task=[
        "Reorder `Entry`'s fields so the struct occupies at most 16 bytes.",
        "Keep every field, with its name and its type.",
        "`Size` must keep reporting the type's real size.",
    ],
    examples=[
        ("sizeof(Entry) as declared", "24", "Padding after each narrow field."),
        ("sizeof(Entry) reordered", "16", None),
        ("the field set", "unchanged", "Only the order moves."),
    ],
    topics=[
        ("Alignment drives padding", "Each field starts at a multiple of its alignment; the gaps are wasted."),
        ("Widest first", "Descending field width usually eliminates internal padding."),
        ("Tail padding", "The struct's size is rounded up to its own alignment."),
    ],
    hint="`int64` first. Then `int32`. Then the bytes.",
    intuition="The compiler will not reorder your fields, so the layout is your responsibility. Sorting from widest to narrowest lets each field land on its natural boundary with nothing skipped.",
    approach=[
        "Declare `Ref int64` first.",
        "Then `Seq int32`.",
        "Then the two `byte` fields.",
    ],
    walkthrough="As declared: Flag at 0, seven bytes of padding, Ref at 8, Kind at 16, three bytes of padding, Seq at 20 — 24 bytes. Reordered: Ref at 0, Seq at 8, Flag at 12, Kind at 13, then two bytes of tail padding — 16.",
    pitfalls=[
        "Changing `int64` to `int32` to save space — that changes the type, and the test checks it.",
        "Assuming the saving is always this large; a struct of same-width fields has no padding to remove.",
    ],
)

P(
    "middle",
    name="alignment",
    title="Is This Address Aligned",
    sig="func Aligned(b []byte, n uintptr) bool",
    doc="""Aligned reports whether b's first byte sits at an address that is a
multiple of n.

n must be a power of two; anything else, or an empty slice, reports
false.

Examples:

	Aligned(buf, 8) => true when buf starts on an 8-byte boundary""",
    imports=['"unsafe"'],
    solution="""if len(b) == 0 || n == 0 || n&(n-1) != 0 {
	return false
}
return uintptr(unsafe.Pointer(unsafe.SliceData(b)))&(n-1) == 0""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestAlignedOnAWideSlice(t *testing.T) {
	u := make([]uint64, 4)
	b := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), 32)
	if !Aligned(b, 8) {
		t.Error("a []uint64's storage must be 8-byte aligned")
	}
	if !Aligned(b, 1) {
		t.Error("everything is 1-byte aligned")
	}
}

func TestAlignedDetectsAnOffset(t *testing.T) {
	u := make([]uint64, 4)
	b := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), 32)
	if Aligned(b[1:], 8) {
		t.Error("a slice starting one byte in cannot be 8-byte aligned")
	}
	if !Aligned(b[8:], 8) {
		t.Error("eight bytes further along is aligned again")
	}
}

func TestAlignedRejectsBadInput(t *testing.T) {
	b := make([]byte, 8)
	if Aligned(nil, 8) {
		t.Error("Aligned(nil) = true, want false")
	}
	if Aligned(b, 0) {
		t.Error("Aligned(b, 0) = true, want false")
	}
	if Aligned(b, 3) {
		t.Error("Aligned(b, 3) = true, want false: n must be a power of two")
	}
	if Aligned(b, 6) {
		t.Error("Aligned(b, 6) = true, want false")
	}
}
""",
    context="A decoder reinterprets an incoming buffer as a slice of `uint64`. On x86 it is merely slow when the buffer is misaligned; on other architectures it faults.",
    task=[
        "Report whether `b`'s first byte is at an address that is a multiple of `n`.",
        "Return false for an empty slice, for `n == 0`, and for any `n` that is not a power of two.",
    ],
    examples=[
        ("Aligned(bufFromUint64s, 8)", "true", None),
        ("Aligned(buf[1:], 8)", "false", "One byte in is no longer aligned."),
        ("Aligned(buf, 3)", "false", "3 is not a power of two."),
    ],
    topics=[
        ("Pointer to uintptr", "The numeric address is only meaningful for arithmetic like this."),
        ("Power-of-two masking", "`x & (n-1) == 0` is the alignment test when n is a power of two."),
        ("Why alignment matters", "Wide loads on a misaligned address are slow or illegal depending on the machine."),
    ],
    hint="A power of two has exactly one bit set: `n & (n-1)` is zero.",
    intuition="Alignment is a property of the address, so you have to look at the address as a number. The mask trick works because a power-of-two boundary means the low bits of the address are all zero.",
    approach=[
        "Reject an empty slice, `n == 0`, and non-powers of two.",
        "Convert the data pointer to `uintptr` and test `addr & (n-1) == 0`.",
    ],
    walkthrough="A `[]uint64`'s storage is 8-byte aligned, so its address ends in three zero bits and `addr & 7` is 0. Taking `b[1:]` adds 1, so the test fails; `b[8:]` adds 8 and it passes again.",
    pitfalls=[
        "Storing the `uintptr` in a variable and using it later — the collector may move the object, and the number goes stale.",
        "Using `%` with a non-power-of-two `n`, which the guard exists to reject.",
    ],
)

P(
    "middle",
    name="fieldptr",
    title="Reach A Field Through Its Offset",
    sig="func BumpSeq(p *Rec) int64",
    doc="""BumpSeq increments the Seq field of the record p points at, using the
field's offset rather than the field selector, and returns the new value.

This is what a generic marshaller does when it only knows the offset.

Examples:

	r := &Rec{Seq: 1}; BumpSeq(r) => 2, r.Seq is 2""",
    imports=['"unsafe"'],
    extra="""// Rec is a record whose Seq field is reached by offset.
type Rec struct {
	Tag  byte
	Seq  int64
	Name string
}""",
    solution="""q := (*int64)(unsafe.Add(unsafe.Pointer(p), unsafe.Offsetof(p.Seq)))
*q++
return *q""",
    tests="""
import "testing"

func TestBumpSeq(t *testing.T) {
	r := &Rec{Tag: 7, Seq: 1, Name: "n"}
	if got := BumpSeq(r); got != 2 {
		t.Errorf("BumpSeq = %d, want 2", got)
	}
	if r.Seq != 2 {
		t.Errorf("r.Seq = %d, want 2: the write must reach the caller's record", r.Seq)
	}
}

func TestBumpSeqLeavesTheOtherFields(t *testing.T) {
	r := &Rec{Tag: 7, Seq: 0, Name: "name"}
	BumpSeq(r)
	if r.Tag != 7 || r.Name != "name" {
		t.Errorf("r = %+v: only Seq may change", *r)
	}
}

func TestBumpSeqRepeated(t *testing.T) {
	r := &Rec{}
	for i := 1; i <= 100; i++ {
		if got := BumpSeq(r); got != int64(i) {
			t.Fatalf("call %d returned %d", i, got)
		}
	}
}

func TestBumpSeqNegative(t *testing.T) {
	r := &Rec{Seq: -1}
	if got := BumpSeq(r); got != 0 {
		t.Errorf("BumpSeq = %d, want 0", got)
	}
}
""",
    context="A marshaller stores each field's offset once and then writes through it for every record. Getting from a struct pointer plus an offset to a typed pointer is the step that has to be right.",
    task=[
        "Increment `p.Seq` by writing through a pointer built from the struct pointer and the field's offset.",
        "Return the new value.",
        "No other field may change.",
    ],
    examples=[
        ("r := &Rec{Seq: 1}; BumpSeq(r)", "2, r.Seq is 2", None),
        ("other fields after the call", "unchanged", None),
        ("100 calls from zero", "1..100", None),
    ],
    topics=[
        ("unsafe.Add", "Advances an `unsafe.Pointer` by a byte count, keeping it a pointer the collector understands."),
        ("unsafe.Offsetof", "The compile-time offset of the field within its struct."),
        ("Pointer conversion", "Casting `unsafe.Pointer` to `*int64` is what makes the memory typed again."),
    ],
    hint="Struct pointer to `unsafe.Pointer`, add the offset, convert to `*int64`.",
    intuition="`unsafe.Add` is pointer arithmetic that stays legal: the result is still a pointer into the same object, so the garbage collector keeps tracking it. Doing the same arithmetic on a `uintptr` would not be safe.",
    approach=[
        "Convert `p` to `unsafe.Pointer`.",
        "`unsafe.Add` it by `unsafe.Offsetof(p.Seq)`.",
        "Convert to `*int64`, increment through it, and return the value.",
    ],
    walkthrough="`Tag` is at 0, so `Seq` is at offset 8 after padding. Adding 8 to the struct's address and reading it as `*int64` is exactly `&p.Seq`.",
    pitfalls=[
        "Doing the arithmetic in `uintptr` across separate statements — a moving collector would invalidate the number.",
        "Hard-coding the offset instead of asking `Offsetof`, which breaks the moment the struct changes.",
    ],
)

P(
    "middle",
    name="uint32at",
    title="Read A Wide Value Out Of A Byte Buffer",
    sig="func Uint32At(b []byte, off int) (uint32, bool)",
    doc="""Uint32At reads the native-endian uint32 at byte offset off in b.

The read is bounds-checked and alignment-checked; anything out of range
or misaligned reports false rather than faulting.

Examples:

	Uint32At(buf, 0) => the first four bytes as a uint32""",
    imports=['"unsafe"'],
    solution="""if off < 0 || off+4 > len(b) {
	return 0, false
}
p := unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), off)
if uintptr(p)&3 != 0 {
	return 0, false
}
return *(*uint32)(p), true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func aligned32(n int) []byte {
	u := make([]uint32, (n+3)/4)
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), n)
}

func TestUint32At(t *testing.T) {
	b := aligned32(8)
	p := (*uint32)(unsafe.Pointer(unsafe.SliceData(b)))
	*p = 0x01020304
	got, ok := Uint32At(b, 0)
	if !ok || got != 0x01020304 {
		t.Errorf("Uint32At = %#x, %v, want 0x01020304, true", got, ok)
	}
}

func TestUint32AtSecondWord(t *testing.T) {
	b := aligned32(8)
	q := (*uint32)(unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), 4))
	*q = 0xdeadbeef
	got, ok := Uint32At(b, 4)
	if !ok || got != 0xdeadbeef {
		t.Errorf("Uint32At = %#x, %v, want 0xdeadbeef, true", got, ok)
	}
}

func TestUint32AtOutOfRange(t *testing.T) {
	b := aligned32(8)
	for _, off := range []int{-1, 5, 8, 9, 100} {
		if _, ok := Uint32At(b, off); ok {
			t.Errorf("Uint32At(off=%d) reported ok, want false", off)
		}
	}
	if _, ok := Uint32At(nil, 0); ok {
		t.Error("Uint32At(nil) reported ok, want false")
	}
}

func TestUint32AtMisaligned(t *testing.T) {
	b := aligned32(16)
	if _, ok := Uint32At(b[1:], 0); ok {
		t.Error("a misaligned read reported ok, want false")
	}
	if _, ok := Uint32At(b, 2); ok {
		t.Error("an odd offset reported ok, want false")
	}
}
""",
    context="A zero-copy frame reader casts into the middle of a received buffer. On the first misaligned frame it reads garbage on one architecture and crashes on another.",
    task=[
        "Read the native-endian uint32 at byte offset `off`.",
        "Report false for a negative offset, a read that would run past the end, or a misaligned address.",
    ],
    examples=[
        ("Uint32At(buf, 0)", "the first word, true", None),
        ("Uint32At(buf, 5)", "0, false", "The read would run past the end."),
        ("Uint32At(buf, 2)", "0, false", "Not 4-byte aligned."),
    ],
    topics=[
        ("Bounds checks are yours now", "`unsafe` removes the runtime's check, so the code must do it."),
        ("Alignment for wide loads", "A uint32 read wants a 4-byte-aligned address."),
        ("unsafe.Add plus conversion", "Offset first, then reinterpret the pointer."),
    ],
    hint="Three guards before the read: negative offset, past the end, and the low two bits of the address.",
    intuition="Reinterpreting memory is only as safe as the checks you write around it. The compiler stops checking the moment you take an `unsafe.Pointer`, so both the range and the alignment become your responsibility.",
    approach=[
        "Reject `off < 0` and `off+4 > len(b)`.",
        "`unsafe.Add` the data pointer by `off`.",
        "Reject when the address's low two bits are set.",
        "Read through `(*uint32)(p)`.",
    ],
    walkthrough="For an 8-byte buffer and off = 4: the range check passes, the address is 4 bytes past an aligned base so the low bits are clear, and the read returns the second word.",
    pitfalls=[
        "Checking `off < len(b)` instead of `off+4 <= len(b)` — that lets a read run three bytes past the end.",
        "Checking the alignment of `off` alone; the buffer's own start may be misaligned.",
    ],
)

P(
    "middle",
    name="int32view",
    title="Reinterpret Bytes As Wider Values",
    sig="func Int32s(b []byte) ([]int32, bool)",
    doc="""Int32s returns a []int32 view over b's bytes, sharing the storage.

The view is only valid when b's length is a multiple of four and its
first byte is 4-byte aligned; otherwise the second result is false.

Examples:

	Int32s(eightBytes) => a 2-element view, true""",
    imports=['"unsafe"'],
    solution="""if len(b) == 0 || len(b)%4 != 0 {
	return nil, false
}
p := unsafe.Pointer(unsafe.SliceData(b))
if uintptr(p)&3 != 0 {
	return nil, false
}
return unsafe.Slice((*int32)(p), len(b)/4), true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func alignedBytes(n int) []byte {
	u := make([]int32, (n+3)/4)
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), n)
}

func TestInt32sShape(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Int32s(b)
	if !ok {
		t.Fatal("Int32s reported false for an aligned 8-byte buffer")
	}
	if len(v) != 2 || cap(v) != 2 {
		t.Errorf("len, cap = %d, %d, want 2, 2", len(v), cap(v))
	}
}

func TestInt32sSharesStorage(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Int32s(b)
	if !ok {
		t.Fatal("Int32s reported false")
	}
	v[0] = 0x01020304
	if b[0] == 0 && b[1] == 0 && b[2] == 0 && b[3] == 0 {
		t.Error("the view does not share the bytes")
	}
	v[1] = -1
	for _, x := range b[4:8] {
		if x != 0xff {
			t.Errorf("b[4:8] = %v, want all 0xff", b[4:8])
			break
		}
	}
}

func TestInt32sRejectsBadShapes(t *testing.T) {
	b := alignedBytes(16)
	for _, c := range []struct {
		name string
		in   []byte
	}{
		{"nil", nil},
		{"empty", b[:0]},
		{"length not a multiple of 4", b[:6]},
		{"misaligned", b[1:13]},
	} {
		if _, ok := Int32s(c.in); ok {
			t.Errorf("%s: reported ok, want false", c.name)
		}
	}
}

func TestInt32sDoesNotAllocate(t *testing.T) {
	b := alignedBytes(4096)
	var sink []int32
	if n := testing.AllocsPerRun(100, func() { sink, _ = Int32s(b) }); n != 0 {
		t.Errorf("Int32s made %v allocations, want 0", n)
	}
	_ = sink
}
""",
    context="A numeric kernel receives its input as bytes off the wire and wants to run over it as `int32` without copying eight megabytes first.",
    task=[
        "Return an `[]int32` view sharing `b`'s storage.",
        "Report false for an empty slice, a length that is not a multiple of four, or a misaligned start.",
        "Zero allocations.",
    ],
    examples=[
        ("Int32s(alignedBytes(8))", "a 2-element view, true", None),
        ("Int32s(b[:6])", "nil, false", "6 is not a multiple of 4."),
        ("Int32s(b[1:13])", "nil, false", "Misaligned start."),
    ],
    topics=[
        ("unsafe.Slice", "Reinterprets a typed pointer and a length as a slice."),
        ("Element count, not byte count", "The length argument is in elements — `len(b)/4`."),
        ("Two preconditions", "Length divisibility and address alignment are both required."),
    ],
    hint="`unsafe.Slice((*int32)(p), n)` — think carefully about what `n` is.",
    intuition="A slice header is a pointer, a length and a capacity. Reinterpreting means keeping the pointer and changing the element type — so the length has to be recomputed in the new element's units.",
    approach=[
        "Reject an empty slice and a length that is not a multiple of four.",
        "Reject a data pointer whose low two bits are set.",
        "Return `unsafe.Slice((*int32)(p), len(b)/4)`.",
    ],
    walkthrough="Eight aligned bytes become two int32s sharing the same memory: writing `v[1] = -1` sets bytes 4 through 7 to 0xff.",
    pitfalls=[
        "Passing `len(b)` as the element count, which produces a view four times too long — and reads past the buffer.",
        "Forgetting that the result's endianness is the machine's, so it is not a wire format.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="takestr",
    title="The String That Changed Underneath Its Owner",
    mode="bug",
    sig="func Take(buf []byte, n int) string",
    doc="""Take returns the first n bytes of buf as a string the caller keeps.

buf is a scratch buffer the caller reuses, so the result must not be a
view of it — a string that changes is a contradiction the rest of the
program is not prepared for.

Examples:

	Take([]byte("hello"), 2) => "he", independent of buf""",
    imports=['"unsafe"'],
    buggy="""if n <= 0 || len(buf) == 0 {
	return ""
}
if n > len(buf) {
	n = len(buf)
}
return unsafe.String(unsafe.SliceData(buf), n)""",
    solution="""if n <= 0 || len(buf) == 0 {
	return ""
}
if n > len(buf) {
	n = len(buf)
}
out := make([]byte, n)
copy(out, buf)
return unsafe.String(unsafe.SliceData(out), n)""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestTake(t *testing.T) {
	if got := Take([]byte("hello"), 2); got != "he" {
		t.Errorf("Take = %q, want \\"he\\"", got)
	}
	if got := Take([]byte("hi"), 9); got != "hi" {
		t.Errorf("Take = %q, want \\"hi\\"", got)
	}
	if got := Take(nil, 3); got != "" {
		t.Errorf("Take = %q, want empty", got)
	}
	if got := Take([]byte("hi"), 0); got != "" {
		t.Errorf("Take = %q, want empty", got)
	}
}

func TestTakeSurvivesBufferReuse(t *testing.T) {
	buf := make([]byte, 8)
	copy(buf, "first")
	s := Take(buf, 5)
	copy(buf, "SECOND")
	if s != "first" {
		t.Errorf("s = %q, want \\"first\\": the string is a view of the reused buffer", s)
	}
}

func TestTakeResultsAreIndependent(t *testing.T) {
	buf := make([]byte, 4)
	got := make([]string, 0, 26)
	for i := 0; i < 26; i++ {
		for j := range buf {
			buf[j] = byte('a' + i)
		}
		got = append(got, Take(buf, 4))
	}
	for i, s := range got {
		want := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if s != want {
			t.Fatalf("result %d = %q, want %q", i, s, want)
		}
	}
}

func TestTakeStillAvoidsTheDoubleCopy(t *testing.T) {
	buf := make([]byte, 256)
	var sink string
	n := testing.AllocsPerRun(100, func() { sink = Take(buf, 256) })
	_ = sink
	if n > 1 {
		t.Errorf("Take made %v allocations, want 1: one copy, then a header over it", n)
	}
}

func TestTakeDoesNotAliasTheInput(t *testing.T) {
	buf := []byte("abcd")
	s := Take(buf, 4)
	if unsafe.StringData(s) == unsafe.SliceData(buf) {
		t.Error("the result shares the caller's buffer")
	}
}
""",
    context="A reader hands out strings built over its own reusable read buffer. Log lines come out interleaved, map keys stop matching themselves, and the strings in a cache change hours after they were stored.",
    task=[
        "Return the first `n` bytes of `buf` as a string, clamping `n` into `[0, len(buf)]`.",
        "The result must survive later writes to `buf`.",
        "Fix the single bug, and keep the cost at one allocation per call.",
    ],
    examples=[
        ('Take([]byte("hello"), 2)', '"he"', None),
        ('s := Take(buf, 5); copy(buf, "SECOND")', "s is unchanged", "The string owns its bytes."),
        ("Take(nil, 3)", '""', None),
    ],
    topics=[
        ("Strings are assumed immutable", "Maps cache their hashes, and the compiler folds comparisons — a mutating string breaks both."),
        ("unsafe.String does not copy", "It builds a header; ownership of the bytes is the caller's problem."),
        ("Copy once, then wrap", "Allocating the bytes and wrapping them keeps the cost at one allocation."),
    ],
    hint="The conversion is fine. The bytes it points at are the problem.",
    intuition="`unsafe.String` is safe exactly when nothing will ever write to the bytes again. Over a buffer the caller is about to reuse, that condition is false by construction — and nothing will report the violation.",
    approach=[
        "Clamp `n` and handle the empty cases.",
        "Allocate an `n`-byte slice and copy the prefix into it.",
        "Wrap the copy with `unsafe.String`, which adds no second copy.",
    ],
    walkthrough='Before the fix, `Take(buf, 5)` returns a header pointing into `buf`; the next `copy(buf, "SECOND")` rewrites the string in place. After it, the string points at a private array nothing else can reach.',
    pitfalls=[
        "`string(buf[:n])` is also correct here — the point is that the copy is what was missing, not the conversion.",
        "Copying and then converting with `string(out)`, which copies a second time.",
    ],
)

P(
    "senior",
    name="freezebytes",
    title="Hand Out Bytes That Cannot Be Written",
    mode="bug",
    sig="func Snapshot(s string) []byte",
    doc="""Snapshot returns a byte slice holding s's bytes that the caller may
modify freely.

A view over the string's own bytes is not that: strings may live in
read-only memory, and every other holder of s would see the writes.

Examples:

	b := Snapshot("hi"); b[0] = 'H' => s is unaffected""",
    imports=['"unsafe"'],
    sol_imports=[],
    buggy="""if len(s) == 0 {
	return nil
}
return unsafe.Slice(unsafe.StringData(s), len(s))""",
    solution="""if len(s) == 0 {
	return nil
}
out := make([]byte, len(s))
copy(out, s)
return out""",
    tests="""
import (
	"bytes"
	"testing"
	"unsafe"
)

func TestSnapshot(t *testing.T) {
	if got := Snapshot("hello"); !bytes.Equal(got, []byte("hello")) {
		t.Errorf("Snapshot = %q, want \\"hello\\"", got)
	}
	if got := Snapshot(""); len(got) != 0 {
		t.Errorf("Snapshot = %q, want empty", got)
	}
}

func TestSnapshotIsWritable(t *testing.T) {
	s := "abcd"
	b := Snapshot(s)
	if unsafe.SliceData(b) == unsafe.StringData(s) {
		t.Fatal("the result views the string's own bytes: writing to it is undefined")
	}
	b[0] = 'X'
	if s != "abcd" {
		t.Errorf("s = %q, want \\"abcd\\"", s)
	}
}

func TestSnapshotsAreIndependent(t *testing.T) {
	s := "shared"
	a := Snapshot(s)
	b := Snapshot(s)
	if unsafe.SliceData(a) == unsafe.SliceData(b) {
		t.Fatal("two snapshots share storage")
	}
	a[0] = 'X'
	if b[0] != 's' {
		t.Error("two snapshots share storage")
	}
}

func TestSnapshotHasRoomOfItsOwn(t *testing.T) {
	b := Snapshot("abc")
	b = append(b, 'd')
	if string(b) != "abcd" {
		t.Errorf("append gave %q, want \\"abcd\\"", string(b))
	}
}
""",
    context="A helper is described as \"the fast way to get bytes from a string\". A caller sorts the result in place, and a string constant elsewhere in the binary comes back reordered — when the process does not simply fault.",
    task=[
        "Return a byte slice holding `s`'s bytes that the caller may modify.",
        "The empty string yields an empty result.",
        "Fix the single bug: the result must not alias the string.",
    ],
    examples=[
        ('b := Snapshot("abcd"); b[0] = \'X\'', "s is unchanged", None),
        ("two snapshots of one string", "independent", None),
        ('Snapshot("")', "[]", None),
    ],
    topics=[
        ("String memory may be read-only", "Literals are placed in a read-only section; writing there faults."),
        ("Immutability is program-wide", "Every holder of the string would observe a write through an aliased slice."),
        ("The safe direction is asymmetric", "Bytes-to-string can be zero-copy under a no-write promise; string-to-writable-bytes cannot."),
    ],
    hint="The conversion is legal. The promise the function makes is not.",
    intuition="`unsafe.Slice` over `StringData` produces a writable-looking view of memory the whole program treats as constant. There is no way to keep the function's promise without giving the caller bytes of their own.",
    approach=[
        "Return nil for the empty string.",
        "Allocate `len(s)` bytes, copy the string into them, and return the copy.",
    ],
    walkthrough="`Snapshot(\"abcd\")` allocates four bytes and copies. Writing `b[0]` now touches only that array; before the fix it wrote into the string literal's storage.",
    pitfalls=[
        "Assuming a runtime-built string is safe to write — it is still shared with everyone holding that string value.",
        "`[]byte(s)` is the ordinary spelling of the fix; the point is that the copy is required.",
    ],
)

P(
    "senior",
    name="uintptrhold",
    title="An Address Is Not A Reference",
    mode="bug",
    sig="func SecondWord(p *Pair) int64",
    doc="""SecondWord returns the B field of the pair p points at, reached through
the field's offset.

Address arithmetic must stay in unsafe.Pointer. A uintptr is a plain
number: nothing keeps the object alive and nothing updates it.

Examples:

	SecondWord(&Pair{A: 1, B: 2}) => 2""",
    imports=['"unsafe"'],
    extra="""// Pair is two 64-bit words.
type Pair struct {
	A int64
	B int64
}""",
    buggy="""addr := uintptr(unsafe.Pointer(p))
addr += unsafe.Offsetof(p.B)
return *(*int64)(unsafe.Pointer(addr))""",
    solution="""q := unsafe.Add(unsafe.Pointer(p), unsafe.Offsetof(p.B))
return *(*int64)(q)""",
    tests="""
import (
	"go/ast"
	"go/parser"
	"go/token"
	"runtime"
	"testing"
)

func TestSecondWord(t *testing.T) {
	if got := SecondWord(&Pair{A: 1, B: 2}); got != 2 {
		t.Errorf("SecondWord = %d, want 2", got)
	}
	if got := SecondWord(&Pair{}); got != 0 {
		t.Errorf("SecondWord = %d, want 0", got)
	}
	if got := SecondWord(&Pair{A: 9, B: -7}); got != -7 {
		t.Errorf("SecondWord = %d, want -7", got)
	}
}

func TestSecondWordUnderAllocationPressure(t *testing.T) {
	for i := 0; i < 2000; i++ {
		p := &Pair{A: int64(i), B: int64(i * 2)}
		if got := SecondWord(p); got != int64(i*2) {
			t.Fatalf("iteration %d: SecondWord = %d, want %d", i, got, i*2)
		}
		if i%100 == 0 {
			runtime.GC()
		}
		runtime.KeepAlive(p)
	}
}

func TestSecondWordUsesUnsafeAdd(t *testing.T) {
	// A moving collector would invalidate a stored uintptr, and no test can
	// provoke that on demand -- so the rule is checked in the source instead.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "uintptrhold.go", nil, 0)
	if err != nil {
		t.Skipf("cannot parse the source: %v", err)
	}
	usesAdd := false
	holdsUintptr := false
	ast.Inspect(f, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok && id.Name == "unsafe" && sel.Sel.Name == "Add" {
				usesAdd = true
			}
		}
		// A uintptr conversion assigned to a variable outlives its expression.
		record := func(rhs []ast.Expr) {
			for _, e := range rhs {
				ast.Inspect(e, func(m ast.Node) bool {
					c, ok := m.(*ast.CallExpr)
					if !ok {
						return true
					}
					if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "uintptr" {
						holdsUintptr = true
					}
					return true
				})
			}
		}
		switch s := n.(type) {
		case *ast.AssignStmt:
			record(s.Rhs)
		case *ast.ValueSpec:
			record(s.Values)
		}
		return true
	})
	if holdsUintptr {
		t.Error("a uintptr conversion is stored in a variable: the address goes stale outside its expression")
	}
	if !usesAdd {
		t.Error("the offset must be applied with unsafe.Add, which keeps the result a pointer")
	}
}
""",
    context="A struct walker computes field addresses as integers. It works for months, then starts returning zeros on a build with a different inliner — the kind of bug that is never reproduced on demand.",
    task=[
        "Return `p.B`, reached through the field's offset rather than the selector.",
        "Fix the single bug: the arithmetic must never leave `unsafe.Pointer`.",
    ],
    examples=[
        ("SecondWord(&Pair{A:1, B:2})", "2", None),
        ("SecondWord(&Pair{})", "0", None),
        ("2000 iterations with GC", "every read correct", None),
    ],
    topics=[
        ("uintptr is not a pointer", "The collector does not see it, so the object it names can be freed or moved."),
        ("unsafe.Add", "The supported way to offset a pointer; the result stays a pointer."),
        ("The valid uintptr pattern", "Converting to `uintptr` and back is only defined within a single expression."),
        ("runtime.KeepAlive", "Needed when an object's last use is through an address the collector cannot follow."),
    ],
    hint="Two statements is one statement too many. What is the pointer during the second one?",
    intuition="`unsafe.Pointer` is a pointer the garbage collector understands; `uintptr` is an integer that happens to look like one. Splitting the arithmetic across statements gives the runtime a window in which nothing refers to the object.",
    approach=[
        "Offset the pointer with `unsafe.Add(unsafe.Pointer(p), unsafe.Offsetof(p.B))`.",
        "Convert the result to `*int64` and read through it.",
    ],
    walkthrough="`unsafe.Add` produces a pointer into the same `Pair`, so the object stays reachable throughout. The `uintptr` version holds only a number between the two statements — correct today by luck, not by rule.",
    pitfalls=[
        "Believing that Go's collector never moves objects; the rules are written so it may.",
        "Wrapping the broken version in `runtime.KeepAlive` — that addresses liveness, not the address going stale.",
    ],
)

P(
    "senior",
    name="pointerwalk",
    title="Walk An Array By Pointer",
    sig="func SumInt32(p *int32, n int) int64",
    doc="""SumInt32 totals n consecutive int32 values starting at p.

This is the shape a C API hands you: a pointer and a count, with no
slice. n <= 0 or a nil pointer totals 0.

Examples:

	SumInt32(&a[0], 3) => a[0] + a[1] + a[2]""",
    imports=['"unsafe"'],
    solution="""if p == nil || n <= 0 {
	return 0
}
var total int64
for i := 0; i < n; i++ {
	q := (*int32)(unsafe.Add(unsafe.Pointer(p), uintptr(i)*unsafe.Sizeof(*p)))
	total += int64(*q)
}
return total""",
    tests="""
import (
	"testing"
	"unsafe"
)

var sink int64

func TestSumInt32(t *testing.T) {
	a := []int32{1, 2, 3, 4}
	if got := SumInt32(unsafe.SliceData(a), 3); got != 6 {
		t.Errorf("SumInt32 = %d, want 6", got)
	}
	if got := SumInt32(unsafe.SliceData(a), 4); got != 10 {
		t.Errorf("SumInt32 = %d, want 10", got)
	}
}

func TestSumInt32Edges(t *testing.T) {
	a := []int32{1, 2}
	if got := SumInt32(unsafe.SliceData(a), 0); got != 0 {
		t.Errorf("SumInt32 = %d, want 0", got)
	}
	if got := SumInt32(unsafe.SliceData(a), -1); got != 0 {
		t.Errorf("SumInt32 = %d, want 0", got)
	}
	if got := SumInt32(nil, 3); got != 0 {
		t.Errorf("SumInt32(nil) = %d, want 0", got)
	}
}

func TestSumInt32Negative(t *testing.T) {
	a := []int32{-5, 5, -1}
	if got := SumInt32(unsafe.SliceData(a), 3); got != -1 {
		t.Errorf("SumInt32 = %d, want -1", got)
	}
}

func TestSumInt32WideAccumulator(t *testing.T) {
	a := make([]int32, 8)
	for i := range a {
		a[i] = 1 << 30
	}
	if got := SumInt32(unsafe.SliceData(a), 8); got != 8<<30 {
		t.Errorf("SumInt32 = %d, want %d: the total must not overflow", got, int64(8)<<30)
	}
}

func TestSumInt32AllocatesNothing(t *testing.T) {
	a := make([]int32, 1024)
	for i := range a {
		a[i] = int32(i)
	}
	p := unsafe.SliceData(a)
	if n := testing.AllocsPerRun(100, func() { sink = SumInt32(p, 1024) }); n != 0 {
		t.Errorf("SumInt32 made %v allocations, want 0", n)
	}
}
""",
    context="A binding wraps a C library that returns a pointer and a count. The Go side has to read the values without owning the memory or making a slice of it.",
    task=[
        "Total `n` consecutive int32 values starting at `p`.",
        "Return 0 for a nil pointer or `n <= 0`.",
        "Accumulate in int64 so a long run cannot overflow; allocate nothing.",
    ],
    examples=[
        ("SumInt32(&a[0], 3) over [1 2 3 4]", "6", None),
        ("SumInt32(nil, 3)", "0", None),
        ("eight values of 1<<30", "8589934592", "The accumulator is wider than the elements."),
    ],
    topics=[
        ("Pointer stride", "Advance by `i * unsafe.Sizeof(*p)`, not by `i`."),
        ("unsafe.Add", "Keeps the arithmetic in pointer space."),
        ("Accumulator width", "Summing int32 into int32 overflows long before int64 does."),
        ("No bounds check exists", "`n` is a promise from the caller; the runtime cannot verify it."),
    ],
    hint="The step between elements is `unsafe.Sizeof(*p)` bytes, not one.",
    intuition="Pointer arithmetic in Go counts bytes, not elements. Every step has to be scaled by the element size, which is exactly what indexing a slice does for you when you have one.",
    approach=[
        "Guard nil and non-positive `n`.",
        "For each `i`, offset `p` by `i * unsafe.Sizeof(*p)` and read through the typed pointer.",
        "Accumulate into an int64.",
    ],
    walkthrough="For four int32s, the offsets are 0, 4, 8 and 12 bytes. Adding `i` instead of `i*4` would read overlapping, misaligned values.",
    pitfalls=[
        "`unsafe.Add(p, i)` — a one-byte step through a four-byte array.",
        "`unsafe.Slice(p, n)` is the idiomatic answer in real code; here the loop is the exercise.",
    ],
)

P(
    "senior",
    name="growcheck",
    title="Detect Whether Append Reallocated",
    sig="func Grew(before, after []int) bool",
    doc="""Grew reports whether after occupies different storage from before —
that is, whether the append that produced it had to reallocate.

Examples:

	s := make([]int, 0, 1); Grew(s, append(s, 1)) => false""",
    imports=['"unsafe"'],
    solution="""if cap(before) == 0 || cap(after) == 0 {
	return cap(before) != cap(after)
}
return unsafe.SliceData(before) != unsafe.SliceData(after)""",
    tests="""
import "testing"

func TestGrewWhenCapacityWasEnough(t *testing.T) {
	s := make([]int, 0, 4)
	if Grew(s, append(s, 1)) {
		t.Error("Grew = true, want false: the capacity was sufficient")
	}
}

func TestGrewWhenCapacityRanOut(t *testing.T) {
	s := make([]int, 1, 1)
	if !Grew(s, append(s, 2)) {
		t.Error("Grew = false, want true: append had to reallocate")
	}
}

func TestGrewFromNil(t *testing.T) {
	var s []int
	if !Grew(s, append(s, 1)) {
		t.Error("Grew = false, want true: a nil slice has no storage")
	}
}

func TestGrewIdentity(t *testing.T) {
	s := make([]int, 2, 4)
	if Grew(s, s) {
		t.Error("Grew(s, s) = true, want false")
	}
	if Grew(s, s[:1]) {
		t.Error("Grew(s, s[:1]) = true, want false: reslicing does not reallocate")
	}
}

func TestGrewAcrossManyAppends(t *testing.T) {
	s := make([]int, 0, 2)
	grew := 0
	for i := 0; i < 64; i++ {
		next := append(s, i)
		if Grew(s, next) {
			grew++
		}
		s = next
	}
	if grew == 0 {
		t.Error("no growth detected across 64 appends, want several")
	}
	if grew > 10 {
		t.Errorf("detected %d reallocations, want the handful append actually performs", grew)
	}
}
""",
    context="A benchmark blames `append` for a copy nobody can see. Proving which appends actually reallocate needs a way to compare two slices' storage, which the language does not offer.",
    task=[
        "Report whether `after` sits in different storage from `before`.",
        "A slice with no capacity has no storage; treat that consistently.",
        "Reslicing without appending must report false.",
    ],
    examples=[
        ("s := make([]int,0,4); Grew(s, append(s,1))", "false", "The capacity was enough."),
        ("s := make([]int,1,1); Grew(s, append(s,2))", "true", "append reallocated."),
        ("Grew(s, s[:1])", "false", "Reslicing keeps the array."),
    ],
    topics=[
        ("append's growth policy", "It reallocates only when the capacity is exhausted."),
        ("SliceData identity", "The data pointer is the only observable identity of the backing array."),
        ("Zero-capacity slices", "There is no array to point at, so the pointer comparison is not meaningful."),
    ],
    hint="Compare the data pointers — after deciding what a zero-capacity slice means.",
    intuition="Whether `append` copied is invisible through the language's own API: the length and capacity change either way. The backing array's address is the one observable that distinguishes them.",
    approach=[
        "Handle the zero-capacity cases explicitly.",
        "Otherwise compare `unsafe.SliceData(before)` with `unsafe.SliceData(after)`.",
    ],
    walkthrough="Appending to a slice with spare capacity returns a header with the same data pointer, so the comparison is false. When the capacity runs out, `append` allocates and the pointers differ.",
    pitfalls=[
        "Comparing capacities instead — growth changes the capacity, but so does a three-index reslice.",
        "Forgetting the nil case, where both data pointers are nil.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="arenastrings",
    title="Keys That Were Overwritten By The Next Batch",
    mode="bug",
    sig="func Intern(arena []byte, spans [][2]int) []string",
    doc="""Intern returns one string per span of arena, for the caller to keep.

arena is a scratch region the caller refills between batches, so the
strings must own their bytes. Copying them into one block keeps the cost
to a single allocation for the whole batch.

Examples:

	Intern([]byte("abcd"), [][2]int{{0,2},{2,4}}) => []string{"ab", "cd"}""",
    imports=['"unsafe"'],
    buggy="""if len(spans) == 0 {
	return nil
}
out := make([]string, 0, len(spans))
for _, sp := range spans {
	lo, hi := sp[0], sp[1]
	if lo < 0 || hi > len(arena) || lo > hi {
		out = append(out, "")
		continue
	}
	if lo == hi {
		out = append(out, "")
		continue
	}
	out = append(out, unsafe.String(unsafe.SliceData(arena[lo:hi]), hi-lo))
}
return out""",
    solution="""if len(spans) == 0 {
	return nil
}
total := 0
for _, sp := range spans {
	lo, hi := sp[0], sp[1]
	if lo < 0 || hi > len(arena) || lo >= hi {
		continue
	}
	total += hi - lo
}
block := make([]byte, 0, total)
out := make([]string, 0, len(spans))
for _, sp := range spans {
	lo, hi := sp[0], sp[1]
	if lo < 0 || hi > len(arena) || lo >= hi {
		out = append(out, "")
		continue
	}
	start := len(block)
	block = append(block, arena[lo:hi]...)
	out = append(out, unsafe.String(unsafe.SliceData(block[start:]), hi-lo))
}
return out""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestIntern(t *testing.T) {
	got := Intern([]byte("abcd"), [][2]int{{0, 2}, {2, 4}})
	if len(got) != 2 || got[0] != "ab" || got[1] != "cd" {
		t.Errorf("Intern = %q, want [ab cd]", got)
	}
	if got := Intern(nil, nil); got != nil {
		t.Errorf("Intern = %q, want nil", got)
	}
}

func TestInternBadSpans(t *testing.T) {
	got := Intern([]byte("abcd"), [][2]int{{-1, 2}, {0, 99}, {3, 1}, {2, 2}})
	for i, s := range got {
		if s != "" {
			t.Errorf("span %d = %q, want empty", i, s)
		}
	}
}

func TestInternSurvivesArenaReuse(t *testing.T) {
	arena := make([]byte, 8)
	copy(arena, "firstrun")
	keys := Intern(arena, [][2]int{{0, 5}, {5, 8}})
	copy(arena, "OVERWRIT")
	if keys[0] != "first" || keys[1] != "run" {
		t.Errorf("keys = %q, want [first run]: the strings view the reused arena", keys)
	}
}

func TestInternKeysStayValidAsMapKeys(t *testing.T) {
	arena := make([]byte, 4)
	m := make(map[string]int)
	for i := 0; i < 26; i++ {
		for j := range arena {
			arena[j] = byte('a' + i)
		}
		for _, k := range Intern(arena, [][2]int{{0, 4}}) {
			m[k] = i
		}
	}
	if len(m) != 26 {
		t.Fatalf("map has %d keys, want 26: the keys changed after insertion", len(m))
	}
	for i := 0; i < 26; i++ {
		k := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if got, ok := m[k]; !ok || got != i {
			t.Fatalf("m[%q] = %d, %v, want %d, true", k, got, ok, i)
		}
	}
}

func TestInternDoesNotAliasTheArena(t *testing.T) {
	arena := []byte("abcdef")
	got := Intern(arena, [][2]int{{0, 3}})
	if unsafe.StringData(got[0]) == unsafe.SliceData(arena) {
		t.Error("the result aliases the arena")
	}
}

func TestInternUsesOneBlock(t *testing.T) {
	arena := make([]byte, 256)
	spans := make([][2]int, 32)
	for i := range spans {
		spans[i] = [2]int{i * 8, i*8 + 8}
	}
	var sink []string
	n := testing.AllocsPerRun(50, func() { sink = Intern(arena, spans) })
	_ = sink
	if n > 3 {
		t.Errorf("Intern made %v allocations for 32 spans, want a handful: copy into one block", n)
	}
}
""",
    context="A column store slices its read arena into keys and inserts them into a map. Lookups start missing entries that are visibly present, and the map's length grows past the number of distinct keys.",
    task=[
        "Return one string per span, for the caller to keep past the next batch.",
        "An out-of-range, inverted or empty span yields the empty string.",
        "Fix the single bug so the strings do not view the arena.",
        "Copy the batch into one block: a handful of allocations for 32 spans, not one per span.",
    ],
    examples=[
        ('Intern([]byte("abcd"), [][2]int{{0,2},{2,4}})', '["ab" "cd"]', None),
        ('keys := Intern(arena, ...); copy(arena, "OVERWRIT")', "keys unchanged", None),
        ("a span of {3,1}", '""', "Inverted spans are rejected."),
    ],
    topics=[
        ("Map keys must be stable", "A map caches each key's hash; mutating a key's bytes strands the entry."),
        ("Arena copying", "One block plus per-span headers beats one allocation per span."),
        ("unsafe.String over your own block", "Legal precisely because nothing else can write to it."),
        ("Sub-slicing the block", "Each string wraps `block[start:]` with its own length."),
    ],
    hint="The conversion is right; the bytes are the caller's. Where should the batch's bytes live?",
    intuition="`unsafe.String` is only sound over bytes that will never change. An arena is the opposite of that. The fix is not to abandon the technique but to give it memory that qualifies — one block per batch, written once.",
    approach=[
        "Sum the valid spans' lengths for the block's capacity.",
        "Append each valid span's bytes into the block, remembering where it started.",
        "Wrap `block[start:]` with `unsafe.String` and the span's length.",
    ],
    walkthrough="Before the fix, twenty-six batches insert twenty-six keys that all point into a four-byte arena, so the map ends up full of entries whose keys now read \"zzzz\". After it, each batch's bytes are copied once and the keys never change.",
    pitfalls=[
        "Appending to `block` after taking a string over it — a reallocation would leave the earlier strings pointing at the old array. Sizing the block up front is what makes this safe.",
        "Allocating per span, which is correct and defeats the purpose.",
    ],
)

P(
    "staff",
    name="paddedcounters",
    title="Pad To The Line, Not To The Word",
    sig="func Run(workers, iters int) int64",
    doc="""Run gives each of workers goroutines its own Slot, has each increment
its counter iters times, and returns the total.

Slot must be padded so no two counters share a cache line; the padding is
computed from the counter's own size, not hard-coded.

Examples:

	Run(4, 1000) => 4000""",
    imports=['"sync"', '"unsafe"'],
    extra="""// LineSize is the coherence granule the counters must not share.
const LineSize = 64

// Slot is one worker's counter, padded to its own cache line.
type Slot struct {
	N   int64
	Pad [LineSize - unsafe.Sizeof(int64(0))]byte
}""",
    solution="""if workers < 1 || iters < 0 {
	return 0
}
slots := make([]Slot, workers)
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	go func(s *Slot) {
		defer wg.Done()
		for i := 0; i < iters; i++ {
			s.N++
		}
	}(&slots[w])
}
wg.Wait()
var total int64
for i := range slots {
	total += slots[i].N
}
return total""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestRunTotals(t *testing.T) {
	if got := Run(4, 1000); got != 4000 {
		t.Errorf("Run = %d, want 4000", got)
	}
	if got := Run(1, 0); got != 0 {
		t.Errorf("Run = %d, want 0", got)
	}
	if got := Run(0, 10); got != 0 {
		t.Errorf("Run = %d, want 0", got)
	}
	if got := Run(3, -1); got != 0 {
		t.Errorf("Run = %d, want 0", got)
	}
}

func TestRunUnderLoad(t *testing.T) {
	if got := Run(8, 200000); got != 1600000 {
		t.Errorf("Run = %d, want 1600000", got)
	}
}

func TestSlotOccupiesAWholeLine(t *testing.T) {
	if got := unsafe.Sizeof(Slot{}); got != LineSize {
		t.Errorf("sizeof(Slot) = %d, want %d", got, LineSize)
	}
}

func TestSlotsAreALineApart(t *testing.T) {
	s := make([]Slot, 2)
	a := uintptr(unsafe.Pointer(&s[0]))
	b := uintptr(unsafe.Pointer(&s[1]))
	if b-a != LineSize {
		t.Errorf("stride = %d, want %d: neighbouring counters share a line", b-a, LineSize)
	}
}

func TestCounterFieldIsFirst(t *testing.T) {
	if off := unsafe.Offsetof(Slot{}.N); off != 0 {
		t.Errorf("N is at offset %d, want 0", off)
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(4, 20000)
	}
}
""",
    context="A sharded counter scales negatively: eight workers are slower than one. The counters are separate variables, no lock is held, and the race detector is silent.",
    task=[
        "Give each worker its own `Slot` and have it increment `N` `iters` times.",
        "Join, then return the sum of the slots.",
        "Non-positive `workers` or negative `iters` return 0.",
        "The slot stride must be a full cache line — the padding follows `unsafe.Sizeof`, not a magic number.",
    ],
    examples=[
        ("Run(4, 1000)", "4000", None),
        ("sizeof(Slot)", "64", "One counter per line."),
        ("&s[1] - &s[0]", "64", None),
    ],
    topics=[
        ("False sharing", "Coherence works per line, so neighbours on one line serialise."),
        ("unsafe.Sizeof in a constant expression", "The padding adapts if the counter's type changes."),
        ("Array stride", "The element size is the struct's size, which is what separates the slots."),
        ("Join before summing", "`wg.Wait()` is what makes the counters safe to read."),
    ],
    hint="The padded type is written for you. The body is allocate, fan out, join, fold.",
    intuition="The hardware does not know about your variables, only about lines. Making each counter occupy a whole line means each core owns its line outright, and the increments never leave L1.",
    approach=[
        "Validate the arguments.",
        "Allocate `workers` slots and start one goroutine per slot, passing its address.",
        "Increment `s.N` `iters` times; `Wait`, then sum the slots.",
    ],
    walkthrough="Eight unpadded int64 counters fit in one 64-byte line, so 1.6 million increments become 1.6 million coherence transactions. With one slot per line each core writes only its own.",
    pitfalls=[
        "Capturing the loop index and indexing `slots[w]` inside the goroutine — correct, but passing the pointer makes the exclusive ownership explicit.",
        "Writing `Pad [56]byte`; the constant expression keeps the padding right if `N` changes type.",
        "Summing before `Wait`.",
    ],
)

P(
    "staff",
    name="sharedview",
    title="Many Readers Over One Reinterpreted Buffer",
    sig="func SumParallel(b []byte, workers int) (int64, bool)",
    doc="""SumParallel reinterprets b as int64 values and totals them using
workers goroutines over disjoint chunks.

The view shares b's storage, so nothing may write through it. Concurrent
reads of shared memory need no synchronisation at all.

Examples:

	SumParallel(sixteenBytes, 2) => the two int64s summed, true""",
    imports=['"sync"', '"unsafe"'],
    solution="""if len(b) == 0 || len(b)%8 != 0 {
	return 0, false
}
p := unsafe.Pointer(unsafe.SliceData(b))
if uintptr(p)&7 != 0 {
	return 0, false
}
view := unsafe.Slice((*int64)(p), len(b)/8)
if workers < 1 {
	workers = 1
}
if workers > len(view) {
	workers = len(view)
}
partial := make([]int64, workers)
size := (len(view) + workers - 1) / workers
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	start := w * size
	end := start + size
	if start > len(view) {
		start = len(view)
	}
	if end > len(view) {
		end = len(view)
	}
	go func(w int, part []int64) {
		defer wg.Done()
		var sum int64
		for _, v := range part {
			sum += v
		}
		partial[w] = sum
	}(w, view[start:end])
}
wg.Wait()
var total int64
for _, v := range partial {
	total += v
}
return total, true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func makeInt64Bytes(vals []int64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(vals))), len(vals)*8)
}

func TestSumParallelSmall(t *testing.T) {
	vals := []int64{1, 2, 3, 4}
	got, ok := SumParallel(makeInt64Bytes(vals), 2)
	if !ok || got != 10 {
		t.Errorf("SumParallel = %d, %v, want 10, true", got, ok)
	}
}

func TestSumParallelWorkerCounts(t *testing.T) {
	vals := make([]int64, 1001)
	var want int64
	for i := range vals {
		vals[i] = int64(i)
		want += int64(i)
	}
	b := makeInt64Bytes(vals)
	for _, w := range []int{0, 1, 2, 7, 64, 100000} {
		got, ok := SumParallel(b, w)
		if !ok || got != want {
			t.Fatalf("workers=%d: SumParallel = %d, %v, want %d, true", w, got, ok, want)
		}
	}
}

func TestSumParallelRejectsBadShapes(t *testing.T) {
	vals := []int64{1, 2}
	b := makeInt64Bytes(vals)
	for _, c := range []struct {
		name string
		in   []byte
	}{
		{"nil", nil},
		{"empty", b[:0]},
		{"length not a multiple of 8", b[:12]},
		{"misaligned", b[1:9]},
	} {
		if _, ok := SumParallel(c.in, 2); ok {
			t.Errorf("%s: reported ok, want false", c.name)
		}
	}
}

func TestSumParallelIsRepeatable(t *testing.T) {
	vals := make([]int64, 4096)
	for i := range vals {
		vals[i] = int64(i % 97)
	}
	b := makeInt64Bytes(vals)
	first, ok := SumParallel(b, 8)
	if !ok {
		t.Fatal("SumParallel reported false")
	}
	for i := 0; i < 20; i++ {
		got, _ := SumParallel(b, 8)
		if got != first {
			t.Fatalf("run %d = %d, want %d: the workers overlap", i, got, first)
		}
	}
}

func TestSumParallelDoesNotCopy(t *testing.T) {
	vals := make([]int64, 1<<16)
	b := makeInt64Bytes(vals)
	var sink int64
	n := testing.AllocsPerRun(20, func() { sink, _ = SumParallel(b, 4) })
	_ = sink
	if n > 12 {
		t.Errorf("SumParallel made %v allocations for a 512 KiB buffer, want a handful: view, do not copy", n)
	}
}
""",
    context="An analytics kernel receives half a gigabyte of packed int64 records and wants every core summing at once, without first copying the buffer into a typed slice.",
    task=[
        "Reinterpret `b` as `[]int64` and total it across `workers` goroutines over disjoint chunks.",
        "Report false for an empty buffer, a length that is not a multiple of eight, or a misaligned start.",
        "Any worker count must produce the same total; nothing may be counted twice.",
        "The view must share `b`'s storage — no copy of the data.",
    ],
    examples=[
        ("SumParallel(bytesOf([1,2,3,4]), 2)", "10, true", None),
        ("workers = 100000 over 1001 values", "the same total, true", "The worker count is clamped."),
        ("a 12-byte buffer", "0, false", "Not a multiple of 8."),
    ],
    topics=[
        ("Reinterpreting a buffer", "`unsafe.Slice` over the data pointer with an element count."),
        ("Concurrent reads are free", "Disjoint, read-only access needs no synchronisation."),
        ("Chunk clamping", "The last chunk must stop at the view's length."),
        ("Preconditions before parallelism", "Alignment and divisibility are decided once, up front."),
    ],
    hint="Validate and build the view first; from there it is an ordinary parallel sum over a slice.",
    intuition="Reinterpretation and parallelism compose cleanly because both are read-only: the view aliases the caller's bytes, and disjoint readers of shared memory never need to coordinate. The only synchronisation is the join.",
    approach=[
        "Reject an empty or non-multiple-of-eight length and a misaligned data pointer.",
        "Build the view with `unsafe.Slice((*int64)(p), len(b)/8)`.",
        "Clamp `workers`, split into ceiling-divided clamped chunks, sum into per-worker slots.",
        "`Wait`, then fold the partials.",
    ],
    walkthrough="A 512 KiB buffer becomes a 65536-element view with no copy. Four workers each sum 16384 elements from memory they alone read, and the four partials are added after the join.",
    pitfalls=[
        "Passing `len(b)` as the element count, which produces a view eight times too long.",
        "Writing through the view — the caller's bytes may be a read-only mapping, and other readers would see the change.",
        "Folding the partials before `Wait`.",
    ],
)

P(
    "staff",
    name="zerocopykeys",
    title="Look Up Without Copying, Store With A Copy",
    sig="func Count(m map[string]int, keys [][]byte)",
    doc="""Count increments m's counter for each key.

The lookup may borrow the key's bytes, but a key that ends up stored in
the map must own its bytes: the caller reuses the buffers the keys point
into.

Examples:

	Count(m, [][]byte{[]byte("a")}) => m["a"] == 1""",
    imports=['"unsafe"'],
    solution="""for _, k := range keys {
	if len(k) == 0 {
		continue
	}
	view := unsafe.String(unsafe.SliceData(k), len(k))
	if n, ok := m[view]; ok {
		m[view] = n + 1
		continue
	}
	owned := make([]byte, len(k))
	copy(owned, k)
	m[unsafe.String(unsafe.SliceData(owned), len(owned))] = 1
}""",
    tests="""
import (
	"testing"
)

func TestCount(t *testing.T) {
	m := map[string]int{}
	Count(m, [][]byte{[]byte("a"), []byte("b"), []byte("a")})
	if m["a"] != 2 || m["b"] != 1 {
		t.Errorf("m = %v, want map[a:2 b:1]", m)
	}
}

func TestCountSkipsEmptyKeys(t *testing.T) {
	m := map[string]int{}
	Count(m, [][]byte{nil, {}, []byte("x")})
	if len(m) != 1 || m["x"] != 1 {
		t.Errorf("m = %v, want map[x:1]", m)
	}
}

func TestCountKeysSurviveBufferReuse(t *testing.T) {
	m := map[string]int{}
	buf := make([]byte, 4)
	for i := 0; i < 26; i++ {
		for j := range buf {
			buf[j] = byte('a' + i)
		}
		Count(m, [][]byte{buf})
	}
	if len(m) != 26 {
		t.Fatalf("map has %d keys, want 26: the stored keys view the reused buffer", len(m))
	}
	for i := 0; i < 26; i++ {
		k := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if m[k] != 1 {
			t.Fatalf("m[%q] = %d, want 1", k, m[k])
		}
	}
}

func TestCountRepeatedKeysDoNotAllocate(t *testing.T) {
	m := map[string]int{}
	key := []byte("stable-key")
	Count(m, [][]byte{key})
	batch := [][]byte{key}
	if n := testing.AllocsPerRun(200, func() { Count(m, batch) }); n != 0 {
		t.Errorf("Count made %v allocations for an existing key, want 0: borrow the bytes for the lookup", n)
	}
	if m["stable-key"] < 200 {
		t.Errorf("counter = %d, want at least 200", m["stable-key"])
	}
}

func TestCountAccumulates(t *testing.T) {
	m := map[string]int{"a": 5}
	Count(m, [][]byte{[]byte("a")})
	if m["a"] != 6 {
		t.Errorf("m[a] = %d, want 6", m["a"])
	}
}
""",
    context="A counter over a hot byte stream copies every key twice: once to look it up and once again when it is already present. Ninety-nine percent of the keys are already present.",
    task=[
        "Increment `m`'s counter for each key.",
        "An existing key must be counted without allocating.",
        "A key stored for the first time must own its bytes — the caller reuses the buffers.",
        "Empty and nil keys are skipped.",
    ],
    examples=[
        ('Count(m, [][]byte{[]byte("a"), []byte("a")})', "m[a] == 2", None),
        ("200 counts of an existing key", "0 allocations", None),
        ("26 batches through one reused buffer", "26 distinct keys", None),
    ],
    topics=[
        ("Borrowing for a lookup", "The borrowed string dies inside the call, so nothing can observe the aliasing."),
        ("Owning for a store", "A stored key is read again later; it must not be able to change."),
        ("Map keys cache their hash", "A key whose bytes change strands its entry permanently."),
        ("The asymmetry is the whole design", "Read paths may alias; write paths may not."),
    ],
    hint="Two paths through the loop. Only one of them may allocate.",
    intuition="Zero-copy is not a property of the conversion but of the lifetime. A string that lives only long enough to hash and compare can safely borrow; one that the map keeps must outlive the buffer, so it has to own its bytes.",
    approach=[
        "Skip empty keys.",
        "Build a borrowed string view for the lookup.",
        "On a hit, increment through the same view.",
        "On a miss, copy the bytes into a private slice and store a string over the copy.",
    ],
    walkthrough="For a key already in the map, the borrowed view hashes, matches, and is discarded — no allocation. For a new key, four bytes are copied and the stored key points at memory the caller cannot reach.",
    pitfalls=[
        "Storing the borrowed view, which is the bug the reuse test exists to catch.",
        "Copying on every key, which is correct and gives up the entire optimisation.",
        "Reusing the borrowed view as the stored key on the miss path — it must be the copy.",
    ],
)

P(
    "staff",
    name="structbytes",
    title="A Byte View Of A Struct, Only When It Is Safe",
    sig="func Bytes(p *Frame) ([]byte, bool)",
    doc="""Bytes returns a byte view of the frame p points at, for writing to a
socket without an intermediate copy.

This is only defined when the struct contains no pointers: a byte view of
a pointer field would let the bytes outlive what they point at, and would
hand the peer an address. Report false rather than producing one.

Examples:

	Bytes(&Frame{}) => a view of unsafe.Sizeof(Frame{}) bytes, true""",
    imports=['"reflect"', '"unsafe"'],
    extra="""// Frame is a fixed-layout wire frame of scalars only.
type Frame struct {
	Kind  uint32
	Seq   uint32
	Stamp int64
}

// hasPointers reports whether t contains any pointer-shaped field.
func hasPointers(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Pointer, reflect.UnsafePointer, reflect.Slice, reflect.Map,
		reflect.Chan, reflect.Func, reflect.Interface, reflect.String:
		return true
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if hasPointers(t.Field(i).Type) {
				return true
			}
		}
		return false
	case reflect.Array:
		return hasPointers(t.Elem())
	default:
		return false
	}
}""",
    solution="""if p == nil {
	return nil, false
}
t := reflect.TypeOf(*p)
if hasPointers(t) {
	return nil, false
}
return unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p)), true""",
    tests="""
import (
	"reflect"
	"testing"
	"unsafe"
)

func TestBytesShape(t *testing.T) {
	f := &Frame{Kind: 1, Seq: 2, Stamp: 3}
	b, ok := Bytes(f)
	if !ok {
		t.Fatal("Bytes reported false for a pointer-free struct")
	}
	if uintptr(len(b)) != unsafe.Sizeof(*f) {
		t.Errorf("len = %d, want %d", len(b), unsafe.Sizeof(*f))
	}
	if uintptr(cap(b)) != unsafe.Sizeof(*f) {
		t.Errorf("cap = %d, want %d: an append must not run past the struct", cap(b), unsafe.Sizeof(*f))
	}
}

func TestBytesSharesTheStruct(t *testing.T) {
	f := &Frame{}
	b, ok := Bytes(f)
	if !ok {
		t.Fatal("Bytes reported false")
	}
	f.Kind = 0x01020304
	found := false
	for _, x := range b[:4] {
		if x != 0 {
			found = true
		}
	}
	if !found {
		t.Error("the view does not share the struct's memory")
	}
}

func TestBytesNil(t *testing.T) {
	if _, ok := Bytes(nil); ok {
		t.Error("Bytes(nil) reported ok, want false")
	}
}

func TestBytesDoesNotAllocate(t *testing.T) {
	f := &Frame{}
	var sink []byte
	if n := testing.AllocsPerRun(200, func() { sink, _ = Bytes(f) }); n != 0 {
		t.Errorf("Bytes made %v allocations, want 0", n)
	}
	_ = sink
}

func TestHasPointersFixture(t *testing.T) {
	type withPtr struct {
		A int
		B *int
	}
	type withString struct{ S string }
	type nested struct{ Inner withPtr }
	type clean struct {
		A int32
		B [4]byte
	}
	cases := []struct {
		v    any
		want bool
	}{
		{withPtr{}, true},
		{withString{}, true},
		{nested{}, true},
		{clean{}, false},
		{Frame{}, false},
	}
	for _, c := range cases {
		if got := hasPointers(reflect.TypeOf(c.v)); got != c.want {
			t.Errorf("hasPointers(%T) = %v, want %v", c.v, got, c.want)
		}
	}
}
""",
    context="A wire encoder writes frames by copying each field into a buffer. Profiling says the copy is most of the send path, and every frame is scalars only.",
    task=[
        "Return a byte view of `*p`, of exactly `unsafe.Sizeof(*p)` bytes.",
        "Return false for a nil pointer, and for any struct type that contains a pointer-shaped field.",
        "Zero allocations; the view's capacity must equal its length.",
    ],
    examples=[
        ("Bytes(&Frame{})", "a 16-byte view, true", None),
        ("Bytes(nil)", "nil, false", None),
        ("a struct containing a *int", "nil, false", "A byte view would expose an address."),
    ],
    topics=[
        ("Pointer-free is the precondition", "Bytes of a pointer are an address, meaningless to a peer and invisible to the collector."),
        ("unsafe.Slice over a struct pointer", "Length and capacity both come from `Sizeof`."),
        ("Padding is in the view", "The bytes include the struct's padding, whose contents are unspecified."),
        ("Endianness and layout", "The view is the machine's layout, not a portable wire format."),
    ],
    hint="`hasPointers` is written for you. Guard, check, then build the view.",
    intuition="Reinterpreting a struct as bytes is sound only when the bytes mean the same thing to whoever reads them. A pointer's bytes do not: they name an address in this process, and copying them out both leaks a layout detail and hides a reference from the collector.",
    approach=[
        "Reject a nil pointer.",
        "Reject a type that `hasPointers` reports on.",
        "Return `unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p))`.",
    ],
    walkthrough="`Frame` is two uint32s and an int64 — sixteen bytes, no pointers. The view aliases the struct, so writing `f.Kind` changes the first four bytes of the slice.",
    pitfalls=[
        "Sending the view over the wire as a portable format; the layout, padding and endianness are all the local machine's.",
        "Using `len(b)` from a hard-coded number instead of `Sizeof`, which breaks the moment a field is added.",
        "Assuming the padding bytes are zero — they are whatever was there.",
    ],
)
