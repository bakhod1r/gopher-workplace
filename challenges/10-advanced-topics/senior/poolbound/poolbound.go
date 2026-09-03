// Package poolbound — Gopher Workplace challenge.
package poolbound

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
//	Render(16) => 16
func Render(size int) int {
	// CHANGE CODE BELOW THIS LINE
	if size < 0 {
		size = 0
	}
	buf := pool.Get().([]byte)[:0]
	for i := 0; i < size; i++ {
		buf = append(buf, byte(i))
	}
	n := len(buf)
	pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
	return n
	// CHANGE CODE ABOVE THIS LINE
}
