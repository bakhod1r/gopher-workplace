# Ordered Results Without a Lock

## Intuition

A data race needs two goroutines touching the *same* memory. Writing distinct
elements of a preallocated slice never does, so the simplest concurrent map
over a slice needs nothing but a `WaitGroup`.

## Approach

1. `sizes := make([]int, len(paths))`.
2. For each `i, path`: `wg.Add(1)`, then a goroutine writing `sizes[i] = size(path)`.
3. `wg.Wait()` and return `sizes`.

## Solution

```go
import "sync"

func FileSizes(paths []string, size func(string) int) []int {
	sizes := make([]int, len(paths))

	var wg sync.WaitGroup
	for i, path := range paths {
		wg.Add(1)
		go func(i int, path string) {
			defer wg.Done()
			sizes[i] = size(path)
		}(i, path)
	}

	wg.Wait()
	return sizes
}
```

## Walkthrough

For three paths, three goroutines run in any order and write indexes 0, 1 and
2. `wg.Wait()` establishes that all three writes happened before the caller
reads the slice, so the manifest is both correct and stable.

## Pitfalls

- Using `append` from the goroutines: a race on the slice header and a scrambled manifest.
- Calling `wg.Add(1)` inside the goroutine — `Wait` can return before those goroutines even start.
- Reading `sizes` before `wg.Wait()`, where the writes are not yet guaranteed visible.
