// Package sortstablecheckgen — Gopher Workplace challenge.
package sortstablecheckgen

import (
	"cmp"
)

// IsStableBy reports whether s is sorted by key and every group
// of equal keys appears in the same relative order as in original.
func IsStableBy[T comparable, K cmp.Ordered](s, original []T, key func(T) K) bool {
	// TODO(candidate): check the ordering, then the relative order within equal keys.
	panic("not implemented")
}
