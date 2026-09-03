# From ns/op To Requests Per Second

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Capacity planning starts with the benchmark you already have. A handler at 200 µs/op is 5000 requests per second per core — the number that tells you whether an optimisation is worth doing before you do it.

## Task

Implement both functions in [throughputcalc.go](throughputcalc.go):

1. `OpsPerSec` converts ns/op into operations per second on one core.
2. `Capacity` scales that across `cores`, assuming perfect scaling, and truncates to a whole number.
3. Non-positive inputs give `0`.

## Examples

**Example 1:**

```
Input:  OpsPerSec(1000)
Output: 1000000
```

**Example 2:**

```
Input:  Capacity(1000, 8)
Output: 8000000
```

**Example 3:**

```
Input:  Capacity(1000, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Rate is the reciprocal of latency** | One second holds `1e9 / nsPerOp` operations. |
| 2 | **Perfect scaling is a ceiling** | Real systems lose to contention, so the number is an upper bound. |
| 3 | **Truncate a capacity** | Rounding up promises throughput the machine cannot deliver. |

## Topics used again

Float arithmetic, int64 conversion, guards.

## Hint

`Capacity` should reuse `OpsPerSec` rather than repeating the formula.

## Validate

```bash
make verify
```
