// Package movingavg — Gopher Workplace challenge.
package movingavg

// Window returns the mean of every n-sample sliding window, so a slice of m
// samples yields m-n+1 values. The whole thing must be one pass: keep a
// running sum, adding the entering sample and subtracting the leaving one,
// rather than re-summing each window. A non-positive n, or an n larger than
// the input, gives an empty, non-nil slice, and the input is not modified.
//
// Examples:
//
//	Window([]float64{1, 2, 3, 4}, 2) => []float64{1.5, 2.5, 3.5}
func Window(samples []float64, n int) []float64 {
	panic("not implemented")
}

// Smoothest returns the index of the window with the smallest mean, earliest
// index on a tie, and -1 when there are no windows.
//
// Examples:
//
//	Smoothest([]float64{3, 1, 1, 5}) => 1
func Smoothest(means []float64) int {
	panic("not implemented")
}
