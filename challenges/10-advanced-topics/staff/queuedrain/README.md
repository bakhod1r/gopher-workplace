# Close The Queue Without Leaking Its Workers

**Level:** staff
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A service creates a queue per request and closes it in a defer. Goroutine count climbs all day and the process is restarted every night to keep it under control.

## Task

Implement [queuedrain.go](queuedrain.go):

1. Stop the workers, wait for the already-queued values to be processed, and return the total.
2. No worker goroutine may survive `Close`.
3. A second `Close` must return the same total without panicking.

Replace the stub body in [queuedrain.go](queuedrain.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  push 1..100 into NewQueue(4), Close()
Output: 5050
```

**Example 2:**

```
Input:  Close() twice
Output: the same total, no panic
```

_Explanation:_ Closing a closed channel panics; guard it.

**Example 3:**

```
Input:  16 values, one worker
Output: 16
```

_Explanation:_ Close returns only after the backlog drains.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **range over a channel** | The loop ends when the channel is closed and drained — that is the exit signal. |
| 2 | **Closing is the sender's job** | Only the side that stops sending may close. |
| 3 | **WaitGroup as the join** | `Wait` is what makes the accumulated total safe to read. |
| 4 | **sync.Once** | Idempotent shutdown without a flag and a mutex. |

## Hint

A `for range` over a channel exits on exactly one event. Nothing else will ever wake those workers.

## Validate

```bash
make verify
```
