# Explicit Bucket Bounds

## Intuition

The bounds partition the number line into `n+1` intervals. Finding the bucket is a lower-bound search: the first bound the value does not exceed.

## Approach

1. `Index` uses a binary search for the first bound `>= v`.
2. `Counts` allocates `len(bounds)+1` counters and tallies.
3. `Cumulative` walks once with a running sum into a fresh slice.

## Solution

```go
func Index(bounds []float64, v float64) int {
	return sort.SearchFloat64s(bounds, v)
}

func Counts(bounds []float64, samples []float64) []int64 {
	out := make([]int64, len(bounds)+1)
	for _, v := range samples {
		out[Index(bounds, v)]++
	}
	return out
}

func Cumulative(counts []int64) []int64 {
	out := make([]int64, 0, len(counts))
	var running int64
	for _, c := range counts {
		running += c
		out = append(out, running)
	}
	return out
}
```

## Walkthrough

`sort.SearchFloat64s` returns the first index where `bounds[i] >= v`, so a value exactly equal to a bound lands in that bound's own bucket — the inclusive-upper convention metrics systems use.

## Pitfalls

- Searching for `>` instead of `>=`, pushing boundary values one bucket too high.
- Allocating `len(bounds)` counters and panicking on the first overflow sample.
- Accumulating in place over the caller's counts, corrupting the raw histogram.
