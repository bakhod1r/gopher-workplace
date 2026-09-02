# Compact Only Sees Its Neighbours

## Intuition

`slices.CompactFunc` is a single linear pass that only ever compares neighbours. That is exactly a run-length collapse — it removes *consecutive* duplicates. Making it remove all duplicates is the caller's job: sort by the same relation first.

## Approach

1. Clone the input so the caller's slice is untouched.
2. Sort by the case-folded value, stably, so equal values become adjacent.
3. Collapse the runs with `CompactFunc` and `strings.EqualFold`.

## Solution

```go
func DistinctFold(s []string) []string {
	out := slices.Clone(s)
	slices.SortStableFunc(out, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	return slices.CompactFunc(out, strings.EqualFold)
}
```

## Walkthrough

`["a", "b", "A"]` has no adjacent pair that folds equal, so nothing is removed and `A` survives alongside `a`.

## Pitfalls

- Compacting with an equality that does not match the sort relation — the runs then are not runs.
- Compacting in place on the caller's slice, which scrambles their data even when the answer is right.
