# Round To Integer

## Intuition

The conversion sandwich is the standard way to reuse `math` from generic code; the constraint keeps it meaningful by excluding integer types.

## Approach

1. Return `T(math.Round(float64(v)))`.

## Solution

```go
func RoundHalfUp[T Float](v T) T {
	return T(math.Round(float64(v)))
}
```

## Walkthrough

`RoundHalfUp(-2.5)` rounds away from zero to `-3`, not towards zero to `-2`.

## Pitfalls

- Adding `0.5` and truncating, which is wrong for negatives.
- Passing `v` straight to `math.Round`, which does not compile.
- Expecting banker's rounding, which `math.RoundToEven` provides instead.
