# The Accumulator That Rounds Off The Floats

## Intuition

The type set holds two integer types and one floating-point type. Accumulating through `int64` truncates every operand towards zero, so an input entirely below 1.0 sums to exactly nothing — and the integer members hide the defect completely.

## Approach

1. Return 0 for an empty input.
2. Accumulate into a `float64`, converting each element as it is added.
3. Divide by the element count.

## Solution

```go
func Mean[T int | int64 | float64](xs []T) float64 {
	if len(xs) == 0 {
		return 0
	}
	var sum float64
	for _, v := range xs {
		sum += float64(v)
	}
	return sum / float64(len(xs))
}
```

## Walkthrough

`Mean([]float64{0.5, 0.5, 0.5, 0.5})` converts each element to `int64(0.5) == 0` and answers `0` instead of `0.5`.

## Pitfalls

- Concluding that the fix is to drop `float64` from the union rather than widen the accumulator.
- Accumulating into `float32`, which starts losing digits well before three million elements.
