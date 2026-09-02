// Package variancegen — Gopher Workplace challenge.
package variancegen

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// Variance returns the population variance of s.
// It returns 0 for fewer than two samples.
func Variance[T Float](s []T) T {
	// TODO(candidate): average the squared deviations from the mean.
	panic("not implemented")
}
