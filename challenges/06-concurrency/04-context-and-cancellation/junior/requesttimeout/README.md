# Imposing a Request Timeout

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Every route in the API goes through a middleware that caps how long its handler may run, so one slow dependency cannot pin a connection indefinitely. The middleware needs one helper that derives the capped context and hands back the cancel func for the caller to defer.

## Task

Implement the exported function(s) in [requesttimeout.go](requesttimeout.go) so that:

1. It returns `context.WithTimeout(ctx, d)` — both values, unchanged.
2. A generous budget yields a live context that has a deadline.
3. A zero or negative budget yields a context that is already `DeadlineExceeded`.
4. A parent that is already cancelled makes the derived context report `context.Canceled`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  WithRequestTimeout(bg, time.Hour) → Err()
Output: nil, and Deadline() ok is true
```

**Example 2:**

```
Input:  WithRequestTimeout(bg, 0) → Err()
Output: context.DeadlineExceeded
```

**Example 3:**

```
Input:  WithRequestTimeout(cancelled parent, time.Hour) → Err()
Output: context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.CancelFunc`** | The named type for a cancel func, returned so the caller can defer it. |
| 2 | **Returning ctx and cancel together** | Ownership of the release moves to the caller. |
| 3 | **Parent wins** | A finished parent cancels the child regardless of the child's budget. |

## Hint

One line: `return context.WithTimeout(ctx, d)`.

## Validate

```bash
make verify
```
