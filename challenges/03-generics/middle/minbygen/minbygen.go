// Package minbygen — Gopher Workplace challenge.
package minbygen

import (
	"cmp"
)

// MinBy returns the element with the smallest key and true.
// On a tie the earlier element wins.
func MinBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool) {
	// TODO(candidate): track the element whose key is smallest.
	panic("not implemented")
}
