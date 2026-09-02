# Insert

## Intuition

The upper bound is `len(s)`, not `len(s)-1`, because inserting after the last element is meaningful. Getting that boundary wrong is the classic off-by-one here.

## Approach

1. Return `s` when `i` is negative or greater than `len(s)`.
2. Otherwise clone and return `slices.Insert(clone, i, v)`.

## Solution

```go
func InsertAt[T any](s []T, i int, v T) []T {
	if i < 0 || i > len(s) {
		return s
	}
	return slices.Insert(slices.Clone(s), i, v)
}
```

## Walkthrough

`InsertAt([]int{1}, 1, 2)` inserts at the end, producing `[1 2]` without touching the caller's slice.

## Pitfalls

- Rejecting `i == len(s)`, which forbids a legal append.
- Passing the caller's slice to `Insert`, which may write into its backing array.
- Letting a bad index reach `slices.Insert` and panic.
