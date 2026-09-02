# Deploy Release Gate

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A deployment holds every request handler at a gate until the new configuration finishes loading. Handlers that arrive early must park; when the release opens the gate, all of them must be released at once — not one at a time.

## Task

Implement the stubbed functions in [releasegate.go](releasegate.go) so that:

1. `Wait` blocks until the gate is open, and returns immediately if it already is.
2. `Open` opens the gate and releases every waiting goroutine.
3. `IsOpen` reports the gate's state.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  g := NewGate(); g.IsOpen()
Output: false
```

**Example 2:**

```
Input:  g := NewGate(); g.Open(); g.IsOpen()
Output: true
```

**Example 3:**

```
Input:  g := NewGate(); g.Open(); g.Wait()
Output: returns immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Cond** | `Wait` parks the goroutine and releases the mutex while it sleeps. |
| 2 | **Broadcast** | One state change that unblocks *everyone* needs `Broadcast`, not `Signal`. |
| 3 | **Condition loop** | `for !g.open { g.cond.Wait() }` — re-check after every wake. |

## Hint

`Wait` locks, loops `for !g.open { g.cond.Wait() }`, unlocks. `Open` locks, sets the flag, unlocks, then `Broadcast`s.

## Validate

```bash
make verify
go test -race ./...
```
