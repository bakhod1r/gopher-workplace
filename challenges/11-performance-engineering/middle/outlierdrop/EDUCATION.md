# The Interquartile Rule

## Intuition

Take the range covering the middle half of the data. Anything more than `k` of those ranges away from that middle is unlike the rest of the data — and that verdict does not depend on the outlier itself.

## Approach

1. Compute Q1 and Q3 over a sorted copy.
2. Derive the low and high fences from the IQR.
3. Filter the original order into a fresh slice.

## Solution

```go
func Quartiles(samples []float64) (float64, float64) {
	if len(samples) == 0 {
		return 0, 0
	}
	s := slices.Clone(samples)
	slices.Sort(s)
	return s[rank(len(s), 0.25)], s[rank(len(s), 0.75)]
}

func rank(n int, q float64) int {
	r := int(math.Ceil(q * float64(n)))
	return min(max(r, 1), n) - 1
}

func Filter(samples []float64, k float64) []float64 {
	out := make([]float64, 0, len(samples))
	if len(samples) == 0 {
		return out
	}
	if k < 0 {
		k = 0
	}
	q1, q3 := Quartiles(samples)
	iqr := q3 - q1
	lo, hi := q1-k*iqr, q3+k*iqr
	for _, v := range samples {
		if v >= lo && v <= hi {
			out = append(out, v)
		}
	}
	return out
}
```

## Walkthrough

For `[1 2 3 4 1000]` the quartiles are 2 and 4, so the IQR is 2 and the fences are −1 and 7. The 1000 falls outside; every other sample survives, and the fences never noticed the outlier existed.

## Pitfalls

- Using the mean and standard deviation, which the outlier itself inflates.
- Sorting the output and losing the caller's ordering.
- Dropping outliers silently in a report — the discarded requests were real.
