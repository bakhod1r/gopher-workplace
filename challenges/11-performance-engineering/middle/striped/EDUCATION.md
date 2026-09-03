# Splitting The Lock

## Intuition

One lock means everyone queues. Sixteen locks mean a goroutine only queues behind the roughly one-sixteenth of traffic that hashes to its shard.

## Approach

1. `New` allocates the shards, each with its own map, plus one hash seed.
2. `ShardOf` hashes the key and takes it modulo the shard count.
3. `Set` and `Get` lock only their shard; `Len` walks them all under read locks.

## Solution

```go
func New(n int) *Map {
	if n < 1 {
		n = 1
	}
	m := &Map{shards: make([]shard, n), seed: maphash.MakeSeed()}
	for i := range m.shards {
		m.shards[i].m = make(map[string]int)
	}
	return m
}

func (m *Map) Shards() int { return len(m.shards) }

func (m *Map) ShardOf(key string) int {
	return int(maphash.String(m.seed, key) % uint64(len(m.shards)))
}

func (m *Map) Set(key string, v int) {
	s := &m.shards[m.ShardOf(key)]
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = v
}

func (m *Map) Get(key string) (int, bool) {
	s := &m.shards[m.ShardOf(key)]
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func (m *Map) Len() int {
	total := 0
	for i := range m.shards {
		m.shards[i].mu.RLock()
		total += len(m.shards[i].m)
		m.shards[i].mu.RUnlock()
	}
	return total
}
```

## Walkthrough

`s := &m.shards[...]` takes a pointer deliberately: `s := m.shards[...]` would copy the shard, mutex and all, locking a copy while the real map stays unprotected — and `go vet` flags exactly that.

## Pitfalls

- Copying a shard out of the slice, so the lock protects nothing.
- Seeding the hash per call, which sends the same key to different shards on different lookups.
- Offering a `Len` that pretends to be a consistent snapshot; it is a sum of moments, not a moment.
