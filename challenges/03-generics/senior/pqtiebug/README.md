# Insertion That Breaks Stability

**Level:** senior  
**Topic:** 03-generics

## Context

A priority queue built on sorted insertion keeps serving the newest job first whenever two jobs share a priority.

## Task

Fix the single planted bug in [pqtiebug.go](pqtiebug.go):

1. Find and fix the single bug so a new element lands after existing elements with the same key.
2. The slice must stay sorted by key.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  InsertSorted([p1,p3], p2, prio)
Output: [p1 p2 p3]
```

**Example 2:**

```
Input:  a second job at priority 2
Output: goes after the first
```

**Example 3:**

```
Input:  InsertSorted([], j, prio)
Output: [j]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stability** | Equal elements must keep their input order unless the doc says otherwise. |
| 2 | **Scan predicate decides the tie** | `<` stops before equals; `<=` walks past them. |

## Hint

One character in the scan condition.

## Validate

```bash
make verify
```
