# The Filter That Keeps The Whole Batch Alive

**Level:** senior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

An ingest stage filters a 16k-record batch down to three records and hands them to a cache that lives for hours. Resident memory grows by a full batch for every cached result.

## Task

Fix the single planted bug in [filterretain.go](filterretain.go):

1. Return the records with `Size >= min`, in order.
2. The result must own its storage — dropping the batch must free the batch.
3. The input batch must not be modified.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Keep(batch, 100)
Output: the records of size >= 100
```

**Example 2:**

```
Input:  cap of the result
Output: the survivor count, not the batch size
```

_Explanation:_ Otherwise the batch cannot be collected.

**Example 3:**

```
Input:  Keep(nil, 1)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Allocation-granular collection** | The collector frees whole allocations; one live element pins all of them. |
| 2 | **In-place compaction has a cost** | It is allocation-free but returns a view of the input. |
| 3 | **Two-pass sizing** | Counting first gives an exactly-sized result with one allocation. |

## Hint

The compaction loop is efficient and wrong. What does `cap` of the returned slice tell you?

## Validate

```bash
make verify
```
