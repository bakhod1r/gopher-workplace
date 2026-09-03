"""10-advanced-topics / 02-escape-analysis — rotation 1: 5 puzzles per level."""

SUB = "02-escape-analysis"

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
    name="retptr",
    title="A Pointer That Outlives Its Frame",
    sig="func New(v int) *int",
    doc="""New returns a pointer to a fresh int holding v.

The pointer outlives the call, so the int cannot live in the frame — the
compiler moves it to the heap. That is one allocation, and exactly one.

Examples:

	*New(7) => 7""",
    solution="""p := v
return &p""",
    tests="""
import "testing"

var sink *int

func TestNew(t *testing.T) {
	p := New(7)
	if p == nil || *p != 7 {
		t.Fatalf("New(7) = %v, want a pointer to 7", p)
	}
	if got := New(0); *got != 0 {
		t.Errorf("*New(0) = %d, want 0", *got)
	}
}

func TestNewReturnsDistinctPointers(t *testing.T) {
	a, b := New(1), New(1)
	if a == b {
		t.Error("New returned the same pointer twice")
	}
	*a = 99
	if *b != 1 {
		t.Errorf("*b = %d, want 1: the two ints share storage", *b)
	}
}

func TestNewAllocatesOnce(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { sink = New(3) }); n != 1 {
		t.Errorf("New made %v allocations, want exactly 1", n)
	}
}
""",
    context="A new Go developer is told \"pointers are fast, values are copies\" and starts returning pointers everywhere. The allocation profile disagrees.",
    task=[
        "Return a pointer to a new int holding `v`.",
        "Each call must return a distinct pointer.",
        "Exactly one allocation per call.",
    ],
    examples=[
        ("*New(7)", "7", None),
        ("a, b := New(1), New(1)", "a != b", "Every call gets its own int."),
        ("allocations per call", "1", "The int must escape to the heap."),
    ],
    topics=[
        ("Escape analysis", "The compiler decides stack or heap by asking whether the value outlives the frame."),
        ("Returning &local", "Legal in Go — the value is simply moved to the heap."),
        ("Pointer identity", "Distinct allocations mean distinct addresses."),
    ],
    hint="In C this would be a dangling pointer. In Go it is an allocation.",
    intuition="Go has no dangling pointers because the compiler proves what escapes. Taking the address of a local and returning it does not break — it just means the local cannot live on the stack.",
    approach=[
        "Copy `v` into a local.",
        "Return the local's address.",
    ],
    walkthrough="`p := v; return &p` — `p` is referenced after the frame dies, so the compiler allocates it on the heap. `new(int)` plus an assignment does the same thing.",
    pitfalls=[
        "Returning `&v` directly works too; the parameter is a local like any other.",
        "Expecting zero allocations — a heap escape is the point here, not a bug.",
    ],
)

P(
    "junior",
    name="byvalue",
    title="Small Values Stay On The Stack",
    sig="func Scale(p Point, f int) Point",
    doc="""Scale returns p with both coordinates multiplied by f.

Nothing about the result outlives the call, so nothing needs the heap.

Examples:

	Scale(Point{2, 3}, 2) => Point{4, 6}""",
    extra="""// Point is a two-dimensional integer point.
type Point struct {
	X, Y int
}""",
    solution="""p.X *= f
p.Y *= f
return p""",
    tests="""
import "testing"

var sink Point

func TestScale(t *testing.T) {
	if got := Scale(Point{2, 3}, 2); got != (Point{4, 6}) {
		t.Errorf("Scale = %v, want {4 6}", got)
	}
	if got := Scale(Point{1, 1}, 0); got != (Point{0, 0}) {
		t.Errorf("Scale = %v, want {0 0}", got)
	}
	if got := Scale(Point{-1, 2}, 3); got != (Point{-3, 6}) {
		t.Errorf("Scale = %v, want {-3 6}", got)
	}
}

func TestScaleDoesNotTouchTheCaller(t *testing.T) {
	p := Point{2, 3}
	Scale(p, 5)
	if p != (Point{2, 3}) {
		t.Errorf("p = %v, want {2 3}: the parameter is a copy", p)
	}
}

func TestScaleAllocatesNothing(t *testing.T) {
	p := Point{2, 3}
	if n := testing.AllocsPerRun(100, func() { sink = Scale(p, 2) }); n != 0 {
		t.Errorf("Scale made %v allocations, want 0", n)
	}
}
""",
    context="A geometry package returns `*Point` from every helper \"to avoid copying\". Each helper costs an allocation and the copies it avoided were sixteen bytes.",
    task=[
        "Return `p` with both coordinates multiplied by `f`.",
        "The caller's `Point` must not change.",
        "Zero allocations.",
    ],
    examples=[
        ("Scale(Point{2,3}, 2)", "{4 6}", None),
        ("p := Point{2,3}; Scale(p, 5)", "p is still {2 3}", "Structs are passed by value."),
        ("Scale(Point{1,1}, 0)", "{0 0}", None),
    ],
    topics=[
        ("Struct value semantics", "Passing and returning a struct copies its fields."),
        ("Stack allocation", "A value that never escapes needs no heap."),
        ("Copy cost vs allocation cost", "Copying two ints is cheaper than allocating one pointer."),
    ],
    hint="The parameter is already your own copy. Mutate it and hand it back.",
    intuition="A sixteen-byte copy is a couple of register moves. An allocation is a trip to the allocator plus work for the collector later. For small structs the copy wins every time.",
    approach=[
        "Multiply the parameter's fields in place.",
        "Return the parameter.",
    ],
    walkthrough="`Scale(Point{2,3}, 2)` copies two ints in, doubles them, copies two ints out. The escape analyser sees no reference leaving the frame, so the point stays in registers.",
    pitfalls=[
        "Returning `*Point` — an allocation to avoid a two-word copy.",
        "Expecting the caller's point to change; take a pointer parameter if that is the intent.",
    ],
)

P(
    "junior",
    name="outparam",
    title="Let The Caller Own The Buffer",
    sig="func Fill(dst []byte, v byte) int",
    doc="""Fill writes v into every byte of dst and returns how many bytes were
written.

The buffer belongs to the caller, so the function allocates nothing at
all.

Examples:

	Fill(make([]byte, 3), 'x') => 3, buffer is "xxx" """,
    solution="""for i := range dst {
	dst[i] = v
}
return len(dst)""",
    tests="""
import (
	"bytes"
	"testing"
)

func TestFill(t *testing.T) {
	buf := make([]byte, 3)
	if n := Fill(buf, 'x'); n != 3 {
		t.Errorf("Fill = %d, want 3", n)
	}
	if !bytes.Equal(buf, []byte("xxx")) {
		t.Errorf("buf = %q, want \\"xxx\\"", buf)
	}
	if n := Fill(nil, 'x'); n != 0 {
		t.Errorf("Fill(nil) = %d, want 0", n)
	}
}

func TestFillWritesOnlyTheView(t *testing.T) {
	buf := []byte("abcd")
	Fill(buf[1:3], 'z')
	if !bytes.Equal(buf, []byte("azzd")) {
		t.Errorf("buf = %q, want \\"azzd\\"", buf)
	}
}

func TestFillAllocatesNothing(t *testing.T) {
	buf := make([]byte, 256)
	if n := testing.AllocsPerRun(100, func() { Fill(buf, 1) }); n != 0 {
		t.Errorf("Fill made %v allocations, want 0", n)
	}
}
""",
    context="A codec allocates and returns a fresh buffer on every call. The caller has a perfectly good buffer it would rather reuse, but the API gives it no way to say so.",
    task=[
        "Write `v` into every byte of `dst`.",
        "Return the number of bytes written.",
        "Allocate nothing — the storage is the caller's.",
    ],
    examples=[
        ("Fill(make([]byte,3), 'x')", '3, buffer "xxx"', None),
        ('buf := []byte("abcd"); Fill(buf[1:3], \'z\')', '"azzd"', "Only the view is written."),
        ("Fill(nil, 'x')", "0", None),
    ],
    topics=[
        ("Caller-owned buffers", "A `dst` parameter moves the allocation decision to the caller."),
        ("Nothing escapes", "The function keeps no reference after it returns."),
        ("Views", "A sub-slice limits exactly which bytes may be touched."),
    ],
    hint="You are given the memory. Do not ask for more.",
    intuition="An API that returns freshly allocated memory forces an allocation on every call. An API that fills a buffer lets the caller allocate once and reuse forever — the same work, none of the garbage.",
    approach=[
        "Range over `dst` by index and assign `v`.",
        "Return `len(dst)`.",
    ],
    walkthrough="`Fill(buf, 1)` on a 256-byte buffer writes 256 bytes and returns. No reference to `dst` survives the call, so the escape analyser has nothing to move.",
    pitfalls=[
        "Allocating a temporary and copying it into `dst`.",
        "Writing past `len(dst)` into the capacity — that memory is not yours.",
    ],
)

