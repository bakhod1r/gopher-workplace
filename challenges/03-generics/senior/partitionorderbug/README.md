# Partition That Reverses One Half

**Level:** senior  
**Topic:** 03-generics

## Context

A validation step splits records into accepted and rejected. The rejected report lists failures in reverse, so operators fix the newest problem first.

## Task

Fix the single planted bug in [partitionorderbug.go](partitionorderbug.go):

1. Find and fix the single bug so both halves preserve input order.
2. Membership of each half must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Partition([1,2,3,4], even)
Output: [2 4], [1 3]
```

**Example 2:**

```
Input:  Partition([1,3], even)
Output: [], [1 3]
```

**Example 3:**

```
Input:  Partition([], even)
Output: [], []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Order of operations** | Doing the right steps in the wrong order is still a bug. |
| 2 | **Prepending reverses** | `append([]T{v}, rest...)` builds the slice backwards and reallocates every time. |

## Hint

Compare how the two halves are extended.

## Validate

```bash
make verify
```
