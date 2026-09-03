# Trace ID Carrier

## Intuition

A context value is a chain of small immutable frames. Looking a key up walks up the chain until a frame matches — and matching is by *type as well as value*, which is exactly why an unexported struct type is collision-proof.

## Approach

1. `WithTrace`: return the parent when the ID is empty; otherwise `context.WithValue(ctx, traceKey{}, id)`.
2. `TraceID`: comma-ok assert `ctx.Value(traceKey{})` to `string` and return it.
3. `Chain`: capture `TraceID(ctx)` first, tag, then read the tagged context twice.

## Solution

```go
func WithTrace(ctx context.Context, id string) context.Context {
	if id == "" {
		return ctx
	}
	return context.WithValue(ctx, traceKey{}, id)
}

func TraceID(ctx context.Context) string {
	id, _ := ctx.Value(traceKey{}).(string)
	return id
}

func Chain(ctx context.Context, id string) []string {
	before := TraceID(ctx)
	tagged := WithTrace(ctx, id)
	return []string{before, TraceID(tagged), TraceID(tagged)}
}
```

## Walkthrough

`WithTrace(bg, "abc")` pushes a frame `{traceKey{}: "abc"}`. A later `context.WithCancel` on it adds a cancel frame but keeps the parent chain, so `TraceID` still finds the value. A `context.WithValue(bg, "traceKey", ...)` from elsewhere stores a `string` key, which never equals `traceKey{}`, so the lookup misses and returns `""`.

## Pitfalls

- Using a `string` or `int` key — any package can produce the same value and silently overwrite yours.
- A one-value assertion `ctx.Value(k).(string)` panics when the key is missing.
- Passing optional configuration through the context instead of through parameters.