P(
    "junior",
    name="localarray",
    title="A Scratch Array That Never Leaves",
    sig="func Digits(n int) int",
    doc="""Digits returns the number of decimal digits in the absolute value of n,
computing them into a fixed-size local array.

The array never escapes, so the whole function is allocation-free.

Examples:

	Digits(1234) => 4""",
    solution="""var buf [20]byte
if n < 0 {
	n = -n
}
i := 0
for {
	buf[i] = byte('0' + n%10)
	i++
	n /= 10
	if n == 0 {
		return i
	}
}""",
    tests="""
import "testing"

var sink int

func TestDigits(t *testing.T) {
	cases := map[int]int{0: 1, 7: 1, 10: 2, 1234: 4, -99: 2, 1000000: 7}
	for in, want := range cases {
		if got := Digits(in); got != want {
			t.Errorf("Digits(%d) = %d, want %d", in, got, want)
		}
	}
}

func TestDigitsAllocatesNothing(t *testing.T) {
	if n := testing.AllocsPerRun(200, func() { sink = Digits(987654321) }); n != 0 {
		t.Errorf("Digits made %v allocations, want 0: the scratch array must not escape", n)
	}
}
""",
    context="A logging helper counts digits by rendering the number with `strconv.Itoa` and taking the length. One allocation per log line, on the hottest path in the service.",
    task=[
        "Return how many decimal digits `n` has, ignoring the sign.",
        "`Digits(0)` is 1.",
        "Use a fixed-size local array; the function must allocate nothing.",
    ],
    examples=[
        ("Digits(1234)", "4", None),
        ("Digits(0)", "1", "Zero is one digit."),
        ("Digits(-99)", "2", "The sign does not count."),
    ],
    topics=[
        ("Fixed-size locals", "`var buf [20]byte` is a stack object when it does not escape."),
        ("Arrays vs slices", "An array has its size in its type, so the compiler can place it in the frame."),
        ("Avoiding conversions", "`strconv.Itoa` allocates a string you throw away."),
    ],
    hint="Twenty bytes is enough for any int64. Where should those twenty bytes live?",
    intuition="A local array whose address never leaves the function is just part of the frame. That makes it free — no allocator, no collector, no cost beyond the stack pointer already moving.",
    approach=[
        "Declare a fixed local array big enough for any int.",
        "Take the absolute value.",
        "Peel digits with `%10` and `/10`, counting until the number reaches zero.",
    ],
    walkthrough="1234 peels to 4, 3, 2, 1 — four iterations, so the answer is 4. The array is written but never read and never escapes, so the frame is the only memory involved.",
    pitfalls=[
        "Returning a slice of the local array — that makes it escape.",
        "Looping `for n > 0`, which returns 0 for the input 0.",
    ],
)

P(
    "junior",
    name="sentinelerr",
    title="An Error That Costs Nothing To Return",
    sig="func Validate(n int) error",
    doc="""Validate reports whether n is a usable count: it must be non-negative
and no greater than MaxCount.

The failures are fixed conditions, so they must be reported with the
package's sentinel errors rather than a freshly formatted one.

Examples:

	Validate(-1) => ErrNegative""",
    imports=['"errors"'],
    extra="""// MaxCount is the largest count Validate accepts.
const MaxCount = 1000

// The conditions Validate can report.
var (
	ErrNegative = errors.New("count is negative")
	ErrTooLarge = errors.New("count is too large")
)""",
    solution="""if n < 0 {
	return ErrNegative
}
if n > MaxCount {
	return ErrTooLarge
}
return nil""",
    tests="""
import (
	"errors"
	"testing"
)

var sink error

func TestValidate(t *testing.T) {
	if err := Validate(5); err != nil {
		t.Errorf("Validate(5) = %v, want nil", err)
	}
	if err := Validate(-1); !errors.Is(err, ErrNegative) {
		t.Errorf("Validate(-1) = %v, want ErrNegative", err)
	}
	if err := Validate(MaxCount + 1); !errors.Is(err, ErrTooLarge) {
		t.Errorf("Validate = %v, want ErrTooLarge", err)
	}
	if err := Validate(0); err != nil {
		t.Errorf("Validate(0) = %v, want nil", err)
	}
	if err := Validate(MaxCount); err != nil {
		t.Errorf("Validate(MaxCount) = %v, want nil", err)
	}
}

func TestValidateAllocatesNothing(t *testing.T) {
	if n := testing.AllocsPerRun(200, func() { sink = Validate(-1) }); n != 0 {
		t.Errorf("Validate made %v allocations, want 0: return the sentinel, do not format one", n)
	}
}
""",
    context="An input validator on a public endpoint builds its error with `fmt.Errorf` on every rejected request. A burst of bad traffic turns error reporting into the biggest allocator in the process.",
    task=[
        "Return `ErrNegative` for `n < 0` and `ErrTooLarge` for `n > MaxCount`.",
        "Return nil otherwise; the boundaries 0 and `MaxCount` are valid.",
        "Zero allocations, including on the failing paths.",
    ],
    examples=[
        ("Validate(5)", "<nil>", None),
        ("Validate(-1)", "ErrNegative", "The same value every time — comparable with errors.Is."),
        ("Validate(1001)", "ErrTooLarge", None),
    ],
    topics=[
        ("Sentinel errors", "One package-level value, created once at init, returned forever."),
        ("errors.Is", "Comparison against a sentinel is what makes callers able to branch."),
        ("fmt.Errorf allocates", "Formatting builds a new string and a new error value on every call."),
    ],
    hint="The two failures are already declared. Return them.",
    intuition="An error carrying no per-call information does not need to be built per call. A package-level sentinel is allocated once at init and returned by pointer forever after.",
    approach=[
        "Check `n < 0` and return `ErrNegative`.",
        "Check `n > MaxCount` and return `ErrTooLarge`.",
        "Return nil.",
    ],
    walkthrough="`Validate(-1)` returns the interface value pointing at the existing `ErrNegative`. Nothing is constructed, so `AllocsPerRun` reports 0. `fmt.Errorf(\"count is negative\")` would allocate a string and an error struct on every rejection.",
    pitfalls=[
        "Wrapping the sentinel with `fmt.Errorf(\"%w\", ...)` when there is nothing to add.",
        "Comparing with `==` in the caller is fine for a bare sentinel, but `errors.Is` survives later wrapping.",
    ],
)

# ---------------------------------------------------------------- middle -----

P(
    "middle",
    name="appendto",
    title="Append Into The Caller's Slice",
    sig="func AppendSquares(dst []int, n int) []int",
    doc="""AppendSquares appends the squares 0..n-1 to dst and returns the
extended slice.

When dst already has the capacity the call must allocate nothing: the
result is the caller's memory, not the function's.

Examples:

	AppendSquares(nil, 3) => []int{0, 1, 4}""",
    solution="""for i := 0; i < n; i++ {
	dst = append(dst, i*i)
}
return dst""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestAppendSquares(t *testing.T) {
	if got := AppendSquares(nil, 4); !reflect.DeepEqual(got, []int{0, 1, 4, 9}) {
		t.Errorf("AppendSquares = %v, want [0 1 4 9]", got)
	}
	if got := AppendSquares([]int{7}, 2); !reflect.DeepEqual(got, []int{7, 0, 1}) {
		t.Errorf("AppendSquares = %v, want [7 0 1]", got)
	}
	if got := AppendSquares([]int{7}, 0); !reflect.DeepEqual(got, []int{7}) {
		t.Errorf("AppendSquares = %v, want [7]", got)
	}
	if got := AppendSquares(nil, -1); len(got) != 0 {
		t.Errorf("AppendSquares = %v, want empty", got)
	}
}

func TestAppendSquaresReusesCapacity(t *testing.T) {
	dst := make([]int, 0, 64)
	if n := testing.AllocsPerRun(100, func() { _ = AppendSquares(dst[:0], 64) }); n != 0 {
		t.Errorf("AppendSquares made %v allocations, want 0 when dst has room", n)
	}
}

func TestAppendSquaresDoesNotClobberBeyondLen(t *testing.T) {
	backing := make([]int, 4, 8)
	got := AppendSquares(backing[:1], 2)
	if !reflect.DeepEqual(got, []int{0, 0, 1}) {
		t.Errorf("got = %v, want [0 0 1]", got)
	}
}
""",
    context="Every helper in a hot pipeline returns a freshly made slice. The caller immediately copies each result into its own buffer and drops the original.",
    task=[
        "Append the squares of `0..n-1` to `dst` and return the result.",
        "With enough capacity in `dst`, allocate nothing.",
        "`n <= 0` returns `dst` unchanged.",
    ],
    examples=[
        ("AppendSquares(nil, 4)", "[0 1 4 9]", None),
        ("AppendSquares([]int{7}, 2)", "[7 0 1]", "dst is extended, not replaced."),
        ("dst with cap 64, n = 64", "0 allocations", None),
    ],
    topics=[
        ("Append-style APIs", "`f(dst, ...) []T` is the idiom for handing allocation control to the caller."),
        ("Reusing capacity", "`dst[:0]` between calls turns a per-call allocation into none."),
        ("append returns a new header", "The result must be reassigned; the old header may be stale."),
    ],
    hint="The signature already tells you where the output goes.",
    intuition="Who allocates is an API decision. Taking a `dst` and returning the extended slice lets a caller in a loop allocate once and reuse, while a caller that does not care can still pass nil.",
    approach=[
        "Loop `i` from 0 to `n-1`.",
        "`dst = append(dst, i*i)`.",
        "Return `dst`.",
    ],
    walkthrough="With `dst[:0]` over a 64-element array, all 64 appends fit in the existing capacity and the allocator is never called. With `dst == nil`, `append` grows from nothing as usual.",
    pitfalls=[
        "Ignoring `append`'s return value.",
        "Allocating a local result and copying into `dst`, which defeats the whole point.",
    ],
)

