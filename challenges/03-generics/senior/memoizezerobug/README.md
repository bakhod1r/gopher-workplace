# The Cache That Never Fills

**Level:** senior  
**Topic:** 03-generics

## Context

A memoised pricing lookup shows no cache hits at all in the metrics, and the upstream service is rate-limiting the service.

## Task

Fix the single planted bug in [memoizezerobug.go](memoizezerobug.go):

1. Find and fix the single bug so a computed result is remembered.
2. The same key must never call `fn` twice.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  f(1) twice
Output: fn called once
```

**Example 2:**

```
Input:  f(1), f(2)
Output: fn called twice
```

**Example 3:**

```
Input:  returned value
Output: equals fn(k)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures own state** | The map lives in the closure, so every call sees the same cache. |
| 2 | **Read and write** | A lookup-only cache is just an expensive indirection. |

## Hint

What is missing between computing and returning?

## Validate

```bash
make verify
```
