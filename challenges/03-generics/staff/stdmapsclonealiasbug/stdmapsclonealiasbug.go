// Package stdmapsclonealiasbug — Gopher Workplace challenge.
package stdmapsclonealiasbug

import (
	"maps"
)

// CloneTags returns a deep copy of m.
// Mutating a slice in the copy must not be visible through m.
//
// Examples:
//
//	CloneTags(map[string][]string{"a": {"x"}}) => an independent copy
func CloneTags(m map[string][]string) map[string][]string {
	// CHANGE CODE BELOW THIS LINE
	return maps.Clone(m)
	// CHANGE CODE ABOVE THIS LINE
}
