# Circuit Breaker State

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A circuit breaker in front of a flaky upstream is either closed (traffic flows) or open (traffic is rejected). Many goroutines observe failures at once, but only the one that actually trips the breaker should log the outage and start the recovery timer.

## Task

Implement the stubbed functions in [circuitstate.go](circuitstate.go) so that:

1. `Trip` moves the breaker from closed to open and reports whether this call did it.
2. `Reset` moves it from open back to closed and reports whether this call did it.
3. `Open` reports whether the breaker is currently open.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var b Breaker; b.Open()
Output: false
```

**Example 2:**

```
Input:  var b Breaker; b.Trip(); b.Trip()
Output: true, then false
```

**Example 3:**

```
Input:  var b Breaker; b.Trip(); b.Reset(); b.Open()
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **State machine with CAS** | A transition is a `CompareAndSwap` from the expected state to the next one. |
| 2 | **Exactly-one transition** | The CAS return value tells you whether *you* made the change. |
| 3 | **atomic.Int32** | Small integer states swap atomically without a mutex. |

## Hint

`Trip` is `b.state.CompareAndSwap(closed, open)`; `Reset` is the same swap in the other direction.

## Validate

```bash
make verify
go test -race ./...
```
