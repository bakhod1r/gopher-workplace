# Test Sharder

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A CI system splits the test suite across parallel machines. Each machine gets a
consecutive block of `perShard` tests, and the scheduler needs to know how long
each block will take so it can spot the straggler shard. Blocks are independent,
so each total is computed in its own goroutine.

## Task

Implement `ShardDurations` in [testsharder.go](testsharder.go) so that:

1. Return `nil` when `perShard <= 0`.
2. Split `durations` into consecutive shards of `perShard` tests; the last shard may be shorter.
3. Write the total runtime of shard `c` to `out[c]`, computing each shard in its own goroutine.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ShardDurations([]int{10, 20, 30, 40}, 2)
Output: [30 70]
```

**Example 2:**

```
Input:  ShardDurations([]int{10, 20, 30}, 2)
Output: [30 30]
```

**Example 3:**

```
Input:  ShardDurations([]int{10}, 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Disjoint slice windows** | `durations[start:end]` shares the backing array, but the windows never overlap and are only read. |

## Hint

The shard count is `(len(durations) + perShard - 1) / perShard`. Compute
`start`/`end` on the parent and pass the sub-slice into the goroutine.

## Validate

```bash
make verify
```
