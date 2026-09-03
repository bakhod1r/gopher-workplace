// Package attemptdeadline — Gopher Workplace challenge.
package attemptdeadline

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
	// TODO(candidate): implement this.
	panic("not implemented")
}
