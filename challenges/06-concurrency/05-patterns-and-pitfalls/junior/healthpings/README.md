# Health Ping Generator

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The uptime checker probes each registered endpoint a fixed number of times per
round and feeds the probes into a worker pool. Rather than building a slice of
identical URLs, a *repeat generator* streams them, so the pool can start
working before the round is fully enumerated.

## Task

Implement `HealthPings` in [healthpings.go](healthpings.go) so that:

1. It creates a channel and returns its receive-only end immediately.
2. A goroutine sends `endpoint + "/health"` exactly `count` times.
3. The channel is closed afterwards; a `count` of zero or less closes it without sending anything.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  HealthPings("api", 2)
Output: "api/health", "api/health" then closed
```

**Example 2:**

```
Input:  HealthPings("db", 1)
Output: "db/health" then closed
```

**Example 3:**

```
Input:  HealthPings("api", 0)
Output: closed immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Repeat generator** | A producer whose output is derived from parameters, not from an input slice. |
| 2 | **Bounded stream** | The counter is what stops the generator; without it the goroutine runs forever. |
| 3 | **Close on the empty path** | A zero count must still close, or the consumer hangs. |

## Hint

A plain counted `for` loop inside the goroutine, with `defer close(out)`.
Negative counts simply never enter the loop.

## Validate

```bash
make verify
```
