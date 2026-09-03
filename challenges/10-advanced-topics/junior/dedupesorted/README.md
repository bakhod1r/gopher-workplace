# Collapse Runs Without New Memory

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A sorted index is deduplicated into a new slice on every rebuild. The index is large, the duplicates are few, and the copy dominates the rebuild time.

## Task

Implement [dedupesorted.go](dedupesorted.go):

1. Collapse runs of equal elements in the sorted input `s`.
2. Reuse `s`'s storage and return the deduplicated prefix.
3. Zero allocations.

Replace the stub body in [dedupesorted.go](dedupesorted.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Dedupe([]int{1,1,2,3,3,3})
Output: [1 2 3]
```

**Example 2:**

```
Input:  Dedupe([]int{2,2,2})
Output: [2]
```

_Explanation:_ A single run collapses to one element.

**Example 3:**

```
Input:  Dedupe(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two-cursor compaction** | The write cursor never passes the read cursor, so overwriting is safe. |
| 2 | **Sorted input** | Duplicates are adjacent, so one comparison against the last kept element suffices. |
| 3 | **Prefix results** | `s[:k]` is the answer without a second array. |

## Hint

Compare each element with the last one you decided to keep, not with its neighbour.

## Validate

```bash
make verify
```
