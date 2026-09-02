# Atomic Versus Mutex

## Intuition

Both counters are correct; they differ in what contention costs. A mutex parks goroutines and involves the scheduler; an atomic add is a single locked instruction that contends only on the cache line.

## Approach

1. `MutexCounter.Add` takes the lock, mutates, unlocks; `Inc` delegates to `Add`.
2. `Value` must also take the lock — an unsynchronised read is a race even for an `int64`.
3. `AtomicCounter` uses `atomic.Int64`'s `Add` and `Load` directly.
4. `IncAll` loops over the interface values.

## Solution

```go
func (c *MutexCounter) Inc() { c.Add(1) }

func (c *MutexCounter) Add(delta int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n += delta
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func (c *AtomicCounter) Inc() { c.n.Add(1) }

func (c *AtomicCounter) Add(delta int64) { c.n.Add(delta) }

func (c *AtomicCounter) Value() int64 { return c.n.Load() }
```

## Walkthrough

`TestNoLostUpdates` runs both implementations through the same 2000-goroutine gauntlet. A non-atomic `c.n++` would drop increments here — which is the difference between a slow counter and a wrong one.

## Pitfalls

- Reading `c.n` without the lock in `MutexCounter.Value` — `-race` flags it, and a torn read is possible on some architectures.
- `c.n.Store(c.n.Load() + delta)` in the atomic version, which reintroduces the read-modify-write gap.
- Copying a counter by value: both a `sync.Mutex` and an `atomic.Int64` must not be copied after first use.
