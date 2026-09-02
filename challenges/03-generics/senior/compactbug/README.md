# Deduplication That Misses

**Level:** senior  
**Topic:** 03-generics

## Context

A tag list still shows duplicates, but only when the same tag appears twice with something else in between.

## Task

Fix the single planted bug in [compactbug.go](compactbug.go):

1. Find and fix the single bug so every value appears once, in ascending order.
2. Leave the input untouched.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Distinct([]int{3,1,3})
Output: []int{1,3}
```

**Example 2:**

```
Input:  Distinct([]int{1,1,2})
Output: []int{1,2}
```

**Example 3:**

```
Input:  Distinct([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Compact` is local** | It only collapses *adjacent* equal elements. |
| 2 | **Sorting makes it global** | Sorting brings equal values together, which is what turns Compact into a deduplicator. |
| 3 | **Adjacent-only bugs hide** | Inputs that happen to be sorted pass the test. |

## Hint

What does `Compact` guarantee, exactly?

## Validate

```bash
make verify
```
