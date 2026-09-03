# Evict The One Nobody Has Touched

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A cache with a FIFO eviction rule has a hit rate of eleven percent. The working set fits comfortably in the limit, and the entries being evicted are the ones every request needs.

## Task

Implement [lrucache.go](lrucache.go):

1. Return the value for `key` and whether it was present.
2. A hit must make the entry the most recently used, so it is the last to be evicted.
3. Hold the lock for the lookup and the reordering together.

Replace the stub body in [lrucache.go](lrucache.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  c.Put("a",1); c.Get("a")
Output: 1, true
```

**Example 2:**

```
Input:  Put a, Put b, Get a, Put c with limit 2
Output: b is evicted, a survives
```

_Explanation:_ The Get rescued a.

**Example 3:**

```
Input:  no Gets, limit 2, Put a b c
Output: a is evicted
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recency needs a Get-side update** | Without it, LRU degenerates into FIFO. |
| 2 | **list.MoveToFront** | O(1) reordering in a doubly linked list. |
| 3 | **Map to list element** | The map gives O(1) lookup; the list gives O(1) ordering. |
| 4 | **Reads mutate the structure** | So `Get` needs the write lock, not a read lock. |

## Hint

A hit does two things, not one.

## Validate

```bash
make verify
```
