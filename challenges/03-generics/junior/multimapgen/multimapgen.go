// Package multimapgen — Gopher Workplace challenge.
package multimapgen

// MultiMap stores several values of V per key of K.
// Use NewMultiMap to create one.
type MultiMap[K comparable, V any] struct {
	items map[K][]V
}

// NewMultiMap returns a ready-to-use multi map.
func NewMultiMap[K comparable, V any]() *MultiMap[K, V] {
	// TODO(candidate): return a multi map with its map allocated.
	panic("not implemented")
}

// Add appends v to the values stored under k.
func (m *MultiMap[K, V]) Add(k K, v V) {
	// TODO(candidate): append v to the slice for k.
	panic("not implemented")
}

// Get returns the values stored under k, oldest first.
// It returns an empty slice for an unknown key.
func (m *MultiMap[K, V]) Get(k K) []V {
	// TODO(candidate): return the values for k.
	panic("not implemented")
}
