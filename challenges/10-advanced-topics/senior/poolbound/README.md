# The Pool That Kept Every Oversized Buffer

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A service pools its render buffers and runs flat for weeks. One customer uploads a 40 MB document, and from then on the process holds 40 MB per pooled buffer forever.

## Task

Fix the single planted bug in [poolbound.go](poolbound.go):

1. Fill `size` bytes into a pooled buffer and return the count.
2. Return the buffer to the pool only when its capacity is at most `maxScratch`.
3. Fix the single bug; small buffers must still be recycled.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Render(16)
Output: 16
```

**Example 2:**

```
Input:  Render(1<<20) then PooledCap()
Output: at most 4096
```

_Explanation:_ The huge buffer was dropped.

**Example 3:**

```
Input:  Render(32) then PooledCap()
Output: non-zero
```

_Explanation:_ Normal buffers still go back.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pools have no size policy** | `sync.Pool` recycles whatever you hand it, however large. |
| 2 | **Capacity is what persists** | The length resets; the allocation does not. |
| 3 | **Dropping is free** | An unreturned buffer is simply collected. |

## Hint

Everything about the borrow is right. What condition should guard the return?

## Validate

```bash
make verify
```
