# Revenue Report Export

## Intuition

`range ch` is the wrong tool the moment cancellation matters: it can only stop when the channel closes. Rebuilding the loop as `for { select { ... } }` gives every iteration a second exit, so an abandoned export stops within one row instead of at the end of the stream.

## Approach

1. Start `total` at 0.
2. Loop forever with a two-case `select`.
3. Done → return `total, ctx.Err()`.
4. Receive with comma-ok → `!ok` returns `total, nil`; otherwise add and keep looping.

## Solution

```go
// TotalRevenue sums the cent amounts streamed by the report query until the
// stream closes, and returns the total. If the user cancels the export or the
// export budget expires first, it stops immediately and returns the partial
// total together with ctx.Err().
//
// Examples:
//
//	TotalRevenue(live ctx, closed chan 100,250)  => 350, nil
//	TotalRevenue(live ctx, closed empty chan)    => 0, nil
//	TotalRevenue(cancelled ctx, empty chan)      => 0, context.Canceled
func TotalRevenue(ctx context.Context, rows <-chan int) (int, error) {
	total := 0
	for {
		select {
		case <-ctx.Done():
			return total, ctx.Err()
		case amount, ok := <-rows:
			if !ok {
				return total, nil
			}
			total += amount
		}
	}
}
```

## Walkthrough

- With a live context and a closed, pre-filled channel, the loop drains every buffered amount and then sees `ok == false` and returns the sum.
- With a cancelled context and an empty open channel, only the done case is ready on the first iteration, so nothing is summed and `context.Canceled` is returned.
- The tests never make both cases ready at once, because `select` would then choose at random.

## Pitfalls

- `for amount := range rows` cannot be cancelled — the consumer runs until the producer closes.
- Returning `0` instead of the partial total on cancellation throws away work the caller may want to log.
- Adding a `default:` case turns the loop into a spin that burns a whole core.
