# CAS Inventory Reservation

**Level:** middle  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

Sixteen checkout workers race for the last units of a hot SKU. Reserving is a read-decide-write sequence, and between the read and the write another worker may have taken the units you were counting on. Overselling a fulfilment centre means a cancelled order and a refund, so the decision and the write have to be one indivisible step.

## Task

Implement the exported function(s) in [warehouseslots.go](warehouseslots.go) so that:

1. `Reserve` returns `false` immediately when `n <= 0`.
2. `Reserve` loops: load the count, return `false` if fewer than `n` remain, otherwise `CompareAndSwap` from the observed value to `have-n`; retry the whole loop if the swap fails.
3. `Reserve` never lets `available` go negative, no matter how many goroutines race.
4. `Release` ignores a non-positive `n` and otherwise `Add`s the units back.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  s := NewStock(10); s.Reserve(3)
Output: true
```

**Example 2:**

```
Input:  s := NewStock(2); s.Reserve(5)
Output: false
```

**Example 3:**

```
Input:  s := NewStock(1); s.Reserve(1); s.Release(1); s.Available()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Compare-and-swap** | `CompareAndSwap(old, new)` only writes if the value is still `old`; a `false` return means someone beat you. |
| 2 | **Retry loop** | A failed CAS is not an error — reload and decide again with the fresh value. |
| 3 | **Check-then-act** | `Load` then `Store` is two steps and loses races; CAS fuses them into one. |

## Hint

`for { have := s.available.Load(); ...; if s.available.CompareAndSwap(have, have-n) { return true } }` — recompute `have` on every iteration, never outside the loop.

## Validate

```bash
make verify
go test -race ./...
```
