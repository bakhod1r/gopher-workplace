# One Scratch Buffer, Many Goroutines

## Intuition

Passing the loop variables in was half the job. The scratch buffer is still one variable captured by every goroutine, so all of them reset it, append to it and read it at once. `out[i]` is fine — those writes are disjoint.

## Approach

1. Move the buffer's declaration inside the goroutine.
2. Leave the rest as is: each goroutine builds its own text and writes its own slot.

## Solution

```go
import (
	"strconv"
	"sync"
)

// EncodeAll renders every batch concurrently and returns the results in
// input order.
//
// Each goroutine must work in storage of its own; a buffer captured from
// the enclosing scope is shared by all of them.
//
// Examples:
//
// 	EncodeAll([][]int{{1}, {2}}) => []string{"1", "2"}
func EncodeAll(batches [][]int) []string {
	out := make([]string, len(batches))
	var wg sync.WaitGroup
	wg.Add(len(batches))
	for i, b := range batches {
		go func(i int, b []int) {
			defer wg.Done()
			buf := make([]byte, 0, 64)
			for j, v := range b {
				if j > 0 {
					buf = append(buf, ',')
				}
				buf = strconv.AppendInt(buf, int64(v), 10)
			}
			out[i] = string(buf)
		}(i, b)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

With a shared buffer, goroutine 3 can reset it to `[:0]` between goroutine 7's append and its `string(buf)`, so result 7 comes out truncated or holding batch 3's digits. A per-goroutine buffer makes the interleaving irrelevant.

## Pitfalls

- Adding a mutex around the buffer — correct, and it serialises the whole fan-out.
- Assuming `-race` will always catch it; it reports only interleavings it actually observes.
