// Package bufferreset — Gopher Workplace challenge.
package bufferreset

import (
	"bytes"
	"strconv"
	"sync"
)

var (
	mu      sync.Mutex
	scratch bytes.Buffer
)

// Render returns vals as decimal numbers joined by '-'.
//
// The package keeps one scratch buffer to avoid allocating per call. A
// shared buffer has to be emptied before it is written to.
//
// Examples:
//
//	Render([]int{1, 2}) => "1-2"
func Render(vals []int) string {
	// CHANGE CODE BELOW THIS LINE
	mu.Lock()
	defer mu.Unlock()
	for i, v := range vals {
		if i > 0 {
			scratch.WriteByte('-')
		}
		scratch.WriteString(strconv.Itoa(v))
	}
	return scratch.String()
	// CHANGE CODE ABOVE THIS LINE
}
