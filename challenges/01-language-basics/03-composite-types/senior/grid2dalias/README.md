# Shared Grid Rows

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`New` allocates **one** row slice and assigns it to every row, so all rows share
the same backing array — writing `g[0][0]` changes every row.

## Task

Fix the loop between the markers in [grid2dalias.go](grid2dalias.go) to allocate
a fresh row per line.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  rows=2, cols=2
Output: [[0,0],[0,0]] with independent rows
```

**Example 2:**

```
Input:  rows=3, cols=1
Output: [[0],[0],[0]]
```

**Example 3:**

```
Input:  rows=1, cols=3
Output: [[0,0,0]]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
