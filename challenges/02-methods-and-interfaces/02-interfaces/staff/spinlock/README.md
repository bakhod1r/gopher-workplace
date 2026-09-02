# Spin Lock

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A critical section is a handful of instructions. Parking a goroutine on a mutex costs more than the work itself, so the lock spins — with a bounded backoff.

## Task

Implement the stub(s) in [spinlock.go](spinlock.go):

1. Implement `Lock`, `TryLock`, and `Unlock` on `*SpinLock` using CAS on an `atomic.Bool`.
2. `Lock` must yield to the scheduler after a bounded number of spins so it cannot starve the holder.
3. Constraint: `-race` clean, mutual exclusion holds under contention, and `Unlock` of an unlocked lock panics rather than corrupting state.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Lock then TryLock
Output: TryLock is false
```

**Example 2:**

```
Input:  Unlock then TryLock
Output: true
```

**Example 3:**

```
Input:  1000 goroutines incrementing under the lock
Output: exactly 1000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Spinning versus parking** | Spinning wins only when the critical section is shorter than a context switch. |
| 2 | **runtime.Gosched** | Yielding prevents a spinner from starving the holder on a busy P. |
| 3 | **Misuse detection** | Reused: fail loudly on an unlock without a matching lock. |

## Hint

`for !l.held.CompareAndSwap(false, true) { spins++; if spins > N { runtime.Gosched(); spins = 0 } }`.

## Validate

```bash
make verify
```
