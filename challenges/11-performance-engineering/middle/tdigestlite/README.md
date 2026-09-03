# Quantiles Without Keeping The Data

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

An exact p99 needs every sample, which at a million requests per minute is not a metric, it is a data pipeline. Sketches trade exactness for fixed memory: a handful of counters, an estimate that is off by at most one bucket, and constant-time observation. Every latency metric you have ever looked at was one of these.

## Task

Implement the four pieces in [tdigestlite.go](tdigestlite.go):

1. `New` builds a sketch with `len(bounds)+1` counters, and `Add` observes a value.
2. `Count` reports how many values were observed.
3. `Quantile` returns the upper bound of the first bucket whose cumulative count reaches `p` percent; landing in the overflow bucket, or having no data, gives `0, false`.

## Examples

**Example 1:**

```
Input:  five values at 0.5 and five at 5, bounds [1 10]; Quantile(50)
Output: 1, true
```

**Example 2:**

```
Input:  the same sketch; Quantile(51)
Output: 10, true
```

**Example 3:**

```
Input:  a sketch whose p99 lands past the last bound
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fixed memory, bounded error** | The estimate is accurate to the bucket width, forever, at any volume. |
| 2 | **Quantiles come from cumulative counts** | Walk the buckets accumulating until you cross the target rank. |
| 3 | **The overflow bucket has no upper bound** | Reporting the last bound there would understate the tail; say "unknown" instead. |

## Topics used again

Methods on pointer receivers, binary search, running totals, multiple returns.

## Hint

The target is `p/100 * total`; accumulate and stop at the first bucket that reaches it.

## Validate

```bash
make verify
```
