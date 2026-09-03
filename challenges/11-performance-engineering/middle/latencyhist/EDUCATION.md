# Counting Into Buckets

## Intuition

Dividing by the bucket width and truncating turns a value into a bucket index. Everything past the last bucket lands in one shared overflow counter.

## Approach

1. Guard the parameters, allocate `n+1` counters.
2. For each non-negative sample, compute the index and clamp it to `n`.
3. `Busiest` scans for the maximum with a strict `>`.

## Solution

```go
func Histogram(samples []float64, width float64, n int) []int64 {
	if width <= 0 || n <= 0 {
		return []int64{}
	}
	out := make([]int64, n+1)
	for _, v := range samples {
		if v < 0 {
			continue
		}
		i := int(v / width)
		if i > n {
			i = n
		}
		out[i]++
	}
	return out
}

func Busiest(buckets []int64) int {
	best := -1
	var bestN int64
	for i, c := range buckets {
		if c > bestN {
			best, bestN = i, c
		}
	}
	return best
}
```

## Walkthrough

`int(v/width)` truncates, so 9.99 lands in bucket 0 and 10 in bucket 1 — the half-open convention falls out of the arithmetic rather than needing an explicit comparison.

## Pitfalls

- Allocating `n` counters and indexing out of range on the first big sample.
- Rounding instead of truncating, which shifts every sample half a bucket up.
- Reading precise percentiles off coarse buckets; the resolution is the bucket width.
