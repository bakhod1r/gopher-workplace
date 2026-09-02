# Generators and Closing

## Intuition

A *generator* is a function that returns a channel it fills itself. The
caller only receives. Because no indexer is receiving yet when
`StreamShardIDs` runs, an unbuffered channel would block on the first
send — so buffer it with room for every id.

## Approach

1. Clamp `n` to 0 so `make` never sees a negative capacity.
2. Create `make(chan int, n)`.
3. Send `0 .. n-1`; the buffer absorbs them all.
4. `close(ch)` so the indexers' `range` terminates.
5. Return the channel — the `<-chan int` result type makes it receive-only.

## Solution

```go
func StreamShardIDs(n int) <-chan int {
	if n < 0 {
		n = 0
	}
	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
```

## Walkthrough

For `StreamShardIDs(3)`: the buffer holds 3 slots, `0`, `1`, `2` go in
without blocking, `close` marks the end, and `for id := range ch` yields 3
ids then stops.

## Pitfalls

- Using an **unbuffered** channel without a goroutine deadlocks on the first send.
- Forgetting `close` makes `range` block forever after the last id.
- `make(chan int, -1)` panics — clamp negative input.
