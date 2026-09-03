# The Median That Sits One Slot Too High

**Level:** staff  
**Topic:** 03-generics

## Context

A latency dashboard's p50 always reads one bucket pessimistic, and on small samples the reported median is visibly not the middle value. The p100 is correct, which is why nobody believed the report for two quarters.

## Task

Fix the single planted bug in [percentilerankbug.go](percentilerankbug.go):

1. Find and fix the single bug so the nearest-rank index is converted to a zero-based offset.
2. The clamps and the validity checks must keep working, and the caller's slice must stay untouched.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Percentile([]int{1,2,3,4,5}, 50)
Output: 3, true
```

**Example 2:**

```
Input:  Percentile([]int{1,2,3,4}, 25)
Output: 1, true
```

**Example 3:**

```
Input:  Percentile([]int{1,2,3}, 101)
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Rank versus index** | The nearest rank is 1-based; a slice offset is 0-based, so exactly one subtraction stands between them. |
| 2 | **Clamps hide off-by-ones** | The upper clamp makes p100 correct no matter what, which masks the defect at the top end. |
| 3 | **Scale is a requirement** | A graded test asserts the result on millions of elements, so a defect that only shows past a threshold is caught. |

## Hint

The rank of the median of five elements is 3. What index is that?

## Validate

```bash
make verify
```
