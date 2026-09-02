# Worker Pool Snapshot

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A worker pool exposes `/stats`, which must return started and failed counts that belong to the same instant. Two separate reads can be caught mid-update and produce a snapshot with more failures than starts — an impossible report that pages the on-call engineer for nothing.

## Task

Implement the stubbed functions in [workerstats.go](workerstats.go) so that:

1. `Start` and `Fail` record pool events under the write lock.
2. `Snapshot` returns both counters read under a single read lock.
3. The returned `Snapshot` is a plain value, safe to use after the lock is released.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var p Pool; p.Start(); p.Snapshot()
Output: Snapshot{Started: 1, Failed: 0}
```

**Example 2:**

```
Input:  var p Pool; p.Start(); p.Fail(); p.Snapshot()
Output: Snapshot{Started: 1, Failed: 1}
```

**Example 3:**

```
Input:  var p Pool; p.Snapshot()
Output: Snapshot{0, 0}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.RWMutex** | Writers take `Lock`; the stats endpoint takes `RLock`. |
| 2 | **Consistent snapshot** | Read every related field in one lock hold, not one call each. |
| 3 | **Value semantics** | Returning a struct copy means the caller needs no lock at all. |

## Hint

`Snapshot` takes `RLock` once and builds `Snapshot{Started: p.started, Failed: p.failed}` before unlocking.

## Validate

```bash
make verify
go test -race ./...
```
