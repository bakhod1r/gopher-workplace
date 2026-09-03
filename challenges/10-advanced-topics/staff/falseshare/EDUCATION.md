# Counters That Fight Over A Cache Line

## Intuition

Cache coherence works in lines, not variables. Two int64 counters eight bytes apart live on one 64-byte line, so every increment on one core invalidates the other's copy — the counters are logically private and physically shared.

## Approach

1. Reject non-positive input.
2. Allocate one padded `counter` per worker.
3. Start each goroutine with a pointer to its own slot; increment `iters` times.
4. `wg.Wait()`, then sum the slots.

## Solution

```go
import "sync"

// cacheLine is the coherence granule the counters must not share.
const cacheLine = 64

// counter is one worker's slot.
type counter struct {
	n   int64
	pad [cacheLine - 8]byte
}

// Count runs workers goroutines, each incrementing its own counter iters
// times, and returns the total.
//
// Each worker's counter must sit on its own cache line: adjacent counters
// put the cores into a write-invalidate storm over one line.
//
// Examples:
//
// 	Count(4, 1000) => 4000
func Count(workers, iters int) int64 {
	if workers < 1 || iters < 0 {
		return 0
	}
	cs := make([]counter, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(c *counter) {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				c.n++
			}
		}(&cs[w])
	}
	wg.Wait()
	var total int64
	for i := range cs {
		total += cs[i].n
	}
	return total
}
```

## Walkthrough

Eight unpadded counters fit in one line: 800000 increments become 800000 coherence transactions. With 64-byte padding each core owns its line and the increments stay in L1.

## Pitfalls

- Summing before `Wait` — the totals are read while the workers are still writing.
- Capturing the loop index in the closure instead of passing the slot pointer.
- Padding the slice instead of the element; the gap must be inside the struct.
