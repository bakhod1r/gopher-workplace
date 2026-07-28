// Package pairbug sums adjacent differences. A planted bug loops i to len(xs)-1
// inclusive and reads xs[i+1], running one past the end.
package pairbug

// SumDiffs returns the sum of xs[i+1]-xs[i] for all adjacent pairs. For fewer
// than two elements it returns 0.
func SumDiffs(xs []int) int {
	total := 0
	// CHANGE CODE BELOW THIS LINE
	for i := 0; i < len(xs); i++ {
		// CHANGE CODE ABOVE THIS LINE
		total += xs[i+1] - xs[i]
	}
	return total
}
