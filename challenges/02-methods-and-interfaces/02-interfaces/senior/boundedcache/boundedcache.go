// Package boundedcache — Gopher Workplace challenge.
package boundedcache

// Source is the slow backing store.
type Source interface {
	Get(key string) string
}

// CountingSource records how many lookups it served.
type CountingSource struct {
	Data  map[string]string
	Calls int
}

// Get reads the backing data.
func (s *CountingSource) Get(key string) string {
	s.Calls++
	return s.Data[key]
}

// Cache is a read-through cache with a hard entry ceiling.
type Cache struct {
	inner   Source
	Max     int
	entries map[string]string
	order   []string
}

// NewCache wraps a source with a cache of at most max entries.
func NewCache(inner Source, max int) *Cache {
	return &Cache{
		inner:   inner,
		Max:     max,
		entries: make(map[string]string, max),
		order:   make([]string, 0, max),
	}
}

// Get serves from the cache, or fetches and caches, evicting when full.
//
// Examples:
//
//	Max 2; Get("a"), Get("b"), Get("c") => "a" was evicted
func (c *Cache) Get(key string) string {
	// TODO(candidate): hit, else fetch; evict the oldest when at Max.
	panic("not implemented")
}

// Len returns how many entries are cached.
func (c *Cache) Len() int {
	// TODO(candidate): entry count.
	panic("not implemented")
}
