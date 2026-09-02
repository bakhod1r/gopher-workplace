// Package slicesmaxmin — Gopher Workplace challenge.
package slicesmaxmin

import (
	"cmp"
)

// Peak returns the largest element of s and true.
// It returns the zero value and false for an empty slice.
func Peak[T cmp.Ordered](s []T) (T, bool) {
	// TODO(candidate): guard the empty case, then use the stdlib.
	panic("not implemented")
}

// Floor returns the smallest element of s and true.
// It returns the zero value and false for an empty slice.
func Floor[T cmp.Ordered](s []T) (T, bool) {
	// TODO(candidate): guard the empty case, then use the stdlib.
	panic("not implemented")
}
