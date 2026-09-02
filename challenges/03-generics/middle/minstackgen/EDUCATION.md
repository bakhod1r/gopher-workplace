# Stack With Minimum

## Intuition

Pushing unconditionally keeps the two stacks the same height, which is what makes `Pop` a symmetric one-line shrink on both.

## Approach

1. `Push`: append the value, then append the smaller of `v` and the current minimum.
2. `Pop`: shrink both stacks.
3. `Min`: read the top of the minima stack.

## Solution

```go
func (s *MinStack[T]) Push(v T) {
	s.items = append(s.items, v)
	if len(s.mins) == 0 || v < s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, v)
	} else {
		s.mins = append(s.mins, s.mins[len(s.mins)-1])
	}
}

func (s *MinStack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.mins = s.mins[:len(s.mins)-1]
	return top, true
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

`Push(3); Push(1)` records minima `[3 1]`; popping restores `3` as the minimum immediately.

## Pitfalls

- Pushing to `mins` only when a new minimum appears, which desynchronises the stacks.
- Scanning `items` in `Min`, which is O(n).
- Forgetting to shrink `mins` in `Pop`.
