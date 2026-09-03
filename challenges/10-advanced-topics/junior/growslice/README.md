# Make Room Before The Appends

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A decoder knows the record count from the header but still lets `append` discover the size, paying a reallocation and a full copy at every doubling.

## Task

Implement [growslice.go](growslice.go):

1. Return a slice with the same length and contents as `s` and room for `n` more elements.
2. Allocate nothing when the spare capacity already covers `n`.
3. Treat `n < 0` as 0.

Replace the stub body in [growslice.go](growslice.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Grow(make([]int,2,2), 8)
Output: len 2, cap >= 10
```

**Example 2:**

```
Input:  Grow(make([]int,1,32), 4)
Output: the same slice, no allocation
```

_Explanation:_ The room is already there.

**Example 3:**

```
Input:  Grow(s, -5)
Output: s unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Spare capacity** | `cap(s) - len(s)` is what an append can use for free. |
| 2 | **make with length and capacity** | `make([]int, len, cap)` keeps the contents indexable and reserves the rest. |
| 3 | **Amortised growth** | Reserving once beats doubling repeatedly. |

## Hint

Compare `cap(s)-len(s)` with `n` before doing anything.

## Validate

```bash
make verify
```
