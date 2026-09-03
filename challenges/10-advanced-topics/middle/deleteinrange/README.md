# Delete While You Range

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A cleanup pass collects the keys to remove into a slice first, "because deleting during iteration is unsafe". The slice is as large as the map, and the caution was unnecessary.

## Task

Implement [deleteinrange.go](deleteinrange.go):

1. Delete every entry with an even key and return the count.
2. Negative even keys count; 0 is even.
3. Modify the caller's map in place; a nil map removes nothing.

Replace the stub body in [deleteinrange.go](deleteinrange.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  RemoveEven(map[int]int{1:1, 2:2})
Output: 1
```

**Example 2:**

```
Input:  RemoveEven(map[int]int{-2:1, -1:1, 0:1})
Output: 2
```

_Explanation:_ -2 and 0 are even.

**Example 3:**

```
Input:  RemoveEven(nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deleting during range is defined** | The spec allows it; an unreached deleted entry is simply not produced. |
| 2 | **Maps are reference-like** | The deletions reach the caller. |
| 3 | **No second slice needed** | Collecting keys first doubles the memory for nothing. |

## Hint

Range and delete. That is the whole function.

## Validate

```bash
make verify
```
