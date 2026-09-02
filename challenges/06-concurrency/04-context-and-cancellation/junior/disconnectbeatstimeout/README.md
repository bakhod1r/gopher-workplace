# Disconnect Beats the Timeout

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Requests get a generous timeout, but most abandoned ones end far earlier — the client disconnects and the handler cancels. The access log must show a 499 (client closed request), not a 504 (gateway timeout), or the availability dashboard will page someone for a timeout that never happened. Prove which reason a context records when cancellation happens well inside the deadline.

## Task

Implement the exported function(s) in [disconnectbeatstimeout.go](disconnectbeatstimeout.go) so that:

1. It derives a context with a long timeout (an hour) from `context.Background()`.
2. It cancels that context immediately, then waits on `ctx.Done()`.
3. It returns `ctx.Err()`, which is `context.Canceled`, never `context.DeadlineExceeded`.
4. It must not sleep or wait for the real deadline.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  DisconnectDuringTimeout()
Output: context.Canceled
```

**Example 2:**

```
Input:  errors.Is(result, context.Canceled)
Output: true
```

**Example 3:**

```
Input:  errors.Is(result, context.DeadlineExceeded)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **First reason wins** | Whichever event finishes the context first latches its error, permanently. |
| 2 | **Timeout contexts are cancellable too** | `WithTimeout` returns a cancel func that works before the deadline. |
| 3 | **Canceled vs DeadlineExceeded** | A 499 and a 504 are very different lines in an SLO report. |

## Hint

`WithTimeout` for an hour, call `cancel()` right away, then read `ctx.Err()` — the deadline never gets a chance.

## Validate

```bash
make verify
```
