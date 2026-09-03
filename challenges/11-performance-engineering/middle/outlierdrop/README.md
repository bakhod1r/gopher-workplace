# The Interquartile Rule

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Discarding outliers by "more than three standard deviations" fails on latency data, because the standard deviation is itself wrecked by the outlier you are trying to find. The interquartile range is computed from the middle of the distribution, so a single wild sample cannot move the threshold that would exclude it.

## Task

Implement both functions in [outlierdrop.go](outlierdrop.go):

1. `Quartiles` returns the nearest-rank Q1 and Q3 over a sorted copy.
2. `Filter` keeps only the samples inside `[q1 - k*IQR, q3 + k*IQR]`, preserving order.
3. A negative `k` is treated as `0`, no samples gives an empty non-nil slice, and neither function modifies the input.

## Examples

**Example 1:**

```
Input:  Filter([1 2 3 4 1000], 1.5)
Output: [1 2 3 4]
```

**Example 2:**

```
Input:  Filter([4 1 1000 3 2], 1.5)
Output: [4 1 3 2]
```

**Example 3:**

```
Input:  Quartiles([1 2 3 4])
Output: 1, 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Robust statistics** | Quartiles come from the middle, so outliers cannot inflate the threshold. |
| 2 | **`k = 1.5` is the convention** | It comes from the normal distribution, and nothing forces you to keep it. |
| 3 | **Dropping is a decision, not a cleanup** | The samples you discarded happened; say so when you report the number. |

## Topics used again

Sorting a copy, percentile ranks, filtering with order preserved.

## Hint

Compute the bounds once before the filtering loop.

## Validate

```bash
make verify
```
