# Empty The Map, Keep The Map

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A worker reuses one map between batches. Assigning a fresh map inside the helper leaves every other holder of the old map looking at stale data.

## Task

Implement [clearmap.go](clearmap.go):

1. Delete every entry of `m` in place.
2. The caller's map value must stay the same map — do not assign a new one.
3. A nil map must not panic.

Replace the stub body in [clearmap.go](clearmap.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Reset(map[string]int{"a":1})
Output: len == 0
```

**Example 2:**

```
Input:  alias := m; Reset(m)
Output: len(alias) == 0
```

_Explanation:_ Both names still refer to one map.

**Example 3:**

```
Input:  Reset(nil)
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **clear** | `clear(m)` removes all entries from the map you already have. |
| 2 | **Maps are reference-like** | A map value is a pointer to a runtime structure; the parameter is a copy of that pointer. |
| 3 | **Reuse over reallocation** | Emptying keeps the buckets already paid for. |

## Hint

Assigning `m = map[string]int{}` only rebinds the local parameter.

## Validate

```bash
make verify
```
