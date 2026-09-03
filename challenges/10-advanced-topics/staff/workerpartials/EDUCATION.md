# Aggregate In Parallel Without Sharing A Word

## Intuition

Contention is a property of sharing, not of parallelism. Give each worker its own bucket array and the inner loop becomes an ordinary single-threaded loop; the only synchronisation left is the join.

## Approach

1. Validate `buckets`, clamp `workers`, and handle the empty input.
2. Split `data` into clamped chunks and give each goroutine a view plus its own `local` bucket array.
3. Correct negative bins by adding `buckets`.
4. After `Wait`, sum the partials into the result.

## Solution

```go
import "sync"

// Histogram counts data into buckets bins by value modulo buckets, using
// workers goroutines over disjoint chunks.
//
// Workers must not share a counter: each accumulates privately and the
// results are folded once, after the join.
//
// Examples:
//
// 	Histogram([]int{0, 1, 2, 3}, 2, 2) => []int64{2, 2}
func Histogram(data []int, buckets, workers int) []int64 {
	if buckets < 1 {
		return nil
	}
	out := make([]int64, buckets)
	if len(data) == 0 {
		return out
	}
	if workers < 1 {
		workers = 1
	}
	if workers > len(data) {
		workers = len(data)
	}
	partials := make([][]int64, workers)
	size := (len(data) + workers - 1) / workers
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		start := w * size
		end := start + size
		if start > len(data) {
			start = len(data)
		}
		if end > len(data) {
			end = len(data)
		}
		go func(w int, part []int) {
			defer wg.Done()
			local := make([]int64, buckets)
			for _, v := range part {
				b := v % buckets
				if b < 0 {
					b += buckets
				}
				local[b]++
			}
			partials[w] = local
		}(w, data[start:end])
	}
	wg.Wait()
	for _, p := range partials {
		for i, c := range p {
			out[i] += c
		}
	}
	return out
}
```

## Walkthrough

With 100003 values, 13 buckets and 8 workers, each worker fills its own 13-element array; the fold afterwards touches 104 int64s once. A shared array would have needed 100003 atomic increments.

## Pitfalls

- `v % buckets` alone puts negative values at a negative index and panics.
- Folding before `Wait`, which reads partials that are still being written.
- Writing into the shared `out` from the workers, which is the original bug.
