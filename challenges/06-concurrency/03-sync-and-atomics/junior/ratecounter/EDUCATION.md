# Rate Counter

## Intuition

A mutex is a token: only the goroutine holding it may touch the guarded field. `hits++` is really load-add-store, so two request handlers can read the same value and both write back the same result, silently dropping a request from the metric.

## Approach

1. Lock the mutex.
2. Mutate or read `hits`.
3. Unlock (use `defer` when the method returns a value).

## Solution

```go
// Package ratecounter - Gopher Workplace challenge.
package ratecounter

import "sync"

// RateCounter counts gateway requests across many serving goroutines.
type RateCounter struct {
	mu   sync.Mutex
	hits int
}

// Record counts one inbound request.
//
// Examples:
//
//	var c RateCounter; c.Record(); c.Hits()             => 1
//	var c RateCounter; c.Record(); c.Record(); c.Hits() => 2
func (c *RateCounter) Record() {
	c.mu.Lock()
	c.hits++
	c.mu.Unlock()
}

// Hits returns the number of requests recorded so far.
//
// Examples:
//
//	var c RateCounter; c.Hits() => 0
func (c *RateCounter) Hits() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.hits
}
```

## Walkthrough

Two handlers call `Record`. The first takes the lock, does `hits++` (0 to 1), releases. The second blocks until then, sees 1, writes 2. No request is lost from the count.

## Pitfalls

- Locking in `Record` but not in `Hits` - a read racing a write is still a data race.
- Copying a `RateCounter` by value after use; that copies the mutex too.
- Forgetting `Unlock` on an early return - prefer `defer c.mu.Unlock()`.
