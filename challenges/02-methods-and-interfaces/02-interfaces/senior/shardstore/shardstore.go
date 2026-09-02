// Package shardstore — Gopher Workplace challenge.
package shardstore

import "sync"

// Store is a concurrent key/value store.
type Store interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

// shard is one independently locked partition.
type shard struct {
	mu   sync.Mutex
	data map[string]string
}

// ShardedStore partitions keys across independently locked shards.
type ShardedStore struct {
	shards []*shard
}

// NewShardedStore returns a store with n shards (at least 1).
func NewShardedStore(n int) *ShardedStore {
	if n < 1 {
		n = 1
	}
	shards := make([]*shard, n)
	for i := range shards {
		shards[i] = &shard{data: make(map[string]string)}
	}
	return &ShardedStore{shards: shards}
}

// shardFor returns the shard owning key. The same key must always map to the
// same shard.
func (s *ShardedStore) shardFor(key string) *shard {
	// TODO(candidate): deterministic hash modulo the shard count.
	panic("not implemented")
}

// Put stores a value.
func (s *ShardedStore) Put(key, value string) {
	// TODO(candidate): lock only the owning shard.
	panic("not implemented")
}

// Get reads a value.
//
// Examples:
//
//	Put("a", "1"); Get("a") => "1", true
func (s *ShardedStore) Get(key string) (string, bool) {
	// TODO(candidate): lock only the owning shard.
	panic("not implemented")
}

// Len returns the total number of keys across all shards.
func (s *ShardedStore) Len() int {
	// TODO(candidate): sum the shard sizes under their locks.
	panic("not implemented")
}
