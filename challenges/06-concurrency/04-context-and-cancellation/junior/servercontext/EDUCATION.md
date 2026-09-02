# Server Root Context

## Intuition

`context.Background()` is the empty context. It is never cancelled, has no deadline and holds no values. It exists so that every chain in the process has one root to hang cancellation and values off. `main`, `init` and tests are the only places that should create it.

## Approach

1. Return `context.Background()`.

## Solution

```go
// ServerContext returns the root context that main creates once at process
// start and passes to the HTTP server, the database pool and the metrics
// exporter. Nothing above it exists, so it is never cancelled and carries no
// deadline and no values.
//
// Examples:
//
//	ServerContext() != nil     => true
//	ServerContext().Err()      => nil
//	ServerContext().Done()     => nil (never closed)
func ServerContext() context.Context {
	return context.Background()
}
```

## Walkthrough

- `context.Background()` returns a package-level singleton value, so every caller shares the same root.
- Its `Done()` returns a nil channel, so a `select` on it blocks forever — exactly what "never cancelled" means.
- `Err()` is nil because nothing can ever cancel it.

## Pitfalls

- Do not create a fresh cancellable context here — the process root must not be cancellable by a single handler.
- `Done() == nil` is not a bug. Receiving from a nil channel blocks forever, which is correct for a context that never finishes.
