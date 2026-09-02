# Metrics Flush

## Intuition

Between a `Load` and a `Store(0)`, other goroutines keep recording; zeroing afterwards throws their bytes away. `Swap` performs both halves as one instruction, so nothing can slip into the gap — there is no gap.

## Approach

1. Declare `n atomic.Int64`.
2. `Record` calls `Add(v)`.
3. `Drain` returns `m.n.Swap(0)`; `Pending` returns `Load()`.

## Solution

```go
// Package metricsflush - Gopher Workplace challenge.
package metricsflush

import "sync/atomic"

// Meter accumulates a byte counter that is flushed periodically.
type Meter struct {
	n atomic.Int64
}

// Record adds v to the pending total.
//
// Examples:
//
//	var m Meter; m.Record(3); m.Record(4); m.Pending() => 7
func (m *Meter) Record(v int64) {
	m.n.Add(v)
}

// Drain returns the pending total and resets it to zero.
//
// Examples:
//
//	var m Meter; m.Record(3); m.Drain() => 3
//	var m Meter; m.Drain()              => 0
func (m *Meter) Drain() int64 {
	return m.n.Swap(0)
}

// Pending returns the current total without clearing it.
func (m *Meter) Pending() int64 {
	return m.n.Load()
}
```

## Walkthrough

A worker records 3 while the flusher drains. If the add lands first, `Drain` returns 3 and leaves 0; if it lands second, `Drain` returns 0 and the 3 stays pending for the next flush. Either way the byte count is never lost.

## Pitfalls

- `v := m.n.Load(); m.n.Store(0); return v` — loses every concurrent `Record`.
- Draining into a local and forgetting that later reads must go through `Load`.
- Assuming `Swap` returns the *new* value; it returns the old one.
