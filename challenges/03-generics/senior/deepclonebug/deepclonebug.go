// Package deepclonebug — Gopher Workplace challenge.
package deepclonebug

import (
	"maps"
)

// Snapshot returns an independent copy of a map of slices.
// Neither the map nor its values may be shared with the original.
//
// Examples:
//
//	Snapshot(map[string][]int{"a": {1}}) => an independent copy
func Snapshot[K comparable, V any](m map[K][]V) map[K][]V {
	// CHANGE CODE BELOW THIS LINE
	out := maps.Clone(m)
	if out == nil {
		out = make(map[K][]V)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
