# Draining a Notification Buffer

## Intuition

Cancellation is not failure — it is a stop signal, and everything collected before it is still valid work. Returning the partial batch alongside the error lets the caller persist or requeue those notifications instead of dropping them on the floor when a laptop lid closes.

## Approach

1. Start with an empty, non-nil slice.
2. Loop on a two-case `select`.
3. Done → return the slice and `ctx.Err()`.
4. Comma-ok receive → `!ok` returns the slice and nil; otherwise append.

## Solution

```go
// Drain collects every pending notification until the producer closes the
// channel, then returns them in arrival order. If the subscriber's context
// finishes first it returns what it has so far together with ctx.Err(), so the
// caller can requeue the undelivered remainder.
//
// Examples:
//
//	Drain(live ctx, closed chan "a","b")  => ["a", "b"], nil
//	Drain(live ctx, closed empty chan)    => [], nil
//	Drain(cancelled ctx, empty chan)      => [], context.Canceled
func Drain(ctx context.Context, ch <-chan string) ([]string, error) {
	got := make([]string, 0)
	for {
		select {
		case <-ctx.Done():
			return got, ctx.Err()
		case v, ok := <-ch:
			if !ok {
				return got, nil
			}
			got = append(got, v)
		}
	}
}
```

## Walkthrough

- With a live context and a closed, pre-filled channel the loop drains every value in order and then sees `ok == false`.
- With no values buffered the first receive already reports the close, returning an empty slice.
- With a cancelled context and an empty open channel the done case is the only ready one, so an empty slice and `context.Canceled` come back.

## Pitfalls

- Returning `nil` instead of the partial slice on cancellation loses notifications the caller could have requeued.
- `var got []string` starts as nil, which fails `reflect.DeepEqual(got, []string{})` — use `make`.
- `for v := range ch` cannot be cancelled; the subscriber would hang until the producer closes.
