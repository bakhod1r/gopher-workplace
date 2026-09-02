# Job Metrics

## Intuition

`n++` compiles to three steps: load, add, store. An atomic add is one indivisible machine instruction, so concurrent workers cannot interleave and lose each other's increments.

## Approach

1. Declare `processed atomic.Int64`.
2. `Add` calls `m.processed.Add(delta)`.
3. `Processed` returns `m.processed.Load()`; `Reset` calls `Store(0)`.

## Solution

```go
// Package jobmetrics - Gopher Workplace challenge.
package jobmetrics

import "sync/atomic"

// JobMetrics is a lock-free counter of jobs processed by a worker pool.
type JobMetrics struct {
	processed atomic.Int64
}

// Add adds delta to the processed-jobs counter.
//
// Examples:
//
//	var m JobMetrics; m.Add(3); m.Processed()            => 3
//	var m JobMetrics; m.Add(2); m.Add(-5); m.Processed() => -3
func (m *JobMetrics) Add(delta int64) {
	m.processed.Add(delta)
}

// Processed returns the current number of processed jobs.
func (m *JobMetrics) Processed() int64 {
	return m.processed.Load()
}

// Reset zeroes the counter after a metrics flush.
//
// Examples:
//
//	var m JobMetrics; m.Add(7); m.Reset(); m.Processed() => 0
func (m *JobMetrics) Reset() {
	m.processed.Store(0)
}
```

## Walkthrough

Four workers each call `Add(1)` a thousand times. Every add is indivisible, so the reported total is exactly 4000 - not 3987.

## Pitfalls

- Reading the field directly instead of via `Load` - that is a data race.
- Mixing atomic writes with plain reads of the same field.
- Copying the struct after use.
