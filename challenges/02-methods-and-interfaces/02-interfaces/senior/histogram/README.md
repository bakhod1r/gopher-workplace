# Latency Histogram

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Percentiles over 100M latency samples must be computed without storing the samples.

## Task

Implement the stub(s) in [histogram.go](histogram.go):

1. Implement `Observe` on `*Histogram`, bucketing a value by upper bounds.
2. Implement `Count` and `Quantile`, where `Quantile(q)` returns the upper bound of the bucket containing the q-th sample.
3. Constraint: memory is O(buckets), independent of the sample count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  bounds [10 100 1000]; observe 5, 50, 500
Output: one sample in each bucket
```

**Example 2:**

```
Input:  Quantile(0.5) over those samples
Output: 100
```

**Example 3:**

```
Input:  Quantile on an empty histogram
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fixed-bucket histograms** | Bounded memory, approximate percentiles — the Prometheus model. |
| 2 | **Cumulative counts** | A quantile is a prefix scan over the bucket counts. |
| 3 | **Overflow bucket** | Values above the last bound still need a home. |

## Hint

Target rank = `ceil(q * total)`, then walk buckets accumulating counts.

## Validate

```bash
make verify
```
