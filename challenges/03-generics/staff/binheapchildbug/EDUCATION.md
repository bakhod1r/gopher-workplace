# The Heap That Swaps With The Wrong Child

## Intuition

A min-heap must swap a too-large parent with its *smallest* child. Swapping with the larger one puts a value above a smaller sibling, so the invariant is broken one level down and the error migrates deeper on later operations.

## Approach

1. Compute the left child; stop if it does not exist.
2. Take the right child instead only when it is smaller than the left.
3. Stop if the parent already fits; otherwise swap and continue from the child.

## Solution

```go
func (h *Heap[T]) siftDown(i int) {
	n := len(h.items)
	for {
		l := 2*i + 1
		if l >= n {
			return
		}
		c := l
		r := l + 1
		if r < n && h.items[r] < h.items[l] {
			c = r
		}
		if h.items[i] <= h.items[c] {
			return
		}
		h.items[i], h.items[c] = h.items[c], h.items[i]
		i = c
	}
}

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
```

## Walkthrough

With items `[5, 1, 2]` the buggy predicate compares `2 < 5` and picks the right child, giving `[2, 1, 5]` — and `1` now sits under `2`.

## Pitfalls

- Comparing a child against the parent when deciding *which* child to use.
- Guarding only `l < n` and then reading `items[l+1]` out of range.
- Trusting a four-element test: shallow heaps rarely have two comparable children.
