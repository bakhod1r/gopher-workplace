# Backup Uploader

## Intuition

Concurrency does not make results unpredictable by itself — shared writes do.
With a private accumulator and a per-index write, the checksums are identical on
every run.

## Approach

1. Allocate `out := make([]int, len(parts))`.
2. In each goroutine, fold the part's bytes into a local `h`.
3. Write `out[i] = h` once, then `wg.Wait()`.

## Solution

```go
// Package backupuploader — Gopher Workplace challenge.
package backupuploader

import (
	"sync"
)

// PartChecksums returns a rolling checksum for every upload part.
//
// Examples:
//
//	PartChecksums([]string{"a", "b"})  => [97 98]
//	PartChecksums([]string{"ab"})      => [3105]
//	PartChecksums(nil)                 => []
func PartChecksums(parts []string) []int {
	out := make([]int, len(parts))
	var wg sync.WaitGroup
	for i, part := range parts {
		wg.Add(1)
		go func(i int, part string) {
			defer wg.Done()
			h := 0
			for _, b := range []byte(part) {
				h = h*31 + int(b)
			}
			out[i] = h
		}(i, part)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"a"` is byte 97, so `h = 0*31 + 97 = 97`.
- `"ab"` folds to `97*31 + 98 = 3105`.
- An empty part never enters the loop, so its checksum is `0`.

## Pitfalls

- Declaring `h` outside the goroutine, which makes every checksum both racy and wrong.
- Ranging over the string directly, which yields runes and changes the value for non-ASCII parts.
- Appending checksums instead of indexing them, losing the part order.
