# Allocate The Slice Once

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A report builder grows its result slice one `append` at a time. Under profiling most of its time is spent copying the backing array as it doubles.

## Task

Implement [prealloc.go](prealloc.go):

1. Return the squares of `0..n-1` in order.
2. The whole result must come from one allocation — `AllocsPerRun` must see at most 1.

Replace the stub body in [prealloc.go](prealloc.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Squares(4)
Output: [0 1 4 9]
```

**Example 2:**

```
Input:  Squares(1)
Output: [0]
```

**Example 3:**

```
Input:  Squares(0)
Output: []
```

_Explanation:_ n == 0 gives an empty, non-nil-or-nil slice; length is what is checked.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **make with a length** | `make([]T, n)` reserves the final size in one allocation. |
| 2 | **append growth** | Appending past cap allocates a bigger array and copies — repeatedly. |
| 3 | **Allocation counting** | `testing.AllocsPerRun` grades allocation behaviour, not just output. |

## Hint

You already know the final length before the loop starts.

## Validate

```bash
make verify
```
