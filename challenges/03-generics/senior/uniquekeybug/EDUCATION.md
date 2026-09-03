# First Occurrence Not Kept

## Intuition

Storing into `byKey` on every iteration overwrites the earlier element, so the final pass emits the last row for each key while the order list still says first-seen.

## Approach

1. Track seen keys in a set.
2. Append the first element for each key and skip the rest.

## Solution

```go
func UniqueBy[T any, K comparable](s []T, key func(T) K) []T {
	seen := make(map[K]bool, len(s))
	out := make([]T, 0, len(s))
	for _, v := range s {
		k := key(v)
		if seen[k] {
			continue
		}
		seen[k] = true
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

For rows `{1 a}` then `{1 b}`, the map ends holding `b`, which is emitted in `a`'s position.

## Pitfalls

- Assuming order preservation implies the right element was kept.
- Constraining the element type to `comparable` to use it as a key.
- Sorting to deduplicate, which destroys input order.
