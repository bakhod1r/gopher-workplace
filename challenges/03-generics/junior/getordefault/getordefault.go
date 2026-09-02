// Package getordefault — Gopher Workplace challenge.
package getordefault

// GetOr returns m[k] when the key is present, otherwise def.
//
// Examples:
//
//	GetOr(map[string]int{"a": 1}, "a", 9) => 1
//	GetOr(map[string]int{}, "a", 9)        => 9
func GetOr[K comparable, V any](m map[K]V, k K, def V) V {
	// TODO(candidate): use the comma-ok lookup and fall back to def.
	panic("not implemented")
}
