# Disconnect Beats the Timeout

## Intuition

A context created with `WithTimeout` can end in two ways, and only the first one counts. Cancelling an hour before the deadline latches `context.Canceled`, and the timer that fires later can no longer change the answer. That single stored value is what lets the access log tell a client disconnect apart from a real timeout — with no clock assertion anywhere in the test.

## Approach

1. `ctx, cancel := context.WithTimeout(context.Background(), time.Hour)`.
2. `defer cancel()` for hygiene, then call `cancel()` immediately.
3. `<-ctx.Done()`.
4. Return `ctx.Err()`.

## Solution

```go
import (
	"context"
	"time"
)

// DisconnectDuringTimeout models a request that was given a generous timeout
// but whose client hung up long before the deadline: the handler's cancel func
// runs first. It returns the reason the context ultimately recorded.
//
// Examples:
//
//	DisconnectDuringTimeout()                              => context.Canceled
//	errors.Is(DisconnectDuringTimeout(), context.Canceled) => true
//	the result is never context.DeadlineExceeded
func DisconnectDuringTimeout() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	cancel()

	<-ctx.Done()
	return ctx.Err()
}
```

## Walkthrough

- The context is created live, with a deadline an hour away.
- `cancel()` finishes it right now and stores `context.Canceled`.
- `<-ctx.Done()` returns immediately and `Err()` reports the stored reason; the deferred second `cancel()` is a harmless no-op.

## Pitfalls

- Expecting `DeadlineExceeded` because the constructor was `WithTimeout` — the constructor does not decide the reason, the first event does.
- Waiting for the real deadline would make the test take an hour; the point is that no wall-clock wait is needed.
- Calling `cancel()` twice (inline and deferred) is safe by contract — cancel funcs are idempotent.
