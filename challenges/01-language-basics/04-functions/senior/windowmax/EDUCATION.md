# Sliding-window index arithmetic

## Intuition

When the window covers `[i-k+1 .. i]` after adding `xs[i]`, the element that just left is `xs[i-k]`; off-by-one here is a classic bug.

## Approach

1. A sliding window of width `k` drops the element `k` positions back.
2. The bug subtracts `xs[i-k+1]` (off by one); use `xs[i-k]`.

## Solution

```go
func MaxWindow(xs []int, k int) int {
	best := 0
	for i := 0; i < k; i++ {
		best += xs[i]
	}
	cur := best
	for i := k; i < len(xs); i++ {
		cur += xs[i] - xs[i-k]
		if cur > best {
			best = cur
		}
	}
	return best
}
```

## Walkthrough

When the window advances to index `i`, the element leaving is at `i-k`. Subtracting `i-k+1` keeps a stale element in the sum and corrupts the max.

## Pitfalls

- Adding `xs[i]` means removing `xs[i-k]` to keep width k.
- Draw the indices for one step to check the boundary.
