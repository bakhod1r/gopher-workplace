# Labeled break vs continue

## Intuition

A label lets `break`/`continue` act on an enclosing loop; `break Label` terminates it while `continue Label` moves to its next iteration — a one-word difference with opposite effects.

## Approach

1. On a bad cell we should skip to the next row, not stop all rows.
2. The bug `break Rows` exits the outer loop; use `continue Rows`.

## Solution

```go
func CleanRows(grid [][]int) int {
	count := 0
Rows:
	for _, row := range grid {
		for _, v := range row {
			if v < 0 {
				continue Rows
			}
		}
		count++
	}
	return count
}
```

## Walkthrough

`break Rows` abandons all remaining rows on the first negative. `continue Rows` skips only the offending row, so the count reaches 2.

## Pitfalls

- `break Rows` ends the outer loop; `continue Rows` skips the current outer iteration.
- Pick the verb that matches 'skip this row' vs 'stop scanning'.
