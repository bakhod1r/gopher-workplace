# Split The Work, Not The Memory

**Level:** staff
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A parallel aggregation was written by copying each worker's slice "to be safe". It is slower than the serial version and allocates a second copy of the dataset.

## Task

Implement [chunkworkers.go](chunkworkers.go):

1. Sum `s` across `workers` goroutines over disjoint chunks.
2. Every element must be counted exactly once, for any worker count.
3. Pass views into `s`, never copies — under 64 KiB allocated for an 8 MiB input.
4. `workers < 1` behaves as 1; more workers than elements is legal.

Replace the stub body in [chunkworkers.go](chunkworkers.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  SumParallel([]int{1,2,3,4}, 2)
Output: 10
```

**Example 2:**

```
Input:  SumParallel([]int{5}, 8)
Output: 5
```

_Explanation:_ Extra workers must not double-count or panic.

**Example 3:**

```
Input:  SumParallel([]int{1,2,3}, 0)
Output: 6
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Disjoint views** | Concurrent reads of one array need no synchronisation at all. |
| 2 | **Per-worker accumulators** | One shared counter would serialise every element. |
| 3 | **Chunk boundaries** | Ceiling-divided chunks must be clamped so the last one stops at `len(s)`. |
| 4 | **Join before reading** | `wg.Wait()` is the happens-before edge that makes the partials visible. |

## Hint

Reading the same array from many goroutines is free. Writing one accumulator from many goroutines is not.

## Validate

```bash
make verify
```
