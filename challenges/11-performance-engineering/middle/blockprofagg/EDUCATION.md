# Reading A Sampled Block Profile

## Intuition

A sampled profile is a survey: each recorded event represents the many similar events that were skipped, and the rate is the multiplier that puts them back.

## Approach

1. `Scale` guards the rate and multiplies.
2. `Totals` aggregates scaled waits into a map.
3. `FractionBlocked` sums the totals and divides by the window.

## Solution

```go
func Scale(wait int64, rate int) int64 {
	if rate <= 1 {
		return wait
	}
	return wait * int64(rate)
}

func Totals(events []Event, rate int) map[string]int64 {
	out := make(map[string]int64)
	for _, e := range events {
		if e.Wait <= 0 {
			continue
		}
		out[e.Site] += Scale(e.Wait, rate)
	}
	return out
}

func FractionBlocked(events []Event, rate int, windowNS int64) float64 {
	if windowNS <= 0 {
		return 0
	}
	var total int64
	for _, v := range Totals(events, rate) {
		total += v
	}
	return float64(total) / float64(windowNS)
}
```

## Walkthrough

A fraction above 1 is not a bug: blocked time is summed across goroutines, so a program with fifty waiting workers can accumulate fifty seconds of blocking in one second of wall clock.

## Pitfalls

- Reporting raw sampled waits, which understates blocking by the rate.
- Interpreting a fraction above 1 as an error rather than as parallel waiting.
- Leaving the block profile enabled at rate 1 in production, where the recording itself becomes the bottleneck.
