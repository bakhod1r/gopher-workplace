# `b.N` Is Not The Input Size

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

A benchmark's ns/op climbs steadily with `-benchtime`, and someone concludes the code has a scaling problem. It does not. The body is using the iteration count as the size of the data it processes, so every extra iteration is also a bigger workload.

## Task

Fix the single planted bug in [bnassizebug.go](bnassizebug.go):

1. Find and fix the one bug so the total work is `n` iterations over an input of `size`.
2. `PerOp` must return the same value for every `n`.
3. Doubling `size` must double the total work; the existing guards must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Work(3, 10)
Output: 30
```

**Example 2:**

```
Input:  PerOp(1, 10) and PerOp(1000000, 10)
Output: the same number
```

**Example 3:**

```
Input:  Work(5, 20) versus Work(5, 10)
Output: exactly double
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two independent dimensions** | The harness owns the iteration count; the benchmark owns the input size. |
| 2 | **Mixing them makes ns/op grow** | The measurement then describes the harness, not the code. |
| 3 | **Sub-benchmarks are for sizes** | `b.Run("size=1000", ...)` is how you vary the input on purpose. |

## Hint

One of the two factors in the product should not be `n`.

## Validate

```bash
make verify
```
