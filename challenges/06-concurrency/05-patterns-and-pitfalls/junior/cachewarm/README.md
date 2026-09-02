# Cache Warm Broadcast

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

On deploy, every cache shard is warmed at once — but not before the new
configuration has been loaded. The warmers are started early and parked on a
`ready` channel; closing that channel releases all of them simultaneously,
which is the cheapest broadcast Go offers.

## Task

Implement `WarmShards` in [cachewarm.go](cachewarm.go) so that:

1. It starts one goroutine per shard, each blocking on `<-ready` before doing any work.
2. Once `ready` is closed, every warmer calls `warm(shard)` and records the result.
3. After `wg.Wait()` it returns the counts sorted ascending.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WarmShards(closed ready, []string{"a", "bb"}, keyCount)
Output: []int{1, 2}
```

**Example 2:**

```
Input:  WarmShards(closed ready, []string{"cccc", "b", "aa"}, keyCount)
Output: []int{1, 2, 4}
```

**Example 3:**

```
Input:  WarmShards(closed ready, nil, keyCount)
Output: []int{} (empty, length 0)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Close as broadcast** | One `close(ready)` releases every goroutine waiting on it. |
| 2 | **Starting gate** | Goroutines can be created long before they are allowed to run. |
| 3 | **WaitGroup** | The mirror image: `Wait` blocks until the last warmer reports done. |

## Hint

`<-ready` on an open channel blocks; closing the channel makes that receive
succeed immediately in *every* waiting goroutine at once.

## Validate

```bash
make verify
```
