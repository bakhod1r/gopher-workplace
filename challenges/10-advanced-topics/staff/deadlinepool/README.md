# Stop When The Deadline Passes

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A batch endpoint fans work out to a worker pool. When the client disconnects the handler returns, and the workers keep going until the batch is finished.

## Task

Implement [deadlinepool.go](deadlinepool.go):

1. Double every item using `workers` goroutines, results in input order.
2. Return `ctx.Err()` when the context is done, with a nil result.
3. Every goroutine must exit before `Process` returns, cancelled or not.
4. `workers < 1` behaves as 1; more workers than items is legal.

Replace the stub body in [deadlinepool.go](deadlinepool.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Process(ctx, []int{1,2,3}, 2)
Output: [2 4 6], nil
```

**Example 2:**

```
Input:  an already-cancelled context
Output: nil, context.Canceled
```

**Example 3:**

```
Input:  20 cancelled runs of 1000 items
Output: no goroutines left
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cancellation on every blocking operation** | Both the feed's send and the workers' receive need a `ctx.Done()` case. |
| 2 | **Disjoint slot writes** | `out[i]` from one worker needs no lock. |
| 3 | **Close then Wait** | Closing the index channel ends the workers on the normal path. |
| 4 | **Wait before returning** | It is what makes "no goroutine outlives the call" true. |

## Hint

Three places can block: the send of an index, the receive of one, and the join.

## Validate

```bash
make verify
```
