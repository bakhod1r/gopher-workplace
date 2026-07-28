# Row-major indexing

## The idea

Go's 2-D slices are row-major: `grid[row][col]`. To walk a column, hold the column
fixed and vary the row:

```go
for r := range grid { sum += grid[r][c] }
```

Swapping to `grid[c][r]` reads across the wrong axis.

## Why it matters

Index-order confusion is a top source of matrix bugs. On a non-square or ragged
grid the swap doesn't just compute the wrong answer — it can index out of range
and panic.

## Watch out

- Outer index is the row; inner is the column.
- Column access is cache-unfriendly (strided) but semantically fine.
- Validate rectangularity if the grid may be ragged.
