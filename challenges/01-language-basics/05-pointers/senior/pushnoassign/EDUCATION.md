# Reassigning appended slice fields

## Intuition

A struct's slice field must be updated with the append result; ignoring it keeps the old, shorter header.

## Approach

1. `append` returns a new slice header; the bug discards it with `_ =`.
2. Store it back: `s.data = append(s.data, v)`.

## Solution

```go
type Stack struct{ data []int }

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}
```

## Walkthrough

Discarding `append`'s result leaves `s.data` empty despite pushes. Assigning it keeps the grown header, so `len` reflects the pushes.

## Pitfalls

- `_ = append(s.data, v)` builds nothing observable.
- Write `s.data = append(s.data, v)`.
