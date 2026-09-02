# Binary Search Against The Grain

**Level:** staff  
**Topic:** 03-generics

## Context

A leaderboard is stored best-first so the top rows can be rendered without a reverse. Score lookups against it report "no such score" for almost every score that is really there — and the few hits look like luck.

## Task

Fix the single planted bug in [stdbinsearchcmpbug.go](stdbinsearchcmpbug.go):

1. Find and fix the single bug so the search agrees with the board's descending order.
2. A score that is not present must still return -1 and false.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  FindScore([{a 9} {b 5} {c 1}], 5)
Output: 1, true
```

**Example 2:**

```
Input:  FindScore([{a 9} {b 5} {c 1}], 9)
Output: 0, true
```

**Example 3:**

```
Input:  FindScore([{a 9} {b 5} {c 1}], 7)
Output: -1, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Preconditions of stdlib helpers** | `slices` and `maps` helpers are fast because they trust you; violating a precondition is silent, not fatal. |
| 2 | **The comparator defines the order** | `BinarySearchFunc` requires `cmp` to be non-decreasing over the slice; on a descending slice that means the arguments swap. |

## Hint

The board is descending. Which way round does the comparison have to read for the result to rise with the index?

## Validate

```bash
make verify
```
