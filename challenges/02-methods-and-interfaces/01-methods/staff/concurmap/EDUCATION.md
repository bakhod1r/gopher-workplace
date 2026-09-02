# Concurrent Map

## Intuition

Go's map has no internal locking on purpose: most maps are goroutine-local, and
paying for synchronization everywhere would be wasteful. When a map really is
shared, you add exactly the discipline you need — here, a reader/writer lock,
because reads are expected to dominate.

## Approach

1. Guard the read path with `RLock`/`RUnlock`.
2. Guard the write path with `Lock`/`Unlock`.
3. Defer the unlock in both so no return path can leak the lock.

## Solution

```go
func (m *ConcurrentMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *ConcurrentMap) Set(key string, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}
```

## Walkthrough

100 goroutines interleave `Set` and `Get` on one key. Every write is exclusive,
so the runtime never sees two writers; every read either waits for a write to
finish or shares the lock with other readers. After `wg.Wait()` the key exists
with *some* value from 0 to 99 — the test only asserts existence, because the
winner is genuinely unspecified.

## Pitfalls

- **Reading without any lock.** `fatal error: concurrent map read and map write`
  — not a recoverable panic; the process dies.
- **`RLock` in `Set`.** Compiles, and lets two writers in at once. This is the
  dangerous bug: it usually looks fine until it crashes in production.
- **Forgetting `defer` and returning early.** The lock is held forever and every
  later call deadlocks.
- **Copying the struct.** `sync.RWMutex` must not be copied; `go vet` catches it.

## When to reach for `sync.Map` instead

`sync.Map` is optimized for two narrow cases: keys written once and read many
times, or disjoint key sets per goroutine. For general read/write traffic — like
this test's — a plain map behind an `RWMutex` is usually faster and always
easier to reason about.
