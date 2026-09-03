# The Range That Overlaps Its Neighbour

**Level:** staff  
**Topic:** 03-generics

## Context

A booking calendar accumulates load over half-open time ranges. A single reservation reports correctly; two back-to-back reservations both claim the instant between them, so the handover slot looks double-booked and the daily total overshoots by one unit per reservation.

## Task

Fix the single planted bug in [intervalboundbug.go](intervalboundbug.go):

1. Find and fix the single bug so a range covers `lo` up to but not including `hi`.
2. An empty range must add nothing at all.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add(0,3,1); Add(3,6,1); At(3)
Output: 1
```

**Example 2:**

```
Input:  Add(0,3,1); Add(3,6,1); Total()
Output: 6
```

**Example 3:**

```
Input:  Add(4,4,5); At(4)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Half-open ranges** | `[lo, hi)` is the convention that makes adjacent ranges tile without overlap. |
| 2 | **Adjacency is the test case** | A single range hides the error; two ranges sharing a boundary expose it. |
| 3 | **Failures that need scale** | A defect that small inputs cannot express is still a defect; test at size. |

## Hint

Which of `lo` and `hi` belongs to the range?

## Validate

```bash
make verify
```
