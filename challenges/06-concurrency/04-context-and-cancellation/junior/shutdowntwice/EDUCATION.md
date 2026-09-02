# Idempotent Shutdown

## Intuition

`close()` on an already-closed channel panics, so a naive cancel implementation would be a landmine. `context` guards `Done()` behind a mutex and a nil check, making every cancel func idempotent by contract. That is precisely what lets you both `defer cancel()` *and* call `cancel()` early on an error path.

## Approach

1. Create the shutdown context.
2. Call `cancel()` twice.
3. `<-ctx.Done()` and return `ctx.Err()`.

## Solution

```go
// ShutdownTwice models an operator sending SIGTERM and then, impatiently,
// SIGINT: the signal handler calls the shutdown cancel func twice. It returns
// the shutdown context's error afterwards.
//
// Examples:
//
//	ShutdownTwice()                              => context.Canceled
//	calling cancel twice must not panic          => true
//	the second call must not change the reason   => true
func ShutdownTwice() error {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	cancel()
	<-ctx.Done()
	return ctx.Err()
}
```

## Walkthrough

- The first `cancel()` closes `Done()` and stores `context.Canceled`.
- The second `cancel()` sees the context already finished and returns immediately without touching the channel or the error.
- `Err()` still reports the reason from the first call.

## Pitfalls

- Do not add your own `sync.Once` around a cancel func; the contract already guarantees idempotence.
- A later cancel never overwrites an earlier deadline error — whichever reason latched first wins forever.
- Idempotent does not mean re-usable: a cancelled context can never be revived.
