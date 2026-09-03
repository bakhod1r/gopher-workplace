# Three Counters, Three Costs

## Intuition

Contention is about how many cores want the same memory. A lock makes them queue, an atomic makes them fight over a cache line, and sharding gives them each their own.

## Approach

1. Mutex version: lock, increment, unlock.
2. Atomic version: `Add(1)` and `Load`.
3. Sharded version: index a slice of atomics by `id`, and sum on read.

## Solution

```go
func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func (c *AtomicCounter) Inc() { c.n.Add(1) }

func (c *AtomicCounter) Value() int64 { return c.n.Load() }

func NewSharded(n int) *ShardedCounter {
	if n < 1 {
		n = 1
	}
	return &ShardedCounter{shards: make([]atomic.Int64, n)}
}

func (c *ShardedCounter) Inc(id int) {
	n := len(c.shards)
	i := ((id % n) + n) % n
	c.shards[i].Add(1)
}

func (c *ShardedCounter) Value() int64 {
	var total int64
	for i := range c.shards {
		total += c.shards[i].Load()
	}
	return total
}
```

## Walkthrough

`c.shards[i].Load()` indexes rather than ranging by value: `atomic.Int64` contains a `noCopy` marker, and copying one out of the slice would both be flagged by vet and defeat the point.

## Pitfalls

- `id % n` alone, which is negative for a negative `id` and panics on the index.
- Ranging `for _, s := range c.shards`, which copies each atomic.
- Treating a sharded `Value` as exact while writers are running; it is a snapshot of a moving target.
