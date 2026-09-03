# Make Room Before The Appends

## Intuition

Growth is only expensive when it is a surprise. Once you can say how much more is coming, one allocation covers all of it and the appends after it are pure writes.

## Approach

1. Clamp `n` at 0.
2. If `cap(s)-len(s) >= n`, return `s`.
3. Otherwise allocate `make([]int, len(s), len(s)+n)`, copy, and return it.

## Solution

```go
// Grow returns s with capacity for at least n more elements, without
// changing its length or contents.
//
// If s already has the room, it is returned untouched and nothing is
// allocated. n < 0 is treated as 0.
//
// Examples:
//
// 	Grow(make([]int, 2, 2), 8) => length 2, capacity at least 10
func Grow(s []int, n int) []int {
	if n < 0 {
		n = 0
	}
	if cap(s)-len(s) >= n {
		return s
	}
	out := make([]int, len(s), len(s)+n)
	copy(out, s)
	return out
}
```

## Walkthrough

A slice of len 2, cap 2 needs room for 8: spare is 0, so a new array of cap 10 is allocated and the two elements copied. A slice of len 1, cap 32 needs room for 4: spare is 31, so it is returned as is.

## Pitfalls

- `make([]int, len(s)+n)` — that changes the length, not just the capacity.
- Always reallocating; the no-op case is what the test measures.
