# Shared Grid Rows

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`New` allocates **one** row slice and assigns it to every row, so all rows share
the same backing array — writing `g[0][0]` changes every row.

## Task

Fix the loop between the markers in [grid2dalias.go](grid2dalias.go) to allocate
a fresh row per line.

## Examples

```go
g := New(3,2); g[0][0]=9 // g[1][0] must stay 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slices are references** | Assigning shares the backing array. |
| 2 | **Per-row make** | Allocate inside the loop. |
| 3 | **2-D construction** | Each inner slice is independent. |

## Hint

`for i := range grid { grid[i] = make([]int, cols) }`.

## Validate

```bash
make verify
```
