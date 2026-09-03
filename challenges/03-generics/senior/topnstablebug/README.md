# Top-N That Sorts The Caller's Slice

**Level:** senior  
**Topic:** 03-generics

## Context

A "top 3" widget is fed a slice that the caller also renders in arrival order. After the widget runs, the arrival-order list is sorted by score too.

## Task

Fix the single planted bug in [topnstablebug.go](topnstablebug.go):

1. Find and fix the single bug so the caller's slice keeps its original order.
2. The ranking and the clamping of `n` must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  TopN(rows, score, 2)
Output: the two best rows
```

**Example 2:**

```
Input:  rows after the call
Output: unchanged order
```

**Example 3:**

```
Input:  TopN(rows, score, -1)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate your internals. |
| 2 | **slices.SortStableFunc sorts in place** | It has no return value: it rewrites the slice you hand it. |

## Hint

Which slice does the sort rewrite?

## Validate

```bash
make verify
```
