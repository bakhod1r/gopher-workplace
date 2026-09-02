# Wait For Shards

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The search service will not accept traffic until every index shard has
finished warming up. Each shard's warm-up goroutine reports in on a signal
channel; no data is exchanged, only the fact of completion.

## Task

Implement `WaitForShards` in [shardwarmup.go](shardwarmup.go) so that:

1. It creates a `chan struct{}` used purely as a readiness signal.
2. It starts `n` goroutines, each sending one empty struct.
3. It receives exactly `n` signals and returns the count; `n <= 0` returns `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WaitForShards(3)
Output: 3
```

**Example 2:**

```
Input:  WaitForShards(1)
Output: 1
```

**Example 3:**

```
Input:  WaitForShards(0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`chan struct{}`** | A zero-width signal channel — carries no data, allocates nothing. |
| 2 | **Counting receives** | Receiving `n` times is a hand-rolled `WaitGroup`. |
| 3 | **Goroutine launch** | `go func(){...}()` inside a loop starts `n` warm-ups. |

## Hint

`struct{}{}` is the value of the empty struct type. Receive exactly `n`
times — one per shard you started.

## Validate

```bash
make verify
```
