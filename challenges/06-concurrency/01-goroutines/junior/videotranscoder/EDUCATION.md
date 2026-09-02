# Video Transcoder

## Intuition

A value that is written once, before any goroutine starts, and only read
afterwards is safe for every goroutine to touch. Only the per-track result needs
a private home.

## Approach

1. Allocate `out := make([]int, len(bitrates))`.
2. Launch one goroutine per track, passing `i` and the source bitrate.
3. Return `0` early for a non-positive factor; otherwise write `b * factorPct / 100`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package videotranscoder — Gopher Workplace challenge.
package videotranscoder

import (
	"sync"
)

// TargetBitrates scales every source bitrate down to the requested percentage.
//
// Examples:
//
//	TargetBitrates([]int{4000, 2000}, 50)  => [2000 1000]
//	TargetBitrates([]int{4000}, 100)       => [4000]
//	TargetBitrates(nil, 50)                => []
func TargetBitrates(bitrates []int, factorPct int) []int {
	out := make([]int, len(bitrates))
	var wg sync.WaitGroup
	for i, b := range bitrates {
		wg.Add(1)
		go func(i int, b int) {
			defer wg.Done()
			if factorPct <= 0 {
				out[i] = 0
				return
			}
			out[i] = b * factorPct / 100
		}(i, b)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `4000` at 50% becomes `4000 * 50 / 100 = 2000`.
- At 100% the bitrate is unchanged.
- A negative factor is nonsense for an encoder, so the goroutine returns `0`.

## Pitfalls

- Dividing first (`b / 100 * factorPct`), which truncates small bitrates to zero.
- Skipping the guard and emitting a negative bitrate.
- Mutating `bitrates` in place instead of filling a new slice.
