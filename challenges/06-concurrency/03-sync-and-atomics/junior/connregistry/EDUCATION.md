# Connection Registry

## Intuition

The lock only protects code that holds it. If you return the internal map or a slice that aliases it, callers read it after the lock is gone — the race just moved outside your type. Copying inside the lock keeps the guarantee.

## Approach

1. `Register` write-locks and assigns.
2. `Lookup` read-locks and returns comma-ok.
3. `IDs` read-locks and appends the keys into a fresh slice.

## Solution

```go
// Package connregistry - Gopher Workplace challenge.
package connregistry

import "sync"

// Registry maps upstream instance IDs to addresses.
type Registry struct {
	mu    sync.RWMutex
	conns map[string]string
}

// NewRegistry returns an empty registry.
func NewRegistry() *Registry {
	return &Registry{conns: make(map[string]string)}
}

// Register records the address of an upstream instance.
//
// Examples:
//
//	r.Register("a", "10.0.0.1"); r.Lookup("a") => "10.0.0.1", true
func (r *Registry) Register(id, addr string) {
	r.mu.Lock()
	r.conns[id] = addr
	r.mu.Unlock()
}

// Lookup returns the address of id and whether it is registered.
//
// Examples:
//
//	NewRegistry().Lookup("ghost") => "", false
func (r *Registry) Lookup(id string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	addr, ok := r.conns[id]
	return addr, ok
}

// IDs returns a copy of the registered instance IDs, in any order.
//
// Examples:
//
//	r.Register("b", "x"); r.Register("a", "y"); len(r.IDs()) => 2
func (r *Registry) IDs() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	ids := make([]string, 0, len(r.conns))
	for id := range r.conns {
		ids = append(ids, id)
	}
	return ids
}
```

## Walkthrough

Twenty routing goroutines call `Lookup` together under `RLock`. A discovery update then waits for them, takes `Lock` alone, and registers a new instance. A concurrent `IDs` returns its own slice, safe to sort and iterate after the lock is released.

## Pitfalls

- Returning the internal map, letting callers read it unlocked.
- Appending to a slice field and returning it, which still aliases shared storage.
- Using `Lock` for `Lookup`, which throws away the read parallelism.
