# Sampling Fractions And What They Cost You

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`runtime.SetMutexProfileFraction(n)` records one contention event in every `n`, so a fraction of 100 gives you a cheap profile made of one percent of the truth. Reading it means multiplying back — and knowing that a site with three samples behind it is a rumour, not a measurement.

## Task

Implement the three functions in [mutexprofagg.go](mutexprofagg.go):

1. `Scale` multiplies a sampled count and delay by the fraction, reporting `ok = false` when the fraction is 0 or less, which means the profile was off.
2. `Estimate` aggregates the scaled delay per site, dropping records with a non-positive count or a negative delay, and returns an empty map when the profile was off.
3. `Confidence` classifies a sample count: `"low"` under 10, `"medium"` under 100, `"high"` otherwise.

## Examples

**Example 1:**

```
Input:  Scale(3, 300, 5)
Output: 15, 1500, true
```

**Example 2:**

```
Input:  Estimate([{a 1 100} {b 2 40} {a 1 100}], 5)
Output: {a:1000 b:200}
```

**Example 3:**

```
Input:  Estimate([{a 1 100}], 0)
Output: {} (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Off is not zero** | A disabled profile produces no data, which is not evidence of no contention. |
| 2 | **Estimates need sample counts** | The scaled number is only as trustworthy as the samples behind it. |
| 3 | **The fraction is a cost dial** | Lower fraction, better data, more overhead in every lock operation. |

## Topics used again

Multiple return values, map aggregation, `switch` with conditions, guards.

## Hint

`Estimate` can decide once, at the top, whether the profile was usable at all.

## Validate

```bash
make verify
```
