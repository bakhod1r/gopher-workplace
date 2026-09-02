# License Seat Pool

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A desktop product sells a fixed number of concurrent seats. Logins arrive from many goroutines; the pool must hand out at most N seats even when everyone asks at the same instant, and it must do so without a lock on the hot path.

## Task

Implement the stubbed functions in [licenseseats.go](licenseseats.go) so that:

1. `TryAcquire` takes a seat and reports success, refusing when none are free.
2. `Release` returns a seat to the pool.
3. `Free` reports how many seats remain.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  p := NewSeatPool(2); p.TryAcquire()
Output: true, 1 seat left
```

**Example 2:**

```
Input:  p := NewSeatPool(1); p.TryAcquire(); p.TryAcquire()
Output: true, then false
```

**Example 3:**

```
Input:  p := NewSeatPool(1); p.TryAcquire(); p.Release(); p.Free()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **CompareAndSwap** | Decrement only if the count is still what you read; retry when it is not. |
| 2 | **Bounded decrement** | A plain `Add(-1)` cannot refuse — it would go negative. |
| 3 | **Retry loop** | Re-read the counter on every failed CAS. |

## Hint

`for { cur := p.free.Load(); if cur == 0 { return false }; if p.free.CompareAndSwap(cur, cur-1) { return true } }`.

## Validate

```bash
make verify
go test -race ./...
```
