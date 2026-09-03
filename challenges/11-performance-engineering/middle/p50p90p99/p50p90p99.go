// Package p50p90p99 — Gopher Workplace challenge.
package p50p90p99

// Percentile returns the p-th percentile of the samples by the nearest-rank
// method: sort ascending, then take the element at rank ceil(p/100 * n),
// counting from 1. A p at or below 0 gives the smallest sample and a p at or
// above 100 the largest. No samples gives 0, and the input is not modified.
//
// Examples:
//
//	Percentile([]float64{1, 2, 3, 4}, 50) => 2
func Percentile(samples []float64, p float64) float64 {
	panic("not implemented")
}

// Summary returns the three latencies every dashboard shows, sorting the data
// once rather than three times.
//
// Examples:
//
//	Summary([]float64{1, 2, 3, 4}) => 2, 4, 4
func Summary(samples []float64) (p50, p90, p99 float64) {
	panic("not implemented")
}
