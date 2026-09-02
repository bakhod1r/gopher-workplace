# Aggregating Through a Results Channel

## Intuition

There are two ways to combine per-worker numbers: guard a shared counter, or
send the numbers somewhere single-threaded. The second needs no lock at all,
because exactly one goroutine ever performs the addition.

## Approach

1. Create `jobs` and `results` buffered to `len(shards)`.
2. Start `workers` goroutines that upload jobs and send the byte counts on `results`.
3. Feed the jobs, close, wait, close `results`, then sum it in the caller.

## Solution

```go
import "sync"

func TotalUploaded(shards []string, workers int, upload func(string) int) int {
	jobs := make(chan string)
	results := make(chan int, len(shards))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for shard := range jobs {
				results <- upload(shard)
			}
		}()
	}

	for _, shard := range shards {
		jobs <- shard
	}
	close(jobs)
	wg.Wait()
	close(results)

	total := 0
	for n := range results {
		total += n
	}
	return total
}
```

## Walkthrough

With shards `a` and `bb` and two workers, the results channel ends up holding
1 and 2 in whichever order the uploads finished. Addition is commutative, so
the order does not matter and the total is 3.

## Pitfalls

- Summing into a shared `total` from the workers without a mutex — a data race and a wrong total.
- Leaving `results` unbuffered while summing only after `wg.Wait()`: the workers block on send and `Wait` never returns.
- Closing `results` before `wg.Wait()`, which panics in whichever worker sends next.
