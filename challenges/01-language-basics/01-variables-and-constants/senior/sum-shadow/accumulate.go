// Package accumulate sums positives. A planted shadowing bug drops the result.
package accumulate

// SumPositive returns the sum of the positive elements of xs.
func SumPositive(xs []int) int {
	total := 0
	for _, x := range xs {
		if x > 0 {
			// CHANGE CODE BELOW THIS LINE
			total := total + x
			// CHANGE CODE ABOVE THIS LINE
			_ = total
		}
	}
	return total
}
