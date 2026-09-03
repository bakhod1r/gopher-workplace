# Sliding Window Runs Off The End

## Intuition

Running `i` to the end and clamping the window emits short tail windows, which downstream code averages as if they were full.

## Approach

1. Stop the loop while `i+n <= len(s)`.
2. Copy each full window into its own slice.

## Solution

```go
func Windows[T any](s []T, n int) [][]T {
	out := make([][]T, 0)
	if n <= 0 || n > len(s) {
		return out
	}
	for i := 0; i+n <= len(s); i++ {
		w := make([]T, n)
		copy(w, s[i:i+n])
		out = append(out, w)
	}
	return out
}
```

## Walkthrough

For `[1 2 3]` with `n = 2` the correct loop stops at `i = 1`; the buggy one also emits `[3]`.

## Pitfalls

- Clamping the window end instead of stopping the loop.
- Emitting partial windows and letting the caller filter them.
- Checking only the first window in tests.
