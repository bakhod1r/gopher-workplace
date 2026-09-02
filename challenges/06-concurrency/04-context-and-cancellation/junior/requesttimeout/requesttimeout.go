// Package requesttimeout — Gopher Workplace challenge.
package requesttimeout

import (
	"context"
	"time"
)

// WithRequestTimeout is the middleware helper that caps how long a handler may
// run. It derives a context from ctx that finishes after d, and returns it
// together with its cancel func — which the caller must defer.
//
// Examples:
//
//	ctx, cancel := WithRequestTimeout(bg, time.Hour) => deadline set, Err() nil
//	ctx, cancel := WithRequestTimeout(bg, 0)         => Err() is DeadlineExceeded
//	ctx, cancel := WithRequestTimeout(cancelled, time.Hour) => Err() is Canceled
func WithRequestTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
