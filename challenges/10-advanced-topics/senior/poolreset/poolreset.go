// Package poolreset — Gopher Workplace challenge.
package poolreset

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
//	Render([]int{1, 2}) => "1,2"
func Render(vals []int) string {
	// CHANGE CODE BELOW THIS LINE
	buf := pool.Get().([]byte)
	for i, v := range vals {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = strconv.AppendInt(buf, int64(v), 10)
	}
	out := string(buf)
	pool.Put(buf) //nolint:staticcheck // the puzzle keeps the pool API simple
	return out
	// CHANGE CODE ABOVE THIS LINE
}
