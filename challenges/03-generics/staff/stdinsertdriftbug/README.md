# Insertion Points That Drift

**Level:** staff  
**Topic:** 03-generics

## Context

A diff renderer injects separator rows at line offsets computed from the original file. The first separator lands correctly and every later one creeps further up the page.

## Task

Fix the single planted bug in [stdinsertdriftbug.go](stdinsertdriftbug.go):

1. Find and fix the single bug so each position is interpreted against the original slice.
2. Out-of-range positions must still be skipped.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  InsertMarks([1,2,3,4], [1,3], 0)
Output: [1 0 2 3 0 4]
```

**Example 2:**

```
Input:  InsertMarks([1,2], [0,2], 9)
Output: [9 1 2 9]
```

**Example 3:**

```
Input:  InsertMarks([1,2], [5], 9)
Output: [1 2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Coordinate systems** | Positions computed before a mutation are stale after it. |
| 2 | **Accumulating offset** | After k insertions every original index has shifted right by exactly k. |

## Hint

`done` is being counted. Is it being used?

## Validate

```bash
make verify
```
