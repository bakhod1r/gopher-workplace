// Package tracespanoverlap — Gopher Workplace challenge.
package tracespanoverlap

// Span is a half-open time interval [Start, End) in nanoseconds.
type Span struct {
	Start int64
	End   int64
}

// Intersect returns the overlap of two spans and whether they overlap at all.
// Touching spans — one ending exactly where the other starts — do not overlap,
// because the intervals are half-open.
//
// Examples:
//
//	Intersect(Span{0, 10}, Span{5, 20}) => Span{5, 10}, true
func Intersect(a, b Span) (Span, bool) {
	panic("not implemented")
}

// Concurrency returns the maximum number of spans active at any instant — the
// peak parallelism a trace shows. Empty spans are ignored.
//
// Examples:
//
//	Concurrency([{0,10},{5,20},{6,7}]) => 3
func Concurrency(spans []Span) int {
	panic("not implemented")
}

// BusiestAt returns the timestamp at which that peak is first reached, and
// false when there are no valid spans.
//
// Examples:
//
//	BusiestAt([{0,10},{5,20}]) => 5, true
func BusiestAt(spans []Span) (int64, bool) {
	panic("not implemented")
}
