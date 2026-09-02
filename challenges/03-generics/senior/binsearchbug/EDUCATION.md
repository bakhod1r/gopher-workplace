# Search That Skips The First Match

## Intuition

`<=` moves `lo` past every element equal to the target, so the loop ends after the run and the equality check compares the wrong element.

## Approach

1. Advance `lo` only past keys strictly below the target.
2. Report a hit when the landing position matches.

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

Searching `[1 1 2]` for `1` with `<=` lands on index 2, whose key is `2`, so the function reports "not found".

## Pitfalls

- Using `<=` in a lower-bound search.
- Setting `hi = mid - 1`, which can skip the first match entirely.
- Testing only with distinct keys.
