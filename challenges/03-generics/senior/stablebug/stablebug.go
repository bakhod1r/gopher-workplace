// Package stablebug — Gopher Workplace challenge.
package stablebug

import (
	"cmp"
	"slices"
)

// SortedBy returns a copy of s sorted by key, ascending.
// Elements with equal keys keep their input order.
//
// Examples:
//
//	SortedBy(people, ageOf) => youngest first, ties in input order
func SortedBy[T any, K cmp.Ordered](s []T, key func(T) K) []T {
	// CHANGE CODE BELOW THIS LINE
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	slices.SortFunc(out, func(a, b T) int {
		return cmp.Compare(key(a), key(b))
	})
	return out
	// CHANGE CODE ABOVE THIS LINE
}
