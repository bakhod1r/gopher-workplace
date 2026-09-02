// Package searchbygen — Gopher Workplace challenge.
package searchbygen

import (
	"cmp"
)

// SearchBy binary-searches s, which must be sorted by key,
// returning the index of target and whether it was found.
func SearchBy[T any, K cmp.Ordered](s []T, key func(T) K, target K) (int, bool) {
	// TODO(candidate): binary-search using the projected key.
	panic("not implemented")
}
