// Package rankvalueonlybug — Gopher Workplace challenge.
package rankvalueonlybug

import (
	"cmp"
	"slices"
)

// Rank returns the keys of m ordered by value, highest first.
// Keys with equal values are ordered by key, ascending.
//
// Examples:
//
//	Rank(map[string]int{"a": 1, "b": 1}) => []string{"a", "b"}
func Rank[K cmp.Ordered, V cmp.Ordered](m map[K]V) []K {
	// CHANGE CODE BELOW THIS LINE
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortStableFunc(keys, func(a, b K) int {
		return cmp.Compare(m[b], m[a])
	})
	return keys
	// CHANGE CODE ABOVE THIS LINE
}
