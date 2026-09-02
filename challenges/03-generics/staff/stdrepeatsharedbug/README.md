# Rows That Are All The Same Row

**Level:** staff  
**Topic:** 03-generics

## Context

A board allocator hands out an n-row grid. Writing to one cell lights up the same cell on every row, and the caller's template comes back modified too.

## Task

Fix the single planted bug in [stdrepeatsharedbug.go](stdrepeatsharedbug.go):

1. Find and fix the single bug so each row owns its storage.
2. The rows must still equal `proto` element-wise, and `n <= 0` must yield an empty result.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Blank([0,0], 3)
Output: three rows of [0 0]
```

**Example 2:**

```
Input:  b := Blank([0,0], 3); b[0][0] = 7
Output: b[1][0] is still 0
```

**Example 3:**

```
Input:  Blank([0], 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Shallow versus deep** | Cloning a container copies the references it holds, not what they point at. |
| 2 | **Backing-array aliasing** | A slice value is a window onto storage someone else may also hold. |

## Hint

`slices.Repeat` copies elements. What is an element of a `[][]T`?

## Validate

```bash
make verify
```
