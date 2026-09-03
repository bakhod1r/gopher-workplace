// Package bestofn — Gopher Workplace challenge.
package bestofn

// Best returns the smallest sample and its index — the least-disturbed run,
// which is what you report when the noise on a machine is one-sided. Ties go
// to the earliest index; no samples gives 0, -1.
//
// Examples:
//
//	Best([]float64{5, 2, 9}) => 2, 1
func Best(samples []float64) (float64, int) {
	panic("not implemented")
}

// Spread returns the ratio of the largest sample to the smallest, a
// one-number answer to "how noisy was this machine". A smallest sample of
// zero or no samples gives 0.
//
// Examples:
//
//	Spread([]float64{10, 20}) => 2
func Spread(samples []float64) float64 {
	panic("not implemented")
}
