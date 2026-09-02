# Reference Counting

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A shared connection is closed while another caller is still using it. Reference counting decides when the resource is really free.

## Task

Implement the stub(s) in [refcount.go](refcount.go):

1. Implement `Acquire`, `Release`, and `Count` on `*RefCounted`.
2. `Release` closes the underlying resource exactly once, when the count reaches zero; releasing below zero must be rejected.
3. Constraint: race-free under `-race`; concurrent acquire/release must never close early or twice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Acquire, Acquire, Release
Output: not closed yet, count 1
```

**Example 2:**

```
Input:  the final Release
Output: the resource is closed once
```

**Example 3:**

```
Input:  Release when the count is 0
Output: false, and no second close
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reference counting** | Lifetime is a shared decision, not a single owner's. |
| 2 | **Mutex-protected invariant** | Count and close must move together, atomically. |
| 3 | **Exactly-once cleanup** | Reused: a resource must never be closed twice. |

## Hint

Hold the mutex across both the decrement and the close decision — two atomics are not enough.

## Validate

```bash
make verify
```
