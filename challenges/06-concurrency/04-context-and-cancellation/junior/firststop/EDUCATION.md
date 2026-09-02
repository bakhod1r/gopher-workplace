# Request or Shutdown, Whichever Comes First

## Intuition

`select` waits on several channels at once, which makes it the natural join point for independent cancellation sources. A context that can never finish contributes a nil channel — permanently not ready — so it simply never wins, and the other source decides. Returning the *matching* context's error is what keeps the access log honest about whether the client left or the server did.

## Approach

1. `select` with a case per context.
2. In each case, return that context's `Err()`.

## Solution

```go
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
	select {
	case <-reqCtx.Done():
		return reqCtx.Err()
	case <-shutdownCtx.Done():
		return shutdownCtx.Err()
	}
}
```

## Walkthrough

- When the request context is already cancelled its case is ready and `context.Canceled` is returned.
- When only the shutdown context has finished, its branch runs and reports its reason — `Canceled` for a signal, `DeadlineExceeded` for an expired drain window.
- Passing `context.Background()` as one side gives a nil `Done()` channel, which is never ready, so the other side always wins.
- The tests never make both sides ready at once, because `select` would then choose at random.

## Pitfalls

- Returning `reqCtx.Err()` from both branches misreports a deploy as a client disconnect.
- Nesting two sequential receives serialises the wait: the second source could never be observed first.
- For the general case Go offers `context.AfterFunc` and, in newer versions, `context.WithoutCancel`; a two-case select is the right tool at this scale.
