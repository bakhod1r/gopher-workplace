// Package providerrace — Gopher Workplace challenge.
package providerrace

import (
	"context"
	"errors"
)

// Charge is the amount a checkout wants authorized.
type Charge struct {
	Amount   int
	Currency string
}

// ErrAllProvidersDeclined reports that no provider authorized the charge.
var ErrAllProvidersDeclined = errors.New("all providers declined")

// FirstAuthorization asks every payment provider to authorize the charge at
// the same time and returns the first authorization code that comes back.
// The losing providers are cancelled through a derived context the moment a
// winner is found, so a slow provider cannot hold the checkout open.
//
// It returns ctx.Err() if the caller's context is already finished, and
// ErrAllProvidersDeclined when no provider authorizes.
//
// Examples:
//
//	FirstAuthorization(live ctx, charge, [slow-a ok-b slow-c], auth) => "ok-b:auth"
//	FirstAuthorization(live ctx, charge, [no-a no-b], auth)          => ErrAllProvidersDeclined
//	FirstAuthorization(cancelled ctx, charge, [ok-a], auth)          => context.Canceled
func FirstAuthorization(ctx context.Context, charge Charge, providers []string, authorize func(context.Context, string, Charge) (string, error)) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
