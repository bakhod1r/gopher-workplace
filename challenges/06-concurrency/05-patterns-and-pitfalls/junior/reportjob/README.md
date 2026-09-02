# Cancellable Report Job

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Report generation is kicked off by an HTTP handler, and the handler's context
is cancelled the moment the client disconnects or the request deadline passes.
A job that ignores its context keeps burning CPU for a response nobody will
ever read.

## Task

Implement `RunReport` in [reportjob.go](reportjob.go) so that:

1. It checks `ctx.Err()` first and returns `0` with that error when the context is already done.
2. Otherwise it sums the rows concurrently, one goroutine per row, with the total guarded by a mutex.
3. It returns the total and a nil error after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RunReport(live ctx, []int{1, 2, 3})
Output: 6, nil
```

**Example 2:**

```
Input:  RunReport(live ctx, nil)
Output: 0, nil
```

**Example 3:**

```
Input:  RunReport(cancelled ctx, []int{1, 2, 3})
Output: 0, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **context.Context** | The standard way to carry cancellation across API boundaries. |
| 2 | **ctx.Err()** | Non-nil once the context is cancelled or its deadline has passed. |
| 3 | **Mutex-guarded total** | Concurrent `total += row` needs a lock, context or not. |

## Hint

One `ctx.Err()` check before any goroutine is started is enough here — return
the error it gives you rather than inventing your own.

## Validate

```bash
make verify
```
