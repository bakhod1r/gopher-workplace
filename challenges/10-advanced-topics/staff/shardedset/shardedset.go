// Package shardedset — Gopher Workplace challenge.
package shardedset

import (
	"hash/maphash"
	"sync"
)

// bucket is one shard of the set, padded to a cache line.
type bucket struct {
	mu sync.Mutex
	m  map[string]struct{}
	_  [48]byte
}

// Set is a striped concurrent string set.
type Set struct {
	seed    maphash.Seed
	buckets []bucket
}

// NewSet returns a set with n shards.
func NewSet(n int) *Set {
	if n < 1 {
		n = 1
	}
	s := &Set{seed: maphash.MakeSeed(), buckets: make([]bucket, n)}
	for i := range s.buckets {
		s.buckets[i].m = make(map[string]struct{})
	}
	return s
}

// bucketFor returns the shard owning key.
func (s *Set) bucketFor(key string) *bucket {
	h := maphash.String(s.seed, key)
	return &s.buckets[h%uint64(len(s.buckets))]
}

// Has reports whether key is present.
func (s *Set) Has(key string) bool {
	b := s.bucketFor(key)
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.m[key]
	return ok
}

// Len reports the total number of keys.
func (s *Set) Len() int {
	n := 0
	for i := range s.buckets {
		b := &s.buckets[i]
		b.mu.Lock()
		n += len(b.m)
		b.mu.Unlock()
	}
	return n
}

// Add inserts key and reports whether it was newly added.
//
// The set is striped across shards so concurrent writers of different keys
// rarely contend, and each shard is padded onto its own cache line.
//
// Examples:
//
//	s := NewSet(4); s.Add("a") => true, then false
func (s *Set) Add(key string) bool {
	panic("not implemented")
}
