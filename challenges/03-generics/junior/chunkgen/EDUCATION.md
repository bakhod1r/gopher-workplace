# Chunk

## Intuition

Sub-slicing is cheap but aliases the input: a later `append` into a group could overwrite the next group's data. Copying each group keeps the results independent.

## Approach

1. Return an empty result for a non-positive `size`.
2. Step `i` from 0 to `len(s)` in increments of `size`.
3. Clamp `end` to `len(s)`, copy `s[i:end]` into a fresh slice, and append it.

## Solution

```go
func Chunk[T any](s []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	out := make([][]T, 0, (len(s)+size-1)/max(size, 1))
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		group := make([]T, end-i)
		copy(group, s[i:end])
		out = append(out, group)
	}
	return out
}
```

## Walkthrough

`Chunk([]int{1, 2, 3}, 2)` produces `[1 2]` for `i = 0`, then clamps `end` from 4 to 3 for `i = 2`, producing `[3]`.

## Pitfalls

- Forgetting the `size <= 0` guard, which makes the loop never advance and hang.
- Appending `s[i:end]` directly, leaving every group aliased to the input.
- Computing `end` without clamping, which panics on the final group.
