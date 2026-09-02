# Worker Pool

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A pool bounds concurrency: a fixed number of goroutines pull tasks from one
channel, so work queues instead of spawning unbounded goroutines. The pool owns
both the workers and the `WaitGroup` that tracks them.

## Task

Implement `Start` on `*Pool` in [poolworker.go](poolworker.go):

1. Launch `p.Count` goroutines.
2. Each worker ranges over `p.Tasks` and calls every task it receives.
3. Register each worker with `p.wg` **before** launching it, and mark it done
   when it exits.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Count 3; 5 tasks; close(Tasks); Wait()
Output: all 5 tasks ran
```

**Example 2:**

```
Input:  close(Tasks) with no tasks queued
Output: every worker exits, Wait() returns
```

**Example 3:**

```
Input:  more tasks than workers
Output: they queue in the channel and run as workers free up
```

_Explanation:_ `range` over a channel ends only when the channel is closed.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`range` over a channel** | Receives until the channel is closed and drained — that is the worker loop. |
| 2 | **`WaitGroup` ownership** | `Add` before `go`, `Done` (deferred) inside; adding inside the goroutine races with `Wait`. |
| 3 | **Bounded concurrency** | `Count` goroutines, not one per task. |

## Hint

`p.wg.Add(1)` belongs in the loop *before* `go func()`. Calling it inside the
goroutine lets `Wait` return before the worker has even registered.

## Validate

```bash
make verify
```
