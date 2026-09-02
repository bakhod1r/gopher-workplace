# Percentile

## Intuition

The `p = 0` case is the one people miss: the ceiling of zero is zero, which would index `c[-1]`.

## Approach

1. Return zero and `false` for an empty slice.
2. Clamp `p`.
3. Sort a clone, compute the ceiling rank, floor it at 1, and index `rank-1`.

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

`Percentile([]float64{1,2,3,4}, 50)` computes rank `ceil(2.0) = 2` and returns `c[1]`, which is `2`.

## Pitfalls

- Indexing with the rank directly and running off the end.
- Forgetting the rank floor, producing a negative index at `p = 0`.
- Sorting the caller's slice.
