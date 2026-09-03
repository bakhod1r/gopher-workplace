# Prefix Skip Turned Into A Filter

**Level:** senior  
**Topic:** 03-generics

## Context

A CSV reader that should skip leading blank lines is also dropping blank lines in the middle of the file, corrupting the record count.

## Task

Fix the single planted bug in [dropwhilebug.go](dropwhilebug.go):

1. Find and fix the single bug so only the leading run is dropped.
2. Everything from the first non-matching element onward must be kept, matches included.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  DropWhile([]int{2,4,5,6}, isEven)
Output: []int{5,6}
```

**Example 2:**

```
Input:  DropWhile([]int{1,2}, isEven)
Output: []int{1,2}
```

**Example 3:**

```
Input:  DropWhile([]int{2,4}, isEven)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Prefix versus subset** | `DropWhile` cuts at the first failure; a filter examines every element. |
| 2 | **Find the split, then copy** | One index scan is enough — no per-element decision afterwards. |
| 3 | **No aliasing** | Return a fresh slice rather than `s[i:]`. |

## Hint

The trailing `6` in the first example is even. Why does it survive?

## Validate

```bash
make verify
```
