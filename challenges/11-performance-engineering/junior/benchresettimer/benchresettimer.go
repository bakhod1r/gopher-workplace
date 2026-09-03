// Package benchresettimer — Gopher Workplace challenge.
package benchresettimer

// Measured returns the nanoseconds a benchmark should attribute to the work
// itself: setupNS is spent once before the loop and must be excluded, and
// workNS is spent on each of the n iterations. A non-positive n measures 0.
//
// Examples:
//
//	Measured(1000, 7, 3) => 21
func Measured(setupNS, workNS int64, n int64) int64 {
	panic("not implemented")
}
