# Remove That Empties The Bin

## Intuition

`delete` removes the whole bin, so eleven remaining occurrences vanish in a single call, while `n` is only decremented once — the total and the per-key counts immediately disagree.

## Approach

1. Read the current count with the comma-ok form and bail out when absent.
2. Delete the key when this was the last occurrence; otherwise store `c-1`.
3. Decrement the running total exactly once.

## Solution

```go
func (s *Multiset[T]) Remove(v T) bool {
	c, ok := s.m[v]
	if !ok {
		return false
	}
	if c <= 1 {
		delete(s.m, v)
	} else {
		s.m[v] = c - 1
	}
	s.n--
	return true
}

func (s *Multiset[T]) Add(v T) {
	if s.m == nil {
		s.m = make(map[T]int)
	}
	s.m[v]++
	s.n++
}

func (s *Multiset[T]) Count(v T) int {
	return s.m[v]
}

func (s *Multiset[T]) Len() int {
	return s.n
}

func (s *Multiset[T]) Distinct() int {
	return len(s.m)
}
```

## Walkthrough

`Add(x)` three times then one `Remove(x)` leaves `Count(x) == 0` and `Distinct() == 0`, though `Len()` still claims 2.

## Pitfalls

- Writing `s.m[v] = c - 1` unconditionally, which leaves zero-count keys inflating `Distinct`.
- Decrementing `n` even on the not-present path.
- Relying on `s.m[v]` alone for presence — a stored zero is indistinguishable from a missing key.
