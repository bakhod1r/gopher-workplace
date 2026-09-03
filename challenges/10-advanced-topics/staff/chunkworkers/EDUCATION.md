# Split The Work, Not The Memory

## Intuition

Slices are views, so splitting the work costs nothing: every goroutine reads a region nobody else touches, which needs no synchronisation. Only the results have to come back, and one slot per worker plus a `Wait` is the whole protocol.

## Approach

1. Normalise `workers` against the input length.
2. Compute the ceiling-divided chunk size and clamp each chunk's end.
3. Give each goroutine `s[start:end]` and a private slot in `partial`.
4. `Wait`, then sum the partials.

## Solution

```go
import "sync"

// SumParallel sums s using workers goroutines over disjoint chunks of
// the input and returns the total.
//
// The input must not be copied: each worker gets a view. Parallelism has to
// be real — no locking on a shared accumulator per element.
//
// Examples:
//
// 	SumParallel([]int{1, 2, 3, 4}, 2) => 10
func SumParallel(s []int, workers int) int64 {
	if workers < 1 {
		workers = 1
	}
	if len(s) == 0 {
		return 0
	}
	if workers > len(s) {
		workers = len(s)
	}
	partial := make([]int64, workers)
	size := (len(s) + workers - 1) / workers
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		start := w * size
		end := start + size
		if start > len(s) {
			start = len(s)
		}
		if end > len(s) {
			end = len(s)
		}
		go func(w int, part []int) {
			defer wg.Done()
			var sum int64
			for _, v := range part {
				sum += int64(v)
			}
			partial[w] = sum
		}(w, s[start:end])
	}
	wg.Wait()
	var total int64
	for _, p := range partial {
		total += p
	}
	return total
}
```

## Walkthrough

For 100003 elements and 7 workers, the chunk size is 14286; the last chunk is clamped to end at 100003. Each worker sums into its own slot, and the `Wait` publishes all seven before the final fold.

## Pitfalls

- `atomic.AddInt64` per element — correct and slower than the serial loop.
- `start + size` without clamping, which panics on the last chunk.
- Summing `partial` before `Wait`.
