# Propagating a Trace ID

## Intuition

A context value is a linked-list node: `WithValue` wraps the parent and `Value` walks up the chain until a key matches. Because the key is an unexported type only this package can name, a third-party library cannot read or spoof the trace ID by accident — which is why string keys are discouraged in exported APIs.

## Approach

1. Define an unexported key type (`struct{}` costs zero bytes).
2. `WithTraceID`: wrap with `context.WithValue`.
3. `TraceID`: look up the key and type-assert with comma-ok.

## Solution

```go
// WithTraceID returns a copy of ctx carrying the distributed-trace ID read
// from the incoming gRPC metadata.
//
// Examples:
//
//	TraceID(WithTraceID(context.Background(), "4bf92f"))  => "4bf92f", true
//	TraceID(context.Background())                         => "", false
//	TraceID(WithTraceID(WithTraceID(bg, "a"), "b"))       => "b", true
func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceKey{}, id)
}

// TraceID reports the trace ID carried by ctx, if any. Outbound clients call it
// to stamp the header on the next hop.
func TraceID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(traceKey{}).(string)
	return id, ok
}
```

## Walkthrough

- `WithTraceID(bg, "4bf92f")` returns a wrapper around `bg` holding `traceKey{} → "4bf92f"`.
- `TraceID` calls `Value(traceKey{})`, which walks the chain and finds the string.
- For a bare `context.Background()` the walk reaches the root and returns `nil`; asserting `nil` to `string` gives `("", false)`.
- Wrapping twice means the lookup hits the innermost wrapper first, so the most recent hop's ID wins.

## Pitfalls

- A plain `string` key can be written by any package — the test proves the foreign `"traceKey"` string key never reaches `TraceID`.
- A single-value assertion `ctx.Value(k).(string)` panics when the key is absent; always use comma-ok.
- Context values are for request-scoped metadata like trace IDs, not for required business arguments.
