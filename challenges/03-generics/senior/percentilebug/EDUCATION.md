# Percentile Zero Panics

## Intuition

Clamping `p` to `[0, 100]` leaves rank 0 reachable, and `c[-1]` panics — the failure is one subtraction away from the clamp that looked sufficient.

## Approach

1. Clamp `p`.
2. Compute the ceiling rank.
3. Floor the rank at 1 before indexing.

## Solution

```go
func Percentile[T Float](s []T, p float64) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	if p < 0 {
		p = 0
	}
	if p > 100 {
		p = 100
	}
	c := slices.Clone(s)
	slices.Sort(c)
	rank := int(math.Ceil(p / 100 * float64(len(c))))
	if rank < 1 {
		rank = 1
	}
	return c[rank-1], true
}
```

## Walkthrough

`Percentile(s, 0)` yields rank 0, so the index `-1` panics without the floor.

## Pitfalls

- Assuming clamping the percentage covers the rank.
- Clamping the rank to `len(c)` and forgetting the lower end.
- Switching to an interpolated definition, which changes every reported number.
