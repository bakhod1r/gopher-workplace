# Last Value and the Seen Flag

## Intuition

Streams have no index, so "the last one" only means something once the
stream ends. Draining to close gives you that. The `seen` flag is the same
idea as comma-ok: without it, a silent device and a device whose last
reading was `0` are indistinguishable.

## Approach

1. Declare `last` (0) and `seen` (false).
2. `range` over `readings`, assigning `last = v` and `seen = true`.
3. Return both.

## Solution

```go
func LastReading(readings <-chan int) (int, bool) {
	last := 0
	seen := false
	for v := range readings {
		last = v
		seen = true
	}
	return last, seen
}
```

## Walkthrough

For `1, 2, 3`: `last` becomes 1, then 2, then 3, `seen` is true — result
`3, true`. For a device that disconnected without reporting, the loop never
runs and the initial `0, false` is returned.

## Pitfalls

- Returning `last != 0` as the flag misreports a device whose last reading really is `0`.
- Ranging over a stream the gateway never closes blocks the dashboard forever.
- There is no way to peek at the last reading without consuming everything before it.
