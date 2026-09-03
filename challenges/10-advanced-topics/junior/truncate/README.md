# Cut The Tail Without Pinning It

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A queue trims its backlog with `q = q[:n]`. The dropped jobs are gone from the caller's view but the process still holds their memory for as long as the queue lives.

## Task

Implement [truncate.go](truncate.go):

1. Clamp `n` into `[0, len(s)]`.
2. Clear the elements from `n` onward before returning.
3. Return `s[:n]` — the storage is reused, not reallocated.

Replace the stub body in [truncate.go](truncate.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Truncate([]*Node{{1},{2},{3}}, 1)
Output: length 1
```

_Explanation:_ Indices 1 and 2 of the array become nil.

**Example 2:**

```
Input:  Truncate(s, 9)
Output: length 3
```

_Explanation:_ n is clamped up to len(s).

**Example 3:**

```
Input:  Truncate(s, -1)
Output: length 0
```

_Explanation:_ n is clamped to 0.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reslicing does not erase** | Elements past the new length stay in the backing array. |
| 2 | **clear on a sub-slice** | `clear(s[n:])` releases exactly the dropped range. |
| 3 | **Clamping** | Reject out-of-range n before it indexes anything. |

## Hint

`s[:n]` is the answer; something has to happen to `s[n:]` first.

## Validate

```bash
make verify
```
