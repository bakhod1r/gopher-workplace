# Worker Pool

## Intuition

Goroutines are cheap, not free. A pool converts "one goroutine per task" into a fixed cost, and indexing results by position removes the need for any lock on the output.

## Approach

1. Allocate `out` with the task count so every slot exists before the workers start.
2. Start `workers` goroutines, each ranging over an index channel.
3. Each worker writes `out[i]` — a slot no other worker touches.
4. Feed indexes, `close` the channel, then `wg.Wait()`.

## Solution

```go
func (s SquareTask) Run() int { return s.N * s.N }

func RunAll(tasks []Task, workers int) []int {
	out := make([]int, len(tasks))
	if len(tasks) == 0 {
		return out
	}
	if workers < 1 {
		workers = 1
	}

	idx := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idx {
				out[i] = tasks[i].Run()
			}
		}()
	}

	for i := range tasks {
		idx <- i
	}
	close(idx)
	wg.Wait()
	return out
}
```

## Walkthrough

`TestConcurrencyBounded` records peak simultaneous `Run` calls with atomics. With four workers the peak never exceeds four, which is the graded constraint.

## Pitfalls

- `go func() { out = append(out, ...) }()` — concurrent appends race and lose results.
- Spawning one goroutine per task, which ignores the ceiling entirely.
- Forgetting `close(idx)`, so the workers block forever and `wg.Wait()` deadlocks.
