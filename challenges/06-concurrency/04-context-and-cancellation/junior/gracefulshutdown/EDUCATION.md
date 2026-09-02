# Graceful Shutdown Wait

## Intuition

`Done()` is a channel that is *closed*, never written to. A closed channel makes every receive return immediately, so any number of waiting goroutines wake at once — the broadcast primitive cancellation needs. Once it is closed, `Err()` is guaranteed non-nil and tells you which of the two ways the context ended.

## Approach

1. `<-ctx.Done()` to block until the context finishes.
2. Return `ctx.Err()`.

## Solution

```go
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
	<-ctx.Done()
	return ctx.Err()
}
```

## Walkthrough

- For an already-cancelled context the receive returns immediately.
- `Err()` then reports `context.Canceled` for a `cancel()` call and `context.DeadlineExceeded` for an expired deadline.
- A child of a cancelled parent is cancelled too, and reports the parent's reason.

## Pitfalls

- Polling `for ctx.Err() == nil {}` burns a CPU core and is a data-race magnet; block on the channel instead.
- Reading `ctx.Err()` *before* `Done()` fires can return nil — always wait first.
- Never `close(ctx.Done())` yourself; the channel is read-only and owned by the context.
