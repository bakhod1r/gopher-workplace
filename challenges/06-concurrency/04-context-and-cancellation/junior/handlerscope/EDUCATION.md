# Per-Handler Context Scope

## Intuition

`cancel` is not only an abort button. It is the *release* half of an acquire/release pair. Deferring it tears the request context down on the success path, the error path and the panic path alike, so a branch added later cannot forget it — and every goroutine the handler spawned sees `Done()` close.

## Approach

1. `ctx, cancel := context.WithCancel(context.Background())`.
2. `defer cancel()` on the very next line.
3. Return `handler(ctx)`.

## Solution

```go
// ServeRequest derives a per-request cancellable context, runs the handler with
// it, and guarantees the context is cancelled once the response has been
// written — so background lookups started by the handler are torn down instead
// of leaking for the life of the process.
//
// ServeRequest returns whatever the handler returns.
//
// Examples:
//
//	ServeRequest(func(ctx context.Context) error { return nil })     => nil
//	ServeRequest(func(ctx context.Context) error { return errDB })   => errDB
//	after ServeRequest returns, the handler's ctx has Err() == context.Canceled
func ServeRequest(handler func(ctx context.Context) error) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return handler(ctx)
}
```

## Walkthrough

- The handler sees a live context: `Done()` is an open channel and `Err()` is nil.
- When `ServeRequest` returns, the deferred `cancel()` fires.
- That closes `Done()` and sets `Err()` to `context.Canceled`, which the test observes on the captured context.

## Pitfalls

- Calling `cancel()` before `handler(ctx)` kills the context the handler was supposed to use.
- `go vet`'s `lostcancel` check flags a `WithCancel` whose cancel func is never called — never silence it.
- Put `defer cancel()` immediately after the `WithCancel` call so an early return added later cannot skip it.
