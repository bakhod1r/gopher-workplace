# Deterministic Map Output

## Intuition

Reproducible output is the reason to sort here; the randomisation exists precisely so that this bug shows up early rather than in production.

## Approach

1. Collect every entry into a slice.
2. Sort by key with `cmp.Compare`.
3. Return the slice.

## Solution

```go
func Entries[K cmp.Ordered, V any](m map[K]V) []Entry[K, V] {
	out := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		out = append(out, Entry[K, V]{Key: k, Value: v})
	}
	slices.SortFunc(out, func(a, b Entry[K, V]) int {
		return cmp.Compare(a.Key, b.Key)
	})
	return out
}
```

## Walkthrough

Two calls on the same map return the same order, because the sort erases the randomised traversal.

## Pitfalls

- Rendering the map directly and getting a flaky test.
- Sorting by value when the keys are what identify an entry.
- Returning nil for an empty map.
