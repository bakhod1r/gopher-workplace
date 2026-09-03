# The Map That Ships N Zero Values

**Level:** staff  
**Topic:** 03-generics

## Context

A projection step over a batch returns twice as many rows as it was given, and the first half of every batch is blank records.

## Task

Fix the single planted bug in [prealloclenbug.go](prealloclenbug.go):

1. Find and fix the single bug so the result holds exactly one element per input.
2. Order must be preserved and the capacity hint must be kept.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Map([1,2], double)
Output: [2 4]
```

**Example 2:**

```
Input:  len(Map(1e6 rows, f))
Output: 1000000
```

**Example 3:**

```
Input:  Map([], double)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Length versus capacity in `make`** | `make([]T, n)` creates n zero values; `make([]T, 0, n)` reserves room for n. |
| 2 | **Preallocation is not initialisation** | The capacity hint exists to avoid regrowth, not to size the result. |

## Hint

Read the second argument to `make`.

## Validate

```bash
make verify
```
