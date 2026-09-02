// Package orderleakbug — Gopher Workplace challenge.
package orderleakbug

import (
	"cmp"
	"slices"
)

// RankByCount returns the distinct values of s ordered by descending
// count, ties broken by ascending value.
//
// The result must be identical for equal inputs on every run.
//
// Examples:
//
//	RankByCount([]string{"b", "a", "b", "c"}) => [b a c]
func RankByCount[T cmp.Ordered](s []T) []T {
	// CHANGE CODE BELOW THIS LINE
	cnt := make(map[T]int, len(s))
	for _, v := range s {
		cnt[v]++
	}
	keys := make([]T, 0, len(cnt))
	for k := range cnt {
		keys = append(keys, k)
	}
	slices.SortStableFunc(keys, func(a, b T) int {
		return cmp.Compare(cnt[b], cnt[a])
	})
	return keys
	// CHANGE CODE ABOVE THIS LINE
}
