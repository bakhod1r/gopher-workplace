# Rotation That Panics On Negatives

## Intuition

Go's remainder keeps the sign, so `k` stays negative and `s[k:]` panics instead of wrapping to the right.

## Approach

1. Normalise `k` with the double-modulus form.
2. Append the tail then the head.

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

`Rotate([]int{1,2,3}, -1)` needs `k = 2`; a bare modulus leaves `-1` and the slice expression panics.

## Pitfalls

- Assuming Go's `%` is a mathematical modulus.
- Adding a special case for negatives instead of normalising once.
- Testing only non-negative offsets.
