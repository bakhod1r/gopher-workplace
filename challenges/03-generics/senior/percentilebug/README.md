# Percentile Zero Panics

**Level:** senior  
**Topic:** 03-generics

## Context

A dashboard panics whenever a panel is configured with p0, which the minimum-latency tile does by design.

## Task

Fix the single planted bug in [percentilebug.go](percentilebug.go):

1. Find and fix the single bug so `p = 0` returns the smallest sample.
2. Clamping `p` must stay; the empty-slice contract must stay.
3. Leave the input untouched.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Percentile([1,2,3,4], 0)
Output: 1, true
```

**Example 2:**

```
Input:  Percentile([1,2,3,4], 50)
Output: 2, true
```

**Example 3:**

```
Input:  Percentile([], 95)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ranks are 1-based** | The nearest rank is converted to an index with `rank-1`. |
| 2 | **Clamping the input is not enough** | A valid `p` of 0 still produces rank 0. |
| 3 | **Two guards, two purposes** | One keeps `p` in range; the other keeps the rank in range. |

## Hint

What is `ceil(0/100 * n)`?

## Validate

```bash
make verify
```
