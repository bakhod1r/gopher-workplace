# Metric Batch Flusher

## Intuition

A batcher is a loop with one piece of state: the partial batch. Two things can end a batch — it filled up, or the input ended. Both paths must send, and only the second one closes.

## Approach

1. Normalise `size` to at least 1.
2. Make the output channel and start one goroutine with `defer close(out)`.
3. `range` over `in`, appending to `batch`.
4. When `len(batch) == size`, send it and allocate a fresh slice.
5. After the loop, send the tail if it is non-empty.

## Solution

```go
func FlushBatches(in <-chan int, size int) <-chan []int {
	if size < 1 {
		size = 1
	}

	out := make(chan []int)
	go func() {
		defer close(out)

		batch := make([]int, 0, size)
		for s := range in {
			batch = append(batch, s)
			if len(batch) == size {
				out <- batch
				batch = make([]int, 0, size)
			}
		}
		if len(batch) > 0 {
			out <- batch
		}
	}()
	return out
}
```

## Walkthrough

With samples 1..5 and size 2: 1,2 fills a batch and it is sent; a new slice takes 3,4 and is sent; 5 lands in a third slice. `in` closes, the `range` ends, the partial `[5]` goes out, and the deferred `close(out)` lets the consumer's `range` finish.

## Pitfalls

- Reusing the batch slice with `batch[:0]` — the consumer's earlier batch mutates under it.
- Sending an empty tail batch when the input length divides evenly.
- Forgetting to close `out`, which hangs the consumer's `range` forever.
