# Counts That Never Grow

**Level:** senior  
**Topic:** 03-generics

## Context

A duplicate-detection report says every value occurs exactly once, so nothing is ever flagged.

## Task

Fix the single planted bug in [bagzerobug.go](bagzerobug.go):

1. Find and fix the single bug so repeated values accumulate.
2. The map must hold one entry per distinct value.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Count([]string{"a","a","b"})
Output: map[a:2 b:1]
```

**Example 2:**

```
Input:  Count([]int{1,1,1})
Output: map[1:3]
```

**Example 3:**

```
Input:  Count([]int{})
Output: map[]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero values are ambiguous** | `m[k]` yields a zero for a missing key, so presence needs the comma-ok form. |
| 2 | **Assignment versus increment** | `m[k] = 1` throws away whatever was already there. |

## Hint

Look at what happens on the second occurrence.

## Validate

```bash
make verify
```
