// Package lrugen — Gopher Workplace challenge.
package lrugen

// LRU is a cache that evicts the least recently used entry.
// Use NewLRU to create one.
type LRU[K comparable, V any] struct {
	items map[K]V
	order []K
	size  int
}

// NewLRU returns a cache holding at most size entries.
func NewLRU[K comparable, V any](size int) *LRU[K, V] {
	// TODO(candidate): allocate the map and store the capacity.
	panic("not implemented")
}

// Get returns the value for k and marks it most recently used.
func (c *LRU[K, V]) Get(k K) (V, bool) {
	// TODO(candidate): look up k and promote it on a hit.
	panic("not implemented")
}

// Put stores v under k, evicting the least recently used entry
// when the cache is full.
func (c *LRU[K, V]) Put(k K, v V) {
	// TODO(candidate): store the entry, promote it, then evict if needed.
	panic("not implemented")
}

// touch moves k to the most-recent end of the order list.
// It is provided for you.
func (c *LRU[K, V]) touch(k K) {
	for i, key := range c.order {
		if key == k {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	c.order = append(c.order, k)
}
