# Dividing By The Wrong Iterations

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

A function that allocates exactly once per call reports 0 allocs/op, and raising the warmup count makes the number look even better. The counters cover the measured iterations; the divisor does not.

## Task

Fix the single planted bug in [reportallocsscopebug.go](reportallocsscopebug.go):

1. Find and fix the one bug so both columns divide by the iterations the counters actually cover.
2. The reported numbers must not change when the warmup count changes.
3. A run with no measured iterations must still report zeros.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Run{Warmup: 90, Measured: 10, Bytes: 800, Allocs: 20}
Output: 80 B/op, 2 allocs/op
```

**Example 2:**

```
Input:  the same run with Warmup 1000000
Output: unchanged
```

**Example 3:**

```
Input:  Run{Measured: 0, ...}
Output: 0, 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The counters have a scope** | They are reset with the timer, so they describe the measured phase only. |
| 2 | **Numerator and denominator must agree** | Both have to cover the same iterations, or the ratio is meaningless. |
| 3 | **Understating allocations is the dangerous direction** | A benchmark that reports 0 allocs/op ends arguments. |

## Hint

Which iterations do `Bytes` and `Allocs` cover?

## Validate

```bash
make verify
```
