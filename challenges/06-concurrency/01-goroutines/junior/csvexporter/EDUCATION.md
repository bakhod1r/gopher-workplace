# CSV Exporter

## Intuition

Not every value has to be passed as a parameter — only the ones that change as
the loop advances. A setting fixed before the loop is safe to capture.

## Approach

1. Allocate `out := make([]string, len(rows))`.
2. Launch one goroutine per record writing `strings.Join(row, sep)`.
3. `wg.Wait()` before returning.

## Solution

```go
// Package csvexporter — Gopher Workplace challenge.
package csvexporter

import (
	"strings"
	"sync"
)

// RenderRows renders every record as one delimited line.
//
// Examples:
//
//	RenderRows([][]string{{"a", "b"}}, ",")  => [a,b]
//	RenderRows([][]string{{}}, ",")          => []
//	RenderRows(nil, ",")                     => []
func RenderRows(rows [][]string, sep string) []string {
	out := make([]string, len(rows))
	var wg sync.WaitGroup
	for i, row := range rows {
		wg.Add(1)
		go func(i int, row []string) {
			defer wg.Done()
			out[i] = strings.Join(row, sep)
		}(i, row)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `{"a","b"}` with a comma renders as `"a,b"`.
- An empty record renders as `""` and still occupies its line.
- Records finish in any order but each lands in its own line number.

## Pitfalls

- Concatenating into a single shared buffer, which races and interleaves records.
- Skipping empty records, which shortens the export and shifts every later line.
- Capturing `row` from the loop variable rather than passing it in.
