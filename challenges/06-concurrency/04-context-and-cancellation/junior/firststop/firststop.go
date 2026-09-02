// Package firststop — Gopher Workplace challenge.
package firststop

import "context"

// FirstStop blocks until either the per-request context or the process
// shutdown context finishes, and returns the error of whichever finished. At
// least one of the two is always guaranteed to finish.
//
// If the request context is already finished, its error is returned; otherwise
// the shutdown context's error is.
//
// Examples:
//
//	FirstStop(cancelled req, live shutdown)   => context.Canceled
//	FirstStop(live req, cancelled shutdown)   => context.Canceled
//	FirstStop(live req, expired shutdown)     => context.DeadlineExceeded
func FirstStop(reqCtx, shutdownCtx context.Context) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
