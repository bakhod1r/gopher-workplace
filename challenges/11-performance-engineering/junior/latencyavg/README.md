# Averaging Averages Is A Lie

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Two endpoints, one averaging 100ms over a single request and one averaging 1ms over 999: the mean of those two averages is 50.5ms, and the true mean is 1.1ms. Combining per-group averages requires the group sizes, and forgetting them is how dashboards end up reporting numbers no user ever experienced.

## Task

Implement both functions in [latencyavg.go](latencyavg.go):

1. `Mean` returns the arithmetic mean; no samples gives `0`.
2. `WeightedMean` weights each value by its count, ignoring pairs beyond the shorter slice and any non-positive weight.
3. A total weight of zero gives `0`.

## Examples

**Example 1:**

```
Input:  Mean([1 2 3])
Output: 2
```

**Example 2:**

```
Input:  WeightedMean([10 20], [1 3])
Output: 17.5
```

**Example 3:**

```
Input:  WeightedMean([100 1], [1 999])
Output: 1.099
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A mean needs its denominator** | Aggregating averages without counts silently reweights the data. |
| 2 | **Ragged inputs** | Pairing two slices means iterating the shorter of the two. |
| 3 | **The mean hides the tail** | Even a correct mean says nothing about the slow requests; that is what percentiles are for. |

## Topics used again

Slices, `min`, float division, guards.

## Hint

`min(len(values), len(weights))` bounds the loop.

## Validate

```bash
make verify
```
