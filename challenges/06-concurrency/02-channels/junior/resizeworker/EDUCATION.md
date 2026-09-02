# Request and Reply Channels

## Intuition

One channel gives you a one-way handoff; two give you a round trip. The
handler's `req <- width` blocks until the worker takes it, and the
handler's `<-reply` blocks until the worker answers — so the whole exchange
behaves like a synchronous function call implemented with concurrency.

## Approach

1. Make `req` and `reply`, both unbuffered.
2. Start a goroutine: receive from `req`, send `w * 2` on `reply`.
3. Send `width` on `req`.
4. Return the value received from `reply`.

## Solution

```go
func ScaleRequest(width int) int {
	req := make(chan int)
	reply := make(chan int)

	go func() {
		w := <-req
		reply <- w * 2
	}()

	req <- width
	return <-reply
}
```

## Walkthrough

For `ScaleRequest(640)`: the worker parks on `<-req`; the handler's send
hands over `640` and both continue. The worker sends `1280` on `reply` and
the handler receives it.

## Pitfalls

- Sending on `req` before starting the worker deadlocks — no receiver exists yet.
- Reusing one channel for both directions lets the handler receive its own request.
- Closing either channel is unnecessary here and risks a send-on-closed panic.
