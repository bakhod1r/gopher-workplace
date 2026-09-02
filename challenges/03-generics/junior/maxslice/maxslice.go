// Package maxslice — Gopher Workplace challenge.
package maxslice

import (
	"cmp"
)

// MaxOf returns the largest element of s and true.
// It returns the zero value and false for an empty slice.
//
// Examples:
//
//	MaxOf([]int{1, 9, 3}) => 9, true
//	MaxOf([]int{})        => 0, false
func MaxOf[T cmp.Ordered](s []T) (T, bool) {
	// TODO(candidate): track the largest element seen so far.
	panic("not implemented")
}
