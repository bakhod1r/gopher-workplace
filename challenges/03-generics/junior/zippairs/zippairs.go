// Package zippairs — Gopher Workplace challenge.
package zippairs

// Pair holds two values of independent types.
type Pair[A, B any] struct {
	First  A
	Second B
}

// Zip pairs a[i] with b[i], stopping at the shorter slice.
func Zip[A, B any](a []A, b []B) []Pair[A, B] {
	// TODO(candidate): pair matching positions until the shorter slice runs out.
	panic("not implemented")
}
