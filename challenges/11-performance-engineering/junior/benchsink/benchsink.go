// Package benchsink — Gopher Workplace challenge.
package benchsink

// Sink holds the most recent value handed to Consume. Benchmarks assign
// their result here so the compiler cannot delete the work as dead code.
var Sink int

// Consume stores v in Sink and returns the value it replaced.
//
// Examples:
//
//	Consume(7) => previous Sink, and Sink == 7 afterwards
func Consume(v int) int {
	panic("not implemented")
}

// SumTo returns 0+1+...+(n-1); a non-positive n sums to 0.
//
// Examples:
//
//	SumTo(4) => 6
func SumTo(n int) int {
	panic("not implemented")
}
