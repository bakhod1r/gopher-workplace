# A Map That Gives Its Buckets Back

**Level:** middle
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A session table peaks at ten million entries overnight and settles at a few thousand by morning. Resident memory never comes back down.

## Task

Implement [compactmap.go](compactmap.go):

1. Return a new map with the same entries as `m`.
2. Size the new map to the surviving entry count.
3. A nil input returns an empty, usable map.

Replace the stub body in [compactmap.go](compactmap.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Compact(map[string]int{"a":1})
Output: map[a:1]
```

_Explanation:_ Same entries, different map.

**Example 2:**

```
Input:  got := Compact(m); got["b"]=2
Output: m does not gain "b"
```

**Example 3:**

```
Input:  Compact(nil)
Output: map[]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Maps never shrink** | `delete` frees the entry, not the bucket array. |
| 2 | **Rebuild to release** | A fresh map sized to `len(m)` is the only way down. |
| 3 | **Size hints** | `make(map[K]V, len(m))` avoids rehashing on the way in. |

## Hint

`delete` in a loop will not help. What does help is a second map.

## Validate

```bash
make verify
```
