# Column Sum Indexing

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Summing column `c` means visiting `grid[r][c]` for every row `r`. The code writes
`grid[c][r]` — row and column swapped — which reads the wrong cells (and can panic
on a non-square grid).

## Task

Fix the index between the markers in [gridcolsum.go](gridcolsum.go).

## Examples

```go
ColSum([[1,2,3],[4,5,6],[7,8,9]], 0) // => 12
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Row-major indexing** | `grid[row][col]`. |
| 2 | **Column traversal** | Fix col, vary row. |
| 3 | **Shape safety** | Swapped indices panic on ragged/non-square grids. |

## Hint

`sum += grid[r][c]`.

## Validate

```bash
make verify
```
