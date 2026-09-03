# Filter That Forgets To Shrink

**Level:** senior  
**Topic:** 03-generics

## Context

An in-place filter over a hot path leaves stale records at the end of every batch, and consumers keep reprocessing them.

## Task

Fix the single planted bug in [filterinplacebug.go](filterinplacebug.go):

1. Find and fix the single bug so the returned slice ends after the kept elements.
2. The kept elements must stay in order.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  FilterInPlace([1,2,3,4], even)
Output: [2 4]
```

**Example 2:**

```
Input:  FilterInPlace([1,3], even)
Output: []
```

**Example 3:**

```
Input:  FilterInPlace([2], even)
Output: [2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Length versus capacity** | Compaction writes to a prefix; the slice header must be re-sliced to match. |
| 2 | **Leftovers are visible** | Everything past the write cursor is stale data, not padding. |

## Hint

What is `n` for, if the return ignores it?

## Validate

```bash
make verify
```
