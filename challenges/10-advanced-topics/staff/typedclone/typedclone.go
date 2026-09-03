// Package typedclone — Gopher Workplace challenge.
package typedclone

// CloneMap returns a shallow copy of m with the same entries.
//
// A type parameter keeps the keys and values concrete, so nothing is boxed
// and the copy costs one allocation plus the entries.
//
// Examples:
//
//	CloneMap(map[string]int{"a": 1}) => a new map with the same entry
func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	panic("not implemented")
}
