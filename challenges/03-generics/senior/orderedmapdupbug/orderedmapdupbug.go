// Package orderedmapdupbug — Gopher Workplace challenge.
package orderedmapdupbug

// Ordered is a map that remembers insertion order.
// Use NewOrdered to create one.
type Ordered[K comparable, V any] struct {
	items map[K]V
	keys  []K
}

// Set stores v under k. An existing key keeps its position.
func (o *Ordered[K, V]) Set(k K, v V) {
	// CHANGE CODE BELOW THIS LINE
	o.keys = append(o.keys, k)
	o.items[k] = v
	// CHANGE CODE ABOVE THIS LINE
}

// Get returns the value stored under k. It is provided for you.
func (o *Ordered[K, V]) Get(k K) (V, bool) {
	v, ok := o.items[k]
	return v, ok
}

// Keys returns the keys in insertion order. It is provided for you.
func (o *Ordered[K, V]) Keys() []K {
	out := make([]K, len(o.keys))
	copy(out, o.keys)
	return out
}

// NewOrdered returns an empty ordered map. It is provided for you.
func NewOrdered[K comparable, V any]() *Ordered[K, V] {
	return &Ordered[K, V]{items: make(map[K]V), keys: make([]K, 0)}
}
