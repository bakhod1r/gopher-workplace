// Package swaptypegen — Gopher Workplace challenge.
package swaptypegen

import (
	"cmp"
)

// SamePair holds two values of the same ordered type.
type SamePair[T cmp.Ordered] struct {
	First  T
	Second T
}

// Swap exchanges the two stored values.
func (p *SamePair[T]) Swap() {
	// TODO(candidate): exchange the two fields.
	panic("not implemented")
}

// Ordered returns the two values smallest first.
func (p SamePair[T]) Ordered() (T, T) {
	// TODO(candidate): return the smaller value first.
	panic("not implemented")
}
