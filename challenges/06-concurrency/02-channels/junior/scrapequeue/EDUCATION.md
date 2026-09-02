# len and cap on Channels

## Intuition

A buffered channel is a fixed-size queue. `cap` is the size you asked for
and never changes; `len` is how many values are sitting in it right now. A
send blocks only when `len == cap`; a receive blocks only when `len == 0`.

## Approach

1. Clamp `size` to 0.
2. `queue := make(chan int, size)`.
3. Send `size` samples — the buffer takes all of them.
4. Return `len(queue), cap(queue)`.

## Solution

```go
func QueueStats(size int) (int, int) {
	if size < 0 {
		size = 0
	}
	queue := make(chan int, size)
	for i := 0; i < size; i++ {
		queue <- i
	}
	return len(queue), cap(queue)
}
```

## Walkthrough

For `QueueStats(3)`: capacity 3, three sends fill it to `len == 3`, and the
agent reports `3, 3`. A fourth send would block forever here, since no
scraper is draining yet.

## Pitfalls

- `make(chan int)` (no size) has `cap == 0`; the very first send blocks.
- `make(chan int, -1)` panics at runtime.
- `len`/`cap` are snapshots — with a live scraper they are stale the instant you read them.
