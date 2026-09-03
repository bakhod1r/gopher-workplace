# The Ranking That Changes Every Run

## Intuition

Map iteration order is deliberately randomised in Go. Sorting the collected keys by value alone leaves ties exactly where iteration put them, so the output differs between runs even for identical input.

## Approach

1. Collect the keys.
2. Compare by value, descending.
3. Break ties by comparing the keys themselves.

## Solution

```go
func Rank[K cmp.Ordered, V cmp.Ordered](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b K) int {
		if c := cmp.Compare(m[b], m[a]); c != 0 {
			return c
		}
		return cmp.Compare(a, b)
	})
	return keys
}
```

## Walkthrough

For eight keys that all score 1, repeated calls return eight different permutations.

## Pitfalls

- Switching to `SortStableFunc` to "make it deterministic" — stability cannot rescue a randomised input order.
- Sorting the keys first and then sorting by value stably; that works, but one total-order comparator says the intent directly.
