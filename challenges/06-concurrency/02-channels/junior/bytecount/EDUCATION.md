# Ranging Over a Channel

## Intuition

`for v := range ch` is the idiomatic drain loop. Each iteration performs a
receive; the loop exits when the channel is closed *and* empty. It is the
receive-side mirror of `close`.

## Approach

1. Start `total` at `0`.
2. `range` over `sizes`, adding each value.
3. Return `total` after the loop.

## Solution

```go
func TotalBytes(sizes <-chan int) int {
	total := 0
	for n := range sizes {
		total += n
	}
	return total
}
```

## Walkthrough

Given sizes `1200`, `800` then a close: the loop yields 1200 (total 1200)
and 800 (total 2000), then sees the closed channel and exits, returning
`2000`.

## Pitfalls

- If the edge never closes the channel, `range` blocks forever and the window never reports.
- `for v := range ch` gives the *value*, not an index — there is no `i, v` form for channels.
- Do not `close(sizes)` in the accountant; closing is the sender's job.
