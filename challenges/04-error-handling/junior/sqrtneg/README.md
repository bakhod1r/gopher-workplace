# Square Root Guard

**Level:** junior
**Topic:** 04-error-handling

## Context

A geometry service computes distances. Negative inputs are a caller bug, so the function refuses them rather than returning NaN.

## Task

Implement `Sqrt` in [sqrtneg.go](sqrtneg.go):

1. Return the square root of `x` for `x >= 0`.
2. Return `0` and `ErrNegative` for `x < 0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sqrt(9)
Output: 3, nil
```

**Example 2:**

```
Input:  Sqrt(0)
Output: 0, nil
```

**Example 3:**

```
Input:  Sqrt(-1)
Output: 0, ErrNegative
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Validating inputs** | Reject the bad case before doing the work. |
| 2 | **math.Sqrt** | The standard library computes the root; the guard is yours. |
| 3 | **Float results** | `float64` results still pair with an `error`. |

## Hint

`math.Sqrt(-1)` returns NaN silently — that is exactly the outcome the guard prevents.

## Validate

```bash
make verify
```
