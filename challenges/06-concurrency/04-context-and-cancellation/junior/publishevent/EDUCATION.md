# Cancellable Event Publish

## Intuition

Go's `select` treats sends and receives symmetrically, so any blocking send becomes cancellable by pairing it with `<-ctx.Done()`. Under backpressure the handler then fails fast with the request's own error instead of quietly holding resources while the broker recovers.

## Approach

1. `select` with two cases.
2. `<-ctx.Done()` → return `ctx.Err()`.
3. `out <- event` → return nil.

## Solution

```go
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
	select {
	case <-ctx.Done():
		return ctx.Err()
	case out <- event:
		return nil
	}
}
```

## Walkthrough

- With room in the buffer the send case is ready and the event is queued.
- With a full buffer and a finished context only the done case is ready, so the error is returned and nothing is enqueued.
- An unbuffered channel with no receiver behaves like a full one — the done case wins.

## Pitfalls

- Sending outside the `select` blocks forever when the writer is stuck.
- Do not add `default:` — dropping the event whenever the buffer is momentarily full loses data that would have fit a microsecond later.
- The send happens or it does not; never report success on the cancellation path.
