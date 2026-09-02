// Package gcdgen — Gopher Workplace challenge.
package gcdgen

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// GCD returns the greatest common divisor of a and b.
// The result is never negative; GCD(0, 0) is 0.
func GCD[T Integer](a, b T) T {
	// TODO(candidate): run Euclid's algorithm on the magnitudes.
	panic("not implemented")
}
