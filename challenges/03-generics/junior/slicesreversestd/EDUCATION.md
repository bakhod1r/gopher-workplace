# Reverse In Place

## Intuition

Pairing an in-place helper with a clone is the standard way to get a pure function out of the `slices` package.

## Approach

1. Clone the input.
2. Replace a nil clone with an empty slice.
3. Reverse the clone in place and return it.

## Solution

```go
func ReverseCopy[T any](s []T) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	slices.Reverse(out)
	return out
}
```

## Walkthrough

`ReverseCopy([]int{})` clones to an empty slice; had the input been nil, the extra guard turns the nil clone into `[]int{}`.

## Pitfalls

- Calling `slices.Reverse(s)` directly, which reorders the caller's slice.
- Assigning the result of `slices.Reverse`, which returns nothing.
- Returning a nil slice when an empty one is expected.
