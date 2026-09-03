// Package latencyavg — Gopher Workplace challenge.
package latencyavg

// Mean returns the arithmetic mean of the samples. No samples gives 0.
//
// Examples:
//
//	Mean([]float64{1, 2, 3}) => 2
func Mean(samples []float64) float64 {
	panic("not implemented")
}

// WeightedMean returns the mean of values weighted by counts — the way you
// combine per-endpoint averages back into one number. Pairs beyond the
// shorter slice are ignored, non-positive weights are skipped, and a total
// weight of zero gives 0.
//
// Examples:
//
//	WeightedMean([]float64{10, 20}, []int{1, 3}) => 17.5
func WeightedMean(values []float64, weights []int) float64 {
	panic("not implemented")
}
