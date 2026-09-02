# In-Flight Request Gauge

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A load balancer routes to whichever backend has the fewest requests in flight. Each handler bumps the gauge on entry and drops it on exit, from a different goroutine every time, and the router reads the gauge continuously.

## Task

Implement the stubbed functions in [inflight.go](inflight.go) so that:

1. `Enter` increases the in-flight count by one.
2. `Exit` decreases it by one.
3. `Current` returns the number of requests currently in flight.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var g Gauge; g.Enter(); g.Current()
Output: 1
```

**Example 2:**

```
Input:  var g Gauge; g.Enter(); g.Exit(); g.Current()
Output: 0
```

**Example 3:**

```
Input:  var g Gauge; g.Enter(); g.Enter(); g.Current()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Int64** | A gauge goes up and down; `Add(1)` and `Add(-1)` are both indivisible. |
| 2 | **Gauge vs counter** | A counter only rises; a gauge tracks a current level. |
| 3 | **defer with Exit** | Callers pair `Enter` with `defer g.Exit()` so the gauge cannot leak. |

## Hint

`Enter` is `g.n.Add(1)`, `Exit` is `g.n.Add(-1)`, `Current` is `g.n.Load()`.

## Validate

```bash
make verify
go test -race ./...
```
