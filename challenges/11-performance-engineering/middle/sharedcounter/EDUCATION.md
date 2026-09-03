# Read, Modify, Write — Three Chances To Lose

## Intuition

The unit of atomicity has to be the whole update. Anything smaller leaves a window where another goroutine reads the value you are about to overwrite.

## Approach

1. Lock, lazily create the map, apply the delta, unlock.
2. `Get` and `Snapshot` take the same lock; `Snapshot` copies under it.

## Solution

```go
func (c *Counter) Add(key string, delta int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.n == nil {
		c.n = make(map[string]int64)
	}
	c.n[key] += delta
}

func (c *Counter) Get(key string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n[key]
}

func (c *Counter) Snapshot() map[string]int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make(map[string]int64, len(c.n))
	for k, v := range c.n {
		out[k] = v
	}
	return out
}
```

## Walkthrough

The explicit copy loop rather than `maps.Clone` is deliberate: `maps.Clone(nil)` returns nil, which would break the empty-but-non-nil contract for a counter nobody has touched yet.

## Pitfalls

- Returning `c.n` from `Snapshot`, handing the caller a map another goroutine is writing.
- Locking inside `Add` only around the map write, leaving the read unprotected.
- Using `sync.Map` reflexively; it is for read-mostly disjoint keys, not for contended counters.
