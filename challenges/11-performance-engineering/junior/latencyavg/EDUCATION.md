# Averaging Averages Is A Lie

## Intuition

A mean is a sum over a count. When the values are themselves means, the sum must be re-expanded by their counts before dividing.

## Approach

1. `Mean` sums and divides by the length, guarding empty input.
2. `WeightedMean` accumulates `value*weight` and the total weight over the paired prefix.

## Solution

```go
func Mean(samples []float64) float64 {
	if len(samples) == 0 {
		return 0
	}
	sum := 0.0
	for _, s := range samples {
		sum += s
	}
	return sum / float64(len(samples))
}

func WeightedMean(values []float64, weights []int) float64 {
	var sum float64
	var total int
	for i := 0; i < min(len(values), len(weights)); i++ {
		if weights[i] <= 0 {
			continue
		}
		sum += values[i] * float64(weights[i])
		total += weights[i]
	}
	if total == 0 {
		return 0
	}
	return sum / float64(total)
}
```

## Walkthrough

With weights 1 and 999, the fast endpoint dominates and the answer lands at 1.099ms — three orders of magnitude away from the unweighted 50.5ms.

## Pitfalls

- Ranging `values` and indexing `weights`, which panics when the weights are shorter.
- Dividing by `len(values)` instead of the total weight.
- Reporting a mean latency at all when the tail is what users notice.
