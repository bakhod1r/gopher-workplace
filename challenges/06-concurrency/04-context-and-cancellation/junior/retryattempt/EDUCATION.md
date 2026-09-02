# Overriding a Value per Attempt

## Intuition

Contexts are immutable: `WithValue` never changes its parent, it wraps it. `Value` walks from the outermost wrapper inward and stops at the first match, so re-tagging for attempt 2 shadows attempt 1 without the retry loop having to unset anything — and the original request context, which other goroutines may still hold, is unaffected.

## Approach

1. `WithAttempt`: return `context.WithValue(ctx, attemptKey{}, n)`.
2. `Attempt`: assert the lookup to `int`, discarding `ok`, and return it.

## Solution

```go
// WithAttempt returns a copy of ctx tagged with the retry attempt number, so
// the HTTP client can stamp it on the outbound request and the logger can show
// which try failed.
//
// Examples:
//
//	Attempt(WithAttempt(bg, 2))                => 2
//	Attempt(context.Background())              => 0
//	Attempt(WithAttempt(WithAttempt(bg, 1), 2)) => 2
func WithAttempt(ctx context.Context, n int) context.Context {
	return context.WithValue(ctx, attemptKey{}, n)
}

// Attempt reports the retry attempt recorded on ctx, or 0 when none was set.
func Attempt(ctx context.Context) int {
	n, _ := ctx.Value(attemptKey{}).(int)
	return n
}
```

## Walkthrough

- `WithAttempt(bg, 1)` wraps the background context with `attemptKey{} → 1`.
- `WithAttempt(that, 2)` wraps again; the lookup finds `2` first and stops, so the inner `1` is never reached.
- On a context with no attempt the lookup returns `nil`; asserting to `int` yields `0`, the sensible "first try" default.
- A value stored as a `string` fails the assertion and also yields `0`.

## Pitfalls

- The parent context still reports its old attempt — that is by design, not a bug.
- The zero value is doing double duty here as "unset"; when you must tell them apart, keep the `ok` flag.
- Do not use context values to pass a retry budget the callee must honour — that belongs in a real parameter.
