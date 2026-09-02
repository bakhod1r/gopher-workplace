# Idempotent Shutdown

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The signal handler cancels the shutdown context on SIGTERM. Operators routinely follow up with Ctrl-C when the drain feels slow, so the handler fires again — and the same `cancel` func is called a second time. If that were unsafe, every Go server would crash during a slow deploy.

## Task

Implement the exported function(s) in [shutdowntwice.go](shutdowntwice.go) so that:

1. It derives a cancellable context from `context.Background()`.
2. It calls the cancel function twice.
3. It waits on `ctx.Done()` and returns `ctx.Err()`, which is `context.Canceled`.
4. It must not panic.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ShutdownTwice()
Output: context.Canceled
```

**Example 2:**

```
Input:  calling cancel twice
Output: no panic
```

**Example 3:**

```
Input:  the reason after the second call
Output: still context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Idempotent cancel funcs** | The second and later calls are no-ops by contract. |
| 2 | **Sticky `Err()`** | The first cancellation reason is latched and never overwritten. |
| 3 | **Closed channels** | `Done()` is closed once; the context guards against a double close. |

## Hint

Call `cancel()` twice in a row, then read `ctx.Err()`. The standard library makes that safe.

## Validate

```bash
make verify
```
