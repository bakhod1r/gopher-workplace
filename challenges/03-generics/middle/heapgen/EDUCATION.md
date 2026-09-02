# Binary Min-Heap

## Intuition

The heap never sorts fully — it only maintains the parent-child ordering, which is exactly enough to know the minimum in O(1) and repair in O(log n).

## Approach

1. `Push`: append, then swap with the parent while smaller.
2. `Pop`: save the root, move the last element in, shrink, then swap with the smaller child while out of order.

## Solution

```go
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

func (h *Heap[T]) Pop() (T, bool) {
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
		if r < len(h.items) && h.items[r] < h.items[small] {
			small = r
		}
		if small == i {
			break
		}
		h.items[i], h.items[small] = h.items[small], h.items[i]
		i = small
	}
	return top, true
}

func (h *Heap[T]) Len() int {
	return len(h.items)
}
```

## Walkthrough

`Push(3); Push(1)` swaps `1` above `3`, so `Pop` returns `1` and leaves `3` at the root.

## Pitfalls

- Sifting down from index 1 instead of 0 after a pop.
- Forgetting to shrink the slice, leaving a stale copy of the last element.
- Comparing against the wrong child and breaking the invariant silently.
