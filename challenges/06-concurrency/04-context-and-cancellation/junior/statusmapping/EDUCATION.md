# Mapping Context Errors to Status

## Intuition

By the time an error reaches the top of a handler it has usually been wrapped with `%w` for context. `==` only matches the outermost error, so it silently misclassifies wrapped cancellations as internal errors — and a burst of client disconnects would then look like an outage. `errors.Is` walks the whole chain, which is why it is the only correct comparison for sentinels.

## Approach

1. Handle `err == nil` first.
2. `errors.Is(err, context.Canceled)` → client closed.
3. `errors.Is(err, context.DeadlineExceeded)` → gateway timeout.
4. Default → internal error.

## Solution

```go
// Status maps the error a handler finished with onto the label the access log
// and the metrics counter use.
//
//	nil                       => "ok"
//	context.Canceled          => "client_closed_request"
//	context.DeadlineExceeded  => "gateway_timeout"
//	anything else             => "internal_error"
//
// Wrapped errors count: an error wrapping context.DeadlineExceeded with %w
// still maps to "gateway_timeout".
//
// Examples:
//
//	Status(nil)                                        => "ok"
//	Status(context.Canceled)                           => "client_closed_request"
//	Status(fmt.Errorf("query: %w", context.DeadlineExceeded)) => "gateway_timeout"
func Status(err error) string {
	switch {
	case err == nil:
		return "ok"
	case errors.Is(err, context.Canceled):
		return "client_closed_request"
	case errors.Is(err, context.DeadlineExceeded):
		return "gateway_timeout"
	default:
		return "internal_error"
	}
}
```

## Walkthrough

- `fmt.Errorf("query: %w", context.DeadlineExceeded)` produces a wrapper whose `Unwrap()` returns the sentinel.
- `errors.Is` follows `Unwrap` repeatedly, so even a doubly wrapped error matches.
- An unrelated error unwraps to nothing that matches, so it falls through to the default.

## Pitfalls

- `err == context.Canceled` fails on every wrapped error — the test covers exactly that case.
- Order matters only in that `nil` must be checked first; `errors.Is(nil, target)` is false, not a panic.
- Never classify by matching the error's message text; the wording is not part of any API.
