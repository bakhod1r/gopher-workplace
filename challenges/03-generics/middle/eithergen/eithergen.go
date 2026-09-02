// Package eithergen — Gopher Workplace challenge.
package eithergen

// Either holds a value of L or a value of R.
// Its zero value is a right-sided Either holding the zero R.
type Either[L, R any] struct {
	left   L
	right  R
	isLeft bool
}

// Left returns an Either holding a left value.
func Left[L, R any](v L) Either[L, R] {
	// TODO(candidate): build a left-sided Either.
	panic("not implemented")
}

// Right returns an Either holding a right value.
func Right[L, R any](v R) Either[L, R] {
	// TODO(candidate): build a right-sided Either.
	panic("not implemented")
}

// Unwrap returns the left value, the right value, and whether
// the Either is left-sided.
func (e Either[L, R]) Unwrap() (L, R, bool) {
	// TODO(candidate): report both slots and which one is set.
	panic("not implemented")
}
