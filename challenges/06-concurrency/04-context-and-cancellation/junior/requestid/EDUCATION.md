# Request ID for Logging

## Intuition

`ctx.Value` gives back `any`, so two things can go wrong: the key is missing (`nil`) or someone stored the wrong type. The comma-ok assertion collapses both into `ok == false`. Adding the empty-string check on top gives the logger one predictable placeholder for every failure mode, and no code path where logging itself can crash the process.

## Approach

1. `WithRequestID`: wrap with `context.WithValue`.
2. `RequestID`: assert with comma-ok.
3. Return `"unknown"` when `!ok` or the string is empty; otherwise return the ID.

## Solution

```go
// WithRequestID returns a copy of ctx carrying the request ID that the edge
// proxy assigned to this request.
//
// Examples:
//
//	RequestID(WithRequestID(bg, "req-8f21")) => "req-8f21"
//	RequestID(context.Background())          => "unknown"
//	RequestID(WithRequestID(bg, ""))         => "unknown"
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

// RequestID returns the request ID the logger should stamp on every line, or
// "unknown" when the context carries no usable ID.
func RequestID(ctx context.Context) string {
	id, ok := ctx.Value(requestIDKey{}).(string)
	if !ok || id == "" {
		return "unknown"
	}
	return id
}
```

## Walkthrough

- `WithRequestID(bg, "req-8f21")` stores the string; the lookup finds it and returns it.
- On a bare `context.Background()` the lookup returns `nil`, the assertion fails, and the placeholder is used.
- When an `int` was stored under the key by mistake, the assertion also fails — same placeholder, no panic in the logging path.

## Pitfalls

- A one-value assertion panics on the wrong-type case; the test covers it deliberately.
- An empty ID from a misconfigured proxy is a data bug, not a crash — normalise it rather than writing empty brackets into every log line.
