// Package nansafegen — Gopher Workplace challenge.
package nansafegen

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// MinIgnoringNaN returns the smallest non-NaN element and true.
// It returns zero and false when there is no such element.
func MinIgnoringNaN[T Float](s []T) (T, bool) {
	// TODO(candidate): skip NaN values while tracking the smallest.
	panic("not implemented")
}
