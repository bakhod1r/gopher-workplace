# Draining a Channel

## Intuition

Ranging over a channel is a receive loop with an automatic exit condition. It
is also the tool that lets a producer finish: every blocked send needs a
receive, and the drain supplies them all.

## Approach

1. Initialise `count := 0`.
2. `for range msgs { count++ }` — the loop ends when the queue closes.
3. Return `count`.

## Solution

```go
func DrainDeadLetters(msgs <-chan string) int {
	count := 0
	for range msgs {
		count++
	}
	return count
}
```

## Walkthrough

The producer sends `a` and blocks until the drain receives it, and so on for
`b` and `c`. After the third message its loop ends, it closes the channel, and
the drain's range terminates with `count == 3`.

## Pitfalls

- Stopping early (say after a maximum) leaves the producer blocked on its next send — a goroutine leak.
- Using `<-msgs` in a bare loop without the comma-ok form: a closed channel yields zero values forever.
- Expecting the count to include messages sent after the close — sending on a closed channel panics instead.