P(
    "middle",
    name="closurecapture",
    title="What A Closure Drags Onto The Heap",
    sig="func Counter(start int) func() int",
    doc="""Counter returns a function that yields start, start+1, start+2 and so
on, one value per call.

The captured variable outlives Counter's frame, so it must live on the
heap — that is what a closure over a mutable local costs.

Examples:

	c := Counter(1); c(), c() => 1, 2""",
    solution="""n := start
return func() int {
	v := n
	n++
	return v
}""",
    tests="""
import "testing"

var sink func() int

func TestCounter(t *testing.T) {
	c := Counter(1)
	for want := 1; want <= 4; want++ {
		if got := c(); got != want {
			t.Fatalf("call = %d, want %d", got, want)
		}
	}
}

func TestCountersAreIndependent(t *testing.T) {
	a, b := Counter(10), Counter(10)
	a()
	a()
	if got := b(); got != 10 {
		t.Errorf("b() = %d, want 10: the counters share state", got)
	}
}

func TestCounterFromZeroAndNegative(t *testing.T) {
	c := Counter(-2)
	if got := []int{c(), c(), c()}; got[0] != -2 || got[1] != -1 || got[2] != 0 {
		t.Errorf("got %v, want [-2 -1 0]", got)
	}
}

func TestCounterAllocationsAreBounded(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { sink = Counter(0) }); n > 2 {
		t.Errorf("Counter made %v allocations, want at most 2", n)
	}
}
""",
    context="A team bans closures from a hot path after seeing them in an allocation profile, then discovers half the codebase's iterators are closures. Nobody can say which ones actually cost anything.",
    task=[
        "Return a function that yields `start`, then `start+1`, and so on.",
        "Two counters must not share state.",
        "At most two allocations per `Counter` call.",
    ],
    examples=[
        ("c := Counter(1); c(); c()", "1, then 2", None),
        ("a, b := Counter(10), Counter(10); a(); a(); b()", "10", "Independent captured state."),
        ("Counter(-2) called three times", "-2, -1, 0", None),
    ],
    topics=[
        ("Closures capture by reference", "The variable itself is shared with the returned function, not a copy of it."),
        ("Escape via capture", "A captured variable that outlives the frame moves to the heap."),
        ("Func values", "A closure is a pointer to code plus a pointer to its captured environment."),
    ],
    hint="The counter's state has to live somewhere after `Counter` returns.",
    intuition="A closure over a mutable local turns that local into shared state between two frames — the one that made it and every call of the closure. It cannot stay in a frame that is about to disappear, so it is allocated.",
    approach=[
        "Copy `start` into a local.",
        "Return a function that reads it, increments it, and returns the old value.",
    ],
    walkthrough="`Counter(1)` allocates the captured `n` and the closure object. Each `c()` reads and bumps the same heap word, which is why two counters made from the same start stay independent but two calls of one counter do not.",
    pitfalls=[
        "Incrementing before returning, which yields `start+1` first.",
        "Capturing the parameter directly is fine — it is a local too.",
    ],
)

P(
    "middle",
    name="bufinloop",
    title="The Buffer Allocated Once Per Iteration",
    mode="bug",
    sig="func Render(rows [][]int) []string",
    doc="""Render turns each row into a comma-separated string.

The scratch buffer is per-call state, not per-row state: allocating it
inside the loop makes one throwaway buffer for every row.

Examples:

	Render([][]int{{1, 2}}) => []string{"1,2"}""",
    imports=['"strconv"'],
    extra="""// scratchCap is the scratch buffer's capacity. It is a variable, so the
// compiler cannot prove the buffer's size and must allocate it on the heap.
var scratchCap = 64""",
    buggy="""out := make([]string, 0, len(rows))
for _, row := range rows {
	buf := make([]byte, 0, scratchCap)
	for i, v := range row {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(v), 10)
	}
	out = append(out, string(buf))
}
return out""",
    solution="""out := make([]string, 0, len(rows))
buf := make([]byte, 0, scratchCap)
for _, row := range rows {
	buf = buf[:0]
	for i, v := range row {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(v), 10)
	}
	out = append(out, string(buf))
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	got := Render([][]int{{1, 2}, {3}, {}})
	want := []string{"1,2", "3", ""}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Render = %q, want %q", got, want)
	}
	if got := Render(nil); len(got) != 0 {
		t.Errorf("Render = %q, want empty", got)
	}
}

func TestRenderAllocationsScaleWithRowsNotIterations(t *testing.T) {
	rows := make([][]int, 64)
	for i := range rows {
		rows[i] = []int{i, i + 1, i + 2}
	}
	n := testing.AllocsPerRun(50, func() { _ = Render(rows) })
	// one result slice, one scratch buffer, one string per row
	if n > float64(len(rows))+4 {
		t.Errorf("Render made %v allocations for %d rows, want about %d: hoist the scratch buffer", n, len(rows), len(rows)+2)
	}
}
""",
    context="A report renderer allocates a scratch buffer for every row it formats. With a million rows that is a million buffers the collector has to chase, all of them identical and all of them dead immediately.",
    task=[
        "Render each row as its values joined by `,`.",
        "Fix the single bug so the scratch buffer is allocated once per call, not once per row.",
        "The output must not change.",
    ],
    examples=[
        ("Render([][]int{{1,2},{3}})", '["1,2" "3"]', None),
        ("Render([][]int{{}})", '[""]', "An empty row renders as an empty string."),
        ("64 rows", "about 66 allocations, not 130", None),
    ],
    topics=[
        ("Hoisting allocations", "State that is the same every iteration belongs outside the loop."),
        ("Resetting instead of reallocating", "`buf = buf[:0]` gives a clean buffer at no cost."),
        ("string(buf) copies", "The per-row string is a real allocation and is supposed to be."),
    ],
    hint="Only one line has to move. Something else has to be added where it used to be.",
    intuition="Allocating inside a loop is only wrong when the value does not need to be fresh. Here every row overwrites the buffer completely, so one buffer serves all of them — as long as its length is reset first.",
    approach=[
        "Move the `make` above the loop.",
        "Reset with `buf = buf[:0]` at the top of each iteration.",
    ],
    walkthrough="64 rows cost 64 scratch buffers before the fix and one after. The 64 output strings remain — those genuinely escape into the result.",
    pitfalls=[
        "Hoisting without resetting, which concatenates every row onto the last.",
        "Trying to avoid `string(buf)` too; that copy is what makes the result independent of the buffer.",
    ],
)

P(
    "middle",
    name="bytescompare",
    title="Compare Bytes Without Making A String",
    sig="func HasPrefix(b []byte, prefix string) bool",
    doc="""HasPrefix reports whether b begins with prefix.

Converting b to a string copies it. The comparison can be done on the
bytes that are already there.

Examples:

	HasPrefix([]byte("hello"), "he") => true""",
    solution="""if len(prefix) > len(b) {
	return false
}
for i := 0; i < len(prefix); i++ {
	if b[i] != prefix[i] {
		return false
	}
}
return true""",
    tests="""
import (
	"bytes"
	"testing"
)

var sink bool

func TestHasPrefix(t *testing.T) {
	cases := []struct {
		b      string
		prefix string
		want   bool
	}{
		{"hello", "he", true},
		{"hello", "hello", true},
		{"hello", "hello!", false},
		{"hello", "", true},
		{"", "x", false},
		{"hello", "ho", false},
	}
	for _, c := range cases {
		if got := HasPrefix([]byte(c.b), c.prefix); got != c.want {
			t.Errorf("HasPrefix(%q, %q) = %v, want %v", c.b, c.prefix, got, c.want)
		}
	}
}

func TestHasPrefixAllocatesNothing(t *testing.T) {
	b := bytes.Repeat([]byte("payload"), 128)
	if n := testing.AllocsPerRun(200, func() { sink = HasPrefix(b, "payload") }); n != 0 {
		t.Errorf("HasPrefix made %v allocations, want 0: do not convert b to a string", n)
	}
}
""",
    context="A protocol dispatcher checks the frame type with `strings.HasPrefix(string(frame), ...)`. Each check copies the whole frame to compare its first four bytes.",
    task=[
        "Report whether `b` starts with `prefix`.",
        "An empty prefix always matches; a prefix longer than `b` never does.",
        "Zero allocations, whatever the size of `b`.",
    ],
    examples=[
        ('HasPrefix([]byte("hello"), "he")', "true", None),
        ('HasPrefix([]byte("hello"), "hello!")', "false", "A longer prefix cannot match."),
        ('HasPrefix([]byte("hello"), "")', "true", None),
    ],
    topics=[
        ("string([]byte) copies", "The conversion allocates because strings are immutable and slices are not."),
        ("Indexing a string", "`prefix[i]` is a byte — no conversion needed to compare."),
        ("Length guards", "Check the length before indexing, not after."),
    ],
    hint="`b[i]` and `prefix[i]` are both bytes already.",
    intuition="`string(b)` has to copy, because the resulting string must be immutable while `b` is not. But comparing does not need a string at all — both sides index to bytes.",
    approach=[
        "Return false when `prefix` is longer than `b`.",
        "Compare byte by byte up to `len(prefix)`.",
        "Return true if none differed.",
    ],
    walkthrough="For a 896-byte frame and a 7-byte prefix, the conversion would copy 896 bytes; the byte loop reads at most 7 and allocates nothing.",
    pitfalls=[
        "`bytes.HasPrefix` is the real-world answer and also allocation-free — the point here is why.",
        "Indexing before the length check, which panics on a short `b`.",
    ],
)

