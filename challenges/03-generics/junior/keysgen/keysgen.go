// Package keysgen — Gopher Workplace challenge.
package keysgen

// Keys returns the keys of m in unspecified order.
//
// Examples:
//
//	Keys(map[string]int{"a": 1}) => []string{"a"}
//	Keys(map[string]int{})        => []string{}
func Keys[K comparable, V any](m map[K]V) []K {
	// TODO(candidate): collect every key into a slice.
	panic("not implemented")
}
