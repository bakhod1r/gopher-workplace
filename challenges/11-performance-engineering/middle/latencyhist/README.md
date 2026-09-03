# Counting Into Buckets

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Keeping every latency sample costs memory proportional to traffic. A histogram costs a fixed number of counters and still answers "what does the distribution look like" — which is why every metrics system, Prometheus included, stores buckets rather than samples. The overflow bucket is what stops one pathological request from needing an unbounded array.

## Task

Implement both functions in [latencyhist.go](latencyhist.go):

1. `Histogram` counts into `n` buckets of `width`, starting at 0, plus one overflow bucket — `n+1` entries in total.
2. Drop negative samples; a non-positive `width` or `n` gives an empty, non-nil slice.
3. `Busiest` returns the index of the fullest bucket, earliest on a tie, and `-1` when all are empty.

## Examples

**Example 1:**

```
Input:  Histogram([0 5 15], 10, 1)
Output: [2 1]
```

**Example 2:**

```
Input:  Histogram([0 9.99 10 19.99 20 29.99 30 1000], 10, 3)
Output: [2 2 2 2]
```

**Example 3:**

```
Input:  Busiest([1 5 5])
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fixed memory, approximate answers** | A histogram trades exact percentiles for a bounded footprint. |
| 2 | **Half-open buckets** | `[0,10)` and `[10,20)` — a sample landing exactly on a boundary belongs to the upper bucket. |
| 3 | **The overflow bucket** | Without it the tail either gets clamped silently or needs unbounded storage. |

## Topics used again

Float-to-int conversion, slices, `min`, guards.

## Hint

`int(v / width)` gives the bucket; clamp anything at or past `n` into the overflow slot.

## Validate

```bash
make verify
```
