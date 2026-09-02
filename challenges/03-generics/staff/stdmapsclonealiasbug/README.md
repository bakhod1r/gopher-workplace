# The Clone That Shares Its Slices

**Level:** staff  
**Topic:** 03-generics

## Context

A request handler snapshots the shared tag table before editing it, so the global stays clean. Edits are leaking into the global anyway, and only for keys whose tag list already existed.

## Task

Fix the single planted bug in [stdmapsclonealiasbug.go](stdmapsclonealiasbug.go):

1. Find and fix the single bug so the returned map's slice values are independent of the original.
2. A nil map must still clone to nil.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  CloneTags({a:[x]}); edit copy's [x]
Output: original still [x]
```

**Example 2:**

```
Input:  CloneTags({a:[x]})
Output: same keys and values
```

**Example 3:**

```
Input:  CloneTags(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shallow versus deep** | Cloning a container copies the references it holds, not what they point at. |
| 2 | **Backing-array aliasing** | A slice value is a window onto storage someone else may also hold. |

## Hint

What exactly does `maps.Clone` copy for a value of type `[]string`?

## Validate

```bash
make verify
```
