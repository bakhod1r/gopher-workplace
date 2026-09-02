# Graceful Shutdown Wait

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

On SIGTERM the API server cancels a process-wide shutdown context. A background metrics flusher parks on that context; when it fires the flusher writes one last batch and records the shutdown reason, because operators need to tell an orderly SIGTERM apart from a drain window that ran out.

## Task

Implement the exported function(s) in [gracefulshutdown.go](gracefulshutdown.go) so that:

1. It blocks until `ctx.Done()` is closed.
2. It then returns `ctx.Err()`.
3. It must not use `time.Sleep` or poll `ctx.Err()` in a loop.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  WaitForShutdown(cancelled ctx)
Output: context.Canceled
```

**Example 2:**

```
Input:  WaitForShutdown(expired-timeout ctx)
Output: context.DeadlineExceeded
```

**Example 3:**

```
Input:  WaitForShutdown(child of cancelled ctx)
Output: context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`<-ctx.Done()`** | The blocking way to wait for cancellation — no polling, no sleeping. |
| 2 | **`ctx.Err()` after `Done()`** | Guaranteed non-nil once the channel is closed. |
| 3 | **Canceled vs DeadlineExceeded** | The same wait reports both; the error says which happened. |

## Hint

Receive from `ctx.Done()`, then return `ctx.Err()`.

## Validate

```bash
make verify
```
