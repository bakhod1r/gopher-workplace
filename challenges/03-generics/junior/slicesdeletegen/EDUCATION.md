# Delete

## Intuition

Half-open ranges are the same convention as slicing, which keeps `j - i` equal to the number of elements removed.

## Approach

1. Return `s` when `i` is out of range.
2. Otherwise clone and return `slices.Delete(clone, i, i+1)`.

## Solution

```go
func RemoveAt[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	return slices.Delete(slices.Clone(s), i, i+1)
}
```

## Walkthrough

`RemoveAt([]int{1, 2, 3}, 1)` deletes the range `[1, 2)` — just the `2` — leaving `[1 3]`.

## Pitfalls

- Calling `Delete(s, i, i)`, which removes nothing.
- Deleting from the caller's slice, which shifts its elements.
- Treating `i` as valid when it equals `len(s)`.
