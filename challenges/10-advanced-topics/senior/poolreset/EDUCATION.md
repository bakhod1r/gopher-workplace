# The Pooled Buffer Nobody Emptied

## Intuition

A pool recycles values, not blank slates. The buffer arrives with the previous borrower's length, so appending continues their output — the data leaks forward and the buffer grows without bound.

## Approach

1. Reslice the borrowed buffer to `[:0]` before writing.
2. Append as before, convert to a string, return the buffer to the pool.

## Solution

```go
import (
	"strconv"
	"sync"
)

// pool hands out reusable scratch buffers.
var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}

// Render returns vals as decimal numbers joined by ','.
//
// The scratch buffer comes from a sync.Pool and goes back after use. A
// buffer that comes out of a pool carries whatever the last borrower left
// in it.
//
// Examples:
//
// 	Render([]int{1, 2}) => "1,2"
func Render(vals []int) string {
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

First call: `New` gives len 0, output "7", buffer returned at len 1. Second call: len 1, so the output is "77". By call 200 the buffer holds 200 sevens.

## Pitfalls

- Resetting after use instead of before — another goroutine may already have taken it.
- Putting a buffer back while the caller still holds a slice of it.
