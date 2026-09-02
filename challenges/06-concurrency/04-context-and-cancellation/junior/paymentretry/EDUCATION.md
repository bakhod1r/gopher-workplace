# Retry Until the Context Says Stop

## Intuition

A retry loop is an amplifier: one abandoned request becomes N calls to a provider that may already be rate-limiting you. Checking the context every iteration turns that into an early exit — and checking *before* the call, rather than after, is what guarantees no charge is attempted once the request is dead.

## Approach

1. Keep a `last error`.
2. Loop `i` from 0 to `attempts`.
3. Return `ctx.Err()` if it is non-nil.
4. Call `charge()`; return nil on success, otherwise remember the error.
5. Return `last` after the loop.

## Solution

```go
// ChargeWithRetry calls charge up to attempts times, stopping at the first
// success. Before each attempt it checks the request context, so a client that
// hung up or a budget that ran out stops the retry loop immediately instead of
// hammering the payment provider.
//
// It returns nil on success, ctx.Err() if the context finished, or the last
// error from charge once the attempts are used up.
//
// Examples:
//
//	ChargeWithRetry(live ctx, 3, fails twice then succeeds) => nil
//	ChargeWithRetry(live ctx, 2, always fails)              => errDeclined
//	ChargeWithRetry(cancelled ctx, 3, always fails)         => context.Canceled
func ChargeWithRetry(ctx context.Context, attempts int, charge func() error) error {
	var last error
	for i := 0; i < attempts; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		last = charge()
		if last == nil {
			return nil
		}
	}
	return last
}
```

## Walkthrough

- When the charge succeeds on the third try, exactly three calls are made and nil comes back.
- When every try fails, the loop runs `attempts` times and the provider's last error is returned.
- With a cancelled context the guard fires on the first iteration, so the test observes zero calls to `charge`.
- With `attempts == 0` the loop body never runs and `last` is still nil.

## Pitfalls

- Checking `ctx.Err()` only once before the loop lets the remaining attempts run after the client is gone.
- Returning `nil` when the attempts are exhausted silently reports a failed charge as a success.
- Do not sleep between attempts in a test-facing function; backoff belongs behind a clock the tests control.
