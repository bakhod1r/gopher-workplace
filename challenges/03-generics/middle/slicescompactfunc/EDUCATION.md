# Compact By Equality

## Intuition

Supplying equality is what lets you collapse on one field while ignoring others — impossible with the `==`-based variant.

## Approach

1. Clone and normalise nil.
2. Return `slices.CompactFunc` of the clone.

## Solution

```go
func Dedupe[T any](s []T, eq func(a, b T) bool) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	return slices.CompactFunc(out, eq)
}
```

## Walkthrough

`Dedupe([]int{1,2,1}, equal)` collapses nothing: the two `1`s are not adjacent.

## Pitfalls

- Expecting non-adjacent duplicates to be removed.
- Passing the caller's slice, which gets rewritten.
- Writing an equality that is not symmetric, which makes the result order-dependent.
