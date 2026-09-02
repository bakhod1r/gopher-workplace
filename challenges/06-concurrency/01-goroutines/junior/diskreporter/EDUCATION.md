# Disk Reporter

## Intuition

`defer wg.Done()` is what makes every exit path from a goroutine — early return
included — still release the WaitGroup. Without it, one empty volume would hang
the whole report.

## Approach

1. Allocate `out := make([]int, len(volumes))`.
2. Launch one goroutine per volume, passing `i` and the volume's sizes.
3. Return `0` early when the volume is empty; otherwise scan from `vol[0]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package diskreporter — Gopher Workplace challenge.
package diskreporter

import (
	"sync"
)

// LargestFiles returns the largest file size on each volume.
//
// Examples:
//
//	LargestFiles([][]int{{30, 10}, {5}})  => [30 5]
//	LargestFiles([][]int{{}})             => [0]
//	LargestFiles(nil)                     => []
func LargestFiles(volumes [][]int) []int {
	out := make([]int, len(volumes))
	var wg sync.WaitGroup
	for i, vol := range volumes {
		wg.Add(1)
		go func(i int, vol []int) {
			defer wg.Done()
			if len(vol) == 0 {
				out[i] = 0
				return
			}
			largest := vol[0]
			for _, size := range vol[1:] {
				if size > largest {
					largest = size
				}
			}
			out[i] = largest
		}(i, vol)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `{30,10}` seeds `largest = 30` and never improves on it.
- An empty volume returns early, leaving its slot at the zero value `0`.
- `{4,8}` seeds `4` and raises it to `8`.

## Pitfalls

- Calling `wg.Done()` at the bottom instead of deferring it — the early return deadlocks `Wait`.
- Indexing `vol[0]` before checking the length, which panics inside a goroutine and kills the process.
- Tracking one fleet-wide maximum, which both races and answers a different question.
