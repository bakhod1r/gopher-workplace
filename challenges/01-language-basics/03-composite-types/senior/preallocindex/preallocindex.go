// Package preallocindex doubles a slice. A planted bug indexes into a length-0
// slice instead of appending.
package preallocindex

// Doubled returns a slice with each element of xs doubled.
func Doubled(xs []int) []int {
	// CHANGE CODE BELOW THIS LINE
	out := make([]int, 0, len(xs))
	for i, x := range xs {
		out[i] = x * 2
	}
	// CHANGE CODE ABOVE THIS LINE
	return out
}
