# Parallel Map

## Intuition

Goroutines give you concurrency; `GOMAXPROCS` decides how much of it becomes parallelism. Chunking contiguous ranges keeps each worker on its own cache lines and makes the output slots naturally disjoint.

## Approach

1. Allocate the output first so every slot exists before the workers start.
2. Clamp the worker count to `GOMAXPROCS(0)` and to the input length.
3. Compute a ceiling chunk size and give each worker `[start, end)`.
4. Each worker writes only its own indexes; `wg.Wait()` joins them.

## Solution

```go
func (SquareOp) Apply(v int) int { return v * v }

func MapParallel(op Op, vs []int) []int {
	out := make([]int, len(vs))
	n := len(vs)
	if n == 0 {
		return out
	}

	workers := runtime.GOMAXPROCS(0)
	if workers > n {
		workers = n
	}
	chunk := (n + workers - 1) / workers

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		start := w * chunk
		if start >= n {
			break
		}
		end := start + chunk
		if end > n {
			end = n
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				out[i] = op.Apply(vs[i])
			}
		}(start, end)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

The prime-length test catches the classic chunking bug: with a ceiling chunk size, the last worker's range must be clamped to `n`, and some workers may get no work at all.

## Pitfalls

- Spawning one goroutine per element — huge scheduling overhead and no extra parallelism.
- Interleaved striding (`i += workers`), which is correct but touches every cache line from every core.
- Using a floor chunk size, which silently drops the tail elements.
