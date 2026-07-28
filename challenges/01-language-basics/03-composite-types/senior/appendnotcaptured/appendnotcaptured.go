// Package appendnotcaptured builds a slice. A planted bug ignores append's
// return value.
package appendnotcaptured

// Doubled returns a slice with each element of xs doubled.
func Doubled(xs []int) []int {
	out := make([]int, 0, len(xs))
	for _, x := range xs {
		// CHANGE CODE BELOW THIS LINE
		_ = append(out, x*2)
		// CHANGE CODE ABOVE THIS LINE
	}
	return out
}
