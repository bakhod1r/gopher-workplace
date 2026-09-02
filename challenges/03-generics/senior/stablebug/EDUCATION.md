# Rows That Jump Between Renders

## Intuition

An unstable sort is free to permute equal elements, so any two rows with the same key can swap between runs — which is exactly what the UI notices.

## Approach

1. Clone and normalise nil.
2. Sort with the stable variant.
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

Two rows aged 20 compare as `0`; only the stable sort promises they keep their relative order.

## Pitfalls

- Reaching for `SortFunc` by default.
- Adding a tie-breaker field to fake stability, which changes the documented order.
- Concluding the sort is stable because a three-element test passed.
