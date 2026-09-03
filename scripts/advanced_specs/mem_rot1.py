"""10-advanced-topics / 01-memory-management-in-depth — rotation 1.

5 puzzles each for middle, senior and staff (junior was filled first).
"""

SUB = "01-memory-management-in-depth"

SPECS = []


def P(level, **kw):
    kw.setdefault("sub", SUB)
    kw["level"] = level
    kw.setdefault("mode", "stub")
    kw.setdefault("stub", 'panic("not implemented")')
    kw.setdefault("imports", [])
    kw.setdefault("extra", "")
    SPECS.append(kw)


# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="safeappend",
    title="An Append That Cannot Reach The Caller's Tail",
    mode="bug",
    sig="func Add(s []int, v int) []int",
    doc="""Add returns s with v appended.

The caller may be holding a longer slice over the same array. Appending
must never overwrite elements past len(s); the result must get its own
storage whenever that would happen.

Examples:

	Add([]int{1, 2}, 3) => []int{1, 2, 3}""",
    buggy="""return append(s, v)""",
    solution="""return append(s[:len(s):len(s)], v)""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add([]int{1, 2}, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Add = %v, want [1 2 3]", got)
	}
	if got := Add(nil, 1); !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("Add = %v, want [1]", got)
	}
}

func TestAddDoesNotClobberTheTail(t *testing.T) {
	backing := []int{1, 2, 3, 4}
	head := backing[:2]
	Add(head, 99)
	if backing[2] != 3 {
		t.Errorf("backing = %v, want [1 2 3 4]: the append wrote past the head", backing)
	}
}

func TestAddResultIsUsable(t *testing.T) {
	backing := []int{1, 2, 3, 4}
	got := Add(backing[:2], 99)
	if !reflect.DeepEqual(got, []int{1, 2, 99}) {
		t.Errorf("Add = %v, want [1 2 99]", got)
	}
}
""",
    context="Two request handlers hold overlapping views of one buffer. One of them appends, and the other's data silently changes — the classic aliasing corruption that only shows under load.",
    task=[
        "Append `v` to `s` and return the result.",
        "Elements of the backing array past `len(s)` must never be written.",
    ],
    examples=[
        ("Add([]int{1,2}, 3)", "[1 2 3]", None),
        ("b := []int{1,2,3,4}; Add(b[:2], 99)", "b is still [1 2 3 4]", "The spare capacity belongs to someone else."),
        ("Add(nil, 1)", "[1]", None),
    ],
    topics=[
        ("Spare capacity is shared", "`append` writes into `cap` before it reallocates — and that memory may not be yours."),
        ("Three-index slicing", "`s[:len(s):len(s)]` sets cap == len, so the next append must allocate."),
        ("Slice headers", "Two headers over one array is normal; only capacity makes it dangerous."),
    ],
    hint="`append` reallocates only when the capacity runs out. Make it run out.",
    intuition="`append` is allowed to write into the slice's spare capacity in place. When the slice is a prefix view of a longer array, that spare capacity is another view's live data.",
    approach=[
        "Cap the slice's capacity at its own length with a three-index slice.",
        "Append to that — the reallocation is now forced whenever room would have been borrowed.",
    ],
    walkthrough="`b[:2]` has len 2, cap 4. Plain `append` writes 99 into `b[2]`. `b[:2:2]` has cap 2, so `append` allocates a new array and `b` is untouched.",
    pitfalls=[
        "Copying unconditionally — correct, but it allocates even when the slice already owns its capacity.",
        "`s[:len(s)]` is not the fix; the third index is what matters.",
    ],
)

P(
    "middle",
    name="compactmap",
    title="A Map That Gives Its Buckets Back",
    sig="func Compact(m map[string]int) map[string]int",
    doc="""Compact returns a new map holding the same entries as m, sized to the
entries it actually keeps.

A map that grew to millions of entries keeps its bucket array after the
entries are deleted; rebuilding is the only way to release it.

Examples:

	Compact(map[string]int{"a": 1}) => a new map[a:1]""",
    solution="""out := make(map[string]int, len(m))
for k, v := range m {
	out[k] = v
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestCompactKeepsEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := Compact(m)
	if !reflect.DeepEqual(got, m) {
		t.Errorf("Compact = %v, want %v", got, m)
	}
}

func TestCompactReturnsANewMap(t *testing.T) {
	m := map[string]int{"a": 1}
	got := Compact(m)
	got["b"] = 2
	if _, ok := m["b"]; ok {
		t.Error("Compact returned the original map: the buckets are not released")
	}
}

func TestCompactAfterMassDeletion(t *testing.T) {
	m := make(map[string]int, 1<<12)
	for i := 0; i < 1<<12; i++ {
		m[string(rune('a'+i%26))+string(rune('a'+i/26))] = i
	}
	for k := range m {
		if m[k]%8 != 0 {
			delete(m, k)
		}
	}
	got := Compact(m)
	if len(got) != len(m) {
		t.Errorf("len = %d, want %d", len(got), len(m))
	}
	for k, v := range m {
		if got[k] != v {
			t.Fatalf("got[%q] = %d, want %d", k, got[k], v)
		}
	}
}

func TestCompactNil(t *testing.T) {
	if got := Compact(nil); len(got) != 0 {
		t.Errorf("Compact(nil) = %v, want empty", got)
	}
}
""",
    context="A session table peaks at ten million entries overnight and settles at a few thousand by morning. Resident memory never comes back down.",
    task=[
        "Return a new map with the same entries as `m`.",
        "Size the new map to the surviving entry count.",
        "A nil input returns an empty, usable map.",
    ],
    examples=[
        ('Compact(map[string]int{"a":1})', "map[a:1]", "Same entries, different map."),
        ("got := Compact(m); got[\"b\"]=2", "m does not gain \"b\"", None),
        ("Compact(nil)", "map[]", None),
    ],
    topics=[
        ("Maps never shrink", "`delete` frees the entry, not the bucket array."),
        ("Rebuild to release", "A fresh map sized to `len(m)` is the only way down."),
        ("Size hints", "`make(map[K]V, len(m))` avoids rehashing on the way in."),
    ],
    hint="`delete` in a loop will not help. What does help is a second map.",
    intuition="A Go map's bucket array only ever grows. Deleting entries leaves the buckets allocated and empty, so the only way to hand the memory back is to build a new map at the size you now need.",
    approach=[
        "`make(map[string]int, len(m))`.",
        "Range over `m` and copy every entry.",
        "Return the new map.",
    ],
    walkthrough="A map grown to 4096 entries and cut to 512 still holds the 4096-bucket array. Copying the 512 survivors into a map sized 512 lets the big array be collected.",
    pitfalls=[
        "Returning `m` itself when it is already small — the caller asked for a rebuild, and identity is what the test checks.",
        "Forgetting the size hint and rehashing all the way back up.",
    ],
)

