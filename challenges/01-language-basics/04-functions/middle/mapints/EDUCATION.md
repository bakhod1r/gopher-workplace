# First-class and higher-order functions

## Intuition

Functions are ordinary values in Go; a function that takes or returns a function is higher-order, enabling map/filter/reduce.

## Approach

1. Range the input.
2. Append `f(v)` to a fresh slice.

## Solution

```go
func MapInts(xs []int, f func(int) int) []int {
	out := make([]int, 0, len(xs))
	for _, v := range xs {
		out = append(out, f(v))
	}
	return out
}
```

## Walkthrough

`MapInts([1 2 3], square)` applies square to each element → `[1 4 9]`.

## Pitfalls

- Return a new slice; don't overwrite the caller's `xs`.
- Preallocating capacity is optional.
