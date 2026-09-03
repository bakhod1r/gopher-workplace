// Package striped — Gopher Workplace challenge.
package striped

import (
	"hash/maphash"
	"sync"
)

// Map is a key-value store split across independent shards, each with its own
// lock. Two goroutines touching different keys almost always take different
// locks, so throughput scales with the shard count instead of collapsing onto
// one mutex.
type Map struct {
	shards []shard
	seed   maphash.Seed
}

type shard struct {
	mu sync.RWMutex
	m  map[string]int
}

// New returns a map with n shards. A non-positive n gives 1.
//
// Examples:
//
//	New(16)
func New(n int) *Map {
	panic("not implemented")
}

// Shards reports the shard count.
//
// Examples:
//
//	New(16).Shards() => 16
func (m *Map) Shards() int {
	panic("not implemented")
}

// ShardOf returns the index of the shard a key belongs to. The same key must
// always map to the same shard, for the lifetime of this Map.
//
// Examples:
//
//	m.ShardOf("a") == m.ShardOf("a")
func (m *Map) ShardOf(key string) int {
	panic("not implemented")
}

// Set stores a value under its shard's write lock.
//
// Examples:
//
//	m.Set("a", 1)
func (m *Map) Set(key string, v int) {
	panic("not implemented")
}

// Get reads a value under its shard's read lock.
//
// Examples:
//
//	m.Get("a") => 1, true
func (m *Map) Get(key string) (int, bool) {
	panic("not implemented")
}

// Len returns the total number of entries across every shard.
//
// Examples:
//
//	m.Len() => 1
func (m *Map) Len() int {
	panic("not implemented")
}
