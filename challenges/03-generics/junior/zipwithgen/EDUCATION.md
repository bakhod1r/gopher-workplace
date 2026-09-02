# Zip With

## Intuition

Zipping is only defined where both slices have an element. Taking the minimum length up front turns the bounds problem into a single loop with no per-iteration checks.

## Approach

1. Set `n` to the smaller of the two lengths.
2. Allocate `out` as `[]R` with capacity `n`.
3. Append `f(a[i], b[i])` for each `i` below `n`.

## Solution

```go
func ZipWith[T, U, R any](a []T, b []U, f func(T, U) R) []R {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]R, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, f(a[i], b[i]))
	}
	return out
}
```

## Walkthrough

`ZipWith([]int{1, 2, 3}, []int{10}, add)` sets `n = 1`, so only `add(1, 10)` runs and the result is `[11]`.

## Pitfalls

- Ranging over `a` and indexing `b`, which panics when `b` is shorter.
- Using the longer length and padding with zero values.
- Declaring one type parameter for both inputs, which rejects `[]string` zipped with `[]int`.
