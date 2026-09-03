# Connection Cleanup Hook

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A websocket hub registers a teardown for every connection: when the request context ends, the socket is removed from the hub's fan-out list. A normal close must unregister the hook instead of letting it fire, so a reconnecting client is not evicted twice.

## Task

Implement the stubbed functions in [cleanuphook.go](cleanuphook.go) so that:

1. `Register` attaches the teardown with `context.AfterFunc` and keeps its stop function.
2. `Release` tears the connection down early and reports whether it stopped the hook in time.
3. A second `Release`, or one after the context already fired, reports `false`.
4. `Wait` blocks until teardown and reports who did it: `"context"` or `"release"`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  h := Register(ctx); cancel(); h.Wait()
Output: "context"
```

**Example 2:**

```
Input:  h := Register(ctx); h.Release()
Output: true
```

**Example 3:**

```
Input:  h := Register(cancelled ctx); h.Release()
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `context.AfterFunc` | Runs a function in its own goroutine once the context finishes — and returns a stop function. |
| 2 | The stop function's answer | `true` means the hook was cancelled before running; `false` means it already ran or was already stopped. |
| 3 | Exactly-once teardown | Both paths lead to the same close; `sync.Once` keeps a double `close(chan)` from panicking. |
| 4 | No polling | `AfterFunc` replaces a `go func(){ <-ctx.Done(); ... }()` goroutine and its lifetime problems. |

## Hint

`h.stop = context.AfterFunc(ctx, func(){ h.finish("context") })`. `Release` calls `h.stop()` and only finishes as `"release"` when that returned `true`.

## Validate

```bash
make verify
```
