# Streaming Aggregate

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A daily job aggregates 100M readings. The current version loads every reading into a slice before computing, and the box runs out of memory.

## Task

Implement the stub(s) in [streamagg.go](streamagg.go):

1. Implement `Add` and `Result` on `*MeanAgg` and `*MaxAgg` — constant memory per aggregator.
2. Implement `Aggregate`, which drains a `Source` through an `Aggregator` and returns the result.
3. Constraint: memory must not grow with the number of readings. `Aggregate` over 1M readings must allocate a bounded number of times, checked by the test.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  mean of [1 2 3]
Output: 2
```

**Example 2:**

```
Input:  max of [1 5 2]
Output: 5
```

**Example 3:**

```
Input:  mean of an empty stream
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Streaming aggregation** | Keep running state, never the input. |
| 2 | **Constant memory** | Sum and count are two ints regardless of stream length. |
| 3 | **Interface-driven pipelines** | Reused: the source and the aggregator are both contracts. |

## Hint

A mean needs only a running sum and a count — never the values themselves.

## Validate

```bash
make verify
```
