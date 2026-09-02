# Count Positive

## Intuition

Untyped constants adapt to the type parameter at instantiation, so `v > 0` compiles once and means the right thing for `int` and `float64` both.

## Approach

1. Start `n := 0`.
2. Increment for each element greater than zero.
3. Return `n`.

## Solution

```go
func CountPositive[T Number](s []T) int {
	n := 0
	for _, v := range s {
		if v > 0 {
			n++
		}
	}
	return n
}
```

## Walkthrough

`CountPositive([]float64{0.5, -0.5})` counts only `0.5`, returning `1`.

## Pitfalls

- Returning `T` instead of `int`.
- Using `>=`, which counts zeros as positive.
- Writing `T(0)` where the untyped `0` already works.
