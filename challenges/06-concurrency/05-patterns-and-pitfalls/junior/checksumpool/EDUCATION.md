# Attributable Results from a Pool

## Intuition

Results from a pool arrive out of order, so they must be self-describing.
Once each result carries its key, the caller can build the map single-handedly
and the whole program needs no lock.

## Approach

1. Define a small `result` struct holding the file and its checksum.
2. Start `workers` goroutines sending `result` values on a channel buffered to `len(files)`.
3. Close jobs, wait, close results, then range over results filling the map in the caller.

## Solution

```go
import "sync"

// result pairs a file with the checksum a worker computed for it.
type result struct {
	file string
	sum  int
}

func ChecksumFiles(files []string, workers int, sum func(string) int) map[string]int {
	jobs := make(chan string)
	results := make(chan result, len(files))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range jobs {
				results <- result{file: file, sum: sum(file)}
			}
		}()
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)
	wg.Wait()
	close(results)

	sums := make(map[string]int, len(files))
	for r := range results {
		sums[r.file] = r.sum
	}
	return sums
}
```

## Walkthrough

With `a` and `bb`, the two results may arrive in either order. Because each
carries its filename, the caller's loop writes `sums["a"] = 1` and
`sums["bb"] = 2` regardless of which came first.

## Pitfalls

- Writing the map from the workers — concurrent map writes are a fatal runtime error.
- Sending bare ints and trying to match them to files by position; completion order is not job order.
- Closing `results` before `wg.Wait()`, which panics in whichever worker sends next.
