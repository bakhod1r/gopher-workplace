# Cancellation Reaches the Query

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

An HTTP handler derives a narrower context for its database query so it can abandon just that query on retry. An incident review asked whether killing the request really tears the query down too, or whether the query keeps burning a connection after the client is gone. Write the experiment that answers it.

## Task

Implement the exported function(s) in [cancelpropagation.go](cancelpropagation.go) so that:

1. It creates a cancellable request context from `context.Background()`.
2. It derives a second cancellable query context from the request context.
3. It cancels the *request* context, waits on the *query* context's `Done()`, and returns the query context's `Err()`.
4. The result is `context.Canceled`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  QueryErrAfterRequestCancel()
Output: context.Canceled
```

**Example 2:**

```
Input:  errors.Is(result, context.Canceled)
Output: true
```

**Example 3:**

```
Input:  result == nil
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Context trees** | Cancelling a parent cancels every descendant, recursively. |
| 2 | **Propagated `Err()`** | The child reports the same sentinel the parent was cancelled with. |
| 3 | **`defer cancel()` on the child** | Still required, even when the parent will cancel it. |

## Hint

Cancel the parent, then wait on the *child's* `Done()` before reading the child's `Err()`.

## Validate

```bash
make verify
```
