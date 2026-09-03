# Explicit Bucket Bounds

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Fixed-width buckets waste resolution on latency data, where the interesting range spans microseconds to seconds. Real metrics systems take explicit ascending bounds — usually exponential — and store one counter per bound. Prometheus then exposes them cumulatively, because that form merges across instances by simple addition.

## Task

Implement all three in [bucketcounts.go](bucketcounts.go):

1. `Index` binary-searches for the first bound at or above `v`, returning `len(bounds)` for anything past the last one.
2. `Counts` tallies samples into `len(bounds)+1` buckets.
3. `Cumulative` turns per-bucket counts into running totals without modifying the input.

## Examples

**Example 1:**

```
Input:  Index([1 5 10], 3)
Output: 1
```

**Example 2:**

```
Input:  Counts([1 5 10], [1 5 10])
Output: [1 1 1 0]
```

**Example 3:**

```
Input:  Cumulative([1 2 3])
Output: [1 3 6]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Explicit bounds beat fixed width** | Exponential bounds give fine resolution where the data actually is. |
| 2 | **Bounds are inclusive upper limits** | A sample equal to a bound belongs to that bucket, not the next. |
| 3 | **Cumulative counts merge by addition** | Which is what makes histograms aggregatable across instances. |

## Topics used again

Binary search, slices, running totals.

## Hint

`sort.SearchFloat64s` finds the first index whose bound is `>= v` — that is already the answer.

## Validate

```bash
make verify
```
