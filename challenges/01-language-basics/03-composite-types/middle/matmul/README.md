# Matrix Multiply

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The classic triple-loop matrix product.

## Task

Implement `Mul(a, b)`; return nil on dimension mismatch.

## Examples

```go
Mul([[1,2],[3,4]], [[5,6],[7,8]]) // => [[19 22] [43 50]]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Dimensions** | m×n · n×p → m×p. |
| 2 | **Dot product** | Sum over the shared dimension. |
| 3 | **Allocate result** | m rows of p. |

## Hint

`out[i][j] = sum_k a[i][k]*b[k][j]`.

## Validate

```bash
make verify
```
