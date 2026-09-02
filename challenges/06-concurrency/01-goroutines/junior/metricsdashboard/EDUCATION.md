# Metrics Dashboard

## Intuition

Downsampling turns one long scan into several independent scans. Each goroutine
reduces its window to a single number and stores it at its own index — the same
per-index write as every other puzzle here.

## Approach

1. Guard `window <= 0` and compute the window count by ceiling division.
2. Clamp `end` to `len(samples)` and pass `samples[start:end]` into the goroutine.
3. Scan locally from `part[0]`, then write `out[c]` once.

## Solution

```go
// Package metricsdashboard — Gopher Workplace challenge.
package metricsdashboard

import (
	"sync"
)

// PeakPerWindow returns the peak value of each fixed-size time window.
//
// Examples:
//
//	PeakPerWindow([]int{1, 9, 3, 4}, 2)  => [9 4]
//	PeakPerWindow([]int{5, 2, 8}, 2)     => [5 8]
//	PeakPerWindow([]int{1}, 0)           => []
func PeakPerWindow(samples []int, window int) []int {
	if window <= 0 {
		return nil
	}
	n := (len(samples) + window - 1) / window
	out := make([]int, n)
	var wg sync.WaitGroup
	for c := 0; c < n; c++ {
		start := c * window
		end := start + window
		if end > len(samples) {
			end = len(samples)
		}
		wg.Add(1)
		go func(c int, part []int) {
			defer wg.Done()
			peak := part[0]
			for _, v := range part[1:] {
				if v > peak {
					peak = v
				}
			}
			out[c] = peak
		}(c, samples[start:end])
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `{1,9,3,4}` with window 2 reduces `[1 9]` to `9` and `[3 4]` to `4`.
- With three samples the tail window holds one value, giving `[5 8]`.
- An all-negative series peaks at `-4`, which only works because the seed comes from the data.

## Pitfalls

- Seeding the peak with `0`, which is wrong for any window below zero.
- Slicing inside the goroutine from a captured `start` that has already advanced.
- Forgetting to clamp `end`, which panics on the ragged final window.
