"""10-advanced-topics — rotation 4: 5 puzzles each for middle, senior, staff."""

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
    name="deleteat",
    title="Remove One Element, Release Its Slot",
    sig="func DeleteAt(s []*Item, i int) []*Item",
    doc="""DeleteAt removes the element at index i, preserving the order of the
rest, and returns the shortened slice.

The vacated slot at the end must be cleared so the removed item stops
being reachable through the backing array.

Examples:

	DeleteAt([]*Item{a, b, c}, 1) => a slice holding a and c""",
    extra="""// Item is one stored element.
type Item struct {
	ID  int
	Pad [512]byte
}""",
    solution="""if i < 0 || i >= len(s) {
	return s
}
copy(s[i:], s[i+1:])
s[len(s)-1] = nil
return s[:len(s)-1]""",
    tests="""
import "testing"

func TestDeleteAt(t *testing.T) {
	a, b, c := &Item{ID: 1}, &Item{ID: 2}, &Item{ID: 3}
	got := DeleteAt([]*Item{a, b, c}, 1)
	if len(got) != 2 || got[0] != a || got[1] != c {
		t.Errorf("DeleteAt = %v, want [a c]", got)
	}
}

func TestDeleteAtOrderIsPreserved(t *testing.T) {
	s := make([]*Item, 6)
	for i := range s {
		s[i] = &Item{ID: i}
	}
	got := DeleteAt(s, 0)
	for i, p := range got {
		if p.ID != i+1 {
			t.Fatalf("got[%d].ID = %d, want %d", i, p.ID, i+1)
		}
	}
}

func TestDeleteAtClearsTheTail(t *testing.T) {
	s := []*Item{{ID: 1}, {ID: 2}, {ID: 3}}
	backing := s
	DeleteAt(s, 0)
	if backing[2] != nil {
		t.Error("the vacated slot still holds an item: it stays reachable")
	}
}

func TestDeleteAtOutOfRange(t *testing.T) {
	s := []*Item{{ID: 1}}
	for _, i := range []int{-1, 1, 99} {
		if got := DeleteAt(s, i); len(got) != 1 {
			t.Errorf("DeleteAt(_, %d) = %v, want the slice unchanged", i, got)
		}
	}
}

func TestDeleteAtLast(t *testing.T) {
	s := []*Item{{ID: 1}, {ID: 2}}
	got := DeleteAt(s, 1)
	if len(got) != 1 || got[0].ID != 1 {
		t.Errorf("DeleteAt = %v, want [1]", got)
	}
	if s[1] != nil {
		t.Error("the last slot was not cleared")
	}
}

func TestDeleteAtAllocatesNothing(t *testing.T) {
	s := make([]*Item, 64)
	for i := range s {
		s[i] = &Item{ID: i}
	}
	if n := testing.AllocsPerRun(100, func() { _ = DeleteAt(s[:64], 0) }); n != 0 {
		t.Errorf("DeleteAt made %v allocations, want 0", n)
	}
}
""",
    context="A registry removes entries with `append(s[:i], s[i+1:]...)`. The entries are half a kilobyte each, and the removed one stays reachable through the slot the shortening left behind.",
    task=[
        "Remove the element at `i`, keeping the order of the rest.",
        "Clear the vacated slot at the end before shortening.",
        "An out-of-range index returns the slice unchanged; allocate nothing.",
    ],
    examples=[
        ("DeleteAt([]*Item{a,b,c}, 1)", "[a c]", None),
        ("s[2] after DeleteAt(s, 0)", "nil", "The slot the shift emptied is cleared."),
        ("DeleteAt(s, -1)", "s unchanged", None),
    ],
    topics=[
        ("Shift with copy", "`copy(s[i:], s[i+1:])` moves the tail down one place."),
        ("The last slot is now a duplicate", "After the shift it holds a second reference to a live element."),
        ("Reslicing hides, it does not release", "The pointer past the new length is still in the array."),
    ],
    hint="After the shift, what is sitting at the old last index?",
    intuition="Shifting the tail down leaves the final slot holding a copy of the element before it. Shortening the slice hides that slot from you and not from the collector, so a long-lived array keeps one stale reference per deletion.",
    approach=[
        "Reject an out-of-range index.",
        "`copy(s[i:], s[i+1:])` to close the gap.",
        "Set the last slot to nil, then return `s[:len(s)-1]`.",
    ],
    walkthrough="Deleting index 0 of [a b c] shifts b and c down, leaving c also at index 2. Clearing index 2 and returning `s[:2]` releases the duplicate.",
    pitfalls=[
        "`append(s[:i], s[i+1:]...)` — the same shift, with the same uncleared tail.",
        "Clearing after reslicing, which no longer reaches the slot.",
    ],
)

P(
    "middle",
    name="runetruncate",
    title="Cut A String Without Splitting A Character",
    sig="func Truncate(s string, n int) string",
    doc="""Truncate returns the longest prefix of s that is at most n bytes and
does not end in the middle of a UTF-8 character.

The result is a substring, so nothing is copied.

Examples:

	Truncate("héllo", 3) => "hé" """,
    imports=['"unicode/utf8"'],
    solution="""if n <= 0 {
	return ""
}
if n >= len(s) {
	return s
}
for n > 0 && !utf8.RuneStart(s[n]) {
	n--
}
return s[:n]""",
    tests="""
import (
	"testing"
	"unicode/utf8"
)

var sink string

func TestTruncateASCII(t *testing.T) {
	if got := Truncate("hello", 3); got != "hel" {
		t.Errorf("Truncate = %q, want \\"hel\\"", got)
	}
	if got := Truncate("hi", 9); got != "hi" {
		t.Errorf("Truncate = %q, want \\"hi\\"", got)
	}
	if got := Truncate("hi", 0); got != "" {
		t.Errorf("Truncate = %q, want empty", got)
	}
	if got := Truncate("hi", -1); got != "" {
		t.Errorf("Truncate = %q, want empty", got)
	}
}

func TestTruncateMultiByte(t *testing.T) {
	// "héllo": h is 1 byte, é is 2 bytes at offsets 1-2
	s := "héllo"
	cases := map[int]string{1: "h", 2: "h", 3: "hé", 4: "hél"}
	for n, want := range cases {
		if got := Truncate(s, n); got != want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", s, n, got, want)
		}
	}
}

func TestTruncateAlwaysValid(t *testing.T) {
	s := "日本語テキスト"
	for n := 0; n <= len(s)+2; n++ {
		got := Truncate(s, n)
		if !utf8.ValidString(got) {
			t.Fatalf("Truncate(%d) = %q, which is not valid UTF-8", n, got)
		}
		if len(got) > n && n <= len(s) {
			t.Fatalf("Truncate(%d) returned %d bytes", n, len(got))
		}
	}
}

func TestTruncateAllocatesNothing(t *testing.T) {
	s := "a fairly long string with some ünicode in it"
	if n := testing.AllocsPerRun(200, func() { sink = Truncate(s, 20) }); n != 0 {
		t.Errorf("Truncate made %v allocations, want 0: return a substring", n)
	}
}
""",
    context="A log line is cut to 200 bytes for display. Non-ASCII messages end in a broken character, and the terminal shows a replacement glyph for the last one.",
    task=[
        "Return the longest prefix of `s` that is at most `n` bytes and ends on a character boundary.",
        "`n <= 0` returns the empty string; `n >= len(s)` returns `s`.",
        "The result must be a substring — zero allocations.",
    ],
    examples=[
        ('Truncate("hello", 3)', '"hel"', None),
        ('Truncate("héllo", 2)', '"h"', "Cutting at 2 would split é, so back up."),
        ('Truncate("hi", 9)', '"hi"', None),
    ],
    topics=[
        ("UTF-8 continuation bytes", "Every byte of a character after the first has the top bits 10."),
        ("utf8.RuneStart", "Reports whether a byte can begin a character."),
        ("Substrings do not copy", "Backing up an index costs nothing."),
    ],
    hint="Walk `n` backwards while `s[n]` is not the start of a character.",
    intuition="UTF-8 is self-synchronising: from any byte you can tell whether it begins a character. So finding a safe cut point is a short walk backwards, never a re-scan from the start.",
    approach=[
        "Handle `n <= 0` and `n >= len(s)`.",
        "While `s[n]` is not a rune start, decrement `n`.",
        "Return `s[:n]`.",
    ],
    walkthrough='In "héllo", é occupies bytes 1 and 2. Cutting at n = 2 lands on the continuation byte, so n drops to 1 and the result is "h".',
    pitfalls=[
        "Converting to `[]rune` to count — correct, and it allocates the whole string again.",
        "Checking `s[n-1]` instead of `s[n]`; the question is whether the byte you are cutting *before* starts a character.",
    ],
)

