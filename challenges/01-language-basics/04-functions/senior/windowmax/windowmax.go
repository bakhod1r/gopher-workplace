// Package windowmax returns the maximum sum of any k consecutive elements. A
// planted bug in the loop bound reads one element past the window, corrupting
// the result (and risking an out-of-range on the last window).
package windowmax

// MaxWindow returns the largest sum of k consecutive elements of xs.
// Precondition: 0 < k <= len(xs).
func MaxWindow(xs []int, k int) int {
	best := 0
	for i := 0; i < k; i++ {
		best += xs[i]
	}
	cur := best
	for i := k; i < len(xs); i++ {
		// CHANGE CODE BELOW THIS LINE
		cur += xs[i] - xs[i-k+1]
		// CHANGE CODE ABOVE THIS LINE
		if cur > best {
			best = cur
		}
	}
	return best
}
