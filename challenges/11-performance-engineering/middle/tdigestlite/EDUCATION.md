# Quantiles Without Keeping The Data

## Intuition

You cannot answer "what was the 99th value" without the values — but you can answer "which bucket does the 99th value fall in" from counts alone, and that is what a latency SLO actually needs.

## Approach

1. `New` allocates `len(bounds)+1` counters.
2. `Add` finds the bucket by binary search and increments it and the total.
3. `Quantile` walks the counters accumulating until it reaches the target rank.

## Solution

```go
func New(bounds []float64) *Sketch {
	return &Sketch{Bounds: slices.Clone(bounds), counts: make([]int64, len(bounds)+1)}
}

func (s *Sketch) Add(v float64) {
	s.counts[sort.SearchFloat64s(s.Bounds, v)]++
	s.total++
}

func (s *Sketch) Count() int64 { return s.total }

func (s *Sketch) Quantile(p float64) (float64, bool) {
	if s.total == 0 {
		return 0, false
	}
	p = min(max(p, 0), 100)
	target := p / 100 * float64(s.total)
	var running int64
	for i, c := range s.counts {
		running += c
		if float64(running) >= target {
			if i >= len(s.Bounds) {
				return 0, false
			}
			return s.Bounds[i], true
		}
	}
	return 0, false
}
```

## Walkthrough

With five values under 1 and five under 10, `Quantile(50)` has a target of 5, which the first bucket reaches exactly, so the answer is its bound, 1. `Quantile(51)` targets 5.1 and needs the second bucket, giving 10 — the estimate jumps by a whole bucket, which is the accuracy you paid for.

## Pitfalls

- Returning the last bound for the overflow bucket, silently capping the reported tail.
- Comparing `running > target` rather than `>=`, which skips a bucket that lands exactly on the rank.
- Choosing linear bounds for latency data, wasting every bucket on a range nothing occupies.
