// Package copyshort duplicates a slice. A planted bug allocates the destination
// with length 0, so copy copies nothing (copy is bounded by the shorter slice).
package copyshort

// Clone returns an independent copy of xs.
func Clone(xs []int) []int {
	// CHANGE CODE BELOW THIS LINE
	dst := make([]int, 0, len(xs))
	// CHANGE CODE ABOVE THIS LINE
	copy(dst, xs)
	return dst
}
