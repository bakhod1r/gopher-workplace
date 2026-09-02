# Latency Peak

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

An HTTP server tracks the slowest request it has served, in milliseconds. Every handler goroutine reports its own latency. Two handlers could both read the old peak and both overwrite it, so raising the peak has to be a compare-and-swap retry loop.

## Task

Implement the stubbed functions in [latencypeak.go](latencypeak.go) so that:

1. `Observe` raises the recorded peak to ms when ms is larger.
2. Faster requests leave the peak untouched.
3. `Peak` returns the current peak (0 before any request).

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var p PeakTracker; p.Observe(5); p.Peak()
Output: 5
```

**Example 2:**

```
Input:  var p PeakTracker; p.Observe(5); p.Observe(3); p.Peak()
Output: 5
```

**Example 3:**

```
Input:  var p PeakTracker; p.Peak()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **CompareAndSwap** | Swaps only if the value still equals the one you read; returns false if someone beat you. |
| 2 | **Retry loop** | On failure, re-read and try again - that is the lock-free idiom. |
| 3 | **atomic.Int64** | `Load` plus `CompareAndSwap` gives read-modify-write without a mutex. |

## Hint

Loop: `cur := p.ms.Load()`; if `ms <= cur` return; if `p.ms.CompareAndSwap(cur, ms)` return; otherwise loop again.

## Validate

```bash
make verify
go test -race ./...
```
