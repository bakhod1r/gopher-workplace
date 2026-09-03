// Package smallmapcache — Gopher Workplace challenge.
package smallmapcache

// Cache is a bounded key-value cache that evicts the oldest inserted entry
// when it is full — a fixed-memory memo table, not an LRU: a hit does not
// change an entry's eviction order.
type Cache struct {
	Cap   int // maximum entries; a non-positive Cap stores nothing
	items map[string]int
	order []string
	hits  int
	miss  int
}

// Get returns the cached value and whether it was present. A hit must not
// allocate.
//
// Examples:
//
//	c.Get("a") => 1, true
func (c *Cache) Get(key string) (int, bool) {
	panic("not implemented")
}

// Put stores a value, evicting the oldest entry when the cache is full.
// Overwriting an existing key updates it in place and does not change the
// eviction order.
//
// Examples:
//
//	c.Put("a", 1)
func (c *Cache) Put(key string, v int) {
	panic("not implemented")
}

// Len reports the number of entries held, and Stats the hit and miss counts.
//
// Examples:
//
//	c.Len() => 2
func (c *Cache) Len() int { panic("not implemented") }

// Stats reports how many Get calls hit and missed.
//
// Examples:
//
//	c.Stats() => 3, 1
func (c *Cache) Stats() (hits, misses int) { panic("not implemented") }
