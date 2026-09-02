// Package heapgen — Gopher Workplace challenge.
package heapgen

import (
	"cmp"
)

// Heap is a binary min-heap of T.
// Its zero value is an empty heap.
type Heap[T cmp.Ordered] struct {
	items []T
}

// Push adds v to the heap.
func (h *Heap[T]) Push(v T) {
	// TODO(candidate): append, then sift the new element up.
	panic("not implemented")
}

// Pop removes and returns the smallest element and true.
// It returns the zero value and false when the heap is empty.
func (h *Heap[T]) Pop() (T, bool) {
	// TODO(candidate): take the root, move the last element up, then sift down.
	panic("not implemented")
}

// Len returns the number of stored elements.
func (h *Heap[T]) Len() int {
	// TODO(candidate): report how many elements are stored.
	panic("not implemented")
}
