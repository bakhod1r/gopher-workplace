# Averaging Without The Wild Ends

## Intuition

Sort the samples and ignore a fixed fraction at each end. What remains is the bulk of the distribution, and its mean barely moves when a single measurement goes wrong.

## Approach

1. Clone and sort.
2. Clamp `pct`, compute `k`, and slice the middle window.
3. Average it.

## Solution

```go
func TrimmedMean(samples []float64, pct float64) float64 {
	if len(samples) == 0 {
		return 0
	}
	s := slices.Clone(samples)
	slices.Sort(s)
	pct = min(max(pct, 0), 49.999999)
	k := int(math.Floor(pct / 100 * float64(len(s))))
	if 2*k >= len(s) {
		k = (len(s) - 1) / 2
	}
	return Mean(s[k : len(s)-k])
}

func Mean(samples []float64) float64 {
	if len(samples) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range samples {
		sum += v
	}
	return sum / float64(len(samples))
}
```

## Walkthrough

With `[1 1 1 1 1 1 1 1 1 10000]` a 10% trim removes one sample from each end, so the 10000 disappears and the mean drops from 1000.9 to 1 — the number the other nine runs agree on.

## Pitfalls

- Trimming one end only, which systematically biases the result.
- Letting `k` reach `len(s)/2` and slicing an empty window.
- Trimming production latency data, which deletes exactly the requests users complain about.
