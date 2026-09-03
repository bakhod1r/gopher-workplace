# Search That Skips The First Match

**Level:** senior  
**Topic:** 03-generics

## Context

A lookup over a sorted index reports "not found" for keys that are definitely present, but only when the key repeats.

## Task

Fix the single planted bug in [binsearchbug.go](binsearchbug.go):

1. Find and fix the single bug so the first matching index is returned.
2. A missing key must still return the insertion point and `false`.
3. The input is assumed sorted by key.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SearchBy([{1},{1},{2}], idOf, 1)
Output: 0, true
```

**Example 2:**

```
Input:  SearchBy([{1},{3}], idOf, 2)
Output: 1, false
```

**Example 3:**

```
Input:  SearchBy(nil, idOf, 1)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lower bound versus upper bound** | `<` converges on the first match; `<=` converges past the last one. |
| 2 | **The final equality check** | It only succeeds when the loop lands on a matching element. |
| 3 | **Duplicates expose it** | With unique keys the two variants agree. |

## Hint

What does the loop converge on when several elements share the target key?

## Validate

```bash
make verify
```
