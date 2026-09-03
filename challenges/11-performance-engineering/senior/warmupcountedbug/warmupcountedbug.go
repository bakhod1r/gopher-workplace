// Package warmupcountedbug — Gopher Workplace challenge.
package warmupcountedbug

// StableMean averages the samples after discarding the first warmup of them,
// where caches are cold and one-time initialisation is still being paid.
// A non-positive warmup keeps everything; discarding every sample gives 0.
//
// Examples:
//
//	StableMean([]float64{100, 2, 4}, 1) => 3
func StableMean(samples []float64, warmup int) float64 {
	if warmup < 0 {
		warmup = 0
	}
	if warmup > len(samples) {
		warmup = len(samples)
	}
	// CHANGE CODE BELOW THIS LINE
	rest := samples[:len(samples)-warmup]
	// CHANGE CODE ABOVE THIS LINE
	if len(rest) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range rest {
		sum += v
	}
	return sum / float64(len(rest))
}
