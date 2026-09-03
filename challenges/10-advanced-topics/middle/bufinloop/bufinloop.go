// Package bufinloop — Gopher Workplace challenge.
package bufinloop

import "strconv"

// scratchCap is the scratch buffer's capacity. It is a variable, so the
// compiler cannot prove the buffer's size and must allocate it on the heap.
var scratchCap = 64

// Render turns each row into a comma-separated string.
//
// The scratch buffer is per-call state, not per-row state: allocating it
// inside the loop makes one throwaway buffer for every row.
//
// Examples:
//
//	Render([][]int{{1, 2}}) => []string{"1,2"}
func Render(rows [][]int) []string {
	// CHANGE CODE BELOW THIS LINE
	out := make([]string, 0, len(rows))
	for _, row := range rows {
		buf := make([]byte, 0, scratchCap)
		for i, v := range row {
			if i > 0 {
				buf = append(buf, ',')
			}
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		out = append(out, string(buf))
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
