# Matrix Multiply

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

The classic triple-loop matrix product.

## Task

Implement `Mul(a, b)`; return nil on dimension mismatch.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [[1,2],[3,4]] * [[5,6],[7,8]]
Output: [[19,22],[43,50]]
```

**Example 2:**

```
Input:  [[1,2,3]] * [[1]]
Output: nil
```

_Explanation:_ inner dims 3 vs 1 mismatch

**Example 3:**

```
Input:  [[2]] * [[3]]
Output: [[6]]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
