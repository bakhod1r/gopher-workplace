# Query Context Before and After Disconnect

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The analytics endpoint runs a multi-second Postgres query. `database/sql` passes the request context into the driver, which kills the query the moment that context finishes. A reviewer claims the driver can just check `ctx.Err() != nil` at any time to know whether the client is gone. Write the experiment that shows what `Err()` actually reports on both sides of the disconnect.

## Task

Implement the exported function(s) in [querystate.go](querystate.go) so that:

1. It derives a cancellable context from `context.Background()`.
2. `connected` is `ctx.Err()` read before cancelling — that is `nil`.
3. `disconnected` is `ctx.Err()` read after `cancel()` — that is `context.Canceled`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  connected from QueryState()
Output: nil
```

**Example 2:**

```
Input:  disconnected from QueryState()
Output: context.Canceled
```

**Example 3:**

```
Input:  connected == disconnected
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`Err()` transitions** | nil → a non-nil sentinel, exactly once, and never back. |
| 2 | **`context.WithCancel`** | The cancel func performs the transition. |
| 3 | **Named return values** | `(connected, disconnected error)` documents the pair in the signature. |

## Hint

Read `ctx.Err()` into `connected`, call `cancel()`, then read it again into `disconnected`.

## Validate

```bash
make verify
```
