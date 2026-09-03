# The Two Percentage Columns

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`pprof -top` prints each function's own share and a running total beside it. The running column is the one that ends arguments: when the first three rows already carry 80% of the profile, everything below them is not worth a sprint.

## Task

Implement both functions in [profiletotalpct.go](profiletotalpct.go):

1. `Top` orders rows by flat descending then name ascending, filling in each row's own share and the running total, both rounded to two decimals.
2. Drop non-positive values; an empty profile gives an empty, non-nil slice.
3. `CoveringCount` returns how many leading rows are needed to reach at least `pct` percent — `0` for a non-positive `pct`, all rows for more than 100.

## Examples

**Example 1:**

```
Input:  Top({a:3 b:1})
Output: [{a 3 75 75} {b 1 25 100}]
```

**Example 2:**

```
Input:  Top({a:1 b:1 c:1})
Output: three rows at 33.33, final CumPct 100
```

**Example 3:**

```
Input:  CoveringCount({a:50 b:30 c:20}, 80)
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cum% is a running sum** | It accumulates down the sorted listing, so it only means anything in order. |
| 2 | **Rounding is per row, not cumulative** | Accumulate the exact values and round for display, or the last row misses 100. |
| 3 | **The covering count sets scope** | It turns "the profile is long" into "there are three things to fix". |

## Topics used again

Sorting with ties, `math.Round`, running totals, guards.

## Hint

Keep an exact running sum and round only when you store it in the row.

## Validate

```bash
make verify
```
