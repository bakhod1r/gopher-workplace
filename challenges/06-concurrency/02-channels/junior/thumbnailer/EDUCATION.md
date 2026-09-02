# Channels Preserve Order

## Intuition

A channel is a FIFO queue: values come out in the order they went in. That
makes it safe to build an ordered result from a single producer. Doing both
halves in one goroutine works *only* if the buffer holds everything, since
nothing is receiving while you send.

## Approach

1. Make a channel with capacity `len(sides)`.
2. Send `side * side` for each request.
3. `close` the channel.
4. `range` over it, appending into a non-nil slice.
5. Return the slice.

## Solution

```go
func ThumbAreas(sides []int) []int {
	areas := make(chan int, len(sides))
	for _, side := range sides {
		areas <- side * side
	}
	close(areas)

	out := []int{}
	for area := range areas {
		out = append(out, area)
	}
	return out
}
```

## Walkthrough

For `[64, 128]`: sends `4096` and `16384` into the buffer, closes, then the
range yields them in that same order — `[4096 16384]`.

## Pitfalls

- An unbuffered channel here deadlocks: the first send waits for a receiver that only runs later.
- Ranging before closing blocks forever once the buffer is drained.
- `var out []int` returns `nil` for an empty request list, failing a `DeepEqual` against `[]int{}`.
