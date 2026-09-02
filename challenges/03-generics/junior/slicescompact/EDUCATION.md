# Compact

## Intuition

`Compact` is a linear neighbour-collapse, which is why sorting first turns it into a full deduplicator and why calling it on unsorted data usually surprises people.

## Approach

1. Clone the input, replacing a nil clone with an empty slice.
2. Return `slices.Compact` of the clone.

## Solution

```go
func Squash[T comparable](s []T) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	return slices.Compact(out)
}
```

## Walkthrough

`Squash([]int{1, 1, 2, 2, 1})` collapses the two runs but keeps the final `1`, since it is not adjacent to the first.

## Pitfalls

- Passing the caller's slice straight in, which mutates it.
- Expecting `[1 2 1]` to become `[1 2]`.
- Ignoring the return value — the original slice header still has the old length.
