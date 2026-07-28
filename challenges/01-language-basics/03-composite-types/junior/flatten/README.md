# Flatten a Grid

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Turning a 2-D grid into a flat list in row-major order.

## Task

Implement `Flatten(grid)`.

## Examples

```go
Flatten([][]int{{1,2},{3},{4,5}}) // => [1 2 3 4 5]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Row-major** | Iterate rows, then columns. |
| 2 | **append spread** | `append(out, row...)`. |
| 3 | **Ragged safe** | Handles rows of any length. |

## Hint

`for _, row := range grid { out = append(out, row...) }`.

## Validate

```bash
make verify
```
