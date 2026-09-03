// Package timerbeforesetupbug — Gopher Workplace challenge.
package timerbeforesetupbug

// Measured returns the nanoseconds a benchmark should report for n
// iterations: the per-iteration work only. The one-time setup happens before
// the loop and must not be charged to any iteration.
//
// Examples:
//
//	Measured(1000, 7, 3) => 21
func Measured(setupNS, workNS, n int64) int64 {
	if n <= 0 {
		return 0
	}
	// CHANGE CODE BELOW THIS LINE
	return setupNS + workNS*n
	// CHANGE CODE ABOVE THIS LINE
}

// PerOp divides the measured total by the iteration count, the ns/op the tool
// prints.
//
// Examples:
//
//	PerOp(1000, 7, 3) => 7
func PerOp(setupNS, workNS, n int64) int64 {
	if n <= 0 {
		return 0
	}
	return Measured(setupNS, workNS, n) / n
}
