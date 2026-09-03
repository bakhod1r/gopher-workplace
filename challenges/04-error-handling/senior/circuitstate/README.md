# Open Circuit

**Level:** senior
**Topic:** 04-error-handling

## Context

A breaker stops calling a failing dependency after a threshold, and reports that it is open rather than attempting the call.

## Task

Implement `Breaker` in [circuitstate.go](circuitstate.go):

1. Call `f` and reset the failure count on success.
2. Count consecutive failures and return the error from `f`.
3. Return `ErrOpen` without calling `f` once the count reaches `Threshold`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  threshold 2, two failures then Call
Output: ErrOpen, f not called
```

**Example 2:**

```
Input:  success resets the count
Output: true
```

**Example 3:**

```
Input:  Call with threshold 0
Output: ErrOpen immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fail fast** | An open circuit protects the dependency and the caller. |
| 2 | **Consecutive vs total** | A success clears the accumulated state. |
| 3 | **State on a receiver** | The breaker mutates through a pointer. |

## Hint

The check happens before the call, and a success must clear the counter entirely.

## Validate

```bash
make verify
```
