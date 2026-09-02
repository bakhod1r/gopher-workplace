# Concat That Returns Its Input

**Level:** senior  
**Topic:** 03-generics

## Context

A request-assembly helper takes a shortcut for the single-slice case. Callers that mutate the result are corrupting the source data.

## Task

Fix the single planted bug in [concatalias.go](concatalias.go):

1. Find and fix the single bug so the result never aliases an input.
2. The concatenation itself must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Concat([1],[2])
Output: [1 2]
```

**Example 2:**

```
Input:  Concat([1,2])
Output: a copy, not the input
```

**Example 3:**

```
Input:  Concat()
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate your internals. |
| 2 | **Fast paths break contracts** | An optimisation that returns the argument changes the function's ownership rules. |

## Hint

Look at the single-argument shortcut.

## Validate

```bash
make verify
```
