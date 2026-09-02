// Package shardedmap — Gopher Workplace challenge.
package shardedmap

import (
	"hash/fnv"
	"sync"
)

type shard struct {
	mu   sync.RWMutex
	data map[string]int
}

// ShardedMap reduces lock contention by sharding keys.
type ShardedMap struct {
	shards []*shard
	num    int
}

// New creates a map with `num` shards.
func New(num int) *ShardedMap {
	s := &ShardedMap{num: num, shards: make([]*shard, num)}
	for i := 0; i < num; i++ {
		s.shards[i] = &shard{data: make(map[string]int)}
	}
	return s
}

func (m *ShardedMap) getShard(key string) *shard {
	h := fnv.New32a()
	h.Write([]byte(key))
	return m.shards[int(h.Sum32())%m.num]
}

// Set stores val under key in that key's shard, holding only that shard's
// write lock.
func (m *ShardedMap) Set(key string, val int) {
	// TODO(candidate): select the shard, take its write lock, store, release.
	panic("not implemented")
}

// Get returns the value for key and whether it was present, holding only that
// shard's read lock.
func (m *ShardedMap) Get(key string) (int, bool) {
	// TODO(candidate): select the shard, take its read lock, look up, release.
	panic("not implemented")
}

// Len returns the total number of keys across all shards.
func (m *ShardedMap) Len() int {
	n := 0
	for _, s := range m.shards {
		s.mu.RLock()
		n += len(s.data)
		s.mu.RUnlock()
	}
	return n
}
