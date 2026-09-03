# Pruning That Walks A Stale Key List

**Level:** staff  
**Topic:** 03-generics

## Context

A cache eviction pass reports which keys it dropped, and the audit log never matches the cache. The map itself ends up right; the report is fiction.

## Task

Fix the single planted bug in [stdmapsdeletewalkbug.go](stdmapsdeletewalkbug.go):

1. Find and fix the single bug so the returned list names exactly the keys that were removed.
2. Entries at or above `limit` must survive.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Prune({a:1, b:5, c:2}, 3)
Output: [a c], map is {b:5}
```

**Example 2:**

```
Input:  Prune({a:5}, 3)
Output: [], map is {a:5}
```

**Example 3:**

```
Input:  Prune({}, 3)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bulk versus per-item** | `maps.DeleteFunc` sweeps the whole map in one call; calling it inside a per-key loop deletes everything on the first iteration. |
| 2 | **Snapshots go stale** | A key list taken before the mutation says nothing about what is still present. |

## Hint

How many entries does one call to `maps.DeleteFunc` remove?

## Validate

```bash
make verify
```
