# Sliding Window Runs Off The End

**Level:** senior  
**Topic:** 03-generics

## Context

A three-sample smoother produces two extra readings at the end of every series, and they are always too high.

## Task

Fix the single planted bug in [windowbug.go](windowbug.go):

1. Find and fix the single bug so only full windows are emitted.
2. The result must hold exactly `len(s)-n+1` windows.
3. Each window must remain independent of the input.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Windows([]int{1,2,3}, 2)
Output: [][]int{{1,2},{2,3}}
```

**Example 2:**

```
Input:  Windows([]int{1,2,3}, 3)
Output: [][]int{{1,2,3}}
```

**Example 3:**

```
Input:  Windows([]int{1,2}, 3)
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Loop bound defines the shape** | `i+n <= len(s)` is what stops before a partial window. |
| 2 | **Clamping hides the problem** | Trimming the tail window produces short windows instead of no window. |
| 3 | **Count the results** | `len(s)-n+1` is the quickest check that the bound is right. |

## Hint

How many windows of size 2 fit in a slice of 3?

## Validate

```bash
make verify
```
