// Package copyemptydst clones a slice. A planted bug sizes the destination
// with length 0, so copy copies nothing.
package copyemptydst

// Clone returns an independent copy of xs.
func Clone(xs []int) []int {
	// CHANGE CODE BELOW THIS LINE
	dst := make([]int, 0, len(xs))
	// CHANGE CODE ABOVE THIS LINE
	copy(dst, xs)
	return dst
}
