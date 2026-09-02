# Filtering in a Goroutine

## Intuition

A filter stage sends fewer values than it receives, and sometimes none at
all. The consumer cannot know that in advance — which is why `close` must
be unconditional. Close is the only signal that says "there is nothing more
coming", regardless of how much came before.

## Approach

1. Make an unbuffered channel.
2. In a goroutine, iterate `seats` and send `seat` when `seat%2 == 0`.
3. After the loop, `close(aisle)` — always.
4. `range` in the allocator, appending seats.
5. Return the collected slice.

## Solution

```go
func AisleSeats(seats []int) []int {
	aisle := make(chan int)
	go func() {
		for _, seat := range seats {
			if seat%2 == 0 {
				aisle <- seat
			}
		}
		close(aisle)
	}()

	out := []int{}
	for seat := range aisle {
		out = append(out, seat)
	}
	return out
}
```

## Walkthrough

For `[1, 2, 3, 4]`: `1` skipped, `2` sent, `3` skipped, `4` sent, then
close — the allocator collects `[2 4]`. For `[1, 3]` nothing is sent, the
close still fires, and the allocator gets `[]`.

## Pitfalls

- Putting `close(aisle)` inside the `if` closes early and later sends panic.
- Skipping `close` when no seat matched hangs the allocator forever.
- In Go, `-3 % 2` is `-1`, not `1` — compare against `0` rather than testing `== 1`.
