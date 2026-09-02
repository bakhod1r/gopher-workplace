# Progress Bar

## Intuition

Unrecoverable failure is more expensive in concurrent code: the caller cannot
wrap a goroutine in `recover`, so the guard has to live inside the goroutine
itself.

## Approach

1. Allocate `out := make([]string, len(percents))`.
2. Inside each goroutine, clamp the percentage into `[0, 100]`.
3. Compute the filled length and write `strings.Repeat("#", filled)` to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package progressbar — Gopher Workplace challenge.
package progressbar

import (
	"strings"
	"sync"
)

// Bars renders a fixed-width bar for every job's completion percentage.
//
// Examples:
//
//	Bars([]int{50, 100}, 10)  => [##### ##########]
//	Bars([]int{0}, 10)        => []
//	Bars(nil, 10)             => []
func Bars(percents []int, width int) []string {
	out := make([]string, len(percents))
	var wg sync.WaitGroup
	for i, pct := range percents {
		wg.Add(1)
		go func(i int, pct int) {
			defer wg.Done()
			if pct < 0 {
				pct = 0
			}
			if pct > 100 {
				pct = 100
			}
			filled := pct * width / 100
			if filled < 0 {
				filled = 0
			}
			out[i] = strings.Repeat("#", filled)
		}(i, pct)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `50` of a 10-wide bar fills `50 * 10 / 100 = 5` characters.
- `150` clamps to `100` and fills the bar completely; `-20` clamps to `0`.
- A 4-wide bar at 25% fills a single character.

## Pitfalls

- Calling `strings.Repeat` with a negative count, which panics and kills the process.
- Trying to `recover` in the caller — a goroutine's panic is not catchable there.
- Dividing before multiplying, which truncates every partial bar to empty.
