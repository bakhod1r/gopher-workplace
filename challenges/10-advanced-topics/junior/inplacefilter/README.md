# Filter Without A Second Slice

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A stream stage filters batches of a few hundred thousand records and allocates a fresh result slice for every batch. The batches are short-lived and the collector is doing all the work.

## Task

Implement [inplacefilter.go](inplacefilter.go):

1. Return the even elements of `s` in order.
2. Reuse `s`'s storage — the function must allocate nothing.

Replace the stub body in [inplacefilter.go](inplacefilter.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  KeepEven([]int{1,2,3,4,6})
Output: [2 4 6]
```

**Example 2:**

```
Input:  KeepEven([]int{1,3})
Output: []
```

_Explanation:_ Nothing survives; the length is 0.

**Example 3:**

```
Input:  KeepEven([]int{-2,-1,0})
Output: [-2 0]
```

_Explanation:_ 0 and negative evens count.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **In-place compaction** | A write index trails the read index and never overtakes it. |
| 2 | **Reslicing as the result** | `s[:k]` reports the kept count without allocating. |
| 3 | **Destructive helpers** | Reusing the caller's array is a documented contract, not a secret. |

## Hint

Two indices, one array: where you are reading and where you are writing.

## Validate

```bash
make verify
```
