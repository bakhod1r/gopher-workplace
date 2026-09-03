# Two Definitions Of "The Median"

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Ask two tools for p50 of `[1 2 3 4]` and you get 2 or 2.5 depending on which definition they implement. Neither is wrong; they answer different questions. Interpolation gives a smooth estimate of the underlying distribution, nearest-rank gives a latency some request actually experienced. Knowing which one your dashboard uses is the difference between a real regression and a rounding artefact.

## Task

Implement both functions in [quantileinterp.go](quantileinterp.go):

1. `Interpolated` computes the zero-based position `(n-1) * p/100` and blends the two neighbouring samples linearly.
2. `NearestRank` returns the sample at rank `ceil(p/100 * n)`, counted from 1.
3. Both clamp `p` into `[0,100]`, return `0` for no samples, and must not modify the input.

## Examples

**Example 1:**

```
Input:  Interpolated([1 2 3 4], 50)
Output: 2.5
```

**Example 2:**

```
Input:  NearestRank([1 2 3 4], 50)
Output: 2
```

**Example 3:**

```
Input:  Interpolated([1 2 3 4], 25)
Output: 1.75
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interpolation invents values** | 2.5ms may be a fine estimate and is certainly not a request that happened. |
| 2 | **Position versus rank** | One is zero-based over `n-1` gaps, the other is one-based over `n` samples. |
| 3 | **Comparability beats correctness** | Two systems using different definitions cannot have their percentiles compared at all. |

## Topics used again

`slices.Clone`, `slices.Sort`, `math.Floor`, linear interpolation.

## Hint

Split the position into its integer and fractional parts; the fraction is the blend weight.

## Validate

```bash
make verify
```
