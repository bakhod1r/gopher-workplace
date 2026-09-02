# Aspect Ratio

## Intuition

Unlike a slice, a `[2]int` array is copied on every assignment and call. Passing
one into a goroutine hands it a fully private value, with no backing array
shared with anyone.

## Approach

1. Allocate `out := make([]int, len(sizes))`.
2. In each goroutine, copy the pair into locals and take absolute values.
3. Run Euclid's loop, store `a` at `out[i]`, then `wg.Wait()`.

## Solution

```go
// Package aspectratio — Gopher Workplace challenge.
package aspectratio

import (
	"sync"
)

// Divisors returns the greatest common divisor of each width/height pair.
//
// Examples:
//
//	Divisors([][2]int{{1920, 1080}})  => [120]
//	Divisors([][2]int{{0, 720}})      => [720]
//	Divisors(nil)                     => []
func Divisors(sizes [][2]int) []int {
	out := make([]int, len(sizes))
	var wg sync.WaitGroup
	for i, size := range sizes {
		wg.Add(1)
		go func(i int, size [2]int) {
			defer wg.Done()
			a, b := size[0], size[1]
			if a < 0 {
				a = -a
			}
			if b < 0 {
				b = -b
			}
			for b != 0 {
				a, b = b, a%b
			}
			out[i] = a
		}(i, size)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `gcd(1920, 1080)` is `120`, which reduces the resolution to 16:9.
- `gcd(0, 720)` ends immediately with `720` after one swap.
- A square resolution reduces by its own side length, giving 1:1.

## Pitfalls

- Skipping the absolute value, which can return a negative divisor.
- Looping while `a != 0` instead of `b != 0`, which is off by one step.
- Assuming `[][2]int` elements alias like slices do; they do not.
