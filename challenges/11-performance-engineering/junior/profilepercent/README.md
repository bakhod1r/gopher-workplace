# The Percentage Columns

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Absolute numbers in a profile are meaningless on their own — 400ms is huge in a request path and invisible in a batch job. Every pprof line therefore carries a share of the total, printed to two decimals.

## Task

Implement both functions in [profilepercent.go](profilepercent.go):

1. `Percent` returns `value` as a share of `total`, rounded to two decimal places.
2. `Percent` returns `0` when `total` is non-positive.
3. `Format` renders the same share as pprof prints it: two decimals and a `%`, e.g. `"33.33%"`.

## Examples

**Example 1:**

```
Input:  Percent(1, 3)
Output: 33.33
```

**Example 2:**

```
Input:  Format(3, 3)
Output: "100.00%"
```

**Example 3:**

```
Input:  Percent(10, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Share beats absolute** | Percentages survive a change in workload size; nanoseconds do not. |
| 2 | **Rounding to a fixed scale** | Multiply, `math.Round`, divide — the standard two-decimal idiom. |
| 3 | **Formatting is not rounding** | `%.2f` rounds for display; the returned number should already be rounded. |

## Topics used again

`math.Round`, `fmt.Sprintf`, float conversion.

## Hint

`math.Round(x*100) / 100` gives two decimals; do the percentage scaling before that.

## Validate

```bash
make verify
```
