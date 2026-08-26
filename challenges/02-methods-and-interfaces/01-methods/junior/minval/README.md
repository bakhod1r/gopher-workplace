# Min Value

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The analytics dashboard needs the minimum value in a dataset.

## Task

Implement `Min` on `Stats` in [minval.go](minval.go):

1. Return the smallest value in `Values`.
2. Return `math.Inf(1)` (`+Inf`) for empty stats.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stats{Values: []float64{3, 1, 2}}.Min()
Output: 1
```

**Example 2:**

```
Input:  Stats{}.Min()
Output: +Inf
```

**Example 3:**

```
Input:  Stats{Values: []float64{-5, -1, 0}}.Min()
Output: -5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only scan. |
| 2 | **math.Inf** | `math.Inf(1)` returns positive infinity — identity for min. |

## Hint

Start with `math.Inf(1)` as the initial minimum, then scan for smaller values.

## Validate

```bash
make verify
```
