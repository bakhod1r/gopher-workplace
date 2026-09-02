// Package indexmax — Gopher Workplace challenge.
package indexmax

import (
	"cmp"
)

// IndexOfMax returns the index of the largest element,
// or -1 when s is empty. The first maximum wins.
//
// Examples:
//
//	IndexOfMax([]int{1, 9, 9}) => 1
//	IndexOfMax([]int{})        => -1
func IndexOfMax[T cmp.Ordered](s []T) int {
	// TODO(candidate): track the index of the largest element.
	panic("not implemented")
}
