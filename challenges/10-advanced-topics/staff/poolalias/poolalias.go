// Package poolalias — Gopher Workplace challenge.
package poolalias

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
//	Encode([]int{1, 2}) => []byte("1,2")
func Encode(vals []int) []byte {
	// CHANGE CODE BELOW THIS LINE
	buf := pool.Get().([]byte)[:0]
	for i, v := range vals {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(v), 10)
	}
	out := buf
	pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
	return out
	// CHANGE CODE ABOVE THIS LINE
}
