# The Ranking That Changes Between Runs

## Intuition

Collecting keys by ranging over a map yields a different order on every run. A stable sort faithfully preserves that random order among equal counts, so the tie order is random too. Only a comparator that also orders ties makes the output a function of the input.

## Approach

1. Count the values.
2. Collect the distinct keys.
3. Sort by descending count, falling back to ascending key when counts tie.

## Solution

```go
func RankByCount[T cmp.Ordered](s []T) []T {
	cnt := make(map[T]int, len(s))
	for _, v := range s {
		cnt[v]++
	}
	keys := make([]T, 0, len(cnt))
	for k := range cnt {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b T) int {
		if c := cmp.Compare(cnt[b], cnt[a]); c != 0 {
			return c
		}
		return cmp.Compare(a, b)
	})
	return keys
}
```

## Walkthrough

With twelve distinct values each occurring once, every comparison returns 0 and the output is exactly the map's randomised iteration order — a different ranking on every run.

## Pitfalls

- Swapping in `SortStableFunc` as the fix; stability cannot repair a non-total comparator.
- Sorting the keys first and then stably sorting by count — correct, but one comparator says it plainly.
