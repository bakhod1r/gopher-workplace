# 2-D grids

## The idea

A 2-D structure is a slice of slices, `[][]T`, indexed `grid[row][col]` in
row-major order. Rows can be independent (and even ragged):

```go
for r := range grid { for c := range grid[r] { use(grid[r][c]) } }
```

## Why it matters

Boards, matrices, and tables are modeled this way. Each row is a separate slice,
so constructing a grid means allocating each row (a shared row is a classic bug).

## Watch out

- Allocate each row with its own `make`; assigning one row to all aliases them.
- `grid[row][col]` — outer is row, inner is column.
- Ragged rows are allowed; don't assume a fixed width.
