// Package treemapgen — Gopher Workplace challenge.
package treemapgen

import (
	"cmp"
)

// SortedMap keeps its keys in ascending order.
// Use NewSorted to create one.
type SortedMap[K cmp.Ordered, V any] struct {
	items map[K]V
	keys  []K
}

// NewSorted returns an empty sorted map.
func NewSorted[K cmp.Ordered, V any]() *SortedMap[K, V] {
	// TODO(candidate): allocate the backing map and key list.
	panic("not implemented")
}

// Set stores v under k, keeping the key order sorted.
func (m *SortedMap[K, V]) Set(k K, v V) {
	// TODO(candidate): insert the key in order when new, then store the value.
	panic("not implemented")
}

// Get returns the value stored under k.
func (m *SortedMap[K, V]) Get(k K) (V, bool) {
	// TODO(candidate): look up the key.
	panic("not implemented")
}

// Keys returns the keys in ascending order.
func (m *SortedMap[K, V]) Keys() []K {
	// TODO(candidate): return a copy of the sorted key list.
	panic("not implemented")
}
