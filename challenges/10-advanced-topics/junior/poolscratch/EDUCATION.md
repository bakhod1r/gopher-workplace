# Borrow A Buffer Instead Of Allocating One

## Intuition

A pool turns "allocate, use, drop" into "borrow, use, return". The catch is that a borrowed buffer is not empty — it is whatever the last borrower left, so you must reset its length before you write.

## Approach

1. `pool.Get().([]byte)` and reslice to `[:0]`.
2. Append the numbers and separators.
3. Convert to a string, put the buffer back, return the string.

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
// The scratch buffer used to build the text must come from the package's
// sync.Pool and go back into it, so repeated calls do not each allocate a
// buffer.
//
// Examples:
//
// 	Encode([]int{1, 2}) => "1,2"
func Encode(vals []int) string {
	buf := pool.Get().([]byte)[:0]
	for i, v := range vals {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(v), 10)
	}
	out := string(buf)
	pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
	return out
}
```

## Walkthrough

The first call finds the pool empty, so `New` makes a 64-byte buffer. Every later call reuses it: 100 calls, one buffer.

## Pitfalls

- Returning `buf` itself instead of a string — the caller would hold memory the pool is handing to someone else.
- Forgetting `[:0]`, which appends onto the previous call's output.
