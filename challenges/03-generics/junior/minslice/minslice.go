// Package minslice — Gopher Workplace challenge.
package minslice

import (
	"cmp"
)

// MinOf returns the smallest element of s and true.
// It returns the zero value and false for an empty slice.
//
// Examples:
//
//	MinOf([]int{4, 1, 3}) => 1, true
//	MinOf([]int{})        => 0, false
func MinOf[T cmp.Ordered](s []T) (T, bool) {
	// TODO(candidate): track the smallest element seen so far.
	panic("not implemented")
}
