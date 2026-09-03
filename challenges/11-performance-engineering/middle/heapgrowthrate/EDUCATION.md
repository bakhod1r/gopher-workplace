# Is The Heap Growing, Or Just Breathing?

## Intuition

Two endpoints define a line. The line's slope answers the leak question; the sawtooth in between is just the collector doing its job.

## Approach

1. `NextTarget` applies the GOGC ratio, rejecting a non-positive setting.
2. `GrowthPerSec` validates the span and divides the byte difference by the seconds.
3. `Doubling` divides the current live heap by the slope.

## Solution

```go
func NextTarget(liveBytes int64, gogc int) (int64, bool) {
	if gogc <= 0 {
		return 0, false
	}
	return liveBytes + liveBytes*int64(gogc)/100, true
}

func GrowthPerSec(samples []Sample) (float64, bool) {
	if len(samples) < 2 {
		return 0, false
	}
	first, last := samples[0], samples[len(samples)-1]
	if last.NS <= first.NS {
		return 0, false
	}
	seconds := float64(last.NS-first.NS) / 1e9
	return float64(last.Live-first.Live) / seconds, true
}

func Doubling(samples []Sample) (float64, bool) {
	slope, ok := GrowthPerSec(samples)
	if !ok || slope <= 0 {
		return 0, false
	}
	live := float64(samples[len(samples)-1].Live)
	if live <= 0 {
		return 0, false
	}
	return live / slope, true
}
```

## Walkthrough

`GOGC=0` does not mean "collect immediately" here — in the runtime it is the value that disables the percentage trigger entirely, which is why it is reported as "no target" rather than as zero bytes.

## Pitfalls

- Comparing consecutive samples and calling every rise a leak; the sawtooth rises constantly.
- Dividing by a zero or negative time span when two samples share a timestamp.
- Reporting a doubling time from a slope measured over a few seconds of an application's warm-up.
