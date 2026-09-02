# A Snapshot That Keeps Changing

**Level:** senior  
**Topic:** 03-generics

## Context

A metrics endpoint serves a snapshot while the collector keeps writing. Readers see values that change mid-response.

## Task

Fix the single planted bug in [deepclonebug.go](deepclonebug.go):

1. Find and fix the single bug so the snapshot shares nothing with the original.
2. Writing into a copied slice must not affect the source map.
3. Return an empty (non-nil) map for empty or nil input.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Snapshot({a:[1,2]})
Output: {a:[1,2]}
```

**Example 2:**

```
Input:  write into the copy
Output: original unchanged
```

**Example 3:**

```
Input:  Snapshot(nil)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cloning is shallow** | `maps.Clone` copies the map; the values still point at the same arrays. |
| 2 | **One level deeper** | Maps of slices need each value copied explicitly. |
| 3 | **Cost is honest** | The fix is O(total elements), which is what a real snapshot costs. |

## Hint

`maps.Clone` copies the map. What about the slices inside it?

## Validate

```bash
make verify
```
