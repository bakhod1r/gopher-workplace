# Object Store Upload Window

**Level:** middle
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A backup agent pushes objects to blob storage, which throttles an account above a fixed number of concurrent uploads. The agent keeps a window of permits: a goroutine that finds the window full blocks until a running upload gives its permit back.

## Task

Implement the stubbed functions in [uploadwindow.go](uploadwindow.go) so that:

1. `NewWindow` clamps the limit to at least 1 and attaches a `sync.Cond` to the mutex.
2. `Acquire` waits while the window is full, then takes a permit.
3. `Release` gives a permit back and wakes waiting goroutines.
4. Concurrency must never exceed the limit, and no waiter may be stranded.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewWindow(3).InUse()
Output: 0
```

**Example 2:**

```
Input:  w := NewWindow(1); w.Acquire(); w.InUse()
Output: 1
```

**Example 3:**

```
Input:  w.Release(); w.InUse()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.Cond` as a gate | `Wait` sleeps and releases the lock; `Broadcast` wakes every sleeper to re-test the predicate. |
| 2 | Predicate in a `for` | `for w.inUse >= w.limit { Wait() }` — the state may change again between the wake and the relock. |
| 3 | Why not a buffered channel | A channel semaphore works too, but a `Cond` lets `InUse` and the limit live in one lock-protected struct. |
| 4 | Broadcast over Signal | Several uploaders may be queued; `Signal` wakes one and leaves the accounting correct but the wake-ups fragile. |

## Hint

`sync.NewCond(&w.mu)` in the constructor. `Acquire` loops on the predicate before incrementing; `Release` decrements then broadcasts.

## Validate

```bash
make verify
```
