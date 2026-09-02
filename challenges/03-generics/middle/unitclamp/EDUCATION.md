# Named Unit Types

## Intuition

The named types exist to prevent unit confusion. A helper taking plain `float64` would quietly undo that, while `~float64` keeps the guarantee intact.

## Approach

1. Return `lo` below the range, `hi` above it, and `v` otherwise.

## Solution

```go
func ClampUnit[T ~float64](v, lo, hi T) T {
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

`ClampUnit(Meters(5), 0, 3)` returns `Meters(3)`; passing a `Seconds` bound alongside a `Meters` value would not compile.

## Pitfalls

- Writing `float64` instead of `~float64`, which rejects the named types.
- Converting to `float64` inside and returning that.
- Declaring separate parameters for the bounds, which lets units be mixed.
