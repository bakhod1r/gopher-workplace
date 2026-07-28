// Package sums accumulates a decimal that only float64 represents tightly.
// Planted narrow float type loses precision.
package sums

// SumTenths adds 0.1 exactly n times and returns the total.
func SumTenths(n int) float64 {
	// CHANGE CODE BELOW THIS LINE
	var total float32
	// CHANGE CODE ABOVE THIS LINE
	for i := 0; i < n; i++ {
		total += 0.1
	}
	return float64(total)
}
