# Two Definitions Of "The Median"

## Intuition

Lay the sorted samples along a line from 0 to 1. Interpolation asks where `p` falls on that line and blends its neighbours; nearest-rank rounds up to the next real sample.

## Approach

1. Clone, sort, clamp `p`.
2. For the interpolated form, compute `(n-1) * p/100`, split it, and blend.
3. For the rank form, ceil and clamp into `[1, n]`.

## Solution

```go
func Interpolated(samples []float64, p float64) float64 {
	if len(samples) == 0 {
		return 0
	}
	s := slices.Clone(samples)
	slices.Sort(s)
	p = min(max(p, 0), 100)
	pos := float64(len(s)-1) * p / 100
	lo := int(math.Floor(pos))
	if lo >= len(s)-1 {
		return s[len(s)-1]
	}
	frac := pos - float64(lo)
	return s[lo] + (s[lo+1]-s[lo])*frac
}

func NearestRank(samples []float64, p float64) float64 {
	if len(samples) == 0 {
		return 0
	}
	s := slices.Clone(samples)
	slices.Sort(s)
	p = min(max(p, 0), 100)
	rank := int(math.Ceil(p / 100 * float64(len(s))))
	rank = min(max(rank, 1), len(s))
	return s[rank-1]
}
```

## Walkthrough

For `[1 2 3 4]` at p50 the position is `3 * 0.5 = 1.5`: halfway between `s[1]=2` and `s[2]=3`, giving 2.5. Nearest-rank instead takes rank `ceil(2) = 2`, which is the sample 2.

## Pitfalls

- Using `n` instead of `n-1` for the position, which pushes p100 out of range.
- Blending with the wrong neighbour when the position lands exactly on an index.
- Comparing a p99 from an interpolating library against one from a nearest-rank library.
