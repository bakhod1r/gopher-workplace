# Retry With Attempt Count

**Level:** middle
**Topic:** 04-error-handling

## Context

A client retries a flaky call a fixed number of times. When it finally gives up, the error says how many attempts were spent.

## Task

Implement `Retry` in [retryattempts.go](retryattempts.go):

1. Call `f` up to `attempts` times, stopping at the first success.
2. Return an error wrapping the last failure as `"after <n> attempts: <err>"`.
3. Return `ErrNoAttempts` without calling `f` when `attempts <= 0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Retry(3, alwaysOK)
Output: nil, 1 call
```

**Example 2:**

```
Input:  Retry(3, alwaysFails)
Output: "after 3 attempts: …"
```

**Example 3:**

```
Input:  Retry(0, f)
Output: ErrNoAttempts, 0 calls
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded loops** | The attempt count is the loop bound. |
| 2 | **Wrapping the final error** | The last failure is the one worth keeping. |
| 3 | **Degenerate input** | Zero attempts is its own failure. |

## Hint

The tests count calls — a successful first attempt must not trigger a second.

## Validate

```bash
make verify
```
