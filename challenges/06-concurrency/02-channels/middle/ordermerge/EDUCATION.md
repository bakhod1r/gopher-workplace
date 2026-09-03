# Fan-In With a Single Closer

## Intuition

Fan-in is many senders, one channel. Go's rule is that the *sender* closes —
but with N senders there is no single sender to do it. The fix is to invent
one: a goroutine whose only job is to wait for every forwarder to finish and
then close. That goroutine is the sole owner of `close`.

Returning the channel before any work happens is the other half. The caller
ranges over the result; if `MergeOrderFeeds` drained the feeds itself first it
would need unbounded buffering and would stop being a stream.

## Approach

1. Make the merged channel.
2. `wg.Add(len(feeds))`, then start one forwarder goroutine per feed.
3. Each forwarder ranges its feed, sends onto `merged`, and calls `wg.Done`.
4. Start the closer goroutine: `wg.Wait()` then `close(merged)`.
5. Return `merged`.

## Solution

```go
// MergeOrderFeeds fans every regional order feed into one merged channel and
// returns it. One goroutine per feed forwards ids as they arrive; the merged
// channel is closed exactly once, after every feed has been drained, so the
// checkout reconciler can simply range over the result.
//
// Interleaving between feeds is not defined — only that every id arrives
// exactly once and that the merged channel is closed at the end.
//
// Examples:
//
//	MergeOrderFeeds(chan eu-1,eu-2 | chan us-1) => channel yielding 3 ids, then closed
//	MergeOrderFeeds(closed empty | chan ap-1)   => channel yielding ap-1, then closed
//	MergeOrderFeeds()                           => already closed, empty channel
func MergeOrderFeeds(feeds ...<-chan string) <-chan string {
	merged := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(feeds))
	for _, feed := range feeds {
		go func(feed <-chan string) {
			defer wg.Done()
			for id := range feed {
				merged <- id
			}
		}(feed)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
```

## Walkthrough

- With two feeds the counter starts at 2. Each forwarder decrements once its
  feed is closed and drained.
- The closer goroutine is parked in `wg.Wait()` the whole time, so `merged`
  stays open while any id can still arrive.
- With zero feeds the counter is already 0: `wg.Wait()` returns immediately and
  the caller's `range` sees a closed channel on the first receive.
- The forwarders block on `merged <- id` until the caller receives — that
  backpressure is what keeps the unbuffered merge from growing memory.

## Pitfalls

- Calling `close(merged)` inside a forwarder panics as soon as a second
  forwarder finishes, or sends on a closed channel.
- Calling `wg.Wait()` on the caller's goroutine deadlocks: the forwarders are
  blocked sending to a channel nobody is receiving from yet.
- Forgetting to pass `feed` into the goroutine was a classic bug before Go 1.22
  changed loop-variable scoping; passing it explicitly is still the clearest form.
- Buffering `merged` does not remove the need for the closer — it only delays
  the deadlock.
