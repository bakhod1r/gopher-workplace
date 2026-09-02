# Epoch Reclamation

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A lock-free structure cannot free a node while another goroutine may still be reading it. Epoch-based reclamation defers the free until every reader has moved on.

## Task

Implement the stub(s) in [epochreclaim.go](epochreclaim.go):

1. Implement `Enter`, `Exit`, and `Retire` on `*Epoch`.
2. Implement `Reclaim`, which frees retired objects only once no reader is inside the epoch they were retired in.
3. Constraint: `-race` clean, an object is never freed while a reader that entered before its retirement is still active, and nothing leaks once all readers exit.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Retire while a reader is inside
Output: not reclaimed yet
```

**Example 2:**

```
Input:  the reader exits, then Reclaim
Output: the object is freed
```

**Example 3:**

```
Input:  Retire with no active readers
Output: reclaimable immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deferred reclamation** | Safe memory reclamation for lock-free structures. |
| 2 | **Grace periods** | An object is safe once every reader that could have seen it has left. |
| 3 | **Reader counting** | Reused: a mutex-protected invariant spanning several fields. |

## Hint

Track active readers per epoch. An object retired in epoch E is safe once no reader from E or earlier is active.

## Validate

```bash
make verify
```
