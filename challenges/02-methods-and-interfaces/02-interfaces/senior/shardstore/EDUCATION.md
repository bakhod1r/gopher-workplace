# Sharded Store

## Intuition

One mutex serialises every operation regardless of key. Sharding partitions the keyspace so unrelated keys take unrelated locks — the same total work, far less waiting.

## Approach

1. `shardFor` computes an FNV-1a style hash over the key bytes and indexes modulo the shard count.
2. `Put` and `Get` lock only the owning shard.
3. `Len` locks each shard in turn and sums the sizes.
4. The constructor clamps the shard count to at least 1.

## Solution

```go
func (s *ShardedStore) shardFor(key string) *shard {
	var h uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		h ^= uint32(key[i])
		h *= 16777619
	}
	return s.shards[int(h%uint32(len(s.shards)))]
}

func (s *ShardedStore) Put(key, value string) {
	sh := s.shardFor(key)
	sh.mu.Lock()
	defer sh.mu.Unlock()
	sh.data[key] = value
}

func (s *ShardedStore) Get(key string) (string, bool) {
	sh := s.shardFor(key)
	sh.mu.Lock()
	defer sh.mu.Unlock()
	v, ok := sh.data[key]
	return v, ok
}
```

## Walkthrough

`Len` is not a snapshot: it locks shards one at a time, so a concurrent writer can change a later shard after an earlier one was counted. The test calls it only after `wg.Wait()`, which is the honest way to use it.

## Pitfalls

- Using Go's map iteration or a random source for shard selection — the same key must always land in the same shard.
- Locking all shards in `Put` to be safe, which restores the bottleneck.
- Reading `len(sh.data)` without the lock, which `-race` reports immediately.
