# Sort By Key

## Intuition

Two decisions make this UI-safe: the clone keeps the caller's order intact, and stability keeps equal rows in their original sequence.

## Approach

1. Clone and normalise nil.
2. Sort stably comparing projected keys.
3. Return the copy.

## Solution

```go
func SortedBy[T any, K cmp.Ordered](s []T, key func(T) K) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	slices.SortStableFunc(out, func(a, b T) int {
		return cmp.Compare(key(a), key(b))
	})
	return out
}
```

## Walkthrough

Two people of the same age compare as `0`, so the stable sort leaves them in input order.

## Pitfalls

- Sorting the caller's slice in place.
- Using `slices.SortFunc`, which may reorder equal elements.
- Returning a `bool` from the comparison.
