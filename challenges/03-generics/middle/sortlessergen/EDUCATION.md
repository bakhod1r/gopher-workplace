# Sort By Less

## Intuition

A boolean `Less` cannot distinguish "equal" from "greater" on its own, so the second call is what makes stability possible.

## Approach

1. Clone and normalise nil.
2. Sort stably with a comparator built from two `Less` calls.
3. Return the copy.

## Solution

```go
func SortedLess[T Lesser[T]](s []T) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	slices.SortStableFunc(out, func(a, b T) int {
		switch {
		case a.Less(b):
			return -1
		case b.Less(a):
			return 1
		default:
			return 0
		}
	})
	return out
}
```

## Walkthrough

For two equal versions both `Less` calls are false, so the comparator returns 0 and the stable sort leaves them alone.

## Pitfalls

- Returning `1` in the default branch, which breaks the sort's invariants.
- Calling `Less` once and guessing the other direction.
- Sorting the caller's slice in place.
