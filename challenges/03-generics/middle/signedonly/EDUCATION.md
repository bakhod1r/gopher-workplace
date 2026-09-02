# Signed Only

## Intuition

If `~uint` were in the set, `AbsDiff(2, 5)` would compute a wrap-around near the type's maximum and the `d < 0` guard would never fire.

## Approach

1. `Negate`: return `-v`.
2. `AbsDiff`: subtract, then flip the sign when the result is negative.

## Solution

```go
func Negate[T Signed](v T) T {
	return -v
}

func AbsDiff[T Signed](a, b T) T {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d
}
```

## Walkthrough

`AbsDiff(2, 5)` computes `-3`, sees it is negative, and returns `3`.

## Pitfalls

- Adding unsigned types to the constraint for "flexibility".
- Comparing before subtracting and duplicating the logic.
- Assuming `AbsDiff` is safe at the type's extremes — `AbsDiff(minInt, 0)` still overflows.
