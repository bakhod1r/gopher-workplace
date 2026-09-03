# Publish A Snapshot Without Tearing It

## Intuition

Tearing comes from updating a shared value field by field. Replace the value instead: build the new snapshot privately, then swap one pointer. Readers either load the old address or the new one, and both point at a complete struct.

## Approach

1. Copy the parameter into a local so the store owns it.
2. `s.v.Store(&cp)` to publish the address.

## Solution

```go
import "sync/atomic"

// Config is one immutable settings snapshot.
type Config struct {
	Retries int
	Timeout int
}

// Store holds the current Config for concurrent readers.
type Store struct {
	v atomic.Pointer[Config]
}

// Get returns the current snapshot, or the zero Config if none is set.
func (s *Store) Get() Config {
	if p := s.v.Load(); p != nil {
		return *p
	}
	return Config{}
}

// Set publishes c as the current configuration.
//
// Readers must see either the old snapshot or the new one, never a mix.
// The snapshot escapes by construction — it outlives the call and is shared
// with every reader.
//
// Examples:
//
// 	s.Set(Config{Retries: 3}); s.Get().Retries => 3
func (s *Store) Set(c Config) {
	cp := c
	s.v.Store(&cp)
}
```

## Walkthrough

2000 concurrent updates with four readers spinning: every `Get` dereferences one of the 2001 complete snapshots. Writing `s.cfg.Retries` and `s.cfg.Timeout` separately would let a reader land between the two stores.

## Pitfalls

- `s.v.Store(&c)` publishes the parameter's address — legal, but then the caller's later writes are visible to readers.
- Mutating a snapshot after publishing it; the whole scheme depends on published values being immutable.
