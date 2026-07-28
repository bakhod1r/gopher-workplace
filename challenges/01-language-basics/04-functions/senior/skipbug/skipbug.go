// Package skipbug sums only the positive numbers using a range loop with
// continue. A planted bug continues on the POSITIVE elements (skipping the ones
// it should add) instead of skipping the non-positive ones.
package skipbug

// SumPositive returns the sum of the positive elements of xs.
func SumPositive(xs []int) int {
	sum := 0
	for _, v := range xs {
		// CHANGE CODE BELOW THIS LINE
		if v > 0 {
			continue
		}
		// CHANGE CODE ABOVE THIS LINE
		sum += v
	}
	return sum
}
