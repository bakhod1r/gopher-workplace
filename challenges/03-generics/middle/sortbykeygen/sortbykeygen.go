// Package sortbykeygen — Gopher Workplace challenge.
package sortbykeygen

import (
	"cmp"
)

// SortedBy returns a copy of s sorted by key, ascending.
// Elements with equal keys keep their input order.
func SortedBy[T any, K cmp.Ordered](s []T, key func(T) K) []T {
	// TODO(candidate): clone, then stably sort by the projected key.
	panic("not implemented")
}
