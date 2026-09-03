// Package latencytrim — Gopher Workplace challenge.
package latencytrim

// TrimmedMean averages the samples after discarding the lowest and highest
// pct percent — the standard defence against a couple of wild measurements
// dragging the average around. The count discarded from each end is
// floor(pct/100 * n). pct is clamped into [0,50), trimming everything gives 0,
// and the input is not modified.
//
// Examples:
//
//	TrimmedMean([]float64{1, 2, 3, 4, 100}, 20) => 3
func TrimmedMean(samples []float64, pct float64) float64 {
	panic("not implemented")
}

// Mean is the untrimmed average, kept alongside so the two can be compared.
//
// Examples:
//
//	Mean([]float64{1, 2, 3, 4, 100}) => 22
func Mean(samples []float64) float64 {
	panic("not implemented")
}
