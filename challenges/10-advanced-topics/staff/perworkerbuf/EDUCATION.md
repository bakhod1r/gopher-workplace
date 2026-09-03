# Give Every Worker Its Own Frame

## Intuition

Per-goroutine state does not have to be heap state. A fixed-size array declared inside the goroutine lives in that goroutine's frame — private without a `make`, and free without a pool.

## Approach

1. Preallocate the result slice to `len(rows)`.
2. Start one goroutine per row, passing `i` and `row` as parameters.
3. Declare `var scratch [64]byte` inside the goroutine and build into `scratch[:0]`.
4. Write `out[i]`, then `wg.Wait()`.

## Solution

```go
import (
	"strconv"
	"sync"
)

// RenderAll renders each row concurrently as comma-separated decimals
// and returns the results in input order.
//
// Each goroutine's scratch buffer must be a local that does not escape:
// one shared buffer is a race, and one heap buffer per row is garbage.
//
// Examples:
//
// 	RenderAll([][]int{{1, 2}}) => []string{"1,2"}
func RenderAll(rows [][]int) []string {
	out := make([]string, len(rows))
	var wg sync.WaitGroup
	wg.Add(len(rows))
	for i, row := range rows {
		go func(i int, row []int) {
			defer wg.Done()
			var scratch [64]byte
			buf := scratch[:0]
			for j, v := range row {
				if j > 0 {
					buf = append(buf, ',')
				}
				buf = strconv.AppendInt(buf, int64(v), 10)
			}
			out[i] = string(buf)
		}(i, row)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

Each goroutine renders into its own 64-byte frame array; only `string(buf)` allocates, which is the result the caller keeps. With 128 rows that is 129 allocations instead of 257.

## Pitfalls

- Rows longer than the scratch array make `append` allocate — correct, just no longer free.
- Hoisting the array above the loop, which is the race the fix was for.
