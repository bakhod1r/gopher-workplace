# Group By

## Intuition

The two type parameters carry different jobs: elements are never compared, keys are hashed. Constraining each to exactly what it needs keeps the function usable with struct elements.

## Approach

1. Allocate the result map.
2. For each element compute its key and append the element to that bucket.
3. Return the map.

## Solution

```go
func GroupBy[T any, K comparable](s []T, key func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, v := range s {
		k := key(v)
		out[k] = append(out[k], v)
	}
	return out
}
```

## Walkthrough

`GroupBy([]int{1,2,3}, parity)` appends `1` to a fresh nil slice for `odd`, then `2` under `even`, then `3` back under `odd`.

## Pitfalls

- Constraining `T` to `comparable` for no reason, rejecting struct elements.
- Forgetting to assign the `append` result back into the map.
- Returning a nil map for empty input.
