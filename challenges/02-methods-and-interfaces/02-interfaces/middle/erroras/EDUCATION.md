# errors.As

## Intuition

`errors.Is` answers "is this that specific error?"; `errors.As` answers "is there an error of this *type* in the chain, and can I have it?" — which is what you need when the payload matters.

## Approach

1. `Error` formats the status with `fmt.Sprintf`.
2. `Call` returns `nil` for 200 and wraps `&HTTPError{Status: status}` otherwise.
3. `StatusOf` declares `var he *HTTPError` and passes `&he` to `errors.As`.
4. `Retryable` is `StatusOf(err) >= 500`, which handles nil for free.

## Solution

```go
func (e *HTTPError) Error() string { return fmt.Sprintf("http %d", e.Status) }

func Call(status int) error {
	if status == 200 {
		return nil
	}
	return fmt.Errorf("call: %w", &HTTPError{Status: status})
}

func StatusOf(err error) int {
	var he *HTTPError
	if errors.As(err, &he) {
		return he.Status
	}
	return 0
}

func Retryable(err error) bool { return StatusOf(err) >= 500 }
```

## Walkthrough

`fmt.Errorf("outer: %w", Call(503))` nests two wrappers. `errors.As` walks down through both and binds the `*HTTPError`, so `StatusOf` still reports 503.

## Pitfalls

- Passing `he` instead of `&he` to `errors.As` — it panics because the target must be a pointer.
- Type-asserting `err.(*HTTPError)` instead, which fails as soon as the error is wrapped.
- Using `errors.Is` here: it compares identity, not type, so it cannot retrieve the status.
