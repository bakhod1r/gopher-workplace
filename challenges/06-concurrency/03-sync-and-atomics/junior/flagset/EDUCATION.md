# Feature Flag Set

## Intuition

An `RWMutex` has two modes. Many readers may hold it at once because reads do not conflict. A writer needs it alone, so it waits for the readers to drain and blocks new ones until the flag push is done.

## Approach

1. `Set` write-locks and assigns the map entry.
2. `Enabled` read-locks and returns `f.flags[name]` (missing means false).
3. `Len` read-locks and returns `len(f.flags)`.

## Solution

```go
// Package flagset - Gopher Workplace challenge.
package flagset

import "sync"

// FlagSet holds feature flags: read on every request, written rarely.
type FlagSet struct {
	mu    sync.RWMutex
	flags map[string]bool
}

// NewFlagSet returns an empty flag set.
func NewFlagSet() *FlagSet {
	return &FlagSet{flags: make(map[string]bool)}
}

// Set records the state of a feature flag.
//
// Examples:
//
//	f.Set("new_ui", true); f.Enabled("new_ui") => true
func (f *FlagSet) Set(name string, on bool) {
	f.mu.Lock()
	f.flags[name] = on
	f.mu.Unlock()
}

// Enabled reports whether the named flag is on. Unknown flags are off.
//
// Examples:
//
//	f.Enabled("unknown") => false
func (f *FlagSet) Enabled(name string) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.flags[name]
}

// Len returns the number of configured flags.
//
// Examples:
//
//	f.Set("a", true); f.Set("b", false); f.Len() => 2
func (f *FlagSet) Len() int {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return len(f.flags)
}
```

## Walkthrough

Twenty request goroutines call `Enabled` concurrently and all hold the read lock together. An operator's `Set` then waits for them to finish, takes the write lock alone, and updates the map.

## Pitfalls

- Calling `RUnlock` on a lock taken with `Lock` - that corrupts the lock state.
- Taking only `RLock` in `Set`; readers do not exclude each other, so two writers could collide.
- Holding the write lock while doing slow work, which starves every request.
