# A Buffer That Allocates Once, Ever

**Level:** middle
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A crash reporter keeps the last 256 log lines in a slice it appends to and reslices. Under a log storm the slice grows without bound before the reslice ever runs.

## Task

Implement [ringbuf.go](ringbuf.go):

1. Store `v` in the ring's existing buffer.
2. When the ring is full, overwrite the oldest element and advance the head.
3. `Push` must never allocate.

Replace the stub body in [ringbuf.go](ringbuf.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewRing(3), push 1,2,3
Output: Items() is [1 2 3]
```

**Example 2:**

```
Input:  then push 4
Output: Items() is [2 3 4]
```

_Explanation:_ The oldest element is overwritten.

**Example 3:**

```
Input:  NewRing(1), push 1 then 2
Output: Items() is [2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Modular indexing** | `(head+n) % len(buf)` wraps the write position without branching on the end. |
| 2 | **Fixed capacity** | Bounded memory is a design choice enforced by never calling `append`. |
| 3 | **Pointer receivers** | `*Ring` is what lets `Push` mutate the ring the caller holds. |

## Hint

Two cases: the ring has room, or it does not. Only the second one moves the head.

## Validate

```bash
make verify
```
