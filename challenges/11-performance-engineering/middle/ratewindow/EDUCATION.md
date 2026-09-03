# Rates Over A Time Window

## Intuition

The events are sorted, so the window is a contiguous slice of them. Two binary searches find its ends, and everything else is arithmetic on that slice.

## Approach

1. Guard the width.
2. Binary search for the first event at or after `from` and the first at or after `from+width`.
3. Count, sum, or divide by the width in seconds.

## Solution

```go
func bounds(events []Event, fromNS, widthNS int64) (int, int) {
	lo := sort.Search(len(events), func(i int) bool { return events[i].NS >= fromNS })
	hi := sort.Search(len(events), func(i int) bool { return events[i].NS >= fromNS+widthNS })
	return lo, hi
}

func CountIn(events []Event, fromNS, widthNS int64) int {
	if widthNS <= 0 {
		return 0
	}
	lo, hi := bounds(events, fromNS, widthNS)
	return hi - lo
}

func RatePerSec(events []Event, fromNS, widthNS int64) float64 {
	if widthNS <= 0 {
		return 0
	}
	return float64(CountIn(events, fromNS, widthNS)) / (float64(widthNS) / 1e9)
}

func SumIn(events []Event, fromNS, widthNS int64) int64 {
	if widthNS <= 0 {
		return 0
	}
	lo, hi := bounds(events, fromNS, widthNS)
	var total int64
	for _, e := range events[lo:hi] {
		total += e.Value
	}
	return total
}
```

## Walkthrough

Both searches use `>=`, which places the event at exactly `from+width` outside the window and at exactly `from` inside it — the half-open convention that makes adjacent windows tile without overlap.

## Pitfalls

- Using `>` for the upper bound, so the boundary event is counted twice across two windows.
- Dividing by the width in nanoseconds and reporting a rate a billion times too small.
- Scanning the whole series per window, which turns a dashboard query into an O(n·windows) loop.
