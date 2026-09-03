# Emptying Without Freeing

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A server that decodes a request per connection can allocate one scratch buffer and empty it between rounds instead of allocating a new one each time. "Empty" here means length zero — the array stays, and so does the win.

## Task

Implement both functions in [sliceresetreuse.go](sliceresetreuse.go):

1. `Reset` returns `s` with length `0` and its capacity intact; nil resets to an empty, non-nil slice.
2. `FillEvens` overwrites `buf` with `0, 2, 4, ... 2*(n-1)`.
3. `FillEvens` must not allocate when `buf`'s capacity already suffices.

## Examples

**Example 1:**

```
Input:  Reset(make([]int, 3, 64))
Output: len 0, cap 64
```

**Example 2:**

```
Input:  FillEvens(nil, 3)
Output: [0 2 4]
```

**Example 3:**

```
Input:  FillEvens([9 9 9 9 9], 2)
Output: [0 2]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`s[:0]` keeps the array** | Length drops to zero, capacity and backing storage are untouched. |
| 2 | **Reuse is the allocation win** | The buffer is allocated once and refilled forever after. |
| 3 | **Stale elements are invisible** | Beyond the new length the old values still exist, which matters if they hold pointers. |

## Topics used again

Slice length and capacity, `append`, loops.

## Hint

`FillEvens` should start from `Reset(buf)`.

## Validate

```bash
make verify
```