P(
    "middle",
    name="reflectlen",
    title="Length Of Whatever You Are Handed",
    sig="func Length(v any) (int, bool)",
    doc="""Length returns the length of v when it has one — a string, slice,
array, map or channel — and reports false otherwise.

Examples:

	Length([]int{1, 2}) => 2, true""",
    imports=['"reflect"'],
    solution="""rv := reflect.ValueOf(v)
switch rv.Kind() {
case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
	return rv.Len(), true
default:
	return 0, false
}""",
    tests="""
import "testing"

func TestLength(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	cases := []struct {
		in   any
		n    int
		ok   bool
	}{
		{[]int{1, 2}, 2, true},
		{"héllo", 6, true},
		{[3]int{}, 3, true},
		{map[string]int{"a": 1, "b": 2}, 2, true},
		{ch, 1, true},
		{[]int(nil), 0, true},
		{map[string]int(nil), 0, true},
		{"", 0, true},
		{3, 0, false},
		{nil, 0, false},
		{struct{}{}, 0, false},
		{&[3]int{}, 0, false},
	}
	for _, c := range cases {
		n, ok := Length(c.in)
		if n != c.n || ok != c.ok {
			t.Errorf("Length(%#v) = %d, %v, want %d, %v", c.in, n, ok, c.n, c.ok)
		}
	}
}

func TestLengthCountsBytesNotRunes(t *testing.T) {
	if n, _ := Length("日本"); n != 6 {
		t.Errorf("Length = %d, want 6: a string's length is in bytes", n)
	}
}

func TestLengthDoesNotPanic(t *testing.T) {
	for _, in := range []any{nil, 3, 1.5, true, func() {}, struct{ A int }{}} {
		_, _ = Length(in)
	}
}
""",
    context="A validation layer checks that a request field is not empty. The field can be a string, a slice or a map depending on the endpoint, and each one got its own copy of the same check.",
    task=[
        "Return `v`'s length when it has one, and whether it did.",
        "Strings, slices, arrays, maps and channels have lengths; nothing else does.",
        "Never panic — a nil interface reports 0, false.",
    ],
    examples=[
        ("Length([]int{1,2})", "2, true", None),
        ('Length("日本")', "6, true", "A string's length is in bytes."),
        ("Length(3)", "0, false", None),
    ],
    topics=[
        ("Value.Len", "Defined only for the five kinds that have a length; it panics on the rest."),
        ("Kind switching as a guard", "The switch is both the dispatch and the safety check."),
        ("Nil containers have length 0", "A nil slice or map is still a valid Value of that kind."),
    ],
    hint="A kind switch with five cases in it and a default that says no.",
    intuition="Reflection will answer \"how long is this\" for any type that has an answer — but only if you ask about the right kinds. The switch turns a potential panic into a second return value.",
    approach=[
        "Take `reflect.ValueOf(v)`.",
        "Switch on the kind; for the five sized kinds return `rv.Len(), true`.",
        "Default to `0, false`.",
    ],
    walkthrough="`Length(nil)` gives an invalid Value whose kind is `Invalid`, which falls to the default. `Length([]int(nil))` gives a valid slice Value of length 0.",
    pitfalls=[
        "Calling `rv.Len()` before the switch, which panics on an int.",
        "Forgetting that `*[3]int` is a pointer, not an array — it has no length."
    ],
)

P(
    "middle",
    name="typesizes",
    title="Size And Alignment Of A Run-Time Type",
    sig="func Sizes(v any) (size, align uintptr, ok bool)",
    doc="""Sizes returns the size and alignment of v's dynamic type.

A nil interface has no type, so it reports false.

Examples:

	Sizes(int64(0)) => 8, 8, true""",
    imports=['"reflect"'],
    solution="""t := reflect.TypeOf(v)
if t == nil {
	return 0, 0, false
}
return t.Size(), uintptr(t.Align()), true""",
    tests="""
import (
	"testing"
	"unsafe"
)

type small struct {
	A byte
	B byte
}

type wide struct {
	A byte
	B int64
}

func TestSizes(t *testing.T) {
	cases := []struct {
		in    any
		size  uintptr
		align uintptr
	}{
		{int64(0), unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0))},
		{byte(0), 1, 1},
		{"", unsafe.Sizeof(""), unsafe.Alignof("")},
		{small{}, unsafe.Sizeof(small{}), unsafe.Alignof(small{})},
		{wide{}, unsafe.Sizeof(wide{}), unsafe.Alignof(wide{})},
	}
	for _, c := range cases {
		size, align, ok := Sizes(c.in)
		if !ok {
			t.Errorf("Sizes(%T) reported false", c.in)
			continue
		}
		if size != c.size {
			t.Errorf("Sizes(%T) size = %d, want %d", c.in, size, c.size)
		}
		if align != c.align {
			t.Errorf("Sizes(%T) align = %d, want %d", c.in, align, c.align)
		}
	}
}

func TestSizesNil(t *testing.T) {
	if _, _, ok := Sizes(nil); ok {
		t.Error("Sizes(nil) reported ok, want false")
	}
}

func TestSizesSliceHeader(t *testing.T) {
	size, _, ok := Sizes([]int{1, 2, 3})
	if !ok {
		t.Fatal("Sizes reported false")
	}
	if size != unsafe.Sizeof([]int(nil)) {
		t.Errorf("size = %d, want the header size %d: the elements do not count",
			size, unsafe.Sizeof([]int(nil)))
	}
}

func TestSizesAlignmentDividesSize(t *testing.T) {
	for _, in := range []any{int64(0), small{}, wide{}, "", []int(nil)} {
		size, align, _ := Sizes(in)
		if align == 0 || size%align != 0 {
			t.Errorf("%T: size %d is not a multiple of align %d", in, size, align)
		}
	}
}
""",
    context="A layout auditor wants a table of every registered type's footprint. `unsafe.Sizeof` needs the type at compile time, and the registry only has values.",
    task=[
        "Return the size and alignment of `v`'s dynamic type.",
        "Report false for a nil interface.",
        "Measure the type itself — a slice's size is its header, not its elements.",
    ],
    examples=[
        ("Sizes(int64(0))", "8, 8, true", None),
        ("Sizes([]int{1,2,3})", "24, 8, true", "The header, not the three ints."),
        ("Sizes(nil)", "0, 0, false", None),
    ],
    topics=[
        ("Type.Size and Type.Align", "The run-time equivalents of `unsafe.Sizeof` and `unsafe.Alignof`."),
        ("Align returns an int", "Unlike `Size`, so it needs a conversion."),
        ("Headers vs payload", "Reflection measures the same thing the compiler would."),
    ],
    hint="`Size` is already a uintptr. `Align` is not.",
    intuition="`unsafe.Sizeof` is a compile-time constant, so it cannot answer for a type you only meet at run time. `reflect.Type` carries the same numbers, computed by the same rules.",
    approach=[
        "Take `reflect.TypeOf(v)`; return false if it is nil.",
        "Return `t.Size()` and `uintptr(t.Align())`.",
    ],
    walkthrough="`Sizes(wide{})` reports 16 and 8: one byte, seven bytes of padding, then the int64. The alignment is the widest field's.",
    pitfalls=[
        "Expecting a slice's size to include its elements.",
        "`reflect.TypeOf(nil)` returns nil, and calling `Size` on it panics.",
    ],
)

