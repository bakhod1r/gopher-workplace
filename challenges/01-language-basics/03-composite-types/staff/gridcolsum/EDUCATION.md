# Row-major indexing

## Intuition

Go's 2-D slices are row-major: `grid[row][col]`. To walk a column, hold the column
fixed and vary the row:

```go
for r := range grid { sum += grid[r][c] }
```

Swapping to `grid[c][r]` reads across the wrong axis.

## Approach

1. Bug: `grid[c][r]` swaps row/column roles, indexing the wrong cells (and can panic on non-square grids).
2. Fix: `grid[r][c]` walks each row r and sums the fixed column c.

## Solution

```go
func ColSum(grid [][]int, c int) int {
	sum := 0
	for r := range grid {
		sum += grid[r][c]
	}
	return sum
}
```

## Walkthrough

c=1, rows 0..2: grid[0][1]=2, grid[1][1]=4, grid[2][1]=6 -> 12. Buggy grid[1][r] would read row 1 across columns instead.

## Pitfalls

- Outer index is the row; inner is the column.
- Column access is cache-unfriendly (strided) but semantically fine.
- Validate rectangularity if the grid may be ragged.
