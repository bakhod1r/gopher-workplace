# Atomic Versus Mutex

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A counter guarded by a mutex became the contention hotspot. Both variants must be correct; the benchmarks decide which ships.

## Task

Implement the stub(s) in [atomicvsmutex.go](atomicvsmutex.go):

1. Implement `Inc`, `Add`, and `Value` on `*MutexCounter` and `*AtomicCounter`.
2. Implement `IncAll`, which increments a set of counters through the interface.
3. Constraint: `-race` clean, exact totals under heavy concurrency, and `AtomicCounter.Inc` must allocate zero times.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  1000 goroutines each Inc once
Output: Value 1000
```

**Example 2:**

```
Input:  Add(5) three times
Output: 15
```

**Example 3:**

```
Input:  IncAll over two counters
Output: both incremented
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Atomics versus locks** | An atomic add is one instruction; a mutex is a full critical section. |
| 2 | **Interface parity** | Both implementations satisfy one contract, so the choice is a benchmark, not a rewrite. |
| 3 | **Exactness under contention** | Reused: a lost update is a correctness bug, not a performance one. |

## Hint

`atomic.Int64.Add` is the whole implementation — no lock, no read-modify-write gap.

## Validate

```bash
make verify
```
