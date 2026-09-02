// Package cachegen — Gopher Workplace challenge.
package cachegen

// Cache stores at most size entries, evicting the oldest first.
// Use NewCache to create one.
type Cache[K comparable, V any] struct {
	items map[K]V
	order []K
	size  int
}

// NewCache returns a cache holding at most size entries.
// A size <= 0 means the cache stores nothing.
func NewCache[K comparable, V any](size int) *Cache[K, V] {
	// TODO(candidate): allocate the map and store the size.
	panic("not implemented")
}

// Put stores v under k, evicting the oldest entry when full.
// Re-putting an existing key updates it without changing its age.
func (c *Cache[K, V]) Put(k K, v V) {
	// TODO(candidate): store the entry, evicting the oldest when over capacity.
	panic("not implemented")
}

// Get returns the value stored under k and whether it was found.
func (c *Cache[K, V]) Get(k K) (V, bool) {
	// TODO(candidate): look up the key.
	panic("not implemented")
}
