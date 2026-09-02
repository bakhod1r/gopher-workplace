# Retry Once

**Level:** junior
**Topic:** 04-error-handling

## Context

A flaky network call succeeds on a second try often enough to be worth one retry. More than that is the caller's problem.

## Task

Implement `RetryTwice` in [retrytwice.go](retrytwice.go):

1. Call `f` and return nil when it succeeds.
2. Call `f` a second time when the first call fails.
3. Return the second call's error when both attempts fail.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RetryTwice(alwaysOK)
Output: nil
```

**Example 2:**

```
Input:  RetryTwice(failsOnce)
Output: nil
```

**Example 3:**

```
Input:  RetryTwice(alwaysFails)
Output: the last error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Function values** | A `func() error` can be stored and called. |
| 2 | **Retry semantics** | Attempt count is fixed and small. |
| 3 | **Returning the last error** | The most recent failure is the informative one. |

## Hint

Two attempts total, not two retries. The tests count the calls.

## Validate

```bash
make verify
```
