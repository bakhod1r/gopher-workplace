# Retry Until the Context Says Stop

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

Charges against the payment provider are retried on transient declines. Two things make an unguarded retry loop dangerous: the client may already be gone, and the provider rate-limits aggressively. Checking the context at the top of every attempt bounds the loop by the request's own lifetime.

## Task

Implement the exported function(s) in [paymentretry.go](paymentretry.go) so that:

1. It loops at most `attempts` times.
2. Before each attempt it returns `ctx.Err()` if the context has finished.
3. It returns `nil` at the first successful charge.
4. After the attempts are used up it returns the last error from `charge`; with `attempts <= 0` it calls nothing and returns `nil`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ChargeWithRetry(live ctx, 3, fails twice then succeeds)
Output: nil
```

**Example 2:**

```
Input:  ChargeWithRetry(live ctx, 2, always fails)
Output: errDeclined
```

**Example 3:**

```
Input:  ChargeWithRetry(cancelled ctx, 3, always fails)
Output: context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Context check inside the loop** | Re-checked every iteration, not once before the loop. |
| 2 | **Error precedence** | A finished context outranks a retryable provider error. |
| 3 | **Bounded retries** | `attempts` caps the work; the context caps the wall-clock exposure. |

## Hint

Check `ctx.Err()` at the top of each iteration, before calling `charge`.

## Validate

```bash
make verify
```
