# Imposing a Request Timeout

## Intuition

A timeout context is a child that finishes at the earlier of two events: its own deadline, or anything that ends its parent. A helper that returns `(ctx, cancel)` hands the release duty to the caller, which is why every such call site in Go reads `ctx, cancel := ...` immediately followed by `defer cancel()`.

## Approach

1. Return `context.WithTimeout(ctx, d)` directly, forwarding both results.

## Solution

```go
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
	return context.WithTimeout(ctx, d)
}
```

## Walkthrough

- With a one-hour budget the derived context is live, has a deadline, and has a real `Done()` channel; calling `cancel` early makes it report `context.Canceled`.
- With a zero or negative budget the deadline is already past, so `Done()` is closed at construction and `Err()` is `context.DeadlineExceeded`.
- With an already-cancelled parent, the child inherits that cancellation immediately — the parent's reason wins over the generous timeout.

## Pitfalls

- Passing `context.Background()` instead of the incoming `ctx` severs the chain: the handler would keep running after the client disconnects.
- Discarding the cancel func with `_` leaks the timer and trips `go vet`'s `lostcancel` check.
- A child cannot extend its parent's life; a longer child timeout is silently capped by the parent.
