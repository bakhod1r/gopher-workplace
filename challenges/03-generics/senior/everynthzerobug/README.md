# Sampling That Skips The First

**Level:** senior  
**Topic:** 03-generics

## Context

A metrics down-sampler is supposed to keep one point in ten. Every chart is missing its first data point, so ranges start late.

## Task

Fix the single planted bug in [everynthzerobug.go](everynthzerobug.go):

1. Find and fix the single bug so sampling starts at index 0.
2. A non-positive step must still yield an empty slice.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  EveryNth([0,1,2,3], 2)
Output: [0 2]
```

**Example 2:**

```
Input:  EveryNth([1,2,3], 1)
Output: [1 2 3]
```

**Example 3:**

```
Input:  EveryNth([1,2], 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Loop initialisation** | The starting index is a separate decision from the step. |
| 2 | **Guard first** | The non-positive step must be rejected before any loop. |

## Hint

Which index does the loop start from?

## Validate

```bash
make verify
```
