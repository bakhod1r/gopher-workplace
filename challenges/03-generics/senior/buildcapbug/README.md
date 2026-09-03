# The Builder That Hands Out Its Buffer

**Level:** senior  
**Topic:** 03-generics

## Context

A query builder is reused across requests. Callers that append to a built result find their additions leaking into the next request's query.

## Task

Fix the single planted bug in [buildcapbug.go](buildcapbug.go):

1. Find and fix the single bug so the built slice is independent of the builder.
2. The contents must still match what was added.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add(1).Add(2).Build()
Output: [1 2]
```

**Example 2:**

```
Input:  append to result, then Add(3).Build()
Output: [1 2 3]
```

**Example 3:**

```
Input:  Build() on an empty builder
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate your internals. |
| 2 | **Defensive copies** | A getter that exposes internal storage cannot enforce any invariant. |

## Hint

What storage does the returned slice point at?

## Validate

```bash
make verify
```
