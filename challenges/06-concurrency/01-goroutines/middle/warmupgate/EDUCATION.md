# Cache Warm-Up Gate

## Intuition

"Return the first error" is a trap in concurrent code: whichever goroutine happens to finish first is not a property of your program, it is a property of the scheduler that day. Pinning "first" to the input index turns a flaky error message into a reproducible one.

## Approach

1. Allocate `errs := make([]error, len(keys))`.
2. Launch a goroutine per key writing into its own slot.
3. `wg.Wait()` so every warm has completed.
4. Scan `errs` in index order and return the first non-nil error, else nil.

## Solution

```go
// Package warmupgate — Gopher Workplace challenge.
package warmupgate

import "sync"

// WarmCaches warms every key concurrently and returns a single error for the
// deploy gate: the failure belonging to the lowest-indexed key that could not be
// warmed, or nil when the whole set is hot. Every warm is allowed to finish
// first, so the pod never starts serving with half-warm caches and never leaks
// a goroutine into the next release.
//
// Examples:
//
//	WarmCaches([]string{"tags", "prices"}, warm)  => <nil>
//	WarmCaches([]string{"tags", "bad"}, warm)     => errColdSource
//	WarmCaches(nil, warm)                         => <nil>
func WarmCaches(keys []string, warm func(key string) error) error {
	errs := make([]error, len(keys))
	var wg sync.WaitGroup
	for i, key := range keys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			errs[i] = warm(key)
		}(i, key)
	}
	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}
```

## Walkthrough

- With all keys warm the scan finds nothing and returns nil.
- In `lowest_index_wins` both `cold-a` and `cold-z` fail; the index scan always returns `cold-a`, whichever goroutine finished first.
- `cold_in_middle` shows the rule is about index order, not about position of the first success.
- The call-count assertion proves the fan-out is not short-circuited: every key is still attempted.

## Pitfalls

- Writing to a shared `var firstErr error` from the goroutines — a race, and the winner varies per run.
- Returning as soon as one goroutine reports failure, leaving the rest running after the function returned.
- Wrapping the error in a new `errors.New` string, which breaks `errors.Is` for the caller.
- Scanning `errs` before `wg.Wait()` and reading slots that are still empty.
