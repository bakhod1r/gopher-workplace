# The Minimum That Goes Stale

## Intuition

Popping the value stack without popping the minimum stack leaves a minimum belonging to a removed element.

## Approach

1. Pop the value stack.
2. If the popped value equals the current minimum, pop the minimum stack too.

## Solution

```go
func (s *MinStack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	if len(s.mins) > 0 && s.mins[len(s.mins)-1] == v {
		s.mins = s.mins[:len(s.mins)-1]
	}
	return v, true
}

func (s *MinStack[T]) Push(v T) {
	s.items = append(s.items, v)
	if len(s.mins) == 0 || v <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, v)
	}
}

func (s *MinStack[T]) Min() (T, bool) {
	if len(s.mins) == 0 {
		var zero T
		return zero, false
	}
	return s.mins[len(s.mins)-1], true
}
```

## Walkthrough

`Push(3); Push(1); Pop()` leaves `mins` as `[3 1]`, so `Min` answers 1 though only 3 remains.

## Pitfalls

- Comparing against `s.mins[0]` instead of the top.
- Pushing to `mins` only on a strict decrease, which breaks on duplicate minima.
