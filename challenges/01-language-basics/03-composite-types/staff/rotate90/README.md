# Rotate Matrix 90°

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A 90° clockwise rotation is **transpose, then reverse each row**. The transpose is
done; the row reversal step is missing, so the result is only transposed.

## Task

Add the row-reversal between the markers in [rotate90.go](rotate90.go).

## Examples

```go
Rotate([[1,2,3],[4,5,6],[7,8,9]]) // => [[7 4 1] [8 5 2] [9 6 3]]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Transpose** | `out[i][j] = m[j][i]`. |
| 2 | **Reverse rows** | Two-pointer swap per row. |
| 3 | **Compose** | Rotation = transpose ∘ reverse-rows. |

## Hint

For each row of `out`, swap ends inward:
`for i := range out { for a, b := 0, n-1; a < b; a, b = a+1, b-1 { out[i][a], out[i][b] = out[i][b], out[i][a] } }`.

## Validate

```bash
make verify
```
