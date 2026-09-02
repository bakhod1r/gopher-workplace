# Insert At The End Rejected

**Level:** senior  
**Topic:** 03-generics

## Context

A playlist silently ignores tracks dropped after the last item, and only that position.

## Task

Fix the single planted bug in [insertboundbug.go](insertboundbug.go):

1. Find and fix the single bug so inserting at `len(s)` appends.
2. Genuinely out-of-range indexes must still return `s` unchanged.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  InsertAt([]int{1}, 1, 2)
Output: []int{1,2}
```

**Example 2:**

```
Input:  InsertAt([]int{1,3}, 1, 2)
Output: []int{1,2,3}
```

**Example 3:**

```
Input:  InsertAt([]int{1}, 5, 2)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Insertion positions** | There are `len(s)+1` gaps in a slice of `len(s)` elements. |
| 2 | **Off-by-one in a guard** | The upper bound for insertion is `len(s)`, not `len(s)-1`. |
| 3 | **Silent rejection** | Returning the input unchanged makes the bug easy to miss. |

## Hint

How many places can a value be inserted into a slice of length 3?

## Validate

```bash
make verify
```
