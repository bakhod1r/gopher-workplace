// Package cartesiangen — Gopher Workplace challenge.
package cartesiangen

// Pair holds two values of independent types.
type Pair[A, B any] struct {
	First  A
	Second B
}

// Product returns every (a, b) combination, a-major.
func Product[A, B any](as []A, bs []B) []Pair[A, B] {
	// TODO(candidate): pair every element of as with every element of bs.
	panic("not implemented")
}
