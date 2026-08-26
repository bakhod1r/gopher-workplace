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

// Set sets a value in the correct shard.
func (m *ShardedMap) Set(key string, val int) {
	// TODO(candidate): get shard, lock, set, unlock
	panic("not implemented")
}
