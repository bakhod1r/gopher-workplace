# Shard Backup Upload

## Intuition

Two goroutines writing different elements of the same slice do not race: each element is its own memory. That is what lets a fan-out collect results in input order with no lock and no channel — the index *is* the coordination.

## Approach

1. Allocate `errs := make([]error, len(shards))`.
2. For every shard, `wg.Add(1)` then launch a goroutine taking `i` and the shard as parameters.
3. Inside, `defer wg.Done()` and store `upload(sh)` into `errs[i]`.
4. `wg.Wait()`, then return `errs`.

## Solution

```go
// Package shardbackup — Gopher Workplace challenge.
package shardbackup

import "sync"

// Shard is one piece of a nightly database backup.
type Shard struct {
	ID   string
	Size int
}

// UploadShards uploads every shard in its own goroutine and returns one slot
// per shard, in input order: nil when that shard landed, the upload error when
// it did not. Every shard is attempted even after an earlier one fails, so a
// single bad object store key never truncates the backup set.
//
// Examples:
//
//	UploadShards([]Shard{{"a", 1}, {"b", 0}}, upload)  => [<nil> errEmptyShard]
//	UploadShards([]Shard{{"a", 1}}, upload)            => [<nil>]
//	UploadShards(nil, upload)                          => []
func UploadShards(shards []Shard, upload func(Shard) error) []error {
	errs := make([]error, len(shards))
	var wg sync.WaitGroup
	for i, sh := range shards {
		wg.Add(1)
		go func(i int, sh Shard) {
			defer wg.Done()
			errs[i] = upload(sh)
		}(i, sh)
	}
	wg.Wait()
	return errs
}
```

## Walkthrough

- With three healthy shards all three goroutines write `nil` and the caller sees a clean report.
- When shard `b` is empty only slot 1 holds an error; slots 0 and 2 still record their own outcome.
- The `first_failure_does_not_stop_rest` case proves the fan-out is not short-circuited: `upload` is still called three times.
- For a nil slice the loop body never runs, `wg.Wait()` returns immediately, and an empty non-nil slice comes back.

## Pitfalls

- `errs = append(errs, upload(sh))` from many goroutines both races and scrambles the order.
- Returning at the first error abandons the remaining goroutines' results — and the WaitGroup they still hold.
- Calling `wg.Add(1)` inside the goroutine: `Wait` may run before any `Add` and return instantly.
- Reading `errs` before `wg.Wait()` reports a half-finished backup as complete.
