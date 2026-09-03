# One Builder, Many Lines

## Intuition

A builder's value is the buffer it accumulates. Constructing one per row throws that buffer away every time — the reuse is the entire point of the type.

## Approach

1. Preallocate the result slice.
2. Estimate the widest row and `Grow` the builder once.
3. Per row: `Reset`, write the values with separators, append `b.String()`.

## Solution

```go
import (
	"strconv"
	"strings"
)

// RenderLines renders each row as its values joined by '-'.
//
// The builder is per-call state: reset it between rows instead of
// constructing one per row, and reserve its capacity once.
//
// Examples:
//
// 	RenderLines([][]int{{1, 2}}) => []string{"1-2"}
func RenderLines(rows [][]int) []string {
	out := make([]string, 0, len(rows))
	widest := 0
	for _, row := range rows {
		if n := len(row) * 12; n > widest {
			widest = n
		}
	}
	var b strings.Builder
	b.Grow(widest)
	for _, row := range rows {
		b.Reset()
		for i, v := range row {
			if i > 0 {
				b.WriteByte('-')
			}
			b.WriteString(strconv.Itoa(v))
		}
		out = append(out, b.String())
	}
	return out
}
```

## Walkthrough

For 64 rows of three values, one builder is grown once and reset 64 times. Only the 64 result strings and the result slice are allocated.

## Pitfalls

- Forgetting `Reset`, which concatenates every row onto the previous ones.
- Reusing a builder across calls — `strings.Builder` must not be copied, and sharing one is not concurrency-safe.
