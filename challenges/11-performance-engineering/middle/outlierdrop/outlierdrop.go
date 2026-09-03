// Package outlierdrop — Gopher Workplace challenge.
package outlierdrop

// Quartiles returns the first and third quartiles by the nearest-rank rule
// used elsewhere in this topic: the samples at ranks ceil(0.25*n) and
// ceil(0.75*n), counting from 1 over a sorted copy. No samples gives 0, 0.
//
// Examples:
//
//	Quartiles([]float64{1, 2, 3, 4}) => 1, 3
func Quartiles(samples []float64) (q1, q3 float64) {
	panic("not implemented")
}

// Filter drops the samples outside [q1 - k*IQR, q3 + k*IQR], the standard
// interquartile-range rule for outliers. Order is preserved, the input is not
// modified, a negative k is treated as 0, and no samples gives an empty,
// non-nil slice.
//
// Examples:
//
//	Filter([]float64{1, 2, 3, 4, 1000}, 1.5) => []float64{1, 2, 3, 4}
func Filter(samples []float64, k float64) []float64 {
	panic("not implemented")
}
