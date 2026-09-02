# Cancellable Event Publish

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The orders service publishes domain events to Kafka through a buffered channel drained by a writer goroutine. When the broker is slow the buffer fills, and a handler that blocks on the send holds its request goroutine and its database transaction open indefinitely. The send has to be bounded by the request context.

## Task

Implement the exported function(s) in [publishevent.go](publishevent.go) so that:

1. It selects on `ctx.Done()` and on the send `out <- event`.
2. On a successful send it returns `nil`.
3. On context completion it returns `ctx.Err()` and publishes nothing.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Publish(live ctx, chan with room, "order.created")
Output: nil
```

**Example 2:**

```
Input:  Publish(cancelled ctx, full chan, ...)
Output: context.Canceled
```

**Example 3:**

```
Input:  Publish(expired ctx, full chan, ...)
Output: context.DeadlineExceeded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sends in a `select`** | `case out <- v:` makes a send cancellable, exactly like a receive. |
| 2 | **Channel backpressure** | A full buffer blocks the producer — that is the signal, not a bug. |
| 3 | **`ctx.Err()` propagation** | Tells the caller whether to retry or to drop the event. |

## Hint

A send can be a `select` case: `case out <- event:`. Pair it with `case <-ctx.Done():`.

## Validate

```bash
make verify
```
