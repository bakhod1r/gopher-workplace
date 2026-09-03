// Package latencyhist — Gopher Workplace challenge.
package latencyhist

// Histogram counts samples into n fixed-width buckets of the given width,
// starting at 0, plus one final overflow bucket for everything at or above
// n*width — so the result has n+1 entries. Negative samples are dropped, and a
// non-positive width or n gives an empty, non-nil slice.
//
// Examples:
//
//	Histogram([]float64{0, 5, 15}, 10, 1) => []int64{2, 1}
func Histogram(samples []float64, width float64, n int) []int64 {
	panic("not implemented")
}

// Busiest returns the index of the fullest bucket, earliest index on a tie,
// and -1 when every bucket is empty.
//
// Examples:
//
//	Busiest([]int64{1, 5, 5}) => 1
func Busiest(buckets []int64) int {
	panic("not implemented")
}
