# A Generator That Can Be Cancelled

## Intuition

An unbounded producer and a bounded consumer only work together if the
producer is told when to stop. Putting the send inside a `select` alongside
`done` turns "blocked forever" into "blocked until either the consumer reads
or the subscription ends".

## Approach

1. Create `out` and start a goroutine with `defer close(out)`.
2. Loop from `base` upwards with no termination condition.
3. `select` on `out <- price` and `<-done`, returning on the latter.

## Solution

```go
func PriceFeed(done <-chan struct{}, base int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for price := base; ; price++ {
			select {
			case out <- price:
			case <-done:
				return
			}
		}
	}()
	return out
}
```

## Walkthrough

The consumer takes three quotes, so the goroutine completes three sends and is
blocked in the select on the fourth. Closing `done` makes the second case
ready, the goroutine returns, the deferred close fires, and the test's drain
loop ends immediately.

## Pitfalls

- A bare `out <- price` with a separate `done` check: the goroutine blocks on the send and never reaches the check.
- Returning without closing `out` — a consumer still ranging over it hangs forever.
- Buffering the channel to "avoid the block": it delays the leak by exactly the buffer size instead of preventing it.
