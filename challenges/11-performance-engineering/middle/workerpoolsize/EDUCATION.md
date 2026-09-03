# How Many Workers

## Intuition

Parallelism is a resource with a ceiling. Below the ceiling you gain throughput; above it you pay for scheduling and get nothing back.

## Approach

1. Allocate the result up front so every worker has its own slot.
2. Send indexes down a channel; each worker ranges it and writes its own results.
3. Wait for the workers, then return.

## Solution

```go
func Map(items []int, workers int, f func(int) int) []int {
	out := make([]int, len(items))
	if len(items) == 0 {
		return out
	}
	if workers < 1 {
		workers = 1
	}
	workers = min(workers, len(items))
	idx := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idx {
				out[i] = f(items[i])
			}
		}()
	}
	for i := range items {
		idx <- i
	}
	close(idx)
	wg.Wait()
	return out
}

func Sizing(cpus int, blocked float64) int {
	if cpus < 1 {
		return 1
	}
	if blocked < 0 || blocked >= 1 {
		blocked = 0
	}
	return max(int(float64(cpus)/(1-blocked)), 1)
}
```

## Walkthrough

Treating an out-of-range `blocked` as 0 rather than clamping it to just under 1 is the safer failure: a nonsense fraction falls back to the CPU-bound answer instead of asking for eight million goroutines.

Writing `out[i]` from several goroutines is safe because the slots are distinct memory: no two workers ever hold the same `i`, so there is no shared word to race on. `-race` confirms it.

## Pitfalls

- Collecting results by appending from workers, which both races and scrambles the order.
- Forgetting `close(idx)`, so the workers block forever and `wg.Wait()` never returns.
- Sizing a pool for I/O-bound work by core count, leaving the CPUs idle behind blocked goroutines.
