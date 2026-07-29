# Column Sum Indexing

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Summing column `c` means visiting `grid[r][c]` for every row `r`. The code writes
`grid[c][r]` — row and column swapped — which reads the wrong cells (and can panic
on a non-square grid).

## Task

Fix the index between the markers in [gridcolsum.go](gridcolsum.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  grid=[[1 2][3 4][5 6]], c=1
Output: 12
```

**Example 2:**

```
Input:  grid=[[1 2][3 4]], c=0
Output: 4
```

**Example 3:**

```
Input:  grid=[[7]], c=0
Output: 7
```

_Explanation:_ single cell.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
