# Scheduler Leader Lease

**Level:** junior
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A scheduler runs on several replicas, but only one may dispatch jobs at a time. Leadership is a monotonically increasing term number: a replica takes over by moving the term from `n-1` to `n`. Every replica races for the same term at once.

## Task

Implement the stubbed functions in [leaderlease.go](leaderlease.go) so that:

1. `Claim` moves the term from `term-1` to `term` and reports whether this caller won.
2. A term that is already taken, or one that skips ahead, must lose.
3. `Term` returns the current term.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var l Lease; l.Claim(1)
Output: true
```

**Example 2:**

```
Input:  l.Claim(1) again
Output: false
```

**Example 3:**

```
Input:  var l Lease; l.Claim(5)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `CompareAndSwap` | Swap only if the current value still equals the expected one — the whole check-and-set is one atomic step. |
| 2 | Check-then-act is a race | `if Load()==x { Store(y) }` lets two goroutines both pass the check. |
| 3 | Monotonic state | Terms only move forward, so a stale or skipping claim must fail rather than overwrite. |

## Hint

One `CompareAndSwap(term-1, term)` call does the whole job — no mutex, no `if`.

## Validate

```bash
make verify
```
