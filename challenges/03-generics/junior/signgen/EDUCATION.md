# Sign

## Intuition

`Sign(v)` is `Compare(v, 0)` specialised. Keeping the constraint signed is what makes the negative branch reachable at all.

## Approach

1. Return `-1` when `v < 0`.
2. Return `1` when `v > 0`.
3. Otherwise return `0`.

## Solution

```go
func Sign[T Signed](v T) int {
	switch {
	case v < 0:
		return -1
	case v > 0:
		return 1
	default:
		return 0
	}
}
```

## Walkthrough

`Sign(1.5)` instantiates `T = float64`, skips the negative case, matches `v > 0`, and returns `1`.

## Pitfalls

- Adding an unsigned type to the constraint, making the `-1` branch dead code.
- Returning `T` instead of `int`.
- Treating `-0.0` as negative — it compares equal to zero, so the result is `0`.
