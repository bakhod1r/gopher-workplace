// Package profilecache — Gopher Workplace challenge.
package profilecache

import (
	"sync"
	"sync/atomic"
)

// Cache holds user profiles in front of the identity service. Reads outnumber
// writes by orders of magnitude, so lookups take a read lock and only the
// rare refresh takes the write lock. The hit/miss counters are atomic so a
// lookup never has to upgrade to an exclusive lock just to record a stat.
type Cache struct {
	mu       sync.RWMutex
	profiles map[string]string
	hits     atomic.Int64
	misses   atomic.Int64
}

// NewCache returns an empty cache.
//
// Examples:
//
//	h, m := NewCache().Stats() => 0, 0
func NewCache() *Cache {
	return &Cache{profiles: make(map[string]string)}
}

// Put stores or replaces a user's profile.
//
// Examples:
//
//	c := NewCache(); c.Put("u1", "ada"); c.Get("u1") => "ada", true
func (c *Cache) Put(userID, profile string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Get returns a cached profile and records a hit or a miss.
//
// Examples:
//
//	c.Put("u1", "ada"); c.Get("u1")  => "ada", true
//	NewCache().Get("nobody")         => "", false
func (c *Cache) Get(userID string) (string, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Invalidate drops a user's entry and reports whether one was present.
//
// Examples:
//
//	c.Put("u1", "ada"); c.Invalidate("u1") => true
//	NewCache().Invalidate("u1")            => false
func (c *Cache) Invalidate(userID string) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Stats returns the hit and miss totals observed so far.
//
// Examples:
//
//	c.Put("u1", "ada"); c.Get("u1"); c.Get("u2"); c.Stats() => 1, 1
func (c *Cache) Stats() (hits, misses int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Len returns the number of cached profiles.
func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.profiles)
}
