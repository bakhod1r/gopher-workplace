# Stream Shard IDs

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The search-index rebuild job at a marketplace splits the corpus into `n`
shards. The coordinator publishes shard ids on a channel and a pool of
indexers ranges over it until the coordinator signals there is no more work.

## Task

Implement `StreamShardIDs` in [shardids.go](shardids.go) so that:

1. It returns a receive-only channel carrying `0 .. n-1` in order.
2. The channel is **closed** once every id has been sent, so `range` terminates.
3. `n <= 0` yields a closed channel with no values.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  StreamShardIDs(3)
Output: 0, 1, 2 then closed
```

**Example 2:**

```
Input:  StreamShardIDs(1)
Output: 0 then closed
```

**Example 3:**

```
Input:  StreamShardIDs(0)
Output: closed immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`make(chan T, n)`** | A buffered channel of capacity `n` accepts `n` sends without a receiver. |
| 2 | **`close`** | Closing tells receivers no more values are coming; `range` stops. |
| 3 | **Directional return** | `<-chan int` lets indexers receive but never send or close. |

## Hint

Buffer the channel with capacity `n` so all sends complete before you
return, then `close` it.

## Validate

```bash
make verify
```
