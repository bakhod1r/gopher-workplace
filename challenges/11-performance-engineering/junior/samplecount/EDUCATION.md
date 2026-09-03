# Samples Are Not Nanoseconds

## Intuition

The runtime sets a timer, and every time it fires the profiler records the current stack. The "time" column is just that count multiplied by how long each tick represents.

## Approach

1. Range the samples once.
2. Skip anything with a non-positive count or period.
3. Accumulate the count and the product.

## Solution

```go
func Totals(samples []Sample) (count int64, nanos int64) {
	for _, s := range samples {
		if s.Count <= 0 || s.Period <= 0 {
			continue
		}
		count += s.Count
		nanos += s.Count * s.Period
	}
	return count, nanos
}
```

## Walkthrough

At the default 100 Hz, a period of 10ms means a function credited with 3 samples is reported as 30ms — a resolution that cannot distinguish 25ms from 34ms.

## Pitfalls

- Adding `Period` once per sample instead of once per observation.
- Trusting a fine-grained number derived from a handful of samples.
- Skipping only one of the two guards, letting a negative product through.
