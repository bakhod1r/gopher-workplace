// Package appendignored builds a slice of squares. A planted bug ignores the
// value append returns, so growth past capacity is lost.
package appendignored

// Squares returns the squares 1..n.
func Squares(n int) []int {
	out := make([]int, 0)
	for i := 1; i <= n; i++ {
		// CHANGE CODE BELOW THIS LINE
		_ = append(out, i*i)
		// CHANGE CODE ABOVE THIS LINE
	}
	return out
}
