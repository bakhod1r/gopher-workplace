// Package orderedmapgen — Gopher Workplace challenge.
package orderedmapgen

// Ordered is a map that remembers insertion order.
// Use NewOrdered to create one.
type Ordered[K comparable, V any] struct {
	items map[K]V
	keys  []K
}

// NewOrdered returns an empty ordered map.
func NewOrdered[K comparable, V any]() *Ordered[K, V] {
	// TODO(candidate): allocate the map and the key list.
	panic("not implemented")
}

// Set stores v under k. An existing key keeps its position.
func (o *Ordered[K, V]) Set(k K, v V) {
	// TODO(candidate): store the value, recording the key only when new.
	panic("not implemented")
}

// Get returns the value stored under k.
func (o *Ordered[K, V]) Get(k K) (V, bool) {
	// TODO(candidate): look up the key.
	panic("not implemented")
}

// Keys returns the keys in insertion order.
func (o *Ordered[K, V]) Keys() []K {
	// TODO(candidate): return a copy of the key order.
	panic("not implemented")
}
