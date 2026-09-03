# Sliding Windows In One Pass

## Intuition

Consecutive windows overlap in `n-1` samples. Recomputing that overlap is the entire waste.

## Approach

1. Guard the window size.
2. Sum the first window.
3. Slide: add `samples[i]`, subtract `samples[i-n]`, record the mean.

## Solution

```go
func Window(samples []float64, n int) []float64 {
	out := make([]float64, 0, max(len(samples)-n+1, 0))
	if n <= 0 || n > len(samples) {
		return out
	}
	sum := 0.0
	for _, v := range samples[:n] {
		sum += v
	}
	out = append(out, sum/float64(n))
	for i := n; i < len(samples); i++ {
		sum += samples[i] - samples[i-n]
		out = append(out, sum/float64(n))
	}
	return out
}

func Smoothest(means []float64) int {
	best := -1
	for i, v := range means {
		if best == -1 || v < means[best] {
			best = i
		}
	}
	return best
}
```

## Walkthrough

The loop starts at `i = n` because the first window is already summed; each iteration then costs one addition and one subtraction regardless of how wide the window is.

## Pitfalls

- Re-summing `samples[i-n+1 : i+1]` each step, which is the O(n·k) version.
- Emitting `len(samples)-n` windows and dropping the last one.
- Accumulating a running sum over millions of float samples without noticing the drift; for long series, resum periodically.
