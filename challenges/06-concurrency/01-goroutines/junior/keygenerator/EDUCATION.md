# Key Generator

## Intuition

Primality of one number is completely independent of another, which makes the
pool embarrassingly parallel. Each goroutine reduces its candidate to one
boolean at one index.

## Approach

1. Allocate `out := make([]bool, len(candidates))`.
2. In each goroutine, set `prime := n >= 2` and trial-divide up to the square root.
3. Store the verdict at `out[i]`, then `wg.Wait()`.

## Solution

```go
// Package keygenerator — Gopher Workplace challenge.
package keygenerator

import (
	"sync"
)

// PrimeCandidates reports which candidates survive a primality test.
//
// Examples:
//
//	PrimeCandidates([]int{2, 4, 7})   => [true false true]
//	PrimeCandidates([]int{1, 0, -5})  => [false false false]
//	PrimeCandidates(nil)              => []
func PrimeCandidates(candidates []int) []bool {
	out := make([]bool, len(candidates))
	var wg sync.WaitGroup
	for i, n := range candidates {
		wg.Add(1)
		go func(i int, n int) {
			defer wg.Done()
			prime := n >= 2
			for d := 2; d*d <= n; d++ {
				if n%d == 0 {
					prime = false
					break
				}
			}
			out[i] = prime
		}(i, n)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `7` survives division by `2` because `3*3 > 7`, so it stays `true`.
- `9` is divisible by `3`, so the verdict flips to `false` and the loop breaks.
- `1`, `0` and `-5` start as `false` thanks to the `n >= 2` seed.

## Pitfalls

- Declaring `prime` outside the goroutine, making every verdict shared and racy.
- Forgetting that `1` is not prime.
- Looping `d < n`, which is correct but does far more work than needed.
