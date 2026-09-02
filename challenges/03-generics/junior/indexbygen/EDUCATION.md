# Index By

## Intuition

This is the sibling of `GroupBy`: same traversal, but each key holds one element rather than a bucket. Which behaviour you want depends on whether keys are unique.

## Approach

1. Allocate the map with capacity `len(s)`.
2. Assign `out[key(v)] = v` per element.
3. Return the map.

## Solution

```go
func IndexBy[T any, K comparable](s []T, key func(T) K) map[K]T {
	out := make(map[K]T, len(s))
	for _, v := range s {
		out[key(v)] = v
	}
	return out
}
```

## Walkthrough

`IndexBy([]int{1, 11}, mod10)` writes `1` under key 1, then overwrites it with `11`.

## Pitfalls

- Skipping elements whose key already exists, which makes the first duplicate win instead.
- Using `GroupBy` semantics and returning `map[K][]T`.
- Returning nil for empty input.
