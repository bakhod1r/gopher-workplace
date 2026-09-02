# Fill

## Intuition

`for i := 0; i < n; i++` already runs zero times when `n <= 0`, so the only care needed is not handing a negative number to `make`.

## Approach

1. Allocate `out` with capacity `max(n, 0)`.
2. Append `v` exactly `n` times.
3. Return `out`.

## Solution

```go
func Fill[T any](n int, v T) []T {
	out := make([]T, 0, max(n, 0))
	for i := 0; i < n; i++ {
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`Fill(0, 7)` allocates an empty slice and skips the loop entirely, returning `[]int{}` rather than `nil`.

## Pitfalls

- Calling `make([]T, n)` with a negative `n`, which panics.
- Using `make([]T, n)` and then appending, which produces `2n` elements.
- Returning `nil` for `n == 0` when an empty slice is expected.
