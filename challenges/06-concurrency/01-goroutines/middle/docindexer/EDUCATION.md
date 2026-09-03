# Partial Document Indexing

## Intuition

Every concurrent aggregation has a seam where parallel work becomes a single answer. Put the seam as late as possible — one `[]bool` written by disjoint indices — and everything after it is ordinary sequential code you can reason about and sort.

## Approach

1. Allocate `ok := make([]bool, len(docs))`.
2. Launch a goroutine per document storing `index(doc) == nil` into `ok[i]`.
3. `wg.Wait()`.
4. Initialise both results to empty slices and walk the docs in order, appending each ID to the matching list.
5. Sort both lists and return them.

## Solution

```go
// Package docindexer — Gopher Workplace challenge.
package docindexer

import (
	"sort"
	"sync"
)

// Doc is one document queued for the search index.
type Doc struct {
	ID   string
	Body string
}

// IndexDocuments indexes every document concurrently and splits the outcome in
// two: the IDs that made it into the index and the IDs that must be replayed.
// Both lists come back sorted. A reindex job is judged by these two lists, so
// neither may depend on which goroutine finished first.
//
// Examples:
//
//	IndexDocuments([]Doc{{"a", "x"}}, index)              => ["a"], []
//	IndexDocuments([]Doc{{"a", "x"}, {"b", ""}}, index)   => ["a"], ["b"]
//	IndexDocuments(nil, index)                            => [], []
func IndexDocuments(docs []Doc, index func(Doc) error) (indexed []string, failed []string) {
	ok := make([]bool, len(docs))

	var wg sync.WaitGroup
	for i, doc := range docs {
		wg.Add(1)
		go func(i int, doc Doc) {
			defer wg.Done()
			ok[i] = index(doc) == nil
		}(i, doc)
	}
	wg.Wait()

	indexed, failed = []string{}, []string{}
	for i, doc := range docs {
		if ok[i] {
			indexed = append(indexed, doc.ID)
		} else {
			failed = append(failed, doc.ID)
		}
	}
	sort.Strings(indexed)
	sort.Strings(failed)
	return indexed, failed
}
```

## Walkthrough

- In `both_lists_sorted` the goroutines finish in arbitrary order, yet the partition walks the input and the sorts fix the final order to `[alpha zed]` and `[beta yank]`.
- `all_rejected` shows the indexed list surviving as an empty slice rather than nil.
- The call counter proves no document was skipped because of a neighbour's rejection.
- For a nil input no goroutine starts and both empty slices are returned.

## Pitfalls

- Appending to the two lists from the goroutines without a mutex — a race on two slices instead of one.
- Adding a mutex but still not sorting, which leaves the report order scheduler-dependent.
- Leaving one of the lists nil, so the job report serialises `null` instead of `[]`.
- Treating a non-nil error as a reason to stop the batch, which loses the documents that would have indexed.
