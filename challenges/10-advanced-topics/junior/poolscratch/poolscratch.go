// Package poolscratch — Gopher Workplace challenge.
package poolscratch

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
//	Encode([]int{1, 2}) => "1,2"
func Encode(vals []int) string {
	panic("not implemented")
}
