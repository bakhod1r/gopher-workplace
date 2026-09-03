# Aggregate In Parallel Without Sharing A Word

**Level:** staff
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A parallel histogram increments a shared bucket slice from every worker. It is fast, wrong, and wrong differently on every run.

## Task

Implement [workerpartials.go](workerpartials.go):

1. Count `data` into `buckets` bins by value modulo `buckets`, using `workers` goroutines over disjoint chunks.
2. Negative values must land in a valid bin.
3. Each worker accumulates privately; fold the partials after the join.
4. `buckets < 1` returns nil; `workers` outside `[1, len(data)]` is clamped.

Replace the stub body in [workerpartials.go](workerpartials.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Histogram([]int{0,1,2,3}, 2, 2)
Output: [2 2]
```

**Example 2:**

```
Input:  Histogram([]int{-1,-2,-3}, 3, 2)
Output: [1 1 1]
```

_Explanation:_ Go's `%` keeps the sign of the dividend.

**Example 3:**

```
Input:  Histogram([]int{1}, 0, 2)
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Private accumulation** | Per-worker state removes contention and the need for atomics entirely. |
| 2 | **Fold after join** | `wg.Wait()` is the happens-before edge that makes the partials safe to read. |
| 3 | **Chunk arithmetic** | Ceiling-divided chunks, clamped, cover the input exactly once. |
| 4 | **Go's modulo sign** | `-1 % 3` is -1, not 2 — the bin must be corrected. |

## Hint

Shared counters need atomics; private counters need nothing. Which is cheaper to fold once at the end?

## Validate

```bash
make verify
```
