// Package queuegen — Gopher Workplace challenge.
package queuegen

// Queue is a first-in-first-out collection of T.
// Its zero value is an empty queue.
type Queue[T any] struct {
	items []T
}

// Enqueue adds v to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	// TODO(candidate): append v to the items.
	panic("not implemented")
}

// Dequeue removes and returns the front element and true.
// It returns the zero value and false when the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	// TODO(candidate): remove and return the first item.
	panic("not implemented")
}

// Len returns the number of queued elements.
func (q *Queue[T]) Len() int {
	// TODO(candidate): report how many items are stored.
	panic("not implemented")
}
