# Email Normalizer

## Intuition

Because both string calls return new values rather than mutating anything, the
only shared memory in the whole function is the output slice — and each
goroutine owns exactly one slot of it.

## Approach

1. Allocate `out := make([]string, len(addrs))`.
2. Launch one goroutine per address, passing `i` and the address.
3. Write `strings.ToLower(strings.TrimSpace(addr))` to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package emailnormalizer — Gopher Workplace challenge.
package emailnormalizer

import (
	"strings"
	"sync"
)

// Normalize trims and lowercases every address in an import batch.
//
// Examples:
//
//	Normalize([]string{" A@X.io "})  => [a@x.io]
//	Normalize([]string{"   "})       => []
//	Normalize(nil)                   => []
func Normalize(addrs []string) []string {
	out := make([]string, len(addrs))
	var wg sync.WaitGroup
	for i, addr := range addrs {
		wg.Add(1)
		go func(i int, addr string) {
			defer wg.Done()
			out[i] = strings.ToLower(strings.TrimSpace(addr))
		}(i, addr)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `" A@X.io "` trims to `"A@X.io"` and lowercases to `"a@x.io"`.
- `"\tB@Y.IO\n"` shows that tabs and newlines are trimmed too.
- An all-whitespace entry becomes `""` but keeps its slot.

## Pitfalls

- Using `strings.Trim(addr, " ")`, which misses tabs and newlines.
- Writing back into `addrs`, mutating the caller's batch.
- Dropping blank entries, which misaligns the batch with its source rows.
