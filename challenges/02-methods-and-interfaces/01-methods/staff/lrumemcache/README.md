# LRU Cache

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An LRU cache needs O(1) lookup *and* O(1) recency updates. The standard answer
is a map for lookup plus a doubly linked list for order, with sentinel head and
tail nodes so no insert or remove ever needs a nil check. The list helpers
`remove` and `insert` are already written — the policy is yours.

## Task

Implement `Get` and `Put` on `*LRU` in [lrumemcache.go](lrumemcache.go):

1. `Get(key)`: on a miss return `(0, false)`. On a hit, move the node to the
   front (`remove` then `insert`) and return `(val, true)`.
2. `Put(key, val)`: if the key exists, update its value and move it to the
   front. Otherwise create a node, insert it at the front, add it to the map,
   and — if the map now exceeds `l.cap` — evict `l.tail.prev`, deleting it from
   the map too.

**Constraint (staff):** memory is bounded — after 100,000 distinct keys, both the map and the list must hold exactly `cap` entries.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  New(2); Put("a",1); Put("b",2); Get("a")
Output: (1, true) — and "a" is now the most recent
```

**Example 2:**

```
Input:  then Put("c",3)
Output: "b" is evicted (it was least recently used)
```

**Example 3:**

```
Input:  Get("b")
Output: (0, false)
```

_Explanation:_ the `Get("a")` in step 1 is what saved "a" and doomed "b".

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map + linked list** | The map gives O(1) find; the list gives O(1) reorder. Neither alone is enough. |
| 2 | **Sentinel nodes** | `head` and `tail` are never real entries, which is why `remove`/`insert` need no nil handling. |
| 3 | **Evict from the tail** | `l.tail.prev` is the least recently used real node. |

## Hint

"Move to front" is exactly `l.remove(n)` followed by `l.insert(n)` — reuse the
helpers rather than repointing links yourself. And evict *after* inserting, so
a `Put` that overwrites an existing key never triggers an eviction.

## Validate

```bash
make verify
```
