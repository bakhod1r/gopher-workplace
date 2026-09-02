# Pub/Sub

## Intuition

A channel per subscriber decouples in two directions: the publisher does not
know who listens, and the listener does not have to be running when the message
is sent. The buffer is what buys that second property.

The shared structure is the topic map, and it has one writer path
(`Subscribe`) and many reader paths (`Publish`) — the exact shape `RWMutex` is
built for.

## Approach

1. Acquire the read lock; defer its release.
2. Range the slice for the topic.
3. Send the message on each channel.

## Solution

```go
func (ps *PubSub) Publish(topic, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}
```

## Walkthrough

Two `Subscribe("news")` calls each allocate a buffered channel and append it
under the key. `Publish` reads that slice under `RLock` and sends `"hello"` to
both. Because each channel has room for 10 messages, the sends complete
immediately and the test's receives find the values already queued.

A missing topic yields a nil slice, so the loop simply does not run.

## Pitfalls

- **Using `Lock` instead of `RLock`.** Correct, but it makes concurrent
  publishers wait on each other for no reason.
- **Unbuffered channels.** A send would block until a receiver arrived —
  while holding the lock — and `Subscribe` would deadlock behind it.
- **Sending after `RUnlock`.** The slice could be reallocated by a concurrent
  `Subscribe`; do the sends inside the critical section, or copy the slice
  first.
- **Blocking forever on a full buffer.** A slow subscriber stalls every
  publisher. Production code uses a `select` with a `default` to drop.

## Why the deferred unlock matters

If a channel is closed elsewhere, `ch <- msg` panics. Without `defer`, the read
lock would be held forever and the whole bus would freeze; with it, the panic
unwinds cleanly.
