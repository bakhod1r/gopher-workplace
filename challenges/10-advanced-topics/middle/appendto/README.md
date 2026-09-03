# Append Into The Caller's Slice

**Level:** middle
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

Every helper in a hot pipeline returns a freshly made slice. The caller immediately copies each result into its own buffer and drops the original.

## Task

Implement [appendto.go](appendto.go):

1. Append the squares of `0..n-1` to `dst` and return the result.
2. With enough capacity in `dst`, allocate nothing.
3. `n <= 0` returns `dst` unchanged.

Replace the stub body in [appendto.go](appendto.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  AppendSquares(nil, 4)
Output: [0 1 4 9]
```

**Example 2:**

```
Input:  AppendSquares([]int{7}, 2)
Output: [7 0 1]
```

_Explanation:_ dst is extended, not replaced.

**Example 3:**

```
Input:  dst with cap 64, n = 64
Output: 0 allocations
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Append-style APIs** | `f(dst, ...) []T` is the idiom for handing allocation control to the caller. |
| 2 | **Reusing capacity** | `dst[:0]` between calls turns a per-call allocation into none. |
| 3 | **append returns a new header** | The result must be reassigned; the old header may be stale. |

## Hint

The signature already tells you where the output goes.

## Validate

```bash
make verify
```
