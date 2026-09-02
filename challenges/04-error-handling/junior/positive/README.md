# Positive Amount

**Level:** junior
**Topic:** 04-error-handling

## Context

An invoice line rejects non-positive quantities. Zero is not a valid order line either.

## Task

Implement `Positive` in [positive.go](positive.go):

1. Return `n` and nil when `n > 0`.
2. Return `0` and `ErrNotPositive` when `n <= 0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Positive(5)
Output: 5, nil
```

**Example 2:**

```
Input:  Positive(0)
Output: 0, ErrNotPositive
```

**Example 3:**

```
Input:  Positive(-3)
Output: 0, ErrNotPositive
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Strict comparison** | Positive excludes zero. |
| 2 | **Value plus error** | Return the zero value alongside a failure. |
| 3 | **Guard clauses** | The invalid case exits first. |

## Hint

Zero belongs on the failing side of the comparison.

## Validate

```bash
make verify
```
