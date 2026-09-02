# Storage Reporter

## Intuition

Goroutines nest naturally — a goroutine can start goroutines of its own. The
rule at every level is the same: wait for your own children before reporting
your result upwards.

## Approach

1. Allocate `out` and declare the outer WaitGroup on the parent.
2. Each outer goroutine allocates `kb := make([]int, len(files))` and its own inner WaitGroup.
3. Inner goroutines write `kb[j] = bytes / 1024`.
4. After `inner.Wait()`, the outer goroutine sums `kb` into `out[i]`; the parent then calls `outer.Wait()`.

## Solution

```go
// Package storagereporter — Gopher Workplace challenge.
package storagereporter

import (
	"sync"
)

// SectionKilobytes returns the total kilobytes stored in each report section.
//
// Examples:
//
//	SectionKilobytes([][]int{{2048, 1024}, {4096}})  => [3 4]
//	SectionKilobytes([][]int{{}})                    => [0]
//	SectionKilobytes(nil)                            => []
func SectionKilobytes(sections [][]int) []int {
	out := make([]int, len(sections))
	var outer sync.WaitGroup
	for i, files := range sections {
		outer.Add(1)
		go func(i int, files []int) {
			defer outer.Done()
			kb := make([]int, len(files))
			var inner sync.WaitGroup
			for j, bytes := range files {
				inner.Add(1)
				go func(j, bytes int) {
					defer inner.Done()
					kb[j] = bytes / 1024
				}(j, bytes)
			}
			inner.Wait()
			total := 0
			for _, v := range kb {
				total += v
			}
			out[i] = total
		}(i, files)
	}
	outer.Wait()
	return out
}
```

## Walkthrough

- For `{2048, 1024}` the inner goroutines write `2` and `1` into `kb`.
- `inner.Wait()` returns, the outer goroutine sums to `3` and stores it at `out[0]`.
- `outer.Wait()` returns only once every section has finished.

## Pitfalls

- Summing `kb` before `inner.Wait()` — you read zeros and the total is wrong.
- Sharing one WaitGroup across both levels, which makes the accounting meaningless.
- Declaring `kb` outside the outer goroutine, so every section scribbles on the same slice.