P(
    "middle",
    name="valuereceiver",
    title="The Method That Copies Its Receiver",
    sig="func (c *Config) Timeouts() (read, write int)",
    doc="""Timeouts returns the read and write timeouts from c.

The receiver is a pointer because Config is large: a value receiver would
copy the whole struct on every call.

Examples:

	(&Config{Read: 1, Write: 2}).Timeouts() => 1, 2""",
    extra="""// Config is a deliberately large settings block.
type Config struct {
	Read  int
	Write int
	Pad   [512]byte
}""",
    solution="""return c.Read, c.Write""",
    tests="""
import "testing"

var sinkA, sinkB int

func TestTimeouts(t *testing.T) {
	c := &Config{Read: 5, Write: 9}
	if r, w := c.Timeouts(); r != 5 || w != 9 {
		t.Errorf("Timeouts = %d, %d, want 5, 9", r, w)
	}
	if r, w := (&Config{}).Timeouts(); r != 0 || w != 0 {
		t.Errorf("Timeouts = %d, %d, want 0, 0", r, w)
	}
}

func TestTimeoutsSeesLaterWrites(t *testing.T) {
	c := &Config{Read: 1}
	c.Read = 42
	if r, _ := c.Timeouts(); r != 42 {
		t.Errorf("read = %d, want 42: the receiver must be the caller's Config", r)
	}
}

func TestTimeoutsAllocatesNothing(t *testing.T) {
	c := &Config{Read: 1, Write: 2}
	if n := testing.AllocsPerRun(200, func() { sinkA, sinkB = c.Timeouts() }); n != 0 {
		t.Errorf("Timeouts made %v allocations, want 0", n)
	}
}
""",
    context="A config struct grows a 512-byte field. Every accessor on it still takes a value receiver, and a benchmark that never touched the allocator starts copying half a kilobyte per call.",
    task=[
        "Return the read and write timeouts from the receiver.",
        "The receiver must be the caller's `Config`, so later writes are visible.",
        "Zero allocations, zero large copies.",
    ],
    examples=[
        ("(&Config{Read:5, Write:9}).Timeouts()", "5, 9", None),
        ("c.Read = 42; c.Timeouts()", "42, ...", "The method sees the caller's struct."),
        ("(&Config{}).Timeouts()", "0, 0", None),
    ],
    topics=[
        ("Pointer vs value receivers", "A value receiver copies the struct on every call."),
        ("Copy cost scales with the struct", "Two words is free; 528 bytes is not."),
        ("Receiver consistency", "Mixing receiver kinds on one type is a readability trap."),
    ],
    hint="The signature is already given. The body is two field reads.",
    intuition="Receivers follow the same rules as parameters: a value receiver is a copy. That is fine for small types and quietly expensive for large ones — and it also means the method cannot see later writes.",
    approach=[
        "Return `c.Read` and `c.Write`.",
    ],
    walkthrough="With a pointer receiver the call passes one word. With a value receiver it would push 528 bytes onto the stack, and the copy would be a snapshot rather than a live view.",
    pitfalls=[
        "Dereferencing into a local `cfg := *c` — that reintroduces the copy you avoided.",
        "Nil receivers: a pointer receiver may be nil, so guard if the API allows it.",
    ],
)

# ---------------------------------------------------------------- senior -----

P(
    "senior",
    name="splitcount",
    title="Count Fields Without Splitting Anything",
    mode="bug",
    sig="func CountFields(line []byte, sep byte) (fields, size int)",
    doc="""CountFields returns how many sep-separated fields line holds and how
many bytes those fields occupy in total, excluding the separators.

Splitting builds a string and a slice of headers that are thrown away
immediately. One scan over the bytes answers both questions.

Examples:

	CountFields([]byte("ab,c"), ',') => 2, 3""",
    imports=['"strings"'],
    sol_imports=[],
    buggy="""if len(line) == 0 {
	return 0, 0
}
parts := strings.Split(string(line), string(sep))
n := 0
for _, p := range parts {
	n += len(p)
}
return len(parts), n""",
    solution="""if len(line) == 0 {
	return 0, 0
}
fields = 1
for _, b := range line {
	if b == sep {
		fields++
		continue
	}
	size++
}
return fields, size""",
    tests="""
import (
	"bytes"
	"testing"
)

var (
	sinkF int
	sinkS int
)

func TestCountFields(t *testing.T) {
	cases := []struct {
		in           string
		fields, size int
	}{
		{"ab,c", 2, 3},
		{"", 0, 0},
		{"abc", 1, 3},
		{",", 2, 0},
		{"a,,b", 3, 2},
		{",a", 2, 1},
	}
	for _, c := range cases {
		f, s := CountFields([]byte(c.in), ',')
		if f != c.fields || s != c.size {
			t.Errorf("CountFields(%q) = %d, %d, want %d, %d", c.in, f, s, c.fields, c.size)
		}
	}
}

func TestCountFieldsAllocatesNothing(t *testing.T) {
	line := bytes.Repeat([]byte("column,"), 256)
	line = line[:len(line)-1]
	n := testing.AllocsPerRun(100, func() { sinkF, sinkS = CountFields(line, ',') })
	if n != 0 {
		t.Errorf("CountFields made %v allocations, want 0: scan the bytes, do not split them", n)
	}
}

func BenchmarkCountFields(b *testing.B) {
	line := bytes.Repeat([]byte("column,"), 256)
	line = line[:len(line)-1]
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkF, sinkS = CountFields(line, ',')
	}
}
""",
    context="A CSV pre-pass reports each line's field count and payload size. It splits every line to count the pieces, then throws the pieces away — two allocations per line, on a file with a hundred million lines.",
    task=[
        "Return the number of sep-separated fields and the total byte size of those fields.",
        "An empty input is 0, 0; a line with no separator is one field.",
        "Fix the single bug so the function allocates nothing.",
    ],
    examples=[
        ('CountFields([]byte("ab,c"), \',\')', "2, 3", "Two fields, three payload bytes."),
        ('CountFields([]byte("a,,b"), \',\')', "3, 2", "The empty middle field still counts."),
        ('CountFields(nil, \',\')', "0, 0", None),
    ],
    topics=[
        ("Split allocates twice", "One copy for `string(line)` and one slice for the headers."),
        ("Counting is not collecting", "Both answers are running totals; the pieces never need to exist."),
        ("Separator arithmetic", "A line with n separators has n+1 fields."),
    ],
    hint="You are asked for two numbers. How many of the pieces do you actually need to hold?",
    intuition="Splitting materialises every field so you can measure them. But the field count is just the separator count plus one, and the payload size is the length minus the separators — both are running totals over a single scan.",
    approach=[
        "Return 0, 0 for an empty line.",
        "Start the field count at 1.",
        "Scan the bytes: a separator increments the field count, anything else increments the size.",
    ],
    walkthrough='"ab,c" has one separator, so two fields, and three non-separator bytes. The split version would allocate a 4-byte string and a 2-element header slice to reach the same numbers.',
    pitfalls=[
        "Starting the field count at 0, which is off by one for every non-empty line.",
        "Counting the separators in the size total.",
    ],
)

