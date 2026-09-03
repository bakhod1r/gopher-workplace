# The Pool That Kept Every Oversized Buffer

## Intuition

A pool is a cache of allocations, and a cache with no eviction rule keeps its worst entry forever. One outlier request permanently raises the memory floor unless the return path checks the size.

## Approach

1. Fill the buffer as before.
2. Return it to the pool only when `cap(buf) <= maxScratch`.

## Solution

```go
import "sync"

// maxScratch is the largest buffer worth keeping in the pool.
const maxScratch = 4096

var pool = sync.Pool{New: func() any { return make([]byte, 0, 64) }}

// PooledCap reports the capacity of a buffer currently in the pool, or 0.
func PooledCap() int {
	v := pool.Get()
	if v == nil {
		return 0
	}
	b := v.([]byte)
	c := cap(b)
	pool.Put(b) //nolint:staticcheck // the puzzle keeps the pool API simple
	return c
}

// Render borrows a scratch buffer, fills size bytes of it, returns the
// buffer to the pool and reports how many bytes it wrote.
//
// Occasional huge requests must not leave the pool holding huge buffers
// forever: a buffer larger than maxScratch is dropped instead of returned.
//
// Examples:
//
// 	Render(16) => 16
func Render(size int) int {
	if size < 0 {
		size = 0
	}
	buf := pool.Get().([]byte)[:0]
	for i := 0; i < size; i++ {
		buf = append(buf, byte(i))
	}
	n := len(buf)
	if cap(buf) <= maxScratch {
		pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
	}
	return n
}
```

## Walkthrough

A 1 MiB render grows the borrowed buffer to at least 1 MiB. Putting it back leaves that megabyte pinned by the pool; dropping it lets the collector reclaim it and the next `Get` starts from 64 bytes again.

## Pitfalls

- Checking `len(buf)` instead of `cap(buf)` — the length is zeroed on the next borrow.
- Dropping every buffer, which disables the pool entirely.
