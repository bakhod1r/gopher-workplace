# A View That Lets Callers Write

**Level:** senior  
**Topic:** 03-generics

## Context

A handler returns the first page of a shared buffer. Callers append a footer row and the second page silently loses its first record.

## Task

Fix the single planted bug in [clipbug.go](clipbug.go):

1. Find and fix the single bug so appending to the result cannot touch `s`.
2. Clamp `n` into `[0, len(s)]` as the current code already does.
3. The returned elements must still be the first `n`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Head([]int{1,2,3}, 2)
Output: []int{1,2}
```

**Example 2:**

```
Input:  append to the result
Output: s unchanged
```

**Example 3:**

```
Input:  Head([]int{1}, 9)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Spare capacity is shared storage** | `s[:n]` keeps `cap(s)`, so `append` writes into `s` itself. |
| 2 | **`slices.Clip`** | Caps the capacity at the length, forcing the next append to allocate. |
| 3 | **Reads are fine, writes are not** | The bug is invisible until someone appends. |

## Hint

The values are right; the capacity is not.

## Validate

```bash
make verify
```
