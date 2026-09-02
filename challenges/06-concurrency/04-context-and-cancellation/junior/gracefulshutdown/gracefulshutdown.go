// Package gracefulshutdown — Gopher Workplace challenge.
package gracefulshutdown

import "context"

// WaitForShutdown blocks the background metrics flusher until the process
// shutdown context finishes, then reports why the process is going down:
// context.Canceled for a SIGTERM, context.DeadlineExceeded when the drain
// window expired.
//
// Examples:
//
//	ctx cancelled by the signal handler  => context.Canceled
//	ctx built with an expired drain window => context.DeadlineExceeded
//	ctx cancelled twice                  => context.Canceled
func WaitForShutdown(ctx context.Context) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
