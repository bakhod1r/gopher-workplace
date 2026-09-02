# Max By Key

## Intuition

Structs are not ordered, so the ordering has to come from a projection. Caching `bestKey` keeps the number of `key` calls linear.

## Approach

1. Return `zero, false` for an empty slice.
2. Seed `best` and `bestKey` from `s[0]`.
3. Replace both whenever a strictly larger key appears.

## Solution

```go
func MaxBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	bestKey := key(best)
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

`MaxBy` over two people scoring 2 each never takes the update branch, so the earlier element is returned.

## Pitfalls

- Constraining `T` to `cmp.Ordered`, which rejects structs.
- Recomputing `key(best)` inside the loop, doubling the calls.
- Using `>=`, which lets a later tie win.
