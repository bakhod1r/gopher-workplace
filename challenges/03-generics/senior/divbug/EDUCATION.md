# Truncation In A Closed Form

## Intuition

Dividing first throws away the fractional half before it can be multiplied back, so every range with an odd sum comes out short.

## Approach

1. Guard the empty range.
2. Compute the count.
3. Multiply the endpoint sum by the count, then divide by two.

## Solution

```go
func SumRange[T Integer](lo, hi T) T {
	if hi < lo {
		var zero T
		return zero
	}
	n := hi - lo + 1
	return (lo + hi) * n / 2
}
```

## Walkthrough

`SumRange(1, 2)` should be 3; dividing first gives `(3/2)*2 = 2`.

## Pitfalls

- Reordering integer arithmetic as if it were exact.
- Switching to floats to dodge the issue and introducing rounding error.
- Testing only even-length ranges.
