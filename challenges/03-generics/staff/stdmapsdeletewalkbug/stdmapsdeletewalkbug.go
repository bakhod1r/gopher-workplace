// Package stdmapsdeletewalkbug — Gopher Workplace challenge.
package stdmapsdeletewalkbug

import (
	"maps"
	"slices"
)

// Prune removes every entry of m whose value is below limit.
// It returns the removed keys in ascending order.
//
// Examples:
//
//	Prune(map[string]int{"a": 1, "b": 5}, 3) => []string{"a"}
func Prune(m map[string]int, limit int) []string {
	// CHANGE CODE BELOW THIS LINE
	keys := slices.Sorted(maps.Keys(m))
	removed := make([]string, 0)
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			continue
		}
		maps.DeleteFunc(m, func(_ string, v int) bool { return v < limit })
		removed = append(removed, k)
	}
	return removed
	// CHANGE CODE ABOVE THIS LINE
}
