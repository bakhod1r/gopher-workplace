# Rotate

## Intuition

Once `k` is normalised, the rotation is just two appends — the tail followed by the head — which also gives a fresh backing array for free.

## Approach

1. Return empty for an empty input.
2. Normalise `k` with the double-modulus formula.
3. Append `s[k:]` then `s[:k]`.

## Solution

```go
func Rotate[T any](s []T, k int) []T {
	out := make([]T, 0, len(s))
	if len(s) == 0 {
		return out
	}
	k = ((k % len(s)) + len(s)) % len(s)
	out = append(out, s[k:]...)
	out = append(out, s[:k]...)
	return out
}
```

## Walkthrough

`Rotate([]int{1,2,3}, -1)` normalises `k` to `2`, appending `[3]` then `[1 2]`.

## Pitfalls

- Using `k % len(s)` alone and indexing with a negative result.
- Dividing by zero (or panicking) on an empty slice.
- Rotating in place and disturbing the caller's slice.
