# Most Frequent

## Intuition

Ranging the counts map would return a different winner between runs on ties. Walking the input in order makes "earliest" well-defined.

## Approach

1. Return `zero, false` for an empty slice.
2. Tally every element.
3. Walk `s` in order, keeping the first element with a strictly higher count.

## Solution

```go
func Mode[T comparable](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	counts := make(map[T]int, len(s))
	for _, v := range s {
		counts[v]++
	}
	best, bestN := s[0], counts[s[0]]
	for _, v := range s[1:] {
		if counts[v] > bestN {
			best, bestN = v, counts[v]
		}
	}
	return best, true
}
```

## Walkthrough

`Mode([]int{1,1,2,2})` sees both counts as 2, and the strict `>` keeps the earlier `1`.

## Pitfalls

- Picking the winner by ranging the map, which is non-deterministic.
- Using `>=`, which lets a later tie win.
- Returning the count rather than the element.
