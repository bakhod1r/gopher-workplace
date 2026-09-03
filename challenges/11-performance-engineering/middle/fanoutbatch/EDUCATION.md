# Fan Out Over Batches, Not Items

## Intuition

Split once, then run one goroutine per piece. The pieces are views into the same array, so splitting is nearly free.

## Approach

1. `Chunks` walks the input in steps of `size`, clamping the last batch.
2. `SumBatches` chunks, then launches one goroutine per batch writing to its own slot.

## Solution

```go
func Chunks(items []int, size int) [][]int {
	out := make([][]int, 0)
	if size <= 0 || len(items) == 0 {
		return out
	}
	for i := 0; i < len(items); i += size {
		out = append(out, items[i:min(i+size, len(items))])
	}
	return out
}

func SumBatches(items []int, size int) []int {
	batches := Chunks(items, size)
	out := make([]int, len(batches))
	var wg sync.WaitGroup
	for i, batch := range batches {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for _, v := range batch {
				sum += v
			}
			out[i] = sum
		}()
	}
	wg.Wait()
	return out
}
```

## Walkthrough

Since Go 1.22 each loop iteration gets its own `i` and `batch`, so the goroutines capture the right values without the old `i := i` dance — but the results still land in order only because each goroutine owns its slot.

## Pitfalls

- Copying each batch, which turns a free split into an O(n) one.
- Appending results from goroutines, which races and randomises the order.
- Batching so coarsely that one slow batch leaves every other worker idle.
