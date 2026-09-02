# Sharding Work, Merging Under a Lock

## Intuition

Lock contention is proportional to how often you take the lock, not to how
much data you merge. Counting locally and merging once turns thousands of
contended increments into one short critical section per chunk.

## Approach

1. Create the shared `totals` map, a `WaitGroup`, and a `sync.Mutex`.
2. Per chunk, start a goroutine that builds a local count map.
3. Lock, add every local entry into `totals`, unlock; `wg.Wait()` then return `totals`.

## Solution

```go
import "sync"

func CountLevels(chunks [][]string, level func(string) string) map[string]int {
	totals := make(map[string]int)

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	for _, chunk := range chunks {
		wg.Add(1)
		go func(chunk []string) {
			defer wg.Done()

			local := make(map[string]int)
			for _, line := range chunk {
				local[level(line)]++
			}

			mu.Lock()
			defer mu.Unlock()
			for lvl, n := range local {
				totals[lvl] += n
			}
		}(chunk)
	}

	wg.Wait()
	return totals
}
```

## Walkthrough

Two chunks each holding one `ERR` line produce two local maps of `{ERR: 1}`.
The first goroutine to get the lock sets `totals["ERR"] = 1`; the second adds
to it, giving 2. Whichever order they run in, the total is the same.

## Pitfalls

- Writing `totals[level(line)]++` directly from the goroutines — that is a concurrent map write and crashes the process.
- Holding the lock for the whole chunk, which serialises all the work and removes any benefit from concurrency.
- Reading `totals` before `wg.Wait()`, which can miss the merges still in flight.
