# The Numbers On The Dashboard

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Nobody experiences the mean latency. Users experience their own request, and the tail is where the complaints come from — which is why every latency panel shows p50, p90 and p99 instead of an average. The nearest-rank definition is the simplest one that never invents a value the system did not actually produce.

## Task

Implement both functions in [p50p90p99.go](p50p90p99.go):

1. `Percentile` sorts a copy ascending and returns the element at rank `ceil(p/100 * n)`, counted from 1.
2. A `p` at or below 0 gives the smallest sample, at or above 100 the largest, and no samples gives `0`; the input must not be modified.
3. `Summary` returns p50, p90 and p99, sorting once rather than three times.

## Examples

**Example 1:**

```
Input:  Percentile([1 2 3 4], 50)
Output: 2
```

**Example 2:**

```
Input:  Percentile([4 1 3 2], 75)
Output: 3
```

**Example 3:**

```
Input:  Summary([1 2 3 4])
Output: 2, 4, 4
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nearest rank returns real data** | The answer is always an observed sample, never an interpolation. |
| 2 | **p99 of 100 samples is rank 99** | Not the maximum — one sample in a hundred is allowed to be worse. |
| 3 | **Sort once** | Three percentiles over one sorted copy, not three sorts. |

## Topics used again

`slices.Clone`, `slices.Sort`, `math.Ceil`, guards.

## Hint

Rank `ceil(p/100*n)` counts from 1, so the index is one less — and it needs clamping into `[0, n-1]`.

## Validate

```bash
make verify
```
