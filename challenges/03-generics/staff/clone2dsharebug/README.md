# The Clone That Shares Its Rows

**Level:** staff  
**Topic:** 03-generics

## Context

A board-game engine clones the grid before exploring a move. Every speculative move is being written back into the real board, so the search corrupts the position it was searching from.

## Task

Fix the single planted bug in [clone2dsharebug.go](clone2dsharebug.go):

1. Find and fix the single bug so writing to the clone cannot affect the source.
2. The clone's shape and contents must match the source exactly.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Clone2D([[1],[2]])
Output: [[1] [2]]
```

**Example 2:**

```
Input:  c[0][0] = 99; m[0][0]
Output: unchanged
```

**Example 3:**

```
Input:  Clone2D(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate the source. |
| 2 | **One level is not deep** | Copying the outer slice duplicates row *headers*, which still point at the original arrays. |
| 3 | **Copy per row** | Each inner slice needs its own `make` plus `copy`. |

## Hint

What exactly does `out[i] = row` duplicate?

## Validate

```bash
make verify
```
