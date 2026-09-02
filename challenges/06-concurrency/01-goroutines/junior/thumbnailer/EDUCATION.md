# Thumbnailer

## Intuition

`defer wg.Done()` is not decoration — it is what makes every exit path from the
goroutine, including an early `return`, still count down the group.

## Approach

1. Allocate `out := make([]int, len(images))`.
2. Launch one goroutine per image, passing `i` and the `Image` value.
3. Guard `img.Width <= 0`, otherwise write the scaled height to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package thumbnailer — Gopher Workplace challenge.
package thumbnailer

import (
	"sync"
)

// Image is a source image with its pixel dimensions.
type Image struct {
	Width  int
	Height int
}

// TargetHeights returns the proportional height of every image scaled to maxWidth.
//
// Examples:
//
//	TargetHeights([]Image{{100, 50}, {200, 200}}, 100)  => [50 100]
//	TargetHeights([]Image{{0, 40}}, 100)                => [0]
//	TargetHeights(nil, 100)                             => []
func TargetHeights(images []Image, maxWidth int) []int {
	out := make([]int, len(images))
	var wg sync.WaitGroup
	for i, img := range images {
		wg.Add(1)
		go func(i int, img Image) {
			defer wg.Done()
			if img.Width <= 0 {
				out[i] = 0
				return
			}
			out[i] = img.Height * maxWidth / img.Width
		}(i, img)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- A 100×50 image scaled to width 100 keeps height `50`.
- A 50×20 image scaled to width 200 becomes `20 * 200 / 50 = 80`.
- A zero-width image would divide by zero, so the goroutine returns `0` early.

## Pitfalls

- Dividing before multiplying — integer truncation then loses the ratio.
- Calling `wg.Done()` as the last statement instead of deferring it; the early return deadlocks `Wait`.
- Passing `*Image` and mutating the caller's data instead of returning new values.
