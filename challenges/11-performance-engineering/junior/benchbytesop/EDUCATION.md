# The MB/s Column

## Intuition

Bytes per nanosecond is already gigabytes per second; the megabyte and second conversions differ by exactly 1000, so the whole formula collapses to `bytes / ns * 1000`.

## Approach

1. Guard `elapsedNS <= 0`.
2. Convert both operands to `float64` and divide, scaling by 1000.

## Solution

```go
func ThroughputMBs(totalBytes int64, elapsedNS int64) float64 {
	if elapsedNS <= 0 {
		return 0
	}
	return float64(totalBytes) / float64(elapsedNS) * 1000
}
```

## Walkthrough

`1e6 bytes / 1e9 ns = 1e-3 bytes/ns`; multiplied by 1000 that is 1 MB/s, matching the tool's output.

## Pitfalls

- `totalBytes / elapsedNS` in integer arithmetic, which is 0 for every realistic input.
- Using `1 << 20` for a megabyte, which disagrees with the benchmark tool.
- Dividing by zero when the benchmark never ran.
