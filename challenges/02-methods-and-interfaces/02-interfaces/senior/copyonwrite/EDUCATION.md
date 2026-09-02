# Copy On Write

## Intuition

Copy-on-write inverts the usual cost: readers do a single atomic load with no lock, and the rare writer pays for a full copy. It only works if published snapshots are treated as immutable.

## Approach

1. `Load` type-asserts the value out of `atomic.Value`.
2. `Store` publishes a snapshot with one atomic write.
3. `Update` takes the writer mutex, copies the current map, applies the mutator to the copy, and stores it.
4. The writer lock serialises concurrent updates so no update is lost.

## Solution

```go
func (c *Config) Load() Snapshot { return c.v.Load().(Snapshot) }

func (c *Config) Store(s Snapshot) { c.v.Store(s) }

func (c *Config) Update(m Mutator) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cur := c.Load()
	next := make(Snapshot, len(cur)+1)
	for k, v := range cur {
		next[k] = v
	}
	m.Mutate(next)
	c.Store(next)
}
```

## Walkthrough

`TestOldSnapshotUnchanged` holds a snapshot across an update. Because `Update` built a copy, the old map is still exactly what the reader saw — the guarantee readers rely on.

## Pitfalls

- Mutating `c.Load()` directly, which races with every reader holding that map.
- Dropping the writer mutex: two concurrent updates read the same base and one loses.
- Storing different dynamic types into one `atomic.Value`, which panics at runtime.
