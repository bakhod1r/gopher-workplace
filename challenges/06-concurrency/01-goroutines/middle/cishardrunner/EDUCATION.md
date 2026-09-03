# CI Shard Runner

## Intuition

There are two ways to aggregate a fan-out: lock a shared accumulator, or give every worker a private slot and fold afterwards. The second is faster, race-free by construction, and — because you fold in index order — deterministic. Concurrency for the work, sequential code for the summary.

## Approach

1. Allocate `passes := make([]int, len(shards))` and `failed := make([]bool, len(shards))`.
2. Launch a goroutine per shard; on error set `failed[i] = true` and return without touching `passes[i]`.
3. Otherwise store the returned count in `passes[i]`.
4. After `wg.Wait()`, loop `i` in order summing `passes[i]` and appending failing indices.
5. Return the total and the failures slice, initialised to `[]int{}`.

## Solution

```go
// Package cishardrunner — Gopher Workplace challenge.
package cishardrunner

import "sync"

// RunShards runs every CI shard in its own goroutine and reports the total
// number of tests that passed together with the indices of the shards that
// failed, sorted ascending. A failing shard contributes no passes: a crashed
// runner cannot be trusted to have counted anything.
//
// Examples:
//
//	RunShards([][]string{{"a"}, {"b"}}, run)  => 2, []
//	RunShards([][]string{{}, {"b"}}, run)     => 1, [0]
//	RunShards(nil, run)                       => 0, []
func RunShards(shards [][]string, run func(shard []string) (passed int, err error)) (int, []int) {
	passes := make([]int, len(shards))
	failed := make([]bool, len(shards))

	var wg sync.WaitGroup
	for i, shard := range shards {
		wg.Add(1)
		go func(i int, shard []string) {
			defer wg.Done()
			n, err := run(shard)
			if err != nil {
				failed[i] = true
				return
			}
			passes[i] = n
		}(i, shard)
	}
	wg.Wait()

	total := 0
	failures := []int{}
	for i := range shards {
		total += passes[i]
		if failed[i] {
			failures = append(failures, i)
		}
	}
	return total, failures
}
```

## Walkthrough

- In `all_green` both goroutines store their counts and the fold gives `1 + 2 = 3`.
- The crashed shard in `one_crashed` returns `7, errShardCrashed`; `passes[1]` stays zero, so the total is still 3 and index 1 is reported.
- `crashed_shard_contributes_no_passes` isolates that rule: the runner offers 7 passes and the result is 0.
- `failures_sorted` has failures at 0 and 2 discovered in arbitrary goroutine order, yet the index-ordered fold emits `[0 2]`.

## Pitfalls

- `total += n` from inside the goroutines — a data race, and a wrong number even when it does not crash.
- Recording the count before checking the error, which folds a crashed runner's numbers into the build summary.
- Appending failing indices from inside the goroutines: the order then depends on scheduling.
- Returning a nil failures slice where the caller compares against `[]int{}`.
