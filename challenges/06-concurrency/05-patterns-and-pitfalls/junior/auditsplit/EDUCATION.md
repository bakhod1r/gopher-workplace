# Teeing a Channel

## Intuition

A send on a `nil` channel blocks forever, which sounds useless until you put
it in a `select`: a `nil` case can never be chosen, so it is effectively
switched off. That is how one select loop delivers exactly one copy to each
output, in whatever order the consumers happen to be ready.

## Approach

1. Create `archive` and `alerts`, and start a goroutine with a deferred close for each.
2. For every event, set `a, b := archive, alerts` and loop twice over a `select` on `a <- ev` and `b <- ev`.
3. Set the chosen channel to `nil` after its send so the second iteration must pick the other.

## Solution

```go
func TeeAudit(events <-chan string) (<-chan string, <-chan string) {
	archive := make(chan string)
	alerts := make(chan string)

	go func() {
		defer close(archive)
		defer close(alerts)
		for ev := range events {
			a, b := archive, alerts
			for i := 0; i < 2; i++ {
				select {
				case a <- ev:
					a = nil
				case b <- ev:
					b = nil
				}
			}
		}
	}()

	return archive, alerts
}
```

## Walkthrough

For event `login`: whichever consumer is ready first wins the select — say the
archive. `a` becomes `nil`, so on the second iteration only `b <- ev` is
possible, and the alerting consumer gets its copy. Then the loop moves to the
next event.

## Pitfalls

- Sending sequentially (`archive <- ev` then `alerts <- ev`) works only if the archive consumer never lags; the select version tolerates either order.
- Forgetting to nil out the chosen case — the select could then send both copies to the same channel.
- Reading only one of the two outputs in the caller: the tee blocks and both consumers stall.
