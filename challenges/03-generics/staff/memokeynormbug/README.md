# The Cache Whose Keys Never Match

**Level:** staff  
**Topic:** 03-generics

## Context

An upstream-resolution cache reports a hit rate of zero and the cache itself grows without bound. Hosts arrive from several sources with different capitalisation and stray whitespace, and a struct key compares field by field.

## Task

Fix the single planted bug in [memokeynormbug.go](memokeynormbug.go):

1. Find and fix the single bug so equivalent keys land on one cache entry.
2. The value handed to `fn` must be the normalised key.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Get({"API.example.com",443}) then Get({"api.example.com",443})
Output: fn called once
```

**Example 2:**

```
Input:  Get({" api.example.com ",443})
Output: same entry
```

**Example 3:**

```
Input:  different ports
Output: different entries
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct keys compare field by field** | `Key{"API",443} != Key{"api",443}` — the map is behaving exactly as specified. |
| 2 | **Normalise at the boundary** | Every path into the cache must funnel through one canonical form. |
| 3 | **A zero hit rate is a bug signal** | A cache that never hits is pure overhead plus a memory leak. |

## Hint

Compare the key that is looked up with the key `Norm` would produce.

## Validate

```bash
make verify
```
