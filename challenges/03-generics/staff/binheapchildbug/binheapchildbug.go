// Package binheapchildbug — Gopher Workplace challenge.
package binheapchildbug

import (
	"cmp"
)

// Heap is a binary min-heap over ordered values.
type Heap[T cmp.Ordered] struct {
	items []T
}

// Len reports how many elements the heap holds.
func (h *Heap[T]) Len() int {
	return len(h.items)
}

// siftDown restores the heap property downward from index i.
// It repeatedly swaps i with its smallest child.
func (h *Heap[T]) siftDown(i int) {
	// CHANGE CODE BELOW THIS LINE
	n := len(h.items)
	for {
		l := 2*i + 1
		if l >= n {
			return
		}
		c := l
		r := l + 1
		if r < n && h.items[r] < h.items[i] {
			c = r
		}
		if h.items[i] <= h.items[c] {
			return
		}
		h.items[i], h.items[c] = h.items[c], h.items[i]
		i = c
	}
	// CHANGE CODE ABOVE THIS LINE
}

// Push adds v to the heap.
func (h *Heap[T]) Push(v T) {
	h.items = append(h.items, v)
	i := len(h.items) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.items[p] <= h.items[i] {
			break
		}
		h.items[p], h.items[i] = h.items[i], h.items[p]
		i = p
	}
}

// Pop removes and returns the smallest element and true.
// It returns the zero value and false when the heap is empty.
func (h *Heap[T]) Pop() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	top := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]
	h.siftDown(0)
	return top, true
}
