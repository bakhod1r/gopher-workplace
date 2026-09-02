# Generic Stack

## Intuition

A slice-backed stack needs no initialisation: appending to a nil slice works. The receiver must be a pointer, or `Push` would grow a copy and drop it.

## Approach

1. `Push`: append to `s.items`.
2. `Pop`: guard the empty case, read the last element, reslice to drop it.
3. `Len`: return `len(s.items)`.

## Solution

```go
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
```

## Walkthrough

`Push(1); Push(2); Pop()` reslices `items` from length 2 to 1 and hands back `2`.

## Pitfalls

- Using a value receiver, so pushes are lost.
- Popping from the front, which makes it a queue.
- Forgetting to shrink `items`, so `Pop` keeps returning the same element.
