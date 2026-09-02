# Allocate Once

## Intuition

Using `make([]T, n)` and appending would produce `2n` elements; using no hint at all would reallocate roughly log₂(n) times.

## Approach

1. Clamp a negative `n`.
2. Allocate with length 0 and capacity `n`.
3. Append `f(i)` for each index.

## Solution

```go
func Build[T any](n int, f func(int) T) []T {
	if n < 0 {
		n = 0
	}
	out := make([]T, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, f(i))
	}
	return out
}
```

## Walkthrough

`Build(3, f)` reserves three slots up front, so all three appends write into the same array.

## Pitfalls

- Writing `make([]T, n)` and then appending.
- Omitting the capacity and letting append grow the slice.
- Passing a negative capacity to `make`, which panics.
