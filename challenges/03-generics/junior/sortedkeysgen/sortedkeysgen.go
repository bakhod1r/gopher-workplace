// Package sortedkeysgen — Gopher Workplace challenge.
package sortedkeysgen

import (
	"cmp"
)

// SortedKeys returns the keys of m in ascending order.
//
// Examples:
//
//	SortedKeys(map[string]int{"b": 1, "a": 2}) => []string{"a", "b"}
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	// TODO(candidate): collect the keys, then sort them.
	panic("not implemented")
}
