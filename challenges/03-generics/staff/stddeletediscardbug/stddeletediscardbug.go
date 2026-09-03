// Package stddeletediscardbug — Gopher Workplace challenge.
package stddeletediscardbug

import (
	"slices"
)

// Purge removes every item for which drop reports true.
// It reuses s's storage and returns the shortened slice.
//
// Examples:
//
//	Purge([]int{1, 2, 3, 4}, even) => []int{1, 3}
func Purge[T any](s []T, drop func(T) bool) []T {
	// CHANGE CODE BELOW THIS LINE
	kept := slices.DeleteFunc(s, drop)
	if len(kept) == 0 {
		return s[:0]
	}
	return s
	// CHANGE CODE ABOVE THIS LINE
}
