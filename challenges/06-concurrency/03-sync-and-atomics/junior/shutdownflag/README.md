# Graceful Shutdown Flag

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

On SIGTERM a server must stop accepting work and drain. Every worker checks a shared shutdown flag on each loop iteration while the signal handler flips it - and exactly one goroutine must win the right to run the drain sequence.

## Task

Implement the stubbed functions in [shutdownflag.go](shutdownflag.go) so that:

1. `Request` marks shutdown as requested.
2. `Requested` reports whether shutdown has been requested.
3. `ClaimShutdown` requests shutdown and reports whether *this* call was the one that started it.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var f ShutdownFlag; f.Requested()
Output: false
```

**Example 2:**

```
Input:  var f ShutdownFlag; f.Request(); f.Requested()
Output: true
```

**Example 3:**

```
Input:  var f ShutdownFlag; f.ClaimShutdown(); f.ClaimShutdown()
Output: true, then false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **atomic.Bool** | `Load`, `Store` and `CompareAndSwap` on a boolean. |
| 2 | **CompareAndSwap** | `CompareAndSwap(false, true)` succeeds for exactly one caller. |
| 3 | **Flag pattern** | A shared boolean polled in a loop must be atomic, not plain. |

## Hint

`ClaimShutdown` is `f.down.CompareAndSwap(false, true)` - return its result directly.

## Validate

```bash
make verify
go test -race ./...
```
