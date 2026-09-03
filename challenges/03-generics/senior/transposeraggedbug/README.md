# Transpose That Assumes A Rectangle

**Level:** senior  
**Topic:** 03-generics

## Context

A CSV pivot silently truncates every column past the width of the first row, and the widest rows are exactly the ones analysts care about.

## Task

Fix the single planted bug in [transposeraggedbug.go](transposeraggedbug.go):

1. Find and fix the single bug so the result is as wide as the longest row.
2. Short rows must simply be skipped in the columns they lack.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Transpose([[1,2],[3]])
Output: [[1 3],[2]]
```

**Example 2:**

```
Input:  Transpose([[1],[2,3]])
Output: [[1 2],[3]]
```

**Example 3:**

```
Input:  Transpose([])
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ragged input** | The first row is not the shape of the data. |
| 2 | **Bounds per row** | Each access needs its own length check. |

## Hint

Where does `width` come from?

## Validate

```bash
make verify
```
