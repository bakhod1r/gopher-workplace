# Test Sharder

## Intuition

A shard reduction is still a per-index write: shard `c` produces exactly one
number and stores it at `out[c]`. The only new part is computing the window
boundaries before launching each goroutine.

## Approach

1. Guard `perShard <= 0` before allocating anything.
2. Compute the shard count with ceiling division and allocate `out`.
3. For each shard, clamp `end` to `len(durations)` and pass `durations[start:end]` in.
4. Sum locally inside the goroutine, then write `out[c]` once.

## Solution

```go
// Package testsharder — Gopher Workplace challenge.
package testsharder

import (
	"sync"
)

// ShardDurations returns the total runtime of each CI shard.
//
// Examples:
//
//	ShardDurations([]int{10, 20, 30, 40}, 2)  => [30 70]
//	ShardDurations([]int{10, 20, 30}, 2)      => [30 30]
//	ShardDurations([]int{10}, 0)              => []
func ShardDurations(durations []int, perShard int) []int {
	if perShard <= 0 {
		return nil
	}
	n := (len(durations) + perShard - 1) / perShard
	out := make([]int, n)
	var wg sync.WaitGroup
	for c := 0; c < n; c++ {
		start := c * perShard
		end := start + perShard
		if end > len(durations) {
			end = len(durations)
		}
		wg.Add(1)
		go func(c int, part []int) {
			defer wg.Done()
			total := 0
			for _, d := range part {
				total += d
			}
			out[c] = total
		}(c, durations[start:end])
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `{10,20,30,40}` with 2 tests per shard gives windows `[10 20]` and `[30 40]`.
- The two goroutines write `30` and `70` to their own indices.
- With three durations the tail shard holds a single test, so the totals are `[30 30]`.

## Pitfalls

- Slicing inside the goroutine with a captured `start` that has already moved on.
- Forgetting to clamp `end`, which panics on the ragged final shard.
- Accumulating into one shared `total` — a data race, and the wrong answer.
