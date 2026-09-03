# The Heap That Swaps With The Wrong Child

**Level:** staff  
**Topic:** 03-generics

## Context

A scheduler drains a priority queue of ten thousand jobs. The first few come out in order, then jobs start appearing out of turn — and the unit tests, which push four elements, are green.

## Task

Fix the single planted bug in [binheapchildbug.go](binheapchildbug.go):

1. Find and fix the single bug so `siftDown` always swaps with the *smaller* of the two children.
2. Draining the heap must yield a non-decreasing sequence for any input size.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Push 5,3,8,1 then drain
Output: 1 3 5 8
```

**Example 2:**

```
Input:  Push 2,2,1 then drain
Output: 1 2 2
```

**Example 3:**

```
Input:  10000 pseudo-random pushes then drain
Output: sorted ascending
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structural invariants** | Every operation must restore what the type promises about itself. |
| 2 | **Failures that need scale** | A defect that small inputs cannot express is still a defect; test at size. |
| 3 | **Choosing a child** | The candidate is decided by comparing the two children with each other, not with the parent. |

## Hint

The comparison that picks between the left and the right child mentions the wrong index.

## Validate

```bash
make verify
```
