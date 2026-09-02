// Package clampslice — Gopher Workplace challenge.
package clampslice

import (
	"cmp"
)

// ClampAll returns a new slice with every element limited to [lo, hi].
//
// Examples:
//
//	ClampAll([]int{-1, 2, 9}, 0, 3) => []int{0, 2, 3}
func ClampAll[T cmp.Ordered](s []T, lo, hi T) []T {
	// TODO(candidate): clamp each element into a new slice.
	panic("not implemented")
}
