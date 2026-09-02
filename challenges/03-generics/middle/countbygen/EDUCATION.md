# Count By Key

## Intuition

When callers only need sizes, tallying avoids retaining every element — a real memory difference on large inputs.

## Approach

1. Allocate the map.
2. Increment the tally for each element's key.
3. Return the map.

## Solution

```go
func CountBy[T any, K comparable](s []T, key func(T) K) map[K]int {
	out := make(map[K]int)
	for _, v := range s {
		out[key(v)]++
	}
	return out
}
```

## Walkthrough

`CountBy([]int{1,2,3}, parity)` increments `odd`, `even`, then `odd` again.

## Pitfalls

- Building a `map[K][]T` and taking lengths afterwards.
- Returning a nil map for an empty input.
- Counting distinct keys rather than elements.
