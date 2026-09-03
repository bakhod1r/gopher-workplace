# Truncation In A Closed Form

**Level:** senior  
**Topic:** 03-generics

## Context

A capacity estimator is short by exactly half a row whenever the range has an odd number of entries.

## Task

Fix the single planted bug in [divbug.go](divbug.go):

1. Find and fix the single bug so the closed form stays exact.
2. Keep the closed form — do not loop.
3. The empty-range contract must stay.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SumRange(1, 4)
Output: 10
```

**Example 2:**

```
Input:  SumRange(1, 2)
Output: 3
```

**Example 3:**

```
Input:  SumRange(5, 1)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Integer division truncates** | Dividing before multiplying discards the remainder permanently. |
| 2 | **Exactness argument** | `(lo+hi)` and `n` are never both odd, so dividing last is always exact. |
| 3 | **Order of operations is semantics** | Both expressions are algebraically equal over the rationals, not over integers. |

## Hint

`(1+2)/2` is `1` in Go. What does that do to the result?

## Validate

```bash
make verify
```
