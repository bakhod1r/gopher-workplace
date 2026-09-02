// Package maxbykey — Gopher Workplace challenge.
package maxbykey

import (
	"cmp"
)

// MaxBy returns the element of s with the largest key and true.
// On a tie the earlier element wins.
// It returns the zero value and false for an empty slice.
//
// Examples:
//
//	MaxBy(people, ageOf) => the oldest person, true
func MaxBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool) {
	// TODO(candidate): track the element whose key is largest.
	panic("not implemented")
}
