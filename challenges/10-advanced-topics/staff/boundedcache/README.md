# A Cache That Cannot Outgrow Its Limit

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A response cache is added with a size limit and a mutex. Memory still climbs: the limit counts entries in the map, and the eviction path was never reached because every key was new.

## Task

Implement [boundedcache.go](boundedcache.go):

1. Store a copy of `val` under `key` — the caller reuses its buffer.
2. Evict the oldest entry when inserting a new key would exceed `limit`.
3. Overwriting an existing key must not evict anything.
4. Safe for concurrent use; copy outside the lock.

Replace the stub body in [boundedcache.go](boundedcache.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  c.Put("k", buf); copy(buf, "SECON")
Output: Get("k") still returns the original
```

**Example 2:**

```
Input:  NewCache(2), put a, b, c
Output: a is evicted, Len is 2
```

**Example 3:**

```
Input:  overwriting an existing key
Output: no eviction
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ownership at the boundary** | A cache that stores a caller's slice stores a promise the caller will break. |
| 2 | **Bounded state** | The eviction path must run on the insert of a new key, not on every put. |
| 3 | **Lock scope** | Allocating and copying outside the lock keeps the critical section short. |
| 4 | **FIFO order tracking** | The order slice is what makes "oldest" meaningful. |

## Hint

Two decisions per put: does this key already exist, and is the cache full?

## Validate

```bash
make verify
```
