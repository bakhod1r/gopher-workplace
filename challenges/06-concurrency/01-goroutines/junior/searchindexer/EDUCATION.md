# Search Indexer

## Intuition

Term frequency for one document says nothing about another. That independence,
plus a preallocated result, removes every reason to synchronise beyond the final
`Wait`.

## Approach

1. Allocate `out := make([]int, len(docs))`.
2. Launch one goroutine per document, passing `i` and `doc`.
3. Write `strings.Count(doc, term)` to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package searchindexer — Gopher Workplace challenge.
package searchindexer

import (
	"strings"
	"sync"
)

// TermCounts counts how often a term occurs in each document.
//
// Examples:
//
//	TermCounts([]string{"banana", "abc"}, "a")  => [3 1]
//	TermCounts([]string{"go go"}, "go")         => [2]
//	TermCounts(nil, "a")                        => []
func TermCounts(docs []string, term string) []int {
	out := make([]int, len(docs))
	var wg sync.WaitGroup
	for i, doc := range docs {
		wg.Add(1)
		go func(i int, doc string) {
			defer wg.Done()
			out[i] = strings.Count(doc, term)
		}(i, doc)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"banana"` contains `"a"` three times, `"abc"` once.
- `strings.Count("aaa", "aa")` is `1` — matches do not overlap.
- Every goroutine writes one index, so the posting list keeps document order.

## Pitfalls

- Expecting overlapping matches to be counted; they are not.
- Accumulating a shared corpus-wide total from the goroutines — that is a race.
- Passing `""` as the term, which counts runes plus one rather than zero.
