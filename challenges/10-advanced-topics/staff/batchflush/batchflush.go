// Package batchflush — Gopher Workplace challenge.
package batchflush

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
//	b := NewBatcher(2, sink); b.Add(1); b.Add(2) => sink received [1 2]
func (b *Batcher) Add(v int) error {
	panic("not implemented")
}
