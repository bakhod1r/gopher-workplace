# Metrics Counter Registry

## Intuition

Two different access patterns share one map: rare registration (a write) and constant increments (a read of the map, plus an atomic write inside the value). Split them — the RWMutex protects the map's shape, the atomic protects the number.

## Approach

1. `counter`: read-lock, look up, unlock; return it if found.
2. Otherwise write-lock, look up **again** (someone may have created it), create when still missing, return it.
3. `Add` calls `counter(name).Add(n)`.
4. `Value` and `Snapshot` read under `RLock`, loading each counter atomically.

## Solution

```go
func NewRegistry() *Registry {
	return &Registry{counters: make(map[string]*atomic.Int64)}
}

func (r *Registry) counter(name string) *atomic.Int64 {
	r.mu.RLock()
	c, ok := r.counters[name]
	r.mu.RUnlock()
	if ok {
		return c
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if c, ok := r.counters[name]; ok { // another goroutine won the race
		return c
	}
	c = &atomic.Int64{}
	r.counters[name] = c
	return c
}

func (r *Registry) Add(name string, n int64) int64 {
	return r.counter(name).Add(n)
}

func (r *Registry) Value(name string) int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.counters[name]
	if !ok {
		return 0
	}
	return c.Load()
}

func (r *Registry) Snapshot() map[string]int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[string]int64, len(r.counters))
	for name, c := range r.counters {
		out[name] = c.Load()
	}
	return out
}
```

## Walkthrough

250 goroutines bump `metric_3`. The first few miss under `RLock`, queue on `Lock`, and exactly one creates the counter — the rest hit the re-check and reuse it. From then on every increment takes only the read lock, so all eight metric names proceed in parallel and each ends at exactly 250.

## Pitfalls

- Skipping the re-check after upgrading to `Lock` — two goroutines create two counters and half the increments vanish.
- Storing `atomic.Int64` by value in the map: every read copies it and the increments are lost.
- Holding `RLock` while writing to the map — read locks do not exclude each other, and the map corrupts.
