# Borrow A Buffer Instead Of Allocating One

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A serialiser allocates a fresh scratch buffer for every record it renders. The buffers are identical in size and die immediately — textbook pool material.

## Task

Implement [poolscratch.go](poolscratch.go):

1. Render `vals` as decimal numbers joined by `,`.
2. Take the scratch buffer from `pool` and put it back before returning.
3. Reset the borrowed buffer's length before writing into it.

Replace the stub body in [poolscratch.go](poolscratch.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Encode([]int{1,2,3})
Output: "1,2,3"
```

**Example 2:**

```
Input:  Encode(nil)
Output: ""
```

**Example 3:**

```
Input:  Encode([]int{-7})
Output: "-7"
```

_Explanation:_ The sign is part of the rendering.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Pool** | A free list of reusable values; `Get` may return a recycled one or call `New`. |
| 2 | **Resetting a borrowed buffer** | A pooled buffer arrives with whatever length it was returned at. |
| 3 | **strconv.Append*** | Appends the rendering into an existing buffer instead of allocating a string. |

## Hint

What is the length of a buffer that came back from the pool?

## Validate

```bash
make verify
```
