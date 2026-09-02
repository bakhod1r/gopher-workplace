# Clip Applied To The Wrong Slice

**Level:** staff  
**Topic:** 03-generics

## Context

A pagination helper hands the first page of a result set to a renderer that appends a footer row. The second page's first record is replaced by that footer row.

## Task

Fix the single planted bug in [clipheaderbug.go](clipheaderbug.go):

1. Find and fix the single bug so the returned slice's capacity equals its length.
2. The clamping behaviour must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  cap(Shrink([]int{1,2,3,4}, 2))
Output: 2
```

**Example 2:**

```
Input:  append to the page, then read s[2]
Output: unchanged
```

**Example 3:**

```
Input:  Shrink([]int{1,2}, 9)
Output: [1 2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Length versus capacity** | A slice header carries spare capacity that `append` will happily write into. |
| 2 | **Three-index slicing** | `s[a:b:b]` caps the result so `append` must allocate instead of overwriting. |
| 3 | **Order of operations on headers** | `Clip(s)[:n]` re-widens the capacity that `Clip` just removed. |

## Hint

`slices.Clip` is called. On which slice, and what happens to its result afterwards?

## Validate

```bash
make verify
```
