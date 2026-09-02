# Min By Key

## Intuition

Recomputing `key(best)` on every iteration doubles the calls — invisible for a field read, expensive when the projection does real work.

## Approach

1. Return `zero, false` for an empty slice.
2. Seed `best` and `bestKey` from `s[0]`.
3. Replace both on a strictly smaller key.

## Solution

```go
func MinBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best, bestKey := s[0], key(s[0])
	for _, v := range s[1:] {
		k := key(v)
		if k < bestKey {
			best, bestKey = v, k
		}
	}
	return best, true
}
```

## Walkthrough

`MinBy([]int{3,1,1}, self)` drops to the first `1` and then refuses the tie, keeping the earlier element.

## Pitfalls

- Recomputing the key for the incumbent each iteration.
- Using `<=`, which lets later ties win.
- Seeding from a zero-valued key, which breaks for all-positive keys.
