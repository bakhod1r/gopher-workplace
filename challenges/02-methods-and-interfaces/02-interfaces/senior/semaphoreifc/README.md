# Semaphore

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A shared dependency tolerates a limited number of in-flight requests. A counting semaphore enforces the ceiling.

## Task

Implement the stub(s) in [semaphoreifc.go](semaphoreifc.go):

1. Implement `Acquire`, `TryAcquire`, and `Release` on `*Semaphore`.
2. Implement `RunLimited`, which runs every job concurrently but never lets more than `limit` run at once.
3. Constraint: race-free under `-race`, peak concurrency never exceeds `limit`, and `RunLimited` must not deadlock.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  limit 2, 100 jobs
Output: peak concurrency 2
```

**Example 2:**

```
Input:  TryAcquire on a full semaphore
Output: false
```

**Example 3:**

```
Input:  Release then TryAcquire
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counting semaphore** | A buffered channel is the idiomatic Go implementation. |
| 2 | **Non-blocking acquire** | `select` with `default` gives `TryAcquire`. |
| 3 | **Bounded concurrency** | Reused: the ceiling is enforced, not hoped for. |

## Hint

`ch := make(chan struct{}, n)` — sending acquires, receiving releases.

## Validate

```bash
make verify
```
