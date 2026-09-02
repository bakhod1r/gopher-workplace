# In-Flight Request Gauge

## Intuition

Every increment must be matched by exactly one decrement. Because both directions are indivisible atomic adds, the level is correct no matter how the handlers interleave — and it returns to zero once every request has finished.

## Approach

1. Declare `n atomic.Int64`.
2. `Enter` adds 1; `Exit` adds -1.
3. `Current` loads the value.

## Solution

```go
// Package inflight - Gopher Workplace challenge.
package inflight

import "sync/atomic"

// Gauge tracks how many requests are currently in flight.
type Gauge struct {
	n atomic.Int64
}

// Enter records that a request started.
//
// Examples:
//
//	var g Gauge; g.Enter(); g.Current()            => 1
//	var g Gauge; g.Enter(); g.Enter(); g.Current() => 2
func (g *Gauge) Enter() {
	g.n.Add(1)
}

// Exit records that a request finished.
//
// Examples:
//
//	var g Gauge; g.Enter(); g.Exit(); g.Current() => 0
func (g *Gauge) Exit() {
	g.n.Add(-1)
}

// Current returns the number of requests in flight.
func (g *Gauge) Current() int64 {
	return g.n.Load()
}
```

## Walkthrough

Sixteen handlers each `Enter`, do work, and `Exit`. During the burst the gauge fluctuates, but each of the 16 adds is paired with exactly one subtract, so once all handlers return `Current` is exactly 0.

## Pitfalls

- Using `Add(1)` on entry and `Store(0)` on exit, which wipes out other in-flight requests.
- Forgetting `defer g.Exit()` on an early-return path, leaking the count upward forever.
- Reading the field directly instead of via `Load`.
