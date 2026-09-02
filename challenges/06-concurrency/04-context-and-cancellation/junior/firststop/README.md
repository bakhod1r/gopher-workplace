# Request or Shutdown, Whichever Comes First

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A long-poll endpoint holds a connection open until an event arrives. Two independent things can end that wait: the client's own request context, and the process-wide shutdown context that the signal handler cancels during a deploy. The handler must react to whichever fires first and report that one's reason.

## Task

Implement the exported function(s) in [firststop.go](firststop.go) so that:

1. It selects on both `reqCtx.Done()` and `shutdownCtx.Done()`.
2. It returns the `Err()` of whichever context finished.
3. It must not poll or sleep.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  FirstStop(cancelled req, live shutdown)
Output: context.Canceled
```

**Example 2:**

```
Input:  FirstStop(live req, cancelled shutdown)
Output: context.Canceled
```

**Example 3:**

```
Input:  FirstStop(live req, expired shutdown)
Output: context.DeadlineExceeded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`select` over two `Done()` channels** | Waiting on several cancellation sources at once. |
| 2 | **Per-case error** | Each branch returns *its own* context's error, not a shared one. |
| 3 | **A never-ending source** | `context.Background().Done()` is nil and simply never wins the select. |

## Hint

Two cases, each `<-someCtx.Done()`, each returning that context's own `Err()`.

## Validate

```bash
make verify
```
