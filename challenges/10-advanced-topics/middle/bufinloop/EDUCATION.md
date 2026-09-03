# The Buffer Allocated Once Per Iteration

## Intuition

Allocating inside a loop is only wrong when the value does not need to be fresh. Here every row overwrites the buffer completely, so one buffer serves all of them — as long as its length is reset first.

## Approach

1. Move the `make` above the loop.
2. Reset with `buf = buf[:0]` at the top of each iteration.

## Solution

```go
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
// 	Render([][]int{{1, 2}}) => []string{"1,2"}
func Render(rows [][]int) []string {
	out := make([]string, 0, len(rows))
	buf := make([]byte, 0, scratchCap)
	for _, row := range rows {
		buf = buf[:0]
		for i, v := range row {
			if i > 0 {
				buf = append(buf, ',')
			}
			buf = strconv.AppendInt(buf, int64(v), 10)
		}
		out = append(out, string(buf))
	}
	return out
}
```

## Walkthrough

64 rows cost 64 scratch buffers before the fix and one after. The 64 output strings remain — those genuinely escape into the result.

## Pitfalls

- Hoisting without resetting, which concatenates every row onto the last.
- Trying to avoid `string(buf)` too; that copy is what makes the result independent of the buffer.
