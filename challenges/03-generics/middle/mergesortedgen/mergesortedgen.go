// Package mergesortedgen — Gopher Workplace challenge.
package mergesortedgen

import (
	"cmp"
)

// Merge returns the elements of two sorted slices, sorted.
// Equal elements keep a before b.
func Merge[T cmp.Ordered](a, b []T) []T {
	// TODO(candidate): walk both slices, always taking the smaller head.
	panic("not implemented")
}
