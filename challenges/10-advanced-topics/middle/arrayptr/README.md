# A Pointer To An Array Is Not A Slice

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A fixed-size buffer is passed as `[8]int` and copied on every call. Switching to a slice would work and would also lose the compiler's guarantee that the length is exactly eight.

## Task

Implement [arrayptr.go](arrayptr.go):

1. Total the array `a` points at.
2. A nil pointer totals 0.
3. Zero allocations; the caller's array must be the one read.

Replace the stub body in [arrayptr.go](arrayptr.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Sum(&[8]int{1,2})
Output: 3
```

**Example 2:**

```
Input:  a[0] = 5 after taking &a
Output: Sum sees 5
```

_Explanation:_ The pointer is a live view.

**Example 3:**

```
Input:  Sum(nil)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | ***[N]T can be ranged** | The length is in the type, so `range a` works on the pointer directly. |
| 2 | **No header, no copy** | One word is passed, and the array is not duplicated. |
| 3 | **Length as a type guarantee** | A `*[8]int` cannot be given a seven-element array. |

## Hint

`for _, v := range a` works on the pointer. No dereference needed.

## Validate

```bash
make verify
```
