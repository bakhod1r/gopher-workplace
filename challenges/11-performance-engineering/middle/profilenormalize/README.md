# Making Two Profiles Comparable

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A 30-second profile has twice the samples of a 15-second one from the same binary doing the same work. Comparing the raw numbers measures how long you profiled, not what the code does. Normalising — to a rate, or to shares of the total — is the step that makes any cross-profile comparison meaningful.

## Task

Implement both functions in [profilenormalize.go](profilenormalize.go):

1. `Rate` divides each value by the duration in seconds, giving per-second rates; a non-positive duration gives an empty, non-nil map.
2. `Fractions` divides each value by the profile total, giving shares that sum to 1.
3. `Fractions` drops non-positive values, and a zero total gives an empty, non-nil map. Neither function modifies its input.

## Examples

**Example 1:**

```
Input:  Rate({a:60}, 30)
Output: {a:2}
```

**Example 2:**

```
Input:  Fractions({a:3 b:1})
Output: {a:0.75 b:0.25}
```

**Example 3:**

```
Input:  Fractions({a:0})
Output: {} (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two normalisations, two questions** | Rates compare absolute work; fractions compare shape. |
| 2 | **Shape survives a workload change** | Fractions stay stable when traffic doubles; rates do not. |
| 3 | **Normalise before diffing** | Otherwise the diff reports the difference in profiling duration. |

## Topics used again

Maps, float conversion, two-pass aggregation, guards.

## Hint

`Fractions` needs the total before it can emit anything — two passes, or one pass into a slice.

## Validate

```bash
make verify
```
