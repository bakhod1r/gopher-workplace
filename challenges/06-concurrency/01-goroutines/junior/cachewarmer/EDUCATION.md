# Cache Warmer

## Intuition

Concurrency is only dangerous where writes are shared. Here every goroutine
reads the same `capBytes` and writes a slot nobody else touches, so no locking
is needed at all.

## Approach

1. Allocate `out := make([]int, len(keys))`.
2. Launch one goroutine per key, passing `i` and `key`.
3. Load into a local `size`, clamp it to `[0, capBytes]`, then store it at `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package cachewarmer — Gopher Workplace challenge.
package cachewarmer

import (
	"sync"
)

// WarmAll loads every key and reports the cached size, clamped to the entry cap.
//
// Examples:
//
//	WarmAll([]string{"a", "bb"}, loader, 1000)  => [100 200]
//	WarmAll([]string{"huge"}, loader, 150)      => [150]
//	WarmAll(nil, loader, 100)                   => []
func WarmAll(keys []string, load func(key string) int, capBytes int) []int {
	out := make([]int, len(keys))
	var wg sync.WaitGroup
	for i, key := range keys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			size := load(key)
			if size < 0 {
				size = 0
			}
			if size > capBytes {
				size = capBytes
			}
			out[i] = size
		}(i, key)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"a"` loads 100 bytes, which is under the cap and passes through.
- `"huge"` loads 400 bytes against a 150-byte cap, so it is recorded as `150`.
- `"missing"` loads `-1`, which is raised to `0`.

## Pitfalls

- Accumulating a shared `totalBytes` counter from goroutines — that is a data race.
- Clamping on the parent after `Wait`, which works but hides the per-key logic.
- Assuming the cap is always larger than the payload; the tests exercise both.
