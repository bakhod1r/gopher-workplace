// Package signedonly — Gopher Workplace challenge.
package signedonly

// Signed is the set of signed integer types.
type Signed interface {
	~int | ~int64
}

// Negate returns -v.
// The constraint excludes unsigned types on purpose.
func Negate[T Signed](v T) T {
	// TODO(candidate): return the negation.
	panic("not implemented")
}

// AbsDiff returns the distance between a and b.
func AbsDiff[T Signed](a, b T) T {
	// TODO(candidate): return the magnitude of the difference.
	panic("not implemented")
}
