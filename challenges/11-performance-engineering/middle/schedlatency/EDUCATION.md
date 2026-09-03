# Waiting For A Core

## Intuition

Every goroutine that becomes runnable joins a queue. The time it spends there is time your program is not doing anything, and no CPU profile will show it — the goroutine is not on a CPU.

## Approach

1. `Delay` validates and subtracts.
2. `Delays` filters through `Delay`.
3. `Stats` walks the delays once, accumulating the sum and the maximum.

## Solution

```go
func Delay(e Event) (int64, bool) {
	if e.Running < e.Runnable {
		return 0, false
	}
	return e.Running - e.Runnable, true
}

func Delays(events []Event) []int64 {
	out := make([]int64, 0, len(events))
	for _, e := range events {
		if d, ok := Delay(e); ok {
			out = append(out, d)
		}
	}
	return out
}

func Stats(events []Event) (float64, int64, bool) {
	ds := Delays(events)
	if len(ds) == 0 {
		return 0, 0, false
	}
	var sum, worst int64
	for _, d := range ds {
		sum += d
		worst = max(worst, d)
	}
	return float64(sum) / float64(len(ds)), worst, true
}
```

## Walkthrough

A zero delay is valid and common — a goroutine made runnable on an idle P starts immediately — so the validity check has to be `Running < Runnable`, not `<=`.

## Pitfalls

- Rejecting zero-delay events and dropping the healthy majority of the data.
- Reporting only the mean, which stays flat while the tail degrades.
- Responding to high scheduling latency by launching more goroutines.
