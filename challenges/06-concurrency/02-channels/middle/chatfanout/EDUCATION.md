# One Producer, Every Subscriber

## Intuition

Fan-out is the mirror of fan-in: instead of many senders sharing one channel,
one sender owns N channels. That makes close ownership easy — the broadcaster
created every output, so the broadcaster closes every output, once, when its
input is exhausted.

The subtle part is that `FanOutRoom` must *return before broadcasting*. The
outputs are unbuffered, so the first send only completes once a client is
receiving. If you tried to broadcast on the calling goroutine you would block
forever, holding channels nobody has been given yet.

## Approach

1. Allocate `subscribers` bidirectional channels, plus a parallel `[]<-chan string` view.
2. Start one broadcaster goroutine.
3. In it, `for msg := range messages`, and for each subscriber `out <- msg`.
4. After the range ends, close every output.
5. Return the receive-only view.

## Solution

```go
// FanOutRoom broadcasts every room message to each of the subscribers and
// returns one receive-only channel per subscriber. Every subscriber sees the
// full message history in publication order, and each subscriber channel is
// closed once the room feed closes so a client goroutine can range over it.
//
// The returned channels are unbuffered: a slow subscriber applies backpressure
// to the whole room, which is the intended fairness rule here.
//
// Examples:
//
//	FanOutRoom(chan hi,bye, 2) => 2 channels, each yielding hi then bye, then closed
//	FanOutRoom(chan hi, 1)     => 1 channel yielding hi, then closed
//	FanOutRoom(chan hi, 0)     => no channels; the feed is still drained
func FanOutRoom(messages <-chan string, subscribers int) []<-chan string {
	outs := make([]chan string, subscribers)
	view := make([]<-chan string, subscribers)
	for i := range outs {
		outs[i] = make(chan string)
		view[i] = outs[i]
	}

	go func() {
		for msg := range messages {
			for _, out := range outs {
				out <- msg
			}
		}
		for _, out := range outs {
			close(out)
		}
	}()

	return view
}
```

## Walkthrough

- With two subscribers the broadcaster sends `hi` to subscriber 0, waits for it
  to be received, then sends `hi` to subscriber 1 — so order is preserved per
  subscriber even though delivery is staggered.
- When the room feed closes, `range` ends and the close loop runs, ending both
  client `range` loops.
- With `subscribers == 0` the inner loop does nothing, but the outer `range`
  still drains the feed to completion — no blocked publisher, no leak.
- With an empty room feed, the broadcaster goes straight to the close loop, so
  every client sees a closed channel on its first receive.

## Pitfalls

- Broadcasting on the caller's goroutine deadlocks on the first unbuffered send.
- Returning `[]chan string` instead of `[]<-chan string` invites a client to
  send or close on a channel it does not own.
- Closing inside the message loop closes the outputs after the first message,
  and the second message panics with "send on closed channel".
- Skipping the drain when there are no subscribers strands whoever is filling
  the room feed.
