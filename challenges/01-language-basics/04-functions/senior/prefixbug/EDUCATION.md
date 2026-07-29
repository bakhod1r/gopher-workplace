# Prefix-sum range queries

## Intuition

With `pre[i]` the sum of the first i elements, any range `[l:r)` is `pre[r] - pre[l]`; dropping the `pre[l]` term over-counts.

## Approach

1. A prefix-sum range is `pre[r] - pre[l]`.
2. The bug returns only `pre[r]`, ignoring the left offset.

## Solution

```go
func RangeSum(xs []int, l, r int) int {
	pre := make([]int, len(xs)+1)
	for i, v := range xs {
		pre[i+1] = pre[i] + v
	}
	return pre[r] - pre[l]
}
```

## Walkthrough

`pre[r]` sums from index 0, not `l`. Subtracting `pre[l]` removes the prefix before `l`, giving the sub-range sum.

## Pitfalls

- `pre[r]` alone is the sum from 0, not from l.
- Match the half-open convention: `pre[r] - pre[l]` for `xs[l:r]`.
