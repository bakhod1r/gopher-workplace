// Package unionprecisionbug — Gopher Workplace challenge.
package unionprecisionbug

// Mean returns the arithmetic mean of xs as a float64.
// It returns 0 for an empty slice.
//
// Examples:
//
//	Mean([]float64{0.5, 0.5}) => 0.5
func Mean[T int | int64 | float64](xs []T) float64 {
	// CHANGE CODE BELOW THIS LINE
	if len(xs) == 0 {
		return 0
	}
	var sum int64
	for _, v := range xs {
		sum += int64(v)
	}
	return float64(sum) / float64(len(xs))
	// CHANGE CODE ABOVE THIS LINE
}
