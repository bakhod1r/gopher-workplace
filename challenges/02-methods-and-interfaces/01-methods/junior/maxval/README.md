# Max Value

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The analytics dashboard needs the maximum value in a dataset.

## Task

Implement `Max` on `Stats` in [maxval.go](maxval.go):

1. Return the largest value in `Values`.
2. Return `math.Inf(-1)` (`-Inf`) for empty stats.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stats{Values: []float64{3, 1, 2}}.Max()
Output: 3
```

**Example 2:**

```
Input:  Stats{}.Max()
Output: -Inf
```

**Example 3:**

```
Input:  Stats{Values: []float64{-5, -1, 0}}.Max()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only scan. |
| 2 | **math.Inf(-1)** | Negative infinity — identity for max. |

## Hint

Mirror of Min: start with `-Inf`, scan for larger values.

## Validate

```bash
make verify
```
