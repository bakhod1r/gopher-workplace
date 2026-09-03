// Package tracespanmerge — Gopher Workplace challenge.
package tracespanmerge

// Span is a half-open time interval [Start, End) in nanoseconds.
type Span struct {
	Start int64
	End   int64
}

// Merge combines overlapping and touching spans into the smallest set that
// covers the same time, ordered by start. Spans with End at or before Start
// are dropped, and the input is not modified.
//
// Examples:
//
//	Merge([{0,10},{5,20}]) => [{0,20}]
func Merge(spans []Span) []Span {
	panic("not implemented")
}

// Covered returns the total time the merged spans occupy — the wall-clock
// time something was happening, which is not the sum of the span durations
// when they overlap.
//
// Examples:
//
//	Covered([{0,10},{5,20}]) => 20
func Covered(spans []Span) int64 {
	panic("not implemented")
}
