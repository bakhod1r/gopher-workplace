# Pad To The Line, Not To The Word

## Intuition

The hardware does not know about your variables, only about lines. Making each counter occupy a whole line means each core owns its line outright, and the increments never leave L1.

## Approach

1. Validate the arguments.
2. Allocate `workers` slots and start one goroutine per slot, passing its address.
3. Increment `s.N` `iters` times; `Wait`, then sum the slots.

## Solution

```go
import (
	"sync"
	"unsafe"
)

// LineSize is the coherence granule the counters must not share.
const LineSize = 64

// Slot is one worker's counter, padded to its own cache line.
type Slot struct {
	N   int64
	Pad [LineSize - unsafe.Sizeof(int64(0))]byte
}

// Run gives each of workers goroutines its own Slot, has each increment
// its counter iters times, and returns the total.
//
// Slot must be padded so no two counters share a cache line; the padding is
// computed from the counter's own size, not hard-coded.
//
// Examples:
//
// 	Run(4, 1000) => 4000
func Run(workers, iters int) int64 {
	if workers < 1 || iters < 0 {
		return 0
	}
	slots := make([]Slot, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(s *Slot) {
			defer wg.Done()
			for i := 0; i < iters; i++ {
				s.N++
			}
		}(&slots[w])
	}
	wg.Wait()
	var total int64
	for i := range slots {
		total += slots[i].N
	}
	return total
}
```

## Walkthrough

Eight unpadded int64 counters fit in one 64-byte line, so 1.6 million increments become 1.6 million coherence transactions. With one slot per line each core writes only its own.

## Pitfalls

- Capturing the loop index and indexing `slots[w]` inside the goroutine — correct, but passing the pointer makes the exclusive ownership explicit.
- Writing `Pad [56]byte`; the constant expression keeps the padding right if `N` changes type.
- Summing before `Wait`.
