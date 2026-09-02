# Is Sorted By

## Intuition

Reusing a single comparison for both sorting and verification is what keeps the two definitions from drifting apart.

## Approach

1. Return `slices.IsSortedFunc` with `cmp.Compare` on the name.

## Solution

```go
func ByName(items []Item) bool {
	return slices.IsSortedFunc(items, func(a, b Item) int {
		return cmp.Compare(a.Name, b.Name)
	})
}
```

## Walkthrough

A slice with two equal names still counts as sorted, because the comparison returns `0`, not a positive number.

## Pitfalls

- Writing a comparison that returns a `bool`.
- Rejecting equal neighbours.
- Checking only the first and last elements.
