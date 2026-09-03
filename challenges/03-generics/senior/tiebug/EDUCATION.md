# The Wrong Winner On A Tie

## Intuition

`>=` replaces the incumbent whenever the keys are equal, so the *last* maximal element ends up winning — the opposite of the documented rule.

## Approach

1. Keep the seed from `s[0]`.
2. Replace the incumbent only on a strictly greater key.

## Solution

```go
func MaxBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best, bestKey := s[0], key(s[0])
	for _, v := range s[1:] {
		k := key(v)
		if k > bestKey {
			best, bestKey = v, k
		}
	}
	return best, true
}
```

## Walkthrough

With two rows scoring 3, the strict comparison leaves the first in place; `>=` would swap it for the second.

## Pitfalls

- Assuming ties never happen in real data.
- Fixing it by sorting, which is slower and still needs a stable sort.
- Flipping the comparison and computing the minimum.
