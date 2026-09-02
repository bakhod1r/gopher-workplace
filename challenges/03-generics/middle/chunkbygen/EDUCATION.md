# Chunk By Boundary

## Intuition

The pairwise predicate makes the split local: no key needs computing, and runs are found in one pass.

## Approach

1. Return empty for empty input.
2. Seed a group with the first element.
3. Append to the group or flush and restart at each boundary.
4. Flush the final group.

## Solution

```go
func ChunkBy[T any](s []T, together func(prev, cur T) bool) [][]T {
	out := make([][]T, 0)
	if len(s) == 0 {
		return out
	}
	group := []T{s[0]}
	for i := 1; i < len(s); i++ {
		if together(s[i-1], s[i]) {
			group = append(group, s[i])
		} else {
			out = append(out, group)
			group = []T{s[i]}
		}
	}
	out = append(out, group)
	return out
}
```

## Walkthrough

`ChunkBy([]int{1,1,2}, equal)` keeps `1,1` together, then splits before `2`, and the trailing flush emits `[2]`.

## Pitfalls

- Forgetting the final flush, dropping the last group.
- Starting the loop at index 0 and reading `s[-1]`.
- Emitting an empty group for empty input.
