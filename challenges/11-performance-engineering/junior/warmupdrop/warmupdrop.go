// Package warmupdrop — Gopher Workplace challenge.
package warmupdrop

// Drop discards the first n samples — the warm-up rounds, where caches are
// cold and the JIT-free but still lazily-initialised parts of a program pay
// one-off costs. A non-positive n keeps everything, and an n at or beyond the
// length leaves an empty, non-nil slice. The input is not modified.
//
// Examples:
//
//	Drop([]float64{9, 1, 1}, 1) => []float64{1, 1}
func Drop(samples []float64, n int) []float64 {
	panic("not implemented")
}

// StableMean drops the warm-up samples and averages what is left. Nothing
// left gives 0.
//
// Examples:
//
//	StableMean([]float64{100, 2, 4}, 1) => 3
func StableMean(samples []float64, n int) float64 {
	panic("not implemented")
}
