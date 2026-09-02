# Contains From Stdlib

## Intuition

The generic stdlib means membership, search, and sort no longer need a copy per element type. Reaching for it first is the habit worth building.

## Approach

1. Return `slices.Contains(tags, tag)`.

## Solution

```go
func HasTag[T comparable](tags []T, tag T) bool {
	return slices.Contains(tags, tag)
}
```

## Walkthrough

`HasTag([]int{1, 2}, 9)` forwards straight to the stdlib, which scans and returns `false`.

## Pitfalls

- Writing the loop by hand anyway.
- Using `slices.Index(...) != -1`, which works but says less.
- Declaring `[T any]` — `Contains` needs `comparable`.
