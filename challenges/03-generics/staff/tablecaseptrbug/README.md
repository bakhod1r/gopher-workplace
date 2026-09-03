# Every Case Is The Last Case

**Level:** staff  
**Topic:** 03-generics

## Context

A generic table-test runner collects a pointer per case so subtests can report and mutate their own case. Every subtest reports the *last* case's name, and a failure in case 1 is attributed to case 40.

## Task

Fix the single planted bug in [tablecaseptrbug.go](tablecaseptrbug.go):

1. Find and fix the single bug so each returned pointer addresses its own case.
2. The pointers must come back in input order.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Pointers(3 cases)[0].Name
Output: the first name
```

**Example 2:**

```
Input:  all three names
Output: distinct
```

**Example 3:**

```
Input:  Pointers(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **One variable, many iterations** | A variable declared *outside* the loop is assigned in place, so every `&v` is the same address. |
| 2 | **Per-iteration scope** | Go 1.22 gives loop variables declared *in* the `range` clause a fresh instance each pass — this loop opts out of that by declaring `c` above. |
| 3 | **Index over value** | `&cases[i]` addresses the real element, which is usually what a runner wants anyway. |

## Hint

How many distinct addresses does the loop produce?

## Validate

```bash
make verify
```
