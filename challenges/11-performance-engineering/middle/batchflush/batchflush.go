// Package batchflush — Gopher Workplace challenge.
package batchflush

// Batcher accumulates items and hands them to Flush in groups of Size,
// amortising a fixed per-call cost — a syscall, a network round trip, a
// transaction — across many items.
type Batcher struct {
	Size  int               // items per batch; a non-positive Size batches 1
	Flush func(batch []int) // called with each full batch, and by Close
	buf   []int
	sent  int
}

// Add appends one item, flushing when the batch is full. The slice handed to
// Flush is the batcher's own buffer: it is valid only for the duration of the
// call, and callers who keep it must copy.
//
// Examples:
//
//	b.Add(1)
func (b *Batcher) Add(v int) {
	panic("not implemented")
}

// Close flushes a partial batch, if any. It is safe to call more than once,
// and must never call Flush with an empty batch.
//
// Examples:
//
//	b.Close()
func (b *Batcher) Close() {
	panic("not implemented")
}

// Flushes reports how many times Flush has been called.
//
// Examples:
//
//	b.Flushes() => 2
func (b *Batcher) Flushes() int {
	panic("not implemented")
}
