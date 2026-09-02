# Clamp All

## Intuition

The per-element logic is identical to a scalar clamp; the only new decision is producing a new slice so the caller's samples stay intact.

## Approach

1. Allocate `out` with capacity `len(s)`.
2. Append `lo`, `hi`, or `v` per element.
3. Return `out`.

## Solution

```go
func ClampAll[T cmp.Ordered](s []T, lo, hi T) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		switch {
		case v < lo:
			out = append(out, lo)
		case v > hi:
			out = append(out, hi)
		default:
			out = append(out, v)
		}
	}
	return out
}
```

## Walkthrough

`ClampAll([]int{-1, 2, 9}, 0, 3)` maps `-1` up to `0`, keeps `2`, and pulls `9` down to `3`.

## Pitfalls

- Writing back into `s[i]`, which mutates the caller's data.
- Dropping out-of-range elements instead of clamping them.
- Checking the high bound before the low bound with inverted bounds, which changes the result.
