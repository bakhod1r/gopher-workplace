# Query Timeout Budget

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The orders service gives each request a 200 ms budget and passes what is left of it to every database query. On a loaded box the budget is sometimes already gone before the query is dispatched. The CI test for that path used to assert "the call took at least 200 ms" and flaked constantly; the reliable assertion is the *error value*, reproduced with a budget that is already spent.

## Task

Implement the exported function(s) in [querytimeout.go](querytimeout.go) so that:

1. It derives a context from `context.Background()` with a timeout of `0` (or a negative duration).
2. It waits on `ctx.Done()`.
3. It returns `ctx.Err()`, which is `context.DeadlineExceeded`.
4. It must not call `time.Sleep`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ExhaustedQueryBudget()
Output: context.DeadlineExceeded
```

**Example 2:**

```
Input:  errors.Is(ExhaustedQueryBudget(), context.DeadlineExceeded)
Output: true
```

**Example 3:**

```
Input:  errors.Is(ExhaustedQueryBudget(), context.Canceled)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`context.WithTimeout`** | Deadline = now + duration; a duration of 0 is already past. |
| 2 | **`context.DeadlineExceeded`** | The sentinel for "the budget ran out", distinct from `Canceled`. |
| 3 | **`<-ctx.Done()`** | Blocks until the context finishes; here it returns immediately. |

## Hint

`context.WithTimeout(context.Background(), 0)` is already expired. Receive from `ctx.Done()`, then return `ctx.Err()`.

## Validate

```bash
make verify
```
