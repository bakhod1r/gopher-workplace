# A Sub-Deadline for Every Attempt

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A card authorisation is retried up to three times inside an HTTP handler that has its own request budget. Retrying against the raw request context is a trap: one provider that accepts the connection and then hangs consumes the entire budget in the first attempt, and the retries never happen. Each attempt therefore gets its own short deadline derived from the request context.

## Task

Implement the exported function(s) in [attemptdeadline.go](attemptdeadline.go) so that:

1. It loops at most `attempts` times and returns `ctx.Err()` at the top of any iteration where the request budget has finished.
2. It derives `attemptCtx` with `context.WithTimeout(ctx, perAttempt)` for each attempt and passes it to `charge`.
3. It releases the attempt's timer as soon as the attempt is over — not with a `defer` that piles up until the function returns.
4. It returns `nil` at the first success, and the last attempt's error once the attempts are exhausted.
5. With `attempts <= 0` it calls nothing and returns `nil`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ChargeWithAttemptDeadline(ctx, 3, time.Hour, declines twice then succeeds)
Output: nil  (3 calls)
```

**Example 2:**

```
Input:  ChargeWithAttemptDeadline(ctx, 3, 0, provider honouring its context)
Output: context.DeadlineExceeded  (3 calls)
```

**Example 3:**

```
Input:  ChargeWithAttemptDeadline(cancelled ctx, 3, time.Hour, anything)
Output: context.Canceled  (0 calls)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deadline inheritance** | `WithTimeout(ctx, d)` can only shorten: the child dies at `min(parent deadline, now+d)`. |
| 2 | **`cancel` inside a loop** | `defer cancel()` in a loop body defers to *function* exit, holding every timer alive. |
| 3 | **Budget vs attempt error** | A finished request budget outranks a retryable decline and ends the loop. |

## Hint

Two clocks: the request budget checked at the top of the loop, and a fresh `WithTimeout(ctx, perAttempt)` per iteration whose `cancel` runs before the next one starts.

## Validate

```bash
make verify
```
