# CI Shard Runner

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

CI splits the test suite into shards that run on separate machines in parallel. The build summary needs two numbers that must agree: how many tests passed overall, and exactly which shards failed. A shard whose runner crashed may still have printed a count, but that count is not evidence — it is discarded.

## Task

Implement the exported function(s) in [cishardrunner.go](cishardrunner.go) so that:

1. Run each shard in its own goroutine, joined with a `sync.WaitGroup`.
2. Sum the `passed` counts of shards that returned a nil error.
3. Ignore the count of any shard that returned an error.
4. Return the indices of failing shards in ascending order.
5. Return `0` and an empty non-nil slice when there are no shards.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  RunShards([][]string{{"a"}, {"b", "c"}}, run)
Output: 3, []
```

**Example 2:**

```
Input:  RunShards([][]string{{"a"}, {}, {"c", "d"}}, run)  // shard 1 crashes
Output: 3, [1]
```

**Example 3:**

```
Input:  RunShards(nil, run)
Output: 0, []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Aggregating across goroutines** | Write per-index, then fold sequentially after `Wait` — no lock, and the fold order is deterministic. |
| 2 | **Partial failure semantics** | A crashed worker's numbers are discarded rather than merged into the total. |
| 3 | **Ordering without sorting** | Walking the index range in order produces sorted failures for free. |
| 4 | **Multiple return values** | One goroutine can report both a count and an error into separate per-index slots. |

## Hint

Give each goroutine two private slots — a count and a failure flag. Fold them into the totals in the parent, after `wg.Wait()`.

## Validate

```bash
make verify
```
