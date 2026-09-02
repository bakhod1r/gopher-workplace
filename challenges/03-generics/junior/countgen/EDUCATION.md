# Count Occurrences

## Intuition

A counter accumulates across the whole slice, so the loop runs to completion. Declaring `n` inside the loop would reset it on every iteration.

## Approach

1. Start `n := 0`.
2. Increment `n` for each element equal to `v`.
3. Return `n`.

## Solution

```go
func Count[T comparable](s []T, v T) int {
	n := 0
	for _, e := range s {
		if e == v {
			n++
		}
	}
	return n
}
```

## Walkthrough

`Count([]int{1, 2, 1}, 1)` increments on index 0, skips index 1, increments on index 2, and returns `2`.

## Pitfalls

- Returning inside the loop, which stops at the first match.
- Declaring the counter inside the loop body.
- Counting indexes rather than matching elements.
