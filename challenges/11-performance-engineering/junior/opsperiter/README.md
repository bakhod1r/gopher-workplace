# Per Operation, Not Per Iteration

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

When a benchmark body processes a batch — 64 records, 8 keys, a whole packet — the ns/op the tool prints is per *iteration*, not per record. Dividing the work count by the iteration count is how you get back to a comparable unit.

## Task

Implement `OpsPerIter` in [opsperiter.go](opsperiter.go):

1. Divide `totalOps` by `iters`.
2. Round half away from zero: `2.5` becomes `3`, `2.25` becomes `2`.
3. A non-positive `iters` returns `0`.

## Examples

**Example 1:**

```
Input:  OpsPerIter(10, 4)
Output: 3
```

**Example 2:**

```
Input:  OpsPerIter(9, 4)
Output: 2
```

**Example 3:**

```
Input:  OpsPerIter(100, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Choosing the unit** | ns/op only compares across changes if "op" means the same thing each run. |
| 2 | **Integer rounding** | `(a + b/2) / b` rounds half up without any floating point. |
| 3 | **Guard before divide** | A zero iteration count is a real case, not a bug to crash on. |

## Hint

Add half the divisor before dividing.

## Validate

```bash
make verify
```
