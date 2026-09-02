# Slugifier

## Intuition

The standard library's string functions return new strings and never mutate
their input. That purity is exactly what makes calling them from many goroutines
at once safe.

## Approach

1. Allocate `out := make([]string, len(titles))`.
2. Launch one goroutine per title, passing `i` and the title.
3. Split, join with hyphens, lowercase, and store at `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package slugifier — Gopher Workplace challenge.
package slugifier

import (
	"strings"
	"sync"
)

// Slugs turns every article title into a lowercase URL slug.
//
// Examples:
//
//	Slugs([]string{"Hello World"})  => [hello-world]
//	Slugs([]string{"   "})          => []
//	Slugs(nil)                      => []
func Slugs(titles []string) []string {
	out := make([]string, len(titles))
	var wg sync.WaitGroup
	for i, title := range titles {
		wg.Add(1)
		go func(i int, title string) {
			defer wg.Done()
			out[i] = strings.ToLower(strings.Join(strings.Fields(title), "-"))
		}(i, title)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"Hello World"` splits into two fields and joins to `"hello-world"`.
- `"Go   Is  Fast"` collapses the runs and becomes `"go-is-fast"`.
- `"   "` splits into zero fields, so the join yields `""`.

## Pitfalls

- Using `strings.Split(title, " ")`, which keeps empty fields and produces double hyphens.
- Building the slug with `+=` into a variable shared by all goroutines.
- Dropping blank titles instead of keeping their slot, which misaligns the result.
