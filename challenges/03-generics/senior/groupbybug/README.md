# Groups That Keep Only The Last

**Level:** senior  
**Topic:** 03-generics

## Context

A grouping report shows exactly one row per category no matter how much data goes in.

## Task

Fix the single planted bug in [groupbybug.go](groupbybug.go):

1. Find and fix the single bug so every element lands in its bucket.
2. Elements within a bucket must keep their input order.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  GroupBy([1,2,3,4], parity)
Output: map[0:[2 4] 1:[1 3]]
```

**Example 2:**

```
Input:  GroupBy([1], parity)
Output: map[1:[1]]
```

**Example 3:**

```
Input:  GroupBy([], parity)
Output: map[]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil slices append fine** | `append(out[k], v)` works even when the key is absent — the zero value is a usable nil slice. |
| 2 | **Assignment discards** | Replacing the bucket throws away everything collected so far. |

## Hint

What is in the bucket before the assignment?

## Validate

```bash
make verify
```
