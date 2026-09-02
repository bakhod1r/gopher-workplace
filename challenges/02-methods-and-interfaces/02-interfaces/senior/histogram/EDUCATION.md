# Latency Histogram

## Intuition

A histogram trades exactness for a fixed memory footprint: you cannot recover individual samples, but percentile *shape* survives, and the cost stops depending on traffic.

## Approach

1. `Observe` increments `total`, finds the first bound `>= v`, and counts that bucket; otherwise counts overflow.
2. `Quantile` clamps `q`, computes `target = ceil(q * total)` with a floor of 1.
3. Walk the buckets accumulating counts and return the first bound whose cumulative count reaches the target.
4. Fall back to the last bound when the target lands in overflow.

## Solution

```go
func (h *Histogram) Observe(v int) {
	h.total++
	for i, b := range h.Bounds {
		if v <= b {
			h.counts[i]++
			return
		}
	}
	h.overflow++
}

func (h *Histogram) Quantile(q float64) int {
	if h.total == 0 || len(h.Bounds) == 0 {
		return 0
	}
	if q < 0 {
		q = 0
	}
	if q > 1 {
		q = 1
	}

	target := int(math.Ceil(q * float64(h.total)))
	if target < 1 {
		target = 1
	}

	seen := 0
	for i, c := range h.counts {
		seen += c
		if seen >= target {
			return h.Bounds[i]
		}
	}
	return h.Bounds[len(h.Bounds)-1]
}
```

## Walkthrough

With one sample per bucket, `Quantile(0.5)` targets rank `ceil(1.5) = 2`; the cumulative count reaches 2 at bucket 1, so the answer is its bound, 100.

## Pitfalls

- Storing samples to compute exact percentiles — the memory problem returns at scale.
- Using `v < b` instead of `v <= b`, which pushes boundary values into the wrong bucket.
- Forgetting the overflow bucket, so large values vanish from `total` and skew every quantile.
