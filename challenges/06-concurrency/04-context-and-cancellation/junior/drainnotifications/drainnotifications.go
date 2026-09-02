// Package drainnotifications — Gopher Workplace challenge.
package drainnotifications

import "context"

// Drain collects every pending notification until the producer closes the
// channel, then returns them in arrival order. If the subscriber's context
// finishes first it returns what it has so far together with ctx.Err(), so the
// caller can requeue the undelivered remainder.
//
// Examples:
//
//	Drain(live ctx, closed chan "a","b")  => ["a", "b"], nil
//	Drain(live ctx, closed empty chan)    => [], nil
//	Drain(cancelled ctx, empty chan)      => [], context.Canceled
func Drain(ctx context.Context, ch <-chan string) ([]string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
