# Group Consecutive

## Intuition

Runs are positional, not value-based: the last `1` in `[1 1 2 1]` starts a third run, which is exactly what run-length encoding needs.

## Approach

1. Return empty for empty input.
2. Seed a run with `s[0]`.
3. Extend or flush on each neighbour comparison.
4. Flush the last run.

## Solution

```go
func GroupRuns[T comparable](s []T) [][]T {
	out := make([][]T, 0)
	if len(s) == 0 {
		return out
	}
	run := []T{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			run = append(run, s[i])
		} else {
			out = append(out, run)
			run = []T{s[i]}
		}
	}
	out = append(out, run)
	return out
}
```

## Walkthrough

`GroupRuns([]int{1,1,2,1})` flushes `[1 1]` at the `2`, then `[2]` at the final `1`, then flushes `[1]`.

## Pitfalls

- Grouping all equal values together, ignoring adjacency.
- Dropping the final run by forgetting the flush.
- Emitting an empty run for empty input.
