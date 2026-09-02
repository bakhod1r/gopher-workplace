# Per-Handler Context Scope

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A profiler showed the API server holding thousands of live cancellable contexts hours after the requests had finished: handlers fanned out to the cache and the search index but nobody released the per-request context. Every derived context holds an entry in its parent's child list until cancelled, so the fix is the discipline every Go server applies — `defer cancel()`.

## Task

Implement the exported function(s) in [handlerscope.go](handlerscope.go) so that:

1. It derives a cancellable context from `context.Background()`.
2. It arranges for that context to be cancelled when `ServeRequest` returns, on every path.
3. It returns the error from `handler`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ServeRequest(func(ctx context.Context) error { return nil })
Output: nil
```

**Example 2:**

```
Input:  ServeRequest(func(ctx context.Context) error { return errDB })
Output: errDB
```

**Example 3:**

```
Input:  captured ctx after ServeRequest returns → ctx.Err()
Output: context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`defer cancel()`** | Releases a derived context's resources on every return path. |
| 2 | **Context leaks** | A derived context stays attached to its parent until cancelled. |
| 3 | **ctx as the first parameter** | `handler(ctx)` follows the standard convention. |

## Hint

Two lines before the call: create the context, then `defer cancel()`. Return `handler(ctx)`.

## Validate

```bash
make verify
```
