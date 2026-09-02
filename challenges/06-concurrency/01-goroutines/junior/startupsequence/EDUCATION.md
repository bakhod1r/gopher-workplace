# Startup Sequence

## Intuition

Result order and completion order are different things. Storing by index means
the slowest check still prints on its own line, exactly where the operator
expects it.

## Approach

1. Allocate `out := make([]int, len(checks))`.
2. For each `i, check`, launch `go func(i int, check func() int)` that writes `out[i] = check()`.
3. `wg.Wait()`, then return `out`.

## Solution

```go
// Package startupsequence — Gopher Workplace challenge.
package startupsequence

import (
	"sync"
)

// RunChecks runs every preflight check and reports its status code in order.
//
// Examples:
//
//	RunChecks([]func() int{configOK, diskFull})  => [0 28]
//	RunChecks([]func() int{configOK})            => [0]
//	RunChecks(nil)                               => []
func RunChecks(checks []func() int) []int {
	out := make([]int, len(checks))
	var wg sync.WaitGroup
	for i, check := range checks {
		wg.Add(1)
		go func(i int, check func() int) {
			defer wg.Done()
			out[i] = check()
		}(i, check)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- Three checks returning `101`, `0` and `28` start in parallel.
- Whichever finishes first, each writes only its own index.
- The boot log reads `[101 0 28]` — configured order, not finish order.

## Pitfalls

- Collecting with `append`, which produces a randomly ordered boot log.
- Capturing `check` from the loop variable instead of passing it in.
- Calling `check()` on the parent goroutine, which serialises the whole sequence.
