# Union That Links The Wrong Nodes

**Level:** senior  
**Topic:** 03-generics

## Context

A cluster-membership check says two machines are unrelated even though a chain of merges joined them minutes ago.

## Task

Fix the single planted bug in [unionfindbug.go](unionfindbug.go):

1. Find and fix the single bug so merging works through existing roots.
2. Chained unions must leave every member connected.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Union(1,2); Connected(1,2)
Output: true
```

**Example 2:**

```
Input:  Union(1,2); Union(2,3); Connected(1,3)
Output: true
```

**Example 3:**

```
Input:  Union(1,1)
Output: no change
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structural invariants** | Every operation must restore what the type promises about itself. |
| 2 | **Roots, not members** | Repointing a non-root node detaches whatever hung below it. |

## Hint

Which nodes should `parent` be updated for?

## Validate

```bash
make verify
```
