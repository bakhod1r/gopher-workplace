# Heap That Loses Its Order

## Intuition

Comparing the right child with `h.items[i]` rather than with `h.items[small]` can pick the larger of the two children, leaving the smaller one above it and breaking the ordering.

## Approach

1. Track `small` as the index of the smallest of the node and its children.
2. Compare each child against `h.items[small]`, not against the original parent.
3. Swap and descend while `small` moves.

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

With children `2` and `3` under a parent `5`, the buggy comparison can select the right child `3`, leaving `2` buried below it.

## Pitfalls

- Comparing children against the parent instead of the running minimum.
- Forgetting the bounds check on the right child.
- Testing only with already-ordered input, which never exercises the branch.
