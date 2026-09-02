# Sharded Map

## Intuition

Contention is not about how many goroutines there are; it is about how many want
the *same* lock. Splitting one map into 32 gives 32 locks, and a well-spread
hash makes the chance that two writers meet on the same shard about 1 in 32.
Nothing about the data structure changes — only how finely it is guarded.

## Approach

1. Route the key to its shard.
2. Take that shard's lock — write for `Set`, read for `Get`.
3. Do the map operation and release.

## Solution

```go
func (m *ShardedMap) Set(key string, val int) {
	s := m.getShard(key)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

func (m *ShardedMap) Get(key string) (int, bool) {
	s := m.getShard(key)
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok
}
```

## Walkthrough

FNV-1a hashes the key to a `uint32`; `% m.num` picks the shard. Because the
function is deterministic, `Set("a", 1)` and `Get("a")` always land on the same
shard — which is the invariant the whole design depends on.

In the 1000-key test, keys spread across all 32 shards, so most writers hold
different locks and run genuinely in parallel. `Len()` sums each shard's size
under its own read lock and finds 1000.

## Pitfalls

- **Locking one shard and writing another.** Usually written as calling
  `getShard` twice with different keys; the race detector may not even see it if
  the shards differ, but the data is unguarded.
- **`Lock` in `Get`.** Correct but wasteful — concurrent readers of the same
  shard now serialize.
- **A map-wide mutex "just to be safe".** Undoes the sharding entirely.
- **`% m.num` on a negative int.** `h.Sum32()` is unsigned, so the conversion is
  safe here; converting a signed hash without masking can produce a negative
  index and a panic.

## Choosing the shard count

Powers of two let `& (num-1)` replace the modulo, and a shard count a few times
the expected parallelism is usually enough. Beyond that you are paying for cache
lines — each `sync.RWMutex` is 24 bytes, and shards that share a cache line
suffer false sharing, which is why high-performance implementations pad them.
