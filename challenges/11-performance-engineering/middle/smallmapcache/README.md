# A Cache That Cannot Grow

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

An unbounded memo table is a memory leak that looks like an optimisation: it is fast right up to the point where it holds every key the service has ever seen. Adding a capacity and an eviction rule is what turns it into a cache. First-in-first-out is the simplest rule that works, and it is what you want when the value distribution is flat.

## Task

Implement the four methods in [smallmapcache.go](smallmapcache.go):

1. `Get` returns the value and whether it was present, counting hits and misses, and must not allocate on a hit.
2. `Put` stores a value and evicts the oldest *inserted* entry when full; a hit never changes eviction order.
3. Overwriting an existing key updates in place without changing the order or the size; a non-positive `Cap` stores nothing.

## Examples

**Example 1:**

```
Input:  Cap 2; Put(a), Put(b), Get(a), Put(c)
Output: a evicted, b and c held
```

**Example 2:**

```
Input:  Cap 2; Put(a), Put(b), Put(a, 99)
Output: Len is 2 and a is still the oldest
```

**Example 3:**

```
Input:  Cap 0; Put(a)
Output: Len 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A cache without a bound is a leak** | The capacity is the feature, not the limitation. |
| 2 | **FIFO is not LRU** | Reads do not refresh an entry; know which one you have promised. |
| 3 | **Two structures, one invariant** | The map and the order slice must be updated together or the eviction goes wrong. |

## Topics used again

Maps plus a slice, `delete`, methods on pointer receivers, invariants.

## Hint

Only a genuine insert appends to `order`; an overwrite touches the map alone.

## Validate

```bash
make verify
```
