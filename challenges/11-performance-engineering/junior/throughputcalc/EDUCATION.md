# From ns/op To Requests Per Second

## Intuition

Latency and throughput are the same measurement upside down: if one operation takes `t` seconds, a second fits `1/t` of them.

## Approach

1. `OpsPerSec` guards and returns `1e9 / nsPerOp`.
2. `Capacity` multiplies by the core count and converts to `int64`, which truncates.

## Solution

```go
func OpsPerSec(nsPerOp float64) float64 {
	if nsPerOp <= 0 {
		return 0
	}
	return 1e9 / nsPerOp
}

func Capacity(nsPerOp float64, cores int) int64 {
	if cores <= 0 {
		return 0
	}
	return int64(OpsPerSec(nsPerOp) * float64(cores))
}
```

## Walkthrough

`int64(...)` truncates toward zero, so 4.5 ops/sec becomes 4 — a capacity estimate should never round up.

## Pitfalls

- Using `1e6` and reporting per-millisecond rates by accident.
- Multiplying latency by cores instead of dividing — more cores do not make one operation faster.
- Treating the perfect-scaling number as achievable; locks and memory bandwidth take their cut.
