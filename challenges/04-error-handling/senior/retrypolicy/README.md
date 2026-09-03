# Retry Only What Is Retryable

**Level:** senior
**Topic:** 04-error-handling

## Context

A client retries transient failures and gives up immediately on rejected input, so bad requests are not hammered against the server.

## Task

Implement `Retry` in [retrypolicy.go](retrypolicy.go):

1. Return nil as soon as `f` succeeds.
2. Retry up to `attempts` times only while the failure matches `ErrTransient`.
3. Return a non-retryable failure immediately, unwrapped.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Retry(3, alwaysTransient)
Output: the last failure after 3 calls
```

**Example 2:**

```
Input:  Retry(3, invalidFn)
Output: ErrInvalid after 1 call
```

**Example 3:**

```
Input:  Retry(3, okFn)
Output: nil after 1 call
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Classifying before retrying** | Not every failure deserves another attempt. |
| 2 | **Attempt budgets** | Bounded loops with an early exit. |
| 3 | **Returning the original error** | A fatal failure is not annotated away. |

## Hint

The tests count calls: a rejected input must cost exactly one attempt.

## Validate

```bash
make verify
```
