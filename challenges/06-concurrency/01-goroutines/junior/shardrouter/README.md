# Shard Router

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A storage proxy decides which shard each key belongs to before issuing the
writes. Routing is a pure hash of the key modulo the shard count, and a batch of
keys is routed concurrently so the proxy can dispatch as soon as the batch is
planned.

## Task

Implement `ShardIDs` in [shardrouter.go](shardrouter.go) so that:

1. Return `nil` when `shards <= 0`.
2. Return a slice of shard IDs the same length as `keys`.
3. ID `i` is `h % shards`, where `h` starts at `0` and folds each byte of the key with `h = h*31 + int(b)`.
4. Route each key in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ShardIDs([]string{"a", "b"}, 4)
Output: [1 2]
```

**Example 2:**

```
Input:  ShardIDs([]string{"a", "b"}, 1)
Output: [0 0]
```

**Example 3:**

```
Input:  ShardIDs([]string{"a"}, 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Guard before you fan out** | The `shards` check happens on the parent, before a single goroutine is started. |

## Hint

Check `shards` first and return early — validating inside the goroutines would
mean starting work you already know is invalid.

## Validate

```bash
make verify
```
