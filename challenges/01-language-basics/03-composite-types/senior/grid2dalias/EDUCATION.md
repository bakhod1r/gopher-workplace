# Independent rows in a 2-D grid

## The idea

A slice value is a header pointing at a backing array. Assigning the *same* slice
to every row makes them aliases. Allocate a new row each iteration:

```go
for i := range grid { grid[i] = make([]int, cols) }
```

## Why it matters

The "all rows are the same row" bug is a classic Go gotcha for 2-D structures.
Writes to one cell mysteriously appear everywhere, corrupting matrices, boards,
and buffers.

## Watch out

- One `make` outside the loop = shared row; one `make` per iteration = independent.
- The same trap hits maps of slices reused across keys.
- Ragged grids need per-row sizing anyway.
