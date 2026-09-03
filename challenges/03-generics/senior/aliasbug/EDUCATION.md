# Chunks That Share Storage

## Intuition

`s[i:end]` has length `end-i` but capacity all the way to the end of `s`. Appending therefore writes over the following elements in place instead of allocating.

## Approach

1. Allocate a group of the right length.
2. Copy the window into it.
3. Append the copy.

## Solution

```go
func Chunk[T any](s []T, size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}
	out := make([][]T, 0)
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

Appending to the first group of `Chunk([]int{1,2,3,4}, 2)` overwrites `s[2]`, which is the first element of the second group.

## Pitfalls

- Returning sub-slices from any API whose callers may append.
- Fixing it with `s[i:end:end]`, which caps capacity but still shares reads — acceptable, but the copy is clearer here.
- Assuming `copy` allocates: it does not, the destination must exist first.
