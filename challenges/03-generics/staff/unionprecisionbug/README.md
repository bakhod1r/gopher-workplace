# The Accumulator That Rounds Off The Floats

**Level:** staff  
**Topic:** 03-generics

## Context

One statistics helper serves counters and latencies alike. The integer instantiations are exact; the float64 one reports a mean of zero for every sub-second latency in the fleet.

## Task

Fix the single planted bug in [unionprecisionbug.go](unionprecisionbug.go):

1. Find and fix the single bug so the accumulator is wide enough for every member of the union.
2. The integer instantiations must stay exact.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Mean([]float64{0.5, 0.5, 0.5, 0.5})
Output: 0.5
```

**Example 2:**

```
Input:  Mean([]int{1, 2, 3, 4})
Output: 2.5
```

**Example 3:**

```
Input:  Mean([]float64{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A union is only as permissive as its accumulator** | The body must be correct for the *widest* member of the type set, not the most convenient one. |
| 2 | **Type sets** | A constraint names a *set* of types; the body must be correct for every member of it. |
| 3 | **Scale is a requirement** | A graded test asserts the result on millions of elements, so a defect that only shows past a threshold is caught. |

## Hint

What is `int64(0.5)`?

## Validate

```bash
make verify
```
