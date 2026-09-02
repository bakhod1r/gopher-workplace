// Package chargeguard — Gopher Workplace challenge.
package chargeguard

import "context"

// Charge runs the card capture only if the request context is still alive.
// Capturing a payment for a caller that has already gone away bills a customer
// nobody will ever show a receipt to, so the guard runs before the side effect,
// not after.
//
// Examples:
//
//	Charge(live ctx, ok)        => "captured", nil
//	Charge(cancelled ctx, ok)   => "", context.Canceled  (capture never runs)
//	Charge(expired ctx, ok)     => "", context.DeadlineExceeded
func Charge(ctx context.Context, capture func() string) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
