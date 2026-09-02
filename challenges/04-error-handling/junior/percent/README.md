# Percentage Of Total

**Level:** junior
**Topic:** 04-error-handling

## Context

A progress bar shows how far a job has run. A job with no work at all has no meaningful percentage.

## Task

Implement `Percent` in [percent.go](percent.go):

1. Return `part / total * 100` as a `float64`.
2. Return `0` and `ErrZeroTotal` when `total` is zero.
3. Return `0` and `ErrNegative` when either argument is negative.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Percent(1, 4)
Output: 25, nil
```

**Example 2:**

```
Input:  Percent(0, 4)
Output: 0, nil
```

**Example 3:**

```
Input:  Percent(1, 0)
Output: 0, ErrZeroTotal
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Integer to float conversion** | `float64(a)` before dividing, not after. |
| 2 | **Division guard** | A zero total has no percentage. |
| 3 | **Argument validation** | Both arguments are checked. |

## Hint

`float64(part / total)` divides as integers first and throws the fraction away — convert before dividing.

## Validate

```bash
make verify
```
