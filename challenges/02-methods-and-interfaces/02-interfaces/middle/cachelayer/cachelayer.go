// Package cachelayer — Gopher Workplace challenge.
package cachelayer

// Source looks a key up.
type Source interface {
	Get(key string) (string, bool)
}

// SlowSource is the backing store; it counts the lookups it serves.
type SlowSource struct {
	Data  map[string]string
	Calls int
}

// Get reads from the backing data and counts the call.
func (s *SlowSource) Get(key string) (string, bool) {
	// TODO(candidate): count the call, then read the map.
	panic("not implemented")
}

// entry is one remembered lookup.
type entry struct {
	value string
	found bool
}

// Cache is a read-through cache over any Source.
type Cache struct {
	inner   Source
	entries map[string]entry
}

// NewCache wraps a source.
func NewCache(inner Source) *Cache {
	return &Cache{inner: inner, entries: make(map[string]entry)}
}

// Get serves from memory or fetches once from the wrapped source.
//
// Examples:
//
//	c.Get("a") twice => the source is consulted only once
func (c *Cache) Get(key string) (string, bool) {
	// TODO(candidate): remember hits AND misses.
	panic("not implemented")
}
