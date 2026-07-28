// Package sharedcell returns independent counters. A planted bug captures a
// single shared variable across all counters, so they interfere.
package sharedcell

// Counters returns n independent counter closures; each returns 1,2,3,... on its
// own successive calls, independently of the others.
func Counters(n int) []func() int {
	out := make([]func() int, 0, n)
	// CHANGE CODE BELOW THIS LINE
	c := 0
	for i := 0; i < n; i++ {
		out = append(out, func() int { c++; return c })
	}
	// CHANGE CODE ABOVE THIS LINE
	return out
}
