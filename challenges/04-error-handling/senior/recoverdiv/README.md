# Panic To Error

**Level:** senior
**Topic:** 04-error-handling

## Context

A rules engine evaluates user-supplied expressions. A division by zero inside one must fail that expression, not take the process down.

## Task

Implement `SafeDivide` in [recoverdiv.go](recoverdiv.go):

1. Return `a / b` and nil when the division succeeds.
2. Recover any panic and return it as `ErrPanic`, with a zero result.
3. Never let a panic escape the function.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SafeDivide(10, 2)
Output: 5, nil
```

**Example 2:**

```
Input:  SafeDivide(1, 0)
Output: 0, ErrPanic
```

**Example 3:**

```
Input:  SafeDivide(-8, 4)
Output: -2, nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **defer and recover** | `recover` only works inside a deferred function. |
| 2 | **Named return values** | The deferred function writes to the named results. |
| 3 | **Boundary recovery** | Panics are converted to errors at a controlled edge. |

## Hint

A deferred closure can assign to named results — that is the only way to change what a panicking function returns.

## Validate

```bash
make verify
```
