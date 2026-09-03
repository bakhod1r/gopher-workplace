# Graceful Shutdown Drain

**Level:** junior
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An HTTP server refuses new connections on SIGTERM, then waits for the requests already running to finish before the process exits. Handlers mark themselves in and out, and shutdown — possibly from more than one place — blocks until the server is idle.

## Task

Implement the stubbed functions in [drainwait.go](drainwait.go) so that:

1. `NewDrain` builds the Drain and attaches a `sync.Cond` to its mutex.
2. `Start` and `Done` adjust the in-flight count under the lock.
3. `Wait` blocks until the count reaches zero; every waiter must wake.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewDrain().Inflight()
Output: 0
```

**Example 2:**

```
Input:  d.Start(); d.Inflight()
Output: 1
```

**Example 3:**

```
Input:  NewDrain().Wait()
Output: returns immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.Cond` | A waiting room attached to a mutex: `Wait` releases the lock and sleeps, `Signal`/`Broadcast` wakes sleepers. |
| 2 | `for`, never `if` | `Wait` may return without the condition holding. Re-check the predicate in a loop. |
| 3 | `Broadcast` vs `Signal` | Several goroutines may be waiting for the same event — `Signal` would wake only one and strand the rest. |

## Hint

`sync.NewCond(&d.mu)` ties the waiting room to the mutex you already have. `Wait` must be called with the lock held, inside `for d.inflight > 0`.

## Validate

```bash
make verify
```
