# The Pause Budget

## Intuition

Two questions, two numbers: how much of the window did the collector take in total, and what was the longest single interruption a request could have landed in.

## Approach

1. `Total` and `Worst` walk the pauses once each, skipping negatives.
2. `FractionOf` guards the window and divides.
3. `WithinBudget` combines the two checks with inclusive comparisons.

## Solution

```go
func Total(pauses []int64) int64 {
	var total int64
	for _, p := range pauses {
		if p >= 0 {
			total += p
		}
	}
	return total
}

func Worst(pauses []int64) (int64, int) {
	best, at := int64(0), -1
	for i, p := range pauses {
		if p < 0 {
			continue
		}
		if at == -1 || p > best {
			best, at = p, i
		}
	}
	return best, at
}

func FractionOf(pauses []int64, windowNS int64) float64 {
	if windowNS <= 0 {
		return 0
	}
	return float64(Total(pauses)) / float64(windowNS)
}

func WithinBudget(pauses []int64, windowNS int64, maxFraction float64, maxPauseNS int64) bool {
	if FractionOf(pauses, windowNS) > maxFraction {
		return false
	}
	worst, at := Worst(pauses)
	return at == -1 || worst <= maxPauseNS
}
```

## Walkthrough

Tracking `at == -1` rather than comparing against a zero initial value is what keeps the earliest index on a tie and distinguishes "no pauses" from "a pause of zero".

## Pitfalls

- Checking only the total, so one 200ms stop passes a budget it should blow.
- Strict `<` comparisons, which reject a value sitting exactly on the documented limit.
- Comparing pause totals across runs of different lengths without normalising by the window.
