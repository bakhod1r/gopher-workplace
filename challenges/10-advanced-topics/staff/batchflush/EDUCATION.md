# Flush At The Threshold, Not At The End

## Intuition

A size threshold is the only bound that holds when the consumer is slow. Flushing on the way in means the pending buffer's capacity is the batcher's entire memory footprint.

## Approach

1. Under the lock, append `v`.
2. Return nil when the batch is not yet full.
3. Copy the pending values into a fresh batch, reset `pending` to `[:0]`, and flush the copy.

## Solution

```go
import "sync"

// Batcher accumulates values and flushes them in fixed-size batches.
type Batcher struct {
	mu      sync.Mutex
	limit   int
	pending []int
	flush   func([]int) error
}

// NewBatcher returns a batcher that calls flush with each full batch.
func NewBatcher(limit int, flush func([]int) error) *Batcher {
	if limit < 1 {
		limit = 1
	}
	return &Batcher{limit: limit, pending: make([]int, 0, limit), flush: flush}
}

// Pending reports how many values are waiting.
func (b *Batcher) Pending() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.pending)
}

// Close flushes whatever is left.
func (b *Batcher) Close() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.pending) == 0 {
		return nil
	}
	batch := make([]int, len(b.pending))
	copy(batch, b.pending)
	b.pending = b.pending[:0]
	return b.flush(batch)
}

// Add appends v to the pending batch and flushes when the batch reaches
// its limit.
//
// The pending slice must never grow past the limit: an unbounded batcher
// turns a slow consumer into an out-of-memory kill.
//
// Examples:
//
// 	b := NewBatcher(2, sink); b.Add(1); b.Add(2) => sink received [1 2]
func (b *Batcher) Add(v int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.pending = append(b.pending, v)
	if len(b.pending) < b.limit {
		return nil
	}
	batch := make([]int, len(b.pending))
	copy(batch, b.pending)
	b.pending = b.pending[:0]
	return b.flush(batch)
}
```

## Walkthrough

With limit 2, the second `Add` copies [1 2] into a new slice, empties the pending buffer, and calls flush. The pending buffer keeps its capacity of 2 and is reused forever.

## Pitfalls

- Passing `b.pending` straight to `flush` — the next batch overwrites what the callee kept.
- `b.pending = nil` instead of `[:0]`, which throws away the capacity every batch.
- Flushing outside the lock without taking the batch out first, which lets a concurrent `Add` join the batch mid-flush.
