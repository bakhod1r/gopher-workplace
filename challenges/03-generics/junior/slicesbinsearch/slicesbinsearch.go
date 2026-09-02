// Package slicesbinsearch — Gopher Workplace challenge.
package slicesbinsearch

import (
	"cmp"
)

// Find returns the index of v in the sorted slice s and true,
// or the insertion point and false when v is absent.
func Find[T cmp.Ordered](s []T, v T) (int, bool) {
	// TODO(candidate): use the stdlib binary search.
	panic("not implemented")
}
