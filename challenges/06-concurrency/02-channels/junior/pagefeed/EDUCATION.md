# A Goroutine Generator

## Intuition

Pair a goroutine with an unbuffered channel and you get a lazy sequence:
the producer runs exactly one step ahead of the consumer, blocking on each
send until the consumer takes the value. No buffer, no memory growth, and
`close` at the end terminates the consumer's `range`.

## Approach

1. `pages := make(chan int)` — unbuffered.
2. Start a goroutine that sends `1..n` and then `close(pages)`.
3. In the exporter, `range` over `pages` appending each number.
4. Return the collected slice.

## Solution

```go
func StreamPages(n int) []int {
	pages := make(chan int)
	go func() {
		for page := 1; page <= n; page++ {
			pages <- page
		}
		close(pages)
	}()

	out := []int{}
	for page := range pages {
		out = append(out, page)
	}
	return out
}
```

## Walkthrough

For `n = 3`: the goroutine blocks on `pages <- 1` until the range receives
it, then `2`, then `3`, then closes. The range sees the close and exits with
`[1 2 3]`. For `n = 0` the goroutine closes immediately and the loop never
runs.

## Pitfalls

- Closing outside the goroutine (right after `go func`) closes while sends are still coming — panic.
- Forgetting `close` leaves the `range` blocked and the export deadlocks.
- Sending without a goroutine on an unbuffered channel deadlocks on the first send.
