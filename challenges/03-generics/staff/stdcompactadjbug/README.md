# Compact Only Sees Its Neighbours

**Level:** staff  
**Topic:** 03-generics

## Context

A tag de-duplicator returns the right answer in every unit test and the wrong answer in production. The tests happen to feed it tags that arrive already grouped.

## Task

Fix the single planted bug in [stdcompactadjbug.go](stdcompactadjbug.go):

1. Find and fix the single bug so every case-insensitive duplicate is removed, not just the adjacent ones.
2. The caller's slice must not be modified.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  DistinctFold(["b","B","a"])
Output: ["a","b"]
```

**Example 2:**

```
Input:  DistinctFold(["a","b","A"])
Output: ["a","b"]
```

**Example 3:**

```
Input:  DistinctFold([])
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sorted-ness is a precondition** | Adjacency-based helpers only see neighbours, so unsorted input hides duplicates. |
| 2 | **Backing-array aliasing** | A slice value is a window onto storage someone else may also hold. |

## Hint

`CompactFunc` compares each element with the one immediately before it. What must be true of the input for that to be enough?

## Validate

```bash
make verify
```
