# A Sub-Deadline for Every Attempt

## Intuition

A retry loop needs two deadlines, not one. The outer budget answers "is this request still worth anything?", the inner sub-deadline answers "has *this* attempt earned another try?". Without the inner one a single hung provider swallows the entire budget and the word "retry" becomes decorative. Deriving the sub-deadline from `ctx` keeps the guarantee that a child can never outlive its parent — `WithTimeout` only ever shortens.

## Approach

1. Keep `last error`.
2. Loop `i` from 0 to `attempts`.
3. Return `ctx.Err()` if the request budget is already finished.
4. `attemptCtx, cancel := context.WithTimeout(ctx, perAttempt)`; call `charge(attemptCtx)`; call `cancel()` immediately after.
5. Return nil on success; return `last` after the loop.

## Solution

```go
import (
	"context"
	"time"
)
// Charge attempts a single authorisation against the payment provider.
type Charge func(ctx context.Context) error

// ChargeWithAttemptDeadline retries a card authorisation under two clocks at
// once: the caller's request budget in ctx, and a fresh perAttempt sub-deadline
// derived for each individual attempt. One provider hang can then burn a single
// attempt instead of the whole request budget.
//
// It returns nil at the first success, ctx.Err() if the request budget finished
// before an attempt started, or the error from the final attempt.
//
// Examples:
//
//	ChargeWithAttemptDeadline(ctx, 3, time.Second, succeeds on the 3rd try) => nil
//	ChargeWithAttemptDeadline(ctx, 3, 0, provider that honours its context)  => context.DeadlineExceeded
//	ChargeWithAttemptDeadline(cancelled ctx, 3, time.Second, anything)       => context.Canceled
func ChargeWithAttemptDeadline(ctx context.Context, attempts int, perAttempt time.Duration, charge Charge) error {
	var last error
	for i := 0; i < attempts; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}

		attemptCtx, cancel := context.WithTimeout(ctx, perAttempt)
		last = charge(attemptCtx)
		cancel()

		if last == nil {
			return nil
		}
	}
	return last
}
```

## Walkthrough

- On the happy path the provider succeeds on attempt three; each attempt ran under its own hour-long sub-deadline, and the test's `sawDeadline` flag confirms the attempt context actually carried one.
- With `perAttempt == 0` every attempt context is born expired, so the provider returns `context.DeadlineExceeded` three times and the loop reports the last one — the *request* budget was never touched.
- A cancelled request returns `context.Canceled` with zero provider calls: the guard runs before the sub-context is derived.
- With `attempts <= 0` the body never runs and `last` is still nil.

## Pitfalls

- `defer cancel()` inside the loop: every attempt's timer stays alive until the function returns, which `go vet`'s lostcancel check will not catch but a profiler will.
- Dropping `cancel` entirely leaks a timer per attempt until the deadline fires.
- Deriving the attempt context from `context.Background()` instead of `ctx` — the attempt then outlives a cancelled request.
- Returning the attempt's `DeadlineExceeded` as if the whole request had timed out, when there are still attempts left in the budget.
