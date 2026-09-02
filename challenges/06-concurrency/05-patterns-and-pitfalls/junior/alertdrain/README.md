# Alert Drain Shutdown

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The alerting consumer runs for the lifetime of the process, so it must not
block forever when the service is asked to shut down. The *done channel*
convention gives it a second thing to wait on: a channel that carries no data
and is only ever closed, which every goroutine can watch at once.

## Task

Implement `CountAlerts` in [alertdrain.go](alertdrain.go) so that:

1. It loops with a `select` over a receive from `alerts` and a receive from `done`.
2. Receiving from `alerts` with the two-value form detects the close and returns the count.
3. A closed `done` returns the count so far immediately, without draining the rest.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  done open, alerts yields 3 alerts then closes
Output: 3
```

**Example 2:**

```
Input:  done open, alerts closed with no alerts
Output: 0
```

**Example 3:**

```
Input:  done already closed, alerts still open
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Done channel** | `chan struct{}` closed to broadcast "stop" to every receiver at once. |
| 2 | **select** | Waiting on several channel operations, taking whichever is ready first. |
| 3 | **Comma-ok receive** | `v, ok := <-ch` distinguishes a real value from a closed channel. |

## Hint

A closed channel is always ready to receive, so `case <-done` fires every
time once closed — that is exactly the broadcast you want for shutdown.

## Validate

```bash
make verify
```
