# Query Context Before and After Disconnect

## Intuition

A context's error is a one-way latch. It reads nil for as long as the query may run, then flips to a sentinel the instant the client disconnects, and stays there forever. That is why a driver goroutine can read `Err()` safely from any goroutine after `Done()` fires and trust the answer.

## Approach

1. Create the cancellable query context.
2. Record `ctx.Err()` as `connected`.
3. Call `cancel()` to model the disconnect.
4. Record `ctx.Err()` as `disconnected` and return both.

## Solution

```go
import "context"

// QueryState models the reporting endpoint's database query context. It
// returns the context's error twice: once while the client is still connected
// and the query is allowed to run, and once after the client disconnected and
// the handler cancelled it.
//
// Examples:
//
//	connected, _ := QueryState()      => connected is nil
//	_, disconnected := QueryState()   => disconnected is context.Canceled
//	connected != disconnected         => true
func QueryState() (connected, disconnected error) {
	ctx, cancel := context.WithCancel(context.Background())
	connected = ctx.Err()
	cancel()
	disconnected = ctx.Err()
	return connected, disconnected
}
```

## Walkthrough

- Fresh context: nothing has cancelled it, so `Err()` is nil and the query is allowed to proceed.
- `cancel()` closes `Done()` and stores `context.Canceled`.
- The second read observes the stored sentinel, which is what the driver reports back as the query's failure reason.

## Pitfalls

- `defer cancel()` would run after both reads, so `disconnected` would be nil — call `cancel()` inline.
- Do not assume `Err()` is non-nil merely because a context exists; a live context reports nil.
