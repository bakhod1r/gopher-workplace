# Log Shipper

## Intuition

Measuring one line tells you nothing about the next, so the batch decomposes
perfectly. Preallocating the result gives every goroutine a private destination.

## Approach

1. Allocate `out := make([]int, len(lines))`.
2. Loop with `for i, line := range lines`, launching one goroutine per pair.
3. `wg.Wait()`, then return `out`.

## Solution

```go
// Package logshipper — Gopher Workplace challenge.
package logshipper

import (
	"sync"
)

// PayloadSizes returns the wire size of every log line, including its newline.
//
// Examples:
//
//	PayloadSizes([]string{"ok", "boom"})  => [3 5]
//	PayloadSizes([]string{""})            => [1]
//	PayloadSizes(nil)                     => []
func PayloadSizes(lines []string) []int {
	out := make([]int, len(lines))
	var wg sync.WaitGroup
	for i, line := range lines {
		wg.Add(1)
		go func(i int, line string) {
			defer wg.Done()
			out[i] = len(line) + 1
		}(i, line)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"ok"` is 2 bytes, so its wire size is `3`.
- A blank line still costs `1` byte for the newline.
- `"añ"` is 3 bytes (ñ is two), so its wire size is `4`.

## Pitfalls

- Using a rune count, which under-reports the bytes actually sent.
- Summing into one shared total instead of writing per index.
- Forgetting `defer wg.Done()`, which deadlocks `Wait`.
