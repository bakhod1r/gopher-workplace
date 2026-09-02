# Cancellable Export Pipeline

## Intuition

Cancelling only the consumer solves half the problem: the producer would still be parked on a send into a channel nobody reads, and that goroutine — plus everything it references — lives until the process dies. Wrapping the send in a `select` with `<-ctx.Done()` gives the producer its own exit, which is the discipline that keeps a long-running service's goroutine count flat.

## Approach

1. Create a buffered channel sized to `len(ids)`.
2. Start the producer: for each ID, `select` on `ctx.Done()` and the send; `defer close(out)`.
3. In the consumer, loop on a `select` over `ctx.Done()` and a comma-ok receive.
4. Closed channel → return the rows and nil; done → return nil and `ctx.Err()`.

## Solution

```go
// ExportRows renders each account ID into a CSV line, streaming the work
// through a goroutine so the consumer can abandon the export the moment the
// request context finishes.
//
// On cancellation it returns nil and ctx.Err(); otherwise the full slice.
//
// Examples:
//
//	ExportRows(live ctx, []int{1, 2})  => ["row-1", "row-2"], nil
//	ExportRows(live ctx, nil)          => [], nil
//	ExportRows(cancelled ctx, []int{1}) => nil, context.Canceled
func ExportRows(ctx context.Context, ids []int) ([]string, error) {
	out := make(chan string, len(ids))

	go func() {
		defer close(out)
		for _, id := range ids {
			select {
			case <-ctx.Done():
				return
			case out <- fmt.Sprintf("row-%d", id):
			}
		}
	}()

	rows := make([]string, 0, len(ids))
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case row, ok := <-out:
			if !ok {
				return rows, nil
			}
			rows = append(rows, row)
		}
	}
}
```

## Walkthrough

- With a live context the producer sends every row and closes the channel; the consumer drains them and sees `ok == false`.
- With no IDs the producer closes immediately and the consumer returns an empty, non-nil slice.
- With an already-cancelled context the consumer returns `ctx.Err()` at once, and the producer's own select lets it return rather than block.

## Pitfalls

- Closing the channel from the consumer side panics the producer's next send; only the producer closes.
- A bare `out <- row` in the producer leaks the goroutine whenever the consumer walks away.
- `make([]string, 0, n)` returns an empty non-nil slice, which `reflect.DeepEqual` distinguishes from `nil`.
- Never send `ctx.Err()` down the data channel; return it from the function instead.
