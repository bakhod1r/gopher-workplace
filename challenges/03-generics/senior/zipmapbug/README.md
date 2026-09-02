# Zip That Panics On Short Input

**Level:** senior  
**Topic:** 03-generics

## Context

A CSV reader panics on rows that have fewer fields than the header, which is exactly the malformed input it was supposed to survive.

## Task

Fix the single planted bug in [zipmapbug.go](zipmapbug.go):

1. Find and fix the single bug so pairing stops at the shorter slice.
2. Never panic, whichever side is shorter.
3. Return an empty (non-nil) map when either slice is empty.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  ZipMap([]string{"a","b"}, []int{1})
Output: {a:1}
```

**Example 2:**

```
Input:  ZipMap([]string{"a"}, []int{1,2})
Output: {a:1}
```

**Example 3:**

```
Input:  ZipMap([]string{"a"}, []int{})
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two lengths, one loop** | Ranging one slice while indexing another assumes they match. |
| 2 | **Bound by the minimum** | Computing the smaller length once removes every bounds question. |
| 3 | **Panics are the loud failure** | The reverse case — a longer values slice — silently drops data instead. |

## Hint

The loop bound comes from `keys`. What guarantees `vals` is that long?

## Validate

```bash
make verify
```
