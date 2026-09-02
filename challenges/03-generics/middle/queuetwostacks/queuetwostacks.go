// Package queuetwostacks — Gopher Workplace challenge.
package queuetwostacks

// SQueue is a FIFO queue built from two stacks.
// Its zero value is an empty queue.
type SQueue[T any] struct {
	in  []T
	out []T
}

// Enqueue adds v to the back of the queue.
func (q *SQueue[T]) Enqueue(v T) {
	// TODO(candidate): push onto the inbound stack.
	panic("not implemented")
}

// Dequeue removes and returns the front element and true.
func (q *SQueue[T]) Dequeue() (T, bool) {
	// TODO(candidate): refill the outbound stack when empty, then pop it.
	panic("not implemented")
}

// Len returns the number of queued elements.
func (q *SQueue[T]) Len() int {
	// TODO(candidate): count both stacks.
	panic("not implemented")
}
