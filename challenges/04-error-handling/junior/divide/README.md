# Safe Division

**Level:** junior
**Topic:** 04-error-handling

## Context

A billing report divides totals by the number of invoices. When a customer has zero invoices the report must fail loudly instead of crashing.

## Task

Implement `Divide` in [divide.go](divide.go):

1. Return `a / b` and a nil error for a non-zero divisor.
2. Return `0` and `ErrDivideByZero` when `b` is zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Divide(10, 2)
Output: 5, nil
```

**Example 2:**

```
Input:  Divide(7, 2)
Output: 3, nil
```

**Example 3:**

```
Input:  Divide(1, 0)
Output: 0, ErrDivideByZero
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Errors are values** | `error` is an ordinary return value, not an exception. |
| 2 | **Multiple return values** | Go returns the result and the error side by side. |
| 3 | **Sentinel comparison** | Callers compare against a package-level error variable. |

## Hint

Check the divisor before dividing. Integer division by zero panics at runtime — the guard must come first.

## Validate

```bash
make verify
```
