# Shut Services Down in Reverse Order

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The API server starts the database pool, then the cache, then the HTTP listener. Shutdown must run the other way: stop accepting requests first, then the cache, then the pool — otherwise an in-flight request finds its database gone. The deploy system gives the process a bounded drain window, and the shutdown log needs to say exactly how far the sequence got.

## Task

Implement the exported function(s) in [shutdownorder.go](shutdownorder.go) so that:

1. It iterates `services` from the last index down to the first — reverse startup order.
2. Before each `Stop` it returns the names collected so far plus `ctx.Err()` if the drain window has finished.
3. It calls `Stop(ctx)` and, on error, returns the names collected so far plus that error without touching the remaining services.
4. On success it appends the service name and continues.
5. It always returns a non-nil slice, even when nothing was stopped.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ShutdownServices(ctx, [database, cache, http])
Output: ["http" "cache" "database"], nil
```

**Example 2:**

```
Input:  ShutdownServices(ctx, [database, stuckCache, http])
Output: ["http"], errStopFailed
```

**Example 3:**

```
Input:  ShutdownServices(cancelled ctx, [database, cache, http])
Output: [], context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shutdown ordering** | Dependencies are torn down in the reverse of the order they were built. |
| 2 | **Partial progress reporting** | The names collected so far are part of the answer, not something to discard on error. |
| 3 | **Drain window as a context** | One deadline covers the whole sequence and is re-checked before each step. |

## Hint

Walk the slice backwards with `for i := len(services) - 1; i >= 0; i--`, and check `ctx.Err()` at the top of each step so a closed drain window stops the sequence between services.

## Validate

```bash
make verify
```