P(
    "senior",
    name="escapecallback",
    title="The Accumulator That Escaped Through A Callback",
    mode="bug",
    sig="func Sum(s []int) int64",
    doc="""Sum returns the total of s.

The helper Each is not inlinable, so any closure handed to it escapes —
and with it everything the closure captures.

Examples:

	Sum([]int{1, 2, 3}) => 6""",
    extra="""// Each calls f for every element of s.
//
// It is a package-level variable, so the compiler cannot see which function
// runs and must assume the callback it is given escapes.
var Each = func(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}""",
    buggy="""var total int64
Each(s, func(v int) { total += int64(v) })
return total""",
    solution="""var total int64
for _, v := range s {
	total += int64(v)
}
return total""",
    tests="""
import "testing"

var sink int64

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
	if got := Sum([]int{-5, 5}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}

func TestSumLarge(t *testing.T) {
	s := make([]int, 1000)
	var want int64
	for i := range s {
		s[i] = i
		want += int64(i)
	}
	if got := Sum(s); got != want {
		t.Errorf("Sum = %d, want %d", got, want)
	}
}

func TestSumAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	for i := range s {
		s[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { sink = Sum(s) }); n != 0 {
		t.Errorf("Sum made %v allocations, want 0: the accumulator is escaping through the callback", n)
	}
}
""",
    context="A tidy refactor replaces a loop with a callback helper. The behaviour is identical, the benchmark is 4x slower, and the allocation count went from zero to two.",
    task=[
        "Return the sum of `s` as an int64.",
        "Fix the single bug so the function allocates nothing.",
        "`Each` must stay as it is — it is a fixed part of the package.",
        "Do not change `Each` into a plain function to dodge the problem.",
    ],
    examples=[
        ("Sum([]int{1,2,3})", "6", None),
        ("Sum(nil)", "0", None),
        ("512 elements", "0 allocations", "The accumulator must stay in the frame."),
    ],
    topics=[
        ("Closures escape into non-inlined calls", "The callee may store the func value, so the compiler must assume it does."),
        ("Capture drags the variable along", "`total` is captured by reference, so it escapes with the closure."),
        ("Inlining is the enabling optimisation", "An inlined callback often costs nothing; a `//go:noinline` one never does."),
    ],
    hint="`Each` is a func variable, not a func. What does that force the compiler to assume about the closure it is given?",
    intuition="Escape analysis is per-function and conservative across call boundaries. Once a closure is passed to a function the compiler cannot see into, it has to assume the func value may be stored — so the closure and everything it captures go to the heap.",
    approach=[
        "Drop the callback and range over `s` directly.",
        "Accumulate into a local, which now provably does not escape.",
    ],
    walkthrough="The callback version allocates the closure object and moves `total` to the heap — two allocations, plus an indirect call per element. The direct loop keeps `total` in a register.",
    pitfalls=[
        "Assuming all closures allocate; one passed to a directly-called, inlinable function usually does not.",
        "Turning `Each` back into a plain function instead of removing the callback — the helper is part of the fixture.",
    ],
)

P(
    "senior",
    name="ifaceescape",
    title="The Buffer That An Interface Sent To The Heap",
    mode="bug",
    sig="func Checksum(vals []int) int",
    doc="""Checksum renders vals as decimal digits into a scratch buffer and
returns the sum of the bytes written.

Passing the scratch buffer to an interface makes it escape. Everything
here can stay in the frame.

Examples:

	Checksum([]int{1}) => 49""",
    imports=['"strconv"'],
    extra="""// sink accepts the rendering; it is a stand-in for a real writer.
type sink interface{ Write(p []byte) (int, error) }

type counter struct{ n int }

func (c *counter) Write(p []byte) (int, error) {
	for _, b := range p {
		c.n += int(b)
	}
	return len(p), nil
}

// newSink is a variable, so the compiler cannot devirtualise the interface
// value it produces.
var newSink = func(c *counter) sink { return c }""",
    buggy="""var buf [64]byte
b := buf[:0]
for _, v := range vals {
	b = strconv.AppendInt(b, int64(v), 10)
}
var c counter
w := newSink(&c)
w.Write(b)
return c.n""",
    solution="""var buf [64]byte
b := buf[:0]
for _, v := range vals {
	b = strconv.AppendInt(b, int64(v), 10)
}
n := 0
for _, x := range b {
	n += int(x)
}
return n""",
    tests="""
import "testing"

var sinkN int

func TestChecksum(t *testing.T) {
	if got := Checksum([]int{1}); got != int('1') {
		t.Errorf("Checksum = %d, want %d", got, int('1'))
	}
	if got := Checksum(nil); got != 0 {
		t.Errorf("Checksum = %d, want 0", got)
	}
	if got := Checksum([]int{12}); got != int('1')+int('2') {
		t.Errorf("Checksum = %d, want %d", got, int('1')+int('2'))
	}
	if got := Checksum([]int{-1}); got != int('-')+int('1') {
		t.Errorf("Checksum = %d, want %d", got, int('-')+int('1'))
	}
}

func TestChecksumAllocatesNothing(t *testing.T) {
	vals := []int{1, 22, 333, 4444}
	if n := testing.AllocsPerRun(200, func() { sinkN = Checksum(vals) }); n != 0 {
		t.Errorf("Checksum made %v allocations, want 0: the scratch buffer is escaping", n)
	}
}
""",
    context="A hot encoder keeps its scratch buffer as a fixed-size local, exactly as the style guide says. The allocation profile still shows the buffer on the heap, once per call.",
    task=[
        "Render each value as decimal digits into the local scratch buffer.",
        "Return the sum of the bytes written.",
        "Fix the single bug so the function allocates nothing.",
    ],
    examples=[
        ('Checksum([]int{1})', "49", "'1' is byte 49."),
        ("Checksum([]int{12})", "97", "'1' + '2'."),
        ("Checksum(nil)", "0", None),
    ],
    topics=[
        ("Interface arguments escape", "A value passed as an interface may be stored by the callee, so it goes to the heap."),
        ("The receiver escapes too", "`&c` behind an interface makes `c` heap-allocated as well."),
        ("Concrete calls keep the frame", "Direct code on a local slice keeps everything in the frame."),
    ],
    hint="The buffer is a fixed-size local. Which line takes its address out of the function's sight?",
    intuition="Escape analysis works on what the compiler can prove. Handing a slice to an interface method destroys the proof: the dynamic implementation is not known at compile time, so the argument must be assumed to escape.",
    approach=[
        "Render into the local buffer as before.",
        "Sum the bytes with a direct loop instead of routing them through the interface.",
    ],
    walkthrough="`w.Write(b)` makes both `buf` and `c` escape — two allocations per call. Summing `b` inline keeps the 64-byte array in the frame and reports 0 allocations.",
    pitfalls=[
        "Calling `c.Write(b)` on the concrete type instead — better, but the fixture still exists to be removed.",
        "Assuming a small local array can never escape; its address leaving the function is what matters, not its size.",
    ],
)

P(
    "senior",
    name="growonce",
    title="One Allocation For A Stream Of Unknown Length",
    sig="func Collect(r io.Reader, hint int) ([]byte, error)",
    doc="""Collect reads r to EOF and returns its bytes.

hint is the caller's estimate of the size. When it is accurate the whole
read must cost a single allocation instead of a chain of doublings.

Examples:

	Collect(strings.NewReader("abc"), 3) => []byte("abc"), nil""",
    imports=['"io"'],
    solution="""if hint < 0 {
	hint = 0
}
buf := make([]byte, 0, hint+1)
for {
	if len(buf) == cap(buf) {
		buf = append(buf, 0)[:len(buf)]
	}
	n, err := r.Read(buf[len(buf):cap(buf)])
	buf = buf[:len(buf)+n]
	if err == io.EOF {
		return buf, nil
	}
	if err != nil {
		return buf, err
	}
}""",
    tests="""
import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestCollect(t *testing.T) {
	got, err := Collect(strings.NewReader("abc"), 3)
	if err != nil || !bytes.Equal(got, []byte("abc")) {
		t.Errorf("Collect = %q, %v, want \\"abc\\", nil", got, err)
	}
	if got, err := Collect(strings.NewReader(""), 0); err != nil || len(got) != 0 {
		t.Errorf("Collect = %q, %v, want empty, nil", got, err)
	}
}

func TestCollectHandlesAWrongHint(t *testing.T) {
	in := strings.Repeat("x", 5000)
	got, err := Collect(strings.NewReader(in), 4)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != in {
		t.Errorf("len = %d, want %d: an underestimate must still read everything", len(got), len(in))
	}
	got, err = Collect(strings.NewReader("ab"), 9999)
	if err != nil || string(got) != "ab" {
		t.Errorf("Collect = %q, %v, want \\"ab\\", nil", got, err)
	}
}

type badReader struct{}

func (badReader) Read(p []byte) (int, error) { return 0, errors.New("boom") }

func TestCollectPropagatesErrors(t *testing.T) {
	if _, err := Collect(badReader{}, 4); err == nil {
		t.Error("want an error, got nil")
	}
}

func TestCollectWithAGoodHintAllocatesOnce(t *testing.T) {
	data := strings.Repeat("y", 8192)
	n := testing.AllocsPerRun(50, func() {
		_, _ = Collect(strings.NewReader(data), len(data))
	})
	if n > 2 {
		t.Errorf("Collect made %v allocations for an exact hint, want at most 2", n)
	}
}

func TestCollectDoesNotLoseBytesAtTheEOF(t *testing.T) {
	got, err := Collect(io.MultiReader(strings.NewReader("ab"), strings.NewReader("cd")), 2)
	if err != nil || string(got) != "abcd" {
		t.Errorf("Collect = %q, %v, want \\"abcd\\", nil", got, err)
	}
}
""",
    context="An RPC layer knows the payload size from the frame header but still calls `io.ReadAll`, which discovers the size by doubling — five allocations and four copies for a payload it could have sized exactly.",
    task=[
        "Read `r` to EOF and return its bytes.",
        "With an accurate `hint`, use at most two allocations.",
        "An under- or over-estimated hint must still return the correct bytes; read errors propagate.",
    ],
    examples=[
        ('Collect(strings.NewReader("abc"), 3)', '"abc", nil', None),
        ('Collect(strings.NewReader(strings.Repeat("x",5000)), 4)', "all 5000 bytes", "A bad hint costs speed, not correctness."),
        ("a failing reader", "the error", None),
    ],
    topics=[
        ("Read into spare capacity", "`r.Read(buf[len(buf):cap(buf)])` fills the room you already own."),
        ("append as the growth policy", "Appending one byte and reslicing back is the idiomatic \"grow if full\"."),
        ("Partial reads with EOF", "`Read` may return bytes and `io.EOF` together."),
    ],
    hint="Read into `buf[len(buf):cap(buf)]` and only grow when that window is empty.",
    intuition="`io.ReadAll` cannot know the size, so it doubles. When the caller does know, the whole doubling chain collapses into one `make` — and the loop only needs a fallback for when the hint was wrong.",
    approach=[
        "Allocate `make([]byte, 0, hint+1)`.",
        "Loop: if the buffer is full, grow it by appending one byte and reslicing back.",
        "Read into the spare capacity and extend the length by what was read.",
        "Return on `io.EOF`; propagate other errors.",
    ],
    walkthrough="An 8192-byte payload with an exact hint allocates once and reads until EOF. `io.ReadAll` on the same payload allocates at 512, 1024, 2048, 4096 and 8192 bytes, copying each time.",
    pitfalls=[
        "Reading into `buf[len(buf):]` — that window has length 0 when len < cap, so `Read` returns immediately and the loop spins.",
        "Discarding the bytes returned alongside `io.EOF`.",
    ],
)

