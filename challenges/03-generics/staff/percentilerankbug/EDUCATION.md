# The Median That Sits One Slot Too High

## Intuition

The nearest-rank definition gives a *rank* — a count of elements at or below the answer. Indexing a slice with a rank reads one element too far, and the upper clamp quietly repairs only the p100 case, leaving every other percentile a notch high.

## Approach

1. Reject empty inputs and out-of-range percentiles.
2. Sort a clone so the caller's slice is untouched.
3. Compute the ceiling of `p*n/100`, subtract one to get an index, and clamp into range.

## Solution

```go
func Percentile[T cmp.Ordered](xs []T, p int) (T, bool) {
	var zero T
	if len(xs) == 0 || p < 0 || p > 100 {
		return zero, false
	}
	sorted := slices.Clone(xs)
	slices.Sort(sorted)
	idx := (p*len(sorted)+99)/100 - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(sorted) {
		idx = len(sorted) - 1
	}
	return sorted[idx], true
}
```

## Walkthrough

For `[1 2 3 4 5]` at p50 the rank is 3, so the index must be 2 and the answer 3. Without the subtraction the index is 3 and the answer is 4.

## Pitfalls

- Sorting `xs` in place, which reorders the caller's data.
- Using `p*n/100` with truncating division, which is a *different* off-by-one at ranks that do not divide evenly.
