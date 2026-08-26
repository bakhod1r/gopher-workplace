# Average

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The analytics dashboard also needs the arithmetic mean.

## Task

Implement `Average` on `Stats` in [average.go](average.go):

1. Return `Sum / Count`.
2. Return `0` for empty stats (avoid division by zero).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Stats{Values: []float64{2, 4, 6}}.Average()
Output: 4
```

**Example 2:**

```
Input:  Stats{}.Average()
Output: 0
```

**Example 3:**

```
Input:  Stats{Values: []float64{-2, 2}}.Average()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Read-only computation. |
| 2 | **Division by zero guard** | Check length before dividing. |

## Hint

Sum the values, divide by `len(s.Values)`. Guard against empty.

## Validate

```bash
make verify
```
