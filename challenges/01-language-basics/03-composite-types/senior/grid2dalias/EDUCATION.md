# Independent rows in a 2-D grid

## Intuition

A slice value is a header pointing at a backing array. Assigning the *same* slice
to every row makes them aliases. Allocate a new row each iteration:

```go
for i := range grid { grid[i] = make([]int, cols) }
```

## Approach

1. Bug: one row slice is made once and assigned to every grid[i], so all rows alias the same backing array. 2. Writing grid[0][0] then shows up in every row. 3. Fix: allocate a fresh make([]int, cols) inside the loop per row.

## Solution

```go
func New(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return grid
}
```

## Walkthrough

With the bug, grid[0] and grid[1] point at the same array; grid[0][0]=9 makes grid[1][0]==9 too. Per-row make gives distinct arrays so writes are isolated.

## Pitfalls

- One `make` outside the loop = shared row; one `make` per iteration = independent.
- The same trap hits maps of slices reused across keys.
- Ragged grids need per-row sizing anyway.
