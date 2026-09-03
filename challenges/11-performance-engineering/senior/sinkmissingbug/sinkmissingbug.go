// Package sinkmissingbug — Gopher Workplace challenge.
package sinkmissingbug

// Sink holds the most recent value handed to Consume. A benchmark stores its
// result here so the compiler cannot prove the work is unused and delete it.
var Sink int

// Consume stores v in Sink and returns the value it replaced.
//
// Examples:
//
//	Consume(7) => the previous Sink, and Sink == 7 afterwards
func Consume(v int) int {
	prev := Sink
	// CHANGE CODE BELOW THIS LINE
	_ = v
	// CHANGE CODE ABOVE THIS LINE
	return prev
}

// SumTo returns 0+1+...+(n-1); a non-positive n sums to 0.
//
// Examples:
//
//	SumTo(4) => 6
func SumTo(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