P(
    "middle",
    name="fieldsview",
    title="Split Without Copying The Text",
    sig="func Fields(s string, sep byte) []string",
    doc="""Fields splits s on sep and returns the pieces.

Substrings of a string share the original bytes, so the pieces cost
nothing but their headers. Only the header slice may be allocated.

Examples:

	Fields("a,b,c", ',') => []string{"a", "b", "c"}""",
    imports=['"strings"'],
    solution="""n := strings.Count(s, string(sep)) + 1
out := make([]string, 0, n)
start := 0
for i := 0; i < len(s); i++ {
	if s[i] == sep {
		out = append(out, s[start:i])
		start = i + 1
	}
}
return append(out, s[start:])""",
    tests="""
import (
	"reflect"
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	if got := Fields("a,b,c", ','); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("Fields = %q, want [a b c]", got)
	}
	if got := Fields("", ','); !reflect.DeepEqual(got, []string{""}) {
		t.Errorf("Fields = %q, want [\\"\\"]", got)
	}
	if got := Fields("a,,b", ','); !reflect.DeepEqual(got, []string{"a", "", "b"}) {
		t.Errorf("Fields = %q, want [a  b]", got)
	}
	if got := Fields(",x", ','); !reflect.DeepEqual(got, []string{"", "x"}) {
		t.Errorf("Fields = %q, want [ x]", got)
	}
}

func TestFieldsAllocatesOnlyTheHeaders(t *testing.T) {
	line := strings.Repeat("column,", 63) + "last"
	if n := testing.AllocsPerRun(50, func() { _ = Fields(line, ',') }); n > 1 {
		t.Errorf("Fields made %v allocations, want 1: the pieces must be substrings", n)
	}
}
""",
    context="A CSV reader turns every field into a fresh string. For a wide file that is one small allocation per cell, and the cells outnumber everything else in the profile.",
    task=[
        "Split `s` on `sep` and return every piece, including empty ones.",
        "The pieces must be substrings of `s` — exactly one allocation per call, for the result slice.",
        "An empty input yields one empty piece.",
    ],
    examples=[
        ('Fields("a,b,c", \',\')', '["a" "b" "c"]', None),
        ('Fields("a,,b", \',\')', '["a" "" "b"]', "Empty fields are preserved."),
        ('Fields("", \',\')', '[""]', "No separator means one piece."),
    ],
    topics=[
        ("Substrings share bytes", "`s[a:b]` is a new header over the same immutable bytes."),
        ("Counting before allocating", "`strings.Count` gives the exact piece count for one `make`."),
        ("Immutability is what makes it safe", "Nobody can write through a string, so sharing needs no copy."),
    ],
    hint="A string cannot be modified, so why would a piece of one need its own bytes?",
    intuition="Strings are immutable, so a substring can point straight into the original bytes with no risk. Slicing is free; only the slice of headers has to be allocated, and its size is knowable in advance.",
    approach=[
        "Count the separators to size the result exactly.",
        "Walk the bytes, cutting a substring at each separator.",
        "Append the final piece after the loop.",
    ],
    walkthrough='For "a,b,c": cuts at index 1 and 3 give "a" and "b"; the tail "c" is appended after the loop. Three headers, zero bytes copied.',
    pitfalls=[
        "`string([]byte(...))` anywhere in the loop — that is the copy you are avoiding.",
        "Dropping the trailing piece by forgetting the append after the loop.",
    ],
)

P(
    "middle",
    name="appendints",
    title="Render Numbers Without Boxing Them",
    sig="func AppendInts(dst []byte, vals []int) []byte",
    doc="""AppendInts renders vals as decimal numbers separated by ' ' and
appends them to dst.

Passing an int to a variadic any parameter puts it on the heap. Rendering
must go straight into dst instead.

Examples:

	AppendInts(nil, []int{1, 2}) => []byte("1 2")""",
    imports=['"strconv"'],
    solution="""for i, v := range vals {
	if i > 0 {
		dst = append(dst, ' ')
	}
	dst = strconv.AppendInt(dst, int64(v), 10)
}
return dst""",
    tests="""
import (
	"bytes"
	"testing"
)

func TestAppendInts(t *testing.T) {
	if got := AppendInts(nil, []int{1, 2, 3}); !bytes.Equal(got, []byte("1 2 3")) {
		t.Errorf("AppendInts = %q, want \\"1 2 3\\"", got)
	}
	if got := AppendInts([]byte("x:"), []int{-4}); !bytes.Equal(got, []byte("x:-4")) {
		t.Errorf("AppendInts = %q, want \\"x:-4\\"", got)
	}
	if got := AppendInts([]byte("keep"), nil); !bytes.Equal(got, []byte("keep")) {
		t.Errorf("AppendInts = %q, want \\"keep\\"", got)
	}
}

func TestAppendIntsAllocatesNothingWithRoom(t *testing.T) {
	vals := make([]int, 32)
	for i := range vals {
		vals[i] = i * 7
	}
	dst := make([]byte, 0, 512)
	if n := testing.AllocsPerRun(100, func() { _ = AppendInts(dst[:0], vals) }); n != 0 {
		t.Errorf("AppendInts made %v allocations, want 0: render into dst", n)
	}
}
""",
    context="A metrics exporter builds every line with `fmt.Sprintf`. The formatter's variadic `any` parameter forces each number onto the heap, once per metric, once per scrape.",
    task=[
        "Append the decimal rendering of each value to `dst`, separated by a single space.",
        "Return the extended slice.",
        "With enough capacity in `dst`, the call must allocate nothing.",
    ],
    examples=[
        ("AppendInts(nil, []int{1,2,3})", '"1 2 3"', None),
        ('AppendInts([]byte("x:"), []int{-4})', '"x:-4"', "dst is extended, not replaced."),
        ('AppendInts([]byte("keep"), nil)', '"keep"', None),
    ],
    topics=[
        ("Interface boxing", "Storing an int in an `any` needs a heap word to point at."),
        ("strconv.Append*", "Renders straight into a byte slice with no intermediate string."),
        ("Caller-owned buffers", "An `Append`-style API lets the caller keep the memory."),
    ],
    hint="`strconv` has an `Append` twin for every `Format` function.",
    intuition="`fmt` takes `...any`. Every argument has to become an interface value, and an interface value needs a pointer — so the int gets a heap home just to be printed. `strconv.AppendInt` writes the digits directly.",
    approach=[
        "Range the values, writing a space before every one but the first.",
        "`dst = strconv.AppendInt(dst, int64(v), 10)`.",
        "Return `dst`.",
    ],
    walkthrough="Rendering 32 numbers with `fmt.Sprintf` boxes 32 ints and allocates 32 strings. The append form writes about 100 bytes into a buffer that already had room, allocating nothing.",
    pitfalls=[
        "Forgetting to reassign `dst` — `append`'s result is the only valid slice afterwards.",
        "Emitting a trailing separator; the first element is the special case, not the last.",
    ],
)

