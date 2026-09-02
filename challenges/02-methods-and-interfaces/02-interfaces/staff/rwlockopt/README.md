# Read-Write Lock

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A config map is read on every request and written once a minute. A plain mutex made every reader wait for every other reader.

## Task

Implement the stub(s) in [rwlockopt.go](rwlockopt.go):

1. Implement `Get`, `Set`, and `Len` on `*RWStore` with `sync.RWMutex`, and the same on `*MutexStore` with `sync.Mutex`.
2. Implement `Snapshot` on `*RWStore`, returning an independent copy taken under a read lock.
3. Constraint: `-race` clean, both stores behave identically, and readers must be able to overlap — the test observes concurrent readers under the read lock.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Set("a", 1); Get("a")
Output: 1, true
```

**Example 2:**

```
Input:  many concurrent readers
Output: they overlap under RLock
```

**Example 3:**

```
Input:  Snapshot then Set
Output: the snapshot is unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **RWMutex semantics** | Many readers or one writer; readers do not exclude each other. |
| 2 | **Snapshot under a read lock** | Copy while holding RLock so no writer can interleave. |
| 3 | **Interface parity** | Reused: two implementations, one contract, decided by measurement. |

## Hint

`RLock` for reads, `Lock` for writes. Never upgrade an RLock to a Lock — release first.

## Validate

```bash
make verify
```
