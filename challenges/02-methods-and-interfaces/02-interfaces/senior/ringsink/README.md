# Ring Buffer Sink

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A crash-dump buffer must keep only the last N log lines, whatever the volume, with no allocation after startup.

## Task

Implement the stub(s) in [ringsink.go](ringsink.go):

1. Implement `Write` on `*RingSink`, overwriting the oldest entry once full.
2. Implement `Snapshot`, returning the entries oldest-first.
3. Constraint: after construction, `Write` must allocate zero times — the ring is fixed size.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  size 3; write a, b, c, d
Output: Snapshot is [b c d]
```

**Example 2:**

```
Input:  fewer writes than the size
Output: only what was written
```

**Example 3:**

```
Input:  1M writes into a size-4 ring
Output: still 4 entries, no growth
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ring buffer** | Fixed memory with O(1) writes; the oldest entry is evicted implicitly. |
| 2 | **Modular indexing** | `pos % len` wraps without branching. |
| 3 | **Allocation discipline** | Reused: the steady state must not allocate. |

## Hint

Track a monotonically increasing write count; the oldest entry is at `count-len` when full.

## Validate

```bash
make verify
```
