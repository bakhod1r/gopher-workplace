# Clamp

## Intuition

Clamping is two one-sided limits applied in sequence. Written this way it needs only `<` and `>`, so `cmp.Ordered` is exactly the right constraint.

## Approach

1. Return `lo` when `v < lo`.
2. Return `hi` when `v > hi`.
3. Otherwise return `v`.

## Solution

```go
func Clamp[T cmp.Ordered](v, lo, hi T) T {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
```

## Walkthrough

`Clamp(5, 0, 3)` passes the low check, fails the high check, and returns `3`.

## Pitfalls

- Returning `v` before checking the bounds.
- Assuming `lo <= hi` without saying so — with inverted bounds this returns `lo`.
- Using arithmetic like `min(max(...))` written by hand with the wrong nesting.
