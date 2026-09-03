# Walk An Array By Pointer

**Level:** senior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A binding wraps a C library that returns a pointer and a count. The Go side has to read the values without owning the memory or making a slice of it.

## Task

Implement [pointerwalk.go](pointerwalk.go):

1. Total `n` consecutive int32 values starting at `p`.
2. Return 0 for a nil pointer or `n <= 0`.
3. Accumulate in int64 so a long run cannot overflow; allocate nothing.

Replace the stub body in [pointerwalk.go](pointerwalk.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  SumInt32(&a[0], 3) over [1 2 3 4]
Output: 6
```

**Example 2:**

```
Input:  SumInt32(nil, 3)
Output: 0
```

**Example 3:**

```
Input:  eight values of 1<<30
Output: 8589934592
```

_Explanation:_ The accumulator is wider than the elements.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer stride** | Advance by `i * unsafe.Sizeof(*p)`, not by `i`. |
| 2 | **unsafe.Add** | Keeps the arithmetic in pointer space. |
| 3 | **Accumulator width** | Summing int32 into int32 overflows long before int64 does. |
| 4 | **No bounds check exists** | `n` is a promise from the caller; the runtime cannot verify it. |

## Hint

The step between elements is `unsafe.Sizeof(*p)` bytes, not one.

## Validate

```bash
make verify
```
