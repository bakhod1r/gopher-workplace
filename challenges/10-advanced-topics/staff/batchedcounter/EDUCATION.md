# Accumulate Locally, Publish Rarely

## Intuition

An atomic increment is cheap in isolation and ruinous under contention, because every core wants the same cache line. Accumulating privately turns sixty-four coherence events into one.

## Approach

1. Add `n` to `local.n`.
2. If the magnitude has reached `batchSize`, `total.Add(local.n)` and zero the local.

## Solution

```go
import "sync/atomic"

// batchSize is how much a Local may accumulate before publishing.
const batchSize = 64

// Local is one goroutine's private accumulator. It must not be shared.
type Local struct {
	n int64
}

// Counter is a shared total fed by batched local accumulators.
type Counter struct {
	total atomic.Int64
}

// Flush publishes whatever the local still holds.
func (c *Counter) Flush(local *Local) {
	if local.n != 0 {
		c.total.Add(local.n)
		local.n = 0
	}
}

// Total returns the published total.
func (c *Counter) Total() int64 { return c.total.Load() }

// Add adds n to the caller's local accumulator, flushing it into the
// shared total when it reaches the batch threshold.
//
// The shared atomic is the contended resource; touching it once per batch
// instead of once per event is the whole point.
//
// Examples:
//
// 	c.Add(local, 1) a thousand times => the total is 1000 after Flush
func (c *Counter) Add(local *Local, n int64) {
	local.n += n
	if local.n >= batchSize || local.n <= -batchSize {
		c.total.Add(local.n)
		local.n = 0
	}
}
```

## Walkthrough

Sixteen workers each accumulate to 64 before touching the shared counter, so a million increments cost about sixteen thousand atomic operations instead of a million.

## Pitfalls

- Checking only `>= batchSize`, so a decreasing counter never publishes.
- Sharing one `Local` between goroutines, which is a data race the type's documentation forbids.
- Forgetting `Flush`, which loses up to `batchSize-1` per goroutine.
