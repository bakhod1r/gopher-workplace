// Package requestbudget — Gopher Workplace challenge.
package requestbudget

import (
	"context"
	"time"
)

// Budget reports the deadline the caller imposed on this request and whether
// one exists at all. The object-storage client calls it before an upload to
// decide between a single PUT and a slower multipart transfer.
//
// Examples:
//
//	Budget(context.Background())            => zero time, false
//	Budget(ctx with deadline t)             => t, true
//	Budget(context.WithValue(bg, k, v))     => zero time, false
func Budget(ctx context.Context) (time.Time, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
