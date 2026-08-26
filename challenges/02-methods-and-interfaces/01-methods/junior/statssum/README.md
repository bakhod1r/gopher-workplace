# Stats Sum

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An analytics dashboard computes aggregate statistics. `Sum` is the most basic.

## Task

Implement `Sum` on `Stats` in [statssum.go](statssum.go):

1. Return the sum of all `Values`.
2. Empty or nil → return `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stats{Values: []float64{1, 2, 3}}.Sum()
Output: 6
```

**Example 2:**

```
Input:  Stats{}.Sum()
Output: 0
```

**Example 3:**

```
Input:  Stats{Values: []float64{-1, -2, 3}}.Sum()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only aggregation. |
| 2 | **Range loop** | Iterate over the slice. |

## Hint

Loop over `s.Values`, accumulate into a `total`, return it.

## Validate

```bash
make verify
```
