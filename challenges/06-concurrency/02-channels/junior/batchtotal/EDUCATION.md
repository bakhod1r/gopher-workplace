# Send, Close, Range

## Intuition

When one goroutine both fills and drains a channel, the sequence is fixed:
send everything (the buffer must be large enough), close, then range. Any
other order deadlocks or panics. It is the simplest complete channel
lifecycle you can write.

## Approach

1. Make a channel of capacity `len(amounts)`.
2. Send every amount.
3. `close` the channel.
4. `range` over it, accumulating.
5. Return the total.

## Solution

```go
func BatchTotal(amounts []int) int {
	staged := make(chan int, len(amounts))
	for _, amount := range amounts {
		staged <- amount
	}
	close(staged)

	total := 0
	for amount := range staged {
		total += amount
	}
	return total
}
```

## Walkthrough

For `[100, 250, 99]`: three sends fill a 3-slot buffer, the close marks the
end, and the range yields the three amounts for a total of `449`. For an
empty day the channel has capacity 0, is closed immediately, and the loop
never runs.

## Pitfalls

- Closing before sending panics on the first send.
- A buffer smaller than `len(amounts)` blocks — nothing is receiving yet.
- Skipping the close leaves `range` waiting after the last amount.
