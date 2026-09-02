# Lock Escalation

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The common case is a read. Taking the write lock every time would serialize
callers that only wanted to look. So: check under the read lock, and escalate to
the write lock only when a change is actually needed — then re-check, because
the value may have changed while no lock was held.

## Task

Implement `IncrementIfZero` on `*OptLock` in [rwlockopt.go](rwlockopt.go):

1. `RLock`, read `o.v`. If it is non-zero, `RUnlock` and return it.
2. `RUnlock`, then `Lock`.
3. Re-check: if `o.v` is still 0, increment it.
4. `Unlock` and return the value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  fresh OptLock; IncrementIfZero()
Output: 1
```

**Example 2:**

```
Input:  IncrementIfZero() again
Output: 1  (fast path: non-zero, no write lock taken)
```

**Example 3:**

```
Input:  two goroutines racing on a zero value
Output: 1 for both — the second finds it non-zero after escalating
```

_Explanation:_ the re-check under the write lock is what prevents a double increment.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`RWMutex` is not upgradable** | You must fully release the read lock before taking the write lock; holding both deadlocks. |
| 2 | **Double-checked locking** | The gap between the two locks is real, so the condition must be re-tested. |
| 3 | **Reading the result under a lock** | Capture the value before unlocking, not after. |

## Hint

Do not `defer` here — the read lock has to be released explicitly before the
write lock is taken. Calling `Lock` while holding `RLock` on the same `RWMutex`
deadlocks immediately.

## Validate

```bash
make verify
```
