# Flush At The Threshold, Not At The End

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A metrics batcher accumulates until a timer fires. When the downstream stalls, the pending slice grows to gigabytes and the process is killed before the timer ever helps.

## Task

Implement [batchflush.go](batchflush.go):

1. Append `v` to the pending batch, flushing when it reaches the limit.
2. The pending slice must never hold `limit` or more values after `Add` returns.
3. The flushed batch must not alias the pending buffer — the callee may keep it.
4. Propagate the flush error; safe for concurrent use.

Replace the stub body in [batchflush.go](batchflush.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewBatcher(2, f); Add(1); Add(2)
Output: f received [1 2], Pending is 0
```

**Example 2:**

```
Input:  1000 adds with limit 4
Output: Pending stays under 4 throughout
```

**Example 3:**

```
Input:  a flush that fails
Output: Add returns the error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded accumulation** | The threshold, not a timer, is what caps the memory. |
| 2 | **Copy before handing over** | The pending buffer is reused; the batch is not. |
| 3 | **Reset with [:0]** | The pending buffer keeps its capacity across batches. |
| 4 | **Lock discipline** | One lock covers append, threshold check and reset. |

## Hint

Append, compare, and then three things happen in order.

## Validate

```bash
make verify
```
