# Accumulate Locally, Publish Rarely

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A request counter uses one atomic increment per event. At several million events a second the atomic's cache line bounces between every core in the machine.

## Task

Implement [batchedcounter.go](batchedcounter.go):

1. Add `n` to the caller's local accumulator.
2. Publish into the shared total when the local reaches `batchSize` in either direction, then reset it.
3. No increment may be lost: `Flush` plus the batched publishes must account for everything.
4. Correct under concurrent use, with one `Local` per goroutine.

Replace the stub body in [batchedcounter.go](batchedcounter.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  63 Adds of 1, then Total()
Output: 0
```

_Explanation:_ Still local.

**Example 2:**

```
Input:  64 Adds of 1
Output: the total is 64
```

_Explanation:_ The batch published.

**Example 3:**

```
Input:  16 workers x 1000 Adds, each flushed
Output: 16000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Contention scales with sharing** | Batching cuts atomic traffic by the batch size. |
| 2 | **Private state needs no synchronisation** | The `Local` belongs to one goroutine. |
| 3 | **Threshold in both directions** | A negative accumulator must publish too. |
| 4 | **Flush is part of the contract** | A partial batch is only counted when the caller flushes. |

## Hint

Add to the local, then decide whether it is time to publish.

## Validate

```bash
make verify
```
