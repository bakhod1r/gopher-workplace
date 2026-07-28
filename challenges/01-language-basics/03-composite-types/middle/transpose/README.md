# Matrix Transpose

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Turning rows into columns — pivoting a table.

## Task

Implement `Transpose(grid)` for a rectangular grid.

## Examples

```go
Transpose([][]int{{1,2,3},{4,5,6}}) // => [[1 4] [2 5] [3 6]]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Dimensions swap** | r×c becomes c×r. |
| 2 | **Allocate result** | `make` each output row. |
| 3 | **Index swap** | `out[j][i] = grid[i][j]`. |

## Hint

`rows, cols := len(grid), len(grid[0])`; allocate cols rows; `out[j][i]=grid[i][j]`.

## Validate

```bash
make verify
```
