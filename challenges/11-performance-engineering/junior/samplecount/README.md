# Samples Are Not Nanoseconds

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A CPU profile does not measure time; it counts interrupts. Each sample carries a count and a sampling period, and "1.2s" in the pprof output is a multiplication the tool did for you. Do it yourself once and the profile stops being magic.

## Task

Implement `Totals` in [samplecount.go](samplecount.go):

1. Return the total sample count.
2. Return the total nanoseconds: the sum of `Count * Period` over the samples.
3. A sample with a non-positive `Count` or `Period` contributes to neither total.

## Examples

**Example 1:**

```
Input:  [{a 3 10} {b 2 10}]
Output: 5, 50
```

**Example 2:**

```
Input:  [{a 2 10000000} {b 1 5000000}]
Output: 3, 25000000
```

**Example 3:**

```
Input:  [{a 0 10} {b 3 0}]
Output: 0, 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sampling, not tracing** | The profiler wakes ~100 times a second; everything between wakeups is inferred. |
| 2 | **Period is the multiplier** | Time equals samples times the period, so a low sample count means a noisy profile. |
| 3 | **Named results** | Returning two totals from one pass beats two traversals. |

## Topics used again

Named return values, structs, `range`.

## Hint

One loop, one guard, two accumulators.

## Validate

```bash
make verify
```
