# Detached Audit Write

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

When a request is cancelled, its audit record must still be written — compliance does not care that the client hung up. The audit call therefore runs on a context that keeps the request's trace values but drops its cancellation.

## Task

Implement the stubbed functions in [detachedaudit.go](detachedaudit.go) so that:

1. Read the actor off the request context.
2. Give `write` a context that keeps the request's values but **not** its cancellation.
3. The audit write must succeed even when the request context is already cancelled or past its deadline.
4. Return the actor and whatever `write` reported.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Record(live ctx with "u1", write)
Output: "u1", nil
```

**Example 2:**

```
Input:  Record(cancelled ctx with "u1", write)
Output: "u1", nil
```

**Example 3:**

```
Input:  Record(ctx without actor, write)
Output: "", nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `context.WithoutCancel` | Returns a context sharing the parent's values whose `Done` never closes and whose `Err` stays nil. |
| 2 | Values outlive cancellation | Trace IDs and actors remain useful after the request is gone — cancellation is the only thing to shed. |
| 3 | Deliberate detachment | Detaching is a decision, not a default: unbounded detached work is how goroutines leak. |
| 4 | Comma-ok lookup | A missing actor yields the zero value rather than a panic. |

## Hint

`auditCtx := context.WithoutCancel(ctx)` — one line. Pass that to `write`, not the original context.

## Validate

```bash
make verify
```
