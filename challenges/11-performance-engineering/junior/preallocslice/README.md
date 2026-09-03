# One Allocation, Not Eleven

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`append` to a slice with no spare capacity allocates a bigger array and copies everything across. Build a 1000-element result that way and you pay for roughly eleven arrays and a thousand copied elements — all avoidable when you already know the final size.

## Task

Implement `Squares` in [preallocslice.go](preallocslice.go):

1. Return `1, 4, 9, ... n*n`.
2. Reach the final size in exactly one allocation, whatever `n` is.
3. A non-positive `n` returns an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  Squares(3)
Output: [1 4 9]
```

**Example 2:**

```
Input:  Squares(1)
Output: [1]
```

**Example 3:**

```
Input:  Squares(-4)
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Growth is amortised, not free** | Each regrow copies every element already written. |
| 2 | **`make([]T, 0, n)`** | Length zero, capacity `n`: append fills it without ever reallocating. |
| 3 | **Length vs capacity** | `make([]T, n)` then appending leaves `n` zeros in front of your data. |

## Topics used again

`make`, `append`, loops.

## Hint

The test asserts on the allocation count, so the capacity hint is not optional.

## Validate

```bash
make verify
```
