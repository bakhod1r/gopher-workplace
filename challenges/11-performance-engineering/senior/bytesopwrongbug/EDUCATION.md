# Megabytes That Are Not Megabytes

## Intuition

Bytes per nanosecond times 1000 is already megabytes per second in decimal units. Anything else in the expression is converting to a different unit than the one the column claims.

## Approach

1. Remove the binary-megabyte conversion factor.

## Solution

```go
func ThroughputMBs(totalBytes, elapsedNS int64) float64 {
	if elapsedNS <= 0 {
		return 0
	}
	return float64(totalBytes) / float64(elapsedNS) * 1000
}
```

## Walkthrough

The extra factor `1e6/1048576` is about 0.9537, so every number came out 4.63% low — small enough to look like noise, consistent enough to survive every rerun, and just large enough to bury a real 5% win.

## Pitfalls

- Using `1 << 20` for a megabyte anywhere near benchmark output.
- Mixing MB and MiB in one dashboard, so two panels disagree by a constant.
- Chasing a fixed-percentage discrepancy as if it were a performance problem.
