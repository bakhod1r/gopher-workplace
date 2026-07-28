// Package sliceleak returns a small head of a large slice. A planted bug returns
// a sub-slice that retains the entire backing array (a memory leak).
package sliceleak

// Head3 returns the first 3 elements as an independent slice that does NOT keep
// the (potentially huge) source backing array alive.
func Head3(xs []int) []int {
	// CHANGE CODE BELOW THIS LINE
	return xs[:3]
	// CHANGE CODE ABOVE THIS LINE
}
