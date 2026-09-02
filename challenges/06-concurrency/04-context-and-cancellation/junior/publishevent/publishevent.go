// Package publishevent — Gopher Workplace challenge.
package publishevent

import "context"

// Publish hands one domain event to the outbound Kafka writer's channel, or
// gives up if the request context finishes first. A full channel means the
// writer is backed up; the caller must not block on it past the request's
// lifetime.
//
// Examples:
//
//	Publish(live ctx, chan with room, "order.created")  => nil
//	Publish(cancelled ctx, full chan, "order.created")  => context.Canceled
//	Publish(expired ctx, full chan, "order.created")    => context.DeadlineExceeded
func Publish(ctx context.Context, out chan<- string, event string) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
