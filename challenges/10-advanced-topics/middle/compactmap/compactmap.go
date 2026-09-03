// Package compactmap — Gopher Workplace challenge.
package compactmap

// Compact returns a new map holding the same entries as m, sized to the
// entries it actually keeps.
//
// A map that grew to millions of entries keeps its bucket array after the
// entries are deleted; rebuilding is the only way to release it.
//
// Examples:
//
//	Compact(map[string]int{"a": 1}) => a new map[a:1]
func Compact(m map[string]int) map[string]int {
	panic("not implemented")
}