P(
    "senior",
    name="noescapesort",
    title="Sort A Small Set Without Reaching For The Heap",
    sig="func Median3(a, b, c int) int",
    doc="""Median3 returns the middle value of a, b and c.

Sorting three ints through a general sorter would box them behind an
interface. Comparisons alone are enough.

Examples:

	Median3(3, 1, 2) => 2""",
    solution="""if a > b {
	a, b = b, a
}
if b > c {
	b = c
	if a > b {
		b = a
	}
}
return b""",
    tests="""
import (
	"sort"
	"testing"
)

var sink int

func TestMedian3(t *testing.T) {
	cases := [][4]int{
		{1, 2, 3, 2}, {3, 2, 1, 2}, {2, 1, 3, 2},
		{1, 1, 1, 1}, {5, 5, 1, 5}, {-1, 0, 1, 0},
	}
	for _, c := range cases {
		if got := Median3(c[0], c[1], c[2]); got != c[3] {
			t.Errorf("Median3(%d,%d,%d) = %d, want %d", c[0], c[1], c[2], got, c[3])
		}
	}
}

func TestMedian3MatchesSorting(t *testing.T) {
	for a := -3; a <= 3; a++ {
		for b := -3; b <= 3; b++ {
			for c := -3; c <= 3; c++ {
				s := []int{a, b, c}
				sort.Ints(s)
				if got := Median3(a, b, c); got != s[1] {
					t.Fatalf("Median3(%d,%d,%d) = %d, want %d", a, b, c, got, s[1])
				}
			}
		}
	}
}

func TestMedian3AllocatesNothing(t *testing.T) {
	if n := testing.AllocsPerRun(500, func() { sink = Median3(9, 2, 5) }); n != 0 {
		t.Errorf("Median3 made %v allocations, want 0: no slice, no interface", n)
	}
}
""",
    context="A quicksort's pivot selection calls `sort.Ints` on a three-element slice. The slice escapes into `sort.Interface`, so the hottest function in the sort allocates twice per partition.",
    task=[
        "Return the median of three ints.",
        "Handle duplicates and negatives correctly.",
        "Zero allocations — no slice, no interface, no sort package.",
    ],
    examples=[
        ("Median3(3, 1, 2)", "2", None),
        ("Median3(5, 5, 1)", "5", "Duplicates make the median one of the repeated values."),
        ("Median3(-1, 0, 1)", "0", None),
    ],
    topics=[
        ("sort.Interface boxes its argument", "A slice passed as an interface escapes to the heap."),
        ("Comparison networks", "Three values need at most three comparisons and no storage."),
        ("Hot-path specialisation", "A general tool is the wrong shape when n is fixed and tiny."),
    ],
    hint="Three comparisons and two swaps of local variables. Nothing else.",
    intuition="A general sorter must work for any length, so it takes an interface and the data escapes. When the length is three and known at compile time, the whole problem collapses into comparisons between registers.",
    approach=[
        "Order `a` and `b` so `a <= b`.",
        "If `b > c`, replace `b` with `c`, then lift it back to `a` if `a` is now larger.",
        "Return `b`.",
    ],
    walkthrough="For (3,1,2): the first swap gives a=1, b=3. b > c so b becomes 2; a (1) is not greater, so the answer is 2.",
    pitfalls=[
        "`a+b+c-min-max` overflows for large ints.",
        "Missing the equal cases; the median of (5,5,1) is 5, not 1.",
    ],
)

# ----------------------------------------------------------------- staff -----

P(
    "staff",
    name="hotpathgeneric",
    title="Type Parameters Instead Of Interface Boxing",
    sig="func Total[T ~int | ~int32 | ~int64](vals []T) int64",
    doc="""Total sums vals.

An `[]any` version would box every element. A type parameter gives the
compiler the concrete type, so nothing is boxed and nothing escapes.

Examples:

	Total([]int{1, 2, 3}) => 6""",
    solution="""var sum int64
for _, v := range vals {
	sum += int64(v)
}
return sum""",
    tests="""
import "testing"

var sink int64

type myInt int

func TestTotal(t *testing.T) {
	if got := Total([]int{1, 2, 3}); got != 6 {
		t.Errorf("Total = %d, want 6", got)
	}
	if got := Total([]int64{1 << 40, 1 << 40}); got != 1<<41 {
		t.Errorf("Total = %d, want %d", got, int64(1)<<41)
	}
	if got := Total([]int32{-5, 5}); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
	if got := Total[myInt](nil); got != 0 {
		t.Errorf("Total = %d, want 0", got)
	}
}

func TestTotalAcceptsNamedTypes(t *testing.T) {
	if got := Total([]myInt{2, 3}); got != 5 {
		t.Errorf("Total = %d, want 5", got)
	}
}

func TestTotalAllocatesNothing(t *testing.T) {
	vals := make([]int, 1024)
	for i := range vals {
		vals[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { sink = Total(vals) }); n != 0 {
		t.Errorf("Total made %v allocations, want 0", n)
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
    context="A metrics aggregator takes `[]any` so it can serve every integer width. Boxing the values costs one allocation per element, and the aggregator is called once per scrape per series.",
    task=[
        "Sum `vals` and return the total as an int64.",
        "Accept `int`, `int32`, `int64` and any named type based on them.",
        "Zero allocations — nothing may be boxed.",
    ],
    examples=[
        ("Total([]int{1,2,3})", "6", None),
        ("Total([]int64{1<<40, 1<<40})", "2199023255552", "The accumulator must be wide enough."),
        ("Total([]myInt{2,3})", "5", "`~int` admits named types."),
    ],
    topics=[
        ("Type parameters vs interfaces", "A generic call is compiled against the concrete type; an interface call boxes."),
        ("Approximation constraints", "`~int` covers every type whose underlying type is int."),
        ("Accumulator width", "Summing int32 into int32 overflows; int64 does not."),
        ("Escape analysis with generics", "No boxing means no escape, so the loop stays in registers."),
    ],
    hint="The constraint is already written. The body is the obvious loop — the point is what the signature bought you.",
    intuition="`any` erases the type, so every value needs a heap word and a type word. A type parameter keeps the type, so the compiler emits code for the real element width and never touches the allocator.",
    approach=[
        "Declare an int64 accumulator.",
        "Range the values, converting each to int64 before adding.",
        "Return the total.",
    ],
    walkthrough="Summing 1024 ints via `[]any` allocates about 1024 times, since values above 255 do not hit the runtime's small-integer cache. The generic version allocates zero and is several times faster.",
    pitfalls=[
        "Accumulating in `T` — an `[]int32` of large values overflows before the conversion.",
        "Adding `~uint` to the constraint; the conversion to int64 would then be lossy.",
    ],
)

P(
    "staff",
    name="atomicsnapshot",
    title="Publish A Snapshot Without Tearing It",
    sig="func (s *Store) Set(c Config)",
    doc="""Set publishes c as the current configuration.

Readers must see either the old snapshot or the new one, never a mix.
The snapshot escapes by construction — it outlives the call and is shared
with every reader.

Examples:

	s.Set(Config{Retries: 3}); s.Get().Retries => 3""",
    imports=['"sync/atomic"'],
    extra="""// Config is one immutable settings snapshot.
type Config struct {
	Retries int
	Timeout int
}

// Store holds the current Config for concurrent readers.
type Store struct {
	v atomic.Pointer[Config]
}

// Get returns the current snapshot, or the zero Config if none is set.
func (s *Store) Get() Config {
	if p := s.v.Load(); p != nil {
		return *p
	}
	return Config{}
}""",
    solution="""cp := c
s.v.Store(&cp)""",
    tests="""
import (
	"sync"
	"testing"
)

func TestSetThenGet(t *testing.T) {
	var s Store
	if got := s.Get(); got != (Config{}) {
		t.Errorf("Get = %v, want the zero Config", got)
	}
	s.Set(Config{Retries: 3, Timeout: 10})
	if got := s.Get(); got != (Config{Retries: 3, Timeout: 10}) {
		t.Errorf("Get = %v, want {3 10}", got)
	}
	s.Set(Config{Retries: 1})
	if got := s.Get(); got.Retries != 1 || got.Timeout != 0 {
		t.Errorf("Get = %v, want {1 0}", got)
	}
}

func TestSetDoesNotAliasTheCaller(t *testing.T) {
	var s Store
	c := Config{Retries: 1, Timeout: 2}
	s.Set(c)
	c.Retries = 99
	if got := s.Get(); got.Retries != 1 {
		t.Errorf("Retries = %d, want 1: the store aliases the caller's variable", got.Retries)
	}
}

func TestSnapshotsNeverTear(t *testing.T) {
	var s Store
	s.Set(Config{Retries: 1, Timeout: 10})
	var wg sync.WaitGroup
	stop := make(chan struct{})
	for w := 0; w < 4; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
				}
				c := s.Get()
				if c.Timeout != c.Retries*10 {
					panic("torn snapshot")
				}
			}
		}()
	}
	for i := 1; i <= 2000; i++ {
		s.Set(Config{Retries: i, Timeout: i * 10})
	}
	close(stop)
	wg.Wait()
}
""",
    context="A config reloader updates two fields of a shared struct in place. Under load a request occasionally sees the new retry count with the old timeout, and the combination is not one anybody configured.",
    task=[
        "Publish `c` as the current snapshot.",
        "Readers must never observe a half-updated Config.",
        "The store must not alias the caller's variable — later writes to it are invisible.",
    ],
    examples=[
        ("s.Set(Config{Retries:3, Timeout:10}); s.Get()", "{3 10}", None),
        ("c := Config{...}; s.Set(c); c.Retries = 99; s.Get()", "the value as of Set", "The store holds its own copy."),
        ("Get before any Set", "the zero Config", None),
    ],
    topics=[
        ("Atomic pointer swap", "Publishing one word replaces a whole struct indivisibly."),
        ("Immutable snapshots", "Never mutate a published value; publish a new one."),
        ("Escape by design", "The snapshot must be heap-allocated — it outlives the call and is shared."),
        ("Happens-before", "`Store`/`Load` on an atomic gives readers the fully written struct."),
    ],
    hint="Two fields cannot be written atomically. One pointer can.",
    intuition="Tearing comes from updating a shared value field by field. Replace the value instead: build the new snapshot privately, then swap one pointer. Readers either load the old address or the new one, and both point at a complete struct.",
    approach=[
        "Copy the parameter into a local so the store owns it.",
        "`s.v.Store(&cp)` to publish the address.",
    ],
    walkthrough="2000 concurrent updates with four readers spinning: every `Get` dereferences one of the 2001 complete snapshots. Writing `s.cfg.Retries` and `s.cfg.Timeout` separately would let a reader land between the two stores.",
    pitfalls=[
        "`s.v.Store(&c)` publishes the parameter's address — legal, but then the caller's later writes are visible to readers.",
        "Mutating a snapshot after publishing it; the whole scheme depends on published values being immutable.",
    ],
)

P(
    "staff",
    name="perworkerbuf",
    title="Give Every Worker Its Own Frame",
    sig="func RenderAll(rows [][]int) []string",
    doc="""RenderAll renders each row concurrently as comma-separated decimals
