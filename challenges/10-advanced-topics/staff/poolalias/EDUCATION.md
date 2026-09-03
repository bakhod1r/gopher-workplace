# The Result That Still Belongs To The Pool

## Intuition

Returning a slice of a pooled buffer publishes memory you have just given away. The next `Get` hands the same array to another goroutine, which appends over the caller's result. Nothing is racy in the detector's sense — the write is simply legitimate and catastrophic.

## Approach

1. Build the text in the borrowed buffer as before.
2. Copy it into a right-sized slice of its own.
3. Return the buffer to the pool and return the copy.

## Solution

```go
import (
	"strconv"
	"sync"
)

// pool hands out reusable scratch buffers.
var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}

// Encode returns vals rendered as decimal numbers joined by ','.
//
// The scratch buffer is borrowed from a pool and returned before Encode
// exits, so the result may not be a view of it: the next borrower would
// overwrite the caller's data.
//
// Examples:
//
// 	Encode([]int{1, 2}) => []byte("1,2")
func Encode(vals []int) []byte {
	buf := pool.Get().([]byte)[:0]
	for i, v := range vals {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(v), 10)
	}
	out := make([]byte, len(buf))
	copy(out, buf)
	pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
	return out
}
```

## Walkthrough

Encode([]int{111,222}) writes seven bytes and returns a view. `Put` releases the array; the next call resets it to len 0 and writes "999,888" — and the first caller's slice now reads "999,888".

## Pitfalls

- `string(buf)` then converting back — an extra copy for the same result.
- Putting the buffer back in a `defer` and returning the view anyway; the ordering does not save you.
