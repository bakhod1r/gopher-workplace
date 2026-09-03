# The Same Work Every Run

## Intuition

A linear congruential generator is a pure function of its state, so the same seed replays the same sequence — on any machine, in any run order.

## Approach

1. Allocate `n` slots (clamping a negative `n` to zero).
2. Step the state and store it, `n` times.

## Solution

```go
func FixedInput(seed uint32, n int) []uint32 {
	if n < 0 {
		n = 0
	}
	out := make([]uint32, 0, n)
	state := seed
	for i := 0; i < n; i++ {
		state = state*1664525 + 1013904223
		out = append(out, state)
	}
	return out
}
```

## Walkthrough

`uint32` multiplication wraps silently, which is the whole point: the sequence is well defined without any explicit masking.

## Pitfalls

- Using `math/rand` without a fixed seed, which changes the workload per run.
- Emitting the state before stepping it, which leaks the raw seed as the first value.
- Doing the generation inside the timed loop, so you benchmark the generator.
