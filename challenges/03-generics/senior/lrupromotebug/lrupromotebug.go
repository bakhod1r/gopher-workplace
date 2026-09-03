// Package lrupromotebug — Gopher Workplace challenge.
package lrupromotebug

// LRU is a cache that evicts the least recently used entry.
// Use NewLRU to create one.
type LRU[K comparable, V any] struct {
	items map[K]V
	order []K
	size  int
}

// NewLRU returns a cache holding at most size entries.
func NewLRU[K comparable, V any](size int) *LRU[K, V] {
	if size < 0 {
		size = 0
	}
	return &LRU[K, V]{items: make(map[K]V), order: make([]K, 0), size: size}
}

// Get returns the value for k and marks it most recently used.
func (c *LRU[K, V]) Get(k K) (V, bool) {
	// CHANGE CODE BELOW THIS LINE
	v, ok := c.items[k]
	if !ok {
		var zero V
		return zero, false
	}
	return v, true
	// CHANGE CODE ABOVE THIS LINE
}

// Put stores v under k, evicting the least recently used entry
// when the cache is full. It is provided for you.
func (c *LRU[K, V]) Put(k K, v V) {
	if c.size == 0 {
		return
	}
	c.items[k] = v
	c.touch(k)
	if len(c.order) > c.size {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.items, oldest)
	}
}

// touch moves k to the most-recent end. It is provided for you.
func (c *LRU[K, V]) touch(k K) {
	for i, key := range c.order {
		if key == k {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	c.order = append(c.order, k)
}
