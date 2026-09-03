# Cancellation Cause Reporting

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

When a report export is cancelled, the operator wants to know *why*: the client hung up, the quota ran out, or the row scan failed. `context.Canceled` alone cannot say. The exporter cancels with a cause so the API layer can report the real reason.

## Task

Implement the stubbed functions in [failurecause.go](failurecause.go) so that:

1. Derive a run context with `context.WithCancelCause`.
2. Stop at the quota, cancelling with `ErrQuotaExceeded` as the cause.
3. Report the reason with `context.Cause`, never the bare `context.Canceled`.
4. A parent that is already finished yields 0 rows and the parent's error.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Export(ctx, 5 rows, quota 10)
Output: 5, nil
```

**Example 2:**

```
Input:  Export(ctx, 5 rows, quota 3)
Output: 3, ErrQuotaExceeded
```

**Example 3:**

```
Input:  Export(cancelled ctx, rows)
Output: 0, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `WithCancelCause` | The cancel func takes an error; that error becomes the context's *cause*. |
| 2 | `context.Cause` | Returns the cause when one was set, otherwise the same value as `ctx.Err()` — a deadline still reports `DeadlineExceeded`. |
| 3 | `defer cancel(nil)` | Cancelling with a nil cause on the happy path releases resources without inventing a failure. |
| 4 | `errors.Is` at the boundary | The API layer distinguishes quota from disconnect by the cause, not by the string. |

## Hint

`runCtx, cancel := context.WithCancelCause(ctx)`, `defer cancel(nil)`. On the quota path call `cancel(ErrQuotaExceeded)` and return `context.Cause(runCtx)`.

## Validate

```bash
make verify
```
