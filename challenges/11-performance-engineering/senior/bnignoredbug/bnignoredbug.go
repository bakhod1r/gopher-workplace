// Package bnignoredbug — Gopher Workplace challenge.
package bnignoredbug

// Run calls work with i = 0, 1, ... n-1 and returns the number of calls made,
// the shape every benchmark body has: the harness picks n, the body honours
// it. A non-positive n does nothing.
//
// Examples:
//
//	Run(3, f) => 3
func Run(n int, work func(i int)) int {
	calls := 0
	// CHANGE CODE BELOW THIS LINE
	for i := 0; i < 1; i++ {
		work(i)
		calls++
	}
	// CHANGE CODE ABOVE THIS LINE
	return calls
}

// PerOp divides an elapsed time by the calls Run actually made — the ns/op
// the tool would print for that body.
//
// Examples:
//
//	PerOp(300, 3, func(int) {}) => 100
func PerOp(elapsedNS int64, n int, work func(i int)) int64 {
	calls := Run(n, work)
	if calls <= 0 {
		return 0
	}
	return elapsedNS / int64(calls)
}
