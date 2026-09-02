# Binary Search That Overshoots The Duplicates

## Intuition

The invariant is that every index below `lo` holds an element strictly less than `v`. Moving `lo` past an element equal to `v` breaks it: the search then converges on the first element strictly *greater* than `v`, which is the upper bound.

## Approach

1. Keep a half-open interval `[lo, hi)`.
2. Compute the midpoint without overflow.
3. Discard the left half only when `s[mid]` is strictly less than `v`; otherwise shrink the right end to `mid`.

## Solution

```go
func LowerBound[T cmp.Ordered](s []T, v T) int {
	lo, hi := 0, len(s)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if s[mid] < v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
```

## Walkthrough

For `[1 2 2 2 3]` and `v = 2` the buggy loop walks past all three twos and returns 4, so a range query starting at 2 skips every sample stamped 2.

## Pitfalls

- Using `mid := (lo + hi) / 2`, which can overflow once the indices are large.
- Returning `-1` for a miss; `len(s)` is the insertion point the callers need.
