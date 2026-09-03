# The Numbers On The Dashboard

## Intuition

Sort the samples and count in: the p-th percentile is the first value at least `p` percent of the data is below.

## Approach

1. Clone and sort.
2. Convert the percentile into a 1-based rank with `math.Ceil`, clamp it, and index.
3. `Summary` sorts once and indexes three times.

## Solution

```go
func Percentile(samples []float64, p float64) float64 {
	if len(samples) == 0 {
		return 0
	}
	s := slices.Clone(samples)
	slices.Sort(s)
	return percentileSorted(s, p)
}

func percentileSorted(s []float64, p float64) float64 {
	rank := int(math.Ceil(p / 100 * float64(len(s))))
	if rank < 1 {
		rank = 1
	}
	if rank > len(s) {
		rank = len(s)
	}
	return s[rank-1]
}

func Summary(samples []float64) (p50, p90, p99 float64) {
	if len(samples) == 0 {
		return 0, 0, 0
	}
	s := slices.Clone(samples)
	slices.Sort(s)
	return percentileSorted(s, 50), percentileSorted(s, 90), percentileSorted(s, 99)
}
```

## Walkthrough

With 99 requests at 1ms and one at 1000ms, p99 is rank 99 — still 1ms. Only p100 sees the outlier, which is exactly why a p99 panel can look calm while one user in a hundred suffers.

## Pitfalls

- Sorting the caller's slice in place.
- Using `p/100 * n` as a 0-based index, which shifts every percentile down by one rank.
- Reporting p99 as "the worst case"; it is the threshold 1% of requests exceed.
