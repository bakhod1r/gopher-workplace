// Package lazymapgen — Gopher Workplace challenge.
package lazymapgen

// Store is a map-backed store of V keyed by K.
// Its zero value is ready to use.
type Store[K comparable, V any] struct {
	items map[K]V
}

// Set stores v under k, allocating the map on first use.
func (s *Store[K, V]) Set(k K, v V) {
	// TODO(candidate): allocate the map when nil, then store.
	panic("not implemented")
}

// Get returns the value stored under k.
func (s *Store[K, V]) Get(k K) (V, bool) {
	// TODO(candidate): look the key up.
	panic("not implemented")
}

// Len returns how many entries are stored.
func (s *Store[K, V]) Len() int {
	// TODO(candidate): report the number of entries.
	panic("not implemented")
}
