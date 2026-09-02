# Binary Search By Key

## Intuition

Splitting the search from the equality check is what yields the insertion point for free, and choosing `hi = mid` is what makes it the *first* match rather than any match.

## Approach

1. Narrow `[lo, hi)` while `lo < hi`, moving `lo` past keys below the target.
2. Report `true` when `s[lo]` matches; otherwise return `lo` as the insertion point.

## Solution

```go
func SearchBy[T any, K cmp.Ordered](s []T, key func(T) K, target K) (int, bool) {
	lo, hi := 0, len(s)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if key(s[mid]) < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(s) && key(s[lo]) == target {
		return lo, true
	}
	return lo, false
}
```

## Walkthrough

Searching for a missing key leaves `lo` at the first larger element — exactly where it would be inserted.

## Pitfalls

- Using `hi = mid - 1`, which can skip the first match.
- Computing the midpoint as `(lo+hi)/2` and overflowing on huge slices.
- Returning `-1` when the puzzle asks for the insertion point.
