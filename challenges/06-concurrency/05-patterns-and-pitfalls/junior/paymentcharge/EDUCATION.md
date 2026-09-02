# Aggregating Failures Concurrently

## Intuition

`append` may reallocate and always rewrites the slice header, so two goroutines
appending at once can lose an element or corrupt the slice. One small mutex
around the append makes the collection safe without serialising the charges.

## Approach

1. Declare `wg`, `mu`, and the `msgs` slice.
2. Per provider, start a goroutine that calls `charge` with no lock held and returns early on success.
3. Lock, append `err.Error()`, unlock; after `wg.Wait()` call `sort.Strings` and return.

## Solution

```go
import (
	"sort"
	"sync"
)

func ChargeAll(providers []string, charge func(string) error) []string {
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		msgs []string
	)

	for _, provider := range providers {
		wg.Add(1)
		go func(provider string) {
			defer wg.Done()

			err := charge(provider)
			if err == nil {
				return
			}

			mu.Lock()
			msgs = append(msgs, err.Error())
			mu.Unlock()
		}(provider)
	}

	wg.Wait()
	sort.Strings(msgs)
	return msgs
}
```

## Walkthrough

For `bad-z` and `bad-a`, both goroutines fail and append their messages in
whatever order they finish. `sort.Strings` then puts `bad-a declined` first,
so the report is identical on every run.

## Pitfalls

- Appending without the mutex — the race detector flags it and messages go missing.
- Holding the lock across `charge`, which turns parallel charges into sequential ones.
- Returning unsorted results: the same failures produce a different report each run.
