// Package batchergen — Gopher Workplace challenge.
package batchergen

// Batcher groups items of T into fixed-size batches.
// Use NewBatcher to create one.
type Batcher[T any] struct {
	buf  []T
	size int
}

// NewBatcher returns a batcher flushing every size items.
func NewBatcher[T any](size int) *Batcher[T] {
	// TODO(candidate): store the batch size.
	panic("not implemented")
}

// Add buffers v and returns a full batch when one is ready.
// It returns nil and false while the batch is still filling.
func (b *Batcher[T]) Add(v T) ([]T, bool) {
	// TODO(candidate): buffer v, returning a full batch when it is complete.
	panic("not implemented")
}

// Flush returns the buffered items, if any, and clears the buffer.
func (b *Batcher[T]) Flush() ([]T, bool) {
	// TODO(candidate): return whatever is buffered and clear it.
	panic("not implemented")
}
