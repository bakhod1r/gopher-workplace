# The Override That Never Applies

**Level:** senior  
**Topic:** 03-generics

## Context

Configuration is layered defaults, then file, then environment. Environment variables are being ignored.

## Task

Fix the single planted bug in [chainmergebug.go](chainmergebug.go):

1. Find and fix the single bug so later maps override earlier ones.
2. Keys present in only one map must survive.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Merge({a:1}, {a:2})
Output: map[a:2]
```

**Example 2:**

```
Input:  Merge({a:1}, {b:2})
Output: map[a:1 b:2]
```

**Example 3:**

```
Input:  Merge()
Output: map[]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Last write wins** | Plain assignment already implements the override rule. |
| 2 | **Presence checks invert precedence** | Skipping existing keys makes the *first* map win. |

## Hint

Read the condition guarding the assignment.

## Validate

```bash
make verify
```