P(
    "middle",
    name="ringbuf",
    title="A Buffer That Allocates Once, Ever",
    sig="func (r *Ring) Push(v int)",
    doc="""Push adds v to the ring, overwriting the oldest element once the ring
is full.

The ring never grows: it was given its capacity at construction and every
later Push must reuse that storage.

Examples:

	r := NewRing(2); r.Push(1); r.Push(2); r.Push(3) => Items() is [2 3]""",
    extra="""// Ring is a fixed-capacity circular buffer of ints.
type Ring struct {
	buf  []int
	head int
	n    int
}

// NewRing returns a ring that holds at most cap elements.
func NewRing(cap int) *Ring {
	if cap < 1 {
		cap = 1
	}
	return &Ring{buf: make([]int, cap)}
}

// Len reports how many elements the ring currently holds.
func (r *Ring) Len() int { return r.n }

// Items returns the ring's contents from oldest to newest.
func (r *Ring) Items() []int {
	out := make([]int, 0, r.n)
	for i := 0; i < r.n; i++ {
		out = append(out, r.buf[(r.head+i)%len(r.buf)])
	}
	return out
}""",
    solution="""if r.n < len(r.buf) {
	r.buf[(r.head+r.n)%len(r.buf)] = v
	r.n++
	return
}
r.buf[r.head] = v
r.head = (r.head + 1) % len(r.buf)""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestRingFillsThenOverwrites(t *testing.T) {
	r := NewRing(3)
	for _, v := range []int{1, 2, 3} {
		r.Push(v)
	}
	if got := r.Items(); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("Items = %v, want [1 2 3]", got)
	}
	r.Push(4)
	if got := r.Items(); !reflect.DeepEqual(got, []int{2, 3, 4}) {
		t.Errorf("Items = %v, want [2 3 4]", got)
	}
	if r.Len() != 3 {
		t.Errorf("Len = %d, want 3", r.Len())
	}
}

func TestRingCapacityOne(t *testing.T) {
	r := NewRing(1)
	r.Push(1)
	r.Push(2)
	if got := r.Items(); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("Items = %v, want [2]", got)
	}
}

func TestPushNeverAllocates(t *testing.T) {
	r := NewRing(8)
	if n := testing.AllocsPerRun(1000, func() { r.Push(1) }); n != 0 {
		t.Errorf("Push made %v allocations, want 0: the ring must not grow", n)
	}
}
""",
    context="A crash reporter keeps the last 256 log lines in a slice it appends to and reslices. Under a log storm the slice grows without bound before the reslice ever runs.",
    task=[
        "Store `v` in the ring's existing buffer.",
        "When the ring is full, overwrite the oldest element and advance the head.",
        "`Push` must never allocate.",
    ],
    examples=[
        ("NewRing(3), push 1,2,3", "Items() is [1 2 3]", None),
        ("then push 4", "Items() is [2 3 4]", "The oldest element is overwritten."),
        ("NewRing(1), push 1 then 2", "Items() is [2]", None),
    ],
    topics=[
        ("Modular indexing", "`(head+n) % len(buf)` wraps the write position without branching on the end."),
        ("Fixed capacity", "Bounded memory is a design choice enforced by never calling `append`."),
        ("Pointer receivers", "`*Ring` is what lets `Push` mutate the ring the caller holds."),
    ],
    hint="Two cases: the ring has room, or it does not. Only the second one moves the head.",
    intuition="A ring turns \"keep the last N\" into pure index arithmetic. Nothing is ever appended, moved or freed — the write position simply walks around a fixed array.",
    approach=[
        "If `r.n < len(r.buf)`, write at `(head+n) % len(buf)` and increment `n`.",
        "Otherwise overwrite `buf[head]` and advance `head` by one, modulo the capacity.",
    ],
    walkthrough="Capacity 3, pushes 1,2,3 fill indices 0,1,2 with head 0. Push 4 overwrites index 0 and moves head to 1, so the oldest is now index 1 — [2 3 4].",
    pitfalls=[
        "Incrementing `n` past the capacity — the length is capped once the ring is full.",
        "Using `append` for the not-yet-full case; the buffer already has the slots.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="streamcount",
    title="Count The Lines Without Holding The File",
    sig="func CountLines(r io.Reader) (int, error)",
    doc="""CountLines returns the number of '\\n' bytes in r.

The reader may deliver gigabytes. The function must work in one pass over
a fixed-size buffer and must never hold the whole stream in memory.

Examples:

	CountLines(strings.NewReader("a\\nb\\n")) => 2, nil""",
    imports=['"bytes"', '"io"'],
    solution="""buf := make([]byte, 32*1024)
n := 0
for {
	c, err := r.Read(buf)
	n += bytes.Count(buf[:c], []byte{'\\n'})
	if err == io.EOF {
		return n, nil
	}
	if err != nil {
		return n, err
	}
}""",
    tests="""
import (
	"errors"
	"io"
	"runtime"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	n, err := CountLines(strings.NewReader("a\\nb\\nc"))
	if err != nil || n != 2 {
		t.Errorf("CountLines = %d, %v, want 2, nil", n, err)
	}
	if n, err := CountLines(strings.NewReader("")); err != nil || n != 0 {
		t.Errorf("CountLines = %d, %v, want 0, nil", n, err)
	}
	if n, err := CountLines(strings.NewReader("\\n\\n\\n")); err != nil || n != 3 {
		t.Errorf("CountLines = %d, %v, want 3, nil", n, err)
	}
}

type errReader struct{ n int }

func (e *errReader) Read(p []byte) (int, error) {
	if e.n == 0 {
		return 0, errors.New("boom")
	}
	e.n--
	p[0] = '\\n'
	return 1, nil
}

func TestCountLinesPropagatesErrors(t *testing.T) {
	if _, err := CountLines(&errReader{n: 2}); err == nil {
		t.Error("want the reader's error, got nil")
	}
}

// zeros yields n bytes, one '\\n' every 1024, without allocating.
type zeros struct{ left int64 }

func (z *zeros) Read(p []byte) (int, error) {
	if z.left <= 0 {
		return 0, io.EOF
	}
	c := int64(len(p))
	if c > z.left {
		c = z.left
	}
	for i := range p[:c] {
		if (z.left-int64(i))%1024 == 0 {
			p[i] = '\\n'
		} else {
			p[i] = 'x'
		}
	}
	z.left -= c
	return int(c), nil
}

func TestCountLinesStaysUnderTheMemoryCeiling(t *testing.T) {
	const size = 64 << 20
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	if _, err := CountLines(&zeros{left: size}); err != nil {
		t.Fatal(err)
	}
	runtime.ReadMemStats(&after)
	if used := after.TotalAlloc - before.TotalAlloc; used > 1<<20 {
		t.Errorf("allocated %d bytes for a %d byte stream, want under 1 MiB: do not buffer the input", used, size)
	}
}
""",
    context="A log tool calls `io.ReadAll` and then counts newlines. It works on the developer's sample and is killed by the OOM reaper on the first real file.",
    task=[
        "Return the number of `\\n` bytes in the stream.",
        "Read through a fixed-size buffer — total allocation must stay under 1 MiB regardless of the stream size.",
        "Return any read error other than `io.EOF` along with the count so far.",
    ],
    examples=[
        ('CountLines(strings.NewReader("a\\nb\\nc"))', "2, nil", "The final unterminated line is not counted."),
        ('CountLines(strings.NewReader(""))', "0, nil", None),
        ("a reader that fails", "the error", None),
    ],
    topics=[
        ("Streaming vs buffering", "Memory should track the buffer size, not the input size."),
        ("io.Reader contract", "`Read` may return `n > 0` together with an error; count those bytes first."),
        ("io.EOF is not a failure", "It ends the loop and returns a nil error."),
    ],
    hint="`io.ReadAll` is the bug, not the tool. What size is your working set allowed to be?",
    intuition="Counting needs one byte of context at a time, so there is no reason to hold the stream. A fixed buffer turns the memory cost from O(input) into O(1).",
    approach=[
        "Allocate one buffer of a fixed size, e.g. 32 KiB.",
        "Loop: `Read` into it, count newlines in `buf[:c]`.",
        "Stop on `io.EOF` with a nil error; return other errors with the count so far.",
    ],
    walkthrough="A 64 MiB stream is consumed in 2048 reads through one 32 KiB buffer. Total allocation is 32 KiB; `io.ReadAll` would have allocated well over 100 MiB across its doublings.",
    pitfalls=[
        "Ignoring the bytes returned alongside a non-nil error — those are real data.",
        "Allocating the buffer inside the loop, which turns O(1) memory into O(input) garbage.",
    ],
)

P(
    "senior",
    name="filterretain",
    title="The Filter That Keeps The Whole Batch Alive",
    mode="bug",
    sig="func Keep(records []Record, min int) []Record",
    doc="""Keep returns the records whose Size is at least min.

Typical batches are huge and typical results are tiny. The result must
not keep the batch's storage alive once the caller drops the batch.

Examples:

	Keep(batch, 100) => only the large records""",
    extra="""// Record is one ingested item.
type Record struct {
	ID   int
	Size int
	Pad  [64]byte
}""",
    buggy="""k := 0
for _, r := range records {
	if r.Size >= min {
		records[k] = r
		k++
	}
}
return records[:k]""",
    solution="""n := 0
for _, r := range records {
	if r.Size >= min {
		n++
	}
}
out := make([]Record, 0, n)
for _, r := range records {
	if r.Size >= min {
		out = append(out, r)
	}
}
return out""",
    tests="""
import "testing"

func TestKeepSelects(t *testing.T) {
	in := []Record{{ID: 1, Size: 10}, {ID: 2, Size: 200}, {ID: 3, Size: 300}}
	got := Keep(in, 100)
	if len(got) != 2 || got[0].ID != 2 || got[1].ID != 3 {
		t.Fatalf("Keep = %v, want records 2 and 3", got)
	}
	if got := Keep(nil, 1); len(got) != 0 {
		t.Errorf("Keep(nil) = %v, want empty", got)
	}
	if got := Keep([]Record{{Size: 1}}, 100); len(got) != 0 {
		t.Errorf("Keep = %v, want empty", got)
	}
}

func TestKeepDoesNotMutateTheInput(t *testing.T) {
	in := []Record{{ID: 1, Size: 10}, {ID: 2, Size: 200}}
	Keep(in, 100)
	if in[0].ID != 1 {
		t.Errorf("in[0].ID = %d, want 1: the batch was rewritten in place", in[0].ID)
	}
}

func TestKeepReleasesTheBatch(t *testing.T) {
	in := make([]Record, 1<<14)
	for i := range in {
		in[i] = Record{ID: i, Size: i}
	}
	got := Keep(in, 1<<14-4)
	if cap(got) > 64 {
		t.Errorf("cap = %d, want a right-sized result: it still owns the batch's array", cap(got))
	}
}
""",
    context="An ingest stage filters a 16k-record batch down to three records and hands them to a cache that lives for hours. Resident memory grows by a full batch for every cached result.",
    task=[
        "Return the records with `Size >= min`, in order.",
        "The result must own its storage — dropping the batch must free the batch.",
        "The input batch must not be modified.",
    ],
    examples=[
        ("Keep(batch, 100)", "the records of size >= 100", None),
        ("cap of the result", "the survivor count, not the batch size", "Otherwise the batch cannot be collected."),
        ("Keep(nil, 1)", "[]", None),
    ],
    topics=[
        ("Allocation-granular collection", "The collector frees whole allocations; one live element pins all of them."),
        ("In-place compaction has a cost", "It is allocation-free but returns a view of the input."),
        ("Two-pass sizing", "Counting first gives an exactly-sized result with one allocation."),
    ],
    hint="The compaction loop is efficient and wrong. What does `cap` of the returned slice tell you?",
    intuition="In-place compaction is the right move when the input dies with the result. Here the result outlives the batch, so a view of the batch is a leak of the batch — the survivors have to move to storage of their own.",
    approach=[
        "First pass: count the survivors.",
        "Allocate a result with exactly that capacity.",
        "Second pass: append the survivors.",
    ],
    walkthrough="16384 records of 80 bytes is about 1.3 MiB. Returning `records[:3]` keeps every byte of it reachable; copying three records into a 3-element array keeps 240 bytes.",
    pitfalls=[
        "`records[:k:k]` — caps the capacity, still points at the batch.",
        "Appending to a nil slice, which is correct but allocates through several growth steps.",
    ],
)

P(
    "senior",
    name="poolreset",
    title="The Pooled Buffer Nobody Emptied",
    mode="bug",
    sig="func Render(vals []int) string",
    doc="""Render returns vals as decimal numbers joined by ','.

The scratch buffer comes from a sync.Pool and goes back after use. A
buffer that comes out of a pool carries whatever the last borrower left
in it.

Examples:

	Render([]int{1, 2}) => "1,2" """,
    imports=['"strconv"', '"sync"'],
    extra="""// pool hands out reusable scratch buffers.
var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}""",
    buggy="""buf := pool.Get().([]byte)
for i, v := range vals {
	if i > 0 {
		buf = append(buf, ',')
	}
	buf = strconv.AppendInt(buf, int64(v), 10)
}
out := string(buf)
pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
return out""",
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

func TestRenderOnce(t *testing.T) {
	if got := Render([]int{1, 2, 3}); got != "1,2,3" {
		t.Errorf("Render = %q, want \\"1,2,3\\"", got)
	}
}

func TestRenderRepeatedly(t *testing.T) {
	for i := 0; i < 200; i++ {
		if got := Render([]int{7}); got != "7" {
			t.Fatalf("call %d: Render = %q, want \\"7\\": the pooled buffer was not reset", i, got)
		}
	}
}

func TestRenderStaysBounded(t *testing.T) {
	for i := 0; i < 500; i++ {
		Render([]int{1, 2, 3, 4})
	}
	b := pool.Get().([]byte)
	if cap(b) > 4096 {
		t.Errorf("pooled buffer grew to cap %d: every call is appending to the last one's output", cap(b))
	}
}
""",
    context="A hot serialiser was converted to a `sync.Pool` and the throughput improved. Two days later the output of one request starts appearing at the front of another's.",
    task=[
        "Render `vals` as decimal numbers joined by `,`.",
        "Fix the single bug so a borrowed buffer starts empty.",
        "The buffer must still go back to the pool.",
    ],
    examples=[
        ("Render([]int{1,2,3})", '"1,2,3"', None),
        ('200 calls of Render([]int{7})', 'every call returns "7"', "A pooled buffer is not a fresh one."),
        ("the pooled buffer's capacity after 500 calls", "bounded", "Otherwise every call appends to the last."),
    ],
    topics=[
        ("sync.Pool semantics", "`Get` returns a value someone else put back, in whatever state they left it."),
        ("Length vs capacity on reuse", "`[:0]` is what makes the capacity reusable and the contents gone."),
        ("Cross-request contamination", "Reuse bugs leak one caller's data into another's output."),
    ],
    hint="The first call is correct. The second one is not. What is different about the buffer it gets?",
    intuition="A pool recycles values, not blank slates. The buffer arrives with the previous borrower's length, so appending continues their output — the data leaks forward and the buffer grows without bound.",
    approach=[
        "Reslice the borrowed buffer to `[:0]` before writing.",
        "Append as before, convert to a string, return the buffer to the pool.",
    ],
    walkthrough='First call: `New` gives len 0, output "7", buffer returned at len 1. Second call: len 1, so the output is "77". By call 200 the buffer holds 200 sevens.',
    pitfalls=[
        "Resetting after use instead of before — another goroutine may already have taken it.",
        "Putting a buffer back while the caller still holds a slice of it.",
    ],
)

P(
    "senior",
    name="deferloop",
    title="Deferred Cleanup That Waits For The Whole Loop",
    mode="bug",
    sig="func Process(items []int, release func(int)) []int",
    doc="""Process doubles each item and calls release with the item as soon as
that item is finished.

release returns the item's resources. Holding every item until the
function returns is what makes a batch job run out of them.

Examples:

	Process([]int{1, 2}, rel) => []int{2, 4}, rel called after each item""",
    buggy="""out := make([]int, 0, len(items))
for _, v := range items {
	defer release(v)
	out = append(out, v*2)
}
return out""",
    solution="""out := make([]int, 0, len(items))
for _, v := range items {
	out = append(out, v*2)
	release(v)
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestProcessDoubles(t *testing.T) {
	got := Process([]int{1, 2, 3}, func(int) {})
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("Process = %v, want [2 4 6]", got)
	}
}

func TestReleaseHappensPerItem(t *testing.T) {
	held := 0
	peak := 0
	Process([]int{1, 2, 3, 4, 5}, func(int) { held-- })
	_ = peak

	held = 0
	Process([]int{1, 2, 3, 4, 5}, func(int) {
		held--
	})
	if held != -5 {
		t.Fatalf("release called %d times, want 5", -held)
	}
}

func TestReleaseOrderIsForward(t *testing.T) {
	var order []int
	Process([]int{1, 2, 3}, func(v int) { order = append(order, v) })
	if !reflect.DeepEqual(order, []int{1, 2, 3}) {
		t.Errorf("release order = %v, want [1 2 3]: each item must be released as it finishes", order)
	}
}

func TestReleaseIsNotStackedToTheEnd(t *testing.T) {
	outstanding := 0
	max := 0
	Process([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(int) { outstanding-- })
	_ = max
	if outstanding != -8 {
		t.Fatalf("release called %d times, want 8", -outstanding)
	}

	seen := make([]int, 0, 8)
	Process([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(v int) { seen = append(seen, v) })
	for i, v := range seen {
		if v != i+1 {
			t.Fatalf("release sequence = %v, want ascending: the cleanups piled up until the return", seen)
		}
	}
}
""",
    context="A batch importer opens a handle per row and defers the close. At row 60000 the process dies on \"too many open files\" — the deferred closes were all still waiting.",
    task=[
        "Return each item doubled, in order.",
        "Call `release` for each item as soon as that item is processed, in input order.",
        "Fix the single bug that keeps every item outstanding until the function returns.",
    ],
    examples=[
        ("Process([]int{1,2,3}, rel)", "[2 4 6]", None),
        ("release call order", "1, 2, 3", "Forward order, one per iteration."),
        ("outstanding items at any moment", "at most one", None),
    ],
    topics=[
        ("defer is function-scoped", "It runs when the function returns, not when the loop body ends."),
        ("LIFO order", "Stacked defers run in reverse, which is the wrong order here."),
        ("Resource lifetime", "Holding N resources instead of 1 is a memory and handle leak."),
    ],
    hint="Where does a deferred call actually run? Not where you wrote it.",
    intuition="`defer` binds to the function, not the block. In a loop it builds a stack of pending calls that all fire at the return — so the peak resource count is the whole loop, and the order is reversed.",
    approach=[
        "Drop the `defer`.",
        "Call `release(v)` directly at the end of each iteration.",
    ],
    walkthrough="With `defer`, releases for 1,2,3 run after the return in the order 3,2,1, and all three items are held at once. Called inline, item 1 is released before item 2 is touched.",
    pitfalls=[
        "Wrapping the body in a closure to keep `defer` — it works, but it is a function call per iteration to preserve a keyword you did not need.",
        "Assuming the reversed order is harmless; it usually is not for dependent resources.",
    ],
)

P(
    "senior",
    name="maxline",
    title="Longest Line, Bounded Working Set",
    sig="func MaxLine(r io.Reader) (int, error)",
    doc="""MaxLine returns the length in bytes of the longest '\\n'-separated line
in r, not counting the newline itself.

Lines may be longer than any single read, and the stream may be far
larger than memory. Only a fixed-size buffer may be held.

Examples:

	MaxLine(strings.NewReader("ab\\ncdef\\n")) => 4, nil""",
    imports=['"io"'],
    solution="""buf := make([]byte, 32*1024)
best, cur := 0, 0
for {
	c, err := r.Read(buf)
	for _, b := range buf[:c] {
		if b == '\\n' {
			if cur > best {
				best = cur
			}
			cur = 0
			continue
		}
		cur++
	}
	if err == io.EOF {
		if cur > best {
			best = cur
		}
		return best, nil
	}
	if err != nil {
		return 0, err
	}
}""",
    tests="""
import (
	"errors"
	"io"
	"runtime"
	"strings"
	"testing"
)

func TestMaxLine(t *testing.T) {
	if n, err := MaxLine(strings.NewReader("ab\\ncdef\\ng")); err != nil || n != 4 {
		t.Errorf("MaxLine = %d, %v, want 4, nil", n, err)
	}
	if n, err := MaxLine(strings.NewReader("")); err != nil || n != 0 {
		t.Errorf("MaxLine = %d, %v, want 0, nil", n, err)
	}
	if n, err := MaxLine(strings.NewReader("\\n\\n")); err != nil || n != 0 {
		t.Errorf("MaxLine = %d, %v, want 0, nil", n, err)
	}
	if n, err := MaxLine(strings.NewReader("tail-without-newline")); err != nil || n != 20 {
		t.Errorf("MaxLine = %d, %v, want 20, nil", n, err)
	}
}

type boom struct{}

func (boom) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestMaxLinePropagatesErrors(t *testing.T) {
	if _, err := MaxLine(boom{}); err == nil {
		t.Error("want an error, got nil")
	}
}

// long yields one line of size bytes then EOF.
type long struct{ left int64 }

func (l *long) Read(p []byte) (int, error) {
	if l.left <= 0 {
		return 0, io.EOF
	}
	c := int64(len(p))
	if c > l.left {
		c = l.left
	}
	for i := range p[:c] {
		p[i] = 'x'
	}
	l.left -= c
	return int(c), nil
}

func TestMaxLineDoesNotBufferTheLine(t *testing.T) {
	const size = 32 << 20
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	n, err := MaxLine(&long{left: size})
	if err != nil {
		t.Fatal(err)
	}
	if n != size {
		t.Fatalf("MaxLine = %d, want %d", n, size)
	}
	runtime.ReadMemStats(&after)
	if used := after.TotalAlloc - before.TotalAlloc; used > 1<<20 {
		t.Errorf("allocated %d bytes for one %d byte line, want under 1 MiB: count, do not accumulate", used, size)
	}
}
""",
    context="A validator uses `bufio.Scanner` to find the longest record and dies with `token too long` on a 32 MiB line — then dies of memory when the buffer limit is raised.",
    task=[
        "Return the byte length of the longest line, excluding the newline.",
        "A trailing line without a newline still counts.",
        "Hold only a fixed-size buffer: under 1 MiB total allocation for a 32 MiB line.",
    ],
    examples=[
        ('MaxLine(strings.NewReader("ab\\ncdef\\ng"))', "4, nil", "\"cdef\" is the longest."),
        ('MaxLine(strings.NewReader("tail-without-newline"))', "20, nil", "The last line needs no terminator."),
        ('MaxLine(strings.NewReader("\\n\\n"))', "0, nil", "Two empty lines."),
    ],
    topics=[
        ("Counting instead of collecting", "The answer is a number, so the line never has to exist in memory."),
        ("Reads do not align with lines", "The running count must survive across `Read` calls."),
        ("EOF finalisation", "The last line is only complete once the stream ends."),
    ],
    hint="You are asked for a length, not for the line.",
    intuition="The temptation is to accumulate the line so you can measure it. But the length is a running counter — it survives buffer boundaries for free, and the line itself never needs to be stored.",
    approach=[
        "Keep `cur` (current line length) and `best` across reads.",
        "For each byte: on `\\n`, fold `cur` into `best` and reset; otherwise increment `cur`.",
        "At EOF fold the final `cur` and return `best`.",
    ],
    walkthrough="A 32 MiB single line is consumed in 1024 reads through one 32 KiB buffer; `cur` reaches 33554432 and is folded into `best` at EOF. Nothing beyond the buffer is ever allocated.",
    pitfalls=[
        "Resetting `cur` at the start of each read instead of at each newline.",
        "Forgetting the unterminated final line.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="poolalias",
    title="The Result That Still Belongs To The Pool",
    mode="bug",
    sig="func Encode(vals []int) []byte",
    doc="""Encode returns vals rendered as decimal numbers joined by ','.

The scratch buffer is borrowed from a pool and returned before Encode
exits, so the result may not be a view of it: the next borrower would
overwrite the caller's data.

Examples:

	Encode([]int{1, 2}) => []byte("1,2")""",
    imports=['"strconv"', '"sync"'],
    extra="""// pool hands out reusable scratch buffers.
var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}""",
    buggy="""buf := pool.Get().([]byte)[:0]
for i, v := range vals {
	if i > 0 {
		buf = append(buf, ',')
	}
	buf = strconv.AppendInt(buf, int64(v), 10)
}
out := buf
pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
return out""",
    solution="""buf := pool.Get().([]byte)[:0]
for i, v := range vals {
	if i > 0 {
		buf = append(buf, ',')
	}
	buf = strconv.AppendInt(buf, int64(v), 10)
}
out := make([]byte, len(buf))
copy(out, buf)
pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
return out""",
    tests="""
import (
	"bytes"
	"strconv"
	"testing"
)

func TestEncode(t *testing.T) {
	if got := Encode([]int{1, 2, 3}); !bytes.Equal(got, []byte("1,2,3")) {
		t.Errorf("Encode = %q, want \\"1,2,3\\"", got)
	}
	if got := Encode(nil); len(got) != 0 {
		t.Errorf("Encode = %q, want empty", got)
	}
}

func TestEarlierResultsSurviveLaterCalls(t *testing.T) {
	first := Encode([]int{111, 222})
	for i := 0; i < 50; i++ {
		Encode([]int{999, 888})
	}
	if !bytes.Equal(first, []byte("111,222")) {
		t.Errorf("first = %q, want \\"111,222\\": the result was a view of a pooled buffer", first)
	}
}

func TestResultsAreIndependentOfEachOther(t *testing.T) {
	got := make([][]byte, 0, 32)
	for i := 0; i < 32; i++ {
		got = append(got, Encode([]int{i, i * 2}))
	}
	for i, b := range got {
		want := Reference(i)
		if !bytes.Equal(b, want) {
			t.Fatalf("result %d = %q, want %q", i, b, want)
		}
	}
}

// Reference renders the expected output without the pool.
func Reference(i int) []byte {
	var out []byte
	out = strconv.AppendInt(out, int64(i), 10)
	out = append(out, ',')
	return strconv.AppendInt(out, int64(i*2), 10)
}
""",
    context="A wire encoder was pooled to cut allocations. Latency improved and, days later, a small fraction of responses started carrying another request's payload.",
    task=[
        "Render `vals` as decimal numbers joined by `,`.",
        "The returned slice must be storage the caller owns.",
        "The scratch buffer must still go back to the pool.",
    ],
    examples=[
        ("Encode([]int{1,2,3})", '"1,2,3"', None),
        ("a result held across 50 later calls", "unchanged", "Otherwise the pool handed its bytes to someone else."),
        ("Encode(nil)", '""', None),
    ],
    topics=[
        ("Ownership across a pool boundary", "`Put` transfers the buffer; anything still viewing it is now a dangling reference in spirit."),
        ("Escape analysis and lifetime", "The result outlives the call; the scratch buffer must not."),
        ("Silent corruption", "No panic, no race detector hit — just wrong bytes under concurrency."),
    ],
    hint="Everything about the buffer is correct. Ask instead what the caller is holding when `Put` runs.",
    intuition="Returning a slice of a pooled buffer publishes memory you have just given away. The next `Get` hands the same array to another goroutine, which appends over the caller's result. Nothing is racy in the detector's sense — the write is simply legitimate and catastrophic.",
    approach=[
        "Build the text in the borrowed buffer as before.",
        "Copy it into a right-sized slice of its own.",
        "Return the buffer to the pool and return the copy.",
    ],
    walkthrough='Encode([]int{111,222}) writes seven bytes and returns a view. `Put` releases the array; the next call resets it to len 0 and writes "999,888" — and the first caller\'s slice now reads "999,888".',
    pitfalls=[
        "`string(buf)` then converting back — an extra copy for the same result.",
        "Putting the buffer back in a `defer` and returning the view anyway; the ordering does not save you.",
    ],
)

P(
    "staff",
    name="falseshare",
    title="Counters That Fight Over A Cache Line",
    sig="func Count(workers, iters int) int64",
    doc="""Count runs workers goroutines, each incrementing its own counter iters
times, and returns the total.

Each worker's counter must sit on its own cache line: adjacent counters
put the cores into a write-invalidate storm over one line.

Examples:

	Count(4, 1000) => 4000""",
    imports=['"sync"'],
    extra="""// cacheLine is the coherence granule the counters must not share.
const cacheLine = 64

// counter is one worker's slot.
type counter struct {
	n   int64
	pad [cacheLine - 8]byte
}""",
    solution="""if workers < 1 || iters < 0 {
	return 0
}
cs := make([]counter, workers)
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	go func(c *counter) {
		defer wg.Done()
		for i := 0; i < iters; i++ {
			c.n++
		}
	}(&cs[w])
}
wg.Wait()
var total int64
for i := range cs {
	total += cs[i].n
}
return total""",
    tests="""
import (
	"testing"
	"unsafe"
)

func TestCountTotal(t *testing.T) {
	if got := Count(4, 1000); got != 4000 {
		t.Errorf("Count = %d, want 4000", got)
	}
	if got := Count(1, 0); got != 0 {
		t.Errorf("Count = %d, want 0", got)
	}
	if got := Count(0, 10); got != 0 {
		t.Errorf("Count = %d, want 0", got)
	}
}

func TestCountUnderLoad(t *testing.T) {
	if got := Count(8, 100000); got != 800000 {
		t.Errorf("Count = %d, want 800000", got)
	}
}

func TestCountersDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(counter{}); got < cacheLine {
		t.Errorf("sizeof(counter) = %d, want at least %d: neighbouring counters share a cache line", got, cacheLine)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(4, 10000)
	}
}
""",
    context="A profiler shows a parallel counter loop running slower with eight cores than with one. No lock is contended and the race detector is silent.",
    task=[
        "Run `workers` goroutines, each incrementing its own `counter` `iters` times.",
        "Wait for all of them, then return the sum of the counters.",
        "Each goroutine must own its slot exclusively — no shared writes, no race.",
        "Non-positive `workers` or negative `iters` return 0.",
    ],
    examples=[
        ("Count(4, 1000)", "4000", None),
        ("Count(8, 100000)", "800000", "Correct under real parallelism."),
        ("Count(0, 10)", "0", None),
    ],
    topics=[
        ("False sharing", "Independent variables on one cache line are not independent to the hardware."),
        ("Padding", "`pad [56]byte` after an int64 pushes the next counter to its own line."),
        ("sync.WaitGroup", "The join point that makes the accumulation safe to read."),
        ("Loop-variable capture", "Pass the slot as a parameter rather than closing over the index."),
    ],
    hint="Every worker writes to its own variable and the cores still serialise. What is the unit of coherence?",
    intuition="Cache coherence works in lines, not variables. Two int64 counters eight bytes apart live on one 64-byte line, so every increment on one core invalidates the other's copy — the counters are logically private and physically shared.",
    approach=[
        "Reject non-positive input.",
        "Allocate one padded `counter` per worker.",
        "Start each goroutine with a pointer to its own slot; increment `iters` times.",
        "`wg.Wait()`, then sum the slots.",
    ],
    walkthrough="Eight unpadded counters fit in one line: 800000 increments become 800000 coherence transactions. With 64-byte padding each core owns its line and the increments stay in L1.",
    pitfalls=[
        "Summing before `Wait` — the totals are read while the workers are still writing.",
        "Capturing the loop index in the closure instead of passing the slot pointer.",
        "Padding the slice instead of the element; the gap must be inside the struct.",
    ],
)

P(
    "staff",
    name="sharedscratch",
    title="One Scratch Buffer, Many Goroutines",
    mode="bug",
    sig="func EncodeAll(batches [][]int) []string",
    doc="""EncodeAll renders every batch concurrently and returns the results in
input order.

Each goroutine must work in storage of its own; a buffer captured from
the enclosing scope is shared by all of them.

Examples:

	EncodeAll([][]int{{1}, {2}}) => []string{"1", "2"}""",
    imports=['"strconv"', '"sync"'],
    buggy="""out := make([]string, len(batches))
buf := make([]byte, 0, 64)
var wg sync.WaitGroup
wg.Add(len(batches))
for i, b := range batches {
	go func(i int, b []int) {
		defer wg.Done()
		buf = buf[:0]
		for j, v := range b {
			if j > 0 {
				buf = append(buf, ',')
			}
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		out[i] = string(buf)
	}(i, b)
}
wg.Wait()
return out""",
    solution="""out := make([]string, len(batches))
var wg sync.WaitGroup
wg.Add(len(batches))
for i, b := range batches {
	go func(i int, b []int) {
		defer wg.Done()
		buf := make([]byte, 0, 64)
		for j, v := range b {
			if j > 0 {
				buf = append(buf, ',')
			}
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		out[i] = string(buf)
	}(i, b)
}
wg.Wait()
return out""",
    tests="""
import (
	"strconv"
	"strings"
	"testing"
)

func TestEncodeAllSmall(t *testing.T) {
	got := EncodeAll([][]int{{1, 2}, {3}})
	if len(got) != 2 || got[0] != "1,2" || got[1] != "3" {
		t.Errorf("EncodeAll = %q, want [1,2 3]", got)
	}
	if got := EncodeAll(nil); len(got) != 0 {
		t.Errorf("EncodeAll = %q, want empty", got)
	}
}

func TestEncodeAllUnderConcurrency(t *testing.T) {
	const n = 64
	batches := make([][]int, n)
	for i := range batches {
		batches[i] = []int{i, i * 2, i * 3}
	}
	for round := 0; round < 20; round++ {
		got := EncodeAll(batches)
		for i := range batches {
			want := strings.Join([]string{
				strconv.Itoa(i), strconv.Itoa(i * 2), strconv.Itoa(i * 3),
			}, ",")
			if got[i] != want {
				t.Fatalf("round %d: result %d = %q, want %q: the goroutines share one buffer", round, i, got[i], want)
			}
		}
	}
}
""",
    context="A fan-out encoder passes its unit tests, passes review, and produces interleaved garbage the first time it runs with more than one core busy.",
    task=[
        "Render every batch concurrently as decimal numbers joined by `,`.",
        "Return the results in input order.",
        "Fix the single bug so each goroutine writes only to memory it owns.",
    ],
    examples=[
        ("EncodeAll([][]int{{1,2},{3}})", '["1,2" "3"]', None),
        ("64 batches, 20 rounds", "every result correct every time", "No goroutine may touch another's scratch."),
        ("EncodeAll(nil)", "[]", None),
    ],
    topics=[
        ("Captured variables are shared", "A closure over an outer variable gives every goroutine the same variable."),
        ("Data race vs logic race", "Two goroutines appending to one slice corrupt both results even when no write tears."),
        ("Disjoint writes are safe", "Writing `out[i]` from goroutine i needs no lock — the elements do not overlap."),
    ],
    hint="`i` and `b` were passed in as parameters for a reason. What else does the closure use?",
    intuition="Passing the loop variables in was half the job. The scratch buffer is still one variable captured by every goroutine, so all of them reset it, append to it and read it at once. `out[i]` is fine — those writes are disjoint.",
    approach=[
        "Move the buffer's declaration inside the goroutine.",
        "Leave the rest as is: each goroutine builds its own text and writes its own slot.",
    ],
    walkthrough="With a shared buffer, goroutine 3 can reset it to `[:0]` between goroutine 7's append and its `string(buf)`, so result 7 comes out truncated or holding batch 3's digits. A per-goroutine buffer makes the interleaving irrelevant.",
    pitfalls=[
        "Adding a mutex around the buffer — correct, and it serialises the whole fan-out.",
        "Assuming `-race` will always catch it; it reports only interleavings it actually observes.",
    ],
)

P(
    "staff",
    name="chunkworkers",
    title="Split The Work, Not The Memory",
    sig="func SumParallel(s []int, workers int) int64",
    doc="""SumParallel sums s using workers goroutines over disjoint chunks of
the input and returns the total.

The input must not be copied: each worker gets a view. Parallelism has to
be real — no locking on a shared accumulator per element.

Examples:

	SumParallel([]int{1, 2, 3, 4}, 2) => 10""",
    imports=['"sync"'],
    solution="""if workers < 1 {
	workers = 1
}
if len(s) == 0 {
	return 0
}
if workers > len(s) {
	workers = len(s)
}
partial := make([]int64, workers)
size := (len(s) + workers - 1) / workers
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	start := w * size
	end := start + size
	if start > len(s) {
		start = len(s)
	}
	if end > len(s) {
		end = len(s)
	}
	go func(w int, part []int) {
		defer wg.Done()
		var sum int64
		for _, v := range part {
			sum += int64(v)
		}
		partial[w] = sum
	}(w, s[start:end])
}
wg.Wait()
var total int64
for _, p := range partial {
	total += p
}
return total""",
    tests="""
import (
	"runtime"
	"testing"
)

func TestSumParallel(t *testing.T) {
	if got := SumParallel([]int{1, 2, 3, 4}, 2); got != 10 {
		t.Errorf("SumParallel = %d, want 10", got)
	}
	if got := SumParallel(nil, 4); got != 0 {
		t.Errorf("SumParallel = %d, want 0", got)
	}
	if got := SumParallel([]int{5}, 8); got != 5 {
		t.Errorf("SumParallel = %d, want 5: more workers than elements", got)
	}
	if got := SumParallel([]int{1, 2, 3}, 0); got != 6 {
		t.Errorf("SumParallel = %d, want 6: workers < 1 must still work", got)
	}
}

func TestSumParallelMatchesSerial(t *testing.T) {
	s := make([]int, 100003)
	var want int64
	for i := range s {
		s[i] = i % 977
		want += int64(s[i])
	}
	for _, w := range []int{1, 2, 3, 7, 16} {
		if got := SumParallel(s, w); got != want {
			t.Fatalf("SumParallel(_, %d) = %d, want %d: the chunks do not cover the input exactly once", w, got, want)
		}
	}
}

func TestSumParallelDoesNotCopyTheInput(t *testing.T) {
	s := make([]int, 1<<20)
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	SumParallel(s, 8)
	runtime.ReadMemStats(&after)
	if used := after.TotalAlloc - before.TotalAlloc; used > 1<<16 {
		t.Errorf("allocated %d bytes for an 8 MiB input, want under 64 KiB: pass views, not copies", used)
	}
}
""",
    context="A parallel aggregation was written by copying each worker's slice \"to be safe\". It is slower than the serial version and allocates a second copy of the dataset.",
    task=[
        "Sum `s` across `workers` goroutines over disjoint chunks.",
        "Every element must be counted exactly once, for any worker count.",
        "Pass views into `s`, never copies — under 64 KiB allocated for an 8 MiB input.",
        "`workers < 1` behaves as 1; more workers than elements is legal.",
    ],
    examples=[
        ("SumParallel([]int{1,2,3,4}, 2)", "10", None),
        ("SumParallel([]int{5}, 8)", "5", "Extra workers must not double-count or panic."),
        ("SumParallel([]int{1,2,3}, 0)", "6", None),
    ],
    topics=[
        ("Disjoint views", "Concurrent reads of one array need no synchronisation at all."),
        ("Per-worker accumulators", "One shared counter would serialise every element."),
        ("Chunk boundaries", "Ceiling-divided chunks must be clamped so the last one stops at `len(s)`."),
        ("Join before reading", "`wg.Wait()` is the happens-before edge that makes the partials visible."),
    ],
    hint="Reading the same array from many goroutines is free. Writing one accumulator from many goroutines is not.",
    intuition="Slices are views, so splitting the work costs nothing: every goroutine reads a region nobody else touches, which needs no synchronisation. Only the results have to come back, and one slot per worker plus a `Wait` is the whole protocol.",
    approach=[
        "Normalise `workers` against the input length.",
        "Compute the ceiling-divided chunk size and clamp each chunk's end.",
        "Give each goroutine `s[start:end]` and a private slot in `partial`.",
        "`Wait`, then sum the partials.",
    ],
    walkthrough="For 100003 elements and 7 workers, the chunk size is 14286; the last chunk is clamped to end at 100003. Each worker sums into its own slot, and the `Wait` publishes all seven before the final fold.",
    pitfalls=[
        "`atomic.AddInt64` per element — correct and slower than the serial loop.",
        "`start + size` without clamping, which panics on the last chunk.",
        "Summing `partial` before `Wait`.",
    ],
)

P(
    "staff",
    name="queuedrain",
    title="Close The Queue Without Leaking Its Workers",
    sig="func (q *Queue) Close() int64",
    doc="""Close stops accepting work, waits for the workers to finish what is
already queued, and returns the total they processed.

Every goroutine the Queue started must have exited by the time Close
returns; nothing may be left blocked on the channel.

Examples:

	q := NewQueue(4); q.Push(1); q.Close() => 1""",
    imports=['"sync"', '"sync/atomic"'],
    extra="""// Queue fans work out to a fixed set of workers.
type Queue struct {
	ch    chan int
	wg    sync.WaitGroup
	total atomic.Int64
	once  sync.Once
}

// NewQueue starts n workers, each accumulating the values it receives.
func NewQueue(n int) *Queue {
	if n < 1 {
		n = 1
	}
	q := &Queue{ch: make(chan int, 16)}
	q.wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer q.wg.Done()
			for v := range q.ch {
				q.total.Add(int64(v))
			}
		}()
	}
	return q
}

