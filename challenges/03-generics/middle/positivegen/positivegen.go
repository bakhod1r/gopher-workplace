// Package positivegen — Gopher Workplace challenge.
package positivegen

// Signed is the set of signed integer types.
type Signed interface {
	~int | ~int64
}

// AllPositive reports whether every element is greater than zero.
// An empty slice counts as positive.
func AllPositive[T Signed](s []T) bool {
	// TODO(candidate): report whether every element exceeds zero.
	panic("not implemented")
}

// FirstNonPositive returns the index of the first element that is
// not greater than zero, or -1.
func FirstNonPositive[T Signed](s []T) int {
	// TODO(candidate): return the index of the first element that is not positive.
	panic("not implemented")
}
