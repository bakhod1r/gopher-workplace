# DNS Resolver

## Intuition

Position in the result is the contract here: row `i` of the table must describe
host `i`. Writing by index gives you that for free, no matter which lookup
returns first.

## Approach

1. Allocate `out := make([]string, len(hosts))`.
2. Launch one goroutine per host, passing `i` and `host`.
3. Resolve into a local `addr`, replace `""` with `"0.0.0.0"`, then write `out[i] = addr`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package dnsresolver — Gopher Workplace challenge.
package dnsresolver

import (
	"sync"
)

// ResolveAll resolves every host, substituting a placeholder for failures.
//
// Examples:
//
//	ResolveAll([]string{"a.io", "b.io"}, lookup)  => [10.0.0.1 10.0.0.2]
//	ResolveAll([]string{"ghost.io"}, lookup)      => [0.0.0.0]
//	ResolveAll(nil, lookup)                       => []
func ResolveAll(hosts []string, resolve func(host string) string) []string {
	out := make([]string, len(hosts))
	var wg sync.WaitGroup
	for i, host := range hosts {
		wg.Add(1)
		go func(i int, host string) {
			defer wg.Done()
			addr := resolve(host)
			if addr == "" {
				addr = "0.0.0.0"
			}
			out[i] = addr
		}(i, host)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"a.io"` resolves to `"10.0.0.1"` and lands at index 0.
- `"ghost.io"` resolves to `""`, so the goroutine substitutes `"0.0.0.0"`.
- Duplicated hosts each get their own goroutine and their own slot.

## Pitfalls

- Skipping failed lookups instead of filling the slot — the table then misaligns with the config.
- Appending to a shared slice from goroutines, which races and randomises the order.
- Capturing `host` from the loop variable rather than passing it in.
