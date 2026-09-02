# Or Done Tick Feed

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

A trading dashboard subscribes to a market data feed, but the user can close
the panel at any moment. If the forwarding goroutine only knows how to send,
it blocks forever the instant nobody reads — a classic leak. The *or-done*
wrapper makes both the receive and the send cancellable.

## Task

Implement `LiveTicks` in [tickfeed.go](tickfeed.go) so that:

1. It returns a new channel immediately and forwards ticks on it.
2. The receive from `ticks` is wrapped in a `select` with `done`, and the two-value form detects the feed closing.
3. The *send* on the output channel is also wrapped in a `select` with `done`, and the output channel is closed on every exit path.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  done open, ticks yields 1, 2, 3 then closes
Output: 1, 2, 3 then closed
```

**Example 2:**

```
Input:  done open, ticks closed with no values
Output: closed immediately
```

**Example 3:**

```
Input:  done already closed
Output: closed with no ticks
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Or-done channel** | Wrapping a stream so cancellation ends it from either side. |
| 2 | **Cancellable send** | A bare `out <- v` blocks forever if the consumer is gone; select on `done` too. |
| 3 | **defer close on all paths** | Every `return` inside the goroutine must still close the output. |

## Hint

Two selects, not one: the outer one guards the receive, the inner one guards
the send. `defer close(out)` covers both exits.

## Validate

```bash
make verify
```
