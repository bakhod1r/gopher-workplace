// Package paymentretry — Gopher Workplace challenge.
package paymentretry

import "context"

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
	// TODO(candidate): implement this.
	panic("not implemented")
}
