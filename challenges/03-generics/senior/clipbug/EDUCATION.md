# A View That Lets Callers Write

## Intuition

Slicing narrows the length but keeps the capacity, so the caller's first `append` reuses your buffer's next cell rather than allocating.

## Approach

1. Clamp `n`.
2. Return the prefix with its capacity clipped to its length.

## Solution

```go
func Head[T any](s []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	return slices.Clip(s[:n])
}
```

## Walkthrough

`append(Head(s, 2), x)` writes `x` into `s[2]` when the capacity is not clipped.

## Pitfalls

- Returning raw sub-slices from any exported API.
- Copying when clipping is enough — correct but needlessly expensive.
- Assuming `Clip` copies the data: it only narrows the capacity.
