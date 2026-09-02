# Deduplicate With Stdlib

## Intuition

Three stdlib calls replace the set-plus-sort you would otherwise write, and the ordering requirement is what makes this composition valid.

## Approach

1. Clone the input and normalise nil to empty.
2. Sort the clone in place.
3. Return `slices.Compact` of it.

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

`Distinct([]int{3, 1, 3})` sorts to `[1 3 3]`, then collapses the adjacent pair to `[1 3]`.

## Pitfalls

- Calling `Compact` without sorting, which leaves non-adjacent duplicates.
- Sorting the caller's slice in place.
- Discarding `Compact`'s return value.
