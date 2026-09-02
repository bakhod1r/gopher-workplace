# Variance

## Intuition

Restricting to floats is what makes the result meaningful: with `~int` in the set, every deviation below 1 would truncate to zero and the variance would collapse.

## Approach

1. Return zero for fewer than two samples.
2. Compute the mean.
3. Average the squared deviations.

## Solution

```go
func Variance[T Float](s []T) T {
	if len(s) < 2 {
		var zero T
		return zero
	}
	var sum T
	for _, v := range s {
		sum += v
	}
	mean := sum / T(len(s))
	var acc T
	for _, v := range s {
		d := v - mean
		acc += d * d
	}
	return acc / T(len(s))
}
```

## Walkthrough

`Variance([]float64{2, 4})` has mean 3, deviations ±1, and squared deviations averaging to 1.

## Pitfalls

- Dividing by `len(s)-1` when the doc says population variance.
- Allowing integer types, which truncate every deviation.
- Computing the mean inside the second loop, making it quadratic.
