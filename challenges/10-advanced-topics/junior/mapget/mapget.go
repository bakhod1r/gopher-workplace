// Package mapget — Gopher Workplace challenge.
package mapget

// Get returns the value stored under key and whether the key was present.
//
// A missing key reads as 0, which is also a value a key can hold — only the
// second result tells them apart.
//
// Examples:
//
//	Get(map[string]int{"a": 0}, "a") => 0, true
func Get(m map[string]int, key string) (int, bool) {
	panic("not implemented")
}
