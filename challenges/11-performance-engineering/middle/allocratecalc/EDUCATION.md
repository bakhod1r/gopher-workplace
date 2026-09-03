# Rates From Cumulative Counters

## Intuition

A rate is a difference over a duration. With unsigned counters the difference is only meaningful when you have already established which snapshot is later.

## Approach

1. `Delta` validates the timestamp order and every counter's direction before subtracting.
2. `BytesPerSec` divides the byte delta by the elapsed seconds.
3. `LiveObjects` guards the subtraction the same way.

## Solution

```go
func Delta(from, to Stats) (uint64, uint64, uint64, bool) {
	if to.NS <= from.NS {
		return 0, 0, 0, false
	}
	if to.TotalAlloc < from.TotalAlloc || to.Mallocs < from.Mallocs || to.Frees < from.Frees {
		return 0, 0, 0, false
	}
	return to.TotalAlloc - from.TotalAlloc, to.Mallocs - from.Mallocs, to.Frees - from.Frees, true
}

func BytesPerSec(from, to Stats) (float64, bool) {
	bytes, _, _, ok := Delta(from, to)
	if !ok {
		return 0, false
	}
	seconds := float64(to.NS-from.NS) / 1e9
	return float64(bytes) / seconds, true
}

func LiveObjects(s Stats) (uint64, bool) {
	if s.Frees > s.Mallocs {
		return 0, false
	}
	return s.Mallocs - s.Frees, true
}
```

## Walkthrough

Every guard here is a comparison placed *before* a subtraction, because after the subtraction the evidence is gone: the wrapped result is a perfectly ordinary large `uint64` that no later check can distinguish from real data.

## Pitfalls

- Subtracting first and checking the result for "unreasonably large", which is a guess dressed as validation.
- Casting to `int64` before subtracting, which merely moves the overflow.
- Graphing across a process restart, producing a spike that sends people hunting for a leak that never existed.
