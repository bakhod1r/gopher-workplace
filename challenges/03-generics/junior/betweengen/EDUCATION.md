# Between

## Intuition

Go has no chained comparison, so the range test is two comparisons joined with `&&`. Both operators come from `cmp.Ordered`.

## Approach

1. Return `v >= lo && v <= hi`.

## Solution

```go
func Between[T cmp.Ordered](v, lo, hi T) bool {
	return v >= lo && v <= hi
}
```

## Walkthrough

`Between(1, 1, 3)` evaluates `1 >= 1` (true) and `1 <= 3` (true), so the endpoint counts as inside.

## Pitfalls

- Writing `lo <= v <= hi`, which does not compile in Go.
- Using strict inequalities and rejecting the endpoints.
- Swapping `lo` and `hi` in the call, which silently returns `false`.
