# Receiving After Close

## Intuition

`close` is a marker, not a delete. Buffered frames stay and are delivered
first. Only when the buffer empties does the channel start returning the
zero value — instantly, forever, and never blocking. That is precisely why
the comma-ok form exists.

## Approach

1. Clamp `extra` to 0.
2. Buffer all of `sizes` on a channel of capacity `len(sizes)`.
3. `close` it.
4. Receive `len(sizes)+extra` times, appending each result.
5. Return the slice.

## Solution

```go
func ReadFrames(sizes []int, extra int) []int {
	if extra < 0 {
		extra = 0
	}
	frames := make(chan int, len(sizes))
	for _, size := range sizes {
		frames <- size
	}
	close(frames)

	got := []int{}
	for i := 0; i < len(sizes)+extra; i++ {
		got = append(got, <-frames)
	}
	return got
}
```

## Walkthrough

For `[]int{1024, 512}` with `extra = 1`: receives yield `1024`, then `512`
from the buffer, then `0` because the closed channel is drained —
`[1024 512 0]`.

## Pitfalls

- Assuming `close` throws away buffered frames — it does not.
- Looping on bare receives from a closed channel spins at full speed forever.
- Sending on a closed channel panics; only receiving is safe.
