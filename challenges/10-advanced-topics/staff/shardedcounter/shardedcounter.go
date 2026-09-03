// Package shardedcounter — Gopher Workplace challenge.
package shardedcounter

import (
	"hash/maphash"
	"sync"
)

// shard is one lock-protected slice of the key space, padded so two
// shards never share a cache line.
type shard struct {
	mu sync.Mutex
	m  map[string]int64
	_  [48]byte
}

// Counter is a sharded string counter.
type Counter struct {
	seed   maphash.Seed
	shards []shard
}

// NewCounter returns a counter with n shards (rounded up to at least 1).
func NewCounter(n int) *Counter {
	if n < 1 {
		n = 1
	}
	c := &Counter{seed: maphash.MakeSeed(), shards: make([]shard, n)}
	for i := range c.shards {
		c.shards[i].m = make(map[string]int64)
	}
	return c
}

// shardFor returns the shard that owns key.
func (c *Counter) shardFor(key string) *shard {
	h := maphash.String(c.seed, key)
	return &c.shards[h%uint64(len(c.shards))]
}

// Total folds every shard into one map. Call it after the writers are done.
func (c *Counter) Total() map[string]int64 {
	out := make(map[string]int64)
	for i := range c.shards {
		s := &c.shards[i]
		s.mu.Lock()
		for k, v := range s.m {
			out[k] += v
		}
		s.mu.Unlock()
	}
	return out
}

// Add increments the counter for key.
//
// Counters are sharded by the key's hash so concurrent writers rarely touch
// the same lock; Total folds the shards after the writers are done.
//
// Examples:
//
//	c := NewCounter(4); c.Add("a", 1) => Total()["a"] == 1
func (c *Counter) Add(key string, n int64) {
	panic("not implemented")
}
