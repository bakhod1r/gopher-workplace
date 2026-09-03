# Nanoseconds Into Cycles

## Intuition

Gigahertz means "billions of cycles per second", and a nanosecond is a billionth of a second, so the two conversion factors cancel: cycles equal nanoseconds times GHz.

## Approach

1. Guard both inputs, multiply.
2. Classify with a tagless `switch`, ordered from the smallest band up.

## Solution

```go
func CyclesPerOp(nsPerOp, ghz float64) float64 {
	if nsPerOp <= 0 || ghz <= 0 {
		return 0
	}
	return nsPerOp * ghz
}

func Verdict(cycles float64) string {
	switch {
	case cycles < 10:
		return "register"
	case cycles < 100:
		return "l1"
	case cycles < 1000:
		return "memory"
	default:
		return "syscall"
	}
}
```

## Walkthrough

A "fast" 40 ns/op function is 120 cycles at 3 GHz — the `memory` band, which says the time is going to cache misses, not arithmetic.

## Pitfalls

- Dividing by GHz instead of multiplying.
- Overlapping bands with `<=`, which puts 100 in two categories.
- Reading a cycle count as portable across machines with different clocks.
