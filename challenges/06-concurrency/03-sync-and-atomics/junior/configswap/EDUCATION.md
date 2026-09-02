# Hot Config Swap

## Intuition

Updating two fields of a shared struct in place lets a reader catch it half-updated. `atomic.Value` swaps a pointer to an immutable snapshot instead: readers get whichever complete snapshot was current when they looked.

## Approach

1. Hold an `atomic.Value`.
2. `Store` calls `s.v.Store(c)` with the whole `Config`.
3. `Load` handles the `nil` case, then type-asserts to `Config`.

## Solution

```go
// Package configswap - Gopher Workplace challenge.
package configswap

import "sync/atomic"

// Config is a service configuration snapshot.
type Config struct {
	Version int
	Region  string
}

// Store publishes configuration snapshots to concurrent readers.
type Store struct {
	v atomic.Value
}

// Store publishes c as the current configuration.
//
// Examples:
//
//	var s Store; s.Store(Config{Version: 2, Region: "eu"}); s.Load().Region => "eu"
func (s *Store) Store(c Config) {
	s.v.Store(c)
}

// Load returns the current configuration, or the zero Config if none.
//
// Examples:
//
//	var s Store; s.Load() => Config{}
func (s *Store) Load() Config {
	v := s.v.Load()
	if v == nil {
		return Config{}
	}
	return v.(Config)
}

// Version returns the published configuration version.
//
// Examples:
//
//	s.Store(Config{Version: 3}); s.Version() => 3
func (s *Store) Version() int {
	return s.Load().Version
}
```

## Walkthrough

A reload stores `Config{3, "eu"}` while a handler loads. The handler sees either the entire old config or the entire new one — the two fields can never be torn apart, because a single pointer is swapped.

## Pitfalls

- Type-asserting without the `nil` check, which panics before the first `Store`.
- Storing a `Config` sometimes and a `*Config` other times — `atomic.Value` panics on an inconsistent type.
- Mutating a stored `Config` afterwards instead of storing a fresh snapshot.
