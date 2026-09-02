// Package triplegen — Gopher Workplace challenge.
package triplegen

// Triple holds three values of independent types.
type Triple[A, B, C any] struct {
	First  A
	Second B
	Third  C
}

// MakeTriple returns a triple holding a, b and c.
func MakeTriple[A, B, C any](a A, b B, c C) Triple[A, B, C] {
	// TODO(candidate): build a triple from the three arguments.
	panic("not implemented")
}

// Rotated returns the triple with the fields rotated left:
// (a, b, c) becomes (b, c, a).
func (t Triple[A, B, C]) Rotated() Triple[B, C, A] {
	// TODO(candidate): rotate the fields one position left.
	panic("not implemented")
}
