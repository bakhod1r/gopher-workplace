# LRU Cache

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Insertion-order eviction throws away hot keys. The cache now evicts the least *recently used* entry instead.

## Task

Implement the stub(s) in [lrucache.go](lrucache.go):

1. Implement `Get` and `Put` on `*LRU` with an entry ceiling of `Cap`.
2. A `Get` hit and a `Put` of an existing key both mark the entry as most recently used.
3. Constraint: `Get` and `Put` must be O(1) — no scan of all entries per operation.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cap 2; Put(a), Put(b), Get(a), Put(c)
Output: b evicted, a and c remain
```

**Example 2:**

```
Input:  Get of a missing key
Output: "", false
```

**Example 3:**

```
Input:  Len() with Cap 2 after 3 puts
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **LRU eviction** | Recency, not insertion order, decides what survives. |
| 2 | **O(1) with a map + linked list** | A doubly linked list gives constant-time move-to-front and eviction. |
| 3 | **Pointer bookkeeping** | Reused: careful node relinking; a dropped pointer leaks or corrupts. |

## Hint

Keep a `map[string]*node` plus head/tail sentinels. Touch on hit, evict at the tail.

## Validate

```bash
make verify
```
