# Latency Peak

## Intuition

CAS says "write ms only if the value is still cur". If another handler changed the peak in between, the CAS fails and you retry with the fresh value. That is how lock-free read-modify-write works.

## Approach

1. Read the current peak.
2. If the new latency is not larger, stop.
3. Try `CompareAndSwap(cur, ms)`; if it fails, loop and re-read.

## Solution

```go
// Package latencypeak - Gopher Workplace challenge.
package latencypeak

import "sync/atomic"

// PeakTracker records the slowest request latency observed, in milliseconds.
type PeakTracker struct {
	ms atomic.Int64
}

// Observe raises the tracked peak to ms when ms is larger.
//
// Examples:
//
//	var p PeakTracker; p.Observe(5); p.Peak()               => 5
//	var p PeakTracker; p.Observe(5); p.Observe(3); p.Peak() => 5
func (p *PeakTracker) Observe(ms int64) {
	for {
		cur := p.ms.Load()
		if ms <= cur {
			return
		}
		if p.ms.CompareAndSwap(cur, ms) {
			return
		}
	}
}

// Peak returns the slowest observed latency, or 0 if none.
//
// Examples:
//
//	var p PeakTracker; p.Peak() => 0
func (p *PeakTracker) Peak() int64 {
	return p.ms.Load()
}
```

## Walkthrough

Two handlers report 5 ms and 9 ms with peak 0. Both read 0; the first CASes to 5. The second's CAS(0, 9) fails, it re-reads 5, 9 is still larger, CAS(5, 9) succeeds. Peak: 9.

## Pitfalls

- Doing `if ms > p.ms.Load() { p.ms.Store(ms) }` - the check and the store are separate, so a peak can be lost.
- Forgetting to re-read inside the loop, spinning forever on a stale value.
- Not returning on the `ms <= cur` path, which loops forever.
