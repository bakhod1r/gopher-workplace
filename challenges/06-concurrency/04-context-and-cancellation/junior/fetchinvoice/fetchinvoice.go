// Package fetchinvoice — Gopher Workplace challenge.
package fetchinvoice

import (
	"context"
	"errors"
)

// ErrNotFound reports that no invoice exists for the requested ID.
var ErrNotFound = errors.New("invoice not found")

// FetchInvoice loads one invoice from the billing store. Following the standard
// convention, the context is the first parameter and is checked before the
// lookup runs.
//
// It returns ErrNotFound for a non-positive ID.
//
// Examples:
//
//	FetchInvoice(live ctx, 7)       => "invoice-7", nil
//	FetchInvoice(live ctx, 0)       => "", ErrNotFound
//	FetchInvoice(cancelled ctx, 7)  => "", context.Canceled
func FetchInvoice(ctx context.Context, id int) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
