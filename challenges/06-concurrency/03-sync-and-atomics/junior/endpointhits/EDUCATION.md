# Per-Route Hit Counter

## Intuition

Incrementing a map entry loads, adds and stores. Splitting that across two lock holds lets another handler slip in between and lose a hit, so the whole operation goes in one critical section.

## Approach

1. `NewHitCounter` makes the map.
2. `Record` locks, does `h.hits[route]++`, unlocks.
3. `Hits` locks, returns `h.hits[route]`.

## Solution

```go
// Package endpointhits - Gopher Workplace challenge.
package endpointhits

import "sync"

// HitCounter counts gateway requests per route.
type HitCounter struct {
	mu   sync.Mutex
	hits map[string]int
}

// NewHitCounter returns an empty per-route counter.
func NewHitCounter() *HitCounter {
	return &HitCounter{hits: make(map[string]int)}
}

// Record counts one request served by route.
//
// Examples:
//
//	h.Record("/users"); h.Hits("/users")                     => 1
//	h.Record("/users"); h.Record("/users"); h.Hits("/users") => 2
func (h *HitCounter) Record(route string) {
	h.mu.Lock()
	h.hits[route]++
	h.mu.Unlock()
}

// Hits returns the number of requests served by route.
//
// Examples:
//
//	NewHitCounter().Hits("/orders") => 0
func (h *HitCounter) Hits(route string) int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.hits[route]
}
```

## Walkthrough

Three handler goroutines each serve 100 requests to `/users`. Each `hits["/users"]++` happens under the lock, so the increments queue up and the final count is exactly 300.

## Pitfalls

- Reading the count, unlocking, then locking again to write - the value can change in between.
- Special-casing a missing route; Go maps return the zero value already.
- Forgetting to lock `Hits`.
