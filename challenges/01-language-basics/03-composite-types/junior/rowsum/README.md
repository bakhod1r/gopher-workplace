# Row Sums

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A spreadsheet totals each row. A 2-D grid is a slice of slices, possibly ragged.

## Task

Implement `RowSums(grid)` returning one sum per row.

## Examples

```go
RowSums([][]int{{1,2,3},{4,5},{}}) // => [6 9 0]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice of slices** | Rows may differ in length. |
| 2 | **Nested range** | Outer rows, inner elements. |
| 3 | **One output per row** | Append a sum per row. |

## Hint

`for _, row := range grid { s := 0; for _, v := range row { s += v }; out = append(out, s) }`.

## Validate

```bash
make verify
```
