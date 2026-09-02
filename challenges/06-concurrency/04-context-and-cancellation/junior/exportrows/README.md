# Cancellable Export Pipeline

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The admin console exports account rows as CSV. Rendering runs in a producer goroutine feeding a channel the handler drains. If the admin closes the tab mid-export, both halves have to stop: the consumer must return promptly, and the producer must not be left blocked on a send forever — that is exactly how a service accumulates leaked goroutines over a week.

## Task

Implement the exported function(s) in [exportrows.go](exportrows.go) so that:

1. It starts a goroutine that renders each ID as `fmt.Sprintf("row-%d", id)` and sends it on a channel, closing the channel when done.
2. The producer's send is itself inside a `select` with `ctx.Done()` so it can never block forever.
3. The consumer loops on a `select` over `ctx.Done()` and the channel, appending rows until the channel closes.
4. On cancellation it returns `nil` and `ctx.Err()`; otherwise the rows and `nil`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ExportRows(live ctx, []int{1, 2})
Output: ["row-1", "row-2"], nil
```

**Example 2:**

```
Input:  ExportRows(live ctx, nil)
Output: [], nil
```

**Example 3:**

```
Input:  ExportRows(cancelled ctx, []int{1})
Output: nil, context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Producer/consumer with a context** | Both ends must watch `ctx.Done()`, not just the consumer. |
| 2 | **Goroutine leaks** | A producer blocked on a send that nobody will ever receive never exits. |
| 3 | **`defer close(out)`** | The producer owns the channel and is the only one allowed to close it. |

## Hint

Give the producer a `select` around its send, and `defer close(out)` so the consumer's comma-ok receive terminates the loop.

## Validate

```bash
make verify
```
