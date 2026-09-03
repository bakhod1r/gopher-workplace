// Package bucketcounts — Gopher Workplace challenge.
package bucketcounts

// Index returns which bucket v falls into, given ascending upper bounds: the
// first bucket whose bound is at or above v, or len(bounds) for anything
// past the last bound. Use a binary search — a metrics hot path observes
// millions of values against dozens of bounds.
//
// Examples:
//
//	Index([]float64{1, 5, 10}, 3) => 1
func Index(bounds []float64, v float64) int {
	panic("not implemented")
}

// Counts tallies the samples into len(bounds)+1 buckets, the last one holding
// everything past the final bound.
//
// Examples:
//
//	Counts([]float64{1, 5}, []float64{0.5, 3, 100}) => []int64{1, 1, 1}
func Counts(bounds []float64, samples []float64) []int64 {
	panic("not implemented")
}

// Cumulative converts bucket counts into the cumulative form Prometheus
// histograms expose, where each entry counts everything at or below its
// bound. The input is not modified.
//
// Examples:
//
//	Cumulative([]int64{1, 2, 3}) => []int64{1, 3, 6}
func Cumulative(counts []int64) []int64 {
	panic("not implemented")
}
