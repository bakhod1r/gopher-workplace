// Package capretain extracts a small head from a large slice for long-term
// keeping. A planted bug returns a plain re-slice, which retains the ENTIRE
// backing array (a memory leak) AND lets an append into spare capacity clobber
// data beyond the head. The test detects the clobber.
package capretain

// Head returns the first k elements as an independent slice whose appends cannot
// reach the original tail.
func Head(xs []int, k int) []int {
	// CHANGE CODE BELOW THIS LINE
	return xs[:k]
	// CHANGE CODE ABOVE THIS LINE
}
