// Package pairgen — Gopher Workplace challenge.
package pairgen

// Pair holds two values of independent types.
type Pair[A, B any] struct {
	First  A
	Second B
}

// Swapped returns a pair with the two fields exchanged.
func (p Pair[A, B]) Swapped() Pair[B, A] {
	// TODO(candidate): return a pair with the fields exchanged.
	panic("not implemented")
}

// MakePair returns a pair holding a and b.
func MakePair[A, B any](a A, b B) Pair[A, B] {
	// TODO(candidate): build a pair from the two arguments.
	panic("not implemented")
}
