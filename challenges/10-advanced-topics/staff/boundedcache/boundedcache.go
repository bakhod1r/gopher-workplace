// Package boundedcache — Gopher Workplace challenge.
package boundedcache

import "sync"

// Cache is a bounded, concurrency-safe byte cache with FIFO eviction.
type Cache struct {
	mu    sync.Mutex
	limit int
	items map[string][]byte
	order []string
}

// NewCache returns a cache holding at most limit entries.
func NewCache(limit int) *Cache {
	if limit < 1 {
		limit = 1
	}
	return &Cache{limit: limit, items: make(map[string][]byte, limit), order: make([]string, 0, limit)}
}

// Get returns the stored bytes for key, if present.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.items[key]
	return v, ok
}

// Len reports how many entries the cache holds.
func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

// Put stores a copy of val under key, evicting the oldest entry when the
// cache is at capacity.
//
// The stored value must own its bytes — callers reuse their buffers — and
// the cache must never hold more than limit entries.
//
// Examples:
//
//	c := NewCache(2); c.Put("a", v) => Get("a") returns a copy of v
func (c *Cache) Put(key string, val []byte) {
	panic("not implemented")
}
