// Package detmapgen — Gopher Workplace challenge.
package detmapgen

import (
	"cmp"
)

// Entry is one key/value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries returns the map's entries sorted by key.
func Entries[K cmp.Ordered, V any](m map[K]V) []Entry[K, V] {
	// TODO(candidate): collect the entries, then sort them by key.
	panic("not implemented")
}
