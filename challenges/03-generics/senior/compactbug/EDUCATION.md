# Deduplication That Misses

## Intuition

`Compact` walks neighbours; two equal values separated by a third are never compared, so they both survive. Sorting first is the precondition it silently assumes.

## Approach

1. Clone the input.
2. Sort the clone.
3. Compact it and return.

## Solution

```go
func Distinct[T cmp.Ordered](s []T) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	slices.Sort(out)
	return slices.Compact(out)
}
```

## Walkthrough

`Distinct([]int{3,1,3})` without the sort compares `3,1` and `1,3` — no neighbours match, so nothing is removed.

## Pitfalls

- Reading `Compact` as "remove duplicates".
- Sorting the caller's slice instead of the clone.
- Testing only with pre-sorted fixtures.
