# Merge That Loses Stability

**Level:** senior  
**Topic:** 03-generics

## Context

A merge step feeds a stable sort. Records with equal keys come out with the second stream ahead of the first, breaking the tie-break contract.

## Task

Fix the single planted bug in [mergetiebug.go](mergetiebug.go):

1. Find and fix the single bug so an element from `a` wins a tie.
2. The output must stay sorted and contain every element.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  MergeSorted([k1,k3],[k2], key)
Output: [k1 k2 k3]
```

**Example 2:**

```
Input:  equal keys
Output: the one from a comes first
```

**Example 3:**

```
Input:  MergeSorted([],[x], key)
Output: [x]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stability** | Equal elements must keep their input order unless the doc says otherwise. |
| 2 | **Which side is strict** | Taking from `b` on equality is exactly what breaks stability. |

## Hint

One comparison operator decides the tie.

## Validate

```bash
make verify
```
