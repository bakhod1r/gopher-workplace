# Telemetry Flusher

## Intuition

The tempting bug here is a single shared `errors++`. Give every goroutine a
private counter and a private destination and the whole class of problem
disappears — no mutex required.

## Approach

1. Guard `batch <= 0` before allocating.
2. Compute the batch count with ceiling division and allocate `out`.
3. Clamp `end` to `len(codes)` and pass the window into the goroutine.
4. Count locally, then perform a single write to `out[c]`.

## Solution

```go
// Package telemetryflusher — Gopher Workplace challenge.
package telemetryflusher

import (
	"sync"
)

// BatchErrorCounts returns how many server errors each flush batch contains.
//
// Examples:
//
//	BatchErrorCounts([]int{200, 500, 503, 200}, 2)  => [1 1]
//	BatchErrorCounts([]int{500, 500, 200}, 2)       => [2 0]
//	BatchErrorCounts([]int{200}, 0)                 => []
func BatchErrorCounts(codes []int, batch int) []int {
	if batch <= 0 {
		return nil
	}
	n := (len(codes) + batch - 1) / batch
	out := make([]int, n)
	var wg sync.WaitGroup
	for c := 0; c < n; c++ {
		start := c * batch
		end := start + batch
		if end > len(codes) {
			end = len(codes)
		}
		wg.Add(1)
		go func(c int, part []int) {
			defer wg.Done()
			errors := 0
			for _, code := range part {
				if code >= 500 {
					errors++
				}
			}
			out[c] = errors
		}(c, codes[start:end])
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `{200,500,503,200}` with batch 2 gives windows `[200 500]` and `[503 200]`.
- Each window holds one server error, so the summaries are `[1 1]`.
- A `404` is a client error and is not counted.

## Pitfalls

- A shared `errors` variable incremented by all goroutines — `-race` will flag it.
- Counting `> 500` instead of `>= 500`, which misses a plain `500`.
- Launching goroutines before checking `batch`.
