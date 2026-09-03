// Package quantileinterp — Gopher Workplace challenge.
package quantileinterp

// Interpolated returns the p-th percentile using linear interpolation between
// the two neighbouring ranks — the definition NumPy and most statistics
// packages use, and the one that reports a value the system may never have
// produced.
//
// With n sorted samples the position is (n-1) * p/100, zero-based: an exact
// integer position returns that sample, and a fractional one blends the two
// around it. p is clamped into [0,100], no samples gives 0, and the input is
// not modified.
//
// Examples:
//
//	Interpolated([]float64{1, 2, 3, 4}, 50) => 2.5
func Interpolated(samples []float64, p float64) float64 {
	panic("not implemented")
}

// NearestRank is the non-interpolating definition, returning an element that
// really occurred: the sample at rank ceil(p/100 * n), counted from 1.
//
// Examples:
//
//	NearestRank([]float64{1, 2, 3, 4}, 50) => 2
func NearestRank(samples []float64, p float64) float64 {
	panic("not implemented")
}
