# First Row or Disconnect

## Intuition

A bare `<-rows` is an unbounded wait. Putting it in a `select` next to `<-ctx.Done()` bounds it by the request's lifetime: whichever event happens first decides. This two-case select is the single most common context shape in production Go.

## Approach

1. `select` with two cases.
2. `<-ctx.Done()` → return the zero value and `ctx.Err()`.
3. `row := <-rows` → return the row and nil.

## Solution

```go
// FirstRow returns the first row streamed by the database, or aborts if the
// request context finishes first because the client disconnected or the
// request budget expired.
//
// On abort it returns the zero row and ctx.Err().
//
// Examples:
//
//	FirstRow(live ctx, chan with "alice")   => "alice", nil
//	FirstRow(cancelled ctx, empty chan)     => "", context.Canceled
//	FirstRow(expired ctx, empty chan)       => "", context.DeadlineExceeded
func FirstRow(ctx context.Context, rows <-chan string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case row := <-rows:
		return row, nil
	}
}
```

## Walkthrough

- With a live context and a buffered channel, only the row case is ready, so the row is returned.
- With a cancelled context and an empty channel, only the done case is ready, so `ctx.Err()` comes back.
- If both were ready, Go picks uniformly at random — which is why the tests never make both ready at once.

## Pitfalls

- Checking `if ctx.Err() != nil` before a bare receive is not enough: the context can finish while you are blocked in the receive.
- Returning a hand-made error instead of `ctx.Err()` loses the distinction between a disconnect and a timeout.
- Do not add a `default:` case — that turns the blocking wait into a busy poll.
