# Delete By Predicate

## Intuition

The in-place compaction is what makes it O(n) with no allocation; the clone is your choice, made explicit rather than hidden.

## Approach

1. Clone the input, normalising nil to empty.
2. Return `slices.DeleteFunc` of the clone.

## Solution

```go
func Purge[T any](s []T, drop func(T) bool) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	return slices.DeleteFunc(out, drop)
}
```

## Walkthrough

`Purge([]int{1,2,3}, isEven)` shifts `3` down over the deleted `2`, yielding `[1 3]`.

## Pitfalls

- Passing the caller's slice straight in, which rewrites it.
- Ignoring the return value — the original header keeps the old length.
- Building a new slice by hand when the stdlib does it in place.
