// Package heapsiftbug — Gopher Workplace challenge.
package heapsiftbug

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
	h.items = append(h.items, v)
	i := len(h.items) - 1
	for i > 0 {
		p := (i - 1) / 2
		if !(h.items[i] < h.items[p]) {
			break
		}
		h.items[i], h.items[p] = h.items[p], h.items[i]
		i = p
	}
}

// Pop removes and returns the smallest element and true.
// It returns the zero value and false when the heap is empty.
func (h *Heap[T]) Pop() (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	top := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]
	i := 0
	for {
		l, r := 2*i+1, 2*i+2
		small := i
		if l < len(h.items) && h.items[l] < h.items[small] {
			small = l
		}
		if r < len(h.items) && h.items[r] < h.items[i] {
			small = r
		}
		if small == i {
			break
		}
		h.items[i], h.items[small] = h.items[small], h.items[i]
		i = small
	}
	return top, true
	// CHANGE CODE ABOVE THIS LINE
}

// Len returns the number of stored elements. It is provided for you.
func (h *Heap[T]) Len() int {
	return len(h.items)
}
