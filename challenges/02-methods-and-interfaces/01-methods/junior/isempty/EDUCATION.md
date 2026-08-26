# Checking State with Methods

## Intuition

`IsEmpty` is a common predicate method. In Go, `len(nil)` is 0, so you don't
need to nil-check the slice before checking its length.

## Approach

1. Check `len(s.items) == 0`.

## Solution

```go
func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}
```

## Walkthrough

For `Stack{}`:
- `s.items` is `nil`.
- `len(nil)` = 0.
- `0 == 0` is `true`.

## Pitfalls

- Checking `s.items == nil` misses the case where `items` is an empty slice
  (`[]int{}`).
- `len` works on nil slices — a Go design choice that simplifies code.
