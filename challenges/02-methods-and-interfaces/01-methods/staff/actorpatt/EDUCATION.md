# Actor Pattern

## Intuition

"Do not communicate by sharing memory; share memory by communicating." An actor
is that slogan made concrete: the counter lives inside one goroutine, and the
only way to affect it is to mail in a function that will be run *by* that
goroutine. Because `run` executes messages one at a time, every operation is
automatically atomic with respect to every other.

## Approach

1. Build a closure that captures `n`.
2. Have it mutate through the pointer it is handed.
3. Send it on the mailbox channel — `Add` returns as soon as the send completes.

## Solution

```go
func (a *CounterActor) Add(n int) {
	a.msgs <- func(c *int) {
		*c += n
	}
}
```

## Walkthrough

100 goroutines call `Add(1)`. Each send hands a closure to the buffered channel;
`run` receives them in FIFO order and calls `fn(&a.count)`, so the increments
happen strictly one after another on a single goroutine.

`wg.Wait()` only proves the 100 *sends* finished — some closures may still be
queued. The test then sends one more message that reports the value back over
`done`. Because the mailbox is FIFO, that message runs after all 100 increments,
so the read is both synchronized and correctly ordered.

## Pitfalls

- **`a.count += n` in `Add`.** The obvious version, and a straight data race —
  `go test -race` reports concurrent writes.
- **Capturing a loop variable.** Not an issue since Go 1.22 (per-iteration `i`),
  but the closure must capture `n`, not read shared state.
- **An unbuffered mailbox.** Correct, just slower: each `Add` blocks until `run`
  picks the message up.
- **Assuming `Add` is synchronous.** It is not — the counter may not have moved
  when `Add` returns. That is why the test drains with a round-trip message.

## Why the read must go through the mailbox

Reading `a.count` directly after `wg.Wait()` is a race *and* likely wrong: the
value may be short by however many messages are still queued. Sending a query
message is the actor-model way to read state.
