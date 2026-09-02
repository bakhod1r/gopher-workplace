# Accumulator

**Level:** junior  
**Topic:** 03-generics

## Context

A latency probe reports a running mean without keeping every sample.

## Task

Implement the stub(s) in [accgen.go](accgen.go):

1. Implement `Add`, `Sum`, and `Mean`.
2. `Mean` returns `0` before anything is added, and never truncates.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(1); Add(2); Sum()
Output: 3
```

**Example 2:**

```
Input:  Add(1); Add(2); Mean()
Output: 1.5
```

**Example 3:**

```
Input:  Mean() before any Add
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A constrained numeric type** | `Acc[T Number]` means `+=` is legal on the field. |
| 2 | **Converting for division** | Reused from earlier: `float64(a.total) / float64(a.n)` avoids integer truncation. |
| 3 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |

## Hint

`Sum` keeps type `T`; `Mean` converts to `float64`.

## Validate

```bash
make verify
```
