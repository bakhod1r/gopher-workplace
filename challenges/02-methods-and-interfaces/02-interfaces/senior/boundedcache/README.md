# Bounded Cache

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A cache in front of a slow source grew without bound and took the process down. It now has a hard entry ceiling.

## Task

Implement the stub(s) in [boundedcache.go](boundedcache.go):

1. Implement `Get` on `*Cache`, which caches lookups but never holds more than `Max` entries.
2. When full, evict the oldest inserted key before adding a new one (insertion-order eviction).
3. Implement `Len`. Constraint: `Len()` must never exceed `Max`, whatever the access pattern.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Max 2; Get("a"), Get("b"), Get("c")
Output: cache holds b and c
```

**Example 2:**

```
Input:  repeated Get of a cached key
Output: the source is not consulted again
```

**Example 3:**

```
Input:  Len() after 1000 distinct keys with Max 10
Output: 10
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded memory** | The ceiling is the design, not an optimisation. |
| 2 | **Eviction order** | A slice of keys records insertion order alongside the map. |
| 3 | **Read-through caching** | Reused: the cache implements the interface it wraps. |

## Hint

Keep `map[string]string` for lookup and `[]string` for insertion order; evict `order[0]` when full.

## Validate

```bash
make verify
```
