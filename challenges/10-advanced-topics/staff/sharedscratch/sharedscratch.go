// Package sharedscratch — Gopher Workplace challenge.
package sharedscratch

import (
	"strconv"
	"sync"
)

// EncodeAll renders every batch concurrently and returns the results in
// input order.
//
// Each goroutine must work in storage of its own; a buffer captured from
// the enclosing scope is shared by all of them.
//
// Examples:
//
//	EncodeAll([][]int{{1}, {2}}) => []string{"1", "2"}
func EncodeAll(batches [][]int) []string {
	// CHANGE CODE BELOW THIS LINE
	out := make([]string, len(batches))
	buf := make([]byte, 0, 64)
	var wg sync.WaitGroup
	wg.Add(len(batches))
	for i, b := range batches {
		go func(i int, b []int) {
			defer wg.Done()
			buf = buf[:0]
			for j, v := range b {
				if j > 0 {
					buf = append(buf, ',')
				}
				buf = strconv.AppendInt(buf, int64(v), 10)
			}
			out[i] = string(buf)
		}(i, b)
	}
	wg.Wait()
	return out
	// CHANGE CODE ABOVE THIS LINE
}
