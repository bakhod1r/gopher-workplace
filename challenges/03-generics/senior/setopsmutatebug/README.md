# Union That Eats Its Input

**Level:** senior  
**Topic:** 03-generics

## Context

A permission check computes the union of a role set and a grant set. After the first check the role set itself has grown, and users keep gaining access they were never given.

## Task

Fix the single planted bug in [setopsmutatebug.go](setopsmutatebug.go):

1. Find and fix the single bug so neither argument is modified.
2. The result must contain every element of both sets.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Union({1,2},{2,3})
Output: {1,2,3}
```

**Example 2:**

```
Input:  a after Union
Output: unchanged
```

**Example 3:**

```
Input:  Union(empty, {1})
Output: {1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Maps are reference types** | Assigning a map copies the header, not the contents. |
| 2 | **Pure functions** | A helper documented as returning a new value must allocate one. |

## Hint

What does `out := a` actually copy?

## Validate

```bash
make verify
```
