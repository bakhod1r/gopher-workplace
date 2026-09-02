# Retry With Attempt Budget

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Transient failures deserve another try — but only a bounded number. The retry
loop owns the budget: it calls the operation until one succeeds or the attempts
run out, and reports the last error if they do.

## Task

Implement `DoWithRetry` on `*Client` in [retrylogic.go](retrylogic.go):

1. Loop at most `maxAttempts` times.
2. Call `c.Do()`; return `nil` immediately on success.
3. If every attempt failed, return the last error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FailInt 2, DoWithRetry(3)
Output: nil; Attempts == 3
```

**Example 2:**

```
Input:  FailInt 5, DoWithRetry(3)
Output: the last error; Attempts == 3
```

**Example 3:**

```
Input:  FailInt 0, DoWithRetry(3)
Output: nil; Attempts == 1 (no retry needed)
```

_Explanation:_ the budget is a maximum, not a quota to spend.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded loop** | `maxAttempts` counts calls, not retries after the first. |
| 2 | **Keeping the last error** | Declared outside the loop so it survives to the return. |
| 3 | **Early return on success** | Stops the loop *and* stops burning attempts. |

## Hint

`var err error` before the loop; assign it each iteration. Returning inside the
loop on failure turns the retry into a single call — the `Attempts` assertions
catch it.

## Validate

```bash
make verify
```
