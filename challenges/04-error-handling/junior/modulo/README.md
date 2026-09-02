# Safe Remainder

**Level:** junior
**Topic:** 04-error-handling

## Context

A sharding helper maps IDs onto buckets with a modulo. A zero bucket count is a configuration error.

## Task

Implement `Mod` in [modulo.go](modulo.go):

1. Return `a % b` and nil for a non-zero `b`.
2. Return `0` and `ErrZeroModulus` when `b` is zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Mod(10, 3)
Output: 1, nil
```

**Example 2:**

```
Input:  Mod(-7, 3)
Output: -1, nil
```

**Example 3:**

```
Input:  Mod(5, 0)
Output: 0, ErrZeroModulus
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Modulo semantics** | Go's `%` keeps the sign of the dividend. |
| 2 | **Runtime panic** | `x % 0` panics exactly like division by zero. |
| 3 | **Guard before use** | Validate the operand, then compute. |

## Hint

`%` panics on a zero operand just as `/` does — and Go's remainder can be negative.

## Validate

```bash
make verify
```
