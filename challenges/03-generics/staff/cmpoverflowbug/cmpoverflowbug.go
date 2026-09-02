// Package cmpoverflowbug — Gopher Workplace challenge.
package cmpoverflowbug

import (
	"cmp"
	"slices"
)

// SortByKey sorts s in ascending key order, in place.
//
// It must be correct for the whole range of int keys, including
// values near math.MinInt and math.MaxInt.
//
// Examples:
//
//	SortByKey(rows, score) => rows ordered by score
func SortByKey[T any](s []T, key func(T) int) {
	// CHANGE CODE BELOW THIS LINE
	slices.SortFunc(s, func(a, b T) int {
		d := key(a) - key(b)
		return cmp.Compare(d, 0)
	})
	// CHANGE CODE ABOVE THIS LINE
}
