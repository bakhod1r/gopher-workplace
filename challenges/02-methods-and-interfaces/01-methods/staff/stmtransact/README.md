# Optimistic Transaction

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Software transactional memory is optimistic: read a value, compute a new one
without holding any lock, and commit only if nobody else changed it meanwhile.
A version counter is what makes "nobody else changed it" checkable. If the
commit is rejected, the transaction re-runs from a fresh read.

`Read` and `Commit` are already written — the retry loop is yours.

## Task

Implement `Tx` on `*TVar` in [stmtransact.go](stmtransact.go):

1. `Read()` the current value and version.
2. Compute `fn(value)`.
3. `Commit(version, newValue)`; on success return the committed value.
4. On failure, loop — reading again, recomputing from the **new** value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TVar(5); Tx(v => v*2)
Output: 10, and Read() sees 10
```

**Example 2:**

```
Input:  Tx(v => v+1) on version 0
Output: version becomes 1
```

**Example 3:**

```
Input:  100 goroutines each running Tx(v => v+1) on TVar(0)
Output: final value 100 — no lost updates
```

_Explanation:_ every conflicting transaction retries rather than overwriting.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Optimistic concurrency** | No lock is held across `fn`, so slow computations do not block other transactions. |
| 2 | **Recompute, do not reuse** | The retry must call `fn` again on the fresh value; reusing the stale result is a lost update. |
| 3 | **Version as a conflict detector** | Comparing values would miss an A→B→A change; comparing versions cannot. |

## Hint

`for { v, ver := tv.Read(); next := fn(v); if tv.Commit(ver, next) { return next } }`.
The critical detail is that `fn` is inside the loop.

## Validate

```bash
make verify
```
