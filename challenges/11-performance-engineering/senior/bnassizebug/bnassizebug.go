// Package bnassizebug — Gopher Workplace challenge.
package bnassizebug

// Work returns how many units of work a benchmark body performs when the
// harness asks for n iterations over a fixed input of size items: n
// iterations, each processing the whole input.
//
// The classic mistake is to use n as the input size instead of the iteration
// count, which makes the "per operation" cost grow with n and the benchmark
// meaningless.
//
// Examples:
//
//	Work(3, 10) => 30
func Work(n, size int) int {
	if n <= 0 || size <= 0 {
		return 0
	}
	// CHANGE CODE BELOW THIS LINE
	return n * n
	// CHANGE CODE ABOVE THIS LINE
}

// PerOp returns the work performed per iteration, which must not depend on n.
//
// Examples:
//
//	PerOp(3, 10) => 10
func PerOp(n, size int) int {
	if n <= 0 {
		return 0
	}
	return Work(n, size) / n
}
