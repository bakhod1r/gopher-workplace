# Refuse to Charge a Dead Request

## Intuition

`Err()` is a cheap, non-blocking snapshot of whether the context is still alive — the right tool at the top of a function that is about to do something expensive or irreversible. It is not a substitute for `select` on `<-ctx.Done()` during a long operation, but as an entry guard it stops obviously wasted, and here obviously harmful, work.

## Approach

1. Read `ctx.Err()` into a local.
2. If it is non-nil, return the zero value and that error.
3. Otherwise call `capture()` and return its result with a nil error.

## Solution

```go
// Charge runs the card capture only if the request context is still alive.
// Capturing a payment for a caller that has already gone away bills a customer
// nobody will ever show a receipt to, so the guard runs before the side effect,
// not after.
//
// Examples:
//
//	Charge(live ctx, ok)        => "captured", nil
//	Charge(cancelled ctx, ok)   => "", context.Canceled  (capture never runs)
//	Charge(expired ctx, ok)     => "", context.DeadlineExceeded
func Charge(ctx context.Context, capture func() string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	return capture(), nil
}
```

## Walkthrough

- A live context reports `Err() == nil`, so the capture runs and its receipt comes back.
- A cancelled context reports `context.Canceled`, so the function returns before the capture — the test asserts the callback was never invoked.
- An expired context is treated identically, but the caller can tell the two apart from the returned sentinel.

## Pitfalls

- Calling `capture()` first and checking afterwards charges the card anyway — the damage is already done.
- Inventing a new error hides whether the customer hung up or the service was too slow.
- `ctx.Err()` is a point-in-time check; for work that takes a while, also select on `<-ctx.Done()` inside the loop.