and returns the results in input order.

Each goroutine's scratch buffer must be a local that does not escape:
one shared buffer is a race, and one heap buffer per row is garbage.

Examples:

	RenderAll([][]int{{1, 2}}) => []string{"1,2"}""",
    imports=['"strconv"', '"sync"'],
    solution="""out := make([]string, len(rows))
var wg sync.WaitGroup
wg.Add(len(rows))
for i, row := range rows {
	go func(i int, row []int) {
		defer wg.Done()
		var scratch [64]byte
		buf := scratch[:0]
		for j, v := range row {
			if j > 0 {
				buf = append(buf, ',')
			}
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		out[i] = string(buf)
	}(i, row)
}
wg.Wait()
return out""",
    tests="""
import (
	"strconv"
	"strings"
	"testing"
)

func TestRenderAll(t *testing.T) {
	got := RenderAll([][]int{{1, 2}, {3}, {}})
	want := []string{"1,2", "3", ""}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
	if got := RenderAll(nil); len(got) != 0 {
		t.Errorf("RenderAll = %q, want empty", got)
	}
}

func TestRenderAllUnderConcurrency(t *testing.T) {
	const n = 128
	rows := make([][]int, n)
	for i := range rows {
		rows[i] = []int{i, i * 2, i * 3}
	}
	for round := 0; round < 20; round++ {
		got := RenderAll(rows)
		for i := range rows {
			want := strings.Join([]string{
				strconv.Itoa(i), strconv.Itoa(i * 2), strconv.Itoa(i * 3),
			}, ",")
			if got[i] != want {
				t.Fatalf("round %d: got[%d] = %q, want %q", round, i, got[i], want)
			}
		}
	}
}

func TestRenderAllOrderIsPreserved(t *testing.T) {
	rows := [][]int{{9}, {8}, {7}, {6}}
	got := RenderAll(rows)
	for i, want := range []string{"9", "8", "7", "6"} {
		if got[i] != want {
			t.Fatalf("got[%d] = %q, want %q", i, got[i], want)
		}
	}
}
""",
    context="A fan-out renderer was fixed for a data race by moving its scratch buffer inside the goroutine. The race is gone and the allocation count per row doubled — the fix reached for `make` when the frame would have done.",
    task=[
        "Render every row concurrently as its values joined by `,`.",
        "Return the results in input order.",
        "Each goroutine's scratch must be a non-escaping local; no shared buffer, no per-row heap buffer.",
    ],
    examples=[
        ("RenderAll([][]int{{1,2},{3}})", '["1,2" "3"]', None),
        ("128 rows, 20 rounds", "every result correct", "No goroutine touches another's scratch."),
        ("RenderAll(nil)", "[]", None),
    ],
    topics=[
        ("Goroutine stacks", "Each goroutine has its own frame, so a local array is per-worker by construction."),
        ("Escape analysis inside a goroutine", "A fixed-size array that only feeds `string(buf)` stays on the stack."),
        ("Disjoint slot writes", "`out[i]` from goroutine i needs no synchronisation."),
        ("Loop variables as parameters", "Passing `i` and `row` in keeps each goroutine's inputs its own."),
    ],
    hint="A goroutine has a stack too. What lives there for free?",
    intuition="Per-goroutine state does not have to be heap state. A fixed-size array declared inside the goroutine lives in that goroutine's frame — private without a `make`, and free without a pool.",
    approach=[
        "Preallocate the result slice to `len(rows)`.",
        "Start one goroutine per row, passing `i` and `row` as parameters.",
        "Declare `var scratch [64]byte` inside the goroutine and build into `scratch[:0]`.",
        "Write `out[i]`, then `wg.Wait()`.",
    ],
    walkthrough="Each goroutine renders into its own 64-byte frame array; only `string(buf)` allocates, which is the result the caller keeps. With 128 rows that is 129 allocations instead of 257.",
    pitfalls=[
        "Rows longer than the scratch array make `append` allocate — correct, just no longer free.",
        "Hoisting the array above the loop, which is the race the fix was for.",
    ],
)

P(
    "staff",
    name="parseints",
    title="Parse Without Making A Single String",
    sig="func ParseInts(line []byte, sep byte) (int64, int, error)",
    doc="""ParseInts sums the decimal integers in line, which are separated by
sep, and returns the total, the count parsed, and any error.

No part of line may be converted to a string: the parse works on the
bytes and allocates nothing.

Examples:

	ParseInts([]byte("1,2,3"), ',') => 6, 3, nil""",
    imports=['"errors"'],
    extra="""// ErrSyntax is returned for a field that is not a decimal integer.
var ErrSyntax = errors.New("invalid integer")""",
    solution="""if len(line) == 0 {
	return 0, 0, nil
}
var total int64
count := 0
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

func TestParseInts(t *testing.T) {
	if total, n, err := ParseInts([]byte("1,2,3"), ','); err != nil || total != 6 || n != 3 {
		t.Errorf("ParseInts = %d, %d, %v, want 6, 3, nil", total, n, err)
	}
	if total, n, err := ParseInts([]byte("-4,+6"), ','); err != nil || total != 2 || n != 2 {
		t.Errorf("ParseInts = %d, %d, %v, want 2, 2, nil", total, n, err)
	}
	if total, n, err := ParseInts(nil, ','); err != nil || total != 0 || n != 0 {
		t.Errorf("ParseInts = %d, %d, %v, want 0, 0, nil", total, n, err)
	}
	if _, _, err := ParseInts([]byte("1,,2"), ','); !errors.Is(err, ErrSyntax) {
		t.Errorf("err = %v, want ErrSyntax", err)
	}
	if _, _, err := ParseInts([]byte("1,x"), ','); !errors.Is(err, ErrSyntax) {
		t.Errorf("err = %v, want ErrSyntax", err)
	}
	if _, _, err := ParseInts([]byte("-"), ','); !errors.Is(err, ErrSyntax) {
		t.Errorf("err = %v, want ErrSyntax", err)
	}
}

func TestParseIntsSingleField(t *testing.T) {
	if total, n, err := ParseInts([]byte("42"), ','); err != nil || total != 42 || n != 1 {
		t.Errorf("ParseInts = %d, %d, %v, want 42, 1, nil", total, n, err)
	}
}

func TestParseIntsAllocatesNothing(t *testing.T) {
	line := bytes.Repeat([]byte("12345,"), 200)
	line = line[:len(line)-1]
	n := testing.AllocsPerRun(100, func() {
		sinkT, sinkC, _ = ParseInts(line, ',')
	})
	if n != 0 {
		t.Errorf("ParseInts made %v allocations, want 0: parse the bytes in place", n)
	}
}

func BenchmarkParseInts(b *testing.B) {
	line := bytes.Repeat([]byte("12345,"), 200)
	line = line[:len(line)-1]
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkT, sinkC, _ = ParseInts(line, ',')
	}
}
""",
    context="An ingest hot path splits each line with `strings.Split(string(line), \",\")` and calls `strconv.Atoi` on every field. At a million lines a second the allocator is the bottleneck, not the parsing.",
    task=[
        "Sum the decimal integers in `line`, separated by `sep`.",
        "Return the total, how many fields were parsed, and `ErrSyntax` for anything that is not a decimal integer.",
        "An empty input is 0, 0, nil; an empty field is a syntax error.",
        "Zero allocations — no string conversion, no split slice.",
    ],
    examples=[
        ('ParseInts([]byte("1,2,3"), \',\')', "6, 3, nil", None),
        ('ParseInts([]byte("-4,+6"), \',\')', "2, 2, nil", "Signs are accepted."),
        ('ParseInts([]byte("1,,2"), \',\')', "ErrSyntax", "An empty field is invalid."),
    ],
    topics=[
        ("Parsing in place", "Digits can be folded into an accumulator straight from the bytes."),
        ("Split allocates twice", "A conversion for the string and a slice for the pieces."),
        ("Sentinel errors on the hot path", "A package-level error keeps the failure path allocation-free too."),
        ("Boundary handling", "Running the loop to `len(line)` inclusive closes the final field."),
    ],
    hint="`v = v*10 + int64(c-'0')` is the whole parser. The rest is where the fields start and end.",
    intuition="Splitting and converting exist to make the data convenient, not to make it parseable. Digits are already digits: one pass with an accumulator and a field start index does the whole job in the caller's memory.",
    approach=[
        "Walk `i` from 0 to `len(line)` inclusive, treating the end as a separator.",
        "For each field, handle an optional sign, then fold digits into an accumulator.",
        "Reject empty fields, a lone sign, and any non-digit byte with `ErrSyntax`.",
        "Add to the total and count the field.",
    ],
    walkthrough='"1,2,3" closes fields at indices 1, 3 and 5 (the virtual separator at the end). Each field folds to 1, 2 and 3, giving 6 over three fields, with no allocation at any point.',
    pitfalls=[
        "Stopping the loop at `len(line)-1`, which silently drops the last field.",
        "Returning a formatted error, which reintroduces an allocation on the failure path.",
        "Accepting an empty field as zero — the spec says it is a syntax error.",
    ],
)

P(
    "staff",
    name="workerpartials",
    title="Aggregate In Parallel Without Sharing A Word",
    sig="func Histogram(data []int, buckets, workers int) []int64",
    doc="""Histogram counts data into buckets bins by value modulo buckets, using
