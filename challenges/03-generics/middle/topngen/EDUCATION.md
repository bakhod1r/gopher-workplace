# Top N

## Intuition

Stability is what turns "first appearance" into the tie-break rule: the sort preserves the order of equal counts exactly as collected.

## Approach

1. Return empty for a non-positive `n`.
2. Tally counts while recording first-seen order.
3. Sort the distinct values stably by descending count, then take the first `n`.

## Solution

```go
func TopN[T comparable](s []T, n int) []T {
	out := make([]T, 0)
	if n <= 0 || len(s) == 0 {
		return out
	}
	counts := make(map[T]int, len(s))
	order := make([]T, 0, len(s))
	for _, v := range s {
		if counts[v] == 0 {
			order = append(order, v)
		}
		counts[v]++
	}
	slices.SortStableFunc(order, func(a, b T) int {
		return counts[b] - counts[a]
	})
	if n > len(order) {
		n = len(order)
	}
	out = append(out, order[:n]...)
	return out
}
```

## Walkthrough

`TopN([]int{1,2,2,3,3,3}, 2)` ranks `3` (three), `2` (two), `1` (one) and returns the first two.

## Pitfalls

- Building the candidate list by ranging the counts map.
- Using an unstable sort, which scrambles equal counts.
- Indexing `order[:n]` without clamping `n`.
