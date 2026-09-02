# Stopping a Worker Pool

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The image service keeps a pool of resize workers parked on the pool context while the queue is empty. A rolling deploy cancels that context, and the supervisor must not exit until every worker has actually stopped — otherwise in-flight temp files are left behind. Each worker records why it stopped so the deploy log can prove it was an orderly shutdown.

## Task

Implement the exported function(s) in [workerpool.go](workerpool.go) so that:

1. It derives one cancellable context from `context.Background()` and starts `n` goroutines that block on `ctx.Done()`.
2. Each worker `i` writes `ctx.Err()` into `reasons[i]`.
3. It cancels the context, waits for all workers with a `sync.WaitGroup`, and returns the slice.
4. It must be race-free under `-race`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  StopWorkers(0)
Output: []
```

**Example 2:**

```
Input:  StopWorkers(1)
Output: [context.Canceled]
```

**Example 3:**

```
Input:  StopWorkers(3)
Output: [context.Canceled, context.Canceled, context.Canceled]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cancellation as a broadcast** | Closing one channel wakes every waiting worker at once. |
| 2 | **`sync.WaitGroup`** | `Add` before the goroutine starts, `defer wg.Done()` inside, `wg.Wait()` after. |
| 3 | **Race-free slice writes** | Distinct indices of a pre-sized slice may be written concurrently. |

## Hint

Pre-size `reasons` with `make([]error, n)` so each worker owns its own index — no mutex needed.

## Validate

```bash
make verify
```
