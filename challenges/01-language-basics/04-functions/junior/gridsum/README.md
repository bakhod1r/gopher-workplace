# 2D Grid Sum

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Nested `for range` loops walk a grid row by row; rows may have different lengths.

## Task

Implement `GridSum` in [gridsum.go](gridsum.go) summing every cell.

Do **not** change the function signature or the tests.

## Examples

```go
GridSum(nil)                  // => 0
GridSum([][]int{{1,2},{3,4}})  // => 10
GridSum([][]int{{5}})          // => 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nested loops** | Outer over rows, inner over cells. |
| 2 | **Ragged rows** | Row lengths can differ; range handles it. |
| 3 | **Accumulator** | One running total across all cells. |

## Hint

Range rows, then range each row's cells adding to `total`.

## Validate

```bash
make verify
```
