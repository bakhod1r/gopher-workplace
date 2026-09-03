# Copy Stops At The Shorter Side

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A fixed-size window is filled from a larger stream. The first attempt indexed `src` up to `len(src)` and wrote past the end of `dst`.

## Task

Implement [copyinto.go](copyinto.go):

1. Copy as many elements as fit from `src` into `dst`.
2. Return the number copied.
3. Nothing is resized, and nothing is allocated.

Replace the stub body in [copyinto.go](copyinto.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  CopyInto(make([]int,2), []int{1,2,3})
Output: 2
```

_Explanation:_ dst is the limit.

**Example 2:**

```
Input:  CopyInto([]int{9,9,9}, []int{1})
Output: 1
```

_Explanation:_ src is the limit; the rest of dst is untouched.

**Example 3:**

```
Input:  CopyInto(make([]int,0,8), []int{1,2})
Output: 0
```

_Explanation:_ Capacity is not length.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **copy is bounded by both** | It moves `min(len(dst), len(src))` elements. |
| 2 | **Length, not capacity** | `copy` writes only what `dst`'s length allows. |
| 3 | **Overlap is well defined** | `copy` behaves like memmove. |

## Hint

The builtin already does exactly this, including the return value.

## Validate

```bash
make verify
```
