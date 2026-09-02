// Package registrygen — Gopher Workplace challenge.
package registrygen

// Registry maps unique keys of K to values of V.
// Use NewRegistry to create one.
type Registry[K comparable, V any] struct {
	items map[K]V
}

// NewRegistry returns a ready-to-use registry.
func NewRegistry[K comparable, V any]() *Registry[K, V] {
	// TODO(candidate): return a registry with its map allocated.
	panic("not implemented")
}

// Register stores v under k and reports whether it was stored.
// An existing key is left untouched and reports false.
func (r *Registry[K, V]) Register(k K, v V) bool {
	// TODO(candidate): store v only when k is free.
	panic("not implemented")
}

// Lookup returns the value stored under k and whether it exists.
func (r *Registry[K, V]) Lookup(k K) (V, bool) {
	// TODO(candidate): look up the key.
	panic("not implemented")
}

// Len returns the number of registered entries.
func (r *Registry[K, V]) Len() int {
	// TODO(candidate): report how many entries are stored.
	panic("not implemented")
}