workers goroutines over disjoint chunks.

Workers must not share a counter: each accumulates privately and the
results are folded once, after the join.

Examples:

	Histogram([]int{0, 1, 2, 3}, 2, 2) => []int64{2, 2}""",
    imports=['"sync"'],
    solution="""if buckets < 1 {
	return nil
}
out := make([]int64, buckets)
if len(data) == 0 {
	return out
}
if workers < 1 {
	workers = 1
}
if workers > len(data) {
	workers = len(data)
}
partials := make([][]int64, workers)
size := (len(data) + workers - 1) / workers
var wg sync.WaitGroup
wg.Add(workers)
for w := 0; w < workers; w++ {
	start := w * size
	end := start + size
	if start > len(data) {
		start = len(data)
	}
	if end > len(data) {
		end = len(data)
	}
	go func(w int, part []int) {
		defer wg.Done()
		local := make([]int64, buckets)
		for _, v := range part {
			b := v % buckets
			if b < 0 {
				b += buckets
			}
			local[b]++
		}
		partials[w] = local
	}(w, data[start:end])
}
wg.Wait()
for _, p := range partials {
	for i, c := range p {
		out[i] += c
	}
}
return out""",
    tests="""
import (
	"reflect"
	"testing"
)

func TestHistogramSmall(t *testing.T) {
	if got := Histogram([]int{0, 1, 2, 3}, 2, 2); !reflect.DeepEqual(got, []int64{2, 2}) {
		t.Errorf("Histogram = %v, want [2 2]", got)
	}
	if got := Histogram(nil, 3, 4); !reflect.DeepEqual(got, []int64{0, 0, 0}) {
		t.Errorf("Histogram = %v, want [0 0 0]", got)
	}
	if got := Histogram([]int{1}, 0, 2); got != nil {
		t.Errorf("Histogram = %v, want nil", got)
	}
}

func TestHistogramNegativeValues(t *testing.T) {
	got := Histogram([]int{-1, -2, -3}, 3, 2)
	if !reflect.DeepEqual(got, []int64{1, 1, 1}) {
		t.Errorf("Histogram = %v, want [1 1 1]: negative values must land in range", got)
	}
}

func TestHistogramMatchesSerial(t *testing.T) {
	const buckets = 13
	data := make([]int, 100003)
	want := make([]int64, buckets)
	for i := range data {
		data[i] = i * 7
		want[data[i]%buckets]++
	}
	for _, w := range []int{1, 2, 3, 8, 32, 1 << 20} {
		got := Histogram(data, buckets, w)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("workers=%d: Histogram = %v, want %v", w, got, want)
		}
	}
}

func TestHistogramRepeatable(t *testing.T) {
	data := make([]int, 50000)
	for i := range data {
		data[i] = i
	}
	first := Histogram(data, 7, 8)
	for i := 0; i < 20; i++ {
		if got := Histogram(data, 7, 8); !reflect.DeepEqual(got, first) {
			t.Fatalf("round %d: %v, want %v: the workers are sharing counters", i, got, first)
		}
	}
}
""",
    context="A parallel histogram increments a shared bucket slice from every worker. It is fast, wrong, and wrong differently on every run.",
    task=[
        "Count `data` into `buckets` bins by value modulo `buckets`, using `workers` goroutines over disjoint chunks.",
        "Negative values must land in a valid bin.",
        "Each worker accumulates privately; fold the partials after the join.",
        "`buckets < 1` returns nil; `workers` outside `[1, len(data)]` is clamped.",
    ],
    examples=[
        ("Histogram([]int{0,1,2,3}, 2, 2)", "[2 2]", None),
        ("Histogram([]int{-1,-2,-3}, 3, 2)", "[1 1 1]", "Go's `%` keeps the sign of the dividend."),
        ("Histogram([]int{1}, 0, 2)", "<nil>", None),
    ],
    topics=[
        ("Private accumulation", "Per-worker state removes contention and the need for atomics entirely."),
        ("Fold after join", "`wg.Wait()` is the happens-before edge that makes the partials safe to read."),
        ("Chunk arithmetic", "Ceiling-divided chunks, clamped, cover the input exactly once."),
        ("Go's modulo sign", "`-1 % 3` is -1, not 2 — the bin must be corrected."),
    ],
    hint="Shared counters need atomics; private counters need nothing. Which is cheaper to fold once at the end?",
    intuition="Contention is a property of sharing, not of parallelism. Give each worker its own bucket array and the inner loop becomes an ordinary single-threaded loop; the only synchronisation left is the join.",
    approach=[
        "Validate `buckets`, clamp `workers`, and handle the empty input.",
        "Split `data` into clamped chunks and give each goroutine a view plus its own `local` bucket array.",
        "Correct negative bins by adding `buckets`.",
        "After `Wait`, sum the partials into the result.",
    ],
    walkthrough="With 100003 values, 13 buckets and 8 workers, each worker fills its own 13-element array; the fold afterwards touches 104 int64s once. A shared array would have needed 100003 atomic increments.",
    pitfalls=[
        "`v % buckets` alone puts negative values at a negative index and panics.",
        "Folding before `Wait`, which reads partials that are still being written.",
        "Writing into the shared `out` from the workers, which is the original bug.",
    ],
)
