# Unique By Key

## Intuition

Constraining the element type would rule out structs holding slices — exactly the records this helper is for. Projecting a key keeps the constraint where it belongs.

## Approach

1. Track seen keys in a map.
2. Skip elements whose key was seen.
3. Append the first of each key.

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

`UniqueBy([]int{1,11,2}, mod10)` keeps `1`, skips `11` (same key), keeps `2`.

## Pitfalls

- Constraining `T comparable` and rejecting valid element types.
- Keeping the last element per key instead of the first.
- Sorting first and losing input order.
