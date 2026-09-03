# Running Total Off By One Step

**Level:** senior  
**Topic:** 03-generics

## Context

A balance chart is drawn one transaction behind: the last payment never appears, and the first bar is always the opening balance.

## Task

Fix the single planted bug in [scanbug.go](scanbug.go):

1. Find and fix the single bug so each output is the accumulator *after* the matching element.
2. The result must still have exactly `len(s)` elements.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Scan([]int{1,2,3}, 0, add)
Output: []int{1,3,6}
```

**Example 2:**

```
Input:  Scan([]int{}, 5, add)
Output: []int{}
```

**Example 3:**

```
Input:  Scan([]int{2}, 1, mul)
Output: []int{2}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Emit after applying** | Appending before the fold shifts every value by one position. |
| 2 | **Length is a weak check** | Both versions return `len(s)` elements — only the values differ. |
| 3 | **Scan versus Reduce** | The final element of a correct `Scan` equals the `Reduce` result. |

## Hint

Compare the last element of the result with what `Reduce` would return.

## Validate

```bash
make verify
```