// Push submits one value. It must not be called after Close.
func (q *Queue) Push(v int) { q.ch <- v }""",
    solution="""q.once.Do(func() { close(q.ch) })
q.wg.Wait()
return q.total.Load()""",
    tests="""
import (
	"runtime"
	"testing"
	"time"
)

func TestQueueTotals(t *testing.T) {
	q := NewQueue(4)
	for i := 1; i <= 100; i++ {
		q.Push(i)
	}
	if got := q.Close(); got != 5050 {
		t.Errorf("Close = %d, want 5050", got)
	}
}

func TestCloseIsIdempotent(t *testing.T) {
	q := NewQueue(2)
	q.Push(7)
	first := q.Close()
	if second := q.Close(); second != first {
		t.Errorf("second Close = %d, want %d", second, first)
	}
}

func TestCloseLeavesNoGoroutines(t *testing.T) {
	runtime.GC()
	base := runtime.NumGoroutine()
	for i := 0; i < 20; i++ {
		q := NewQueue(4)
		q.Push(i)
		q.Close()
	}
	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > base+2 && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if got := runtime.NumGoroutine(); got > base+2 {
		t.Errorf("goroutines = %d, want about %d: the workers are still blocked on the channel", got, base)
	}
}

func TestCloseWaitsForQueuedWork(t *testing.T) {
	q := NewQueue(1)
	for i := 0; i < 16; i++ {
		q.Push(1)
	}
	if got := q.Close(); got != 16 {
		t.Errorf("Close = %d, want 16: Close returned before the queue drained", got)
	}
}
""",
    context="A service creates a queue per request and closes it in a defer. Goroutine count climbs all day and the process is restarted every night to keep it under control.",
    task=[
        "Stop the workers, wait for the already-queued values to be processed, and return the total.",
        "No worker goroutine may survive `Close`.",
        "A second `Close` must return the same total without panicking.",
    ],
    examples=[
        ("push 1..100 into NewQueue(4), Close()", "5050", None),
        ("Close() twice", "the same total, no panic", "Closing a closed channel panics; guard it."),
        ("16 values, one worker", "16", "Close returns only after the backlog drains."),
    ],
    topics=[
        ("range over a channel", "The loop ends when the channel is closed and drained — that is the exit signal."),
        ("Closing is the sender's job", "Only the side that stops sending may close."),
        ("WaitGroup as the join", "`Wait` is what makes the accumulated total safe to read."),
        ("sync.Once", "Idempotent shutdown without a flag and a mutex."),
    ],
    hint="A `for range` over a channel exits on exactly one event. Nothing else will ever wake those workers.",
    intuition="A worker blocked in `range ch` is not idle — it is a live goroutine holding its stack and everything its frame references. Only closing the channel ends the loop, and only the `WaitGroup` tells you it actually ended.",
    approach=[
        "Close the channel exactly once, guarded by `sync.Once`.",
        "`wg.Wait()` so every worker has drained and exited.",
        "Return the accumulated total.",
    ],
    walkthrough="With four workers and a backlog of ten values, `close` lets each `range` drain what is left and then return; `Wait` unblocks once all four have called `Done`, at which point the total is complete and stable.",
    pitfalls=[
        "`Wait` before `close` — the workers never exit and `Close` deadlocks.",
        "Closing without `Once`; the second `Close` panics on a closed channel.",
        "Reading `total` without the `Wait`, which reports a partial sum.",
    ],
)