P(
    "middle",
    name="bytescontains",
    title="Search Bytes Without Building A String",
    sig="func Contains(haystack []byte, needle string) bool",
    doc="""Contains reports whether needle appears in haystack.

An empty needle is always present. Neither operand may be converted:
the search runs over the bytes that are already there.

Examples:

	Contains([]byte("hello"), "ell") => true""",
    solution="""if len(needle) == 0 {
	return true
}
if len(needle) > len(haystack) {
	return false
}
for i := 0; i+len(needle) <= len(haystack); i++ {
	if haystack[i] != needle[0] {
		continue
	}
	match := true
	for j := 1; j < len(needle); j++ {
		if haystack[i+j] != needle[j] {
			match = false
			break
		}
	}
	if match {
		return true
	}
}
return false""",
    tests="""
import (
	"bytes"
	"testing"
)

var sink bool

func TestContains(t *testing.T) {
	cases := []struct {
		hay, needle string
		want        bool
	}{
		{"hello", "ell", true},
		{"hello", "hello", true},
		{"hello", "hellos", false},
		{"hello", "", true},
		{"", "", true},
		{"", "x", false},
		{"hello", "lo", true},
		{"hello", "he", true},
		{"aaa", "aab", false},
		{"abcabd", "abd", true},
	}
	for _, c := range cases {
		if got := Contains([]byte(c.hay), c.needle); got != c.want {
			t.Errorf("Contains(%q, %q) = %v, want %v", c.hay, c.needle, got, c.want)
		}
	}
}

func TestContainsMatchesStdlib(t *testing.T) {
	hay := []byte("the quick brown fox jumps over the lazy dog")
	for _, needle := range []string{"quick", "dog", "cat", "the", "g", "", "o d"} {
		want := bytes.Contains(hay, []byte(needle))
		if got := Contains(hay, needle); got != want {
			t.Errorf("Contains(_, %q) = %v, want %v", needle, got, want)
		}
	}
}

func TestContainsAllocatesNothing(t *testing.T) {
	hay := bytes.Repeat([]byte("payload-"), 256)
	if n := testing.AllocsPerRun(100, func() { sink = Contains(hay, "payload-payload") }); n != 0 {
		t.Errorf("Contains made %v allocations, want 0", n)
	}
}
""",
    context="A filter checks every incoming frame for a marker with `strings.Contains(string(frame), marker)`. Each check copies the whole frame to look for eight bytes.",
    task=[
        "Report whether `needle` appears in `haystack`.",
        "An empty needle is present; a needle longer than the haystack is not.",
        "Convert nothing — zero allocations.",
    ],
    examples=[
        ('Contains([]byte("hello"), "ell")', "true", None),
        ('Contains([]byte("hello"), "")', "true", "The empty string is everywhere."),
        ('Contains([]byte("aaa"), "aab")', "false", "A partial match must not stop the search."),
    ],
    topics=[
        ("string([]byte) copies", "The conversion allocates because strings are immutable."),
        ("Indexing both sides", "`haystack[i]` and `needle[j]` are both bytes already."),
        ("Restarting after a partial match", "The outer loop advances by one, not past the failed match."),
    ],
    hint="A first-byte check before the inner loop skips most positions cheaply.",
    intuition="Searching needs to compare bytes, and both operands already are bytes. The conversion exists only to satisfy a function signature — and it costs a full copy of the larger side.",
    approach=[
        "Handle the empty needle and the too-long needle.",
        "For each start position, skip quickly unless the first byte matches.",
        "Compare the rest; return true on a full match.",
    ],
    walkthrough='Searching "aab" in "aaa": positions 0 and 1 match the first byte and fail at the third; position 2 is too close to the end. The result is false.',
    pitfalls=[
        "Looping `i` to `len(haystack)`, which reads past the end during the inner comparison.",
        "`bytes.Contains` is the real answer and is also allocation-free — the point here is why the conversion was the cost.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="pooledreader",
    title="The Pooled Reader Still Reading The Last Request",
    mode="bug",
    sig="func FirstLine(r io.Reader) (string, error)",
    doc="""FirstLine returns the first '\\n'-terminated line from r, using a
pooled bufio.Reader.

A bufio.Reader taken from a pool is still attached to the previous
source, with the previous source's buffered bytes inside it.

Examples:

	FirstLine(strings.NewReader("a\\nb\\n")) => "a", nil""",
    imports=['"bufio"', '"io"', '"strings"', '"sync"'],
    extra="""var pool = sync.Pool{New: func() any { return bufio.NewReaderSize(nil, 64) }}""",
    buggy="""br := pool.Get().(*bufio.Reader)
defer pool.Put(br)
line, err := br.ReadString('\\n')
if err != nil && err != io.EOF {
	return "", err
}
return strings.TrimSuffix(line, "\\n"), nil""",
    solution="""br := pool.Get().(*bufio.Reader)
br.Reset(r)
defer pool.Put(br)
line, err := br.ReadString('\\n')
if err != nil && err != io.EOF {
	return "", err
}
return strings.TrimSuffix(line, "\\n"), nil""",
    tests="""
import (
	"errors"
	"strings"
	"testing"
)

func TestFirstLine(t *testing.T) {
	got, err := FirstLine(strings.NewReader("alpha\\nbeta\\n"))
	if err != nil || got != "alpha" {
		t.Errorf("FirstLine = %q, %v, want \\"alpha\\", nil", got, err)
	}
}

func TestFirstLineNoTrailingNewline(t *testing.T) {
	got, err := FirstLine(strings.NewReader("only"))
	if err != nil || got != "only" {
		t.Errorf("FirstLine = %q, %v, want \\"only\\", nil", got, err)
	}
}

func TestFirstLineEmpty(t *testing.T) {
	got, err := FirstLine(strings.NewReader(""))
	if err != nil || got != "" {
		t.Errorf("FirstLine = %q, %v, want empty, nil", got, err)
	}
}

func TestFirstLineAcrossManyRequests(t *testing.T) {
	for i := 0; i < 100; i++ {
		want := strings.Repeat("x", i%7+1)
		got, err := FirstLine(strings.NewReader(want + "\\ntail\\n"))
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("request %d: FirstLine = %q, want %q: the pooled reader kept the previous source", i, got, want)
		}
	}
}

type boom struct{}

func (boom) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestFirstLinePropagatesErrors(t *testing.T) {
	if _, err := FirstLine(boom{}); err == nil {
		t.Error("want the reader's error, got nil")
	}
}
""",
    context="A request handler pools its `bufio.Reader` to cut allocations. Under load, responses start containing fragments of other requests' bodies.",
    task=[
        "Return the first line of `r`, without the trailing newline.",
        "A final line with no newline still counts; an empty reader returns the empty string.",
        "Fix the single bug so a pooled reader reads from the right source.",
        "The reader must still go back to the pool.",
    ],
    examples=[
        ('FirstLine(strings.NewReader("alpha\\nbeta\\n"))', '"alpha", nil', None),
        ("100 sequential requests", "each gets its own first line", None),
        ("a failing reader", "the error", None),
    ],
    topics=[
        ("Pooled values carry state", "A `bufio.Reader` holds a source and a buffer of unread bytes."),
        ("Reader.Reset", "Rebinds the reader to a new source and discards the buffer."),
        ("Cross-request contamination", "Reuse bugs leak one caller's data into another's output."),
        ("io.EOF is not a failure", "A final unterminated line arrives with EOF."),
    ],
    hint="The pool hands you a reader. What is it currently reading from?",
    intuition="A pool recycles objects with their state intact. A buffered reader's state includes the source it wraps and whatever it has already read ahead — so a borrowed one keeps serving the previous request until it is rebound.",
    approach=[
        "Get the reader from the pool.",
        "`br.Reset(r)` to bind it to this call's source.",
        "Read the line, trim the newline, tolerate `io.EOF`.",
    ],
    walkthrough="The first call works because `New` builds a reader over a nil source, which fails and is then... not what happens: the reader is never bound, so every call reads from whatever the last `Reset` left — and with no `Reset` at all, from nil.",
    pitfalls=[
        "Resetting after reading — the read has already happened.",
        "Putting the reader back still bound to the caller's source, which keeps that source alive; `Reset(nil)` before `Put` avoids it.",
    ],
)

P(
    "senior",
    name="appendall",
    title="Concatenate Slices With One Allocation",
    mode="bug",
    sig="func AppendAll(parts [][]int) []int",
    doc="""AppendAll returns every part concatenated in order.

The final length is known before the first append, so the result should
be allocated once instead of growing through every doubling.

Examples:

	AppendAll([][]int{{1}, {2, 3}}) => []int{1, 2, 3}""",
    buggy="""var out []int
for _, p := range parts {
	out = append(out, p...)
}
return out""",
    solution="""n := 0
for _, p := range parts {
	n += len(p)
}
if n == 0 {
	return nil
}
out := make([]int, 0, n)
for _, p := range parts {
	out = append(out, p...)
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

var sink []int

func TestAppendAll(t *testing.T) {
	got := AppendAll([][]int{{1}, {2, 3}, {}})
	if !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("AppendAll = %v, want [1 2 3]", got)
	}
	if got := AppendAll(nil); len(got) != 0 {
		t.Errorf("AppendAll = %v, want empty", got)
	}
	if got := AppendAll([][]int{{}, {}}); len(got) != 0 {
		t.Errorf("AppendAll = %v, want empty", got)
	}
}

func TestAppendAllIsIndependent(t *testing.T) {
	a := []int{1, 2}
	got := AppendAll([][]int{a})
	got[0] = 99
	if a[0] != 1 {
		t.Error("the result shares storage with a part")
	}
}

func TestAppendAllAllocatesOnce(t *testing.T) {
	parts := make([][]int, 64)
	for i := range parts {
		parts[i] = []int{i, i, i, i}
	}
	n := testing.AllocsPerRun(50, func() { sink = AppendAll(parts) })
	if n > 1 {
		t.Errorf("AppendAll made %v allocations, want 1: size the result first", n)
	}
}

func TestAppendAllLarge(t *testing.T) {
	parts := make([][]int, 100)
	want := 0
	for i := range parts {
		parts[i] = make([]int, i)
		want += i
	}
	if got := AppendAll(parts); len(got) != want {
		t.Errorf("len = %d, want %d", len(got), want)
	}
}
""",
    context="A merge step concatenates a few dozen slices per batch. It starts from nil and grows, reallocating and copying everything it has so far at every doubling.",
    task=[
        "Concatenate the parts in order.",
        "Fix the single bug so the result costs one allocation.",
        "An empty input, or all-empty parts, returns an empty result.",
    ],
    examples=[
        ("AppendAll([][]int{{1},{2,3}})", "[1 2 3]", None),
        ("64 parts of 4 elements", "1 allocation, not several", None),
        ("AppendAll(nil)", "[]", None),
    ],
    topics=[
        ("Known output size", "The sum of the parts' lengths is the exact final length."),
        ("append's doubling", "Growing from nil to 256 elements reallocates about nine times."),
        ("make with length 0 and a capacity", "Keeps `append` semantics while reserving the space."),
    ],
    hint="One extra loop before the existing one.",
    intuition="`append` grows by guessing because it cannot see what is coming. When the caller can, one counting pass removes every reallocation and every intermediate copy.",
    approach=[
        "Sum the parts' lengths.",
        "Return nil when the total is zero.",
        "`make([]int, 0, n)` and append the parts into it.",
    ],
    walkthrough="64 parts of 4 elements is 256. Starting from nil, `append` reallocates at 1, 2, 4 … 256 and copies 255 elements along the way; the sized version allocates once and copies each element exactly once.",
    pitfalls=[
        "`make([]int, n)` instead of `make([]int, 0, n)`, which prepends n zeros.",
        "Returning `parts[0]` when there is only one part — that shares storage with the caller.",
    ],
)

P(
    "senior",
    name="mapfield",
    title="Create The Map A Struct Field Needs",
    sig="func PutTag(ptr any, key, val string) error",
    doc="""PutTag sets ptr's Tags map entry, creating the map when the field is
nil.

Writing to a nil map panics, and reflection will not create one for you.

Examples:

	PutTag(&doc{}, "a", "1") => nil, doc.Tags["a"] == "1" """,
    imports=['"errors"', '"reflect"'],
    extra="""// ErrTarget is returned when ptr has no settable Tags map field.
var ErrTarget = errors.New("target must be a pointer to a struct with a settable Tags map[string]string")""",
    solution="""rv := reflect.ValueOf(ptr)
if rv.Kind() != reflect.Pointer || rv.IsNil() {
	return ErrTarget
}
rv = rv.Elem()
if rv.Kind() != reflect.Struct {
	return ErrTarget
}
f := rv.FieldByName("Tags")
if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Map {
	return ErrTarget
}
mt := f.Type()
if mt.Key().Kind() != reflect.String || mt.Elem().Kind() != reflect.String {
	return ErrTarget
}
if f.IsNil() {
	f.Set(reflect.MakeMap(mt))
}
f.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
return nil""",
    tests="""
import (
	"errors"
	"testing"
)

type doc struct {
	Tags map[string]string
	Name string
}

func TestPutTagCreatesTheMap(t *testing.T) {
	d := &doc{}
	if err := PutTag(d, "a", "1"); err != nil {
		t.Fatal(err)
	}
	if d.Tags == nil {
		t.Fatal("the map was not created")
	}
	if d.Tags["a"] != "1" {
		t.Errorf("Tags = %v, want map[a:1]", d.Tags)
	}
}

func TestPutTagReusesTheMap(t *testing.T) {
	d := &doc{Tags: map[string]string{"keep": "yes"}}
	existing := d.Tags
	if err := PutTag(d, "a", "1"); err != nil {
		t.Fatal(err)
	}
	if d.Tags["keep"] != "yes" || d.Tags["a"] != "1" {
		t.Errorf("Tags = %v, want both entries", d.Tags)
	}
	d.Tags["direct"] = "x"
	if existing["direct"] != "x" {
		t.Error("the field was replaced with a new map")
	}
}

func TestPutTagOverwrites(t *testing.T) {
	d := &doc{}
	PutTag(d, "a", "1")
	if err := PutTag(d, "a", "2"); err != nil {
		t.Fatal(err)
	}
	if d.Tags["a"] != "2" {
		t.Errorf("Tags[a] = %q, want \\"2\\"", d.Tags["a"])
	}
}

func TestPutTagBadTargets(t *testing.T) {
	type noTags struct{ A int }
	type wrongKind struct{ Tags []string }
	type wrongTypes struct{ Tags map[string]int }
	type unexported struct{ tags map[string]string }

	cases := []any{
		doc{}, nil, (*doc)(nil), new(int),
		&noTags{}, &wrongKind{}, &wrongTypes{}, &unexported{},
	}
	for _, c := range cases {
		if err := PutTag(c, "a", "1"); !errors.Is(err, ErrTarget) {
			t.Errorf("PutTag(%#v) = %v, want ErrTarget", c, err)
		}
	}
}
""",
    context="A metadata helper writes into a struct's `Tags` map. Half the structs it is given have never had a tag set, and the helper panics on the first one.",
    task=[
        "Set `Tags[key] = val` on the struct `ptr` points at.",
        "Create the map when the field is nil; reuse it when it is not.",
        "Return `ErrTarget` for a bad target: not a pointer, not a struct, no `Tags` field, unexported, wrong kind, or wrong key/value types.",
    ],
    examples=[
        ('PutTag(&doc{}, "a", "1")', "nil, the map is created", None),
        ("a doc that already has tags", "the existing map is kept", None),
        ("&struct{Tags []string}{}", "ErrTarget", None),
    ],
    topics=[
        ("Nil maps are read-only", "Writing to one panics, in reflection as in ordinary code."),
        ("reflect.MakeMap", "Builds a map of a type known only at run time."),
        ("SetMapIndex", "The reflective map write; the key and value must be assignable to the map's types."),
        ("Type checks before writes", "Every reflective write panics on a mismatch, so validate first."),
    ],
    hint="Check `f.IsNil()` before writing, and `MakeMap` needs the field's own type.",
    intuition="Reflection mirrors the language exactly: a nil map cannot be written to, and the fix is the same as in ordinary code — make one. The only difference is that the map's type comes from the field rather than from the source.",
    approach=[
        "Validate the pointer, the struct, and the `Tags` field's settability and kind.",
        "Check the map's key and element kinds.",
        "If the field is nil, `Set` it to `reflect.MakeMap(mt)`.",
        "`SetMapIndex` with the key and value.",
    ],
    walkthrough="For a fresh `doc`, `f.IsNil()` is true, so a new `map[string]string` is created and stored in the field; then the entry is written. For a doc that already has tags, the existing map is written through.",
    pitfalls=[
        "Skipping the key/value type check, which panics inside `SetMapIndex`.",
        "Building the map with `reflect.MakeMap(reflect.TypeOf(map[string]string{}))` — it works here and breaks for any other map type.",
    ],
)

P(
    "senior",
    name="slicelen",
    title="The Length Argument Is In Elements",
    mode="bug",
    sig="func Words(b []byte) ([]uint32, bool)",
    doc="""Words returns a []uint32 view over b's bytes.

unsafe.Slice takes a count of elements, not of bytes: passing the byte
length produces a view four times too long, running off the end of the
buffer.

Examples:

	Words(eightBytes) => a 2-element view, true""",
    imports=['"unsafe"'],
    buggy="""if len(b) == 0 || len(b)%4 != 0 {
	return nil, false
}
p := unsafe.Pointer(unsafe.SliceData(b))
if uintptr(p)&3 != 0 {
	return nil, false
}
return unsafe.Slice((*uint32)(p), len(b)), true""",
    solution="""if len(b) == 0 || len(b)%4 != 0 {
	return nil, false
}
p := unsafe.Pointer(unsafe.SliceData(b))
if uintptr(p)&3 != 0 {
	return nil, false
}
return unsafe.Slice((*uint32)(p), len(b)/4), true""",
    tests="""
import (
	"testing"
	"unsafe"
)

func alignedBytes(n int) []byte {
	u := make([]uint32, (n+3)/4)
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(u))), n)
}

func TestWordsLength(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Words(b)
	if !ok {
		t.Fatal("Words reported false for an aligned 8-byte buffer")
	}
	if len(v) != 2 {
		t.Fatalf("len = %d, want 2: the count is in elements, not bytes", len(v))
	}
	if cap(v) != 2 {
		t.Errorf("cap = %d, want 2", cap(v))
	}
}

func TestWordsCoversExactlyTheBuffer(t *testing.T) {
	b := alignedBytes(16)
	v, ok := Words(b)
	if !ok {
		t.Fatal("Words reported false")
	}
	if len(v)*4 != len(b) {
		t.Errorf("the view covers %d bytes, want %d", len(v)*4, len(b))
	}
}

func TestWordsSharesStorage(t *testing.T) {
	b := alignedBytes(8)
	v, ok := Words(b)
	if !ok {
		t.Fatal("Words reported false")
	}
	v[0] = 0xffffffff
	for i := 0; i < 4; i++ {
		if b[i] != 0xff {
			t.Fatalf("b[%d] = %#x, want 0xff: the view does not share the bytes", i, b[i])
		}
	}
}

func TestWordsRejectsBadShapes(t *testing.T) {
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
		if _, ok := Words(c.in); ok {
			t.Errorf("%s: reported ok, want false", c.name)
		}
	}
}
""",
    context="A frame decoder reinterprets a received buffer as words. It reads plausible garbage past the end of every buffer, and the corruption is blamed on the network for a week.",
    task=[
        "Return a `[]uint32` view sharing `b`'s storage.",
        "Report false for an empty buffer, a length that is not a multiple of four, or a misaligned start.",
        "Fix the single bug: the view must cover exactly `b` and no more.",
    ],
    examples=[
        ("Words(alignedBytes(8))", "a 2-element view, true", None),
        ("len(view) * 4", "len(b)", "The view covers exactly the buffer."),
        ("Words(b[:6])", "nil, false", None),
    ],
    topics=[
        ("unsafe.Slice's length is in elements", "It is multiplied by the element size internally."),
        ("Out-of-bounds without a panic", "The runtime cannot check a length you invented."),
        ("Silent corruption", "The extra elements read whatever follows the buffer in memory."),
    ],
    hint="Everything about the pointer is right. Count what the second argument is counting.",
    intuition="`unsafe.Slice(p, n)` builds a slice of `n` elements of `*p`'s type. Handing it a byte count for a four-byte element type asks for four times the buffer — and nothing in the runtime will object.",
    approach=[
        "Keep the validation as it is.",
        "Pass `len(b)/4` as the element count.",
    ],
    walkthrough="For an 8-byte buffer the correct view is 2 elements. The buggy version asks for 8, so `v[2]` through `v[7]` read 24 bytes that belong to something else.",
    pitfalls=[
        "Testing only that the values look right — the first two do.",
        "Fixing it with `[:len(b)/4]` after the fact, which narrows the view but was already constructed out of bounds.",
    ],
)

P(
    "senior",
    name="interncache",
    title="Store One Copy Of Each Repeated String",
    sig="func (p *Pool) Intern(b []byte) string",
    doc="""Intern returns a string with b's contents, reusing a previously stored
one when the same bytes have been seen before.

Repeated values then share one allocation instead of one each, and a
repeat lookup must not allocate at all.

Examples:

	p.Intern([]byte("a")) twice => the same string, one allocation""",
    imports=['"unsafe"'],
    extra="""// Pool holds one canonical string per distinct byte sequence.
type Pool struct {
	m map[string]string
}

// Len reports how many distinct strings the pool holds.
func (p *Pool) Len() int { return len(p.m) }""",
    solution="""if len(b) == 0 {
	return ""
}
view := unsafe.String(unsafe.SliceData(b), len(b))
if s, ok := p.m[view]; ok {
	return s
}
owned := string(b)
if p.m == nil {
	p.m = make(map[string]string)
}
p.m[owned] = owned
return owned""",
    tests="""
import (
	"testing"
	"unsafe"
)

var sink string

func TestInternReturnsTheContents(t *testing.T) {
	var p Pool
	if got := p.Intern([]byte("hello")); got != "hello" {
		t.Errorf("Intern = %q, want \\"hello\\"", got)
	}
	if got := p.Intern(nil); got != "" {
		t.Errorf("Intern(nil) = %q, want empty", got)
	}
	if p.Len() != 1 {
		t.Errorf("Len = %d, want 1: the empty string is not stored", p.Len())
	}
}

func TestInternSharesOneCopy(t *testing.T) {
	var p Pool
	a := p.Intern([]byte("repeated"))
	b := p.Intern([]byte("repeated"))
	if unsafe.StringData(a) != unsafe.StringData(b) {
		t.Error("the two results are separate allocations; they must share one")
	}
	if p.Len() != 1 {
		t.Errorf("Len = %d, want 1", p.Len())
	}
}

func TestInternDistinctValues(t *testing.T) {
	var p Pool
	for i := 0; i < 26; i++ {
		p.Intern([]byte{byte('a' + i)})
	}
	if p.Len() != 26 {
		t.Errorf("Len = %d, want 26", p.Len())
	}
}

func TestInternSurvivesBufferReuse(t *testing.T) {
	var p Pool
	buf := make([]byte, 4)
	got := make([]string, 0, 26)
	for i := 0; i < 26; i++ {
		for j := range buf {
			buf[j] = byte('a' + i)
		}
		got = append(got, p.Intern(buf))
	}
	for i, s := range got {
		want := string([]byte{byte('a' + i), byte('a' + i), byte('a' + i), byte('a' + i)})
		if s != want {
			t.Fatalf("result %d = %q, want %q: the stored string viewed the reused buffer", i, s, want)
		}
	}
	if p.Len() != 26 {
		t.Errorf("Len = %d, want 26", p.Len())
	}
}

func TestInternRepeatDoesNotAllocate(t *testing.T) {
	var p Pool
	key := []byte("a-repeated-token")
	p.Intern(key)
	if n := testing.AllocsPerRun(200, func() { sink = p.Intern(key) }); n != 0 {
		t.Errorf("a repeat Intern made %v allocations, want 0: borrow the bytes for the lookup", n)
	}
}
""",
    context="A parser produces millions of strings, most of them repeats of a few hundred distinct values. Each repeat allocates its own copy, and the heap is full of identical short strings.",
    task=[
        "Return a string with `b`'s contents, reusing the stored one when the bytes have been seen.",
        "A repeat lookup must allocate nothing.",
        "A stored string must own its bytes — the caller reuses its buffer.",
        "The empty input returns `\"\"` without storing anything.",
    ],
    examples=[
        ('p.Intern([]byte("repeated")) twice', "the same string, one allocation", None),
        ("200 repeat lookups", "0 allocations", None),
        ("26 batches through one reused buffer", "26 distinct stored strings", None),
    ],
    topics=[
        ("Borrow for the lookup, own for the store", "The read path may alias; the write path may not."),
        ("Map keys must be immutable", "A stored key that changes strands its entry."),
        ("Interning", "One canonical copy per distinct value, shared by every holder."),
        ("Key and value are the same string", "Storing `owned` under itself keeps one allocation, not two."),
    ],
    hint="Two paths: the hit borrows, the miss copies. Only one of them may allocate.",
    intuition="Interning trades a map lookup for an allocation, which pays off when repeats dominate. The subtlety is that the lookup can be free — a string view built to hash and compare never escapes — while the stored copy cannot be.",
    approach=[
        "Return `\"\"` for an empty input.",
        "Build a borrowed string view and look it up; return the stored string on a hit.",
        "On a miss, copy with `string(b)`, store it as both key and value, and return it.",
    ],
    walkthrough="The first `Intern` of a token copies it once and stores it. Every later call with the same bytes hashes the borrowed view, finds the entry, and returns the stored string — no allocation at all.",
    pitfalls=[
        "Storing the borrowed view as the key, which the reuse test exists to catch.",
        "Storing the key and the value as two separate copies.",
        "Letting the pool grow without bound; a real interner needs an eviction rule.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="ringpool",
    title="Hand Out Buffers Without A Lock On The Fast Path",
    sig="func (p *BufPool) Get() []byte",
    doc="""Get returns a buffer from the ring, or a fresh one when the ring is
empty.

The ring is a fixed array of slots claimed with an atomic index, so
concurrent callers never block each other and never receive the same
buffer twice.

Examples:

	p := NewBufPool(4, 64); p.Get() => a 64-byte buffer""",
    imports=['"sync/atomic"'],
    extra="""// BufPool hands out fixed-size buffers from a bounded ring.
type BufPool struct {
	size  int
	next  atomic.Int64
	slots []atomic.Pointer[[]byte]
}

// NewBufPool returns a pool of n slots holding size-byte buffers.
func NewBufPool(n, size int) *BufPool {
	if n < 1 {
		n = 1
	}
	if size < 1 {
		size = 1
	}
	return &BufPool{size: size, slots: make([]atomic.Pointer[[]byte], n)}
}

// Put returns a buffer to the ring, dropping it if the ring is full or the
// buffer is the wrong size.
func (p *BufPool) Put(b []byte) {
	if cap(b) != p.size {
		return
	}
	b = b[:0]
	i := int(p.next.Add(1)-1) % len(p.slots)
	if i < 0 {
		i = -i
	}
	p.slots[i].CompareAndSwap(nil, &b)
}""",
    solution="""for i := range p.slots {
	if b := p.slots[i].Swap(nil); b != nil {
		return (*b)[:0]
	}
}
return make([]byte, 0, p.size)""",
    tests="""
import (
	"sync"
	"testing"
)

func TestGetReturnsASizedBuffer(t *testing.T) {
	p := NewBufPool(4, 64)
	b := p.Get()
	if len(b) != 0 {
		t.Errorf("len = %d, want 0", len(b))
	}
	if cap(b) != 64 {
		t.Errorf("cap = %d, want 64", cap(b))
	}
}

func TestGetRecyclesAPutBuffer(t *testing.T) {
	p := NewBufPool(4, 64)
	b := p.Get()
	b = append(b, 'x')
	p.Put(b)
	got := p.Get()
	if len(got) != 0 {
		t.Errorf("len = %d, want 0: a recycled buffer must be empty", len(got))
	}
	if cap(got) != 64 {
		t.Errorf("cap = %d, want 64", cap(got))
	}
}

func TestGetNeverHandsOutTheSameBufferTwice(t *testing.T) {
	p := NewBufPool(4, 64)
	for i := 0; i < 4; i++ {
		b := make([]byte, 0, 64)
		p.Put(b)
	}
	seen := map[*byte]bool{}
	for i := 0; i < 4; i++ {
		b := p.Get()
		b = append(b, 0)
		if seen[&b[0]] {
			t.Fatal("the same buffer was handed out twice")
		}
		seen[&b[0]] = true
	}
}

func TestGetFallsBackWhenEmpty(t *testing.T) {
	p := NewBufPool(2, 32)
	for i := 0; i < 10; i++ {
		b := p.Get()
		if cap(b) != 32 {
			t.Fatalf("cap = %d, want 32", cap(b))
		}
	}
}

func TestPutRejectsWrongSizes(t *testing.T) {
	p := NewBufPool(2, 32)
	p.Put(make([]byte, 0, 4096))
	b := p.Get()
	if cap(b) != 32 {
		t.Errorf("cap = %d, want 32: an oversized buffer must not enter the ring", cap(b))
	}
}

func TestConcurrentGetAndPut(t *testing.T) {
	p := NewBufPool(8, 128)
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				b := p.Get()
				if cap(b) != 128 || len(b) != 0 {
					panic("bad buffer from the pool")
				}
				b = append(b, byte(w))
				if b[0] != byte(w) {
					panic("buffer shared between goroutines")
				}
				p.Put(b)
			}
		}(w)
	}
	wg.Wait()
}
""",
    context="A pool behind a mutex is fine at four cores and is the top contended lock at sixty-four. The buffers are all the same size, and there are only ever a handful in flight.",
    task=[
        "Return a buffer from the ring, emptied and ready to write.",
        "Fall back to a fresh `p.size` buffer when every slot is empty.",
        "A slot's buffer must be handed to exactly one caller — never two.",
        "Correct under concurrent `Get` and `Put`.",
    ],
    examples=[
        ("NewBufPool(4, 64).Get()", "len 0, cap 64", None),
        ("Put then Get", "the same storage, emptied", None),
        ("4 slots filled, 4 Gets", "four distinct buffers", None),
    ],
    topics=[
        ("Atomic claim by swap", "`Swap(nil)` both reads the slot and empties it, so only one caller wins."),
        ("Read-modify-write must be one operation", "A load followed by a store lets two callers take one buffer."),
        ("Bounded by construction", "A fixed ring cannot grow, whatever the load."),
        ("Fallback over blocking", "An empty ring allocates rather than waiting."),
    ],
    hint="Taking a slot has to be a single atomic operation. Which one both reads and clears?",
    intuition="A lock-free hand-off works only if claiming is indivisible. `Swap` returns the old value and installs nil in one step, so two goroutines racing on a slot cannot both come away with the buffer.",
    approach=[
        "Walk the slots, `Swap(nil)` on each.",
        "Return the first non-nil buffer, resliced to length 0.",
        "If every slot was empty, allocate a fresh one.",
    ],
    walkthrough="With 16 goroutines and 8 slots, the fast path is one atomic swap; when the ring runs dry a caller allocates rather than blocking, and `Put` refills the ring as buffers come back.",
    pitfalls=[
        "`Load` then `Store(nil)` — two goroutines can both see the same buffer.",
        "Returning the buffer without `[:0]`, so the caller appends onto the last user's bytes.",
        "Assuming the ring is a cache; a buffer that never comes back is simply collected.",
    ],
)

P(
    "staff",
    name="batchflush",
    title="Flush At The Threshold, Not At The End",
    sig="func (b *Batcher) Add(v int) error",
    doc="""Add appends v to the pending batch and flushes when the batch reaches
its limit.

The pending slice must never grow past the limit: an unbounded batcher
turns a slow consumer into an out-of-memory kill.

Examples:

	b := NewBatcher(2, sink); b.Add(1); b.Add(2) => sink received [1 2]""",
    imports=['"sync"'],
    extra="""// Batcher accumulates values and flushes them in fixed-size batches.
type Batcher struct {
	mu      sync.Mutex
	limit   int
	pending []int
	flush   func([]int) error
}

// NewBatcher returns a batcher that calls flush with each full batch.
func NewBatcher(limit int, flush func([]int) error) *Batcher {
	if limit < 1 {
		limit = 1
	}
	return &Batcher{limit: limit, pending: make([]int, 0, limit), flush: flush}
}

// Pending reports how many values are waiting.
func (b *Batcher) Pending() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.pending)
}

// Close flushes whatever is left.
func (b *Batcher) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.pending) == 0 {
		return nil
	}
	batch := make([]int, len(b.pending))
	copy(batch, b.pending)
	b.pending = b.pending[:0]
	return b.flush(batch)
}""",
    solution="""b.mu.Lock()
defer b.mu.Unlock()

b.pending = append(b.pending, v)
if len(b.pending) < b.limit {
	return nil
}
batch := make([]int, len(b.pending))
copy(batch, b.pending)
b.pending = b.pending[:0]
return b.flush(batch)""",
    tests="""
import (
	"errors"
	"sync"
	"testing"
)

func TestAddFlushesAtTheLimit(t *testing.T) {
	var got [][]int
	b := NewBatcher(2, func(batch []int) error {
		got = append(got, batch)
		return nil
	})
	for i := 1; i <= 4; i++ {
		if err := b.Add(i); err != nil {
			t.Fatal(err)
		}
	}
	if len(got) != 2 {
		t.Fatalf("flushed %d batches, want 2", len(got))
	}
	if got[0][0] != 1 || got[0][1] != 2 || got[1][0] != 3 || got[1][1] != 4 {
		t.Errorf("batches = %v, want [[1 2] [3 4]]", got)
	}
	if b.Pending() != 0 {
		t.Errorf("Pending = %d, want 0", b.Pending())
	}
}

func TestPendingNeverExceedsTheLimit(t *testing.T) {
	b := NewBatcher(4, func([]int) error { return nil })
	for i := 0; i < 1000; i++ {
		if err := b.Add(i); err != nil {
			t.Fatal(err)
		}
		if p := b.Pending(); p >= 4 {
			t.Fatalf("Pending = %d after %d adds, want under 4", p, i+1)
		}
	}
}

func TestBatchesAreIndependentOfThePending(t *testing.T) {
	var kept []int
	b := NewBatcher(2, func(batch []int) error {
		kept = batch
		return nil
	})
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	if kept == nil {
		t.Fatal("no batch was kept")
	}
	if kept[0] != 3 || kept[1] != 4 {
		t.Errorf("the kept batch = %v, want [3 4]: it must not alias the pending buffer", kept)
	}
}

func TestAddPropagatesFlushErrors(t *testing.T) {
	boom := errors.New("boom")
	b := NewBatcher(1, func([]int) error { return boom })
	if err := b.Add(1); !errors.Is(err, boom) {
		t.Errorf("Add = %v, want boom", err)
	}
}

func TestCloseFlushesTheRemainder(t *testing.T) {
	var got [][]int
	b := NewBatcher(3, func(batch []int) error {
		got = append(got, batch)
		return nil
	})
	b.Add(1)
	b.Add(2)
	if err := b.Close(); err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || len(got[0]) != 2 {
		t.Errorf("batches = %v, want one batch of 2", got)
	}
}

func TestConcurrentAdds(t *testing.T) {
	var mu sync.Mutex
	total := 0
	b := NewBatcher(8, func(batch []int) error {
		mu.Lock()
		total += len(batch)
		mu.Unlock()
		return nil
	})
	var wg sync.WaitGroup
	const workers = 8
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				b.Add(i)
			}
		}()
	}
	wg.Wait()
	b.Close()
	mu.Lock()
	defer mu.Unlock()
	if total != workers*100 {
		t.Errorf("flushed %d values, want %d", total, workers*100)
	}
}
""",
    context="A metrics batcher accumulates until a timer fires. When the downstream stalls, the pending slice grows to gigabytes and the process is killed before the timer ever helps.",
    task=[
        "Append `v` to the pending batch, flushing when it reaches the limit.",
        "The pending slice must never hold `limit` or more values after `Add` returns.",
        "The flushed batch must not alias the pending buffer — the callee may keep it.",
        "Propagate the flush error; safe for concurrent use.",
    ],
    examples=[
        ("NewBatcher(2, f); Add(1); Add(2)", "f received [1 2], Pending is 0", None),
        ("1000 adds with limit 4", "Pending stays under 4 throughout", None),
        ("a flush that fails", "Add returns the error", None),
    ],
    topics=[
        ("Bounded accumulation", "The threshold, not a timer, is what caps the memory."),
        ("Copy before handing over", "The pending buffer is reused; the batch is not."),
        ("Reset with [:0]", "The pending buffer keeps its capacity across batches."),
        ("Lock discipline", "One lock covers append, threshold check and reset."),
    ],
    hint="Append, compare, and then three things happen in order.",
    intuition="A size threshold is the only bound that holds when the consumer is slow. Flushing on the way in means the pending buffer's capacity is the batcher's entire memory footprint.",
    approach=[
        "Under the lock, append `v`.",
        "Return nil when the batch is not yet full.",
        "Copy the pending values into a fresh batch, reset `pending` to `[:0]`, and flush the copy.",
    ],
    walkthrough="With limit 2, the second `Add` copies [1 2] into a new slice, empties the pending buffer, and calls flush. The pending buffer keeps its capacity of 2 and is reused forever.",
    pitfalls=[
        "Passing `b.pending` straight to `flush` — the next batch overwrites what the callee kept.",
        "`b.pending = nil` instead of `[:0]`, which throws away the capacity every batch.",
        "Flushing outside the lock without taking the batch out first, which lets a concurrent `Add` join the batch mid-flush.",
    ],
)

P(
    "staff",
    name="fastdispatch",
    title="Reflect Only When You Have To",
    sig="func Render(dst []byte, v any) []byte",
    doc="""Render appends v's text form to dst.

The common types are handled by a type switch, which costs nothing;
everything else falls back to reflection. The fast path must not
allocate.

Examples:

	Render(nil, 42) => []byte("42")""",
    imports=['"reflect"', '"strconv"'],
    solution="""switch x := v.(type) {
case nil:
	return append(dst, "<nil>"...)
case string:
	return append(dst, x...)
case int:
	return strconv.AppendInt(dst, int64(x), 10)
case int64:
	return strconv.AppendInt(dst, x, 10)
case bool:
	return strconv.AppendBool(dst, x)
case []byte:
	return append(dst, x...)
}
rv := reflect.ValueOf(v)
switch rv.Kind() {
case reflect.String:
	return append(dst, rv.String()...)
case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	return strconv.AppendInt(dst, rv.Int(), 10)
case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	return strconv.AppendUint(dst, rv.Uint(), 10)
case reflect.Bool:
	return strconv.AppendBool(dst, rv.Bool())
default:
	return append(dst, '?')
}""",
    tests="""
import (
	"bytes"
	"testing"
)

type myInt int
type myString string

var sink []byte

func TestRenderFastPath(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{42, "42"},
		{int64(-7), "-7"},
		{"text", "text"},
		{true, "true"},
		{false, "false"},
		{[]byte("bytes"), "bytes"},
		{nil, "<nil>"},
	}
	for _, c := range cases {
		if got := Render(nil, c.in); !bytes.Equal(got, []byte(c.want)) {
			t.Errorf("Render(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRenderFallback(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{myInt(5), "5"},
		{myString("named"), "named"},
		{uint8(200), "200"},
		{int32(-3), "-3"},
		{uint64(1 << 40), "1099511627776"},
	}
	for _, c := range cases {
		if got := Render(nil, c.in); !bytes.Equal(got, []byte(c.want)) {
			t.Errorf("Render(%#v) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestRenderUnsupported(t *testing.T) {
	if got := Render(nil, struct{ A int }{}); !bytes.Equal(got, []byte("?")) {
		t.Errorf("Render = %q, want \\"?\\"", got)
	}
	if got := Render(nil, 1.5); !bytes.Equal(got, []byte("?")) {
		t.Errorf("Render = %q, want \\"?\\"", got)
	}
}

func TestRenderAppends(t *testing.T) {
	got := Render([]byte("pre:"), 9)
	if !bytes.Equal(got, []byte("pre:9")) {
		t.Errorf("Render = %q, want \\"pre:9\\"", got)
	}
}

func TestRenderFastPathAllocatesNothing(t *testing.T) {
	dst := make([]byte, 0, 64)
	for _, v := range []any{42, "text", true, int64(9)} {
		v := v
		if n := testing.AllocsPerRun(200, func() { sink = Render(dst[:0], v) }); n != 0 {
			t.Errorf("Render(%#v) made %v allocations, want 0", v, n)
		}
	}
}

func BenchmarkRenderInt(b *testing.B) {
	dst := make([]byte, 0, 64)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sink = Render(dst[:0], 42)
	}
}
""",
    context="A structured logger renders every field through reflection. Ninety-nine percent of the fields are ints and strings, and each one pays for a run-time type walk.",
    task=[
        "Append `v`'s text form to `dst`.",
        "Handle nil, `string`, `int`, `int64`, `bool` and `[]byte` with a type switch — no allocations.",
        "Fall back to reflection for other integer, string and bool kinds, including named types.",
        "Anything else renders as `?`.",
    ],
    examples=[
        ("Render(nil, 42)", '"42"', "The fast path."),
        ("Render(nil, myInt(5))", '"5"', "A named type reaches the fallback."),
        ("Render(nil, 1.5)", '"?"', None),
    ],
    topics=[
        ("Type switch vs reflection", "The switch compares type words; reflection walks the type."),
        ("Named types miss the switch", "`myInt` is not `int` to a type assertion, but its kind is int."),
        ("strconv.Append*", "Renders into the destination with no intermediate string."),
        ("Fast path, slow path", "Optimise for the distribution you actually have."),
    ],
    hint="Two switches: one on the concrete type, one on the kind.",
    intuition="A type switch is a handful of comparisons the compiler lays out; reflection is a run-time inspection. Handling the common types first means the general machinery only runs for the cases that need it.",
    approach=[
        "Type-switch the concrete types, appending directly.",
        "Otherwise take `reflect.ValueOf(v)` and switch on the kind.",
        "Use the typed accessors — `String`, `Int`, `Uint`, `Bool`.",
        "Default to `?`.",
    ],
    walkthrough="`Render(dst, 42)` matches `case int` and appends two digits. `Render(dst, myInt(5))` falls through the switch, reaches the fallback, and `rv.Int()` reads it without boxing.",
    pitfalls=[
        "`fmt.Append` in the fallback, which is correct and allocates for the boxing.",
        "Forgetting that `case nil` in a type switch matches a nil interface, not a typed nil.",
        "Using `rv.Interface()` in the fallback, which allocates and undoes the point.",
    ],
)

P(
    "staff",
    name="stripedbuffers",
    title="A Buffer Per Shard, Padded Apart",
    sig="func (s *Striped) With(id int, fn func(buf []byte) []byte) int",
    doc="""With runs fn on the shard's scratch buffer and returns the number of
bytes fn left in it.

Each shard has its own buffer and its own lock, padded so neighbouring
shards do not share a cache line.

Examples:

	s.With(0, func(b []byte) []byte { return append(b, 'x') }) => 1""",
    imports=['"sync"', '"unsafe"'],
    extra="""// lineSize is the coherence granule the shards must not share.
const lineSize = 64

// stripe is one shard: a lock, a scratch buffer, and padding to a line.
type stripe struct {
	mu  sync.Mutex
	buf []byte
	_   [lineSize - unsafe.Sizeof(sync.Mutex{}) - unsafe.Sizeof([]byte(nil))]byte
}

// Striped hands out per-shard scratch buffers.
type Striped struct {
	stripes []stripe
}

// NewStriped returns a Striped with n shards, each holding a size-byte buffer.
func NewStriped(n, size int) *Striped {
	if n < 1 {
		n = 1
	}
	s := &Striped{stripes: make([]stripe, n)}
	for i := range s.stripes {
		s.stripes[i].buf = make([]byte, 0, size)
	}
	return s
}""",
    solution="""i := id % len(s.stripes)
if i < 0 {
	i = -i
}
st := &s.stripes[i]
st.mu.Lock()
defer st.mu.Unlock()
st.buf = fn(st.buf[:0])
return len(st.buf)""",
    tests="""
import (
	"sync"
	"testing"
	"unsafe"
)

func TestWithRunsOnTheBuffer(t *testing.T) {
	s := NewStriped(4, 64)
	got := s.With(0, func(b []byte) []byte { return append(b, 'a', 'b') })
	if got != 2 {
		t.Errorf("With = %d, want 2", got)
	}
}

func TestBufferIsResetEachTime(t *testing.T) {
	s := NewStriped(2, 64)
	for i := 0; i < 10; i++ {
		got := s.With(0, func(b []byte) []byte { return append(b, 'x') })
		if got != 1 {
			t.Fatalf("call %d: With = %d, want 1: the buffer was not reset", i, got)
		}
	}
}

func TestShardsAreIndependent(t *testing.T) {
	s := NewStriped(4, 64)
	s.With(0, func(b []byte) []byte { return append(b, 'a', 'a', 'a') })
	got := s.With(1, func(b []byte) []byte {
		if len(b) != 0 {
			t.Errorf("shard 1 saw %d bytes from another shard", len(b))
		}
		return append(b, 'b')
	})
	if got != 1 {
		t.Errorf("With = %d, want 1", got)
	}
}

func TestNegativeAndLargeIDs(t *testing.T) {
	s := NewStriped(4, 64)
	for _, id := range []int{-1, -7, 0, 3, 4, 1 << 20} {
		got := s.With(id, func(b []byte) []byte { return append(b, 'z') })
		if got != 1 {
			t.Fatalf("id %d: With = %d, want 1", id, got)
		}
	}
}

func TestStripesDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(stripe{}); got != lineSize {
		t.Errorf("sizeof(stripe) = %d, want %d", got, lineSize)
	}
	s := make([]stripe, 2)
	a := uintptr(unsafe.Pointer(&s[0]))
	b := uintptr(unsafe.Pointer(&s[1]))
	if b-a != lineSize {
		t.Errorf("stride = %d, want %d", b-a, lineSize)
	}
}

func TestConcurrentShards(t *testing.T) {
	s := NewStriped(8, 128)
	var wg sync.WaitGroup
	const workers = 16
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(w int) {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				n := s.With(w, func(b []byte) []byte {
					for j := 0; j < 4; j++ {
						b = append(b, byte(w))
					}
					return b
				})
				if n != 4 {
					panic("shard buffer shared between goroutines")
				}
			}
		}(w)
	}
	wg.Wait()
}
""",
    context="Every worker in a pool allocates a scratch buffer per call. Moving to one shared buffer created a race; moving to a `sync.Pool` helped, and the profile still shows cache-line ping-pong between the pool's per-shard state.",
    task=[
        "Run `fn` on the shard's scratch buffer, reset to length 0 first.",
        "Store the returned slice back and report its length.",
        "Route `id` to a shard, including negative and out-of-range ids.",
        "Hold only that shard's lock; safe for concurrent use.",
    ],
    examples=[
        ("s.With(0, appendOneByte)", "1", None),
        ("ten calls on one shard", "1 every time", "The buffer is reset each call."),
        ("sizeof(stripe)", "64", "One shard per cache line."),
    ],
    topics=[
        ("Lock striping", "Independent shards mean independent locks."),
        ("Padding computed from the fields", "`lineSize - Sizeof(mutex) - Sizeof(slice)` adapts if the struct changes."),
        ("Reset before, store after", "`fn` may append past the capacity, so its result is the new buffer."),
        ("Modulo with negative ids", "Go's `%` keeps the sign of the dividend."),
    ],
    hint="Route, lock, reset, call, store, report.",
    intuition="Per-shard state removes contention in software; padding removes it in hardware. Both are needed — eight mutexes on one cache line contend almost as badly as one mutex.",
    approach=[
        "Reduce `id` modulo the shard count, correcting a negative result.",
        "Lock the shard, call `fn(st.buf[:0])`, store the result back.",
        "Return its length.",
    ],
    walkthrough="Sixteen workers over eight shards contend in pairs rather than sixteen ways, and each shard's mutex and buffer header sit alone on their line.",
    pitfalls=[
        "Discarding `fn`'s return value, so a buffer that grew is lost.",
        "`id % len` alone, which indexes negatively and panics.",
        "Padding the slice of stripes instead of the stripe itself.",
    ],
)

P(
    "staff",
    name="pipeline",
    title="A Stage That Stops When Told",
    sig="func Stage(done <-chan struct{}, in <-chan int, workers int) <-chan int",
    doc="""Stage returns a channel carrying each input doubled, computed by
workers goroutines, and closed once the input drains or done is closed.

Every goroutine must exit on done, whether it is blocked receiving or
blocked sending.

Examples:

	Stage(done, in, 4) => a channel of doubled values""",
    imports=['"sync"'],
    solution="""if workers < 1 {
	workers = 1
}
out := make(chan int)
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	go func() {
		defer wg.Done()
		for {
			var v int
			var ok bool
			select {
			case v, ok = <-in:
				if !ok {
					return
				}
			case <-done:
				return
			}
			select {
			case out <- v * 2:
			case <-done:
				return
			}
		}
	}()
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

func feed(vals ...int) <-chan int {
	ch := make(chan int, len(vals))
	for _, v := range vals {
		ch <- v
	}
	close(ch)
	return ch
}

func TestStageDoublesEverything(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(1, 2, 3), 2)
	sum, count := 0, 0
	for v := range out {
		sum += v
		count++
	}
	if count != 3 || sum != 12 {
		t.Errorf("got %d values summing to %d, want 3 and 12", count, sum)
	}
}

func TestStageClosesOnDrain(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(1), 4)
	<-out
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced an extra value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed after the input drained")
	}
}

func TestStageEmptyInput(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(), 3)
	select {
	case _, ok := <-out:
		if ok {
			t.Error("out produced a value")
		}
	case <-time.After(2 * time.Second):
		t.Error("out was never closed for an empty input")
	}
}

func TestStageZeroWorkers(t *testing.T) {
	done := make(chan struct{})
	defer close(done)
	out := Stage(done, feed(5), 0)
	if v, ok := <-out; !ok || v != 10 {
		t.Errorf("got %d, %v, want 10, true", v, ok)
	}
}

func TestStageDoesNotLeak(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()

	for round := 0; round < 20; round++ {
		done := make(chan struct{})
		in := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case in <- i:
				case <-done:
					close(in)
					return
				}
			}
		}()
		out := Stage(done, in, 4)
		<-out // take one value, then abandon the stage
		close(done)
	}

	deadline := time.Now().Add(3 * time.Second)
	for runtime.NumGoroutine() > base+4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+4 {
		t.Errorf("goroutines = %d, want about %d: workers are still blocked", got, base)
	}
}
""",
    context="A pipeline stage is cancelled when the consumer hits an error. The cancellation unblocks the receives and the workers all sit blocked on their sends instead.",
    task=[
        "Return a channel carrying each input doubled, computed by `workers` goroutines.",
        "Close it once the input drains.",
        "Every goroutine must exit when `done` is closed — whether it is receiving or sending.",
        "`workers < 1` behaves as 1.",
    ],
    examples=[
        ("Stage(done, feed(1,2,3), 2)", "a channel yielding 2, 4 and 6 in some order", None),
        ("input drained", "out is closed", None),
        ("consumer abandons out, done closed", "no goroutine left behind", None),
    ],
    topics=[
        ("Both ends can block", "Cancelling a receive is not enough if the send can block too."),
        ("select on every channel operation", "Each one needs the `done` escape."),
        ("WaitGroup then close", "The closer runs after all workers, in its own goroutine."),
        ("Only one closer", "Workers must not close a channel they share."),
    ],
    hint="Two blocking operations per iteration. Both need a way out.",
    intuition="Cancellation has to reach every place a goroutine can wait. Covering the receive and forgetting the send is the classic half-fix: the workers wake up, take a value, and block forever handing it on.",
    approach=[
        "Start `workers` goroutines tracked by a `WaitGroup`.",
        "Each loops: `select` to receive or return on `done`; return when the input is closed.",
        "`select` to send or return on `done`.",
        "A separate goroutine `Wait`s and closes the output.",
    ],
    walkthrough="With four workers and an abandoned consumer, closing `done` releases whichever `select` each worker is sitting in; the closer's `Wait` then returns and the output is closed.",
    pitfalls=[
        "`for v := range in` with a plain send — the range is cancellable only by closing `in`, and the send is not cancellable at all.",
        "Closing `out` from a worker, which panics when the second one does it.",
        "Calling `wg.Wait()` inside `Stage`, which blocks the caller until the pipeline is fully consumed.",
    ],
)
